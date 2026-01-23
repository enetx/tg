package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/reply"
	"github.com/enetx/tg/suggested"
	"github.com/enetx/tg/types/effects"
)

type SendVoice struct {
	ctx         *Context
	doc         gotgbot.InputFileOrString
	opts        *gotgbot.SendVoiceOpts
	file        *g.File
	chatID      g.Option[int64]
	after       g.Option[time.Duration]
	deleteAfter g.Option[time.Duration]
	err         error
}

// CaptionEntities sets custom entities for the voice caption.
func (sv *SendVoice) CaptionEntities(e *entities.Entities) *SendVoice {
	sv.opts.CaptionEntities = e.Std()
	return sv
}

// After schedules the voice message to be sent after the specified duration.
func (sv *SendVoice) After(duration time.Duration) *SendVoice {
	sv.after = g.Some(duration)
	return sv
}

// DeleteAfter schedules the voice message to be deleted after the specified duration.
func (sv *SendVoice) DeleteAfter(duration time.Duration) *SendVoice {
	sv.deleteAfter = g.Some(duration)
	return sv
}

// Caption sets the caption text for the voice message.
func (sv *SendVoice) Caption(caption g.String) *SendVoice {
	sv.opts.Caption = caption.Std()
	return sv
}

// HTML sets the caption parse mode to HTML.
func (sv *SendVoice) HTML() *SendVoice {
	sv.opts.ParseMode = "HTML"
	return sv
}

// Markdown sets the caption parse mode to MarkdownV2.
func (sv *SendVoice) Markdown() *SendVoice {
	sv.opts.ParseMode = "MarkdownV2"
	return sv
}

// Silent disables notification for the voice message.
func (sv *SendVoice) Silent() *SendVoice {
	sv.opts.DisableNotification = true
	return sv
}

// Protect enables content protection for the voice message.
func (sv *SendVoice) Protect() *SendVoice {
	sv.opts.ProtectContent = true
	return sv
}

// AllowPaidBroadcast allows the message to be sent in paid broadcast channels.
func (sv *SendVoice) AllowPaidBroadcast() *SendVoice {
	sv.opts.AllowPaidBroadcast = true
	return sv
}

// Effect sets a message effect for the message.
func (sv *SendVoice) Effect(effect effects.EffectType) *SendVoice {
	sv.opts.MessageEffectId = effect.String()
	return sv
}

// Markup sets the reply markup keyboard for the voice message.
func (sv *SendVoice) Markup(kb keyboard.Keyboard) *SendVoice {
	sv.opts.ReplyMarkup = kb.Markup()
	return sv
}

// Duration sets the voice message duration.
func (sv *SendVoice) Duration(duration time.Duration) *SendVoice {
	sv.opts.Duration = int64(duration.Seconds())
	return sv
}

// Reply sets reply parameters using the reply builder.
func (sv *SendVoice) Reply(params *reply.Parameters) *SendVoice {
	if params != nil {
		sv.opts.ReplyParameters = params.Std()
	}
	return sv
}

// Timeout sets a custom timeout for this request.
func (sv *SendVoice) Timeout(duration time.Duration) *SendVoice {
	if sv.opts.RequestOpts == nil {
		sv.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sv.opts.RequestOpts.Timeout = duration

	return sv
}

// APIURL sets a custom API URL for this request.
func (sv *SendVoice) APIURL(url g.String) *SendVoice {
	if sv.opts.RequestOpts == nil {
		sv.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sv.opts.RequestOpts.APIURL = url.Std()

	return sv
}

// Business sets the business connection ID for the voice message.
func (sv *SendVoice) Business(id g.String) *SendVoice {
	sv.opts.BusinessConnectionId = id.Std()
	return sv
}

// Thread sets the message thread ID for the voice message.
func (sv *SendVoice) Thread(id int64) *SendVoice {
	sv.opts.MessageThreadId = id
	return sv
}

// SuggestedPost sets suggested post parameters for direct messages chats.
func (sv *SendVoice) SuggestedPost(params *suggested.PostParameters) *SendVoice {
	if params != nil {
		sv.opts.SuggestedPostParameters = params.Std()
	}
	return sv
}

// To sets the target chat ID for the voice message.
func (sv *SendVoice) To(chatID int64) *SendVoice {
	sv.chatID = g.Some(chatID)
	return sv
}

// DirectMessagesTopic sets the direct messages topic ID for the message.
func (sv *SendVoice) DirectMessagesTopic(topicID int64) *SendVoice {
	sv.opts.DirectMessagesTopicId = topicID
	return sv
}

// Send sends the voice message to Telegram and returns the result.
func (sv *SendVoice) Send() g.Result[*gotgbot.Message] {
	if sv.err != nil {
		return g.Err[*gotgbot.Message](sv.err)
	}

	if sv.file != nil {
		defer sv.file.Close()
	}

	return sv.ctx.timers(sv.after, sv.deleteAfter, func() g.Result[*gotgbot.Message] {
		chatID := sv.chatID.UnwrapOr(sv.ctx.EffectiveChat.Id)
		return g.ResultOf(sv.ctx.Bot.Raw().SendVoice(chatID, sv.doc, sv.opts))
	})
}
