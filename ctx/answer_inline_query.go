package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/g/ref"
	"github.com/enetx/tg/inline"
)

// AnswerInlineQuery represents a request to answer an inline query.
type AnswerInlineQuery struct {
	ctx           *Context
	inlineQueryID g.String
	results       g.Slice[gotgbot.InlineQueryResult]
	opts          *gotgbot.AnswerInlineQueryOpts
}

// AddResult adds a single result builder to the inline query.
func (aiq *AnswerInlineQuery) AddResult(result inline.QueryResult) *AnswerInlineQuery {
	aiq.results.Push(result.Build())
	return aiq
}

// AddBuilders adds multiple result builders to the inline query.
func (aiq *AnswerInlineQuery) Results(results ...inline.QueryResult) *AnswerInlineQuery {
	for _, builder := range results {
		aiq.results.Push(builder.Build())
	}

	return aiq
}

// CacheFor sets the maximum amount of time the result may be cached on Telegram servers.
func (aiq *AnswerInlineQuery) CacheFor(duration time.Duration) *AnswerInlineQuery {
	aiq.opts.CacheTime = ref.Of(int64(duration.Seconds()))
	return aiq
}

// Personal sets that results may be cached on the server side only for the user that sent the query.
func (aiq *AnswerInlineQuery) Personal() *AnswerInlineQuery {
	aiq.opts.IsPersonal = true
	return aiq
}

// NextOffset sets the offset that a client should send in the next query.
func (aiq *AnswerInlineQuery) NextOffset(nextOffset g.String) *AnswerInlineQuery {
	aiq.opts.NextOffset = nextOffset.Std()
	return aiq
}

// ButtonText sets the text of the button that appears at the top of search results.
func (aiq *AnswerInlineQuery) ButtonText(text g.String) *AnswerInlineQuery {
	if aiq.opts.Button == nil {
		aiq.opts.Button = new(gotgbot.InlineQueryResultsButton)
	}

	aiq.opts.Button.Text = text.Std()

	return aiq
}

// StartParameter sets the parameter for the start message sent to the bot.
func (aiq *AnswerInlineQuery) StartParameter(parameter g.String) *AnswerInlineQuery {
	if aiq.opts.Button == nil {
		aiq.opts.Button = new(gotgbot.InlineQueryResultsButton)
	}

	aiq.opts.Button.StartParameter = parameter.Std()

	return aiq
}

// WebApp sets the Web App that will be launched when the user presses the button.
func (aiq *AnswerInlineQuery) WebApp(url g.String) *AnswerInlineQuery {
	if aiq.opts.Button == nil {
		aiq.opts.Button = new(gotgbot.InlineQueryResultsButton)
	}

	aiq.opts.Button.WebApp = &gotgbot.WebAppInfo{Url: url.Std()}

	return aiq
}

// Timeout sets a custom timeout for this request.
func (aiq *AnswerInlineQuery) Timeout(duration time.Duration) *AnswerInlineQuery {
	if aiq.opts.RequestOpts == nil {
		aiq.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	aiq.opts.RequestOpts.Timeout = duration

	return aiq
}

// APIURL sets a custom API URL for this request.
func (aiq *AnswerInlineQuery) APIURL(url g.String) *AnswerInlineQuery {
	if aiq.opts.RequestOpts == nil {
		aiq.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	aiq.opts.RequestOpts.APIURL = url.Std()

	return aiq
}

// Send answers the inline query and returns the result.
func (aiq *AnswerInlineQuery) Send() g.Result[bool] {
	return g.ResultOf(aiq.ctx.Bot.Raw().AnswerInlineQuery(aiq.inlineQueryID.Std(), aiq.results, aiq.opts))
}
