package ctx

import (
	"errors"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/reply"
	"github.com/enetx/tg/types/effects"
)

type MediaGroup struct {
	ctx         *Context
	media       g.Slice[input.Media]
	opts        *gotgbot.SendMediaGroupOpts
	chatID      g.Option[int64]
	after       g.Option[time.Duration]
	deleteAfter g.Option[time.Duration]
}

// After schedules the media group to be sent after the specified duration.
func (mg *MediaGroup) After(duration time.Duration) *MediaGroup {
	mg.after = g.Some(duration)
	return mg
}

// DeleteAfter schedules the media group messages to be deleted after the specified duration.
func (mg *MediaGroup) DeleteAfter(duration time.Duration) *MediaGroup {
	mg.deleteAfter = g.Some(duration)
	return mg
}

// Photo adds a photo to the media group with optional caption.
func (mg *MediaGroup) Photo(photo input.Media) *MediaGroup {
	mg.media.Push(photo)
	return mg
}

// Video creates a new media video builder.
func (mg *MediaGroup) Video(video input.Media) *MediaGroup {
	if _, ok := video.(*input.MediaVideo); ok {
		mg.media.Push(video)
	}

	return mg
}

// Audio adds an audio file to the media group with optional caption.
func (mg *MediaGroup) Audio(audio input.Media) *MediaGroup {
	if _, ok := audio.(*input.MediaAudio); ok {
		mg.media.Push(audio)
	}

	return mg
}

// Document adds a document to the media group with optional caption.
func (mg *MediaGroup) Document(document input.Media) *MediaGroup {
	if _, ok := document.(*input.MediaDocument); ok {
		mg.media.Push(document)
	}

	return mg
}

// Silent disables notification for the media group messages.
func (mg *MediaGroup) Silent() *MediaGroup {
	mg.opts.DisableNotification = true
	return mg
}

// Protect enables content protection for the media group messages.
func (mg *MediaGroup) Protect() *MediaGroup {
	mg.opts.ProtectContent = true
	return mg
}

// AllowPaidBroadcast allows the media group to be sent in paid broadcast channels.
func (mg *MediaGroup) AllowPaidBroadcast() *MediaGroup {
	mg.opts.AllowPaidBroadcast = true
	return mg
}

// Thread sets the message thread ID for the media group.
func (mg *MediaGroup) Thread(id int64) *MediaGroup {
	mg.opts.MessageThreadId = id
	return mg
}

// Effect sets a message effect for the media group.
func (mg *MediaGroup) Effect(effect effects.EffectType) *MediaGroup {
	mg.opts.MessageEffectId = effect.String()
	return mg
}

// Reply sets reply parameters using the reply builder.
func (mg *MediaGroup) Reply(params *reply.Parameters) *MediaGroup {
	mg.opts.ReplyParameters = params.Std()
	return mg
}

// Business sets the business connection ID for the media group.
func (mg *MediaGroup) Business(id g.String) *MediaGroup {
	mg.opts.BusinessConnectionId = id.Std()
	return mg
}

// To sets the target chat ID for the media group.
func (mg *MediaGroup) To(chatID int64) *MediaGroup {
	mg.chatID = g.Some(chatID)
	return mg
}

// Timeout sets a custom timeout for this request.
func (mg *MediaGroup) Timeout(duration time.Duration) *MediaGroup {
	if mg.opts.RequestOpts == nil {
		mg.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	mg.opts.RequestOpts.Timeout = duration

	return mg
}

// APIURL sets a custom API URL for this request.
func (mg *MediaGroup) APIURL(url g.String) *MediaGroup {
	if mg.opts.RequestOpts == nil {
		mg.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	mg.opts.RequestOpts.APIURL = url.Std()

	return mg
}

// DirectMessagesTopic sets the direct messages topic ID for the message.
func (mg *MediaGroup) DirectMessagesTopic(topicID int64) *MediaGroup {
	mg.opts.DirectMessagesTopicId = topicID
	return mg
}

// Send sends the media group to Telegram and returns the result.
func (mg *MediaGroup) Send() g.Result[g.Slice[gotgbot.Message]] {
	if mg.media.Len() == 0 {
		return g.Err[g.Slice[gotgbot.Message]](errors.New("no media added to media group"))
	}

	chatID := mg.chatID.UnwrapOr(mg.ctx.EffectiveChat.Id)
	media := g.TransformSlice(mg.media, input.Media.Build)

	return g.ResultOf[g.Slice[gotgbot.Message]](mg.ctx.Bot.Raw().SendMediaGroup(chatID, media, mg.opts))
}
