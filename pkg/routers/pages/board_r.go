package pages

import (
	"github.com/gin-gonic/gin"
	"trello-api-client/pkg/handler/pages/board_p"
)

func SetBoardRoutes(router *gin.RouterGroup) {
	router.GET("/boards", board_p.GetAllBoards)
	router.GET("/boards/own", board_p.GetOwnBoard)
	router.GET("/boards/trello/:id", board_p.GetBoardById)
}
