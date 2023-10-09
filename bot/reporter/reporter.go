package reporter

import (
	"antispambot/bot/constants"
	"context"
	"fmt"
	tgBot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/rs/zerolog/log"
	"os"
	"runtime/debug"
)

func ReportToMe(ctx context.Context, b *tgBot.Bot, text string, markdown bool) {
	msg := &tgBot.SendMessageParams{
		ChatID: constants.SaneQTgID,
		Text:   text,
	}
	if markdown {
		msg.ParseMode = "Markdown"
	}

	_, err := b.SendMessage(ctx, msg)
	if err != nil {
		log.Error().Err(err)
	}

}

func HandlePanic(ctx context.Context, b *tgBot.Bot, update *models.Update) {
	if r := recover(); r != nil {
		text := ""
		from := ""
		if update.Message != nil {
			text = update.Message.Text
			from = update.Message.From.Username
		}
		stack := string(debug.Stack())
		pErr := fmt.Sprintf("panic: %v\n%s\nMessage text: %s\nUser info: %v",
			r, stack, text, from)
		fmt.Fprint(os.Stderr, pErr)
		ReportToMe(ctx, b, pErr, false)
	}
}
