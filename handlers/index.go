package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pyke/config"
)

func Index(c *gin.Context) {
	host, port, mode :=
		config.PykeConfigInstant.Host,
		config.PykeConfigInstant.Port,
		config.PykeConfigInstant.Mode

	msg := fmt.Sprintf("hello, pyke is running at %s:%d, at %s mode!", host, port, mode)
	fmt.Fprintln(gin.DefaultWriter, msg)
	c.JSON(http.StatusOK, gin.H{"msg": msg})
}
