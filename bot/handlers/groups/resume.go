package groups

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/gotgcalls/gotgcalls"

	"GoMusicBot/tgcalls"
)

func resume(b *gotgbot.Bot, ctx *ext.Context) error {
	switch result, err := tgcalls.Get().Resume("main", ctx.EffectiveChat.Id); result {
	case gotgcalls.OK:
		_, err = ctx.Message.Reply(b, "Resumed.", nil)
		return err
	case gotgcalls.NOT_STREAMING: // #
		_, err = ctx.Message.Reply(b, "Not paused.", nil)
		return err
	case gotgcalls.NOT_IN_CALL:
		_, err = ctx.Message.Reply(b, "Not in call.", nil)
		return err
	default:
		if err != nil {
			_, err = ctx.Message.Reply(b, "Error resuming: "+err.Error(), nil)
			return err
		} else {
			_, err = ctx.Message.Reply(b, "Could not resume.", nil)
			return err
		}
	}
}

var resumeHandler = handlers.NewCommand("resume", resume)
