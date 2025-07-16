package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/types/permissions"
)

type Restrict struct {
	ctx             *Context
	opts            *gotgbot.RestrictChatMemberOpts
	permissions     *gotgbot.ChatPermissions
	autoPermissions bool
	userID          int64
	chatID          Option[int64]
}

func (r *Restrict) ChatID(id int64) *Restrict {
	r.chatID = Some(id)
	return r
}

func (r *Restrict) Until(t time.Time) *Restrict {
	r.opts.UntilDate = t.Unix()
	return r
}

func (r *Restrict) For(d time.Duration) *Restrict {
	return r.Until(time.Now().Add(d))
}

func (r *Restrict) AutoPermissions() *Restrict {
	r.autoPermissions = true
	return r
}

func (r *Restrict) Permissions(perms ...permissions.Permission) *Restrict {
	r.permissions = permissions.Permissions(perms...)
	return r
}

func (r *Restrict) Timeout(duration time.Duration) *Restrict {
	r.opts.RequestOpts = &gotgbot.RequestOpts{Timeout: duration}
	return r
}

func (r *Restrict) Send() Result[bool] {
	if r.permissions == nil {
		return Err[bool](Errorf("permissions are required"))
	}

	chatID := r.chatID.UnwrapOr(r.ctx.EffectiveChat.Id)
	r.opts.UseIndependentChatPermissions = !r.autoPermissions

	return ResultOf(r.ctx.Bot.Raw().RestrictChatMember(chatID, r.userID, *r.permissions, r.opts))
}
