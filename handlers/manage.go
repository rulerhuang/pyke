package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"pyke/rule"
	"pyke/storage"
)

func RuleGet(c *gin.Context) {
	RuleStorage.Load()
	c.JSON(http.StatusOK, gin.H{"msg": "RuleGet"})
}

func RuleSet(c *gin.Context) {
	tr := rule.Rule{}
	err := c.ShouldBind(&tr)
	if err != nil {
		log.Printf("error:%s", err.Error())
	}
	fmt.Printf("%+v,\n", tr)
	storage.Save(&tr)
}

func RuleUpdate(c *gin.Context) {

}

func RuleDelete(c *gin.Context) {

}
