package board_p

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trello-api-client/pkg/middlewares/session"
	"trello-api-client/pkg/service/own"
	"trello-api-client/pkg/service/trello"
)

func GetAllBoards(c *gin.Context){
	boards, err := trello.GetUserBoards(session.Session.Values["token"].(string))
	if err != nil {
		c.AbortWithError(http.StatusTeapot, err)

		return
	}

	c.HTML(
		http.StatusOK,
		"boards.html",
		gin.H{
			"title": "Boards Page",
			"authorized": len(session.Session.Values) > 0,
			"payload": boards,
		},
	)
}

func GetBoardById(c *gin.Context){
	board, err := trello.GetUserBoard(c.Param("id"), session.Session.Values["token"].(string))
	if err != nil {
		c.AbortWithError(http.StatusTeapot, err)

		return
	}

	cards, err := trello.GetBoardCards(board, session.Session.Values["token"].(string))
	if err != nil {
		c.AbortWithError(http.StatusTeapot, err)

		return
	}

	c.HTML(
		http.StatusOK,
		"board.html",
		gin.H{
			"title": "Board Page",
			"boardName": board.Name,
			"boardDescription": board.Description,
			"boardID": board.ID,
			"authorized": len(session.Session.Values) > 0,
			"lists": board.Lists,
			"payload": cards,
		},
	)
}

func GetOwnBoard(c *gin.Context){
	cards, err := own.GetUserCards(int(session.Session.Values["user_id"].(float64)))
	if err != nil {
		c.AbortWithError(http.StatusTeapot, err)

		return
	}

	lists, err := own.GetUserLists(int(session.Session.Values["user_id"].(float64)))
	if err != nil {
		c.AbortWithError(http.StatusTeapot, err)

		return
	}

	c.HTML(
		http.StatusOK,
		"own_board.html",
		gin.H{
			"title": "Board Page",
			"boardName": "Your own board",
			"boardDescription": "Доска, которая хранится в БД",
			"authorized": len(session.Session.Values) > 0,
			"lists": lists,
			"payload": cards,
		},
	)
}