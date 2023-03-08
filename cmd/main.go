package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	status := "online"

	router.GET("/health/online", func(c *gin.Context) {
		status = "online"
		c.String(http.StatusOK, "do online")
	})

	router.GET("/health/offline", func(c *gin.Context) {
		status = "offline"
		c.String(http.StatusOK, "do offline")
	})

	router.GET("/health/check", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	router.GET("/health/status", func(c *gin.Context) {
		if status != "online" {
			c.String(http.StatusBadRequest, "bad request")
			return
		}
		c.String(http.StatusOK, "ok")
	})

	router.Run()
}
