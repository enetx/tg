package tg

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

type Delete struct {
	ctx       *Context
	chatID    Option[int64]
	messageID Option[int64]
	after     Option[time.Duration]
	opts      *gotgbot.DeleteMessageOpts
}

func (d *Delete) After(duration time.Duration) *Delete {
	d.after = Some(duration)
	return d
}

func (d *Delete) ChatID(id int64) *Delete {
	d.chatID = Some(id)
	return d
}

func (d *Delete) MessageID(id int64) *Delete {
	d.messageID = Some(id)
	return d
}

func (d *Delete) Timeout(duration time.Duration) *Delete {
	d.opts.RequestOpts = &gotgbot.RequestOpts{Timeout: duration}
	return d
}

func (d *Delete) Send() Result[bool] {
	chatID := d.chatID.UnwrapOr(d.ctx.EffectiveChat.Id)
	messageID := d.messageID.UnwrapOr(d.ctx.EffectiveMessage.MessageId)

	if d.after.IsSome() {
		go func(ctx *Context, chatID, messageID int64, opts *gotgbot.DeleteMessageOpts, delay time.Duration) {
			<-time.After(delay)
			ctx.Bot.Raw.DeleteMessage(chatID, messageID, opts)
		}(d.ctx.Copy(), chatID, messageID, d.opts, d.after.Some())

		return Ok(true)
	}

	return ResultOf(d.ctx.Bot.Raw.DeleteMessage(chatID, messageID, d.opts))
}
