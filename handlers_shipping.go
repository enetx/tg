package tg

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

type ShippingHandlers struct{ b *Bot }

func (h *ShippingHandlers) Any(fn Handler) *ShippingHandlers {
	h.b.handleShippingQuery(nil, fn)
	return h
}

func (h *ShippingHandlers) FromUserID(id int64, fn Handler) *ShippingHandlers {
	h.b.handleShippingQuery(func(s *gotgbot.ShippingQuery) bool {
		return s != nil && s.From.Id == id
	}, fn)
	return h
}

func (h *ShippingHandlers) HasPayloadPrefix(prefix String, fn Handler) *ShippingHandlers {
	h.b.handleShippingQuery(func(s *gotgbot.ShippingQuery) bool {
		return s != nil && String(s.InvoicePayload).StartsWith(prefix)
	}, fn)
	return h
}
