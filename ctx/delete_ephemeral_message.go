package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// DeleteEphemeralMessage represents a request to delete an ephemeral message.
type DeleteEphemeralMessage struct {
	ctx        *Context
	chatID     g.Option[int64]
	receiverID g.Option[int64]
	messageID  g.Option[int64]
	opts       *gotgbot.DeleteEphemeralMessageOpts
}

// ChatID sets the target chat ID for the delete action.
func (dem *DeleteEphemeralMessage) ChatID(id int64) *DeleteEphemeralMessage {
	dem.chatID = g.Some(id)
	return dem
}

// ReceiverUser sets the identifier of the user the ephemeral message was sent to.
func (dem *DeleteEphemeralMessage) ReceiverUser(id int64) *DeleteEphemeralMessage {
	dem.receiverID = g.Some(id)
	return dem
}

// MessageID sets the target ephemeral message ID to delete.
func (dem *DeleteEphemeralMessage) MessageID(id int64) *DeleteEphemeralMessage {
	dem.messageID = g.Some(id)
	return dem
}

// Timeout sets a custom timeout for this request.
func (dem *DeleteEphemeralMessage) Timeout(duration time.Duration) *DeleteEphemeralMessage {
	if dem.opts.RequestOpts == nil {
		dem.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dem.opts.RequestOpts.Timeout = duration

	return dem
}

// APIURL sets a custom API URL for this request.
func (dem *DeleteEphemeralMessage) APIURL(url g.String) *DeleteEphemeralMessage {
	if dem.opts.RequestOpts == nil {
		dem.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dem.opts.RequestOpts.APIURL = url.Std()

	return dem
}

// Send deletes the ephemeral message and returns the result.
func (dem *DeleteEphemeralMessage) Send() g.Result[bool] {
	return g.ResultOf(dem.ctx.Bot.Raw().DeleteEphemeralMessage(
		dem.chatID.UnwrapOr(dem.ctx.EffectiveChat.Id),
		dem.receiverID.UnwrapOr(dem.ctx.ephemeralReceiverID()),
		dem.messageID.UnwrapOr(dem.ctx.ephemeralMessageID()),
		dem.opts,
	))
}
