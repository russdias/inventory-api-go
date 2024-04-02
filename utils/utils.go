package utils

import "github.com/gin-gonic/gin"

func FormatErrorAndSend(c *gin.Context, message string) {
	c.JSON(404, gin.H{
		"success": false,
		"error":   message})
}

func FormatJSONAndSend(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"success": true,
		"data":    data,
	})
}
