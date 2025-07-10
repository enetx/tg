package tg

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

type ChosenInlineResultHandlers struct{ b *Bot }

func (h *ChosenInlineResultHandlers) Any(fn Handler) *ChosenInlineResultHandlers {
	h.b.handleChosenInlineResult(nil, fn)
	return h
}

func (h *ChosenInlineResultHandlers) FromUser(id int64, fn Handler) *ChosenInlineResultHandlers {
	h.b.handleChosenInlineResult(func(cir *gotgbot.ChosenInlineResult) bool {
		return cir != nil && cir.From.Id == id
	}, fn)
	return h
}

func (h *ChosenInlineResultHandlers) Query(query String, fn Handler) *ChosenInlineResultHandlers {
	h.b.handleChosenInlineResult(func(cir *gotgbot.ChosenInlineResult) bool {
		return cir != nil && cir.Query == query.Std()
	}, fn)
	return h
}

func (h *ChosenInlineResultHandlers) QueryPrefix(prefix String, fn Handler) *ChosenInlineResultHandlers {
	h.b.handleChosenInlineResult(func(cir *gotgbot.ChosenInlineResult) bool {
		return cir != nil && String(cir.Query).StartsWith(prefix)
	}, fn)
	return h
}

func (h *ChosenInlineResultHandlers) QuerySuffix(suffix String, fn Handler) *ChosenInlineResultHandlers {
	h.b.handleChosenInlineResult(func(cir *gotgbot.ChosenInlineResult) bool {
		return cir != nil && String(cir.Query).EndsWith(suffix)
	}, fn)
	return h
}

func (h *ChosenInlineResultHandlers) InlineMessage(messageID String, fn Handler) *ChosenInlineResultHandlers {
	h.b.handleChosenInlineResult(func(cir *gotgbot.ChosenInlineResult) bool {
		return cir != nil && cir.InlineMessageId == messageID.Std()
	}, fn)
	return h
}

func (h *ChosenInlineResultHandlers) Location(fn Handler) *ChosenInlineResultHandlers {
	h.b.handleChosenInlineResult(func(cir *gotgbot.ChosenInlineResult) bool {
		return cir != nil && cir.Location != nil
	}, fn)
	return h
}
