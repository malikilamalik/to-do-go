package main

import (
	"os"
	"to-do-go/modules/auth"
	"to-do-go/modules/todos"
	"to-do-go/modules/users"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	initRoutes(app)
	app.Listen(":1997")
}

//Routes Initialization
func initRoutes(app *iris.Application) {
	prefix := app.Party(os.Getenv("PREFIX"))

	//setupRoutes
	users.Routes(prefix)
	todos.Routes(prefix)
	auth.Routes(prefix)
}

//
