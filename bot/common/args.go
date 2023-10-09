package common

import (
	"antispambot/bot/config"
	"context"
	tgBot "github.com/go-telegram/bot"
	tgModels "github.com/go-telegram/bot/models"
)

type HandlerArgs struct {
	Ctx    context.Context
	Bot    *tgBot.Bot
	Update *tgModels.Update
	Conf   *config.Config
}
