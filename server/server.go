package server

import (
	"to-do-go/config"
	"to-do-go/modules/users"

	"github.com/kataras/iris/v12"
)

//Routes Initialization
func SetupRoutes(app *iris.Application) {
	prefix := app.Party(config.StringEnvVariable("PREFIX"))
	//setupRoutes
	users.Routes(prefix)
}
