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
		c.JSON(iris.Map{
			"message": err.Error(),
		})
		return
	}
	c.JSON(users, context.
		JSON{Indent: "  "},
	)
	return
}

func createUser(c iris.Context) {
	var body Request
	var user User
	fmt.Println(body.Username)
	err := c.ReadJSON(&body)

	if err != nil {
		c.StatusCode(iris.StatusBadRequest)
		c.JSON(iris.Map{
			"message": err.Error(),
		})
		return
	}
	if result, _ := FindByUsername(body.Username); len(result) > 0 {
		c.StatusCode(iris.StatusBadRequest)
		c.JSON(iris.Map{
			"message": "Username Already Exist",
		})
		return
	}

	user.Hash, _ = config.GenerateFromPassword(body.Password)
	user.Username = body.Username
	if err := Create(user); err != nil {
		c.StatusCode(iris.StatusBadRequest)
		c.JSON(iris.Map{
			"message": err.Error(),
		})
		return
	}
	c.StatusCode(iris.StatusCreated)
	c.JSON(user)
	return
}

func loginUser(c iris.Context) {
	var body Request
	var userArray []User
	if err := c.ReadJSON(&body); err != nil {
		c.StatusCode(iris.StatusBadRequest)
		c.JSON(iris.Map{
			"message": err.Error(),
		})
		return
	}
	if (body.Username == "") || (body.Password == "") {
		c.StatusCode(iris.StatusOK)
		c.JSON(iris.Map{
			"message": "Username Or Password Empty",
		})
		return
	}
	if userArray, _ = FindByUsername(body.Username); len(userArray) == 0 {
		c.StatusCode(iris.StatusOK)
		c.JSON(iris.Map{
			"message": "Username Doesn't Exist",
		})
		return
	}
	authorized, _ := config.ComparePasswordAndHash(body.Password, userArray[0].Hash)
	if authorized {

		signedToken, err := auth.Wrapper.GenerateToken(body.Password)

		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{
				"message": "Error Signing Token",
			})
			return
		}

		tokenResponse := LoginResponse{
			Token: signedToken,
		}

		c.StatusCode(iris.StatusCreated)
		c.JSON(tokenResponse)

	} else {
		c.StatusCode(iris.StatusOK)
		c.JSON(iris.Map{
			"message": "Password Doesn't Match",
		})
	}

}
