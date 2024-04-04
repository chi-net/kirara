package handler

import (
	"context"
	"encoding/json"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func KiraraTelegramBotHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	resp, _ := json.Marshal(update.Message)
	print(string(resp))
	if update.Message.Chat.Type == "private" {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   update.Message.Text,
		})
	}
}
