package groups

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"

	"GoMusicBot/i18n"
	"GoMusicBot/tgcalls"
)

func skip(b *gotgbot.Bot, ctx *ext.Context) error {
	if tgcalls.OnFinish(tgcalls.CLIENT, ctx.EffectiveChat.Id) {
		ctx.Message.Reply(b, i18n.Localize("skipped", nil), nil)
	} else {
		ctx.Message.Reply(b, i18n.Localize("not_streaming", nil), nil)
	}

	return nil
}

var skipHandler = handlers.NewCommand("skip", skip)
