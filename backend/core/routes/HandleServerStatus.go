package routes

import "github.com/gin-gonic/gin"

func HandleServerStatus(c *gin.Context, installationRequired bool) {
	if installationRequired {
		c.JSON(200, gin.H{
			"code":   200,
			"status": "uninstalled",
		})
	} else {
		c.JSON(200, gin.H{
			"code":   200,
			"status": "ok",
		})
	}
}
