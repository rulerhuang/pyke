package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pyke/config"
	"pyke/handlers"
)

var ServerConfig *config.Config

func init() {
	ServerConfig, _ = config.New()
}

func initRouter(router *gin.Engine) {
	router.GET("/", handlers.Index)

	manage := router.Group("/manage")
	{
		manage.GET("/load", handlers.RuleLoad)
		manage.POST("/save", handlers.RuleSave)

		manage.POST("/get", handlers.RuleGet)
		manage.POST("/set", handlers.RuleSet)
	}

	rule := router.Group("/rule")
	{
		rule.POST("/match", handlers.RuleMatch)
	}
}

func main() {
	gin.SetMode(ServerConfig.Mode)
	router := gin.Default()
	initRouter(router)

	serverUri := fmt.Sprintf("%s:%d", ServerConfig.Host, ServerConfig.Port)
	err := router.Run(serverUri)
	if nil != err {

	}
}
