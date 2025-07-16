package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/core"
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
		delay := d.after.Some()
		d.after = None[time.Duration]()

		bot := d.ctx.Bot

		var opts *gotgbot.DeleteMessageOpts
		if d.opts != nil {
			ocp := *d.opts
			opts = &ocp
		}

		go func(bot core.BotAPI, chatID, messageID int64, opts *gotgbot.DeleteMessageOpts, delay time.Duration) {
			<-time.After(delay)
			bot.Raw().DeleteMessage(chatID, messageID, opts)
		}(bot, chatID, messageID, opts, delay)

		return Ok(true)
	}

	return ResultOf(d.ctx.Bot.Raw().DeleteMessage(chatID, messageID, d.opts))
}
