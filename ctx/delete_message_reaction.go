package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// DeleteMessageReaction represents a request to remove a reaction from a message.
type DeleteMessageReaction struct {
	ctx       *Context
	messageID int64
	opts      *gotgbot.DeleteMessageReactionOpts
	chatID    g.Option[int64]
}

// ChatID sets the target chat ID where the message is located.
func (dmr *DeleteMessageReaction) ChatID(id int64) *DeleteMessageReaction {
	dmr.chatID = g.Some(id)
	return dmr
}

// UserID sets the identifier of the user whose reaction will be removed,
// if the reaction was added by a user.
func (dmr *DeleteMessageReaction) UserID(id int64) *DeleteMessageReaction {
	dmr.opts.UserId = id
	return dmr
}

// ActorChatID sets the identifier of the chat whose reaction will be removed,
// if the reaction was added by a chat.
func (dmr *DeleteMessageReaction) ActorChatID(id int64) *DeleteMessageReaction {
	dmr.opts.ActorChatId = id
	return dmr
}

// Timeout sets a custom timeout for this request.
func (dmr *DeleteMessageReaction) Timeout(duration time.Duration) *DeleteMessageReaction {
	if dmr.opts.RequestOpts == nil {
		dmr.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dmr.opts.RequestOpts.Timeout = duration

	return dmr
}

// APIURL sets a custom API URL for this request.
func (dmr *DeleteMessageReaction) APIURL(url g.String) *DeleteMessageReaction {
	if dmr.opts.RequestOpts == nil {
		dmr.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dmr.opts.RequestOpts.APIURL = url.Std()

	return dmr
}

// Send removes the reaction and returns the result.
func (dmr *DeleteMessageReaction) Send() g.Result[bool] {
	chatID := dmr.chatID.UnwrapOr(dmr.ctx.EffectiveChat.Id)
	return g.ResultOf(dmr.ctx.Bot.Raw().DeleteMessageReaction(chatID, dmr.messageID, dmr.opts))
}
