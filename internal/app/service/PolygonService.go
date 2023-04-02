package service

//
//import "it-planet-task/internal/app/model/entity"
//
//type PolygonFig struct {
//	points *[]entity.AreaPoint
//}
//
//func NewPolygon(points *[]entity.AreaPoint) *PolygonFig {
//	return &PolygonFig{points: points}
//}
//
//func (p *PolygonFig) IsConvex() bool {
//	n := len(*p.points)
//	points := *p.points
//	flag := 0
//	for i, point := range *p.points {
//		j := (i + 1) % n
//		k := (i + 2) % n
//		z := (*points[j].Latitude - *point.Latitude) * (*points[k].Longitude - *points[j].Longitude)
//		z -= (*points[j].Longitude - *point.Longitude) * (*points[k].Latitude - *points[j].Latitude)
//
//		if z < 0 {
//			flag |= 1
//		} else if z > 0 {
//			flag |= 2
//		}
//		if flag == 3 {
//			return false
//		}
//	}
//	if flag != 0 {
//		return true
//	}
//	return false
//}
//
//func (p *PolygonFig) IsIntersect(p2 *PolygonFig) bool {
//
//}
