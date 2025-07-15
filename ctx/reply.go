package ctx

import (
	"github.com/PaulSonOfLars/gotgbot/v2"

	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/preview"
	"github.com/enetx/tg/types/effects"

	. "github.com/enetx/g"
)

type Reply struct {
	ctx  *Context
	text String
	opts *gotgbot.SendMessageOpts
}

func (r *Reply) HTML() *Reply {
	r.opts.ParseMode = "HTML"
	return r
}

func (r *Reply) Markdown() *Reply {
	r.opts.ParseMode = "MarkdownV2"
	return r
}

func (r *Reply) Silent() *Reply {
	r.opts.DisableNotification = true
	return r
}

func (r *Reply) Effect(effect effects.EffectType) *Reply {
	r.opts.MessageEffectId = effect.String()
	return r
}

func (r *Reply) ReplyTo(id int64) *Reply {
	r.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: id}
	return r
}

func (r *Reply) Markup(kb keyboard.KeyboardBuilder) *Reply {
	r.opts.ReplyMarkup = kb.Markup()
	return r
}

func (r *Reply) AllowPaidBroadcast() *Reply {
	r.opts.AllowPaidBroadcast = true
	return r
}

func (r *Reply) Thread(id int64) *Reply {
	r.opts.MessageThreadId = id
	return r
}

func (r *Reply) ForceReply() *Reply {
	r.opts.ReplyMarkup = gotgbot.ForceReply{ForceReply: true}
	return r
}

func (r *Reply) RemoveKeyboard() *Reply {
	r.opts.ReplyMarkup = gotgbot.ReplyKeyboardRemove{RemoveKeyboard: true}
	return r
}

func (r *Reply) Preview(preview *preview.Preview) *Reply {
	r.opts.LinkPreviewOptions = preview.Std()
	return r
}

func (r *Reply) Send() Result[*gotgbot.Message] {
	return ResultOf(r.ctx.EffectiveMessage.Reply(r.ctx.Bot.Raw(), r.text.Std(), r.opts))
}
