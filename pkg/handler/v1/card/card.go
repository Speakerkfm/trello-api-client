package card

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trello-api-client/pkg/middlewares/session"
	"trello-api-client/pkg/service/trello"
)

func DeleteCardById(c *gin.Context){
	err := trello.DeleteCardById(c.Param("id"), session.Session.Values["token"].(string))
	if err != nil {
		c.AbortWithError(http.StatusTeapot, err)

		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func UpdateCardStatusById(c *gin.Context){
	listID, ok :=  c.GetPostForm("listID")
	if !ok {
		c.AbortWithStatus(http.StatusTeapot)

		return
	}

	err := trello.UpdateCardStatusById(c.Param("id"), listID, session.Session.Values["token"].(string))
	if err != nil {
		c.AbortWithError(http.StatusTeapot, err)

		return
	}

	c.JSON(http.StatusNoContent, nil)
}
