package main

import (
	"github.com/gin-gonic/gin"
	"pyke/handlers"
	"pyke/utils"
)

var ServerConfig *utils.Config

func init() {
	ServerConfig = utils.GetConfig("./pyke_config.toml")
}

func main() {
	router := gin.Default()
	router.GET("/", handlers.Index)

	router.GET("/rule-get", handlers.RuleGet)
	router.POST("/rule-set", handlers.RuleSet)
	router.POST("/rule-update", handlers.RuleUpdate)
	router.POST("/rule-delete", handlers.RuleDelete)

	router.POST("/rule-match", handlers.RuleMatch)

	err := router.Run(":8000")
	if nil != err {

	}
}
