package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/core"
)

type DeleteMessage struct {
	ctx       *Context
	chatID    Option[int64]
	messageID Option[int64]
	after     Option[time.Duration]
	opts      *gotgbot.DeleteMessageOpts
}

// After schedules the message deletion after the specified duration.
func (dm *DeleteMessage) After(duration time.Duration) *DeleteMessage {
	dm.after = Some(duration)
	return dm
}

// ChatID sets the target chat ID for the delete action.
func (dm *DeleteMessage) ChatID(id int64) *DeleteMessage {
	dm.chatID = Some(id)
	return dm
}

// MessageID sets the target message ID to delete.
func (dm *DeleteMessage) MessageID(id int64) *DeleteMessage {
	dm.messageID = Some(id)
	return dm
}

// Timeout sets a custom timeout for this request.
func (dm *DeleteMessage) Timeout(duration time.Duration) *DeleteMessage {
	if dm.opts.RequestOpts == nil {
		dm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dm.opts.RequestOpts.Timeout = duration

	return dm
}

// APIURL sets a custom API URL for this request.
func (dm *DeleteMessage) APIURL(url String) *DeleteMessage {
	if dm.opts.RequestOpts == nil {
		dm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dm.opts.RequestOpts.APIURL = url.Std()

	return dm
}

// Send deletes the message and returns the result.
func (dm *DeleteMessage) Send() Result[bool] {
	chatID := dm.chatID.UnwrapOr(dm.ctx.EffectiveChat.Id)
	messageID := dm.messageID.UnwrapOr(dm.ctx.EffectiveMessage.MessageId)

	if dm.after.IsSome() {
		delay := dm.after.Some()
		dm.after = None[time.Duration]()

		bot := dm.ctx.Bot

		var opts *gotgbot.DeleteMessageOpts
		if dm.opts != nil {
			ocp := *dm.opts
			opts = &ocp
		}

		go func(bot core.BotAPI, chatID, messageID int64, opts *gotgbot.DeleteMessageOpts, delay time.Duration) {
			<-time.After(delay)
			bot.Raw().DeleteMessage(chatID, messageID, opts)
		}(bot, chatID, messageID, opts, delay)

		return Ok(true)
	}

	return ResultOf(dm.ctx.Bot.Raw().DeleteMessage(chatID, messageID, dm.opts))
}
