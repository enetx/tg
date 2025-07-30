package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/inline"
)

// AnswerWebAppQuery represents a request to answer a web app query.
type AnswerWebAppQuery struct {
	ctx           *Context
	webAppQueryID g.String
	result        inline.QueryResult
	opts          *gotgbot.AnswerWebAppQueryOpts
	err           error
}

// Timeout sets a custom timeout for this request.
func (awaq *AnswerWebAppQuery) Timeout(duration time.Duration) *AnswerWebAppQuery {
	if awaq.opts.RequestOpts == nil {
		awaq.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	awaq.opts.RequestOpts.Timeout = duration

	return awaq
}

// APIURL sets a custom API URL for this request.
func (awaq *AnswerWebAppQuery) APIURL(url g.String) *AnswerWebAppQuery {
	if awaq.opts.RequestOpts == nil {
		awaq.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	awaq.opts.RequestOpts.APIURL = url.Std()

	return awaq
}

// Send answers the web app query and returns the result.
func (awaq *AnswerWebAppQuery) Send() g.Result[*gotgbot.SentWebAppMessage] {
	return g.ResultOf(awaq.ctx.Bot.Raw().AnswerWebAppQuery(
		awaq.webAppQueryID.Std(),
		awaq.result.Build(),
		awaq.opts,
	))
}
