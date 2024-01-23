package routes

import (
    "github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
)

func Web() {
    facades.Route().Get("/", func(ctx http.Context) http.Response {
		return ctx.Response().Redirect(http.StatusMovedPermanently, "weather")
	})

	weatherController := controllers.NewWeatherController()
	facades.Route().Get("weather", weatherController.Index)
	facades.Route().Get("weather/{location}", weatherController.Show)
}
