package login

import (
	"github.com/dghubble/oauth1"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"gitlab.loc/xsolla-login/go-xsolla-login/pkg/log"
	"net/http"
	"time"
	"trello-api-client/pkg/middlewares/session"
	"trello-api-client/pkg/service/trello"
	"trello-api-client/pkg/store"
)

func TrelloLogin(c *gin.Context) {
	var requestToken string
	var requestSecret string
	var err error

	requestToken, requestSecret, err = trello.AuthConfig.RequestToken()
	if err != nil {
		log.Error().Err(err).Msg("getting request token failed")
		c.AbortWithError(http.StatusFailedDependency, err)

		return
	}

	authorizationURL, err := trello.AuthConfig.AuthorizationURL(requestToken)
	if err != nil {
		log.Error().Err(err).Msg("getting auth url failed")
		c.AbortWithError(http.StatusFailedDependency, err)

		return
	}

	store.Set(requestToken, requestSecret, 10 * time.Minute)

	c.Redirect(http.StatusMovedPermanently, authorizationURL.String())
}

func TrelloCallback(c *gin.Context) {
	requestToken, verifier, err := oauth1.ParseAuthorizationCallback(c.Request)
	if err != nil {
		log.Error().Err(err).Msg("getting request token failed")
		c.AbortWithError(http.StatusFailedDependency, err)

		return
	}

	requestSecret := store.Get(requestToken)

	accessToken, accessSecret, err := trello.AuthConfig.AccessToken(requestToken, requestSecret, verifier)
	if err != nil {
		log.Error().Err(err).Msg("getting access token failed")
		c.AbortWithError(http.StatusFailedDependency, err)

		return
	}

	token := oauth1.NewToken(accessToken, accessSecret)

	session.Session.Values["token"] = token.Token
	session.Session.Values["token_secret"] = token.TokenSecret

	sessions.Save(c.Request, c.Writer)

	c.Redirect(http.StatusMovedPermanently, "/")
}

func Logout(c *gin.Context){
	session.Session.Options.MaxAge = -1
	sessions.Save(c.Request, c.Writer)

	c.Redirect(http.StatusMovedPermanently, "/")
}
