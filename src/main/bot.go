package main

import (
	"log"
	"bytes"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	EVENT_KILL    = "1"
	EVENT_ACTION  = "2"
)

const (
	REQ_SEND = iota
	REQ_DELETE
	REQ_EDIT
)

const (
	RES_SEND = iota
	RES_DELETE
	RES_EDIT
)

type API_Request struct {
	request int
	chattable tgbotapi.Chattable
	deleteMsgId int
	editMsgId int
	chatId int64
}

type API_Response struct {
	response int
	message tgbotapi.Message
	ok bool
}

type User struct {
	chatID   int64
	msg      chan *tgbotapi.Message
	event    chan *string
	monMsgID int
}

func NewUser() User{
	return User{
		chatID: 0,
		msg: make(chan *tgbotapi.Message),
		event: make(chan *string),
	}
}

func FindUser(users *map[int]*User, userID int) (*User, bool){
	var user *User
	var ok bool

	user, ok = (*users)[userID]

	if !ok {
		tmp := NewUser()
		user = &tmp
		(*users)[userID] = user
	}

	return user, !ok
}

func ServeNewUser(user *User, send chan *API_Request, response chan *API_Response){
	quit := false

	for !quit {
		select{
		case message := <- user.msg:
			println(message.Text)

			switch message.Command() {
			case "start":
				req := API_Request{
					request: REQ_SEND,
					chattable: tgbotapi.NewMessage(user.chatID, "STARTED!"),
				}


				send <- &req
				resp := <- response
				println("MESSAGE ID: ", resp.message.MessageID)


				req = API_Request{
					request: REQ_DELETE,
					deleteMsgId: resp.message.MessageID,
					chatId: user.chatID,
				}

				send <- &req
				resp = <- response

			case "stop":
			}

		case event := <- user.event:
			switch *event {
			case EVENT_KILL:
				//quit = true
				println("EVENT_KILL")
			case EVENT_ACTION:
				println("EVENT_ACTION")
			default:
				log.Println("Unknown bot event recieved: ", *event)
			}
		}
	}
}

func ServeBot(token string, monChan chan *bytes.Buffer, lastChan chan []string){
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	var(
		ch_send_msg = make(chan *API_Request)
		ch_response_msg = make(chan *API_Response)

		users = make(map[int] *User)
	)

	go func(){
		for true {
			select {
				case req := <- ch_send_msg:
					resp := API_Response{}

					switch req.request {
						case REQ_SEND:
							delivered, _ := bot.Send(req.chattable)
							resp = API_Response{message: delivered, response: RES_SEND}

						case REQ_EDIT:

						case REQ_DELETE:
							tgresp, _ := bot.DeleteMessage(tgbotapi.DeleteMessageConfig{
								ChatID: req.chatId,
								MessageID: req.deleteMsgId,
							})

							resp = API_Response{ok: tgresp.Ok, response: RES_DELETE}
					}

					ch_response_msg <- &resp

				case myplot_bytes:= <-monChan:
					for _, v := range users {
						file := tgbotapi.FileBytes{
							Bytes: myplot_bytes.Bytes(),
							Name:  "myplot.png",
						}

						msg := tgbotapi.NewPhotoUpload(v.chatID, file)
						msg.DisableNotification = true

						if v.monMsgID != 0 {
							bot.DeleteMessage(tgbotapi.DeleteMessageConfig{
								ChatID:    v.chatID,
								MessageID: v.monMsgID,
							})
						}

						delivered, _ := bot.Send(msg)
						v.monMsgID = delivered.MessageID
					}

				case last := <- lastChan:
					s := ""

					for i := range last {
						if i != len(last) - 1 {
							s += last[i] + ", "
						}else{
							s += last[i]
						}
					}

					for _, v := range users {
						msg := tgbotapi.NewMessage(v.chatID, "New log in from new ip: " + s + "\nIs it you?")
						msg.DisableNotification = false
						bot.Send(msg)
					}
			}
		}
	}()


	for update := range updates {
		if update.CallbackQuery != nil {
			chatId := update.CallbackQuery.From.ID

			user, isNew := FindUser(&users, chatId)

			if isNew {
				user.chatID = update.CallbackQuery.Message.Chat.ID
				println("CHATID ON CALLBACK: ", user.chatID)
				go ServeNewUser(user, ch_send_msg, ch_response_msg)
			}

			user.event <- &update.CallbackQuery.Data

		} else if update.Message != nil {
			userID := update.Message.From.ID

			user, isNew := FindUser(&users, userID)

			if isNew {
				user.chatID = update.Message.Chat.ID
				println("CHATID ON MESSAGE: ", user.chatID)
				go ServeNewUser(user, ch_send_msg, ch_response_msg)
			}

			user.msg <- update.Message
		}
	}
}