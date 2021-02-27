package server

import (
	"to-do-go/config"
	"to-do-go/modules/todos"
	"to-do-go/modules/users"

	"github.com/kataras/iris/v12"
)

func StartServer() {
	app := iris.New()
	config.InitDatabase()
	setupRoutes(app)
	app.Listen(":1997")
}

//Routes Initialization
func setupRoutes(app *iris.Application) {
	prefix := app.Party(config.StringEnvVariable("PREFIX"))
	//setupRoutes
	users.Routes(prefix)
	todos.Routes(prefix)
}
