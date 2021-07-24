package groups

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/gotgcalls/gotgcalls"

	"GoMusicBot/i18n"
	"GoMusicBot/tgcalls"
)

func resume(b *gotgbot.Bot, ctx *ext.Context) error {
	switch result, err := tgcalls.Get().Resume(tgcalls.CLIENT, ctx.EffectiveChat.Id); result {
	case gotgcalls.OK:
		_, err = reply(b, ctx.Message, i18n.Localize("resumed", nil))
		return err
	case gotgcalls.NOT_PAUSED:
		_, err = reply(b, ctx.Message, i18n.Localize("not_paused", nil))
		return err
	case gotgcalls.NOT_IN_CALL:
		_, err = reply(b, ctx.Message, i18n.Localize("not_in_call", nil))
		return err
	default:
		if err != nil {
			_, err = reply(b, ctx.Message, i18n.Localize("resume_error", map[string]string{"Error": err.Error()}))
			return err
		}
	}

	return nil
}

var resumeHandler = handlers.NewCommand("resume", resume)
