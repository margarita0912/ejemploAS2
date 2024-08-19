package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"test-apicall/test-apicall/dto"
)

func main() {
	log.Info("starting class 19/08/24")

	args := os.Args
	itemIdParam := args[1]

	resp, err := http.Get(fmt.Sprintf("http://api.mercadolibre.com/items/%s", itemIdParam))
	if err != nil {
		fmt.Println("error making http get request error: ", err)
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println("status code:", resp.StatusCode)

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response data: ", err)
		log.Fatal(err)
	}
	var itemDto dto.ItemDto
	json.Unmarshal(responseData, &itemDto)

	fmt.Println("Id: " + itemDto.Id)
	fmt.Println("Site: " + itemDto.SiteId)
	fmt.Println("Titulo: " + itemDto.Title)

	respC, err := http.Get(fmt.Sprintf("http://api.mercadolibre.com/countries/%s", dto.UserDto{CountryId}))
	if err != nil {
		fmt.Println("error making http get request error: ", err)
		log.Fatal(err)
	}

	var countryDto dto.CountryDto
	json.Unmarshal(responseData, &itemDto)

	fmt.Println("Id: " + countryDto.Id)
	fmt.Println("Name: " + countryDto.Name)
	fmt.Println("Latitude: " + fmt.Sprintf("%f", countryDto.GeoInformation.Location.Latitude))
	fmt.Println("Longitude: " + fmt.Sprintf("%f", countryDto.GeoInformation.Location.Longitude))

	log.Info("---------------------------------------")

	log.Info("finishing class 19/08/24")
}
