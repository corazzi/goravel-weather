package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type WeatherController struct {
	//Dependent services
}

func NewWeatherController() *WeatherController {
	return &WeatherController{
		//Inject services
	}
}

func (r *WeatherController) Index(ctx http.Context) http.Response {
	return nil
}

func (r *WeatherController) Show(ctx http.Context) http.Response {
	location := ctx.Request().Route("location")

	// get the weather report

	return ctx.Response().Success().Json(http.Json{
        "location": location,
    })
}
