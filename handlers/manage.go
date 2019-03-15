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
	num, err := storage.RuleStorageInstant.Load()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "RuleLoad", "rule": num})
}

func RuleSave(c *gin.Context) {

}

func RuleGet(c *gin.Context) {
	d, _ := storage.RuleStorageInstant.Get()
	c.JSON(http.StatusOK, gin.H{"msg": d})
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
