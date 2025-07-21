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
func (mg *MediaGroup) After(duration time.Duration) *MediaGroup {
	mg.after = Some(duration)
	return mg
}

// DeleteAfter schedules the media group messages to be deleted after the specified duration.
func (mg *MediaGroup) DeleteAfter(duration time.Duration) *MediaGroup {
	mg.deleteAfter = Some(duration)
	return mg
}

// Photo adds a photo to the media group with optional caption.
func (mg *MediaGroup) Photo(filename String, caption ...String) *MediaGroup {
	result := tgfile.ProcessFile(filename)
	if result.IsErr() {
		mg.err = result.Err()
		return mg
	}

	mg.files.Push(result.Ok().File)

	media := gotgbot.InputMediaPhoto{Media: result.Ok().Doc}

	if len(caption) > 0 {
		media.Caption = caption[0].Std()
	}

	mg.media.Push(media)

	return mg
}

// Video adds a video to the media group with optional caption.
func (mg *MediaGroup) Video(filename String, caption ...String) *MediaGroup {
	result := tgfile.ProcessFile(filename)
	if result.IsErr() {
		mg.err = result.Err()
		return mg
	}

	mg.files.Push(result.Ok().File)

	media := gotgbot.InputMediaVideo{Media: result.Ok().Doc}

	if len(caption) > 0 {
		media.Caption = caption[0].Std()
	}

	mg.media.Push(media)

	return mg
}

// Audio adds an audio file to the media group with optional caption.
func (mg *MediaGroup) Audio(filename String, caption ...String) *MediaGroup {
	result := tgfile.ProcessFile(filename)
	if result.IsErr() {
		mg.err = result.Err()
		return mg
	}

	mg.files.Push(result.Ok().File)

	media := gotgbot.InputMediaAudio{Media: result.Ok().Doc}

	if len(caption) > 0 {
		media.Caption = caption[0].Std()
	}

	mg.media.Push(media)

	return mg
}

// Document adds a document to the media group with optional caption.
func (mg *MediaGroup) Document(filename String, caption ...String) *MediaGroup {
	result := tgfile.ProcessFile(filename)
	if result.IsErr() {
		mg.err = result.Err()
		return mg
	}

	mg.files.Push(result.Ok().File)

	media := gotgbot.InputMediaDocument{Media: result.Ok().Doc}

	if len(caption) > 0 {
		media.Caption = caption[0].Std()
	}

	mg.media.Push(media)

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

// ReplyTo sets the message ID to reply to.
func (mg *MediaGroup) ReplyTo(messageID int64) *MediaGroup {
	mg.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return mg
}

// Business sets the business connection ID for the media group.
func (mg *MediaGroup) Business(id String) *MediaGroup {
	mg.opts.BusinessConnectionId = id.Std()
	return mg
}

// To sets the target chat ID for the media group.
func (mg *MediaGroup) To(chatID int64) *MediaGroup {
	mg.chatID = Some(chatID)
	return mg
}

// Send sends the media group to Telegram and returns the result.
func (mg *MediaGroup) Send() Result[Slice[gotgbot.Message]] {
	if mg.err != nil {
		return Err[Slice[gotgbot.Message]](mg.err)
	}

	if mg.media.Len() == 0 {
		return Err[Slice[gotgbot.Message]](errors.New("no media added to media group"))
	}

	defer mg.files.Iter().ForEach(func(file *File) { file.Close() })

	chatID := mg.chatID.UnwrapOr(mg.ctx.EffectiveChat.Id)

	return ResultOf[Slice[gotgbot.Message]](mg.ctx.Bot.Raw().SendMediaGroup(chatID, mg.media, mg.opts))
}
