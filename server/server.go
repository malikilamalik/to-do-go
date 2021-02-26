package server

import (
	"to-do-go/modules/users"

	"github.com/kataras/iris/v12"
)

//Routes Initialization
func SetupRoutes(app *iris.Application) {

	//setupRoutes
	users.Routes(app)
}
