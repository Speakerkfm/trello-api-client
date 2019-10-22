package v1

import (
	"github.com/gin-gonic/gin"
	"trello-api-client/pkg/middlewares/authentication"
)

func InitRoutes(g *gin.RouterGroup)  {
	SetTrelloLoginRoutes(g)
	SetLogoutRoute(g)

	g.Use(authentication.AuthenticationRequired())
	{
		SetCardRoutes(g)
		SetBoardRoutes(g)
	}
}