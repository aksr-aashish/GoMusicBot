package groups

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
)

func reply(b *gotgbot.Bot, msg *gotgbot.Message, text string) (*gotgbot.Message, error) {
	return msg.Reply(b, text, &gotgbot.SendMessageOpts{ParseMode: "HTML"})
}

func edit(b *gotgbot.Bot, msg *gotgbot.Message, text string) (*gotgbot.Message, error) {
	return msg.Reply(b, text, &gotgbot.SendMessageOpts{ParseMode: "HTML"})
}
