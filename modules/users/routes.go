package users

import (
	"github.com/kataras/iris/v12"
)

const name = "user"

// Routes init user
func Routes(routes iris.Party) {
	route := routes.Party(name)
	route.Get("/", getUsers)
	route.Post("/", createUser)
	route.Post("/login", loginUser)
}
