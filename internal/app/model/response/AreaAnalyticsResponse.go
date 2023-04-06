package response

type AreaAnalytics struct {
	TotalQuantityAnimals int               `json:"totalQuantityAnimals"`
	TotalAnimalsArrived  int               `json:"totalAnimalsArrived"`
	TotalAnimalsGone     int               `json:"totalAnimalsGone"`
	AnimalsAnalytics     []AnimalAnalytics `json:"animalsAnalytics"`
}
