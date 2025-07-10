package tg

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

type PaidMediaHandlers struct{ b *Bot }

func (h *PaidMediaHandlers) Any(fn Handler) *PaidMediaHandlers {
	h.b.handlePurchasedPaidMedia(nil, fn)
	return h
}

func (h *PaidMediaHandlers) FromUserID(id int64, fn Handler) *PaidMediaHandlers {
	h.b.handlePurchasedPaidMedia(func(pm *gotgbot.PaidMediaPurchased) bool {
		return pm != nil && pm.From.Id == id
	}, fn)
	return h
}

func (h *PaidMediaHandlers) Payload(payload String, fn Handler) *PaidMediaHandlers {
	h.b.handlePurchasedPaidMedia(func(pm *gotgbot.PaidMediaPurchased) bool {
		return pm != nil && pm.PaidMediaPayload == payload.Std()
	}, fn)
	return h
}

func (h *PaidMediaHandlers) PayloadPrefix(prefix String, fn Handler) *PaidMediaHandlers {
	h.b.handlePurchasedPaidMedia(func(pm *gotgbot.PaidMediaPurchased) bool {
		return pm != nil && String(pm.PaidMediaPayload).StartsWith(prefix)
	}, fn)
	return h
}
