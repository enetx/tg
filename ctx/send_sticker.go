package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
)

type SendSticker struct {
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
func (ss *SendSticker) After(duration time.Duration) *SendSticker {
	ss.after = Some(duration)
	return ss
}

// DeleteAfter schedules the sticker message to be deleted after the specified duration.
func (ss *SendSticker) DeleteAfter(duration time.Duration) *SendSticker {
	ss.deleteAfter = Some(duration)
	return ss
}

// Silent disables notification for the sticker message.
func (ss *SendSticker) Silent() *SendSticker {
	ss.opts.DisableNotification = true
	return ss
}

// Protect enables content protection for the sticker message.
func (ss *SendSticker) Protect() *SendSticker {
	ss.opts.ProtectContent = true
	return ss
}

// Markup sets the reply markup keyboard for the sticker message.
func (ss *SendSticker) Markup(kb keyboard.Keyboard) *SendSticker {
	ss.opts.ReplyMarkup = kb.Markup()
	return ss
}

// Emoji sets the emoji associated with the sticker.
func (ss *SendSticker) Emoji(emoji String) *SendSticker {
	ss.opts.Emoji = emoji.Std()
	return ss
}

// ReplyTo sets the message ID to reply to.
func (ss *SendSticker) ReplyTo(messageID int64) *SendSticker {
	ss.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return ss
}

// Timeout sets a custom timeout for this request.
func (ss *SendSticker) Timeout(duration time.Duration) *SendSticker {
	if ss.opts.RequestOpts == nil {
		ss.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ss.opts.RequestOpts.Timeout = duration

	return ss
}

// APIURL sets a custom API URL for this request.
func (ss *SendSticker) APIURL(url String) *SendSticker {
	if ss.opts.RequestOpts == nil {
		ss.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ss.opts.RequestOpts.APIURL = url.Std()

	return ss
}

// Business sets the business connection ID for the sticker message.
func (ss *SendSticker) Business(id String) *SendSticker {
	ss.opts.BusinessConnectionId = id.Std()
	return ss
}

// Thread sets the message thread ID for the sticker message.
func (ss *SendSticker) Thread(id int64) *SendSticker {
	ss.opts.MessageThreadId = id
	return ss
}

// To sets the target chat ID for the sticker message.
func (ss *SendSticker) To(chatID int64) *SendSticker {
	ss.chatID = Some(chatID)
	return ss
}

// Send sends the sticker message to Telegram and returns the result.
func (ss *SendSticker) Send() Result[*gotgbot.Message] {
	if ss.err != nil {
		return Err[*gotgbot.Message](ss.err)
	}

	if ss.file != nil {
		defer ss.file.Close()
	}

	return ss.ctx.timers(ss.after, ss.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := ss.chatID.UnwrapOr(ss.ctx.EffectiveChat.Id)
		return ResultOf(ss.ctx.Bot.Raw().SendSticker(chatID, ss.doc, ss.opts))
	})
}
