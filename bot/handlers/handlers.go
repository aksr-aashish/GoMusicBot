package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"

	"GoMusicBot/bot/handlers/groups"
)

func AddHandlers(dp *ext.Dispatcher) {
	groups.AddHandlers(dp)
}
