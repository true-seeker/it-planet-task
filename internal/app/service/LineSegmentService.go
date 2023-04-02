package service

import (
	"it-planet-task/internal/app/model/entity"
	"math"
)

type LineSegment struct {
	p entity.AreaPoint
	q entity.AreaPoint
}

func NewLineSegment(p entity.AreaPoint, q entity.AreaPoint) *LineSegment {
	return &LineSegment{p: p, q: q}
}

func onSegment(p entity.AreaPoint, q entity.AreaPoint, r entity.AreaPoint) bool {
	if *q.Latitude <= math.Max(*p.Latitude, *r.Latitude) && *q.Latitude >= math.Min(*p.Latitude, *r.Latitude) &&
		*q.Longitude <= math.Max(*p.Longitude, *r.Longitude) && *q.Longitude >= math.Min(*p.Longitude, *r.Longitude) {
		return true
	}
	return false
}

func orientation(p entity.AreaPoint, q entity.AreaPoint, r entity.AreaPoint) int {
	val := (*q.Longitude-*p.Longitude)*(*r.Latitude-*q.Latitude) - (*q.Latitude-*p.Latitude)*(*r.Longitude-*q.Longitude)

	if val > 0 {
		return 1 // Clockwise
	} else if val < 0 {
		return 2 // Counterclockwise
	} else {
		return 0 //Collinear
	}
}

// IsIntersects https://www.geeksforgeeks.org/check-if-two-given-line-segments-intersect/
func (l *LineSegment) IsIntersects(l2 *LineSegment) bool {
	// Find the four orientations needed for general and
	// special cases
	o1 := orientation(l.p, l.q, l2.p)
	o2 := orientation(l.p, l.q, l2.q)
	o3 := orientation(l2.p, l2.q, l.p)
	o4 := orientation(l2.p, l2.q, l.q)

	// General case
	if o1 != o2 && o3 != o4 {
		return true
	}

	// Special Cases
	// p1, q1 and p2 are collinear and p2 lies on segment p1q1
	if o1 == 0 && onSegment(l.p, l2.p, l.q) {
		return true
	}

	// p1, q1 and q2 are collinear and q2 lies on segment p1q1
	if o2 == 0 && onSegment(l.p, l2.q, l.q) {
		return true
	}

	// p2, q2 and p1 are collinear and p1 lies on segment p2q2
	if o3 == 0 && onSegment(l2.p, l.p, l2.q) {
		return true
	}

	// p2, q2 and q1 are collinear and q1 lies on segment p2q2
	if o4 == 0 && onSegment(l2.p, l.q, l2.q) {
		return true
	}

	return false // Doesn't fall in any of the above cases
}

func IsPointInsideArea(point *entity.AreaPoint, lineSegments *[]LineSegment) (i bool) {
	for _, side := range *lineSegments {
		if rayIntersectsSegment(*point, &side) {
			i = !i
		}
	}
	return i
}

// https://rosettacode.org/wiki/Ray-casting_algorithm#Go
func rayIntersectsSegment(p entity.AreaPoint, s *LineSegment) bool {
	var a, b entity.AreaPoint
	if *s.p.Longitude < *s.q.Longitude {
		a, b = s.p, s.q
	} else {
		a, b = s.q, s.p
	}
	for p.Longitude == a.Longitude || p.Longitude == b.Longitude {
		longitude := math.Nextafter(*p.Longitude, math.Inf(1))
		p.Longitude = &longitude
	}
	if *p.Longitude < *a.Longitude || *p.Longitude > *b.Longitude {
		return false
	}
	if *a.Latitude > *b.Latitude {
		if *p.Latitude > *a.Latitude {
			return false
		}
		if *p.Latitude < *b.Latitude {
			return true
		}
	} else {
		if *p.Latitude > *b.Latitude {
			return false
		}
		if *p.Latitude < *a.Latitude {
			return true
		}
	}
	return (*p.Longitude-*a.Longitude)/(*p.Latitude-*a.Latitude) >= (*b.Longitude-*a.Longitude)/(*b.Latitude-*a.Latitude)
}
