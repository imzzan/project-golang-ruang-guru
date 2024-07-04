package middleware

import (
	"a21hc3NpZ25tZW50/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		token, err := ctx.Cookie("session_token")
		if err != nil {
			if ctx.GetHeader("Content-type") == "application/json" {
				ctx.JSON(401, model.NewErrorResponse(err.Error()))
				return
			} else {
				ctx.Redirect(303, "/users/login")
			}
			return
		}

		secrettoken, err := jwt.ParseWithClaims(token, &model.Claims{}, func(t *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})

		if err != nil {
			ctx.JSON(400, model.NewErrorResponse(err.Error()))
			return
		}

		user, ok := secrettoken.Claims.(*model.Claims)
		if !ok || !secrettoken.Valid {
			ctx.JSON(400, model.NewErrorResponse(err.Error()))
			return
		}

		ctx.Set("email", user.Email)
		ctx.Next()
	})
}
