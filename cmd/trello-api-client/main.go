package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gitlab.loc/xsolla-login/go-xsolla-login/pkg/log"
	"os"
	"trello-api-client/pkg/middlewares/cache"
	"trello-api-client/pkg/middlewares/session"
	"trello-api-client/pkg/routers/pages"
	"trello-api-client/pkg/routers/v1"
	"trello-api-client/pkg/service/trello"
	"trello-api-client/pkg/store"
)

var programName = "trello-api-client"

func main(){
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	router := gin.New()
	router.LoadHTMLGlob("public/templates/*")

	log.Config(programName, os.Stderr)

	//store
	store.Config(os.Getenv("SECRET_KEY"))

	trello.Config(os.Getenv("TRELLO_KEY"), os.Getenv("TRELLO_SECRET"))

	//middleware before
	router.Use(session.MiddlewareSession())

	//routers
	router.Static("/public", "./public")
	pages.InitRoutes(router.Group("/"))
	v1.InitRoutes(router.Group("/v1"))

	//middleware after
	router.Use(cache.MiddlewareCache())

	router.Run(":8080")
}

func checkErrFatal(err error, msg string) {
	if err != nil {
		log.Fatal().Err(err).Msg(msg)
		os.Exit(1)
	}
}
