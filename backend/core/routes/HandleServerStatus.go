package routes

import "github.com/gin-gonic/gin"

/*
HandleServerStatus

	Package: Routes
	Path: /kirara/app/KiraraServerStatus.action
	Method: POST
	function: installationRequired(is this server has configured yet?)
	Input: none
	Output: Server Status(Installed or not)
*/
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
