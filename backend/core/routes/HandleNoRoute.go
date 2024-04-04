package routes

import "github.com/gin-gonic/gin"

/*
HandleNoRoute

	Package: Route
	Input: none
	Method: ALL
	Description: Handle all of those route have no route yet, send an aborted connection.
*/
func HandleNoRoute(c *gin.Context) {
	c.AbortWithStatus(403)
}
