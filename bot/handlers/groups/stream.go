package groups

import (
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/gotgcalls/gotgcalls"

	"GoMusicBot/i18n"
	"GoMusicBot/queues"
	"GoMusicBot/tgcalls"
)

func downloadFile(b *gotgbot.Bot, fileId string) error {
	if _, err := os.Stat(fileId); err == nil {
		return nil
	}

	file, err := b.GetFile(fileId)
	if err != nil {
		return err
	}

	downloadUrl := "https://api.telegram.org/file/bot" + b.Token + "/" + file.FilePath
	res, err := http.Get(downloadUrl)
	if err != nil {
		return err
	}

	file2, err := os.Create(fileId)
	if err != nil {
		return err
	}

	defer file2.Close()
	_, err = io.Copy(file2, res.Body)
	return err
}

func getFFmpegArgs(input string, output string) []string {
	return []string{"-y", "-i", input, "-c", "copy", "-acodec", "pcm_s16le", "-f", "s16le", "-ac", "1", "-ar", "65000", output}
}

func convert(input string) error {
	output := input + ".raw"
	if _, err := os.Stat(output); err == nil {
		return nil
	}

	cmd := exec.Command("ffmpeg", getFFmpegArgs(input, output)...)
	err := cmd.Start()
	if err != nil {
		return err
	}

	return cmd.Wait()
}

func stream(b *gotgbot.Bot, ctx *ext.Context) error {
	repliedMessage := ctx.EffectiveMessage.ReplyToMessage
	if repliedMessage == nil || repliedMessage.Audio == nil {
		ctx.Message.Reply(b, i18n.Localize("require_audio_file", nil), nil)
		return nil
	}

	msg, err := ctx.Message.Reply(b, i18n.Localize("downloading", nil), nil)
	if err != nil {
		return err
	}

	err = downloadFile(b, repliedMessage.Audio.FileId)
	if err != nil {
		msg.EditText(b, i18n.Localize("download_error", nil), nil)
		return nil
	}

	msg, err = msg.EditText(b, i18n.Localize("converting", nil), nil)
	if err != nil {
		return nil
	}

	err = convert(repliedMessage.Audio.FileId)
	if err != nil {
		msg.EditText(b, i18n.Localize("convert_error", map[string]string{"Error": err.Error()}), nil)
		return nil
	}

	if isFinished, _ := tgcalls.Get().IsFinished(tgcalls.CLIENT, ctx.EffectiveChat.Id); isFinished != gotgcalls.NOT_FINISHED {
		err = tgcalls.Get().Stream(tgcalls.CLIENT, ctx.EffectiveChat.Id, repliedMessage.Audio.FileId+".raw")
		if err != nil {
			msg.EditText(b, i18n.Localize("stream_error", map[string]string{"Error": err.Error()}), nil)
			return nil
		}

		msg.EditText(b, i18n.Localize("streaming", nil), nil)
	} else {
		position := queues.Push(ctx.EffectiveChat.Id, repliedMessage.Audio.FileId+".raw")
		msg.EditText(b, i18n.Localize("queued_at", map[string]int{"Position": position}), nil)
	}
	return nil
}

var streamHandler = handlers.NewCommand("stream", stream)
