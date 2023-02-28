package response

type Location struct {
	Id        int      `json:"id"`
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
}
