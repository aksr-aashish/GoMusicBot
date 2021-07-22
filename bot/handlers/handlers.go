package handlers

import (
	"GoMusicBot/bot/handlers/groups"

	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func AddHandlers(dp *ext.Dispatcher) {
	dp.AddHandler(handlers.NewCommand("ping", ping))
	groups.AddHandlers(dp)
}
