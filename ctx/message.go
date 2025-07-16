package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/preview"
	"github.com/enetx/tg/types/effects"
)

type Message struct {
	ctx         *Context
	text        String
	chatID      Option[int64]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
	opts        *gotgbot.SendMessageOpts
}

func (m *Message) After(duration time.Duration) *Message {
	m.after = Some(duration)
	return m
}

func (m *Message) DeleteAfter(duration time.Duration) *Message {
	m.deleteAfter = Some(duration)
	return m
}

func (m *Message) To(chatID int64) *Message {
	m.chatID = Some(chatID)
	return m
}

func (m *Message) HTML() *Message {
	m.opts.ParseMode = "HTML"
	return m
}

func (m *Message) Markdown() *Message {
	m.opts.ParseMode = "MarkdownV2"
	return m
}

func (m *Message) Silent() *Message {
	m.opts.DisableNotification = true
	return m
}

func (m *Message) Effect(effect effects.EffectType) *Message {
	m.opts.MessageEffectId = effect.String()
	return m
}

func (m *Message) ReplyTo(messageID int64) *Message {
	m.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return m
}

func (m *Message) Markup(kb keyboard.KeyboardBuilder) *Message {
	m.opts.ReplyMarkup = kb.Markup()
	return m
}

func (m *Message) AllowPaidBroadcast() *Message {
	m.opts.AllowPaidBroadcast = true
	return m
}

func (m *Message) Thread(id int64) *Message {
	m.opts.MessageThreadId = id
	return m
}

func (m *Message) ForceReply() *Message {
	m.opts.ReplyMarkup = gotgbot.ForceReply{ForceReply: true}
	return m
}

func (m *Message) RemoveKeyboard() *Message {
	m.opts.ReplyMarkup = gotgbot.ReplyKeyboardRemove{RemoveKeyboard: true}
	return m
}

func (m *Message) Preview(p *preview.Preview) *Message {
	m.opts.LinkPreviewOptions = p.Std()
	return m
}

func (m *Message) Send() Result[*gotgbot.Message] {
	return m.ctx.timers(m.after, m.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := m.chatID.UnwrapOr(m.ctx.EffectiveChat.Id)
		return ResultOf(m.ctx.Bot.Raw().SendMessage(chatID, m.text.Std(), m.opts))
	})
}
