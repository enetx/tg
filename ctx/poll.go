package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"

	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/types/effects"

	. "github.com/enetx/g"
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

func (p *Poll) After(duration time.Duration) *Poll {
	p.after = Some(duration)
	return p
}

func (p *Poll) DeleteAfter(duration time.Duration) *Poll {
	p.deleteAfter = Some(duration)
	return p
}

func (p *Poll) To(id int64) *Poll {
	p.chatID = Some(id)
	return p
}

func (p *Poll) Option(text String) *Poll {
	opt := gotgbot.InputPollOption{Text: text.Std()}
	p.options.Push(opt)
	return p
}

func (p *Poll) Anonymous() *Poll {
	p.opts.IsAnonymous = true
	return p
}

func (p *Poll) Business(id String) *Poll {
	p.opts.BusinessConnectionId = id.Std()
	return p
}

func (p *Poll) Thread(id int64) *Poll {
	p.opts.MessageThreadId = id
	return p
}

func (p *Poll) AllowPaidBroadcast() *Poll {
	p.opts.AllowPaidBroadcast = true
	return p
}

func (p *Poll) Effect(effect effects.EffectType) *Poll {
	p.opts.MessageEffectId = effect.String()
	return p
}

func (p *Poll) MultipleAnswers() *Poll {
	p.opts.AllowsMultipleAnswers = true
	return p
}

func (p *Poll) Protect() *Poll {
	p.opts.ProtectContent = true
	return p
}

func (p *Poll) Quiz(correct int) *Poll {
	p.opts.Type = "quiz"
	p.opts.CorrectOptionId = int64(correct)
	return p
}

func (p *Poll) Explanation(text String) *Poll {
	p.opts.Explanation = text.Std()
	return p
}

func (p *Poll) ExplanationHTML() *Poll {
	p.opts.ExplanationParseMode = "HTML"
	return p
}

func (p *Poll) ExplanationMarkdown() *Poll {
	p.opts.ExplanationParseMode = "MarkdownV2"
	return p
}

func (p *Poll) Silent() *Poll {
	p.opts.DisableNotification = true
	return p
}

func (p *Poll) CloseIn(seconds int64) *Poll {
	p.opts.OpenPeriod = seconds
	return p
}

func (p *Poll) CloseAt(timestamp int64) *Poll {
	p.opts.CloseDate = timestamp
	return p
}

func (p *Poll) Closed() *Poll {
	p.opts.IsClosed = true
	return p
}

func (p *Poll) Markup(kb keyboard.KeyboardBuilder) *Poll {
	p.opts.ReplyMarkup = kb.Markup()
	return p
}

func (p *Poll) ReplyTo(messageID int64) *Poll {
	p.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return p
}

func (p *Poll) Timeout(duration time.Duration) *Poll {
	p.opts.RequestOpts = &gotgbot.RequestOpts{Timeout: duration}
	return p
}

func (p *Poll) Send() Result[*gotgbot.Message] {
	return p.ctx.timers(p.after, p.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := p.chatID.UnwrapOr(p.ctx.EffectiveChat.Id)
		return ResultOf(p.ctx.Bot.Raw().SendPoll(chatID, p.question.Std(), p.options, p.opts))
	})
}
