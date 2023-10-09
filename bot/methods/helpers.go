package bot

import (
	"antispambot/bot/common"
	"antispambot/bot/reporter"
	"fmt"
	tgModels "github.com/go-telegram/bot/models"
	"strconv"
	"strings"
)

func getFLUName(user *tgModels.User) string {
	return getFirstnameLastnameUsername(user.FirstName, user.LastName, user.Username)
}

func getFirstnameLastnameUsername(firstname string, lastname string, username string) string {
	if firstname != "" || lastname != "" {
		if firstname != "" {
			firstname = firstname + " "
		}
		if lastname != "" {
			lastname = lastname + " "
		}
	} else if firstname == "" && lastname == "" {
		if username != "" {
			username = username + " "
		}
	}

	if (firstname != "" || lastname != "") && username != "" {
		username = "(" + username + ") "
	}
	return firstname + lastname + username
}

func CreateMentionUserText(mentionUserText string, userID int64) string {
	userIDStr := strconv.FormatInt(userID, 10)
	return "[" + strings.TrimSpace(mentionUserText) + "](tg://user?id=" + userIDStr + ")"
}

func ReportToMeWithMention(args *common.HandlerArgs) {
	var fluName string
	var mentionUser string
	fluName = getFLUName(args.Update.Message.From)
	mentionUser = CreateMentionUserText(fluName, args.Update.Message.From.ID)
	msgText := fmt.Sprintf("Banned %s", mentionUser)
	reporter.ReportToMe(args.Ctx, args.Bot, msgText, true)
}
