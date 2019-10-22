package authentication

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trello-api-client/pkg/middlewares/session"
)

func AuthenticationRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := session.Session.Values["token"]
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user needs to be signed in to access this service"})
			c.Abort()
			return
		}

		c.Next()
	}
}