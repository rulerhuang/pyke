package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pyke/rule"
	"pyke/storage"
)

func RuleLoad(c *gin.Context) {
	num, err := storage.PykeStorageInstant.Load()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "RuleLoad", "rule num": num})
}

func RuleGet(c *gin.Context) {
	d, _ := storage.PykeStorageInstant.Get()
	c.JSON(http.StatusOK, gin.H{"msg": "RuleGet", "rules": d})
}

func RuleSave(c *gin.Context) {
	err := storage.PykeStorageInstant.Save()
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

	err = storage.PykeStorageInstant.Set(&tr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "RuleSet"})
}
