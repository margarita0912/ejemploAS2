package dto

type UserDto struct {
	Id        string `json:"id"`
	CountryId string `json:"country_id"`
	Name      string `json:"nickname"`
}
