package roles

import "github.com/PaulSonOfLars/gotgbot/v2"

type Role int

const (
	ManageChat Role = iota
	DeleteMessages
	ManageVideoChats
	RestrictMembers
	PromoteMembers
	ChangeInfo
	InviteUsers
	PostMessages
	EditMessages
	PinMessages
	ManageTopics
)

func Roles(list ...Role) *gotgbot.PromoteChatMemberOpts {
	opts := new(gotgbot.PromoteChatMemberOpts)
	for _, r := range list {
		switch r {
		case ManageChat:
			opts.CanManageChat = true
		case DeleteMessages:
			opts.CanDeleteMessages = true
		case ManageVideoChats:
			opts.CanManageVideoChats = true
		case RestrictMembers:
			opts.CanRestrictMembers = true
		case PromoteMembers:
			opts.CanPromoteMembers = true
		case ChangeInfo:
			opts.CanChangeInfo = true
		case InviteUsers:
			opts.CanInviteUsers = true
		case PostMessages:
			opts.CanPostMessages = true
		case EditMessages:
			opts.CanEditMessages = true
		case PinMessages:
			opts.CanPinMessages = true
		case ManageTopics:
			opts.CanManageTopics = true
		}
	}

	return opts
}
