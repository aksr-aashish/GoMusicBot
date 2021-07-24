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
		reply(b, ctx.Message, i18n.Localize("require_audio_file", nil))
		return nil
	}

	msg, err := reply(b, repliedMessage, i18n.Localize("downloading", nil))
	if err != nil {
		return err
	}

	input, err := downloader.Download(b, repliedMessage.Audio.FileId)
	if err != nil {
		edit(b, msg, i18n.Localize("download_error", nil))
		return nil
	}

	msg, err = edit(b, msg, i18n.Localize("converting", nil))
	if err != nil {
		return nil
	}

	filePath, err := converter.Convert(input)
	if err != nil {
		edit(b, msg, i18n.Localize("convert_error", map[string]string{"Error": err.Error()}))
		return nil
	}

	if isFinished, _ := tgcalls.Get().IsFinished(tgcalls.CLIENT, ctx.EffectiveChat.Id); isFinished != gotgcalls.NOT_FINISHED {
		err = tgcalls.Get().Stream(tgcalls.CLIENT, ctx.EffectiveChat.Id, filePath)
		if err != nil {
			edit(b, msg, i18n.Localize("stream_error", map[string]string{"Error": err.Error()}))
			return nil
		}

		edit(b, msg, i18n.Localize("streaming", nil))
	} else {
		position := queues.Push(ctx.EffectiveChat.Id, filePath)
		edit(b, msg, i18n.Localize("queued_at", map[string]int{"Position": position}))
	}
	return nil
}

var streamHandler = handlers.NewCommand("stream", stream)
