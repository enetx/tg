package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	"github.com/enetx/g"
	"github.com/enetx/tg/core"
)

// PreCheckoutHandlers provides methods to handle pre-checkout query events.
type PreCheckoutHandlers struct{ Bot core.BotAPI }

// handlePreCheckoutQuery registers a pre-checkout query handler with the dispatcher.
func (h *PreCheckoutHandlers) handlePreCheckoutQuery(f filters.PreCheckoutQuery, fn Handler) {
	h.Bot.Dispatcher().AddHandler(handlers.NewPreCheckoutQuery(f, wrap(h.Bot, middlewares(h.Bot), fn)))
}

// Any handles all pre-checkout queries.
func (h *PreCheckoutHandlers) Any(fn Handler) *PreCheckoutHandlers {
	h.handlePreCheckoutQuery(nil, fn)
	return h
}

// FromUserID handles pre-checkout queries from a specific user.
func (h *PreCheckoutHandlers) FromUserID(id int64, fn Handler) *PreCheckoutHandlers {
	h.handlePreCheckoutQuery(func(p *gotgbot.PreCheckoutQuery) bool {
		return p != nil && p.From.Id == id
	}, fn)
	return h
}

// HasPayloadPrefix handles pre-checkout queries where invoice payload starts with the specified prefix.
func (h *PreCheckoutHandlers) HasPayloadPrefix(prefix g.String, fn Handler) *PreCheckoutHandlers {
	h.handlePreCheckoutQuery(func(p *gotgbot.PreCheckoutQuery) bool {
		return p != nil && g.String(p.InvoicePayload).StartsWith(prefix)
	}, fn)
	return h
}

// Currency handles pre-checkout queries that match the given currency.
func (h *PreCheckoutHandlers) Currency(currency g.String, fn Handler) *PreCheckoutHandlers {
	h.handlePreCheckoutQuery(func(p *gotgbot.PreCheckoutQuery) bool {
		return p != nil && g.String(p.Currency).Eq(currency)
	}, fn)
	return h
}
