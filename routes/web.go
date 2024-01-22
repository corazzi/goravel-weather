package routes

import (
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
)

func Web() {
	weatherController := controllers.NewWeatherController()
	facades.Route().Get("weather", weatherController.Index)
	facades.Route().Get("weather/{location}", weatherController.Show)
}
