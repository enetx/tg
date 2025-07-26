package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/inline"
)

// SavePreparedInlineMessage represents a request to save prepared inline message.
type SavePreparedInlineMessage struct {
	ctx    *Context
	userID int64
	result inline.QueryResult
	opts   *gotgbot.SavePreparedInlineMessageOpts
}

// AllowUserChats allows the message to be sent to user chats.
func (spim *SavePreparedInlineMessage) AllowUserChats() *SavePreparedInlineMessage {
	spim.opts.AllowUserChats = true
	return spim
}

// AllowBotChats allows the message to be sent to bot chats.
func (spim *SavePreparedInlineMessage) AllowBotChats() *SavePreparedInlineMessage {
	spim.opts.AllowBotChats = true
	return spim
}

// AllowGroupChats allows the message to be sent to group chats.
func (spim *SavePreparedInlineMessage) AllowGroupChats() *SavePreparedInlineMessage {
	spim.opts.AllowGroupChats = true
	return spim
}

// AllowChannelChats allows the message to be sent to channel chats.
func (spim *SavePreparedInlineMessage) AllowChannelChats() *SavePreparedInlineMessage {
	spim.opts.AllowChannelChats = true
	return spim
}

// Timeout sets a custom timeout for this request.
func (spim *SavePreparedInlineMessage) Timeout(duration time.Duration) *SavePreparedInlineMessage {
	if spim.opts.RequestOpts == nil {
		spim.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	spim.opts.RequestOpts.Timeout = duration

	return spim
}

// APIURL sets a custom API URL for this request.
func (spim *SavePreparedInlineMessage) APIURL(url String) *SavePreparedInlineMessage {
	if spim.opts.RequestOpts == nil {
		spim.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	spim.opts.RequestOpts.APIURL = url.Std()

	return spim
}

// Send saves the prepared inline message.
func (spim *SavePreparedInlineMessage) Send() Result[*gotgbot.PreparedInlineMessage] {
	return ResultOf(spim.ctx.Bot.Raw().SavePreparedInlineMessage(
		spim.userID,
		spim.result.Build(),
		spim.opts,
	))
}
