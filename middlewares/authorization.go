package middlewares

import (
	"to-do-go/modules/auth"

	"github.com/kataras/iris/v12"
)

func Authorization() iris.Handler {
	return func(c iris.Context) {
		clientToken := c.Request().Header.Get("Authorization")

		if clientToken == "" {
			c.StatusCode(iris.StatusForbidden)
			c.JSON(iris.Map{
				"message": "No Authorization header provided",
			})
			return
		}
		if _, err := auth.Wrapper.ValidateToken(clientToken); err != nil {
			c.StatusCode(iris.StatusUnauthorized)
			c.JSON(iris.Map{
				"message": "Incorrect Format of Authorization Token",
			})
			return
		}
		c.Next()
	}
}
