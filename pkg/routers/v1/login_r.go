package v1

import (
	"github.com/gin-gonic/gin"
	"trello-api-client/pkg/handler/v1/login"
)

func SetTrelloLoginRoutes(router *gin.RouterGroup) {
	router.GET("/trello/login_redirect", login.TrelloLogin)
	router.GET("/trello/callback", login.TrelloCallback)
}

func SetLogoutRoute(router *gin.RouterGroup) {
	router.GET("/logout", login.Logout)
}