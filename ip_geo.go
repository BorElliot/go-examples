/*
* References
* 1. https://medium.com/@xcoulon/nested-structs-in-golang-2c750403a007
* 2. https://www.devdungeon.com/content/ip-geolocation-go
 */
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

type GeoIP struct {
	IP            string  `json:"ip"`
	ContinentName string  `json:"continent_name"`
	CountryName   string  `json:"country_name"`
	RegionName    string  `json:"region_name"`
	Latitude      float32 `json:"latitude"`
	Longitude     float32 `json:"longitude"`
	Location      struct {
		Capital          string `json:"capital"`
		CountryFlagEmoji string `json:"country_flag_emoji"`
		Languages        []struct {
			Code   string `json:"code"`
			Name   string `json:"name"`
			Native string `json:"native"`
		} `json:"languages"`
	} `json:"location"`
	TimeZone struct {
		ID               string `json:"id"`
		CurrentTime      string `json:"current_time"`
		GmtOffset        int32  `json:"gmt_offset"`
		Code             string `json:"code"`
		IsDaylightSaving bool   `json:"is_daylight_saving"`
	} `json:"time_zone"`
}

var (
	address  string = "54.219.132.112"
	err      error
	geo      GeoIP
	response *http.Response
	body     []byte
)

func main() {

	if len(os.Args) > 1 {
		address = os.Args[1]
	} else {
		log.Fatal("please input IP address")
	}

	if nil == net.ParseIP(address) {
		log.Fatal("The ip address is not correct")
	}

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

	fmt.Println("\n===== IP Geolocation Info ====")
	fmt.Println("IP address:\t", geo.IP)
	fmt.Println("Continent name:\t", geo.ContinentName)
	fmt.Println("Country name:\t", geo.CountryName)
	fmt.Println("Region name:\t", geo.RegionName)
	fmt.Println("Latitude:\t", geo.Latitude)
	fmt.Println("Longitude:\t", geo.Longitude)
	fmt.Println("Capital:\t", geo.Location.Capital)

	fmt.Println("\n==== Location Info ====")
	fmt.Println("Location Capital:\t", geo.Location.Capital)
	fmt.Println("Location Country Flag emoji:\t", geo.Location.CountryFlagEmoji)

	for _, v := range geo.Location.Languages {
		fmt.Println("Location language name:\t", v.Name)
	}

	fmt.Println("\n==== TimeZone Info ====")
	fmt.Println("ID:\t", geo.TimeZone.ID)
	fmt.Println("CurrentTime:\t", geo.TimeZone.CurrentTime)
	fmt.Println("Is daylight saving:\t", geo.TimeZone.IsDaylightSaving)
}
