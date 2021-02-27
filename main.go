package main

import (
	"to-do-go/config"
	"to-do-go/server"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	config.InitDatabase()
	server.SetupRoutes(app)
	app.Listen(":1997")
}
