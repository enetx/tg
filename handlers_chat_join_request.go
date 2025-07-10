package tg

import "github.com/PaulSonOfLars/gotgbot/v2"

type ChatJoinRequestHandlers struct{ b *Bot }

func (h *ChatJoinRequestHandlers) Any(fn Handler) *ChatJoinRequestHandlers {
	h.b.handleChatJoinRequest(nil, fn)
	return h
}

func (h *ChatJoinRequestHandlers) ChatID(id int64, fn Handler) *ChatJoinRequestHandlers {
	h.b.handleChatJoinRequest(func(r *gotgbot.ChatJoinRequest) bool {
		return r != nil && r.Chat.Id == id
	}, fn)
	return h
}

func (h *ChatJoinRequestHandlers) FromUserID(id int64, fn Handler) *ChatJoinRequestHandlers {
	h.b.handleChatJoinRequest(func(r *gotgbot.ChatJoinRequest) bool {
		return r != nil && r.From.Id == id
	}, fn)
	return h
}

func (h *ChatJoinRequestHandlers) HasInviteLink(fn Handler) *ChatJoinRequestHandlers {
	h.b.handleChatJoinRequest(func(r *gotgbot.ChatJoinRequest) bool {
		return r != nil && r.InviteLink != nil
	}, fn)
	return h
}
