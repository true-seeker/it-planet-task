package test

import (
	"it-planet-task/internal/app/model/entity"
	"it-planet-task/internal/app/service"
	"testing"
)

func makeFloatPtr(a float64) *float64 {
	return &a
}

func TestIntersectingLineSegments(t *testing.T) {
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

	ls1 := service.NewLineSegment(p1, q1)
	ls2 := service.NewLineSegment(p2, q2)

	got := ls1.IsIntersects(ls2)
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestNonIntersectingLineSegments(t *testing.T) {
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

	ls1 := service.NewLineSegment(p1, q1)
	ls2 := service.NewLineSegment(p2, q2)

	got := ls1.IsIntersects(ls2)
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestParallelIntersectingLineSegments(t *testing.T) {
	p1 := entity.AreaPoint{
		Latitude:  makeFloatPtr(0),
		Longitude: makeFloatPtr(1.5),
	}
	q1 := entity.AreaPoint{
		Latitude:  makeFloatPtr(1.5),
		Longitude: makeFloatPtr(2.5),
	}
	p2 := entity.AreaPoint{
		Latitude:  makeFloatPtr(0),
		Longitude: makeFloatPtr(1.5),
	}
	q2 := entity.AreaPoint{
		Latitude:  makeFloatPtr(1.5),
		Longitude: makeFloatPtr(2.5),
	}

	ls1 := service.NewLineSegment(p1, q1)
	ls2 := service.NewLineSegment(p2, q2)

	got := ls1.IsIntersects(ls2)
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestTouchingIntersectingLineSegments(t *testing.T) {
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

	ls1 := service.NewLineSegment(p1, q1)
	ls2 := service.NewLineSegment(p2, q2)

	got := ls1.IsIntersects(ls2)
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestPointInsideConvexPolygon(t *testing.T) {
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
		Latitude:  makeFloatPtr(1.5),
		Longitude: makeFloatPtr(2),
	}
	lss := make([]service.LineSegment, 0)

	lss = append(lss, *service.NewLineSegment(p1, q1))
	lss = append(lss, *service.NewLineSegment(q1, p2))
	lss = append(lss, *service.NewLineSegment(p2, q2))
	lss = append(lss, *service.NewLineSegment(p2, q1))

	got := service.IsPointInsideArea(&point, &lss)
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestPointInsideConveyPolygon(t *testing.T) {
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

	point := entity.AreaPoint{
		Latitude:  makeFloatPtr(1.5),
		Longitude: makeFloatPtr(2.5),
	}
	lss := make([]service.LineSegment, 0)

	lss = append(lss, *service.NewLineSegment(p1, q1))
	lss = append(lss, *service.NewLineSegment(q1, p2))
	lss = append(lss, *service.NewLineSegment(p2, q2))
	lss = append(lss, *service.NewLineSegment(p2, q1))

	got := service.IsPointInsideArea(&point, &lss)
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestPointOutsideConvexPolygon(t *testing.T) {
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
	lss := make([]service.LineSegment, 0)

	lss = append(lss, *service.NewLineSegment(p1, q1))
	lss = append(lss, *service.NewLineSegment(q1, p2))
	lss = append(lss, *service.NewLineSegment(p2, q2))
	lss = append(lss, *service.NewLineSegment(p2, q1))

	got := service.IsPointInsideArea(&point, &lss)
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestPointOnBorderConvexPolygon(t *testing.T) {
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
	lss := make([]service.LineSegment, 0)

	lss = append(lss, *service.NewLineSegment(p1, q1))
	lss = append(lss, *service.NewLineSegment(q1, p2))
	lss = append(lss, *service.NewLineSegment(p2, q2))
	lss = append(lss, *service.NewLineSegment(p2, q1))

	got := service.IsPointInsideArea(&point, &lss)
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
