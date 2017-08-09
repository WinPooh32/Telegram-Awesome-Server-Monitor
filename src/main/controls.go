package main

import "github.com/go-telegram-bot-api/telegram-bot-api"

var(
	inlineKeyboards = make(map [string]tgbotapi.InlineKeyboardMarkup)
)

func MakeMainLayout() {
	btn1 := tgbotapi.NewInlineKeyboardButtonData("🚀Realtime monitor", EVENT_TO_REALTIME)
	btn2 := tgbotapi.NewInlineKeyboardButtonData("📈Stepped monitor", EVENT_TO_STEPPED)
	btn3 := tgbotapi.NewInlineKeyboardButtonData("📜Show last logins", EVENT_TO_LAST)
	btn4 := tgbotapi.NewInlineKeyboardButtonData("⚙️Settings", EVENT_TO_SETTINGS)

	row1 := tgbotapi.NewInlineKeyboardRow(btn1, btn2)
	row2 := tgbotapi.NewInlineKeyboardRow(btn3, btn4)

	inlineKeyboards["MainLayout"] = tgbotapi.NewInlineKeyboardMarkup(row1, row2)
}

func MakeUpdateLayout(){
	btn1 := tgbotapi.NewInlineKeyboardButtonData("🏠Home", EVENT_TO_MAIN)
	btn2 := tgbotapi.NewInlineKeyboardButtonData("🔄Refresh", EVENT_REFRESH)

	row1 := tgbotapi.NewInlineKeyboardRow(btn1, btn2)

	inlineKeyboards["Update"] = tgbotapi.NewInlineKeyboardMarkup(row1)
}

func MakeGoHomeLayout(){
	btn1 := tgbotapi.NewInlineKeyboardButtonData("🏠Home", EVENT_TO_MAIN)
	row1 := tgbotapi.NewInlineKeyboardRow(btn1)
	inlineKeyboards["Home"] = tgbotapi.NewInlineKeyboardMarkup(row1)
}

func MakeSettingsLayout(){
	btn1 := tgbotapi.NewInlineKeyboardButtonData("3 sec", EVENT_SET_3)
	btn2 := tgbotapi.NewInlineKeyboardButtonData("5 sec", EVENT_SET_5)
	btn3 := tgbotapi.NewInlineKeyboardButtonData("10 sec", EVENT_SET_10)

	btn4 := tgbotapi.NewInlineKeyboardButtonData("🔙Cancel", EVENT_TO_MAIN)

	row1 := tgbotapi.NewInlineKeyboardRow(btn1, btn2, btn3)
	row2 := tgbotapi.NewInlineKeyboardRow(btn4)

	inlineKeyboards["Settings"] = tgbotapi.NewInlineKeyboardMarkup(row1, row2)
}

func InitKeyboards(){
	MakeMainLayout()
	MakeUpdateLayout()
	MakeGoHomeLayout()
	MakeSettingsLayout()
}