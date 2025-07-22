package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	. "github.com/enetx/g"
	"github.com/enetx/tg/core"
	"github.com/enetx/tg/types/chat"
)

// InlineQueryHandlers provides methods to handle inline query events.
type InlineQueryHandlers struct{ Bot core.BotAPI }

// handleInlineQuery registers an inline query handler with the dispatcher.
func (h *InlineQueryHandlers) handleInlineQuery(f filters.InlineQuery, fn Handler) {
	h.Bot.Dispatcher().AddHandler(handlers.NewInlineQuery(f, wrap(h.Bot, middlewares(h.Bot), fn)))
}

// Any handles all inline queries.
func (h *InlineQueryHandlers) Any(fn Handler) *InlineQueryHandlers {
	h.handleInlineQuery(nil, fn)
	return h
}

// FromUser handles inline queries from a specific user.
func (h *InlineQueryHandlers) FromUser(id int64, fn Handler) *InlineQueryHandlers {
	h.handleInlineQuery(func(iq *gotgbot.InlineQuery) bool {
		return iq != nil && iq.From.Id == id
	}, fn)
	return h
}

// Query handles inline queries with a specific query string.
func (h *InlineQueryHandlers) Query(query String, fn Handler) *InlineQueryHandlers {
	h.handleInlineQuery(func(iq *gotgbot.InlineQuery) bool {
		return iq != nil && iq.Query == query.Std()
	}, fn)
	return h
}

// QueryPrefix handles inline queries where query starts with the specified prefix.
func (h *InlineQueryHandlers) QueryPrefix(prefix String, fn Handler) *InlineQueryHandlers {
	h.handleInlineQuery(func(iq *gotgbot.InlineQuery) bool {
		return iq != nil && String(iq.Query).StartsWith(prefix)
	}, fn)
	return h
}

// QuerySuffix handles inline queries where query ends with the specified suffix.
func (h *InlineQueryHandlers) QuerySuffix(suffix String, fn Handler) *InlineQueryHandlers {
	h.handleInlineQuery(func(iq *gotgbot.InlineQuery) bool {
		return iq != nil && String(iq.Query).EndsWith(suffix)
	}, fn)
	return h
}

// Location handles inline queries that include location data.
func (h *InlineQueryHandlers) Location(fn Handler) *InlineQueryHandlers {
	h.handleInlineQuery(func(iq *gotgbot.InlineQuery) bool {
		return iq != nil && iq.Location != nil
	}, fn)
	return h
}

// Sender handles inline queries from sender chat type.
func (h *InlineQueryHandlers) Sender(fn Handler) *InlineQueryHandlers {
	h.handleInlineQuery(func(iq *gotgbot.InlineQuery) bool {
		return iq != nil && iq.ChatType == chat.Sender.String()
	}, fn)
	return h
}

// Private handles inline queries from private chats.
func (h *InlineQueryHandlers) Private(fn Handler) *InlineQueryHandlers {
	h.handleInlineQuery(func(iq *gotgbot.InlineQuery) bool {
		return iq != nil && iq.ChatType == chat.Private.String()
	}, fn)
	return h
}

// Group handles inline queries from group chats.
func (h *InlineQueryHandlers) Group(fn Handler) *InlineQueryHandlers {
	h.handleInlineQuery(func(iq *gotgbot.InlineQuery) bool {
		return iq != nil && iq.ChatType == chat.Group.String()
	}, fn)
	return h
}

// Supergroup handles inline queries from supergroup chats.
func (h *InlineQueryHandlers) Supergroup(fn Handler) *InlineQueryHandlers {
	h.handleInlineQuery(func(iq *gotgbot.InlineQuery) bool {
		return iq != nil && iq.ChatType == chat.Supergroup.String()
	}, fn)
	return h
}

// Channel handles inline queries from channels.
func (h *InlineQueryHandlers) Channel(fn Handler) *InlineQueryHandlers {
	h.handleInlineQuery(func(iq *gotgbot.InlineQuery) bool {
		return iq != nil && iq.ChatType == chat.Channel.String()
	}, fn)
	return h
}
