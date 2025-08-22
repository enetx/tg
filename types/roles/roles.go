// Package roles defines Telegram chat administrator role types and utilities.
package roles

import "github.com/PaulSonOfLars/gotgbot/v2"

// Role enumerates all supported Telegram chat administrator roles.
type Role int

const (
	ManageChat           Role = iota // Manage chat settings and info
	DeleteMessages                   // Delete messages from other users
	ManageVideoChats                 // Manage video chats
	RestrictMembers                  // Restrict, unrestrict, kick, and ban chat members
	PromoteMembers                   // Add new administrators with fewer rights than their own
	ChangeInfo                       // Change chat title, photo, and other settings
	InviteUsers                      // Invite new users to the chat
	PostMessages                     // Post messages in channels
	EditMessages                     // Edit messages of other users in channels
	PinMessages                      // Pin messages
	PostStories                      // Post stories in channels
	EditStories                      // Edit stories of other users in channels
	DeleteStories                    // Delete stories of other users in channels
	ManageTopics                     // Manage topics in forum supergroups
	ManageDirectMessages             // Manage direct messages of the channel
)

// Roles creates a PromoteChatMemberOpts object with the specified administrator roles enabled.
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
		case PostStories:
			opts.CanPostStories = true
		case EditStories:
			opts.CanEditStories = true
		case DeleteStories:
			opts.CanDeleteStories = true
		case ManageTopics:
			opts.CanManageTopics = true
		case ManageDirectMessages:
			opts.CanManageDirectMessages = true
		}
	}

	return opts
}
