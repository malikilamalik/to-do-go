package main

import (
	"to-do-go/modules/users"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	users.Routes(app)
	app.Listen(":1997")
}
