package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// SetChatDescription represents a request to set the chat description.
type SetChatDescription struct {
	ctx         *Context
	description String
	opts        *gotgbot.SetChatDescriptionOpts
	chatID      Option[int64]
}

// ChatID sets the target chat ID for this request.
func (scd *SetChatDescription) ChatID(id int64) *SetChatDescription {
	scd.chatID = Some(id)
	return scd
}

// Timeout sets a custom timeout for this request.
func (scd *SetChatDescription) Timeout(duration time.Duration) *SetChatDescription {
	if scd.opts.RequestOpts == nil {
		scd.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	scd.opts.RequestOpts.Timeout = duration

	return scd
}

// APIURL sets a custom API URL for this request.
func (scd *SetChatDescription) APIURL(url String) *SetChatDescription {
	if scd.opts.RequestOpts == nil {
		scd.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	scd.opts.RequestOpts.APIURL = url.Std()

	return scd
}

// Send executes the SetChatDescription request.
func (scd *SetChatDescription) Send() Result[bool] {
	chatID := scd.chatID.UnwrapOr(scd.ctx.EffectiveChat.Id)
	return ResultOf(scd.ctx.Bot.Raw().SetChatDescription(chatID, scd.opts))
}
