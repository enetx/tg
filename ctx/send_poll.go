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
func (sp *SendPoll) QuestionEntities(e *entities.Entities) *SendPoll {
	sp.opts.QuestionEntities = e.Std()
	return sp
}

// ExplanationEntities sets custom entities for the poll explanation.
func (sp *SendPoll) ExplanationEntities(e *entities.Entities) *SendPoll {
	sp.opts.ExplanationEntities = e.Std()
	return sp
}

// After schedules the poll to be sent after the specified duration.
func (sp *SendPoll) After(duration time.Duration) *SendPoll {
	sp.after = Some(duration)
	return sp
}

// DeleteAfter schedules the poll message to be deleted after the specified duration.
func (sp *SendPoll) DeleteAfter(duration time.Duration) *SendPoll {
	sp.deleteAfter = Some(duration)
	return sp
}

// To sets the target chat ID for the poll.
func (sp *SendPoll) To(id int64) *SendPoll {
	sp.chatID = Some(id)
	return sp
}

// Option adds a poll option with the specified text.
func (sp *SendPoll) Option(text String) *SendPoll {
	opt := gotgbot.InputPollOption{Text: text.Std()}
	sp.options.Push(opt)
	return sp
}

// Anonymous makes the poll anonymous.
func (sp *SendPoll) Anonymous() *SendPoll {
	sp.opts.IsAnonymous = true
	return sp
}

// Business sets the business connection ID for the poll.
func (sp *SendPoll) Business(id String) *SendPoll {
	sp.opts.BusinessConnectionId = id.Std()
	return sp
}

// Thread sets the message thread ID for the poll.
func (sp *SendPoll) Thread(id int64) *SendPoll {
	sp.opts.MessageThreadId = id
	return sp
}

// AllowPaidBroadcast allows the poll to be sent in paid broadcast channels.
func (sp *SendPoll) AllowPaidBroadcast() *SendPoll {
	sp.opts.AllowPaidBroadcast = true
	return sp
}

// Effect sets a message effect for the poll.
func (sp *SendPoll) Effect(effect effects.EffectType) *SendPoll {
	sp.opts.MessageEffectId = effect.String()
	return sp
}

// MultipleAnswers allows users to select multiple answers.
func (sp *SendPoll) MultipleAnswers() *SendPoll {
	sp.opts.AllowsMultipleAnswers = true
	return sp
}

// Protect enables content protection for the poll.
func (sp *SendPoll) Protect() *SendPoll {
	sp.opts.ProtectContent = true
	return sp
}

// Quiz converts the poll to a quiz with the specified correct option index.
func (sp *SendPoll) Quiz(correct int) *SendPoll {
	sp.opts.Type = "quiz"
	sp.opts.CorrectOptionId = int64(correct)
	return sp
}

// Explanation sets an explanation text for quiz answers.
func (sp *SendPoll) Explanation(text String) *SendPoll {
	sp.opts.Explanation = text.Std()
	return sp
}

// ExplanationHTML sets the explanation parse mode to HTML.
func (sp *SendPoll) ExplanationHTML() *SendPoll {
	sp.opts.ExplanationParseMode = "HTML"
	return sp
}

// ExplanationMarkdown sets the explanation parse mode to MarkdownV2.
func (sp *SendPoll) ExplanationMarkdown() *SendPoll {
	sp.opts.ExplanationParseMode = "MarkdownV2"
	return sp
}

// Silent disables notification for the poll.
func (sp *SendPoll) Silent() *SendPoll {
	sp.opts.DisableNotification = true
	return sp
}

// ClosesIn sets the poll to close after the specified duration.
func (sp *SendPoll) ClosesIn(duration time.Duration) *SendPoll {
	sp.opts.OpenPeriod = int64(duration.Seconds())
	return sp
}

// ClosesAt sets the poll to close at the specified time.
func (sp *SendPoll) ClosesAt(t time.Time) *SendPoll {
	sp.opts.CloseDate = t.Unix()
	return sp
}

// Closed marks the poll as already closed.
func (sp *SendPoll) Closed() *SendPoll {
	sp.opts.IsClosed = true
	return sp
}

// Markup sets the reply markup keyboard for the poll.
func (sp *SendPoll) Markup(kb keyboard.KeyboardBuilder) *SendPoll {
	sp.opts.ReplyMarkup = kb.Markup()
	return sp
}

// ReplyTo sets the message ID to reply to.
func (sp *SendPoll) ReplyTo(messageID int64) *SendPoll {
	sp.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return sp
}

// Timeout sets a custom timeout for this request.
func (sp *SendPoll) Timeout(duration time.Duration) *SendPoll {
	if sp.opts.RequestOpts == nil {
		sp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sp.opts.RequestOpts.Timeout = duration

	return sp
}

// APIURL sets a custom API URL for this request.
func (sp *SendPoll) APIURL(url String) *SendPoll {
	if sp.opts.RequestOpts == nil {
		sp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sp.opts.RequestOpts.APIURL = url.Std()

	return sp
}

// Send sends the poll to Telegram and returns the result.
func (sp *SendPoll) Send() Result[*gotgbot.Message] {
	return sp.ctx.timers(sp.after, sp.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := sp.chatID.UnwrapOr(sp.ctx.EffectiveChat.Id)
		return ResultOf(sp.ctx.Bot.Raw().SendPoll(chatID, sp.question.Std(), sp.options, sp.opts))
	})
}
