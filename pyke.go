package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pyke/config"
	"pyke/handlers"
)

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
	gin.SetMode(config.PykeConfigInstant.Mode)
	router := gin.Default()
	initRouter(router)

	serverUri := fmt.Sprintf("%s:%d", config.PykeConfigInstant.Host, config.PykeConfigInstant.Port)
	err := router.Run(serverUri)
	if nil != err {

	}
}
