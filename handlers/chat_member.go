package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"

	"github.com/enetx/tg/core"
	"github.com/enetx/tg/types/chatmember"
)

type ChatMemberHandlers struct{ Bot core.BotAPI }

func (h *ChatMemberHandlers) handleChatMember(f filters.ChatMember, fn Handler) {
	h.Bot.Dispatcher().AddHandler(handlers.NewChatMember(f, wrap(h.Bot, middlewares(h.Bot), fn)))
}

func (h *ChatMemberHandlers) Any(fn Handler) *ChatMemberHandlers {
	h.handleChatMember(nil, fn)
	return h
}

func (h *ChatMemberHandlers) StatusChange(from, to chatmember.ChatMemberStatus, fn Handler) *ChatMemberHandlers {
	filter := func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm != nil &&
			cm.OldChatMember != nil &&
			cm.NewChatMember != nil &&
			cm.OldChatMember.GetStatus() == from.String() &&
			cm.NewChatMember.GetStatus() == to.String()
	}

	h.handleChatMember(filter, fn)
	return h
}

func (h *ChatMemberHandlers) Joined(fn Handler) *ChatMemberHandlers {
	return h.StatusChange(chatmember.Left, chatmember.Member, fn)
}

func (h *ChatMemberHandlers) Left(fn Handler) *ChatMemberHandlers {
	return h.StatusChange(chatmember.Member, chatmember.Left, fn)
}

func (h *ChatMemberHandlers) Banned(fn Handler) *ChatMemberHandlers {
	return h.StatusChange(chatmember.Member, chatmember.Kicked, fn)
}

func (h *ChatMemberHandlers) Unbanned(fn Handler) *ChatMemberHandlers {
	return h.StatusChange(chatmember.Kicked, chatmember.Member, fn)
}

func (h *ChatMemberHandlers) Restricted(fn Handler) *ChatMemberHandlers {
	return h.StatusChange(chatmember.Member, chatmember.Restricted, fn)
}

func (h *ChatMemberHandlers) Unrestricted(fn Handler) *ChatMemberHandlers {
	return h.StatusChange(chatmember.Restricted, chatmember.Member, fn)
}

func (h *ChatMemberHandlers) Promoted(fn Handler) *ChatMemberHandlers {
	return h.StatusChange(chatmember.Member, chatmember.Administrator, fn)
}

func (h *ChatMemberHandlers) Demoted(fn Handler) *ChatMemberHandlers {
	return h.StatusChange(chatmember.Administrator, chatmember.Member, fn)
}

func (h *ChatMemberHandlers) ChatID(id int64, fn Handler) *ChatMemberHandlers {
	h.handleChatMember(func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm != nil && cm.Chat.Id == id
	}, fn)
	return h
}

func (h *ChatMemberHandlers) UserID(id int64, fn Handler) *ChatMemberHandlers {
	h.handleChatMember(func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm != nil && cm.NewChatMember.GetUser().Id == id
	}, fn)
	return h
}

func (h *ChatMemberHandlers) FromUserID(id int64, fn Handler) *ChatMemberHandlers {
	h.handleChatMember(func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm != nil && cm.From.Id == id
	}, fn)
	return h
}

func (h *ChatMemberHandlers) NewStatus(status chatmember.ChatMemberStatus, fn Handler) *ChatMemberHandlers {
	h.handleChatMember(func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm != nil && cm.NewChatMember.GetStatus() == status.String()
	}, fn)
	return h
}

func (h *ChatMemberHandlers) OldStatus(status chatmember.ChatMemberStatus, fn Handler) *ChatMemberHandlers {
	h.handleChatMember(func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm != nil && cm.OldChatMember.GetStatus() == status.String()
	}, fn)
	return h
}

func (h *ChatMemberHandlers) HasInviteLink(fn Handler) *ChatMemberHandlers {
	h.handleChatMember(func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm != nil && cm.InviteLink != nil
	}, fn)
	return h
}
