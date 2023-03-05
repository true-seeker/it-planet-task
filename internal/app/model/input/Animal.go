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
