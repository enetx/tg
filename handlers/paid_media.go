package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	. "github.com/enetx/g"
	"github.com/enetx/tg/core"
)

// PaidMediaHandlers provides methods to handle paid media purchase events.
type PaidMediaHandlers struct{ Bot core.BotAPI }

// handlePurchasedPaidMedia registers a paid media purchase handler with the dispatcher.
func (h *PaidMediaHandlers) handlePurchasedPaidMedia(f filters.PurchasedPaidMedia, fn Handler) {
	h.Bot.Dispatcher().AddHandler(handlers.NewPurchasedPaidMedia(f, wrap(h.Bot, middlewares(h.Bot), fn)))
}

// Any handles all paid media purchases.
func (h *PaidMediaHandlers) Any(fn Handler) *PaidMediaHandlers {
	h.handlePurchasedPaidMedia(nil, fn)
	return h
}

// FromUserID handles paid media purchases from a specific user.
func (h *PaidMediaHandlers) FromUserID(id int64, fn Handler) *PaidMediaHandlers {
	h.handlePurchasedPaidMedia(func(pm *gotgbot.PaidMediaPurchased) bool {
		return pm != nil && pm.From.Id == id
	}, fn)
	return h
}

// Payload handles paid media purchases with a specific payload.
func (h *PaidMediaHandlers) Payload(payload String, fn Handler) *PaidMediaHandlers {
	h.handlePurchasedPaidMedia(func(pm *gotgbot.PaidMediaPurchased) bool {
		return pm != nil && pm.PaidMediaPayload == payload.Std()
	}, fn)
	return h
}

// PayloadPrefix handles paid media purchases where payload starts with the specified prefix.
func (h *PaidMediaHandlers) PayloadPrefix(prefix String, fn Handler) *PaidMediaHandlers {
	h.handlePurchasedPaidMedia(func(pm *gotgbot.PaidMediaPurchased) bool {
		return pm != nil && String(pm.PaidMediaPayload).StartsWith(prefix)
	}, fn)
	return h
}
