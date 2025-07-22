package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	. "github.com/enetx/g"
	"github.com/enetx/tg/core"
)

// ShippingHandlers provides methods to handle shipping query events.
type ShippingHandlers struct{ Bot core.BotAPI }

// handleShippingQuery registers a shipping query handler with the dispatcher.
func (h *ShippingHandlers) handleShippingQuery(f filters.ShippingQuery, fn Handler) {
	h.Bot.Dispatcher().AddHandler(handlers.NewShippingQuery(f, wrap(h.Bot, middlewares(h.Bot), fn)))
}

// Any handles all shipping queries.
func (h *ShippingHandlers) Any(fn Handler) *ShippingHandlers {
	h.handleShippingQuery(nil, fn)
	return h
}

// FromUserID handles shipping queries from a specific user.
func (h *ShippingHandlers) FromUserID(id int64, fn Handler) *ShippingHandlers {
	h.handleShippingQuery(func(s *gotgbot.ShippingQuery) bool {
		return s != nil && s.From.Id == id
	}, fn)
	return h
}

// HasPayloadPrefix handles shipping queries where invoice payload starts with the specified prefix.
func (h *ShippingHandlers) HasPayloadPrefix(prefix String, fn Handler) *ShippingHandlers {
	h.handleShippingQuery(func(s *gotgbot.ShippingQuery) bool {
		return s != nil && String(s.InvoicePayload).StartsWith(prefix)
	}, fn)
	return h
}
