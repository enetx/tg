// Package chat defines the chat-related types used by Telegram.
package chatmember

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
		return "creator"
	case Administrator:
		return "administrator"
	case Member:
		return "member"
	case Restricted:
		return "restricted"
	case Left:
		return "left"
	case Kicked:
		return "kicked"
	default:
		return "unknown"
	}
}
