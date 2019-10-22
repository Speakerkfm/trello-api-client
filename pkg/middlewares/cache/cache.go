package cache

import "github.com/gin-gonic/gin"

func MiddlewareCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Cache-Control", "no-cache")

		c.Next()
	}
}
