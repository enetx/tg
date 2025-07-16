package ctx

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

type PreCheckout struct {
	ctx  *Context
	ok   bool
	err  String
	opts *gotgbot.AnswerPreCheckoutQueryOpts
}

func (p *PreCheckout) Ok() *PreCheckout {
	p.ok = true
	return p
}

func (p *PreCheckout) Error(text String) *PreCheckout {
	p.ok = false
	p.err = text
	return p
}

func (p *PreCheckout) Send() Result[bool] {
	query := p.ctx.Update.PreCheckoutQuery
	if query == nil {
		return Err[bool](Errorf("no precheckout query"))
	}

	if !p.ok {
		p.opts = &gotgbot.AnswerPreCheckoutQueryOpts{ErrorMessage: p.err.Std()}
	}

	return ResultOf(query.Answer(p.ctx.Bot.Raw(), p.ok, p.opts))
}
