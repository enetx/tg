package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	. "github.com/enetx/g"
	"github.com/enetx/tg/core"
)

// ChosenInlineResultHandlers provides methods to handle chosen inline result events.
type ChosenInlineResultHandlers struct{ Bot core.BotAPI }

func (h *ChosenInlineResultHandlers) handleChosenInlineResult(f filters.ChosenInlineResult, fn Handler) {
	h.Bot.Dispatcher().AddHandler(handlers.NewChosenInlineResult(f, wrap(h.Bot, middlewares(h.Bot), fn)))
}

// Any handles all chosen inline results.
func (h *ChosenInlineResultHandlers) Any(fn Handler) *ChosenInlineResultHandlers {
	h.handleChosenInlineResult(nil, fn)
	return h
}

// FromUser handles chosen inline results from a specific user.
func (h *ChosenInlineResultHandlers) FromUser(id int64, fn Handler) *ChosenInlineResultHandlers {
	h.handleChosenInlineResult(func(cir *gotgbot.ChosenInlineResult) bool {
		return cir != nil && cir.From.Id == id
	}, fn)
	return h
}

// Query handles chosen inline results with a specific query string.
func (h *ChosenInlineResultHandlers) Query(query String, fn Handler) *ChosenInlineResultHandlers {
	h.handleChosenInlineResult(func(cir *gotgbot.ChosenInlineResult) bool {
		return cir != nil && cir.Query == query.Std()
	}, fn)
	return h
}

// QueryPrefix handles chosen inline results where query starts with the specified prefix.
func (h *ChosenInlineResultHandlers) QueryPrefix(prefix String, fn Handler) *ChosenInlineResultHandlers {
	h.handleChosenInlineResult(func(cir *gotgbot.ChosenInlineResult) bool {
		return cir != nil && String(cir.Query).StartsWith(prefix)
	}, fn)
	return h
}

// QuerySuffix handles chosen inline results where query ends with the specified suffix.
func (h *ChosenInlineResultHandlers) QuerySuffix(suffix String, fn Handler) *ChosenInlineResultHandlers {
	h.handleChosenInlineResult(func(cir *gotgbot.ChosenInlineResult) bool {
		return cir != nil && String(cir.Query).EndsWith(suffix)
	}, fn)
	return h
}

// InlineMessage handles chosen inline results with a specific inline message ID.
func (h *ChosenInlineResultHandlers) InlineMessage(messageID String, fn Handler) *ChosenInlineResultHandlers {
	h.handleChosenInlineResult(func(cir *gotgbot.ChosenInlineResult) bool {
		return cir != nil && cir.InlineMessageId == messageID.Std()
	}, fn)
	return h
}

// Location handles chosen inline results that include location data.
func (h *ChosenInlineResultHandlers) Location(fn Handler) *ChosenInlineResultHandlers {
	h.handleChosenInlineResult(func(cir *gotgbot.ChosenInlineResult) bool {
		return cir != nil && cir.Location != nil
	}, fn)
	return h
}
