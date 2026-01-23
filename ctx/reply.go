package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/preview"
	"github.com/enetx/tg/reply"
	"github.com/enetx/tg/types/effects"
)

type Reply struct {
	ctx         *Context
	text        g.String
	opts        *gotgbot.SendMessageOpts
	after       g.Option[time.Duration]
	deleteAfter g.Option[time.Duration]
}

// Entities sets custom entities for the reply text.
func (r *Reply) Entities(e *entities.Entities) *Reply {
	r.opts.Entities = e.Std()
	return r
}

// After schedules the reply to be sent after the specified duration.
func (r *Reply) After(duration time.Duration) *Reply {
	r.after = g.Some(duration)
	return r
}

// DeleteAfter schedules the reply message to be deleted after the specified duration.
func (r *Reply) DeleteAfter(duration time.Duration) *Reply {
	r.deleteAfter = g.Some(duration)
	return r
}

// HTML sets the reply parse mode to HTML.
func (r *Reply) HTML() *Reply {
	r.opts.ParseMode = "HTML"
	return r
}

// Markdown sets the reply parse mode to MarkdownV2.
func (r *Reply) Markdown() *Reply {
	r.opts.ParseMode = "MarkdownV2"
	return r
}

// Silent disables notification for the reply message.
func (r *Reply) Silent() *Reply {
	r.opts.DisableNotification = true
	return r
}

// Effect sets a message effect for the reply.
func (r *Reply) Effect(effect effects.EffectType) *Reply {
	r.opts.MessageEffectId = effect.String()
	return r
}

// Replay sets reply parameters using the reply builder.
func (r *Reply) Reply(params *reply.Parameters) *Reply {
	if params != nil {
		r.opts.ReplyParameters = params.Std()
	}
	return r
}

// Markup sets the reply markup keyboard for the reply message.
func (r *Reply) Markup(kb keyboard.Keyboard) *Reply {
	r.opts.ReplyMarkup = kb.Markup()
	return r
}

// AllowPaidBroadcast allows the reply to be sent in paid broadcast channels.
func (r *Reply) AllowPaidBroadcast() *Reply {
	r.opts.AllowPaidBroadcast = true
	return r
}

// Thread sets the message thread ID for the reply.
func (r *Reply) Thread(id int64) *Reply {
	r.opts.MessageThreadId = id
	return r
}

// ForceReply forces users to reply to this message.
func (r *Reply) ForceReply() *Reply {
	r.opts.ReplyMarkup = gotgbot.ForceReply{ForceReply: true}
	return r
}

// RemoveKeyboard removes the custom keyboard.
func (r *Reply) RemoveKeyboard() *Reply {
	r.opts.ReplyMarkup = gotgbot.ReplyKeyboardRemove{RemoveKeyboard: true}
	return r
}

// Preview sets link preview options for the reply.
func (r *Reply) Preview(p *preview.Preview) *Reply {
	r.opts.LinkPreviewOptions = p.Std()
	return r
}

// Business sets the business connection ID for the reply.
func (r *Reply) Business(id g.String) *Reply {
	r.opts.BusinessConnectionId = id.Std()
	return r
}

// Protect enables content protection for the reply message.
func (r *Reply) Protect() *Reply {
	r.opts.ProtectContent = true
	return r
}

// Timeout sets a custom timeout for this request.
func (r *Reply) Timeout(duration time.Duration) *Reply {
	if r.opts.RequestOpts == nil {
		r.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	r.opts.RequestOpts.Timeout = duration

	return r
}

// APIURL sets a custom API URL for this request.
func (r *Reply) APIURL(url g.String) *Reply {
	if r.opts.RequestOpts == nil {
		r.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	r.opts.RequestOpts.APIURL = url.Std()

	return r
}

// Send sends the reply message and returns the result.
func (r *Reply) Send() g.Result[*gotgbot.Message] {
	return r.ctx.timers(r.after, r.deleteAfter, func() g.Result[*gotgbot.Message] {
		return g.ResultOf(r.ctx.EffectiveMessage.Reply(r.ctx.Bot.Raw(), r.text.Std(), r.opts))
	})
}
