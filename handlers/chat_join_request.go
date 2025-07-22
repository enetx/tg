package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	"github.com/enetx/tg/core"
)

// ChatJoinRequestHandlers provides methods to handle chat join request events.
type ChatJoinRequestHandlers struct{ Bot core.BotAPI }

// handleChatJoinRequest registers a chat join request handler with the dispatcher.
func (h *ChatJoinRequestHandlers) handleChatJoinRequest(f filters.ChatJoinRequest, fn Handler) {
	h.Bot.Dispatcher().AddHandler(handlers.NewChatJoinRequest(f, wrap(h.Bot, middlewares(h.Bot), fn)))
}

// Any handles all chat join requests.
func (h *ChatJoinRequestHandlers) Any(fn Handler) *ChatJoinRequestHandlers {
	h.handleChatJoinRequest(nil, fn)
	return h
}

// ChatID handles chat join requests for a specific chat ID.
func (h *ChatJoinRequestHandlers) ChatID(id int64, fn Handler) *ChatJoinRequestHandlers {
	h.handleChatJoinRequest(func(r *gotgbot.ChatJoinRequest) bool {
		return r != nil && r.Chat.Id == id
	}, fn)
	return h
}

// FromUserID handles chat join requests from a specific user ID.
func (h *ChatJoinRequestHandlers) FromUserID(id int64, fn Handler) *ChatJoinRequestHandlers {
	h.handleChatJoinRequest(func(r *gotgbot.ChatJoinRequest) bool {
		return r != nil && r.From.Id == id
	}, fn)
	return h
}

// HasInviteLink handles chat join requests that include an invite link.
func (h *ChatJoinRequestHandlers) HasInviteLink(fn Handler) *ChatJoinRequestHandlers {
	h.handleChatJoinRequest(func(r *gotgbot.ChatJoinRequest) bool {
		return r != nil && r.InviteLink != nil
	}, fn)
	return h
}
