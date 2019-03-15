package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"pyke/config"
	"pyke/handlers"
	"pyke/logger"
)

func initGin() {
	gin.SetMode(config.PykeConfigInstant.Mode)
	gin.DisableConsoleColor()

	f, _ := os.Create("./gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func initRouter(router *gin.Engine) {
	router.GET("/", handlers.Index)

	manage := router.Group("/manage")
	{
		manage.GET("/load", handlers.RuleLoad)
		manage.GET("/get", handlers.RuleGet)

		manage.POST("/save", handlers.RuleSave)
		manage.POST("/set", handlers.RuleSet)
	}

	rule := router.Group("/rule")
	{
		rule.POST("/match", handlers.RuleMatch)
	}
}

func main() {
	initGin()

	router := gin.Default()
	initRouter(router)

	serverUri := fmt.Sprintf("%s:%d", config.PykeConfigInstant.Host, config.PykeConfigInstant.Port)
	err := router.Run(serverUri)
	if nil != err {
		logger.PykeError.Println(err)
	}
	logger.PykeInfo.Println("Server running at:", serverUri)
}
