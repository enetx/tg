package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/types/permissions"
)

type RestrictChatMember struct {
	ctx             *Context
	opts            *gotgbot.RestrictChatMemberOpts
	permissions     *gotgbot.ChatPermissions
	autoPermissions bool
	userID          int64
	chatID          Option[int64]
}

// ChatID sets the target chat ID for the restrict action.
func (r *RestrictChatMember) ChatID(id int64) *RestrictChatMember {
	r.chatID = Some(id)
	return r
}

// Until sets the restriction expiration time.
func (r *RestrictChatMember) Until(t time.Time) *RestrictChatMember {
	r.opts.UntilDate = t.Unix()
	return r
}

// For sets the restriction duration from now.
func (r *RestrictChatMember) For(d time.Duration) *RestrictChatMember {
	return r.Until(time.Now().Add(d))
}

// AutoPermissions uses chat default permissions instead of independent permissions.
func (r *RestrictChatMember) AutoPermissions() *RestrictChatMember {
	r.autoPermissions = true
	return r
}

// Permissions sets the allowed permissions for the restricted user.
func (r *RestrictChatMember) Permissions(perms ...permissions.Permission) *RestrictChatMember {
	r.permissions = permissions.Permissions(perms...)
	return r
}

// Timeout sets a custom timeout for this request.
func (r *RestrictChatMember) Timeout(duration time.Duration) *RestrictChatMember {
	if r.opts.RequestOpts == nil {
		r.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	r.opts.RequestOpts.Timeout = duration

	return r
}

// APIURL sets a custom API URL for this request.
func (r *RestrictChatMember) APIURL(url String) *RestrictChatMember {
	if r.opts.RequestOpts == nil {
		r.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	r.opts.RequestOpts.APIURL = url.Std()

	return r
}

// Send restricts the user's permissions and returns the result.
func (r *RestrictChatMember) Send() Result[bool] {
	if r.permissions == nil {
		return Err[bool](Errorf("permissions are required"))
	}

	chatID := r.chatID.UnwrapOr(r.ctx.EffectiveChat.Id)
	r.opts.UseIndependentChatPermissions = !r.autoPermissions

	return ResultOf(r.ctx.Bot.Raw().RestrictChatMember(chatID, r.userID, *r.permissions, r.opts))
}
