package todos

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

const name = "todo"

// Routes init user
func Routes(routes iris.Party) {
	route := routes.Party(name)
	fmt.Println(route)
}
