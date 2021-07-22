package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func ping(b *gotgbot.Bot, ctx *ext.Context) error {
	ctx.Message.Reply(b, "Pong", nil)
	return nil
}
