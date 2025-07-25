package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/types/roles"
)

type PromoteChatMember struct {
	ctx    *Context
	opts   *gotgbot.PromoteChatMemberOpts
	roles  bool
	userID int64
	chatID Option[int64]
}

// ChatID sets the target chat ID for the promote action.
func (p *PromoteChatMember) ChatID(id int64) *PromoteChatMember {
	p.chatID = Some(id)
	return p
}

// Roles sets the administrator roles to grant to the user.
func (p *PromoteChatMember) Roles(r ...roles.Role) *PromoteChatMember {
	p.opts = roles.Roles(r...)
	p.roles = true

	return p
}

// Timeout sets a custom timeout for this request.
func (p *PromoteChatMember) Timeout(duration time.Duration) *PromoteChatMember {
	if p.opts.RequestOpts == nil {
		p.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	p.opts.RequestOpts.Timeout = duration

	return p
}

// APIURL sets a custom API URL for this request.
func (p *PromoteChatMember) APIURL(url String) *PromoteChatMember {
	if p.opts.RequestOpts == nil {
		p.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	p.opts.RequestOpts.APIURL = url.Std()

	return p
}

// Send promotes the user to administrator and returns the result.
func (p *PromoteChatMember) Send() Result[bool] {
	if !p.roles {
		return Err[bool](Errorf("roles are required"))
	}

	chatID := p.chatID.UnwrapOr(p.ctx.EffectiveChat.Id)
	return ResultOf(p.ctx.Bot.Raw().PromoteChatMember(chatID, p.userID, p.opts))
}
