package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

type Forward struct {
	ctx         *Context
	fromChatID  int64
	messageID   int64
	opts        *gotgbot.ForwardMessageOpts
	toChatID    Option[int64]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
}

// After schedules the forward to be sent after the specified duration.
func (f *Forward) After(duration time.Duration) *Forward {
	f.after = Some(duration)
	return f
}

// DeleteAfter schedules the forwarded message to be deleted after the specified duration.
func (f *Forward) DeleteAfter(duration time.Duration) *Forward {
	f.deleteAfter = Some(duration)
	return f
}

// Silent disables notification for the forwarded message.
func (f *Forward) Silent() *Forward {
	f.opts.DisableNotification = true
	return f
}

// Protect enables content protection for the forwarded message.
func (f *Forward) Protect() *Forward {
	f.opts.ProtectContent = true
	return f
}

// Timeout sets the request timeout duration.
func (f *Forward) Timeout(duration time.Duration) *Forward {
	f.opts.RequestOpts = &gotgbot.RequestOpts{Timeout: duration}
	return f
}

// To sets the target chat ID for the forwarded message.
func (f *Forward) To(chatID int64) *Forward {
	f.toChatID = Some(chatID)
	return f
}

// Send forwards the message to the target chat and returns the result.
func (f *Forward) Send() Result[*gotgbot.Message] {
	return f.ctx.timers(f.after, f.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := f.toChatID.UnwrapOr(f.ctx.EffectiveChat.Id)
		return ResultOf(f.ctx.Bot.Raw().ForwardMessage(chatID, f.fromChatID, f.messageID, f.opts))
	})
}
