package main

import (
	"antispambot/config"
	"context"
	"github.com/go-telegram/bot/models"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
)

var conf *config.Config

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}
	conf = config.GetConfig()

	b, err := bot.New(conf.BotToken, opts...)
	if err != nil {
		panic(err)
	}
	b.SetWebhook(ctx, &bot.SetWebhookParams{
		URL: "https://44f5-136-169-243-85.ngrok.io",
	})

	go func() {
		http.ListenAndServe(":2000", b.WebhookHandler())
	}()

	// Use StartWebhook instead of Start
	b.StartWebhook(ctx)
}

type msg struct {
	UserID int64
	ChatID int64
	Text   string
}

var msgSlice []msg
var counter int

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {

	if !update.Message.From.IsBot {
		sl := msgSlice
		message := msg{
			UserID: update.Message.From.ID,
			ChatID: update.Message.Chat.ID,
			Text:   update.Message.Text,
		}

		msgSlice = append(msgSlice, message)
		sl2 := msgSlice
		counter++

		if len(sl) != 0 && check(sl, sl2) {
			b.DeleteMessage(ctx, &bot.DeleteMessageParams{MessageID: update.Message.ID, ChatID: message.ChatID})

		}
	}
}

func check(sl []msg, sl2 []msg) bool {
	msg1 := sl[len(sl)-1]
	msg2 := sl2[len(sl2)-1]
	//if &sl[len(sl)-1].ID == &sl2[len(sl2)-1].ID && &sl[len(sl)-1].Text == &sl2[len(sl2)-1].Text {
	//	return true
	//}
	if msg2 == msg1 {
		return true
	}
	return false
}
