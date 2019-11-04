package pages

import (
	"github.com/gin-gonic/gin"
	"trello-api-client/pkg/handler/pages/card_p"
)

func SetCardRoutes(router *gin.RouterGroup) {
	router.GET("/cards/:id", card_p.GetCard)
}

