package main

import (
	"log"
	"bytes"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func ServeBot(token string, monChan chan *bytes.Buffer){
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	var id = 0
	var started = false

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if !started {
			started = true;

			go func() {
				//fixme no exit condition
				for true {
					myplot_bytes := <-monChan

					file := tgbotapi.FileBytes{
						Bytes: myplot_bytes.Bytes(),
						Name:  "myplot.png",
					}

					msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, file)
					msg.DisableNotification = true
					//msg.Caption = "Проба пера"

					if id != 0 {
						bot.DeleteMessage(tgbotapi.DeleteMessageConfig{
							ChatID:    update.Message.Chat.ID,
							MessageID: id,
						})
					}

					delivered, _ := bot.Send(msg)
					id = delivered.MessageID

					//reset remote image buffer
					myplot_bytes.Reset()
				}
			}()
		} //!started
	}
}