package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Home(c *gin.Context) {
	logger := log.WithFields(log.Fields{
		"url": c.Request.Host + c.Request.URL.Path,
	})
	logger.Info("access log start")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}
