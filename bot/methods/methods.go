package bot

import (
	"antispambot/bot/common"
	"antispambot/bot/models"
	"github.com/go-telegram/bot"
	tgModels "github.com/go-telegram/bot/models"
)

var msgSlice []models.Msg
var counter int

func GetPrevAndLastMsg(update *tgModels.Update) (prevMsg *models.Msg, lastMsg *models.Msg) {
	sl1 := msgSlice
	message := models.Msg{
		UserID: update.Message.From.ID,
		Text:   update.Message.Text,
	}

	msgSlice = append(msgSlice, message)
	sl2 := msgSlice
	if len(sl1) != 0 {
		return &sl1[len(sl1)-1], &sl2[len(sl2)-1]
	} else {
		return nil, &sl2[len(sl2)-1]
	}

}

func DeleteMsgOrBanChatMember(args *common.HandlerArgs, prevMsg *models.Msg, lastMsg *models.Msg) (err error) {
	if check(prevMsg, lastMsg) {
		_, err = args.Bot.DeleteMessage(args.Ctx, &bot.DeleteMessageParams{
			ChatID:    args.Update.Message.Chat.ID,
			MessageID: args.Update.Message.ID,
		})
		if err != nil {
			return err
		}
		counter++
		if counter == 2 {
			//_, err = args.Bot.BanChatMember(args.Ctx, &bot.BanChatMemberParams{
			//	ChatID: args.Update.Message.Chat.ID,
			//	UserID: args.Update.Message.From.ID,
			//})
			//if err != nil {
			//	return err
			//}
			counter = 0
		}
	} else {
		counter = 0
	}
	return nil
}

func check(prevMsg *models.Msg, lastMsg *models.Msg) bool {
	if prevMsg != nil && prevMsg.UserID == lastMsg.UserID && prevMsg.Text == lastMsg.Text {
		return true
	}
	return false
}