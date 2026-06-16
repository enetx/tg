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

// SendLivePhoto represents a request to send a live photo.
// A live photo consists of a short video (up to 10 seconds, 10 MB) and a static cover photo.
type SendLivePhoto struct {
	ctx         *Context
	livePhoto   gotgbot.InputFileOrString
	photo       gotgbot.InputFileOrString
	opts        *gotgbot.SendLivePhotoOpts
	livePhotoFD *g.File
	photoFD     *g.File
	chatID      g.Option[int64]
	after       g.Option[time.Duration]
	deleteAfter g.Option[time.Duration]
	err         error
}

// CaptionEntities sets custom entities for the live photo caption.
func (slp *SendLivePhoto) CaptionEntities(e *entities.Entities) *SendLivePhoto {
	slp.opts.CaptionEntities = e.Std()
	return slp
}

// After schedules the live photo to be sent after the specified duration.
func (slp *SendLivePhoto) After(duration time.Duration) *SendLivePhoto {
	slp.after = g.Some(duration)
	return slp
}

// DeleteAfter schedules the live photo message to be deleted after the specified duration.
func (slp *SendLivePhoto) DeleteAfter(duration time.Duration) *SendLivePhoto {
	slp.deleteAfter = g.Some(duration)
	return slp
}

// Spoiler marks the live photo as a spoiler.
func (slp *SendLivePhoto) Spoiler() *SendLivePhoto {
	slp.opts.HasSpoiler = true
	return slp
}

// Caption sets the caption text for the live photo.
func (slp *SendLivePhoto) Caption(caption g.String) *SendLivePhoto {
	slp.opts.Caption = caption.Std()
	return slp
}

// HTML sets the caption parse mode to HTML.
func (slp *SendLivePhoto) HTML() *SendLivePhoto {
	slp.opts.ParseMode = "HTML"
	return slp
}

// Markdown sets the caption parse mode to MarkdownV2.
func (slp *SendLivePhoto) Markdown() *SendLivePhoto {
	slp.opts.ParseMode = "MarkdownV2"
	return slp
}

// Silent disables notification for the live photo message.
func (slp *SendLivePhoto) Silent() *SendLivePhoto {
	slp.opts.DisableNotification = true
	return slp
}

// Protect enables content protection for the live photo message.
func (slp *SendLivePhoto) Protect() *SendLivePhoto {
	slp.opts.ProtectContent = true
	return slp
}

// AllowPaidBroadcast allows the message to be sent in paid broadcast channels.
func (slp *SendLivePhoto) AllowPaidBroadcast() *SendLivePhoto {
	slp.opts.AllowPaidBroadcast = true
	return slp
}

// Effect sets a message effect for the message.
func (slp *SendLivePhoto) Effect(effect effects.EffectType) *SendLivePhoto {
	slp.opts.MessageEffectId = effect.String()
	return slp
}

// Markup sets the reply markup keyboard for the live photo message.
func (slp *SendLivePhoto) Markup(kb keyboard.Keyboard) *SendLivePhoto {
	slp.opts.ReplyMarkup = kb.Markup()
	return slp
}

// Reply sets reply parameters using the reply builder.
func (slp *SendLivePhoto) Reply(params *reply.Parameters) *SendLivePhoto {
	if params != nil {
		slp.opts.ReplyParameters = params.Std()
	}
	return slp
}

// Timeout sets a custom timeout for this request.
func (slp *SendLivePhoto) Timeout(duration time.Duration) *SendLivePhoto {
	if slp.opts.RequestOpts == nil {
		slp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	slp.opts.RequestOpts.Timeout = duration

	return slp
}

// APIURL sets a custom API URL for this request.
func (slp *SendLivePhoto) APIURL(url g.String) *SendLivePhoto {
	if slp.opts.RequestOpts == nil {
		slp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	slp.opts.RequestOpts.APIURL = url.Std()

	return slp
}

// Business sets the business connection ID for the live photo message.
func (slp *SendLivePhoto) Business(id g.String) *SendLivePhoto {
	slp.opts.BusinessConnectionId = id.Std()
	return slp
}

// Thread sets the message thread ID for the live photo message.
func (slp *SendLivePhoto) Thread(id int64) *SendLivePhoto {
	slp.opts.MessageThreadId = id
	return slp
}

// ShowCaptionAboveMedia displays the caption above the live photo instead of below.
func (slp *SendLivePhoto) ShowCaptionAboveMedia() *SendLivePhoto {
	slp.opts.ShowCaptionAboveMedia = true
	return slp
}

// SuggestedPost sets suggested post parameters for direct messages chats.
func (slp *SendLivePhoto) SuggestedPost(params *suggested.PostParameters) *SendLivePhoto {
	if params != nil {
		slp.opts.SuggestedPostParameters = params.Std()
	}

	return slp
}

// To sets the target chat ID for the live photo message.
func (slp *SendLivePhoto) To(chatID int64) *SendLivePhoto {
	slp.chatID = g.Some(chatID)
	return slp
}

// DirectMessagesTopic sets the direct messages topic ID for the message.
func (slp *SendLivePhoto) DirectMessagesTopic(topicID int64) *SendLivePhoto {
	slp.opts.DirectMessagesTopicId = topicID
	return slp
}

// Send sends the live photo message to Telegram and returns the result.
func (slp *SendLivePhoto) Send() g.Result[*gotgbot.Message] {
	if slp.err != nil {
		return g.Err[*gotgbot.Message](slp.err)
	}

	if slp.livePhotoFD != nil {
		defer slp.livePhotoFD.Close()
	}

	if slp.photoFD != nil {
		defer slp.photoFD.Close()
	}

	return slp.ctx.timers(slp.after, slp.deleteAfter, func() g.Result[*gotgbot.Message] {
		chatID := slp.chatID.UnwrapOr(slp.ctx.EffectiveChat.Id)
		return g.ResultOf(slp.ctx.Bot.Raw().SendLivePhoto(chatID, slp.livePhoto, slp.photo, slp.opts))
	})
}
