package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	. "github.com/enetx/g"
	"github.com/enetx/tg/core"
	"github.com/enetx/tg/ctx"
)

// Handler is a function type that handles bot events and returns an error if processing fails.
type Handler func(*ctx.Context) error

func wrap(bot core.BotAPI, middlewares Slice[Handler], handler Handler) func(*gotgbot.Bot, *ext.Context) error {
	return func(_ *gotgbot.Bot, ectx *ext.Context) error {
		c := ctx.New(bot, ectx)

		final := handler
		for i := len(middlewares) - 1; i >= 0; i-- {
			mw := middlewares[i]
			next := final
			final = func(c *ctx.Context) error {
				if err := mw(c); err != nil {
					return err
				}
				return next(c)
			}
		}

		return final(c)
	}
}

func middlewares(api core.BotAPI) []Handler {
	if b, ok := api.(interface{ Middlewares() []Handler }); ok {
		return b.Middlewares()
	}

	return nil
}
