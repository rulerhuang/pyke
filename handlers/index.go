package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pyke/config"
)

func Index(c *gin.Context) {
	host, port, mode :=
		config.ConfigInstant.Host,
		config.ConfigInstant.Port,
		config.ConfigInstant.Mode

	msg := fmt.Sprintf("hello, pyke is running at %s:%d, at %s mode!", host, port, mode)
	c.JSON(http.StatusOK, gin.H{"msg": msg})
}
