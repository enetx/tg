package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// SendChatJoinRequestWebApp represents a request to show a Mini App to a user
// before deciding the outcome of a chat join request query.
type SendChatJoinRequestWebApp struct {
	ctx       *Context
	queryID   g.Option[g.String]
	webAppURL g.String
	opts      *gotgbot.SendChatJoinRequestWebAppOpts
}

// QueryID sets the join request query ID to process.
// By default the query ID of the effective update is used.
func (scjrwa *SendChatJoinRequestWebApp) QueryID(id g.String) *SendChatJoinRequestWebApp {
	scjrwa.queryID = g.Some(id)
	return scjrwa
}

// URL sets the HTTPS URL of the Web App to be opened.
func (scjrwa *SendChatJoinRequestWebApp) URL(url g.String) *SendChatJoinRequestWebApp {
	scjrwa.webAppURL = url
	return scjrwa
}

// Timeout sets a custom timeout for this request.
func (scjrwa *SendChatJoinRequestWebApp) Timeout(duration time.Duration) *SendChatJoinRequestWebApp {
	if scjrwa.opts.RequestOpts == nil {
		scjrwa.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	scjrwa.opts.RequestOpts.Timeout = duration

	return scjrwa
}

// APIURL sets a custom API URL for this request.
func (scjrwa *SendChatJoinRequestWebApp) APIURL(url g.String) *SendChatJoinRequestWebApp {
	if scjrwa.opts.RequestOpts == nil {
		scjrwa.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	scjrwa.opts.RequestOpts.APIURL = url.Std()

	return scjrwa
}

// Send shows the Mini App to the user and returns the result.
func (scjrwa *SendChatJoinRequestWebApp) Send() g.Result[bool] {
	return g.ResultOf(scjrwa.ctx.Bot.Raw().SendChatJoinRequestWebApp(
		scjrwa.queryID.UnwrapOr(scjrwa.ctx.chatJoinRequestQueryID()).Std(),
		scjrwa.webAppURL.Std(),
		scjrwa.opts,
	))
}
