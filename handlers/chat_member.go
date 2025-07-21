package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	"github.com/enetx/tg/core"
	"github.com/enetx/tg/types/chatmember"
)

// ChatMemberHandlers provides methods to handle chat member update events.
type ChatMemberHandlers struct{ Bot core.BotAPI }

func (h *ChatMemberHandlers) handleChatMember(f filters.ChatMember, fn Handler) {
	h.Bot.Dispatcher().AddHandler(handlers.NewChatMember(f, wrap(h.Bot, middlewares(h.Bot), fn)))
}

// Any handles all chat member updates.
func (h *ChatMemberHandlers) Any(fn Handler) *ChatMemberHandlers {
	h.handleChatMember(nil, fn)
	return h
}

// StatusChange handles chat member status changes from one status to another.
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

// Joined handles when a member joins the chat.
func (h *ChatMemberHandlers) Joined(fn Handler) *ChatMemberHandlers {
	return h.StatusChange(chatmember.Left, chatmember.Member, fn)
}

// Left handles when a member leaves the chat.
func (h *ChatMemberHandlers) Left(fn Handler) *ChatMemberHandlers {
	return h.StatusChange(chatmember.Member, chatmember.Left, fn)
}

// Banned handles when a member is banned from the chat.
func (h *ChatMemberHandlers) Banned(fn Handler) *ChatMemberHandlers {
	return h.StatusChange(chatmember.Member, chatmember.Kicked, fn)
}

// Unbanned handles when a member is unbanned from the chat.
func (h *ChatMemberHandlers) Unbanned(fn Handler) *ChatMemberHandlers {
	return h.StatusChange(chatmember.Kicked, chatmember.Member, fn)
}

// Restricted handles when a member is restricted in the chat.
func (h *ChatMemberHandlers) Restricted(fn Handler) *ChatMemberHandlers {
	return h.StatusChange(chatmember.Member, chatmember.Restricted, fn)
}

// Unrestricted handles when a member's restrictions are removed.
func (h *ChatMemberHandlers) Unrestricted(fn Handler) *ChatMemberHandlers {
	return h.StatusChange(chatmember.Restricted, chatmember.Member, fn)
}

// Promoted handles when a member is promoted to administrator.
func (h *ChatMemberHandlers) Promoted(fn Handler) *ChatMemberHandlers {
	return h.StatusChange(chatmember.Member, chatmember.Administrator, fn)
}

// Demoted handles when an administrator is demoted to regular member.
func (h *ChatMemberHandlers) Demoted(fn Handler) *ChatMemberHandlers {
	return h.StatusChange(chatmember.Administrator, chatmember.Member, fn)
}

// ChatID handles chat member updates in a specific chat.
func (h *ChatMemberHandlers) ChatID(id int64, fn Handler) *ChatMemberHandlers {
	h.handleChatMember(func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm != nil && cm.Chat.Id == id
	}, fn)
	return h
}

// UserID handles chat member updates for a specific user.
func (h *ChatMemberHandlers) UserID(id int64, fn Handler) *ChatMemberHandlers {
	h.handleChatMember(func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm != nil && cm.NewChatMember.GetUser().Id == id
	}, fn)
	return h
}

// FromUserID handles chat member updates initiated by a specific user.
func (h *ChatMemberHandlers) FromUserID(id int64, fn Handler) *ChatMemberHandlers {
	h.handleChatMember(func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm != nil && cm.From.Id == id
	}, fn)
	return h
}

// NewStatus handles chat member updates where the new status matches the specified status.
func (h *ChatMemberHandlers) NewStatus(status chatmember.ChatMemberStatus, fn Handler) *ChatMemberHandlers {
	h.handleChatMember(func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm != nil && cm.NewChatMember.GetStatus() == status.String()
	}, fn)
	return h
}

// OldStatus handles chat member updates where the old status matches the specified status.
func (h *ChatMemberHandlers) OldStatus(status chatmember.ChatMemberStatus, fn Handler) *ChatMemberHandlers {
	h.handleChatMember(func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm != nil && cm.OldChatMember.GetStatus() == status.String()
	}, fn)
	return h
}

// HasInviteLink handles chat member updates that include an invite link.
func (h *ChatMemberHandlers) HasInviteLink(fn Handler) *ChatMemberHandlers {
	h.handleChatMember(func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm != nil && cm.InviteLink != nil
	}, fn)
	return h
}
