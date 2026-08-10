package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// AnswerChatJoinRequestQuery represents a request to process a received chat join request query.
type AnswerChatJoinRequestQuery struct {
	ctx     *Context
	queryID g.Option[g.String]
	result  g.String
	opts    *gotgbot.AnswerChatJoinRequestQueryOpts
}

// QueryID sets the join request query ID to answer.
// By default the query ID of the effective update is used.
func (acjrq *AnswerChatJoinRequestQuery) QueryID(id g.String) *AnswerChatJoinRequestQuery {
	acjrq.queryID = g.Some(id)
	return acjrq
}

// Approve allows the user to join the chat.
func (acjrq *AnswerChatJoinRequestQuery) Approve() *AnswerChatJoinRequestQuery {
	acjrq.result = "approve"
	return acjrq
}

// Decline disallows the user to join the chat.
func (acjrq *AnswerChatJoinRequestQuery) Decline() *AnswerChatJoinRequestQuery {
	acjrq.result = "decline"
	return acjrq
}

// Queue leaves the decision to other administrators.
func (acjrq *AnswerChatJoinRequestQuery) Queue() *AnswerChatJoinRequestQuery {
	acjrq.result = "queue"
	return acjrq
}

// Timeout sets a custom timeout for this request.
func (acjrq *AnswerChatJoinRequestQuery) Timeout(duration time.Duration) *AnswerChatJoinRequestQuery {
	if acjrq.opts.RequestOpts == nil {
		acjrq.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	acjrq.opts.RequestOpts.Timeout = duration

	return acjrq
}

// APIURL sets a custom API URL for this request.
func (acjrq *AnswerChatJoinRequestQuery) APIURL(url g.String) *AnswerChatJoinRequestQuery {
	if acjrq.opts.RequestOpts == nil {
		acjrq.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	acjrq.opts.RequestOpts.APIURL = url.Std()

	return acjrq
}

// Send answers the chat join request query and returns the result.
func (acjrq *AnswerChatJoinRequestQuery) Send() g.Result[bool] {
	return g.ResultOf(acjrq.ctx.Bot.Raw().AnswerChatJoinRequestQuery(
		acjrq.queryID.UnwrapOr(acjrq.ctx.chatJoinRequestQueryID()).Std(),
		acjrq.result.Std(),
		acjrq.opts,
	))
}
