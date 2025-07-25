package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/types/permissions"
)

// SetChatPermissions represents a request to set default chat permissions.
type SetChatPermissions struct {
	ctx             *Context
	permissions     *gotgbot.ChatPermissions
	autoPermissions bool
	opts            *gotgbot.SetChatPermissionsOpts
	chatID          Option[int64]
}

// ChatID sets the target chat ID for this request.
func (scp *SetChatPermissions) ChatID(id int64) *SetChatPermissions {
	scp.chatID = Some(id)
	return scp
}

// AutoPermissions uses chat default permissions instead of independent permissions.
func (scp *SetChatPermissions) AutoPermissions() *SetChatPermissions {
	scp.autoPermissions = true
	return scp
}

// Timeout sets a custom timeout for this request.
func (scp *SetChatPermissions) Timeout(duration time.Duration) *SetChatPermissions {
	if scp.opts.RequestOpts == nil {
		scp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	scp.opts.RequestOpts.Timeout = duration

	return scp
}

// APIURL sets a custom API URL for this request.
func (scp *SetChatPermissions) APIURL(url String) *SetChatPermissions {
	if scp.opts.RequestOpts == nil {
		scp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	scp.opts.RequestOpts.APIURL = url.Std()

	return scp
}

// Permissions sets the allowed permissions for the chat.
func (scp *SetChatPermissions) Permissions(perms ...permissions.Permission) *SetChatPermissions {
	scp.permissions = permissions.Permissions(perms...)
	return scp
}

// Send executes the SetChatPermissions request.
func (scp *SetChatPermissions) Send() Result[bool] {
	if scp.permissions == nil {
		return Err[bool](Errorf("permissions are required"))
	}

	chatID := scp.chatID.UnwrapOr(scp.ctx.EffectiveChat.Id)
	scp.opts.UseIndependentChatPermissions = !scp.autoPermissions

	return ResultOf(scp.ctx.Bot.Raw().SetChatPermissions(chatID, *scp.permissions, scp.opts))
}
