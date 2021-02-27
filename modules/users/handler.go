package users

import (
	"github.com/kataras/iris/v12"
)

func GetUsers(c iris.Context) {
	users := GetAll()
	c.JSON(users)
}

func CreateUser(c iris.Context) {

}
