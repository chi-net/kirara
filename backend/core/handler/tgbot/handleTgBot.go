package tgbot

import (
	"context"
	"encoding/json"
	"github.com/chi-net/kirara/core/utils"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func KiraraTelegramBotHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	resp, _ := json.Marshal(update.Message)
	print(string(resp))
	if update.Message != nil {
		if !utils.CheckKiraraActivationInfo() {
			if update.Message.Chat.Type == "private" {
				b.SendMessage(ctx, &bot.SendMessageParams{
					ChatID: update.Message.Chat.ID,
					Text:   "You are activated Kirara with your telegram bot token successfully!\nHowever, you still have not configured your kirara instance yet.\nPlease refer to https://kirara.chinet.work/ for further settings.\nThis is what you say:\n" + update.Message.Text,
				})
			}
		}
	}
}
