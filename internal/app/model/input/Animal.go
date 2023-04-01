package input

type Animal struct {
	AnimalTypeIds      []int    `json:"animalTypes"`
	Weight             *float32 `json:"weight"`
	Height             *float32 `json:"height"`
	Length             *float32 `json:"length"`
	Gender             *string  `json:"gender"`
	ChipperId          *int     `json:"chipperId"`
	ChippingLocationId *int     `json:"chippingLocationId"`
	LifeStatus         *string  `json:"lifeStatus"`
}

type AnimalTypeUpdate struct {
	OldTypeId *int `json:"oldTypeId"`
	NewTypeId *int `json:"newTypeId"`
}

type AnimalLocationPointUpdate struct {
	VisitedLocationPointId *int `json:"visitedLocationPointId"`
	LocationPointId        *int `json:"locationPointId"`
}

//TODO сделать инпуты на все сущности
