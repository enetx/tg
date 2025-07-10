package tg

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type namedHandler struct {
	name string
	ext.Handler
}

func (n namedHandler) Name() string {
	return n.name
}

type Handler func(*Context) error

func wrap(b *Bot, handler Handler) func(*gotgbot.Bot, *ext.Context) error {
	return func(_ *gotgbot.Bot, rctx *ext.Context) error {
		ctx := newCtx(b, rctx)

		final := handler
		for i := len(b.middlewares) - 1; i >= 0; i-- {
			mw := b.middlewares[i]
			next := final
			final = func(c *Context) error {
				if err := mw(c); err != nil {
					return err
				}

				return next(c)
			}
		}

		return final(ctx)
	}
}
