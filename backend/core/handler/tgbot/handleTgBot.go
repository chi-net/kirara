package tgbot

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chi-net/kirara/core/db/sqlite"
	"github.com/chi-net/kirara/core/utils"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"strconv"
)

type ImageData struct {
	Fileid  []string `json:"fileid"`
	Caption string   `json:"caption"`
}

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
		} else {
			// fixed type: 0 is text, 1 is stickers, 2 is photos
			id := update.Message.ID
			from := update.Message.From.ID
			chat := update.Message.Chat.ID
			date := update.Message.Date
			r := -1
			t := 0
			content := ""
			if update.Message.ReplyToMessage != nil {
				r = update.Message.ReplyToMessage.ID
			}

			if update.Message.Sticker != nil {
				t = 1
				content = update.Message.Sticker.FileID
			}
			if update.Message.Text != "" {
				content = update.Message.Text
			}
			if update.Message.Photo != nil {
				t = 2
				imagelist := []string{}
				for i := range update.Message.Photo {
					imagelist = append(imagelist, update.Message.Photo[i].FileID)
					file, _ := b.GetFile(ctx, &bot.GetFileParams{FileID: update.Message.Photo[i].FileID})
					fmt.Println(b.FileDownloadLink(file))
				}
				photo := ImageData{
					Fileid:  imagelist,
					Caption: update.Message.Caption,
				}
				data, _ := json.Marshal(photo)
				content = string(data)
			}
			insertMessagesSQL := "INSERT INTO `messages`(`id`, `from`, `chat`, `date`, `content`, `quote`, `type`)VALUES (?,?,?,?,?,?,?);"
			log.Println(insertMessagesSQL)
			result, err := sqlite.Exec(insertMessagesSQL, strconv.Itoa(id), strconv.FormatInt(from, 10), strconv.FormatInt(chat, 10), strconv.Itoa(date), content, strconv.Itoa(r), strconv.Itoa(t))
			if err != nil {
				log.Println(insertMessagesSQL)
				log.Println(err.Error())
			}
			if result != nil {
				affect, _ := result.RowsAffected()
				if affect < 1 {
					log.Println("[Kirara TgBot]Can not insert messages")
				}
			} else {
				log.Println("[Kirara TgBot]Can not insert messages")
			}
		}
		//if update.Message.Chat.Type == "private" {
		//	if update.Message.Sticker != nil {
		//		params := &bot.SendPhotoParams{
		//			ChatID:  update.Message.Chat.ID,
		//			Photo:   &models.InputFileString{Data: "AgACAgIAAxkDAAIBOWJimnCJHQJiJ4P3aasQCPNyo6mlAALDuzEbcD0YSxzjB-vmkZ6BAQADAgADbQADJAQ"},
		//			Caption: "Preloaded Facebook logo",
		//		}
		//
		//		b.SendPhoto(ctx, params)
		//		file := models.InputFileString{Data: update.Message.Sticker.FileID}
		//		b.SendSticker(ctx, &bot.SendStickerParams{
		//			ChatID:  update.Message.Chat.ID,
		//			Sticker: &file,
		//		})
		//	}
		//}
		//}
	}
}
