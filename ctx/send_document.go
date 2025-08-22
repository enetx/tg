package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/reply"
	"github.com/enetx/tg/suggested"
)

type SendDocument struct {
	ctx         *Context
	doc         gotgbot.InputFileOrString
	opts        *gotgbot.SendDocumentOpts
	file        *g.File
	thumb       *g.File
	chatID      g.Option[int64]
	after       g.Option[time.Duration]
	deleteAfter g.Option[time.Duration]
	err         error
}

// CaptionEntities sets custom entities for the document caption.
func (sd *SendDocument) CaptionEntities(e *entities.Entities) *SendDocument {
	sd.opts.CaptionEntities = e.Std()
	return sd
}

// After schedules the document to be sent after the specified duration.
func (sd *SendDocument) After(duration time.Duration) *SendDocument {
	sd.after = g.Some(duration)
	return sd
}

// DeleteAfter schedules the document message to be deleted after the specified duration.
func (sd *SendDocument) DeleteAfter(duration time.Duration) *SendDocument {
	sd.deleteAfter = g.Some(duration)
	return sd
}

// Caption sets the caption text for the document.
func (sd *SendDocument) Caption(caption g.String) *SendDocument {
	sd.opts.Caption = caption.Std()
	return sd
}

// HTML sets the caption parse mode to HTML.
func (sd *SendDocument) HTML() *SendDocument {
	sd.opts.ParseMode = "HTML"
	return sd
}

// Markdown sets the caption parse mode to MarkdownV2.
func (sd *SendDocument) Markdown() *SendDocument {
	sd.opts.ParseMode = "MarkdownV2"
	return sd
}

// Silent disables notification for the document message.
func (sd *SendDocument) Silent() *SendDocument {
	sd.opts.DisableNotification = true
	return sd
}

// Protect enables content protection for the document message.
func (sd *SendDocument) Protect() *SendDocument {
	sd.opts.ProtectContent = true
	return sd
}

// Markup sets the reply markup keyboard for the document message.
func (sd *SendDocument) Markup(kb keyboard.Keyboard) *SendDocument {
	sd.opts.ReplyMarkup = kb.Markup()
	return sd
}

// Thumbnail sets a custom thumbnail for the document.
func (sd *SendDocument) Thumbnail(file g.String) *SendDocument {
	sd.thumb = g.NewFile(file)

	reader := sd.thumb.Open()
	if reader.IsErr() {
		sd.err = reader.Err()
		return sd
	}

	sd.opts.Thumbnail = gotgbot.InputFileByReader(sd.thumb.Name().Std(), reader.Ok().Std())
	return sd
}

// Reply sets reply parameters using the reply builder.
func (sd *SendDocument) Reply(params *reply.Parameters) *SendDocument {
	sd.opts.ReplyParameters = params.Std()
	return sd
}

// Timeout sets a custom timeout for this request.
func (sd *SendDocument) Timeout(duration time.Duration) *SendDocument {
	if sd.opts.RequestOpts == nil {
		sd.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sd.opts.RequestOpts.Timeout = duration

	return sd
}

// APIURL sets a custom API URL for this request.
func (sd *SendDocument) APIURL(url g.String) *SendDocument {
	if sd.opts.RequestOpts == nil {
		sd.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sd.opts.RequestOpts.APIURL = url.Std()

	return sd
}

// Business sets the business connection ID for the document message.
func (sd *SendDocument) Business(id g.String) *SendDocument {
	sd.opts.BusinessConnectionId = id.Std()
	return sd
}

// Thread sets the message thread ID for the document message.
func (sd *SendDocument) Thread(id int64) *SendDocument {
	sd.opts.MessageThreadId = id
	return sd
}

// DisableContentTypeDetection disables automatic content type detection for the document.
func (sd *SendDocument) DisableContentTypeDetection() *SendDocument {
	sd.opts.DisableContentTypeDetection = true
	return sd
}

// SuggestedPost sets suggested post parameters for direct messages chats.
func (sd *SendDocument) SuggestedPost(params *suggested.PostParameters) *SendDocument {
	if params != nil {
		sd.opts.SuggestedPostParameters = params.Std()
	}
	return sd
}

// To sets the target chat ID for the document message.
func (sd *SendDocument) To(chatID int64) *SendDocument {
	sd.chatID = g.Some(chatID)
	return sd
}

// DirectMessagesTopic sets the direct messages topic ID for the message.
func (sd *SendDocument) DirectMessagesTopic(topicID int64) *SendDocument {
	sd.opts.DirectMessagesTopicId = topicID
	return sd
}

// Send sends the document message to Telegram and returns the result.
func (sd *SendDocument) Send() g.Result[*gotgbot.Message] {
	if sd.err != nil {
		return g.Err[*gotgbot.Message](sd.err)
	}

	if sd.file != nil {
		defer sd.file.Close()
	}

	if sd.thumb != nil {
		defer sd.thumb.Close()
	}

	return sd.ctx.timers(sd.after, sd.deleteAfter, func() g.Result[*gotgbot.Message] {
		chatID := sd.chatID.UnwrapOr(sd.ctx.EffectiveChat.Id)
		return g.ResultOf(sd.ctx.Bot.Raw().SendDocument(chatID, sd.doc, sd.opts))
	})
}
