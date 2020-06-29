package model

// List holds technical fields
type List struct {
	Count int      `json:"count"`
	Data  []Object `json:"data"`
}
