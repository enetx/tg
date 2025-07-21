package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/types/roles"
)

type Promote struct {
	ctx    *Context
	opts   *gotgbot.PromoteChatMemberOpts
	roles  bool
	userID int64
	chatID Option[int64]
}

// ChatID sets the target chat ID for the promote action.
func (p *Promote) ChatID(id int64) *Promote {
	p.chatID = Some(id)
	return p
}

// Roles sets the administrator roles to grant to the user.
func (p *Promote) Roles(r ...roles.Role) *Promote {
	p.opts = roles.Roles(r...)
	p.roles = true

	return p
}

// Timeout sets the request timeout duration.
func (p *Promote) Timeout(duration time.Duration) *Promote {
	p.opts.RequestOpts = &gotgbot.RequestOpts{Timeout: duration}
	return p
}

// Send promotes the user to administrator and returns the result.
func (p *Promote) Send() Result[bool] {
	if !p.roles {
		return Err[bool](Errorf("roles are required"))
	}

	chatID := p.chatID.UnwrapOr(p.ctx.EffectiveChat.Id)
	return ResultOf(p.ctx.Bot.Raw().PromoteChatMember(chatID, p.userID, p.opts))
}
