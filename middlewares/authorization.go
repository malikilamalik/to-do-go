package middlewares

import (
	"fmt"
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
		claims, err := auth.Wrapper.ValidateToken(clientToken)
		if err != nil {
			c.StatusCode(iris.StatusUnauthorized)
			c.JSON(iris.Map{
				"message": "Incorrect Format of Authorization Token",
			})
			return
		}
		c.SetCookieKV("id", fmt.Sprint(claims.Id))
		c.Next()
	}
}
