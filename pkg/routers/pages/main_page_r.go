package pages

import (
	"github.com/gin-gonic/gin"
	"trello-api-client/pkg/handler/pages/index"
)

func SetMainPageRoutes(router *gin.RouterGroup) {
	router.GET("", index.ShowIndexPage)
}

