package business

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// Message provides tools for managing business messages:
// marking as read and deleting them.
type Message struct {
	bot    Bot
	connID String
}

// Read creates a request to mark a specific business message as read.
func (m *Message) Read(chatID, messageID int64) *Read {
	return &Read{
		bot:       m.bot,
		connID:    m.connID,
		chatID:    chatID,
		messageID: messageID,
		opts:      new(gotgbot.ReadBusinessMessageOpts),
	}
}

// Delete creates a request to delete one or more business messages.
func (m *Message) Delete(messageIDs Slice[int64]) *Delete {
	return &Delete{
		bot:        m.bot,
		connID:     m.connID,
		messageIDs: messageIDs,
		opts:       new(gotgbot.DeleteBusinessMessagesOpts),
	}
}
