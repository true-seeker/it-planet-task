package response

type Area struct {
	Id         int         `json:"id"`
	Name       string      `json:"name"`
	AreaPoints []AreaPoint `json:"areaPoints"`
}
