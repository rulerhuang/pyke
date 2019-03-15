package handlers

import (
	"github.com/gin-gonic/gin"
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
	c.JSON(http.StatusOK, gin.H{"msg": "RuleLoad", "rule num": num})
}

func RuleGet(c *gin.Context) {
	d, _ := storage.RuleStorageInstant.Get()
	c.JSON(http.StatusOK, gin.H{"msg": "RuleGet", "rules": d})
}

func RuleSave(c *gin.Context) {
	err := storage.RuleStorageInstant.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "RuleSave"})

}

func RuleSet(c *gin.Context) {
	tr := rule.Rule{}
	err := c.ShouldBind(&tr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	err = storage.RuleStorageInstant.Set(&tr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "RuleSet"})
}
