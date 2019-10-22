package pages

import (
	"github.com/gin-gonic/gin"
	"trello-api-client/pkg/middlewares/authentication"
)

func InitRoutes(g *gin.RouterGroup)  {
	SetMainPageRoutes(g)

	g.Use(authentication.AuthenticationRequired())
	{
		SetBoardRoutes(g)
	}
}
