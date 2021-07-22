package groups

import (
	"GoMusicBot/tgcalls"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
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
		ctx.Message.Reply(b, "Reply an audio file.", nil)
		return nil
	}

	err := downloadFile(b, repliedMessage.Audio.FileId)
	if err != nil {
		ctx.Message.Reply(b, "Error downloading file: "+err.Error(), nil)
		return nil
	}

	err = convert(repliedMessage.Audio.FileId)
	if err != nil {
		ctx.Message.Reply(b, "Error converting file: "+err.Error(), nil)
		return nil
	}

	err = tgcalls.Get().Stream("main", ctx.EffectiveChat.Id, repliedMessage.Audio.FileId+".raw")
	if err != nil {
		ctx.Message.Reply(b, "Error requesting to stream: "+err.Error(), nil)
		return nil
	}

	ctx.Message.Reply(b, "Streaming...", nil)
	return nil
}

var streamHandler = handlers.NewCommand("stream", stream)
