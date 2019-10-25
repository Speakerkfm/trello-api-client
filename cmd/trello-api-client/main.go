package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"gitlab.loc/xsolla-login/go-xsolla-login/pkg/log"
	"os"
	"time"
	"trello-api-client/pkg/middlewares/cache"
	"trello-api-client/pkg/middlewares/session"
	"trello-api-client/pkg/routers/pages"
	"trello-api-client/pkg/routers/v1"
	"trello-api-client/pkg/routers/v2"
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

	//mysql
	mysqlConf := mysql.NewConfig()
	mysqlConf.Net = "tcp"
	mysqlConf.Addr = os.Getenv("DATABASE_HOST")+ ":" + os.Getenv("DATABASE_PORT")
	mysqlConf.User = os.Getenv("DATABASE_USER")
	mysqlConf.Passwd = os.Getenv("DATABASE_PASSWORD")
	mysqlConf.DBName = os.Getenv("DATABASE_NAME")
	mysqlConf.MultiStatements = true
	mysqlConf.ParseTime = true
	mysqlConf.Loc = time.Local
	mysqlConf.Collation = "utf8mb4_general_ci"

	db, err := gorm.Open("mysql", mysqlConf.FormatDSN())

	checkErrFatal(err, "Mysql connection failed")

	defer db.Close()
	db.SingularTable(true)

	//store
	store.Config(os.Getenv("SECRET_KEY"), os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PWD"), 10, db)

	trello.Config(os.Getenv("APP_HOST"), os.Getenv("APP_PORT"), os.Getenv("TRELLO_KEY"), os.Getenv("TRELLO_SECRET"))

	//middleware before
	router.Use(cors.Default())
	router.Use(session.MiddlewareSession())

	//routers
	router.Static("/public", "./public")
	pages.InitRoutes(router.Group("/"))
	v1.InitRoutes(router.Group("/v1"))
	v2.InitRoutes(router.Group("/v2"))

	//middleware after
	router.Use(cache.MiddlewareCache())

	router.Run(os.Getenv("APP_PORT"))
}

func checkErrFatal(err error, msg string) {
	if err != nil {
		log.Fatal().Err(err).Msg(msg)
		os.Exit(1)
	}
}
