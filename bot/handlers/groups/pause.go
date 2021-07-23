package groups

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/gotgcalls/gotgcalls"

	"GoMusicBot/i18n"
	"GoMusicBot/tgcalls"
)

func pause(b *gotgbot.Bot, ctx *ext.Context) error {
	switch result, err := tgcalls.Get().Pause(tgcalls.CLIENT, ctx.EffectiveChat.Id); result {
	case gotgcalls.OK:
		_, err = ctx.Message.Reply(b, i18n.Localize("paused", nil), nil)
		return err
	case gotgcalls.NOT_STREAMING:
		_, err := ctx.Message.Reply(b, i18n.Localize("not_streaming_to_pause", nil), nil)
		return err
	case gotgcalls.NOT_IN_CALL:
		_, err = ctx.Message.Reply(b, i18n.Localize("not_in_call", nil), nil)
		return err
	default:
		if err != nil {
			_, err = ctx.Message.Reply(b, i18n.Localize("pause_error", map[string]string{"Error": err.Error()}), nil)
			return err
		}
	}

	return nil
}

var pauseHandler = handlers.NewCommand("pause", pause)
