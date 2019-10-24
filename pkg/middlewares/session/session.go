package session

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"gitlab.loc/xsolla-login/go-xsolla-login/pkg/log"
	"trello-api-client/pkg/store"
)

var Session *sessions.Session

func MiddlewareSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		Session, err = store.Storage.SessionStore.Get(c.Request, "X-AUTH-SESSION")
		if err != nil {
			log.Error().Err(err).Msg("getting session failed")
		}

		c.Next()
	}
}