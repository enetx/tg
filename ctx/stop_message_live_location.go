package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
)

// StopMessageLiveLocation represents a request to stop updating a live location message.
type StopMessageLiveLocation struct {
	ctx       *Context
	opts      *gotgbot.StopMessageLiveLocationOpts
	chatID    g.Option[int64]
	messageID g.Option[int64]
}

// ChatID sets the target chat ID.
func (smll *StopMessageLiveLocation) ChatID(chatID int64) *StopMessageLiveLocation {
	smll.chatID = g.Some(chatID)
	return smll
}

// MessageID sets the target message ID.
func (smll *StopMessageLiveLocation) MessageID(messageID int64) *StopMessageLiveLocation {
	smll.messageID = g.Some(messageID)
	return smll
}

// InlineMessageID sets the inline message ID to edit.
func (smll *StopMessageLiveLocation) InlineMessageID(id g.String) *StopMessageLiveLocation {
	smll.opts.InlineMessageId = id.Std()
	return smll
}

// Business sets the business connection ID for stopping the location.
func (smll *StopMessageLiveLocation) Business(id g.String) *StopMessageLiveLocation {
	smll.opts.BusinessConnectionId = id.Std()
	return smll
}

// Markup sets the reply markup keyboard for the message.
func (smll *StopMessageLiveLocation) Markup(kb *keyboard.InlineKeyboard) *StopMessageLiveLocation {
	if markup, ok := kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
		smll.opts.ReplyMarkup = markup
	}

	return smll
}

// Timeout sets a custom timeout for this request.
func (smll *StopMessageLiveLocation) Timeout(duration time.Duration) *StopMessageLiveLocation {
	if smll.opts.RequestOpts == nil {
		smll.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	smll.opts.RequestOpts.Timeout = duration

	return smll
}

// APIURL sets a custom API URL for this request.
func (smll *StopMessageLiveLocation) APIURL(url g.String) *StopMessageLiveLocation {
	if smll.opts.RequestOpts == nil {
		smll.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	smll.opts.RequestOpts.APIURL = url.Std()

	return smll
}

// Send stops updating the live location message.
func (smll *StopMessageLiveLocation) Send() g.Result[*gotgbot.Message] {
	smll.opts.ChatId = smll.chatID.UnwrapOr(smll.ctx.EffectiveChat.Id)
	smll.opts.MessageId = smll.messageID.UnwrapOr(smll.ctx.EffectiveMessage.MessageId)
	msg, _, err := smll.ctx.Bot.Raw().StopMessageLiveLocation(smll.opts)

	return g.ResultOf(msg, err)
}
