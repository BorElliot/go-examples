package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GeoIP struct {
	IP            string  `json:"ip"`
	ContinentName string  `json:"continent_name"`
	CountryName   string  `json:"country_name"`
	RegionName    string  `json:"region_name"`
	Latitude      float32 `json:"latitude"`
	Longitude     float32 `json:"longitude"`
}

var (
	address  string = "54.219.132.112"
	err      error
	geo      GeoIP
	response *http.Response
	body     []byte
)

func main() {

	response, err = http.Get("https://ipstack.com/ipstack_api.php?ip=" + address)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	// response.Body() is a reader type. We have
	// to  use ioutil.ReadAll() to read the data
	// in to a byte slice(string)
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	// Unmarshal the JSON byte slice to a GeoIP struct
	err = json.Unmarshal(body, &geo)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\n===== IP Geolocation Info ====\n")
	fmt.Println("IP address:\t", geo.IP)
	fmt.Println("Continent name:\t", geo.ContinentName)
	fmt.Println("Country name:\t", geo.CountryName)
	fmt.Println("Region name:\t", geo.RegionName)
	fmt.Println("Latitude:\t", geo.Latitude)
	fmt.Println("Longitude:\t", geo.Longitude)
}
