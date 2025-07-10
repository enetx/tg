package tg

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

type PreCheckoutHandlers struct{ b *Bot }

func (h *PreCheckoutHandlers) Any(fn Handler) *PreCheckoutHandlers {
	h.b.handlePreCheckoutQuery(nil, fn)
	return h
}

func (h *PreCheckoutHandlers) FromUserID(id int64, fn Handler) *PreCheckoutHandlers {
	h.b.handlePreCheckoutQuery(func(p *gotgbot.PreCheckoutQuery) bool {
		return p != nil && p.From.Id == id
	}, fn)
	return h
}

func (h *PreCheckoutHandlers) HasPayloadPrefix(prefix String, fn Handler) *PreCheckoutHandlers {
	h.b.handlePreCheckoutQuery(func(p *gotgbot.PreCheckoutQuery) bool {
		return p != nil && String(p.InvoicePayload).StartsWith(prefix)
	}, fn)
	return h
}
