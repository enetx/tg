package handlers

import (
	"fmt"
	"reflect"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"

	"github.com/enetx/tg/core"

	. "github.com/enetx/g"
)

type CallbackHandler struct {
	bot          core.BotAPI
	filter       filters.CallbackQuery
	handler      Handler
	name         string
	allowChannel bool
}

func (h *CallbackHandler) AllowChannel() *CallbackHandler {
	h.allowChannel = true
	return h
}

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

type CallbackHandlers struct{ Bot core.BotAPI }

func (h *CallbackHandlers) handleCallback(f filters.CallbackQuery, fn Handler) *CallbackHandler {
	return (&CallbackHandler{
		bot:     h.Bot,
		filter:  f,
		handler: fn,
		name:    fmt.Sprintf("callback_%x_%x", reflect.ValueOf(f).Pointer(), reflect.ValueOf(fn).Pointer()),
	}).Register()
}

func (h *CallbackHandlers) Any(fn Handler) *CallbackHandler { return h.handleCallback(nil, fn) }

func (h *CallbackHandlers) Equal(data String, fn Handler) *CallbackHandler {
	return h.handleCallback(func(q *gotgbot.CallbackQuery) bool {
		return q != nil && q.Data == data.Std()
	}, fn)
}

func (h *CallbackHandlers) Prefix(prefix String, fn Handler) *CallbackHandler {
	return h.handleCallback(func(q *gotgbot.CallbackQuery) bool {
		return q != nil && String(q.Data).StartsWith(prefix)
	}, fn)
}

func (h *CallbackHandlers) Suffix(suffix String, fn Handler) *CallbackHandler {
	return h.handleCallback(func(q *gotgbot.CallbackQuery) bool {
		return q != nil && String(q.Data).EndsWith(suffix)
	}, fn)
}

func (h *CallbackHandlers) FromUserID(id int64, fn Handler) *CallbackHandler {
	return h.handleCallback(func(q *gotgbot.CallbackQuery) bool {
		return q != nil && q.From.Id == id
	}, fn)
}

func (h *CallbackHandlers) GameName(name String, fn Handler) *CallbackHandler {
	return h.handleCallback(func(q *gotgbot.CallbackQuery) bool {
		return q != nil && q.GameShortName == name.Std()
	}, fn)
}

func (h *CallbackHandlers) Inline(fn Handler) *CallbackHandler {
	return h.handleCallback(func(q *gotgbot.CallbackQuery) bool {
		return q != nil && q.InlineMessageId != ""
	}, fn)
}

func (h *CallbackHandlers) ChatInstance(inst String, fn Handler) *CallbackHandler {
	return h.handleCallback(func(q *gotgbot.CallbackQuery) bool {
		return q != nil && q.ChatInstance == inst.Std()
	}, fn)
}
