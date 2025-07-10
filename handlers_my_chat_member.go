package tg

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/types/chatmember"
)

type MyChatMemberHandlers struct{ b *Bot }

func (h *MyChatMemberHandlers) Any(fn Handler) *MyChatMemberHandlers {
	h.b.handleMyChatMember(nil, fn)
	return h
}

func (h *MyChatMemberHandlers) StatusChange(from, to chatmember.ChatMemberStatus, fn Handler) *MyChatMemberHandlers {
	filter := func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm != nil &&
			cm.OldChatMember != nil &&
			cm.NewChatMember != nil &&
			cm.OldChatMember.GetStatus() == from.String() &&
			cm.NewChatMember.GetStatus() == to.String()
	}

	h.b.handleMyChatMember(filter, fn)
	return h
}

func (h *MyChatMemberHandlers) Joined(fn Handler) *MyChatMemberHandlers {
	return h.StatusChange(chatmember.Left, chatmember.Member, fn)
}

func (h *MyChatMemberHandlers) Left(fn Handler) *MyChatMemberHandlers {
	return h.StatusChange(chatmember.Member, chatmember.Left, fn)
}

func (h *MyChatMemberHandlers) Banned(fn Handler) *MyChatMemberHandlers {
	return h.StatusChange(chatmember.Member, chatmember.Kicked, fn)
}

func (h *MyChatMemberHandlers) Unbanned(fn Handler) *MyChatMemberHandlers {
	return h.StatusChange(chatmember.Kicked, chatmember.Member, fn)
}

func (h *MyChatMemberHandlers) Restricted(fn Handler) *MyChatMemberHandlers {
	return h.StatusChange(chatmember.Member, chatmember.Restricted, fn)
}

func (h *MyChatMemberHandlers) Unrestricted(fn Handler) *MyChatMemberHandlers {
	return h.StatusChange(chatmember.Restricted, chatmember.Member, fn)
}

func (h *MyChatMemberHandlers) Promoted(fn Handler) *MyChatMemberHandlers {
	return h.StatusChange(chatmember.Member, chatmember.Administrator, fn)
}

func (h *MyChatMemberHandlers) Demoted(fn Handler) *MyChatMemberHandlers {
	return h.StatusChange(chatmember.Administrator, chatmember.Member, fn)
}

func (h *MyChatMemberHandlers) ChatID(id int64, fn Handler) *MyChatMemberHandlers {
	h.b.handleMyChatMember(func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm != nil && cm.Chat.Id == id
	}, fn)
	return h
}

func (h *MyChatMemberHandlers) UserID(id int64, fn Handler) *MyChatMemberHandlers {
	h.b.handleMyChatMember(func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm != nil && cm.NewChatMember.GetUser().Id == id
	}, fn)
	return h
}

func (h *MyChatMemberHandlers) FromUserID(id int64, fn Handler) *MyChatMemberHandlers {
	h.b.handleMyChatMember(func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm != nil && cm.From.Id == id
	}, fn)
	return h
}

func (h *MyChatMemberHandlers) NewStatus(status chatmember.ChatMemberStatus, fn Handler) *MyChatMemberHandlers {
	h.b.handleMyChatMember(func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm != nil && cm.NewChatMember.GetStatus() == status.String()
	}, fn)
	return h
}

func (h *MyChatMemberHandlers) OldStatus(status chatmember.ChatMemberStatus, fn Handler) *MyChatMemberHandlers {
	h.b.handleMyChatMember(func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm != nil && cm.OldChatMember.GetStatus() == status.String()
	}, fn)
	return h
}

func (h *MyChatMemberHandlers) HasInviteLink(fn Handler) *MyChatMemberHandlers {
	h.b.handleMyChatMember(func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm != nil && cm.InviteLink != nil
	}, fn)
	return h
}
