package controllers

import (
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/support"

	"goravel/app/services"
	"fmt"
	"math/rand"
	"time"
)

type WeatherController struct {
	//Dependent services
}

func NewWeatherController() *WeatherController {
	return &WeatherController{
		//Inject services
	}
}

func getRandomLocation() string {
    sampleLocations := []string{"Hull", "Edinburgh", "London", "Scarborough", "Wolverhampton", "Dublin"}

    rand.Seed(time.Now().UnixNano())
    randomIndex := rand.Intn(len(sampleLocations))
    return sampleLocations[randomIndex]
}

func (r *WeatherController) Index(ctx http.Context) http.Response {
    location := ctx.Request().Query("location")

    if (location != "") {
        return ctx.Response().Redirect(http.StatusMovedPermanently, fmt.Sprintf("/weather/%s", location))
    }

    return ctx.Response().View().Make("weather.tmpl", map[string]any{
        "version": support.Version,
        "location": getRandomLocation(),
    })
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

	weatherDetail := fmt.Sprintf("The temperature in %s is %.2f°C – feels like %.2f°C", weather.Location, weather.Main.Temperature, weather.Main.FeelsLike)

	return ctx.Response().View().Make("weather.tmpl", map[string]any{
        "version": support.Version,
        "location": location,
        "weatherDetail": weatherDetail,
    })

    /*
	return ctx.Response().String(
	    http.StatusOK,
	    weatherDetail,
	)
	*/
}
