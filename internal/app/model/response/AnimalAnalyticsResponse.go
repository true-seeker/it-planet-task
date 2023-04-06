package response

type AnimalAnalytics struct {
	AnimalType      string `json:"animalType"`
	AnimalTypeId    int    `json:"animalTypeId"`
	QuantityAnimals int    `json:"quantityAnimals"`
	AnimalsArrived  int    `json:"animalsArrived"`
	AnimalsGone     int    `json:"animalsGone"`
}
