package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// DeleteAllMessageReactions represents a request to remove up to 10000 recent
// reactions in a group or supergroup added by a given user or chat.
// The bot must have the 'can_delete_messages' administrator right in the chat.
type DeleteAllMessageReactions struct {
	ctx    *Context
	opts   *gotgbot.DeleteAllMessageReactionsOpts
	chatID g.Option[int64]
}

// ChatID sets the target chat ID where reactions will be removed.
func (damr *DeleteAllMessageReactions) ChatID(id int64) *DeleteAllMessageReactions {
	damr.chatID = g.Some(id)
	return damr
}

// UserID sets the identifier of the user whose reactions will be removed,
// if the reactions were added by a user.
func (damr *DeleteAllMessageReactions) UserID(id int64) *DeleteAllMessageReactions {
	damr.opts.UserId = id
	return damr
}

// ActorChatID sets the identifier of the chat whose reactions will be removed,
// if the reactions were added by a chat.
func (damr *DeleteAllMessageReactions) ActorChatID(id int64) *DeleteAllMessageReactions {
	damr.opts.ActorChatId = id
	return damr
}

// Timeout sets a custom timeout for this request.
func (damr *DeleteAllMessageReactions) Timeout(duration time.Duration) *DeleteAllMessageReactions {
	if damr.opts.RequestOpts == nil {
		damr.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	damr.opts.RequestOpts.Timeout = duration

	return damr
}

// APIURL sets a custom API URL for this request.
func (damr *DeleteAllMessageReactions) APIURL(url g.String) *DeleteAllMessageReactions {
	if damr.opts.RequestOpts == nil {
		damr.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	damr.opts.RequestOpts.APIURL = url.Std()

	return damr
}

// Send removes the reactions and returns the result.
func (damr *DeleteAllMessageReactions) Send() g.Result[bool] {
	chatID := damr.chatID.UnwrapOr(damr.ctx.EffectiveChat.Id)
	return g.ResultOf(damr.ctx.Bot.Raw().DeleteAllMessageReactions(chatID, damr.opts))
}
