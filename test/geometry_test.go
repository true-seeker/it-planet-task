package test

import (
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/service/geometry"
	"testing"
)

func makeFloatPtr(a float64) *float64 {
	return &a
}
func TestIntersectingLineSegments(t *testing.T) {
	geometryService := geometry.NewGeometryService()
	p1 := entity.AreaPoint{
		Latitude:  makeFloatPtr(0),
		Longitude: makeFloatPtr(1.5),
	}
	q1 := entity.AreaPoint{
		Latitude:  makeFloatPtr(1.5),
		Longitude: makeFloatPtr(2.5),
	}
	p2 := entity.AreaPoint{
		Latitude:  makeFloatPtr(1),
		Longitude: makeFloatPtr(1.5),
	}
	q2 := entity.AreaPoint{
		Latitude:  makeFloatPtr(0.5),
		Longitude: makeFloatPtr(3),
	}

	ls1 := geometry.NewLineSegment(p1, q1)
	ls2 := geometry.NewLineSegment(p2, q2)

	got := geometryService.IsIntersects(ls1, ls2)
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestIntersectingLineSegments2(t *testing.T) {
	geometryService := geometry.NewGeometryService()
	p1 := entity.AreaPoint{
		Latitude:  makeFloatPtr(1),
		Longitude: makeFloatPtr(-179.0),
	}
	q1 := entity.AreaPoint{
		Latitude:  makeFloatPtr(1),
		Longitude: makeFloatPtr(-166.0),
	}
	p2 := entity.AreaPoint{
		Latitude:  makeFloatPtr(7),
		Longitude: makeFloatPtr(-166.0),
	}
	q2 := entity.AreaPoint{
		Latitude:  makeFloatPtr(7),
		Longitude: makeFloatPtr(-180),
	}
	lss := make([]geometry.LineSegment, 0)

	lss = append(lss, *geometry.NewLineSegment(p1, q1))
	lss = append(lss, *geometry.NewLineSegment(q1, p2))
	lss = append(lss, *geometry.NewLineSegment(p2, q2))
	lss = append(lss, *geometry.NewLineSegment(q2, p1))

	for i := 0; i < len(lss); i++ {
		for j := i + 1; j < len(lss); j++ {
			if geometryService.IsIntersects(&lss[j], &lss[i]) {
				t.Errorf("got %t, wanted %t", true, false)
				return
			}
		}
	}
}

func TestNonIntersectingLineSegments(t *testing.T) {
	geometryService := geometry.NewGeometryService()
	p1 := entity.AreaPoint{
		Latitude:  makeFloatPtr(0),
		Longitude: makeFloatPtr(1.5),
	}
	q1 := entity.AreaPoint{
		Latitude:  makeFloatPtr(1.5),
		Longitude: makeFloatPtr(2.5),
	}
	p2 := entity.AreaPoint{
		Latitude:  makeFloatPtr(1),
		Longitude: makeFloatPtr(1.5),
	}
	q2 := entity.AreaPoint{
		Latitude:  makeFloatPtr(2),
		Longitude: makeFloatPtr(2),
	}

	ls1 := geometry.NewLineSegment(p1, q1)
	ls2 := geometry.NewLineSegment(p2, q2)

	got := geometryService.IsIntersects(ls2, ls1)
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestTouchingMiddleIntersectingLineSegments(t *testing.T) {
	geometryService := geometry.NewGeometryService()
	p1 := entity.AreaPoint{
		Latitude:  makeFloatPtr(0),
		Longitude: makeFloatPtr(0),
	}
	q1 := entity.AreaPoint{
		Latitude:  makeFloatPtr(0),
		Longitude: makeFloatPtr(2),
	}
	p2 := entity.AreaPoint{
		Latitude:  makeFloatPtr(0),
		Longitude: makeFloatPtr(1),
	}
	q2 := entity.AreaPoint{
		Latitude:  makeFloatPtr(1),
		Longitude: makeFloatPtr(1),
	}

	ls1 := geometry.NewLineSegment(p1, q1)
	ls2 := geometry.NewLineSegment(p2, q2)

	got := geometryService.IsIntersects(ls2, ls1)
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestContainingLineSegments(t *testing.T) {
	geometryService := geometry.NewGeometryService()
	p1 := entity.AreaPoint{
		Latitude:  makeFloatPtr(0),
		Longitude: makeFloatPtr(0),
	}
	q1 := entity.AreaPoint{
		Latitude:  makeFloatPtr(0),
		Longitude: makeFloatPtr(5),
	}
	p2 := entity.AreaPoint{
		Latitude:  makeFloatPtr(0),
		Longitude: makeFloatPtr(1),
	}
	q2 := entity.AreaPoint{
		Latitude:  makeFloatPtr(0),
		Longitude: makeFloatPtr(4),
	}

	ls1 := geometry.NewLineSegment(p1, q1)
	ls2 := geometry.NewLineSegment(p2, q2)

	got := geometryService.IsIntersects(ls2, ls1)
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestTouchingIntersectingLineSegments(t *testing.T) {
	geometryService := geometry.NewGeometryService()
	p1 := entity.AreaPoint{
		Latitude:  makeFloatPtr(0),
		Longitude: makeFloatPtr(1.5),
	}
	q1 := entity.AreaPoint{
		Latitude:  makeFloatPtr(1),
		Longitude: makeFloatPtr(1.5),
	}
	p2 := entity.AreaPoint{
		Latitude:  makeFloatPtr(1),
		Longitude: makeFloatPtr(1.5),
	}
	q2 := entity.AreaPoint{
		Latitude:  makeFloatPtr(2),
		Longitude: makeFloatPtr(2),
	}

	ls1 := geometry.NewLineSegment(p1, q1)
	ls2 := geometry.NewLineSegment(p2, q2)

	got := geometryService.IsIntersects(ls2, ls1)
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestPointInsideConveyPolygon(t *testing.T) {
	geometryService := geometry.NewGeometryService()
	p1 := entity.AreaPoint{
		Latitude:  makeFloatPtr(0),
		Longitude: makeFloatPtr(1.5),
	}
	q1 := entity.AreaPoint{
		Latitude:  makeFloatPtr(1),
		Longitude: makeFloatPtr(1.5),
	}
	q2 := entity.AreaPoint{
		Latitude:  makeFloatPtr(2),
		Longitude: makeFloatPtr(2),
	}
	p2 := entity.AreaPoint{
		Latitude:  makeFloatPtr(1.5),
		Longitude: makeFloatPtr(3),
	}
	areaPoints := make([]entity.AreaPoint, 0)

	areaPoints = append(areaPoints, p1)
	areaPoints = append(areaPoints, q1)
	areaPoints = append(areaPoints, q2)
	areaPoints = append(areaPoints, p2)

	area := entity.Area{
		AreaPoints: areaPoints,
	}

	point := entity.AreaPoint{
		Latitude:  makeFloatPtr(1.5),
		Longitude: makeFloatPtr(2.5),
	}

	got := geometryService.IsPointInsideArea(&point, &area, false)
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestPointOutsideConvexPolygon(t *testing.T) {
	geometryService := geometry.NewGeometryService()
	p1 := entity.AreaPoint{
		Latitude:  makeFloatPtr(0),
		Longitude: makeFloatPtr(1.5),
	}
	q1 := entity.AreaPoint{
		Latitude:  makeFloatPtr(1),
		Longitude: makeFloatPtr(1.5),
	}
	p2 := entity.AreaPoint{
		Latitude:  makeFloatPtr(2),
		Longitude: makeFloatPtr(2),
	}
	q2 := entity.AreaPoint{
		Latitude:  makeFloatPtr(1.5),
		Longitude: makeFloatPtr(3),
	}

	point := entity.AreaPoint{
		Latitude:  makeFloatPtr(5),
		Longitude: makeFloatPtr(2),
	}

	areaPoints := make([]entity.AreaPoint, 0)
	areaPoints = append(areaPoints, p1)
	areaPoints = append(areaPoints, q1)
	areaPoints = append(areaPoints, q2)
	areaPoints = append(areaPoints, p2)

	area := entity.Area{
		AreaPoints: areaPoints,
	}

	got := geometryService.IsPointInsideArea(&point, &area, false)
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestPointOnBorderConvexPolygon(t *testing.T) {
	geometryService := geometry.NewGeometryService()
	p1 := entity.AreaPoint{
		Latitude:  makeFloatPtr(1),
		Longitude: makeFloatPtr(1),
	}
	q1 := entity.AreaPoint{
		Latitude:  makeFloatPtr(2),
		Longitude: makeFloatPtr(1),
	}
	p2 := entity.AreaPoint{
		Latitude:  makeFloatPtr(2),
		Longitude: makeFloatPtr(2),
	}
	q2 := entity.AreaPoint{
		Latitude:  makeFloatPtr(1),
		Longitude: makeFloatPtr(2),
	}

	point := entity.AreaPoint{
		Latitude:  makeFloatPtr(2),
		Longitude: makeFloatPtr(1.5),
	}
	areaPoints := make([]entity.AreaPoint, 0)
	areaPoints = append(areaPoints, p1)
	areaPoints = append(areaPoints, q1)
	areaPoints = append(areaPoints, q2)
	areaPoints = append(areaPoints, p2)

	area := entity.Area{
		AreaPoints: areaPoints,
	}

	got := geometryService.IsPointInsideArea(&point, &area, false)
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestOneLinePolygon(t *testing.T) {
	geometryService := geometry.NewGeometryService()
	points := make([]entity.AreaPoint, 0)
	points = append(points, entity.AreaPoint{
		Latitude:  makeFloatPtr(-29),
		Longitude: makeFloatPtr(-179),
	})
	points = append(points, entity.AreaPoint{
		Latitude:  makeFloatPtr(-29),
		Longitude: makeFloatPtr(-175.75),
	})
	points = append(points, entity.AreaPoint{
		Latitude:  makeFloatPtr(-29),
		Longitude: makeFloatPtr(-172.5),
	})
	points = append(points, entity.AreaPoint{
		Latitude:  makeFloatPtr(-29),
		Longitude: makeFloatPtr(-169.25),
	})

	got := geometryService.IsAllPointsOnOneLine(&points)
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestNotOneLinePolygon(t *testing.T) {
	geometryService := geometry.NewGeometryService()
	points := make([]entity.AreaPoint, 0)
	points = append(points, entity.AreaPoint{
		Latitude:  makeFloatPtr(-28),
		Longitude: makeFloatPtr(-179),
	})
	points = append(points, entity.AreaPoint{
		Latitude:  makeFloatPtr(-26),
		Longitude: makeFloatPtr(-175.75),
	})
	points = append(points, entity.AreaPoint{
		Latitude:  makeFloatPtr(-25),
		Longitude: makeFloatPtr(-172.5),
	})
	points = append(points, entity.AreaPoint{
		Latitude:  makeFloatPtr(-20),
		Longitude: makeFloatPtr(-169.25),
	})

	got := geometryService.IsAllPointsOnOneLine(&points)
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestPointOnLineSegment(t *testing.T) {
	geometryService := geometry.NewGeometryService()
	p1 := entity.AreaPoint{
		Latitude:  makeFloatPtr(1),
		Longitude: makeFloatPtr(1),
	}
	q1 := entity.AreaPoint{
		Latitude:  makeFloatPtr(2),
		Longitude: makeFloatPtr(1),
	}

	ls := geometry.LineSegment{
		P: p1,
		Q: q1,
	}

	point := entity.AreaPoint{
		Latitude:  makeFloatPtr(1.5),
		Longitude: makeFloatPtr(1),
	}

	got := geometryService.IsPointOnLineSegment(&ls, &point)
	want := true
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestPointNotOnLineSegment(t *testing.T) {
	geometryService := geometry.NewGeometryService()
	p1 := entity.AreaPoint{
		Latitude:  makeFloatPtr(1),
		Longitude: makeFloatPtr(1),
	}
	q1 := entity.AreaPoint{
		Latitude:  makeFloatPtr(2),
		Longitude: makeFloatPtr(1),
	}

	ls := geometry.LineSegment{
		P: p1,
		Q: q1,
	}

	point := entity.AreaPoint{
		Latitude:  makeFloatPtr(3),
		Longitude: makeFloatPtr(1.5),
	}

	got := geometryService.IsPointOnLineSegment(&ls, &point)
	want := false
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestPointOnLineSegment2(t *testing.T) {
	geometryService := geometry.NewGeometryService()
	p1 := entity.AreaPoint{
		Latitude:  makeFloatPtr(0),
		Longitude: makeFloatPtr(0),
	}
	q1 := entity.AreaPoint{
		Latitude:  makeFloatPtr(2),
		Longitude: makeFloatPtr(2),
	}

	ls := geometry.LineSegment{
		P: p1,
		Q: q1,
	}

	point := entity.AreaPoint{
		Latitude:  makeFloatPtr(1),
		Longitude: makeFloatPtr(1),
	}

	got := geometryService.IsPointOnLineSegment(&ls, &point)
	want := true
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
