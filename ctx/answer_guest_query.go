package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/inline"
)

// AnswerGuestQuery represents a request to reply to a received guest message.
type AnswerGuestQuery struct {
	ctx          *Context
	guestQueryID g.String
	result       inline.QueryResult
	opts         *gotgbot.AnswerGuestQueryOpts
}

// Timeout sets a custom timeout for this request.
func (agq *AnswerGuestQuery) Timeout(duration time.Duration) *AnswerGuestQuery {
	if agq.opts.RequestOpts == nil {
		agq.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	agq.opts.RequestOpts.Timeout = duration

	return agq
}

// APIURL sets a custom API URL for this request.
func (agq *AnswerGuestQuery) APIURL(url g.String) *AnswerGuestQuery {
	if agq.opts.RequestOpts == nil {
		agq.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	agq.opts.RequestOpts.APIURL = url.Std()

	return agq
}

// Send answers the guest query and returns the result.
func (agq *AnswerGuestQuery) Send() g.Result[*gotgbot.SentGuestMessage] {
	return g.ResultOf(agq.ctx.Bot.Raw().AnswerGuestQuery(
		agq.guestQueryID.Std(),
		agq.result.Build(),
		agq.opts,
	))
}
