package dto

type CountryDto struct {
	Id             string            `json:"id"`
	Name           string            `json:"name"`
	GeoInformation GeoInformationDto `json:"geo_information"`
}

type GeoInformationDto struct {
	Location LocationDto `json:"location"`
}

type LocationDto struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
