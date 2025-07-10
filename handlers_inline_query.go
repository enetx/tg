package tg

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

type InlineQueryHandlers struct{ bot *Bot }

func (h *InlineQueryHandlers) Any(fn Handler) *InlineQueryHandlers {
	h.bot.handleInlineQuery(nil, fn)
	return h
}

func (h *InlineQueryHandlers) FromUser(id int64, fn Handler) *InlineQueryHandlers {
	h.bot.handleInlineQuery(func(iq *gotgbot.InlineQuery) bool {
		return iq != nil && iq.From.Id == id
	}, fn)
	return h
}

func (h *InlineQueryHandlers) Query(query String, fn Handler) *InlineQueryHandlers {
	h.bot.handleInlineQuery(func(iq *gotgbot.InlineQuery) bool {
		return iq != nil && iq.Query == query.Std()
	}, fn)
	return h
}

func (h *InlineQueryHandlers) QueryPrefix(prefix String, fn Handler) *InlineQueryHandlers {
	h.bot.handleInlineQuery(func(iq *gotgbot.InlineQuery) bool {
		return iq != nil && String(iq.Query).StartsWith(prefix)
	}, fn)
	return h
}

func (h *InlineQueryHandlers) QuerySuffix(suffix String, fn Handler) *InlineQueryHandlers {
	h.bot.handleInlineQuery(func(iq *gotgbot.InlineQuery) bool {
		return iq != nil && String(iq.Query).EndsWith(suffix)
	}, fn)
	return h
}

func (h *InlineQueryHandlers) Location(fn Handler) *InlineQueryHandlers {
	h.bot.handleInlineQuery(func(iq *gotgbot.InlineQuery) bool {
		return iq != nil && iq.Location != nil
	}, fn)
	return h
}

func (h *InlineQueryHandlers) Sender(fn Handler) *InlineQueryHandlers {
	h.bot.handleInlineQuery(func(iq *gotgbot.InlineQuery) bool {
		return iq != nil && iq.ChatType == "sender"
	}, fn)
	return h
}

func (h *InlineQueryHandlers) Private(fn Handler) *InlineQueryHandlers {
	h.bot.handleInlineQuery(func(iq *gotgbot.InlineQuery) bool {
		return iq != nil && iq.ChatType == "private"
	}, fn)
	return h
}

func (h *InlineQueryHandlers) Group(fn Handler) *InlineQueryHandlers {
	h.bot.handleInlineQuery(func(iq *gotgbot.InlineQuery) bool {
		return iq != nil && iq.ChatType == "group"
	}, fn)
	return h
}

func (h *InlineQueryHandlers) Supergroup(fn Handler) *InlineQueryHandlers {
	h.bot.handleInlineQuery(func(iq *gotgbot.InlineQuery) bool {
		return iq != nil && iq.ChatType == "supergroup"
	}, fn)
	return h
}

func (h *InlineQueryHandlers) Channel(fn Handler) *InlineQueryHandlers {
	h.bot.handleInlineQuery(func(iq *gotgbot.InlineQuery) bool {
		return iq != nil && iq.ChatType == "channel"
	}, fn)
	return h
}
