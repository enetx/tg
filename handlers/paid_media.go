package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"

	"github.com/enetx/tg/core"

	. "github.com/enetx/g"
)

type PaidMediaHandlers struct{ Bot core.BotAPI }

func (h *PaidMediaHandlers) handlePurchasedPaidMedia(f filters.PurchasedPaidMedia, fn Handler) {
	h.Bot.Dispatcher().AddHandler(handlers.NewPurchasedPaidMedia(f, wrap(h.Bot, middlewares(h.Bot), fn)))
}

func (h *PaidMediaHandlers) Any(fn Handler) *PaidMediaHandlers {
	h.handlePurchasedPaidMedia(nil, fn)
	return h
}

func (h *PaidMediaHandlers) FromUserID(id int64, fn Handler) *PaidMediaHandlers {
	h.handlePurchasedPaidMedia(func(pm *gotgbot.PaidMediaPurchased) bool {
		return pm != nil && pm.From.Id == id
	}, fn)
	return h
}

func (h *PaidMediaHandlers) Payload(payload String, fn Handler) *PaidMediaHandlers {
	h.handlePurchasedPaidMedia(func(pm *gotgbot.PaidMediaPurchased) bool {
		return pm != nil && pm.PaidMediaPayload == payload.Std()
	}, fn)
	return h
}

func (h *PaidMediaHandlers) PayloadPrefix(prefix String, fn Handler) *PaidMediaHandlers {
	h.handlePurchasedPaidMedia(func(pm *gotgbot.PaidMediaPurchased) bool {
		return pm != nil && String(pm.PaidMediaPayload).StartsWith(prefix)
	}, fn)
	return h
}
