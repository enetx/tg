package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"

	"github.com/enetx/tg/core"

	. "github.com/enetx/g"
)

type ShippingHandlers struct{ Bot core.BotAPI }

func (h *ShippingHandlers) handleShippingQuery(f filters.ShippingQuery, fn Handler) {
	h.Bot.Dispatcher().AddHandler(handlers.NewShippingQuery(f, wrap(h.Bot, middlewares(h.Bot), fn)))
}

func (h *ShippingHandlers) Any(fn Handler) *ShippingHandlers {
	h.handleShippingQuery(nil, fn)
	return h
}

func (h *ShippingHandlers) FromUserID(id int64, fn Handler) *ShippingHandlers {
	h.handleShippingQuery(func(s *gotgbot.ShippingQuery) bool {
		return s != nil && s.From.Id == id
	}, fn)
	return h
}

func (h *ShippingHandlers) HasPayloadPrefix(prefix String, fn Handler) *ShippingHandlers {
	h.handleShippingQuery(func(s *gotgbot.ShippingQuery) bool {
		return s != nil && String(s.InvoicePayload).StartsWith(prefix)
	}, fn)
	return h
}
