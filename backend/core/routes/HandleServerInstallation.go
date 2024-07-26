package routes

import (
	"fmt"
	"github.com/chi-net/kirara/core/utils"
	"github.com/gin-gonic/gin"
)

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
		AuthenticationToken: You should authenticate yourself as this application's genuine owner by entering this tokens.
		Type: Precheck(only check tokens) or Install(install this application)
	Output:
		Code: [HTTP Status Code]
		Status: [Successful or not]
		Error: [optional Error messages]
*/

type ServerInstallationForm struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	PasswordRepeat      string `json:"passwordRepeat"`
	ListenPort          int    `json:"listenPort"`
	DatabaseType        string `json:"databaseType"`
	AuthenticationToken string `json:"authenticationToken"`
	TgBotToken          string `json:"tgBotToken"`
	Type                string `json:"type"`
}

func HandleServerInstallation(c *gin.Context, token string) {
	form := ServerInstallationForm{}
	if c.ShouldBind(&form) != nil {
		c.JSON(400, gin.H{
			"code":   400,
			"status": "failed",
		})
		return
	}

	if token != form.AuthenticationToken {
		c.JSON(400, gin.H{
			"code":   400,
			"status": "invalid tokens",
		})
		return
	}

	if form.Password != form.PasswordRepeat && form.Password != "" {
		c.JSON(400, gin.H{
			"code":   400,
			"status": "invalid password",
		})
		return
	}

	// start installation process
	err := utils.InitializeApplication(form.Username, form.Password, form.ListenPort, form.DatabaseType, form.TgBotToken)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{
			"code":   400,
			"status": "failed to initialize application",
		})
		return
	}

	c.JSON(200, gin.H{
		"code":   200,
		"status": "success",
	})

}
