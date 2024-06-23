package main

import (
	"context"
	"github.com/chi-net/kirara/config"
	"github.com/chi-net/kirara/core/db/sqlite"
	"github.com/chi-net/kirara/core/handler"
	"github.com/chi-net/kirara/core/routes"
	"github.com/chi-net/kirara/core/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-telegram/bot"
	"os"
	"os/signal"
	"strconv"
)

// Kirara - A magic tool enables you to communicate with your friends in Telegram everywhere and anytime.
// Licensed under GPL3, Made with love and passion by chi Network Contributors(c)2022-2024.
// The icon of this application is an AIGC content and it was provided by baiyuanneko.

var IsApplicationActivated bool = false
var ApplicationListenPort int = 8080
var KiraraTelegramBotInstance bot.Bot
var KiraraDatabaseInstance sqlite.SQLiteDBInstance

func init() {
	dbPath := ""
	dir, _ := os.Getwd()

	if utils.CheckConfiguration(dir + string(os.PathSeparator) + "kirara.config.json") {

		conf := utils.ReadJSONConfiguration(dir + string(os.PathSeparator) + "kirara.config.json")

		if conf.DbPath != "Failed to GET" {
			IsApplicationActivated = true
		} else {
			dbPath = conf.DbPath
		}

		if conf.ListenPort != -1 {
			ApplicationListenPort = conf.ListenPort
		}

		// we use SQLite to store some settings data including your credentials of MySQL, encrypted password, username API Key etc.
		// before the application run, we should get some configurations from SQLite.
		if IsApplicationActivated {
			KiraraDatabaseInstance, err := sqlite.New(dbPath)
			if err != nil {
				panic(err)
			}

			_, err = KiraraDatabaseInstance.Exec("SELECT * FROM USERS")
			if err != nil {
				panic(err)
			}

		}

		// initialize Telegram Bot Instance
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
		defer cancel()

		opts := []bot.Option{
			bot.WithDefaultHandler(handler.KiraraTelegramBotHandler),
		}

		KiraraTelegramBotInstance, err := bot.New(config.KiraraTelegramBotToken, opts...)
		if err != nil {
			panic(err)
		}

		KiraraTelegramBotInstance.Start(ctx)
	}
}

func main() {
	app := gin.New()
	app.Use(gin.Logger())

	// Installed so that you do not need to install it anymore
	if IsApplicationActivated {
		app.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"code": 200,
			})
		})
		app.POST("/kirara/app/KiraraServerStatus.action", func(c *gin.Context) {
			routes.HandleServerStatus(c, false)
		})
	} else {
		// You have not configured this application so you can not use its features
		// Some installation option routes are listed below
		app.POST("/kirara/app/KiraraServerStatus.action", func(c *gin.Context) {
			routes.HandleServerStatus(c, true)
		})
		app.POST("/kirara/app/KiraraInstallation.action", routes.HandleServerInstallation)
	}
	app.NoRoute(routes.HandleNoRoute)
	err := app.Run(":" + strconv.Itoa(ApplicationListenPort))
	if err != nil {
		panic(err)
	}
}
