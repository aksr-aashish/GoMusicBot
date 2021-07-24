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
		reply(b, ctx.Message, i18n.Localize("skipped", nil))
	} else {
		reply(b, ctx.Message, i18n.Localize("not_streaming", nil))
	}

	return nil
}

var skipHandler = handlers.NewCommand("skip", skip)
