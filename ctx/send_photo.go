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

type SendPhoto struct {
	ctx         *Context
	doc         gotgbot.InputFileOrString
	opts        *gotgbot.SendPhotoOpts
	file        *g.File
	chatID      g.Option[int64]
	after       g.Option[time.Duration]
	deleteAfter g.Option[time.Duration]
	err         error
}

// CaptionEntities sets custom entities for the photo caption.
func (sp *SendPhoto) CaptionEntities(e *entities.Entities) *SendPhoto {
	sp.opts.CaptionEntities = e.Std()
	return sp
}

// After schedules the photo to be sent after the specified duration.
func (sp *SendPhoto) After(duration time.Duration) *SendPhoto {
	sp.after = g.Some(duration)
	return sp
}

// DeleteAfter schedules the photo message to be deleted after the specified duration.
func (sp *SendPhoto) DeleteAfter(duration time.Duration) *SendPhoto {
	sp.deleteAfter = g.Some(duration)
	return sp
}

// Spoiler marks the photo as a spoiler.
func (sp *SendPhoto) Spoiler() *SendPhoto {
	sp.opts.HasSpoiler = true
	return sp
}

// Caption sets the caption text for the photo.
func (sp *SendPhoto) Caption(caption g.String) *SendPhoto {
	sp.opts.Caption = caption.Std()
	return sp
}

// HTML sets the caption parse mode to HTML.
func (sp *SendPhoto) HTML() *SendPhoto {
	sp.opts.ParseMode = "HTML"
	return sp
}

// Markdown sets the caption parse mode to MarkdownV2.
func (sp *SendPhoto) Markdown() *SendPhoto {
	sp.opts.ParseMode = "MarkdownV2"
	return sp
}

// Silent disables notification for the photo message.
func (sp *SendPhoto) Silent() *SendPhoto {
	sp.opts.DisableNotification = true
	return sp
}

// Protect enables content protection for the photo message.
func (sp *SendPhoto) Protect() *SendPhoto {
	sp.opts.ProtectContent = true
	return sp
}

// AllowPaidBroadcast allows the message to be sent in paid broadcast channels.
func (sp *SendPhoto) AllowPaidBroadcast() *SendPhoto {
	sp.opts.AllowPaidBroadcast = true
	return sp
}

// Effect sets a message effect for the message.
func (sp *SendPhoto) Effect(effect effects.EffectType) *SendPhoto {
	sp.opts.MessageEffectId = effect.String()
	return sp
}

// Markup sets the reply markup keyboard for the photo message.
func (sp *SendPhoto) Markup(kb keyboard.Keyboard) *SendPhoto {
	sp.opts.ReplyMarkup = kb.Markup()
	return sp
}

// Reply sets reply parameters using the reply builder.
func (sp *SendPhoto) Reply(params *reply.Parameters) *SendPhoto {
	if params != nil {
		sp.opts.ReplyParameters = params.Std()
	}
	return sp
}

// Timeout sets a custom timeout for this request.
func (sp *SendPhoto) Timeout(duration time.Duration) *SendPhoto {
	if sp.opts.RequestOpts == nil {
		sp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sp.opts.RequestOpts.Timeout = duration

	return sp
}

// APIURL sets a custom API URL for this request.
func (sp *SendPhoto) APIURL(url g.String) *SendPhoto {
	if sp.opts.RequestOpts == nil {
		sp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sp.opts.RequestOpts.APIURL = url.Std()

	return sp
}

// Business sets the business connection ID for the photo message.
func (sp *SendPhoto) Business(id g.String) *SendPhoto {
	sp.opts.BusinessConnectionId = id.Std()
	return sp
}

// Thread sets the message thread ID for the photo message.
func (sp *SendPhoto) Thread(id int64) *SendPhoto {
	sp.opts.MessageThreadId = id
	return sp
}

// ShowCaptionAboveMedia displays the caption above the photo instead of below.
func (sp *SendPhoto) ShowCaptionAboveMedia() *SendPhoto {
	sp.opts.ShowCaptionAboveMedia = true
	return sp
}

// SuggestedPost sets suggested post parameters for direct messages chats.
func (sp *SendPhoto) SuggestedPost(params *suggested.PostParameters) *SendPhoto {
	if params != nil {
		sp.opts.SuggestedPostParameters = params.Std()
	}
	return sp
}

// To sets the target chat ID for the photo message.
func (sp *SendPhoto) To(chatID int64) *SendPhoto {
	sp.chatID = g.Some(chatID)
	return sp
}

// DirectMessagesTopic sets the direct messages topic ID for the message.
func (sp *SendPhoto) DirectMessagesTopic(topicID int64) *SendPhoto {
	sp.opts.DirectMessagesTopicId = topicID
	return sp
}

// Send sends the photo message to Telegram and returns the result.
func (sp *SendPhoto) Send() g.Result[*gotgbot.Message] {
	if sp.err != nil {
		return g.Err[*gotgbot.Message](sp.err)
	}

	if sp.file != nil {
		defer sp.file.Close()
	}

	return sp.ctx.timers(sp.after, sp.deleteAfter, func() g.Result[*gotgbot.Message] {
		chatID := sp.chatID.UnwrapOr(sp.ctx.EffectiveChat.Id)
		return g.ResultOf(sp.ctx.Bot.Raw().SendPhoto(chatID, sp.doc, sp.opts))
	})
}
