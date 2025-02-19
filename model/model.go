package model

type Data struct {
	Id           string      `json:"id"`
	Avg          string      `json:"avg"`
	Date         string      `json:"date"`
	DateStandard string      `json:"dateStandard"`
	Value        float64     `json:"value"`
	Unit         string      `json:"unit"`
	Address      string      `json:"address"`
	Coordinates  Coordinates `json:"coordinates"`
}
type Coordinates struct {
	Lat float64 `json:"lat,omitempty"`
	Lng float64 `json:"lng,omitempty"`
}
