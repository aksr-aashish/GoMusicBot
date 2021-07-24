package downloader

import (
	"io"
	"net/http"
	"os"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

const DOWNLOADED_FILES_DIR = "downloads/"

func Download(b *gotgbot.Bot, fileId string) (string, error) {
	output := DOWNLOADED_FILES_DIR + fileId

	if _, err := os.Stat(output); err == nil {
		return output, nil
	}

	file, err := b.GetFile(fileId)
	if err != nil {
		return "", err
	}

	downloadUrl := "https://api.telegram.org/file/bot" + b.Token + "/" + file.FilePath
	res, err := http.Get(downloadUrl)
	if err != nil {
		return "", err
	}

	file2, err := os.Create(output)
	if err != nil {
		return "", err
	}

	defer file2.Close()
	_, err = io.Copy(file2, res.Body)
	return output, err
}
