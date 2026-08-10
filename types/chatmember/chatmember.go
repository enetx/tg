// Package chat defines the chat-related types used by Telegram.
package chatmember

import "github.com/PaulSonOfLars/gotgbot/v2"

// ChatMemberStatus enumerates all possible member statuses in a chat.
//
// See: https://core.telegram.org/bots/api#chatmember
type ChatMemberStatus int

const (
	Creator       ChatMemberStatus = iota // The user is the owner of the chat.
	Administrator                         // The user has admin rights in the chat.
	Member                                // The user is a normal member of the chat.
	Restricted                            // The user has limited access in the chat.
	Left                                  // The user has left the chat voluntarily.
	Kicked                                // The user was removed and banned from the chat.
)

// String returns the canonical string representation used by the Telegram Bot API.
func (s ChatMemberStatus) String() string {
	switch s {
	case Creator:
		return gotgbot.ChatMemberStatusCreator
	case Administrator:
		return gotgbot.ChatMemberStatusAdministrator
	case Member:
		return gotgbot.ChatMemberStatusMember
	case Restricted:
		return gotgbot.ChatMemberStatusRestricted
	case Left:
		return gotgbot.ChatMemberStatusLeft
	case Kicked:
		return gotgbot.ChatMemberStatusKicked
	default:
		return "unknown"
	}
}
