package geometry

import (
	"it-planet-task/internal/app/model/entity"
	"math"
)

type Geometry interface {
	IsPointOnLineSegment(l *LineSegment, c *entity.AreaPoint) bool
	IsIntersects(l *LineSegment, l2 *LineSegment) bool
	IsPointInsideArea(pt *entity.AreaPoint, pg *entity.Area, isOnEdgeCounts bool) bool
	IsAllPointsOnOneLine(points *[]entity.AreaPoint) bool
}

type GeometryService struct {
}

func NewGeometryService() Geometry {
	return &GeometryService{}
}

type LineSegment struct {
	P entity.AreaPoint
	Q entity.AreaPoint
}

func NewLineSegment(p entity.AreaPoint, q entity.AreaPoint) *LineSegment {
	return &LineSegment{P: p, Q: q}
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

func (g *GeometryService) IsPointOnLineSegment(l *LineSegment, c *entity.AreaPoint) bool {
	crossProduct := (*c.Longitude-*l.P.Longitude)*(*l.Q.Latitude-*l.P.Latitude) - (*c.Latitude-*l.P.Latitude)*(*l.Q.Longitude-*l.P.Longitude)
	if math.Abs(crossProduct) > 0 {
		return false
	}

	dotProduct := (*c.Latitude-*l.P.Latitude)*(*l.Q.Latitude-*l.P.Latitude) + (*c.Longitude-*l.P.Longitude)*(*l.Q.Longitude-*l.P.Longitude)
	if dotProduct < 0 {
		return false
	}

	squaredLength := (*l.Q.Latitude-*l.P.Latitude)*(*l.Q.Latitude-*l.P.Latitude) + (*l.Q.Longitude-*l.P.Longitude)*(*l.Q.Longitude-*l.P.Longitude)
	if dotProduct > squaredLength {
		return false
	}

	return true
}

func (g *GeometryService) IsIntersects(l *LineSegment, l2 *LineSegment) bool {
	if g.IsPointOnLineSegment(l, &l2.P) || g.IsPointOnLineSegment(l, &l2.Q) || g.IsPointOnLineSegment(l2, &l.P) || g.IsPointOnLineSegment(l2, &l.Q) {
		return false
	}

	//if (l.P.IsEqual(&l2.P) && !l.Q.IsEqual(&l2.Q)) ||
	//	(l.Q.IsEqual(&l2.Q) && !l.P.IsEqual(&l2.P)) ||
	//	(l.Q.IsEqual(&l2.P) && !l.P.IsEqual(&l2.Q)) ||
	//	(l.P.IsEqual(&l2.Q) && !l.Q.IsEqual(&l2.Q)) {
	//	return false
	//}
	// Find the four orientations needed for general and
	// special cases
	o1 := orientation(l.P, l.Q, l2.P)
	o2 := orientation(l.P, l.Q, l2.Q)
	o3 := orientation(l2.P, l2.Q, l.P)
	o4 := orientation(l2.P, l2.Q, l.Q)

	// General case
	if o1 != o2 && o3 != o4 {
		return true
	}

	//// Special Cases
	//// p1, q1 and p2 are collinear and p2 lies on segment p1q1
	//if o1 == 0 && onSegment(l.P, l2.P, l.Q) {
	//	return true
	//}
	//
	//// p1, q1 and q2 are collinear and q2 lies on segment p1q1
	//if o2 == 0 && onSegment(l.P, l2.Q, l.Q) {
	//	return true
	//}
	//
	//// p2, q2 and p1 are collinear and p1 lies on segment p2q2
	//if o3 == 0 && onSegment(l2.P, l.P, l2.Q) {
	//	return true
	//}
	//
	//// p2, q2 and q1 are collinear and q1 lies on segment p2q2
	//if o4 == 0 && onSegment(l2.P, l.Q, l2.Q) {
	//	return true
	//}

	return false
}

func (g *GeometryService) IsPointInsideArea(pt *entity.AreaPoint, pg *entity.Area, isOnEdgeCounts bool) bool {
	if len(pg.AreaPoints) == 0 {
		return false
	}

	in := rayIntersectsSegment(*pt, pg.AreaPoints[len(pg.AreaPoints)-1], pg.AreaPoints[0])
	for i := 1; i < len(pg.AreaPoints); i++ {
		ls := NewLineSegment(pg.AreaPoints[i-1], pg.AreaPoints[i])
		if isOnEdgeCounts {
			if g.IsPointOnLineSegment(ls, pt) {
				return true
			}
		} else {
			if g.IsPointOnLineSegment(ls, pt) {
				return false
			}
		}
		if rayIntersectsSegment(*pt, pg.AreaPoints[i-1], pg.AreaPoints[i]) {
			in = !in
		}
	}
	return in

}

func rayIntersectsSegment(p, a, b entity.AreaPoint) bool {
	return (*a.Longitude > *p.Longitude) != (*b.Longitude > *p.Longitude) &&
		*p.Latitude < (*b.Latitude-*a.Latitude)*(*p.Longitude-*a.Longitude)/(*b.Longitude-*a.Longitude)+*a.Latitude
}

func (g *GeometryService) IsAllPointsOnOneLine(points *[]entity.AreaPoint) bool {
	dx := *(*points)[0].Latitude - *(*points)[1].Latitude
	dy := *(*points)[0].Longitude - *(*points)[1].Longitude
	for i := 2; i < len(*points); i++ {
		dx1 := *(*points)[0].Latitude - *(*points)[i].Latitude
		dy1 := *(*points)[0].Longitude - *(*points)[i].Longitude
		if dx1*dy != dy1*dx {
			return false
		}
	}
	return true
}
