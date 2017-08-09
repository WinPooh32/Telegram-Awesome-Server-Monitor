package main

import "github.com/go-telegram-bot-api/telegram-bot-api"

var(
	inlineKeyboards = make(map [string]tgbotapi.InlineKeyboardMarkup)
)

func MakeMainLayout() {
	btn1 := tgbotapi.NewInlineKeyboardButtonData("🚀Realtime monitor", EVENT_TO_REALTIME)
	btn2 := tgbotapi.NewInlineKeyboardButtonData("📈Stepped monitor", EVENT_TO_STEPPED)
	btn3 := tgbotapi.NewInlineKeyboardButtonData("📜Show last logins", EVENT_TO_LAST)

	row1 := tgbotapi.NewInlineKeyboardRow(btn1, btn2)
	row2 := tgbotapi.NewInlineKeyboardRow(btn3)

	inlineKeyboards["MainLayout"] = tgbotapi.NewInlineKeyboardMarkup(row1, row2)
}

func InitKeyboards(){
	MakeMainLayout()
}