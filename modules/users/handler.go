package users

import (
	"fmt"
	"to-do-go/config"

	"to-do-go/modules/auth"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func getUsers(c iris.Context) {
	users, err := GetAll()
	if err != nil {
		c.StatusCode(iris.StatusBadRequest)
		c.WriteString(err.Error())
		return
	}
	c.JSON(users, context.
		JSON{Indent: "  "},
	)
}

func createUser(c iris.Context) {
	var body auth.Request
	var user User
	fmt.Println(body.Username)
	err := c.ReadJSON(&body)

	if err != nil {
		c.StatusCode(iris.StatusBadRequest)
		c.WriteString(err.Error())
		return
	}
	if result, _ := FindByUsername(body.Username); len(result) > 0 {
		c.StatusCode(iris.StatusBadRequest)
		c.WriteString("Username Already Exist")
		return
	}

	user.Hash, _ = config.GenerateFromPassword(body.Password)
	user.Username = body.Username
	if err := Create(user); err != nil {
		c.StatusCode(iris.StatusBadRequest)
		c.WriteString(err.Error())
		return
	}
	c.StatusCode(iris.StatusCreated)
	c.WriteString("User Created")
}

func loginUser(c iris.Context) {
	var body auth.Request
	var userArray []User
	if err := c.ReadJSON(&body); err != nil {
		c.StatusCode(iris.StatusBadRequest)
		c.WriteString(err.Error())
		return
	}
	if (body.Username == "") || (body.Password == "") {
		c.StatusCode(iris.StatusBadRequest)
		c.WriteString("Username Or Password Empty")
		return
	}
	if userArray, _ = FindByUsername(body.Username); len(userArray) == 0 {
		c.StatusCode(iris.StatusBadRequest)
		c.WriteString("Username Doesn't Exist")
		return
	}
	authorized, _ := config.ComparePasswordAndHash(body.Password, userArray[0].Hash)
	if authorized {
		c.StatusCode(iris.StatusOK)
		c.WriteString("MASUK PAK EKO")
	} else {
		c.StatusCode(iris.StatusOK)
		c.WriteString("TIDAK MASUK PAK EKO")
	}

}
