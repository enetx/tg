package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
)

type Sticker struct {
	ctx         *Context
	doc         gotgbot.InputFileOrString
	opts        *gotgbot.SendStickerOpts
	file        *File
	chatID      Option[int64]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
	err         error
}

// After schedules the sticker to be sent after the specified duration.
func (s *Sticker) After(duration time.Duration) *Sticker {
	s.after = Some(duration)
	return s
}

// DeleteAfter schedules the sticker message to be deleted after the specified duration.
func (s *Sticker) DeleteAfter(duration time.Duration) *Sticker {
	s.deleteAfter = Some(duration)
	return s
}

// Silent disables notification for the sticker message.
func (s *Sticker) Silent() *Sticker {
	s.opts.DisableNotification = true
	return s
}

// Protect enables content protection for the sticker message.
func (s *Sticker) Protect() *Sticker {
	s.opts.ProtectContent = true
	return s
}

// Markup sets the reply markup keyboard for the sticker message.
func (s *Sticker) Markup(kb keyboard.KeyboardBuilder) *Sticker {
	s.opts.ReplyMarkup = kb.Markup()
	return s
}

// Emoji sets the emoji associated with the sticker.
func (s *Sticker) Emoji(emoji String) *Sticker {
	s.opts.Emoji = emoji.Std()
	return s
}

// ReplyTo sets the message ID to reply to.
func (s *Sticker) ReplyTo(messageID int64) *Sticker {
	s.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return s
}

// Timeout sets the request timeout duration.
func (s *Sticker) Timeout(duration time.Duration) *Sticker {
	s.opts.RequestOpts = &gotgbot.RequestOpts{Timeout: duration}
	return s
}

// Business sets the business connection ID for the sticker message.
func (s *Sticker) Business(id String) *Sticker {
	s.opts.BusinessConnectionId = id.Std()
	return s
}

// Thread sets the message thread ID for the sticker message.
func (s *Sticker) Thread(id int64) *Sticker {
	s.opts.MessageThreadId = id
	return s
}

// To sets the target chat ID for the sticker message.
func (s *Sticker) To(chatID int64) *Sticker {
	s.chatID = Some(chatID)
	return s
}

// Send sends the sticker message to Telegram and returns the result.
func (s *Sticker) Send() Result[*gotgbot.Message] {
	if s.err != nil {
		return Err[*gotgbot.Message](s.err)
	}

	if s.file != nil {
		defer s.file.Close()
	}

	return s.ctx.timers(s.after, s.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := s.chatID.UnwrapOr(s.ctx.EffectiveChat.Id)
		return ResultOf(s.ctx.Bot.Raw().SendSticker(chatID, s.doc, s.opts))
	})
}
