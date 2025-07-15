package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"

	"github.com/enetx/tg/core"
)

type ChatJoinRequestHandlers struct{ Bot core.BotAPI }

func (h *ChatJoinRequestHandlers) handleChatJoinRequest(f filters.ChatJoinRequest, fn Handler) {
	h.Bot.Dispatcher().AddHandler(handlers.NewChatJoinRequest(f, wrap(h.Bot, middlewares(h.Bot), fn)))
}

func (h *ChatJoinRequestHandlers) Any(fn Handler) *ChatJoinRequestHandlers {
	h.handleChatJoinRequest(nil, fn)
	return h
}

func (h *ChatJoinRequestHandlers) ChatID(id int64, fn Handler) *ChatJoinRequestHandlers {
	h.handleChatJoinRequest(func(r *gotgbot.ChatJoinRequest) bool {
		return r != nil && r.Chat.Id == id
	}, fn)
	return h
}

func (h *ChatJoinRequestHandlers) FromUserID(id int64, fn Handler) *ChatJoinRequestHandlers {
	h.handleChatJoinRequest(func(r *gotgbot.ChatJoinRequest) bool {
		return r != nil && r.From.Id == id
	}, fn)
	return h
}

func (h *ChatJoinRequestHandlers) HasInviteLink(fn Handler) *ChatJoinRequestHandlers {
	h.handleChatJoinRequest(func(r *gotgbot.ChatJoinRequest) bool {
		return r != nil && r.InviteLink != nil
	}, fn)
	return h
}
