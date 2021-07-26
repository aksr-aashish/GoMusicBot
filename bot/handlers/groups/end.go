package groups

import (
	"GoMusicBot/queues"
	"GoMusicBot/tgcalls"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func end(b *gotgbot.Bot, ctx *ext.Context) error {
	queues.Clear(ctx.EffectiveChat.Id)
	tgcalls.GoTGCalls.Stop(tgcalls.CLIENT, ctx.EffectiveChat.Id)
	return nil
}

var endHandler = handlers.NewMessage(
	func(msg *gotgbot.Message) bool {
		return msg.VoiceChatEnded != nil
	},
	end,
)
