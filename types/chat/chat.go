// Package chat defines Telegram chat types.
//
// See: https://core.telegram.org/bots/api#chat
package chat

import "github.com/PaulSonOfLars/gotgbot/v2"

// ChatType enumerates all supported Telegram chat types.
type ChatType int

const (
	Private    ChatType = iota // Private chat with a user
	Group                      // Group chat with multiple users
	Supergroup                 // Supergroup — large group with additional features
	Channel                    // Channel — one-way communication, e.g. news feed
	Sender                     // Sender chat type for special contexts
)

// String returns the canonical string representation used by Telegram Bot API.
func (c ChatType) String() string {
	switch c {
	case Private:
		return gotgbot.ChatTypePrivate
	case Group:
		return gotgbot.ChatTypeGroup
	case Supergroup:
		return gotgbot.ChatTypeSupergroup
	case Channel:
		return gotgbot.ChatTypeChannel
	case Sender:
		return "sender"
	default:
		return "unknown"
	}
}
