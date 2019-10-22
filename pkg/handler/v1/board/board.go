package board

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trello-api-client/pkg/middlewares/session"
	"trello-api-client/pkg/service/trello"
)

func DeleteBoardById(c *gin.Context){
	err := trello.DeleteBoardById(c.Param("id"), session.Session.Values["token"].(string))
	if err != nil {
		c.AbortWithError(http.StatusTeapot, err)

		return
	}

	c.JSON(http.StatusNoContent, nil)
}

