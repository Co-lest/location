package status

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Location struct {
    IP       string 
    City     string
    Region   string
    Country  string
    Loc      string
    Latitude  float32
    Longitude float32
}

func GetLocation() (Location, error) {
    var LocationStruct Location

    resp, err := http.Get("https://ipinfo.io/json")
    if err != nil {
        return LocationStruct, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return LocationStruct, err
    }

    err = json.Unmarshal(body, &LocationStruct)
    if err != nil {
        return LocationStruct, err
    }

    fmt.Sscanf(LocationStruct.Loc, "%f,%f", &LocationStruct.Latitude, &LocationStruct.Longitude)

    return LocationStruct, nil
}