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

// After schedules the message to be sent after the specified duration.
func (m *Message) After(duration time.Duration) *Message {
	m.after = Some(duration)
	return m
}

// DeleteAfter schedules the message to be deleted after the specified duration.
func (m *Message) DeleteAfter(duration time.Duration) *Message {
	m.deleteAfter = Some(duration)
	return m
}

// To sets the target chat ID for the message.
func (m *Message) To(chatID int64) *Message {
	m.chatID = Some(chatID)
	return m
}

// HTML sets the message parse mode to HTML.
func (m *Message) HTML() *Message {
	m.opts.ParseMode = "HTML"
	return m
}

// Markdown sets the message parse mode to MarkdownV2.
func (m *Message) Markdown() *Message {
	m.opts.ParseMode = "MarkdownV2"
	return m
}

// Silent disables notification for the message.
func (m *Message) Silent() *Message {
	m.opts.DisableNotification = true
	return m
}

// Effect sets a message effect for the message.
func (m *Message) Effect(effect effects.EffectType) *Message {
	m.opts.MessageEffectId = effect.String()
	return m
}

// ReplyTo sets the message ID to reply to.
func (m *Message) ReplyTo(messageID int64) *Message {
	m.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return m
}

// Markup sets the reply markup keyboard for the message.
func (m *Message) Markup(kb keyboard.KeyboardBuilder) *Message {
	m.opts.ReplyMarkup = kb.Markup()
	return m
}

// AllowPaidBroadcast allows the message to be sent in paid broadcast channels.
func (m *Message) AllowPaidBroadcast() *Message {
	m.opts.AllowPaidBroadcast = true
	return m
}

// Thread sets the message thread ID for the message.
func (m *Message) Thread(id int64) *Message {
	m.opts.MessageThreadId = id
	return m
}

// ForceReply forces users to reply to the message.
func (m *Message) ForceReply() *Message {
	m.opts.ReplyMarkup = gotgbot.ForceReply{ForceReply: true}
	return m
}

// RemoveKeyboard removes the custom keyboard.
func (m *Message) RemoveKeyboard() *Message {
	m.opts.ReplyMarkup = gotgbot.ReplyKeyboardRemove{RemoveKeyboard: true}
	return m
}

// Preview sets link preview options for the message.
func (m *Message) Preview(p *preview.Preview) *Message {
	m.opts.LinkPreviewOptions = p.Std()
	return m
}

// Business sets the business connection ID for the message.
func (m *Message) Business(id String) *Message {
	m.opts.BusinessConnectionId = id.Std()
	return m
}

// Protect enables content protection for the message.
func (m *Message) Protect() *Message {
	m.opts.ProtectContent = true
	return m
}

// Send sends the message to Telegram and returns the result.
func (m *Message) Send() Result[*gotgbot.Message] {
	return m.ctx.timers(m.after, m.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := m.chatID.UnwrapOr(m.ctx.EffectiveChat.Id)
		return ResultOf(m.ctx.Bot.Raw().SendMessage(chatID, m.text.Std(), m.opts))
	})
}
