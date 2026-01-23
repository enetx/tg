package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/reply"
	"github.com/enetx/tg/suggested"
)

// SendPaidMedia represents a request to send paid media content.
type SendPaidMedia struct {
	ctx       *Context
	opts      *gotgbot.SendPaidMediaOpts
	chatID    g.Option[int64]
	starCount int64
	media     g.Slice[input.PaidMedia]
}

// SuggestedPost sets suggested post parameters for direct messages chats.
func (spm *SendPaidMedia) SuggestedPost(params *suggested.PostParameters) *SendPaidMedia {
	if params != nil {
		spm.opts.SuggestedPostParameters = params.Std()
	}
	return spm
}

// To sets the target chat ID for sending paid media.
func (spm *SendPaidMedia) To(chatID int64) *SendPaidMedia {
	spm.chatID = g.Some(chatID)
	return spm
}

// Photo adds a paid photo to the media list.
func (spm *SendPaidMedia) Photo(photo input.PaidMedia) *SendPaidMedia {
	if _, ok := photo.(*input.PaidMediaPhoto); ok {
		spm.media.Push(photo)
	}

	return spm
}

// Video adds a paid video to the media list.
func (spm *SendPaidMedia) Video(video input.PaidMedia) *SendPaidMedia {
	if _, ok := video.(*input.PaidMediaVideo); ok {
		spm.media.Push(video)
	}

	return spm
}

// Business sets the business connection ID for the paid media.
func (spm *SendPaidMedia) Business(businessConnectionID g.String) *SendPaidMedia {
	spm.opts.BusinessConnectionId = businessConnectionID.Std()
	return spm
}

// Payload sets the bot-defined payload for internal processing.
func (spm *SendPaidMedia) Payload(payload g.String) *SendPaidMedia {
	spm.opts.Payload = payload.Std()
	return spm
}

// Caption sets the media caption.
func (spm *SendPaidMedia) Caption(caption g.String) *SendPaidMedia {
	spm.opts.Caption = caption.Std()
	return spm
}

// HTML sets the caption parse mode to HTML.
func (spm *SendPaidMedia) HTML() *SendPaidMedia {
	spm.opts.ParseMode = "HTML"
	return spm
}

// Markdown sets the caption parse mode to Markdown.
func (spm *SendPaidMedia) Markdown() *SendPaidMedia {
	spm.opts.ParseMode = "MarkdownV2"
	return spm
}

// ShowCaptionAbove shows the caption above the media.
func (spm *SendPaidMedia) ShowCaptionAbove() *SendPaidMedia {
	spm.opts.ShowCaptionAboveMedia = true
	return spm
}

// Silent sends the message silently (no notification sound).
func (spm *SendPaidMedia) Silent() *SendPaidMedia {
	spm.opts.DisableNotification = true
	return spm
}

// Protect protects the media from forwarding and saving.
func (spm *SendPaidMedia) Protect() *SendPaidMedia {
	spm.opts.ProtectContent = true
	return spm
}

// AllowPaidBroadcast allows paid broadcast for high-speed delivery.
func (spm *SendPaidMedia) AllowPaidBroadcast() *SendPaidMedia {
	spm.opts.AllowPaidBroadcast = true
	return spm
}

// Reply sets reply parameters using the reply builder.
func (spm *SendPaidMedia) Reply(params *reply.Parameters) *SendPaidMedia {
	if params != nil {
		spm.opts.ReplyParameters = params.Std()
	}
	return spm
}

// Markup sets the reply markup keyboard.
func (spm *SendPaidMedia) Markup(kb keyboard.Keyboard) *SendPaidMedia {
	if markup, ok := kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
		spm.opts.ReplyMarkup = markup
	}

	return spm
}

// Timeout sets a custom timeout for this request.
func (spm *SendPaidMedia) Timeout(duration time.Duration) *SendPaidMedia {
	if spm.opts.RequestOpts == nil {
		spm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	spm.opts.RequestOpts.Timeout = duration

	return spm
}

// APIURL sets a custom API URL for this request.
func (spm *SendPaidMedia) APIURL(url g.String) *SendPaidMedia {
	if spm.opts.RequestOpts == nil {
		spm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	spm.opts.RequestOpts.APIURL = url.Std()

	return spm
}

// DirectMessagesTopic sets the direct messages topic ID for the message.
func (spm *SendPaidMedia) DirectMessagesTopic(topicID int64) *SendPaidMedia {
	spm.opts.DirectMessagesTopicId = topicID
	return spm
}

// Send sends the paid media and returns the message.
func (spm *SendPaidMedia) Send() g.Result[*gotgbot.Message] {
	if spm.media.IsEmpty() {
		return g.Err[*gotgbot.Message](g.Errorf("no paid media specified"))
	}

	if spm.media.Len() > 10 {
		return g.Err[*gotgbot.Message](g.Errorf("too many media items: {} (maximum 10)", spm.media.Len()))
	}

	if spm.starCount < 1 || spm.starCount > 10000 {
		return g.Err[*gotgbot.Message](g.Errorf("star count must be between 1-10000, got {}", spm.starCount))
	}

	chatID := spm.chatID.UnwrapOr(spm.ctx.EffectiveChat.Id)
	media := g.TransformSlice(spm.media, input.PaidMedia.Build)

	return g.ResultOf(spm.ctx.Bot.Raw().SendPaidMedia(chatID, spm.starCount, media, spm.opts))
}
