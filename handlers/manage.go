package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"pyke/rule"
	"pyke/storage"
)

func RuleLoad(c *gin.Context) {
	num, err := storage.RuleStorage.Load()

	c.JSON(http.StatusOK, gin.H{"msg": "RuleLoad"})
}

func RuleSave(c *gin.Context) {

}

func RuleGet(c *gin.Context) {

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
