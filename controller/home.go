package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}
