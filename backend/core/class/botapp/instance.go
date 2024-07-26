package botapp

import (
	"context"
	"github.com/chi-net/kirara/core/handler/tgbot"
	"github.com/go-telegram/bot"
	"log"
	"os"
)

var KiraraTelegramBotInstance bot.Bot

func InitializeKiraraTgBotService(ctx context.Context, cancel context.CancelFunc, token string) {
	opts := []bot.Option{
		bot.WithDefaultHandler(tgbot.KiraraTelegramBotHandler),
	}

	KiraraTelegramBotInstance, err := bot.New(token, opts...)
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

func GetKiraraBotInstance() bot.Bot {
	return KiraraTelegramBotInstance
}
