package v1

import (
	"github.com/gin-gonic/gin"
	"trello-api-client/pkg/handler/v1/card"
)

func SetCardRoutes(router *gin.RouterGroup){
	router.POST("/cards/:id", card.UpdateCardStatusById)
	router.DELETE("/cards/:id", card.DeleteCardById)
}
