package main

import (
	"context"
	"github.com/chi-net/kirara/config"
	"github.com/chi-net/kirara/core/handler"
	"github.com/chi-net/kirara/core/routes"
	"github.com/chi-net/kirara/core/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-telegram/bot"
	"os"
	"os/signal"
)

// Kirara - A magic tool enables you to communicate with your friends in Telegram everywhere and anytime.
// Licensed under GPL3, Made with love and passion by chi Network Contributors(c)2022-2024.
// The icon of this application is an AIGC content and it was provided by baiyuanneko.

func main() {
	dir, _ := os.Getwd()
	app := gin.New()
	app.Use(gin.Logger())
	if utils.CheckConfiguration(dir + string(os.PathSeparator) + "kirara.config.json") {
		// we use SQLite to store some settings data including your credentials of MySQL, encrypted password, username API Key etc.
		// before the application run, we should get some configurations from SQLite.
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
		defer cancel()

		opts := []bot.Option{
			bot.WithDefaultHandler(handler.KiraraTelegramBotHandler),
		}

		b, err := bot.New(config.KiraraTelegramBotToken, opts...)
		if err != nil {
			panic(err)
		}

		b.Start(ctx)

		// Installed so that you do not need to install it anymore
		app.POST("/api/KiraraServerStatus.action", func(c *gin.Context) {
			routes.HandleServerStatus(c, false)
		})

		app.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"code": 200,
			})
		})
	} else {
		// You have not configured this application so you can not use its features
		// Some installation option routes are listed below
		app.POST("/api/KiraraServerStatus.action", func(c *gin.Context) {
			routes.HandleServerStatus(c, true)
		})
	}
	app.Run(":8080")
}
