package card_p

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trello-api-client/pkg/middlewares/session"
	"trello-api-client/pkg/service/own"
)

func GetCard(c *gin.Context){
	card, err := own.GetCardByID(int(session.Session.Values["user_id"].(float64)), c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusTeapot, err)

		return
	}

	c.HTML(
		http.StatusOK,
		"card.html",
		gin.H{
			"title": "Card Page",
			"ID": card.ID,
			"cardName": card.Name,
			"cardDescription": card.Description,
			"authorized": len(session.Session.Values) > 0,
		},
	)
}
