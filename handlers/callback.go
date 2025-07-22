package handlers

import (
	"fmt"
	"reflect"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	. "github.com/enetx/g"
	"github.com/enetx/tg/core"
)

// CallbackHandler handles callback query events from inline keyboards.
type CallbackHandler struct {
	bot          core.BotAPI
	filter       filters.CallbackQuery
	handler      Handler
	name         string
	allowChannel bool
}

// AllowChannel configures the handler to process callback queries from channels.
func (h *CallbackHandler) AllowChannel() *CallbackHandler {
	h.allowChannel = true
	return h
}

// Register registers the callback handler with the bot dispatcher.
func (h *CallbackHandler) Register() *CallbackHandler {
	h.bot.Dispatcher().RemoveHandlerFromGroup(h.name, 0)

	c := handlers.CallbackQuery{
		AllowChannel: h.allowChannel,
		Filter:       h.filter,
		Response:     wrap(h.bot, middlewares(h.bot), h.handler),
	}

	h.bot.Dispatcher().AddHandlerToGroup(namedHandler{h.name, c}, 0)

	return h
}

// CallbackHandlers provides methods to handle callback query events.
type CallbackHandlers struct{ Bot core.BotAPI }

// handleCallback creates and registers a new callback query handler with the specified filter.
func (h *CallbackHandlers) handleCallback(f filters.CallbackQuery, fn Handler) *CallbackHandler {
	return (&CallbackHandler{
		bot:     h.Bot,
		filter:  f,
		handler: fn,
		name:    fmt.Sprintf("callback_%x_%x", reflect.ValueOf(f).Pointer(), reflect.ValueOf(fn).Pointer()),
	}).Register()
}

// Any handles all callback queries.
func (h *CallbackHandlers) Any(fn Handler) *CallbackHandler { return h.handleCallback(nil, fn) }

// Equal handles callback queries with exact matching data.
func (h *CallbackHandlers) Equal(data String, fn Handler) *CallbackHandler {
	return h.handleCallback(func(q *gotgbot.CallbackQuery) bool {
		return q != nil && q.Data == data.Std()
	}, fn)
}

// Prefix handles callback queries with data starting with the specified prefix.
func (h *CallbackHandlers) Prefix(prefix String, fn Handler) *CallbackHandler {
	return h.handleCallback(func(q *gotgbot.CallbackQuery) bool {
		return q != nil && String(q.Data).StartsWith(prefix)
	}, fn)
}

// Suffix handles callback queries with data ending with the specified suffix.
func (h *CallbackHandlers) Suffix(suffix String, fn Handler) *CallbackHandler {
	return h.handleCallback(func(q *gotgbot.CallbackQuery) bool {
		return q != nil && String(q.Data).EndsWith(suffix)
	}, fn)
}

// FromUserID handles callback queries from a specific user ID.
func (h *CallbackHandlers) FromUserID(id int64, fn Handler) *CallbackHandler {
	return h.handleCallback(func(q *gotgbot.CallbackQuery) bool {
		return q != nil && q.From.Id == id
	}, fn)
}

// GameName handles callback queries from games with the specified short name.
func (h *CallbackHandlers) GameName(name String, fn Handler) *CallbackHandler {
	return h.handleCallback(func(q *gotgbot.CallbackQuery) bool {
		return q != nil && q.GameShortName == name.Std()
	}, fn)
}

// Inline handles callback queries with inline message ID.
func (h *CallbackHandlers) Inline(fn Handler) *CallbackHandler {
	return h.handleCallback(func(q *gotgbot.CallbackQuery) bool {
		return q != nil && q.InlineMessageId != ""
	}, fn)
}

// ChatInstance handles callback queries with the specified chat instance.
func (h *CallbackHandlers) ChatInstance(inst String, fn Handler) *CallbackHandler {
	return h.handleCallback(func(q *gotgbot.CallbackQuery) bool {
		return q != nil && q.ChatInstance == inst.Std()
	}, fn)
}
