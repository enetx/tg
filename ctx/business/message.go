package business

import (
	"time"

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

// Read is a request builder for marking a business message as read.
type Read struct {
	bot       Bot
	connID    String
	chatID    int64
	messageID int64
	opts      *gotgbot.ReadBusinessMessageOpts
}

// Timeout sets a custom timeout for this request.
func (r *Read) Timeout(duration time.Duration) *Read {
	if r.opts.RequestOpts == nil {
		r.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	r.opts.RequestOpts.Timeout = duration

	return r
}

// APIURL sets a custom API URL for this request.
func (r *Read) APIURL(url String) *Read {
	if r.opts.RequestOpts == nil {
		r.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	r.opts.RequestOpts.APIURL = url.Std()

	return r
}

// Send executes the Read request.
func (r *Read) Send() Result[bool] {
	return ResultOf(r.bot.Raw().ReadBusinessMessage(
		r.connID.Std(),
		r.chatID,
		r.messageID,
		r.opts,
	))
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

// Delete is a request builder for deleting business messages.
type Delete struct {
	bot        Bot
	connID     String
	messageIDs Slice[int64]
	opts       *gotgbot.DeleteBusinessMessagesOpts
}

// Timeout sets a custom timeout for this request.
func (d *Delete) Timeout(duration time.Duration) *Delete {
	if d.opts.RequestOpts == nil {
		d.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	d.opts.RequestOpts.Timeout = duration

	return d
}

// APIURL sets a custom API URL for this request.
func (d *Delete) APIURL(url String) *Delete {
	if d.opts.RequestOpts == nil {
		d.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	d.opts.RequestOpts.APIURL = url.Std()

	return d
}

// Send executes the Delete request.
func (d *Delete) Send() Result[bool] {
	return ResultOf(d.bot.Raw().DeleteBusinessMessages(
		d.connID.Std(),
		d.messageIDs,
		d.opts,
	))
}
