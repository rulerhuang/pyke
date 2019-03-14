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

	router.GET("/rule-get", handlers.RuleGet)
	router.POST("/rule-set", handlers.RuleSet)
	router.POST("/rule-update", handlers.RuleUpdate)
	router.POST("/rule-delete", handlers.RuleDelete)

	router.POST("/rule-match", handlers.RuleMatch)

	uri := fmt.Sprintf("%s:%d", ServerConfig.Host, ServerConfig.Port)
	err := router.Run(uri)
	if nil != err {

	}
}
