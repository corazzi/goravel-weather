package controllers

import (
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/contracts/http"
	"goravel/app/services"
	"fmt"
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

    // @todo inject this
	weatherService := services.WeatherService{
	    ApiKey: facades.Config().Env("OPENWEATHERMAP_API_KEY").(string),
	}

	weather, err := weatherService.GetCurrentWeather(location)

	if err != nil {
	    return ctx.Response().String(500, err.Error())
	}

	return ctx.Response().String(
	    http.StatusOK,
	    fmt.Sprintf("The temperature in %s is %.2f°C – feels like %.2f°C", weather.Location, weather.Main.Temperature, weather.Main.FeelsLike),
	)
}
