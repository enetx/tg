package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// GetChatMember represents a request to get information about a chat member.
type GetChatMember struct {
	ctx    *Context
	userID int64
	opts   *gotgbot.GetChatMemberOpts
	chatID Option[int64]
}

// ChatID sets the target chat ID for this request.
func (gcm *GetChatMember) ChatID(id int64) *GetChatMember {
	gcm.chatID = Some(id)
	return gcm
}

// UserID sets the target user ID for this request.
func (gcm *GetChatMember) UserID(id int64) *GetChatMember {
	gcm.userID = id
	return gcm
}

// Timeout sets a custom timeout for this request.
func (gcm *GetChatMember) Timeout(duration time.Duration) *GetChatMember {
	if gcm.opts.RequestOpts == nil {
		gcm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gcm.opts.RequestOpts.Timeout = duration

	return gcm
}

// APIURL sets a custom API URL for this request.
func (gcm *GetChatMember) APIURL(url String) *GetChatMember {
	if gcm.opts.RequestOpts == nil {
		gcm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gcm.opts.RequestOpts.APIURL = url.Std()

	return gcm
}

// Send executes the GetChatMember request and returns chat member information.
func (gcm *GetChatMember) Send() Result[gotgbot.ChatMember] {
	chatID := gcm.chatID.UnwrapOr(gcm.ctx.EffectiveChat.Id)
	return ResultOf(gcm.ctx.Bot.Raw().GetChatMember(chatID, gcm.userID, gcm.opts))
}
