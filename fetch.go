package status

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type WeatherData struct {	
	Hourly struct {
		Temperature2m []float32 `json:"temperature_2m"`
		Humidity []float32 `json:"relative_humidity_2m"`
		Cloud_cover []float32 `json:"cloud_cover"`
		Wind_speed []float32 `json:"wind_speed_10m"`
		Shortwave_radiation []float32 `json:"shortwave_radiation"`
	} `json:"hourly"`
}

func FetchWeather(baseurl string, latitude, longitude float32) (WeatherData, error) {
	var weatherData WeatherData

	latString := fmt.Sprintf("%.2f", latitude)
	longString := fmt.Sprintf("%.2f", longitude)

	finishedurl, err := ParseUrl(baseurl, latString, longString)
	if err != nil {
		return weatherData, nil
	}

	resp, err := http.Get(finishedurl)
	if err != nil {
		return weatherData, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Non-OK HTTP status:", resp.StatusCode)
		panic("Failed to fetch API!")
	}
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return weatherData, err
	}

	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		panic(err)
	}

	return weatherData, nil
}

func ParseUrl(baseurl, latitude, longitude string) (string, error) {
	urlValues := "&hourly=temperature_2m,relative_humidity_2m,cloud_cover,wind_speed_10m,shortwave_radiation"

	//var urlValues = []string {"temperature_2m", "relative_humidity_2m", "cloud_cover,wind_speed_10m", "shortwave_radiation"}

	parsedUrl, err := url.Parse(baseurl)
	if err != nil {
		fmt.Println("Failed to parse url:", err)
		return "", err
	}

	//parsedUrl := strings.Join(urlValues, ",")

	query := url.Values{}
	query.Add("latitude", latitude)
	query.Add("longitude", longitude)
	// query.Add("hourly", urlValues)
	parsedUrl.RawQuery = query.Encode()

	// fmt.Println("Final url:", parsedUrl.String()) 

	//decided not to use the parse methods because of (hourly=temperature_2m%2Crelative_humidity_2m) the %2C addition which is not recommended

	finUrl := parsedUrl.String() + urlValues

	return finUrl, nil
}