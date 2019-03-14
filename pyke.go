package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pyke/handlers"
	"pyke/utils"
)

var ServerConfig *utils.Config

func init() {
	ServerConfig = utils.GetConfig("./pyke_config.toml")
}

func main() {
	gin.SetMode(ServerConfig.Mode)
	router := gin.Default()

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

	serverUri := fmt.Sprintf("%s:%d", ServerConfig.Host, ServerConfig.Port)
	err := router.Run(serverUri)
	if nil != err {

	}
}
