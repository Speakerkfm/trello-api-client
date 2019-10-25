package v2

import (
	"github.com/gin-gonic/gin"
	"trello-api-client/pkg/middlewares/authentication"
)

func InitRoutes(g *gin.RouterGroup)  {
	g.Use(authentication.AuthenticationRequired())
	{
		SetCardRoutes(g)
	}
}