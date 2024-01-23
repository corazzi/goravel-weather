package commands

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"fmt"
)

type GetWeatherInformation struct {
}

//Signature The name and signature of the console command.
func (receiver *GetWeatherInformation) Signature() string {
	return "weather"
}

//Description The console command description.
func (receiver *GetWeatherInformation) Description() string {
	return "What's the weather like?"
}

//Extend The console command extend.
func (receiver *GetWeatherInformation) Extend() command.Extend {
	return command.Extend{}
}

//Handle Execute the console command.
func (receiver *GetWeatherInformation) Handle(ctx console.Context) error {
	location := ctx.Argument(0)

	fmt.Printf("Looking up the weather in %s...", location)

	return nil
}
