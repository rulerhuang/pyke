package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pyke/config"
	"pyke/handlers"
	"pyke/storage"
)

var ServerConfig *config.Config
var RuleStorage storage.Storage

func init() {
	ServerConfig, _ = config.New()
	RuleStorage = storage.New(storage.JSON_MODE)
}

func initRouter(router *gin.Engine) {
	router.GET("/", handlers.Index)

	manage := router.Group("/manage")
	{
		manage.GET("/get", handlers.RuleGet)
		manage.POST("/set", handlers.RuleSet)
		manage.POST("/update", handlers.RuleUpdate)
		manage.POST("/delete", handlers.RuleDelete)
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
