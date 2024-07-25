package routes

import "github.com/gin-gonic/gin"

type UserLoginForm struct {
	Username       string `json:"username" binding:"required"`
	Password       string `json:"password" binding:"required"`
	PasswordRepeat string `json:"passwordRepeat" binding:"required"`
}

func HandleUserLogin(c *gin.Context) {
	loginForm := UserLoginForm{}

	if c.ShouldBind(&loginForm) != nil {
		c.JSON(400, gin.H{
			"code":   400,
			"status": "not enough content",
		})
		return
	}

}
