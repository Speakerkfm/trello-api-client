package v1

import (
	"github.com/gin-gonic/gin"
	"trello-api-client/pkg/handler/v1/board"
)


func SetBoardRoutes(router *gin.RouterGroup){
	router.DELETE("/boards/:id", board.DeleteBoardById)
}

