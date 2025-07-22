package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	"github.com/enetx/tg/core"
	"github.com/enetx/tg/types/chatmember"
)

// MyChatMemberHandlers provides methods to handle bot's own chat member status updates.
type MyChatMemberHandlers struct{ Bot core.BotAPI }

// handleMyChatMember registers a bot chat member update handler with the dispatcher.
func (h *MyChatMemberHandlers) handleMyChatMember(f filters.ChatMember, fn Handler) {
	h.Bot.Dispatcher().AddHandler(handlers.NewMyChatMember(f, wrap(h.Bot, middlewares(h.Bot), fn)))
}

// Any handles all bot chat member updates.
func (h *MyChatMemberHandlers) Any(fn Handler) *MyChatMemberHandlers {
	h.handleMyChatMember(nil, fn)
	return h
}

// StatusChange handles bot chat member status changes from one status to another.
func (h *MyChatMemberHandlers) StatusChange(from, to chatmember.ChatMemberStatus, fn Handler) *MyChatMemberHandlers {
	filter := func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm != nil &&
			cm.OldChatMember != nil &&
			cm.NewChatMember != nil &&
			cm.OldChatMember.GetStatus() == from.String() &&
			cm.NewChatMember.GetStatus() == to.String()
	}

	h.handleMyChatMember(filter, fn)
	return h
}

// Joined handles when the bot joins a chat.
func (h *MyChatMemberHandlers) Joined(fn Handler) *MyChatMemberHandlers {
	return h.StatusChange(chatmember.Left, chatmember.Member, fn)
}

// Left handles when the bot leaves a chat.
func (h *MyChatMemberHandlers) Left(fn Handler) *MyChatMemberHandlers {
	return h.StatusChange(chatmember.Member, chatmember.Left, fn)
}

// Banned handles when the bot is banned from a chat.
func (h *MyChatMemberHandlers) Banned(fn Handler) *MyChatMemberHandlers {
	return h.StatusChange(chatmember.Member, chatmember.Kicked, fn)
}

// Unbanned handles when the bot is unbanned from a chat.
func (h *MyChatMemberHandlers) Unbanned(fn Handler) *MyChatMemberHandlers {
	return h.StatusChange(chatmember.Kicked, chatmember.Member, fn)
}

// Restricted handles when the bot is restricted in a chat.
func (h *MyChatMemberHandlers) Restricted(fn Handler) *MyChatMemberHandlers {
	return h.StatusChange(chatmember.Member, chatmember.Restricted, fn)
}

// Unrestricted handles when the bot's restrictions are removed.
func (h *MyChatMemberHandlers) Unrestricted(fn Handler) *MyChatMemberHandlers {
	return h.StatusChange(chatmember.Restricted, chatmember.Member, fn)
}

// Promoted handles when the bot is promoted to administrator.
func (h *MyChatMemberHandlers) Promoted(fn Handler) *MyChatMemberHandlers {
	return h.StatusChange(chatmember.Member, chatmember.Administrator, fn)
}

// Demoted handles when the bot is demoted from administrator.
func (h *MyChatMemberHandlers) Demoted(fn Handler) *MyChatMemberHandlers {
	return h.StatusChange(chatmember.Administrator, chatmember.Member, fn)
}

// ChatID handles bot chat member updates in a specific chat.
func (h *MyChatMemberHandlers) ChatID(id int64, fn Handler) *MyChatMemberHandlers {
	h.handleMyChatMember(func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm != nil && cm.Chat.Id == id
	}, fn)
	return h
}

// UserID handles bot chat member updates for a specific user.
func (h *MyChatMemberHandlers) UserID(id int64, fn Handler) *MyChatMemberHandlers {
	h.handleMyChatMember(func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm != nil && cm.NewChatMember.GetUser().Id == id
	}, fn)
	return h
}

// FromUserID handles bot chat member updates initiated by a specific user.
func (h *MyChatMemberHandlers) FromUserID(id int64, fn Handler) *MyChatMemberHandlers {
	h.handleMyChatMember(func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm != nil && cm.From.Id == id
	}, fn)
	return h
}

// NewStatus handles bot chat member updates where the new status matches the specified status.
func (h *MyChatMemberHandlers) NewStatus(status chatmember.ChatMemberStatus, fn Handler) *MyChatMemberHandlers {
	h.handleMyChatMember(func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm != nil && cm.NewChatMember.GetStatus() == status.String()
	}, fn)
	return h
}

// OldStatus handles bot chat member updates where the old status matches the specified status.
func (h *MyChatMemberHandlers) OldStatus(status chatmember.ChatMemberStatus, fn Handler) *MyChatMemberHandlers {
	h.handleMyChatMember(func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm != nil && cm.OldChatMember.GetStatus() == status.String()
	}, fn)
	return h
}

// HasInviteLink handles bot chat member updates that include an invite link.
func (h *MyChatMemberHandlers) HasInviteLink(fn Handler) *MyChatMemberHandlers {
	h.handleMyChatMember(func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm != nil && cm.InviteLink != nil
	}, fn)
	return h
}
