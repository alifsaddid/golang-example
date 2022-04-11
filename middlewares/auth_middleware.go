package middlewares

import (
	"Oauth/dto/context"
	"Oauth/dto/response"
	"Oauth/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtService services.JWTService) gin.HandlerFunc {

	return func(c *gin.Context) {
		accessToken := c.Request.Header.Get("Authorization")
		if accessToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.BaseErrorResponse{
				Error:            "Authentication error",
				ErrorDescription: "Missing authorization token",
			})
			return
		}

		accessToken = strings.ReplaceAll(accessToken, "Bearer ", "")
		user, token, err := jwtService.Verify(accessToken)
		if err != nil {
			c.AbortWithStatusJSON(err.Status, response.BaseErrorResponse{
				Error:            err.Err.Error(),
				ErrorDescription: err.Err.Error(),
			})
			return
		}

		c.Set("currentUser", context.UserContext{
			Id:           int(user.ID),
			Username:     user.Username,
			Name:         user.FullName,
			Npm:          user.NPM,
			ClientId:     token.ClientId,
			AccessToken:  token.Token,
			RefreshToken: token.RefreshToken,
		})

		c.Next()
	}
}
