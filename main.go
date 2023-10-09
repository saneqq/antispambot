package main

import (
	"antispambot/bot/common"
	"antispambot/bot/config"
	bot "antispambot/bot/methods"
	"antispambot/bot/reporter"
	"context"
	tgModels "github.com/go-telegram/bot/models"
	"net/http"
	"os"
	"os/signal"

	tgBot "github.com/go-telegram/bot"
)

var conf *config.Config

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []tgBot.Option{
		tgBot.WithDefaultHandler(handlerFunc),
	}
	conf = config.GetConfig()

	b, err := tgBot.New(conf.BotToken, opts...)
	if err != nil {
		panic(err)
	}
	b.SetWebhook(ctx, &tgBot.SetWebhookParams{
		URL: "ngrok",
	})

	go func() {
		http.ListenAndServe(":2000", b.WebhookHandler())
	}()

	b.StartWebhook(ctx)
}

func handlerFunc(ctx context.Context, b *tgBot.Bot, update *tgModels.Update) {
	conf = config.GetConfig()

	args := &common.HandlerArgs{
		Ctx:    ctx,
		Bot:    b,
		Update: update,
		Conf:   conf,
	}

	if !update.Message.From.IsBot {
		prevMsg, lastMsg := bot.GetPrevAndLastMsg(update)
		err := bot.DeleteMsgOrBanChatMember(args, prevMsg, lastMsg)
		if err != nil {
			reporter.ReportToMe(ctx, b, err.Error(), false)
		}
	}
}
