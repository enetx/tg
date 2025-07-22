package ctx

import (
	"errors"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/internal/tgfile"
	"github.com/enetx/tg/types/effects"
)

type MediaGroup struct {
	ctx         *Context
	media       Slice[gotgbot.InputMedia]
	files       Slice[*File]
	opts        *gotgbot.SendMediaGroupOpts
	chatID      Option[int64]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
	err         error
}

// After schedules the media group to be sent after the specified duration.
func (c *MediaGroup) After(duration time.Duration) *MediaGroup {
	c.after = Some(duration)
	return c
}

// DeleteAfter schedules the media group messages to be deleted after the specified duration.
func (c *MediaGroup) DeleteAfter(duration time.Duration) *MediaGroup {
	c.deleteAfter = Some(duration)
	return c
}

// Photo adds a photo to the media group with optional caption.
func (c *MediaGroup) Photo(filename String, caption ...String) *MediaGroup {
	result := tgfile.ProcessFile(filename)
	if result.IsErr() {
		c.err = result.Err()
		return c
	}

	c.files.Push(result.Ok().File)

	media := gotgbot.InputMediaPhoto{Media: result.Ok().Doc}

	if len(caption) > 0 {
		media.Caption = caption[0].Std()
	}

	c.media.Push(media)

	return c
}

// Video adds a video to the media group with optional caption.
func (c *MediaGroup) Video(filename String, caption ...String) *MediaGroup {
	result := tgfile.ProcessFile(filename)
	if result.IsErr() {
		c.err = result.Err()
		return c
	}

	c.files.Push(result.Ok().File)

	media := gotgbot.InputMediaVideo{Media: result.Ok().Doc}

	if len(caption) > 0 {
		media.Caption = caption[0].Std()
	}

	c.media.Push(media)

	return c
}

// Audio adds an audio file to the media group with optional caption.
func (c *MediaGroup) Audio(filename String, caption ...String) *MediaGroup {
	result := tgfile.ProcessFile(filename)
	if result.IsErr() {
		c.err = result.Err()
		return c
	}

	c.files.Push(result.Ok().File)

	media := gotgbot.InputMediaAudio{Media: result.Ok().Doc}

	if len(caption) > 0 {
		media.Caption = caption[0].Std()
	}

	c.media.Push(media)

	return c
}

// Document adds a document to the media group with optional caption.
func (c *MediaGroup) Document(filename String, caption ...String) *MediaGroup {
	result := tgfile.ProcessFile(filename)
	if result.IsErr() {
		c.err = result.Err()
		return c
	}

	c.files.Push(result.Ok().File)

	media := gotgbot.InputMediaDocument{Media: result.Ok().Doc}

	if len(caption) > 0 {
		media.Caption = caption[0].Std()
	}

	c.media.Push(media)

	return c
}

// Silent disables notification for the media group messages.
func (c *MediaGroup) Silent() *MediaGroup {
	c.opts.DisableNotification = true
	return c
}

// Protect enables content protection for the media group messages.
func (c *MediaGroup) Protect() *MediaGroup {
	c.opts.ProtectContent = true
	return c
}

// AllowPaidBroadcast allows the media group to be sent in paid broadcast channels.
func (c *MediaGroup) AllowPaidBroadcast() *MediaGroup {
	c.opts.AllowPaidBroadcast = true
	return c
}

// Thread sets the message thread ID for the media group.
func (c *MediaGroup) Thread(id int64) *MediaGroup {
	c.opts.MessageThreadId = id
	return c
}

// Effect sets a message effect for the media group.
func (c *MediaGroup) Effect(effect effects.EffectType) *MediaGroup {
	c.opts.MessageEffectId = effect.String()
	return c
}

// ReplyTo sets the message ID to reply to.
func (c *MediaGroup) ReplyTo(messageID int64) *MediaGroup {
	c.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return c
}

// Business sets the business connection ID for the media group.
func (c *MediaGroup) Business(id String) *MediaGroup {
	c.opts.BusinessConnectionId = id.Std()
	return c
}

// To sets the target chat ID for the media group.
func (c *MediaGroup) To(chatID int64) *MediaGroup {
	c.chatID = Some(chatID)
	return c
}

// Timeout sets a custom timeout for this request.
func (c *MediaGroup) Timeout(duration time.Duration) *MediaGroup {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *MediaGroup) APIURL(url String) *MediaGroup {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Send sends the media group to Telegram and returns the result.
func (c *MediaGroup) Send() Result[Slice[gotgbot.Message]] {
	if c.err != nil {
		return Err[Slice[gotgbot.Message]](c.err)
	}

	if c.media.Len() == 0 {
		return Err[Slice[gotgbot.Message]](errors.New("no media added to media group"))
	}

	defer c.files.Iter().ForEach(func(file *File) { file.Close() })

	chatID := c.chatID.UnwrapOr(c.ctx.EffectiveChat.Id)

	return ResultOf[Slice[gotgbot.Message]](c.ctx.Bot.Raw().SendMediaGroup(chatID, c.media, c.opts))
}
