package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/types/effects"
)

type SendPoll struct {
	ctx         *Context
	question    String
	chatID      Option[int64]
	options     Slice[gotgbot.InputPollOption]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
	opts        *gotgbot.SendPollOpts
}

// QuestionEntities sets custom entities for the poll question.
func (c *SendPoll) QuestionEntities(e *entities.Entities) *SendPoll {
	c.opts.QuestionEntities = e.Std()
	return c
}

// ExplanationEntities sets custom entities for the poll explanation.
func (c *SendPoll) ExplanationEntities(e *entities.Entities) *SendPoll {
	c.opts.ExplanationEntities = e.Std()
	return c
}

// After schedules the poll to be sent after the specified duration.
func (c *SendPoll) After(duration time.Duration) *SendPoll {
	c.after = Some(duration)
	return c
}

// DeleteAfter schedules the poll message to be deleted after the specified duration.
func (c *SendPoll) DeleteAfter(duration time.Duration) *SendPoll {
	c.deleteAfter = Some(duration)
	return c
}

// To sets the target chat ID for the poll.
func (c *SendPoll) To(id int64) *SendPoll {
	c.chatID = Some(id)
	return c
}

// Option adds a poll option with the specified text.
func (c *SendPoll) Option(text String) *SendPoll {
	opt := gotgbot.InputPollOption{Text: text.Std()}
	c.options.Push(opt)
	return c
}

// Anonymous makes the poll anonymous.
func (c *SendPoll) Anonymous() *SendPoll {
	c.opts.IsAnonymous = true
	return c
}

// Business sets the business connection ID for the poll.
func (c *SendPoll) Business(id String) *SendPoll {
	c.opts.BusinessConnectionId = id.Std()
	return c
}

// Thread sets the message thread ID for the poll.
func (c *SendPoll) Thread(id int64) *SendPoll {
	c.opts.MessageThreadId = id
	return c
}

// AllowPaidBroadcast allows the poll to be sent in paid broadcast channels.
func (c *SendPoll) AllowPaidBroadcast() *SendPoll {
	c.opts.AllowPaidBroadcast = true
	return c
}

// Effect sets a message effect for the poll.
func (c *SendPoll) Effect(effect effects.EffectType) *SendPoll {
	c.opts.MessageEffectId = effect.String()
	return c
}

// MultipleAnswers allows users to select multiple answers.
func (c *SendPoll) MultipleAnswers() *SendPoll {
	c.opts.AllowsMultipleAnswers = true
	return c
}

// Protect enables content protection for the poll.
func (c *SendPoll) Protect() *SendPoll {
	c.opts.ProtectContent = true
	return c
}

// Quiz converts the poll to a quiz with the specified correct option index.
func (c *SendPoll) Quiz(correct int) *SendPoll {
	c.opts.Type = "quiz"
	c.opts.CorrectOptionId = int64(correct)
	return c
}

// Explanation sets an explanation text for quiz answers.
func (c *SendPoll) Explanation(text String) *SendPoll {
	c.opts.Explanation = text.Std()
	return c
}

// ExplanationHTML sets the explanation parse mode to HTML.
func (c *SendPoll) ExplanationHTML() *SendPoll {
	c.opts.ExplanationParseMode = "HTML"
	return c
}

// ExplanationMarkdown sets the explanation parse mode to MarkdownV2.
func (c *SendPoll) ExplanationMarkdown() *SendPoll {
	c.opts.ExplanationParseMode = "MarkdownV2"
	return c
}

// Silent disables notification for the poll.
func (c *SendPoll) Silent() *SendPoll {
	c.opts.DisableNotification = true
	return c
}

// CloseIn sets the poll to close after the specified number of seconds.
func (c *SendPoll) CloseIn(seconds int64) *SendPoll {
	c.opts.OpenPeriod = seconds
	return c
}

// CloseAt sets the poll to close at the specified timestamp.
func (c *SendPoll) CloseAt(timestamp int64) *SendPoll {
	c.opts.CloseDate = timestamp
	return c
}

// Closed marks the poll as already closed.
func (c *SendPoll) Closed() *SendPoll {
	c.opts.IsClosed = true
	return c
}

// Markup sets the reply markup keyboard for the poll.
func (c *SendPoll) Markup(kb keyboard.KeyboardBuilder) *SendPoll {
	c.opts.ReplyMarkup = kb.Markup()
	return c
}

// ReplyTo sets the message ID to reply to.
func (c *SendPoll) ReplyTo(messageID int64) *SendPoll {
	c.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return c
}

// Timeout sets a custom timeout for this request.
func (c *SendPoll) Timeout(duration time.Duration) *SendPoll {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *SendPoll) APIURL(url String) *SendPoll {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Send sends the poll to Telegram and returns the result.
func (c *SendPoll) Send() Result[*gotgbot.Message] {
	return c.ctx.timers(c.after, c.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := c.chatID.UnwrapOr(c.ctx.EffectiveChat.Id)
		return ResultOf(c.ctx.Bot.Raw().SendPoll(chatID, c.question.Std(), c.options, c.opts))
	})
}
