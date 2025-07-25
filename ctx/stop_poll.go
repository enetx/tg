package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
)

// StopPoll represents a request to stop a poll.
type StopPoll struct {
	ctx       *Context
	opts      *gotgbot.StopPollOpts
	chatID    Option[int64]
	messageID Option[int64]
}

// ChatID sets the target chat ID.
func (sp *StopPoll) ChatID(chatID int64) *StopPoll {
	sp.chatID = Some(chatID)
	return sp
}

// MessageID sets the target message ID.
func (sp *StopPoll) MessageID(messageID int64) *StopPoll {
	sp.messageID = Some(messageID)
	return sp
}

// Business sets the business connection ID.
func (sp *StopPoll) Business(id String) *StopPoll {
	sp.opts.BusinessConnectionId = id.Std()
	return sp
}

// Markup sets the reply markup keyboard for the stopped poll.
func (sp *StopPoll) Markup(kb *keyboard.InlineKeyboard) *StopPoll {
	if markup, ok := kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
		sp.opts.ReplyMarkup = markup
	}

	return sp
}

// Timeout sets a custom timeout for this request.
func (sp *StopPoll) Timeout(duration time.Duration) *StopPoll {
	if sp.opts.RequestOpts == nil {
		sp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sp.opts.RequestOpts.Timeout = duration

	return sp
}

// APIURL sets a custom API URL for this request.
func (sp *StopPoll) APIURL(url String) *StopPoll {
	if sp.opts.RequestOpts == nil {
		sp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sp.opts.RequestOpts.APIURL = url.Std()

	return sp
}

// Send stops the poll.
func (sp *StopPoll) Send() Result[*gotgbot.Poll] {
	return ResultOf(sp.ctx.Bot.Raw().StopPoll(
		sp.chatID.UnwrapOr(sp.ctx.EffectiveChat.Id),
		sp.messageID.UnwrapOr(sp.ctx.EffectiveMessage.MessageId),
		sp.opts,
	))
}
