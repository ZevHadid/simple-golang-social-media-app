package middleware

import (
	"net/http"
	"simple-golang-social-media-app/pkg/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			handleUnauthorized(c, "no-auth-token-found")
			return
		}

		email, err := utils.ValidateJWT(tokenString)
		if err != nil {
			handleUnauthorized(c, "invalid-auth-token")
			return
		}

		c.Set("email", email)
		c.Next()
	}
}

func handleUnauthorized(c *gin.Context, reason string) {
	c.Redirect(http.StatusFound, "/login?reason="+reason)
	c.Abort()
}
