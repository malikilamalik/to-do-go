package users

import (
	"github.com/kataras/iris/v12"
)

// Routes init user
func Routes(app *iris.Application) {
	app.Get("/", GetUsers)
}
