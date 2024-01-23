package services

import (
    "encoding/json"
    "fmt"
    "net/http"
    "io/ioutil"
)

type WeatherResponse struct {
    Main struct {
        Temperature float64 `json:"temp"`
        FeelsLike float64 `json:"feels_like"`
    } `json:"main"`

    Location string `json:"name"`
}

type WeatherService struct {
    ApiKey string
}

func (service WeatherService) GetCurrentWeather(location string) (*WeatherResponse, error) {
    endpoint := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", location, service.ApiKey)

    response, e := http.Get(endpoint)

    if e != nil {
        return nil, e
    }

    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("Error response %s from API", response.StatusCode)
    }

    var weatherResponse WeatherResponse
    body, e := ioutil.ReadAll(response.Body)

    if e != nil {
        return nil, e
    }

    e = json.Unmarshal(body, &weatherResponse)

    if e != nil {
        return nil, e
    }

    return &weatherResponse, nil
}
