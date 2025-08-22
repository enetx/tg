package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/reply"
	"github.com/enetx/tg/suggested"
	"github.com/enetx/tg/types/effects"
)

type SendSticker struct {
	ctx         *Context
	doc         gotgbot.InputFileOrString
	opts        *gotgbot.SendStickerOpts
	file        *g.File
	chatID      g.Option[int64]
	after       g.Option[time.Duration]
	deleteAfter g.Option[time.Duration]
	err         error
}

// After schedules the sticker to be sent after the specified duration.
func (ss *SendSticker) After(duration time.Duration) *SendSticker {
	ss.after = g.Some(duration)
	return ss
}

// DeleteAfter schedules the sticker message to be deleted after the specified duration.
func (ss *SendSticker) DeleteAfter(duration time.Duration) *SendSticker {
	ss.deleteAfter = g.Some(duration)
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

// AllowPaidBroadcast allows the message to be sent in paid broadcast channels.
func (ss *SendSticker) AllowPaidBroadcast() *SendSticker {
	ss.opts.AllowPaidBroadcast = true
	return ss
}

// Effect sets a message effect for the message.
func (ss *SendSticker) Effect(effect effects.EffectType) *SendSticker {
	ss.opts.MessageEffectId = effect.String()
	return ss
}

// Markup sets the reply markup keyboard for the sticker message.
func (ss *SendSticker) Markup(kb keyboard.Keyboard) *SendSticker {
	ss.opts.ReplyMarkup = kb.Markup()
	return ss
}

// Emoji sets the emoji associated with the sticker.
func (ss *SendSticker) Emoji(emoji g.String) *SendSticker {
	ss.opts.Emoji = emoji.Std()
	return ss
}

// Reply sets reply parameters using the reply builder.
func (ss *SendSticker) Reply(params *reply.Parameters) *SendSticker {
	ss.opts.ReplyParameters = params.Std()
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
func (ss *SendSticker) APIURL(url g.String) *SendSticker {
	if ss.opts.RequestOpts == nil {
		ss.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ss.opts.RequestOpts.APIURL = url.Std()

	return ss
}

// Business sets the business connection ID for the sticker message.
func (ss *SendSticker) Business(id g.String) *SendSticker {
	ss.opts.BusinessConnectionId = id.Std()
	return ss
}

// Thread sets the message thread ID for the sticker message.
func (ss *SendSticker) Thread(id int64) *SendSticker {
	ss.opts.MessageThreadId = id
	return ss
}

// SuggestedPost sets suggested post parameters for direct messages chats.
func (ss *SendSticker) SuggestedPost(params *suggested.PostParameters) *SendSticker {
	if params != nil {
		ss.opts.SuggestedPostParameters = params.Std()
	}
	return ss
}

// To sets the target chat ID for the sticker message.
func (ss *SendSticker) To(chatID int64) *SendSticker {
	ss.chatID = g.Some(chatID)
	return ss
}

// DirectMessagesTopic sets the direct messages topic ID for the message.
func (ss *SendSticker) DirectMessagesTopic(topicID int64) *SendSticker {
	ss.opts.DirectMessagesTopicId = topicID
	return ss
}

// Send sends the sticker message to Telegram and returns the result.
func (ss *SendSticker) Send() g.Result[*gotgbot.Message] {
	if ss.err != nil {
		return g.Err[*gotgbot.Message](ss.err)
	}

	if ss.file != nil {
		defer ss.file.Close()
	}

	return ss.ctx.timers(ss.after, ss.deleteAfter, func() g.Result[*gotgbot.Message] {
		chatID := ss.chatID.UnwrapOr(ss.ctx.EffectiveChat.Id)
		return g.ResultOf(ss.ctx.Bot.Raw().SendSticker(chatID, ss.doc, ss.opts))
	})
}
