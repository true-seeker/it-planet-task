package response

type AreaAnalytics struct {
	TotalQuantityAnimals int
	TotalAnimalsArrived  int
	TotalAnimalsGone     int
	AnimalAnalytics      []AnimalAnalytics
}
