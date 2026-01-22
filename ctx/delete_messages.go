package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/core"
)

// DeleteMessages represents a request to delete multiple messages simultaneously.
type DeleteMessages struct {
	ctx        *Context
	chatID     g.Option[int64]
	messageIDs g.Slice[int64]
	after      g.Option[time.Duration]
	opts       *gotgbot.DeleteMessagesOpts
}

// ChatID sets the target chat ID for the delete action.
func (dm *DeleteMessages) ChatID(id int64) *DeleteMessages {
	dm.chatID = g.Some(id)
	return dm
}

// MessageIDs sets the message IDs to delete (up to 100 messages).
func (dm *DeleteMessages) MessageIDs(ids g.Slice[int64]) *DeleteMessages {
	dm.messageIDs = ids
	return dm
}

// AddMessages adds multiple message IDs to the delete list.
func (dm *DeleteMessages) AddMessages(ids ...int64) *DeleteMessages {
	dm.messageIDs.Push(ids...)
	return dm
}

// After schedules the messages deletion after the specified duration.
func (dm *DeleteMessages) After(duration time.Duration) *DeleteMessages {
	dm.after = g.Some(duration)
	return dm
}

// Timeout sets a custom timeout for this request.
func (dm *DeleteMessages) Timeout(duration time.Duration) *DeleteMessages {
	if dm.opts.RequestOpts == nil {
		dm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dm.opts.RequestOpts.Timeout = duration

	return dm
}

// APIURL sets a custom API URL for this request.
func (dm *DeleteMessages) APIURL(url g.String) *DeleteMessages {
	if dm.opts.RequestOpts == nil {
		dm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dm.opts.RequestOpts.APIURL = url.Std()

	return dm
}

// Send deletes the messages and returns the result.
func (dm *DeleteMessages) Send() g.Result[bool] {
	if dm.messageIDs.IsEmpty() {
		return g.Err[bool](g.Errorf("no message IDs specified for deletion"))
	}

	if dm.messageIDs.Len() > 100 {
		return g.Err[bool](g.Errorf("too many message IDs: {} (maximum 100)", dm.messageIDs.Len()))
	}

	chatID := dm.chatID.UnwrapOr(dm.ctx.EffectiveChat.Id)

	if dm.after.IsSome() {
		delay := dm.after.Some()
		dm.after = g.None[time.Duration]()

		bot := dm.ctx.Bot

		var opts *gotgbot.DeleteMessagesOpts
		if dm.opts != nil {
			ocp := *dm.opts
			opts = &ocp
		}

		go func(bot core.BotAPI, chatID int64, messageIDs []int64, opts *gotgbot.DeleteMessagesOpts, delay time.Duration) {
			<-time.After(delay)
			bot.Raw().DeleteMessages(chatID, messageIDs, opts)
		}(
			bot,
			chatID,
			dm.messageIDs,
			opts,
			delay,
		)

		return g.Ok(true)
	}

	return g.ResultOf(dm.ctx.Bot.Raw().DeleteMessages(chatID, dm.messageIDs, dm.opts))
}
