// Package rights defines Telegram chat administrator rights types and utilities.
package rights

import "github.com/PaulSonOfLars/gotgbot/v2"

// Right enumerates all supported Telegram chat administrator rights.
type Right int

const (
	Anonymous            Right = iota // Administrator can be anonymous
	ManageChat                        // Manage chat settings and info
	DeleteMessages                    // Delete messages from other users
	ManageVideoChats                  // Manage video chats
	RestrictMembers                   // Restrict, unrestrict, kick, and ban chat members
	PromoteMembers                    // Add new administrators with fewer rights than their own
	ChangeInfo                        // Change chat title, photo, and other settings
	InviteUsers                       // Invite new users to the chat
	PostMessages                      // Post messages in channels
	EditMessages                      // Edit messages of other users in channels
	PinMessages                       // Pin messages
	PostStories                       // Post stories in channels
	EditStories                       // Edit stories of other users in channels
	DeleteStories                     // Delete stories of other users in channels
	ManageTopics                      // Manage topics in forum supergroups
	ManageDirectMessages              // Manage direct messages of the channel
)

// Rights creates a ChatAdministratorRights object with the specified administrator rights enabled.
func Rights(list ...Right) *gotgbot.ChatAdministratorRights {
	rights := new(gotgbot.ChatAdministratorRights)
	for _, r := range list {
		switch r {
		case Anonymous:
			rights.IsAnonymous = true
		case ManageChat:
			rights.CanManageChat = true
		case DeleteMessages:
			rights.CanDeleteMessages = true
		case ManageVideoChats:
			rights.CanManageVideoChats = true
		case RestrictMembers:
			rights.CanRestrictMembers = true
		case PromoteMembers:
			rights.CanPromoteMembers = true
		case ChangeInfo:
			rights.CanChangeInfo = true
		case InviteUsers:
			rights.CanInviteUsers = true
		case PostMessages:
			rights.CanPostMessages = true
		case EditMessages:
			rights.CanEditMessages = true
		case PinMessages:
			rights.CanPinMessages = true
		case PostStories:
			rights.CanPostStories = true
		case EditStories:
			rights.CanEditStories = true
		case DeleteStories:
			rights.CanDeleteStories = true
		case ManageTopics:
			rights.CanManageTopics = true
		case ManageDirectMessages:
			rights.CanManageDirectMessages = true
		}
	}

	return rights
}
