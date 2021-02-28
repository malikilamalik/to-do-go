package todos

import (
	"to-do-go/middlewares"

	"github.com/kataras/iris/v12"
)

const name = "todo"

// Routes init user
func Routes(routes iris.Party) {
	route := routes.Party(name)
	route.Use(middlewares.Authorization())
	route.Get("/", getTasks)
	route.Get("/task", getTaskByTaskId)
	route.Get("/status", getTaskByTaskStatus)
	route.Post("/new", createTask)
	route.Delete("/delete", deleteTask)
	route.Put("/status/finish", taskFinish)
	route.Put("/status/ongoing", taskOngoing)
}
