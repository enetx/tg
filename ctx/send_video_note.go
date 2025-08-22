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

type SendVideoNote struct {
	ctx         *Context
	doc         gotgbot.InputFileOrString
	opts        *gotgbot.SendVideoNoteOpts
	file        *g.File
	thumb       *g.File
	chatID      g.Option[int64]
	after       g.Option[time.Duration]
	deleteAfter g.Option[time.Duration]
	err         error
}

// After schedules the video note to be sent after the specified duration.
func (svn *SendVideoNote) After(duration time.Duration) *SendVideoNote {
	svn.after = g.Some(duration)
	return svn
}

// DeleteAfter schedules the video note message to be deleted after the specified duration.
func (svn *SendVideoNote) DeleteAfter(duration time.Duration) *SendVideoNote {
	svn.deleteAfter = g.Some(duration)
	return svn
}

// Silent disables notification for the video note message.
func (svn *SendVideoNote) Silent() *SendVideoNote {
	svn.opts.DisableNotification = true
	return svn
}

// Protect enables content protection for the video note message.
func (svn *SendVideoNote) Protect() *SendVideoNote {
	svn.opts.ProtectContent = true
	return svn
}

// AllowPaidBroadcast allows the message to be sent in paid broadcast channels.
func (svn *SendVideoNote) AllowPaidBroadcast() *SendVideoNote {
	svn.opts.AllowPaidBroadcast = true
	return svn
}

// Effect sets a message effect for the message.
func (svn *SendVideoNote) Effect(effect effects.EffectType) *SendVideoNote {
	svn.opts.MessageEffectId = effect.String()
	return svn
}

// Markup sets the reply markup keyboard for the video note message.
func (svn *SendVideoNote) Markup(kb keyboard.Keyboard) *SendVideoNote {
	svn.opts.ReplyMarkup = kb.Markup()
	return svn
}

// Duration sets the video note duration in seconds.
func (svn *SendVideoNote) Duration(duration time.Duration) *SendVideoNote {
	svn.opts.Duration = int64(duration.Seconds())
	return svn
}

// Length sets the video note diameter (video notes are square).
func (svn *SendVideoNote) Length(length int64) *SendVideoNote {
	svn.opts.Length = length
	return svn
}

// Thumbnail sets a custom thumbnail for the video note.
func (svn *SendVideoNote) Thumbnail(file g.String) *SendVideoNote {
	svn.thumb = g.NewFile(file)

	reader := svn.thumb.Open()
	if reader.IsErr() {
		svn.err = reader.Err()
		return svn
	}

	svn.opts.Thumbnail = gotgbot.InputFileByReader(svn.thumb.Name().Std(), reader.Ok().Std())
	return svn
}

// Reply sets reply parameters using the reply builder.
func (svn *SendVideoNote) Reply(params *reply.Parameters) *SendVideoNote {
	svn.opts.ReplyParameters = params.Std()
	return svn
}

// Timeout sets a custom timeout for this request.
func (svn *SendVideoNote) Timeout(duration time.Duration) *SendVideoNote {
	if svn.opts.RequestOpts == nil {
		svn.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	svn.opts.RequestOpts.Timeout = duration

	return svn
}

// APIURL sets a custom API URL for this request.
func (svn *SendVideoNote) APIURL(url g.String) *SendVideoNote {
	if svn.opts.RequestOpts == nil {
		svn.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	svn.opts.RequestOpts.APIURL = url.Std()

	return svn
}

// Business sets the business connection ID for the video note message.
func (svn *SendVideoNote) Business(id g.String) *SendVideoNote {
	svn.opts.BusinessConnectionId = id.Std()
	return svn
}

// Thread sets the message thread ID for the video note message.
func (svn *SendVideoNote) Thread(id int64) *SendVideoNote {
	svn.opts.MessageThreadId = id
	return svn
}

// SuggestedPost sets suggested post parameters for direct messages chats.
func (svn *SendVideoNote) SuggestedPost(params *suggested.PostParameters) *SendVideoNote {
	if params != nil {
		svn.opts.SuggestedPostParameters = params.Std()
	}
	return svn
}

// To sets the target chat ID for the video note message.
func (svn *SendVideoNote) To(chatID int64) *SendVideoNote {
	svn.chatID = g.Some(chatID)
	return svn
}

// DirectMessagesTopic sets the direct messages topic ID for the message.
func (svn *SendVideoNote) DirectMessagesTopic(topicID int64) *SendVideoNote {
	svn.opts.DirectMessagesTopicId = topicID
	return svn
}

// Send sends the video note message to Telegram and returns the result.
func (svn *SendVideoNote) Send() g.Result[*gotgbot.Message] {
	if svn.err != nil {
		return g.Err[*gotgbot.Message](svn.err)
	}

	if svn.file != nil {
		defer svn.file.Close()
	}

	if svn.thumb != nil {
		defer svn.thumb.Close()
	}

	return svn.ctx.timers(svn.after, svn.deleteAfter, func() g.Result[*gotgbot.Message] {
		chatID := svn.chatID.UnwrapOr(svn.ctx.EffectiveChat.Id)
		return g.ResultOf(svn.ctx.Bot.Raw().SendVideoNote(chatID, svn.doc, svn.opts))
	})
}
