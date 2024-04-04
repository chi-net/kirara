package main

import (
	"context"
	"fmt"
	"github.com/chi-net/kirara/config"
	"github.com/chi-net/kirara/core/handler"
	"github.com/chi-net/kirara/core/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-telegram/bot"
	"os"
	"os/signal"
)

func main() {
	dir, _ := os.Getwd()
	fmt.Println(dir + string(os.PathSeparator) + "kirara.config.json")
	if utils.CheckConfiguration(dir + string(os.PathSeparator) + "kirara.config.json") {
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

		app := gin.New()
		app.Use(gin.Logger())

		app.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"code": 200,
			})
		})
		app.Run(":8080")
	} else {
		fmt.Println("Error!")
	}
}
