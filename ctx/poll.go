package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/types/effects"
)

type Poll struct {
	ctx         *Context
	question    String
	chatID      Option[int64]
	options     Slice[gotgbot.InputPollOption]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
	opts        *gotgbot.SendPollOpts
}

// After schedules the poll to be sent after the specified duration.
func (p *Poll) After(duration time.Duration) *Poll {
	p.after = Some(duration)
	return p
}

// DeleteAfter schedules the poll message to be deleted after the specified duration.
func (p *Poll) DeleteAfter(duration time.Duration) *Poll {
	p.deleteAfter = Some(duration)
	return p
}

// To sets the target chat ID for the poll.
func (p *Poll) To(id int64) *Poll {
	p.chatID = Some(id)
	return p
}

// Option adds a poll option with the specified text.
func (p *Poll) Option(text String) *Poll {
	opt := gotgbot.InputPollOption{Text: text.Std()}
	p.options.Push(opt)
	return p
}

// Anonymous makes the poll anonymous.
func (p *Poll) Anonymous() *Poll {
	p.opts.IsAnonymous = true
	return p
}

// Business sets the business connection ID for the poll.
func (p *Poll) Business(id String) *Poll {
	p.opts.BusinessConnectionId = id.Std()
	return p
}

// Thread sets the message thread ID for the poll.
func (p *Poll) Thread(id int64) *Poll {
	p.opts.MessageThreadId = id
	return p
}

// AllowPaidBroadcast allows the poll to be sent in paid broadcast channels.
func (p *Poll) AllowPaidBroadcast() *Poll {
	p.opts.AllowPaidBroadcast = true
	return p
}

// Effect sets a message effect for the poll.
func (p *Poll) Effect(effect effects.EffectType) *Poll {
	p.opts.MessageEffectId = effect.String()
	return p
}

// MultipleAnswers allows users to select multiple answers.
func (p *Poll) MultipleAnswers() *Poll {
	p.opts.AllowsMultipleAnswers = true
	return p
}

// Protect enables content protection for the poll.
func (p *Poll) Protect() *Poll {
	p.opts.ProtectContent = true
	return p
}

// Quiz converts the poll to a quiz with the specified correct option index.
func (p *Poll) Quiz(correct int) *Poll {
	p.opts.Type = "quiz"
	p.opts.CorrectOptionId = int64(correct)
	return p
}

// Explanation sets an explanation text for quiz answers.
func (p *Poll) Explanation(text String) *Poll {
	p.opts.Explanation = text.Std()
	return p
}

// ExplanationHTML sets the explanation parse mode to HTML.
func (p *Poll) ExplanationHTML() *Poll {
	p.opts.ExplanationParseMode = "HTML"
	return p
}

// ExplanationMarkdown sets the explanation parse mode to MarkdownV2.
func (p *Poll) ExplanationMarkdown() *Poll {
	p.opts.ExplanationParseMode = "MarkdownV2"
	return p
}

// Silent disables notification for the poll.
func (p *Poll) Silent() *Poll {
	p.opts.DisableNotification = true
	return p
}

// CloseIn sets the poll to close after the specified number of seconds.
func (p *Poll) CloseIn(seconds int64) *Poll {
	p.opts.OpenPeriod = seconds
	return p
}

// CloseAt sets the poll to close at the specified timestamp.
func (p *Poll) CloseAt(timestamp int64) *Poll {
	p.opts.CloseDate = timestamp
	return p
}

// Closed marks the poll as already closed.
func (p *Poll) Closed() *Poll {
	p.opts.IsClosed = true
	return p
}

// Markup sets the reply markup keyboard for the poll.
func (p *Poll) Markup(kb keyboard.KeyboardBuilder) *Poll {
	p.opts.ReplyMarkup = kb.Markup()
	return p
}

// ReplyTo sets the message ID to reply to.
func (p *Poll) ReplyTo(messageID int64) *Poll {
	p.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return p
}

// Timeout sets the request timeout duration.
func (p *Poll) Timeout(duration time.Duration) *Poll {
	p.opts.RequestOpts = &gotgbot.RequestOpts{Timeout: duration}
	return p
}

// Send sends the poll to Telegram and returns the result.
func (p *Poll) Send() Result[*gotgbot.Message] {
	return p.ctx.timers(p.after, p.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := p.chatID.UnwrapOr(p.ctx.EffectiveChat.Id)
		return ResultOf(p.ctx.Bot.Raw().SendPoll(chatID, p.question.Std(), p.options, p.opts))
	})
}
