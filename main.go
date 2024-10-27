package main

import (
	"fmt"
	"loc/status"
	"sync"
)

var once sync.Once

func main() {
	var weatherDataChan = make(chan struct {
		status.WeatherData
		error
	  })

	location, err := status.GetLocation()
    if err != nil {
        fmt.Println("Error fetching location:", err)
        return
    }

	url1 := "https://api.open-meteo.com/v1/forecast?"

	// var err1 error
	// var weatherData status.WeatherData

	go func() {
		once.Do(func() {
		  weatherData, err := status.FetchWeather(url1, location.Latitude, location.Longitude)
		  weatherDataChan <- struct {
			status.WeatherData
			error
		  }{weatherData, err}
		})
	  }()
	  
	  dataAndError := <-weatherDataChan
	  weatherData, err1 := dataAndError.WeatherData, dataAndError.error

	//weatherData, err1 = status.FetchWeather(url1, location.Latitude, location.Longitude)

	if err1 != nil {
		fmt.Println("Error fetching data:", err1)
		return
	}

	fmt.Printf("IP: %s\nCity: %s\nRegion: %s\nCountry: %s\nLatitude: %.4f\nLongitude: %.4f\n",
        location.IP, location.City, location.Region, location.Country, location.Latitude, location.Longitude)

	fmt.Println("Temperature:", weatherData.Hourly.Temperature2m[0], "°C")
	fmt.Println("Humidity:", weatherData.Hourly.Humidity[0], "%")
	fmt.Println("Cloud cover:", weatherData.Hourly.Cloud_cover[0], "%")
	fmt.Println("Wind speed:", weatherData.Hourly.Wind_speed[0] ,"km/h")
	fmt.Println("Shortwave radistion:", weatherData.Hourly.Shortwave_radiation[0], "W/m²")
}
