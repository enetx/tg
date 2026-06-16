package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// SetChatMemberTag represents a request to set a tag for a regular member in a chat.
type SetChatMemberTag struct {
	ctx    *Context
	userID int64
	opts   *gotgbot.SetChatMemberTagOpts
	chatID g.Option[int64]
}

// ChatID sets the target chat ID where the member's tag will be updated.
func (scmt *SetChatMemberTag) ChatID(id int64) *SetChatMemberTag {
	scmt.chatID = g.Some(id)
	return scmt
}

// Tag sets the new tag for the member; 0-16 characters, emoji are not allowed.
// Pass an empty string to remove the tag.
func (scmt *SetChatMemberTag) Tag(tag g.String) *SetChatMemberTag {
	scmt.opts.Tag = tag.Std()
	return scmt
}

// Timeout sets a custom timeout for this request.
func (scmt *SetChatMemberTag) Timeout(duration time.Duration) *SetChatMemberTag {
	if scmt.opts.RequestOpts == nil {
		scmt.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	scmt.opts.RequestOpts.Timeout = duration

	return scmt
}

// APIURL sets a custom API URL for this request.
func (scmt *SetChatMemberTag) APIURL(url g.String) *SetChatMemberTag {
	if scmt.opts.RequestOpts == nil {
		scmt.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	scmt.opts.RequestOpts.APIURL = url.Std()

	return scmt
}

// Send executes the set chat member tag action and returns the result.
func (scmt *SetChatMemberTag) Send() g.Result[bool] {
	chatID := scmt.chatID.UnwrapOr(scmt.ctx.EffectiveChat.Id)
	return g.ResultOf(scmt.ctx.Bot.Raw().SetChatMemberTag(chatID, scmt.userID, scmt.opts))
}
