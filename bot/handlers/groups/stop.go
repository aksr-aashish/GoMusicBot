package groups

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/gotgcalls/gotgcalls"

	"GoMusicBot/i18n"
	"GoMusicBot/queues"
	"GoMusicBot/tgcalls"
)

func stop(b *gotgbot.Bot, ctx *ext.Context) error {
	switch result, _ := tgcalls.GoTGCalls.Stop(tgcalls.CLIENT, ctx.EffectiveChat.Id); result {
	case gotgcalls.OK:
		queues.Clear(ctx.EffectiveChat.Id)
		ctx.Message.Reply(b, i18n.Localize("stopped", nil), nil)
		return nil
	default:
		ctx.Message.Reply(b, i18n.Localize("not_stopped", nil), nil)
	}

	return nil
}

var stopHandler = handlers.NewCommand("stop", stop)
