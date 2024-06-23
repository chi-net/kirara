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

type ServerInstallationForm struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	PasswordRepeat      string `json:"passwordRepeat"`
	ListenPort          string `json:"listenPort"`
	DatabaseType        string `json:"databaseType"`
	AuthenticationToken string `json:"authenticationToken"`
	Type                string `json:"type"`
}

func HandleServerInstallation(c *gin.Context) {
	form := ServerInstallationForm{}
	if c.ShouldBind(&form) == nil {
		c.Status(400)
		c.JSON(400, gin.H{
			"code":   400,
			"status": "failed",
		})
	}

}
