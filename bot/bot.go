package bot

import (
	"GoMusicBot/bot/handlers"
	"os"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Start() error {
	b, err := gotgbot.NewBot(os.Getenv("BOT_TOKEN"), nil)
	if err != nil {
		return err
	}

	updater := ext.NewUpdater(nil)

	handlers.AddHandlers(updater.Dispatcher)

	err = updater.StartPolling(b, &ext.PollingOpts{DropPendingUpdates: true})
	if err != nil {
		return err
	}

	updater.Idle()
	return nil
}
