package app

import (
	"fmt"
	"net/http"
	"strings"

	. "github.com/line/line-bot-sdk-go/linebot"
	"github.com/zkrhm/ja-bot/bot_fn"

	"go.uber.org/zap"
)

func (app *App) handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

func (app *App) handleLine(w http.ResponseWriter, r *http.Request) {
	// data, err := ioutil.ReadAll(r.Body)
	// defer r.Body.Close()
	// if err != nil {
	// 	panic(err)
	// }

	events, err := app.bot.ParseRequest(r)
	if err != nil {

		panic(err)
	}
	for _, event := range events {
		app.logger.Debug("received event: ", zap.Any("event", event))

	}
	// app.logger.Debug(string(data))
}

func (app *App) handleMessage(e *Event) {
	token := e.ReplyToken

	switch e.Message.(type) {
	case *TextMessage:
		msg := e.Message.(*TextMessage)
		chatContent := msg.Text

		c1 := strings.Contains(chatContent, "kerjaan")
		c2 := strings.Contains(chatContent, "pekerjaan")
		c3 := strings.Contains(chatContent, "daftar kerjaan")

		if c1 || c2 || c3 {

			jobs := []SendingMessage{
				NewTextMessage("Hi kak, berikut daftar kerjaan buat kaka"),
				bot_fn.NewJobOffer("Dibutuhkan 5 Orang Usher untuk promo PT. Mayora - Lippo BSD 5 November 2012 - 3jt per hari "),
				bot_fn.NewJobOffer("Dibutuhkan 5 Orang SPG untuk promo Produk Honda - Lippo BSD 5 November 2012 - 5jt per hari "),
			}
			app.bot.ReplyMessage(token, jobs...)
		}

		return
	default:
		app.bot.ReplyMessage(token, NewTextMessage("Kakak dummy gak ngerti permintaan kira, maapin dummy :("))
	}

}

func (app *App) handleFollow(e *Event) {
	app.logger.Info("I am get followed by: ")
}

func (app *App) handleUnfollow(e *Event) {
	app.logger.Info("I am get unfollowed")
}

func (app *App) handleJoin(e *Event) {
	app.logger.Info("I am joining...")
}

func (app *App) handleLeave(e *Event) {
	app.logger.Info("I am leaving...")
}

func (app *App) handleEvent(e *Event) {

	eventTypeHandlers := map[EventType]interface{}{
		EventTypeMessage: app.handleMessage,
	}

	handler := eventTypeHandlers[e.Type].(func(*Event))
	handler(e)
}
