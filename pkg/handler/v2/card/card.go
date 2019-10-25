package card

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"trello-api-client/pkg/middlewares/session"
	"trello-api-client/pkg/service/own"
)

func DeleteCardById(c *gin.Context){
	err := own.DeleteUserCard(c.Param("id"))
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

	err := own.UpdateCardStatus(int(session.Session.Values["user_id"].(float64)), c.Param("id"), listID)
	if err != nil {
		c.AbortWithError(http.StatusTeapot, err)

		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func CreateCard(c *gin.Context){
	cardName, ok := c.GetPostForm("cardName")
	if !ok {
		c.AbortWithStatus(http.StatusTeapot)

		return
	}

	cardDescription, ok := c.GetPostForm("cardDescription")
	if !ok {
		c.AbortWithStatus(http.StatusTeapot)

		return
	}

	err := own.CreateCard(int(session.Session.Values["user_id"].(float64)), cardName, cardDescription)
	if err != nil {
		c.AbortWithError(http.StatusTeapot, err)

		return
	}

	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/boards/own"))
}
