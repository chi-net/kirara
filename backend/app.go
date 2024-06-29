package main

import (
	"context"
	"github.com/chi-net/kirara/config"
	"github.com/chi-net/kirara/core/db/sqlite"
	"github.com/chi-net/kirara/core/handler/tgbot"
	"github.com/chi-net/kirara/core/routes"
	"github.com/chi-net/kirara/core/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-telegram/bot"
	"log"
	"os"
	"os/signal"
	"strconv"
)

// Kirara - A magic tool enables you to communicate with your friends in Telegram everywhere and anytime.
// Licensed under GPL3, Made with love and passion by chi Network Contributors(c)2022-2024.
// The icon of this application is an AIGC content and it was provided by baiyuanneko.

var IsApplicationActivated bool = false
var ApplicationListenPort int = 8080

func initializeKiraraTgBotService(ctx context.Context, cancel context.CancelFunc) {
	opts := []bot.Option{
		bot.WithDefaultHandler(tgbot.KiraraTelegramBotHandler),
	}

	KiraraTelegramBotInstance, err := bot.New(config.KiraraTelegramBotToken, opts...)
	if err != nil {
		panic(err)
	}

	KiraraTelegramBotInstance.Start(ctx)

	<-ctx.Done()
	log.Println("[Kirara TgBotService] shutting down")
	cancel()
	// workaround
	os.Exit(1)
}

func main() {
	app := gin.New()
	app.Use(gin.Logger())

	dbPath := ""
	dir, _ := os.Getwd()

	if utils.CheckConfiguration(dir + string(os.PathSeparator) + "kirara.config.json") {

		conf := utils.ReadJSONConfiguration(dir + string(os.PathSeparator) + "kirara.config.json")
		IsApplicationActivated := utils.CheckKiraraActivationInfo()
		if conf.ListenPort != -1 {
			ApplicationListenPort = conf.ListenPort
		}

		// we use SQLite to store some settings data including your credentials of MySQL, encrypted password, username API Key etc.
		// before the application run, we should get some configurations from SQLite.
		if IsApplicationActivated {
			dbPath = conf.DbPath

			err := sqlite.New(dbPath)
			if err != nil {
				panic(err)
			}

			_, err = sqlite.Exec("SELECT * FROM USERS")
			if err != nil {
				panic(err)
			}

		}
	}

	// initialize Telegram Bot Instance
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	go initializeKiraraTgBotService(ctx, cancel)

	// Installed so that you do not need to install it anymore
	if IsApplicationActivated {
		app.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"code": 200,
			})
		})
		app.POST("/kirara/app/status", func(c *gin.Context) {
			routes.HandleServerStatus(c, false)
		})
	} else {
		// You have not configured this application so you can not use its features
		// Some installation option routes are listed below
		app.POST("/kirara/app/status", func(c *gin.Context) {
			routes.HandleServerStatus(c, true)
		})

		token := utils.GenerateRandomString(32)
		app.POST("/kirara/app/install", func(c *gin.Context) {
			routes.HandleServerInstallation(c, token)
		})

		log.Println("[Kirara Installation] Your token is:", token)
		log.Println("Please DO NOT TELL OTHER YOUR TOKEN!")
		log.Println("If you think the token was stolen, please restart this application.")
	}

	app.NoRoute(routes.HandleNoRoute)
	err := app.Run(":" + strconv.Itoa(ApplicationListenPort))
	if err != nil {
		log.Println("[Kirara Webserver]\nPopup: If it seems that the application is binded a port that has already been used, Please create a `kirara.config.json` and fill it with\n{\n  \"ListenPort\": [An integer Port You want to use],\n  \"DBPath\": \"Failed to GET\"\n}\nWe will overwrite this file when it becomes installing.")
		panic(err.Error())
	}

}
