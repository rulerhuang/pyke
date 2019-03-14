package main

import (
	"github.com/gin-gonic/gin"
	"pyke/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/", handlers.Index)

	router.GET("/rule-get", handlers.RuleGet)
	router.POST("/rule-set", handlers.RuleSet)
	router.POST("/rule-update", handlers.RuleUpdate)
	router.POST("/rule-delete", handlers.RuleDelete)

	err := router.Run(":8000")
	if nil != err {

	}
}
