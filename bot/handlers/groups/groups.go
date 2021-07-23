package groups

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func AddHandlers(dp *ext.Dispatcher) {
	handlers2 := []ext.Handler{pauseHandler, resumeHandler, streamHandler, skipHandler}

	dp.AddHandler(
		handlers.NewMessage(
			func(msg *gotgbot.Message) bool {
				if msg.Chat.Type == "group" || msg.Chat.Type == "supergroup" {
					return true
				}

				return false
			},
			func(b *gotgbot.Bot, ctx *ext.Context) error {
				for _, handler := range handlers2 {
					if handler.CheckUpdate(b, ctx.Update) {
						return handler.HandleUpdate(b, ctx)
					}
				}

				return nil
			},
		),
	)
}
