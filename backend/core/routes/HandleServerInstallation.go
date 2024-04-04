package routes

import "github.com/gin-gonic/gin"

/*
HandleServerInstallation

	Package: Route
	Routes: /kirara/app/KiraraInstallation.action
	Method: POST
	Input:
		Username: Your Administrator Username
		Password: Your Administrator Password
		PasswordRepeat: Repeated Password
		ListenPort: Your Application's Listen Port
		DatabaseType: The type of your Database(SQLite Only now)
		AuthenticationToken: You should authenticate yourself as this application's genuine owner by entering this token.
		Type: Precheck(only check token) or Install(install this application)
	Output:
		Code: [HTTP Status Code]
		Status: [Successful or not]
		Error: [optional Error messages]
*/
func HandleServerInstallation(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":   200,
		"status": "success",
	})
}
