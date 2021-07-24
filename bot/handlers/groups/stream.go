package groups

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/gotgcalls/gotgcalls"

	"GoMusicBot/converter"
	"GoMusicBot/downloader"
	"GoMusicBot/i18n"
	"GoMusicBot/queues"
	"GoMusicBot/tgcalls"
)

func stream(b *gotgbot.Bot, ctx *ext.Context) error {
	repliedMessage := ctx.EffectiveMessage.ReplyToMessage
	if repliedMessage == nil || repliedMessage.Audio == nil {
		ctx.Message.Reply(b, i18n.Localize("require_audio_file", nil), nil)
		return nil
	}

	msg, err := repliedMessage.Reply(b, i18n.Localize("downloading", nil), nil)
	if err != nil {
		return err
	}

	input, err := downloader.Download(b, repliedMessage.Audio.FileId)
	if err != nil {
		msg.EditText(b, i18n.Localize("download_error", nil), nil)
		return nil
	}

	msg, err = msg.EditText(b, i18n.Localize("converting", nil), nil)
	if err != nil {
		return nil
	}

	filePath, err := converter.Convert(input)
	if err != nil {
		msg.EditText(b, i18n.Localize("convert_error", map[string]string{"Error": err.Error()}), nil)
		return nil
	}

	if isFinished, _ := tgcalls.Get().IsFinished(tgcalls.CLIENT, ctx.EffectiveChat.Id); isFinished != gotgcalls.NOT_FINISHED {
		err = tgcalls.Get().Stream(tgcalls.CLIENT, ctx.EffectiveChat.Id, filePath)
		if err != nil {
			msg.EditText(b, i18n.Localize("stream_error", map[string]string{"Error": err.Error()}), nil)
			return nil
		}

		msg.EditText(b, i18n.Localize("streaming", nil), nil)
	} else {
		position := queues.Push(ctx.EffectiveChat.Id, filePath)
		msg.EditText(b, i18n.Localize("queued_at", map[string]int{"Position": position}), nil)
	}
	return nil
}

var streamHandler = handlers.NewCommand("stream", stream)
