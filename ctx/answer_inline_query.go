package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// AnswerInlineQuery represents a request to answer an inline query.
type AnswerInlineQuery struct {
	ctx           *Context
	inlineQueryID String
	results       Slice[gotgbot.InlineQueryResult]
	opts          *gotgbot.AnswerInlineQueryOpts
}

// Results sets the results for the inline query.
func (aiq *AnswerInlineQuery) Results(results Slice[gotgbot.InlineQueryResult]) *AnswerInlineQuery {
	aiq.results = results
	return aiq
}

// AddResult adds a single result to the inline query.
func (aiq *AnswerInlineQuery) AddResult(result gotgbot.InlineQueryResult) *AnswerInlineQuery {
	aiq.results = aiq.results.Append(result)
	return aiq
}

// CacheTime sets the maximum amount of time the result may be cached on Telegram servers.
func (aiq *AnswerInlineQuery) CacheTime(cacheTime int64) *AnswerInlineQuery {
	aiq.opts.CacheTime = cacheTime
	return aiq
}

// IsPersonal sets whether results may be cached on the server side only for the user that sent the query.
func (aiq *AnswerInlineQuery) Personal() *AnswerInlineQuery {
	aiq.opts.IsPersonal = true
	return aiq
}

// NextOffset sets the offset that a client should send in the next query.
func (aiq *AnswerInlineQuery) NextOffset(nextOffset String) *AnswerInlineQuery {
	aiq.opts.NextOffset = nextOffset.Std()
	return aiq
}

// SwitchPmText sets the text of the button that appears at the top of search results.
func (aiq *AnswerInlineQuery) SwitchPmText(text String) *AnswerInlineQuery {
	if aiq.opts.Button == nil {
		aiq.opts.Button = new(gotgbot.InlineQueryResultsButton)
	}

	aiq.opts.Button.Text = text.Std()

	return aiq
}

// SwitchPmParameter sets the parameter for the start message sent to the bot.
func (aiq *AnswerInlineQuery) SwitchPmParameter(parameter String) *AnswerInlineQuery {
	if aiq.opts.Button == nil {
		aiq.opts.Button = new(gotgbot.InlineQueryResultsButton)
	}

	aiq.opts.Button.StartParameter = parameter.Std()

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
func (aiq *AnswerInlineQuery) APIURL(url String) *AnswerInlineQuery {
	if aiq.opts.RequestOpts == nil {
		aiq.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	aiq.opts.RequestOpts.APIURL = url.Std()

	return aiq
}

// Send answers the inline query and returns the result.
func (aiq *AnswerInlineQuery) Send() Result[bool] {
	return ResultOf(aiq.ctx.Bot.Raw().AnswerInlineQuery(aiq.inlineQueryID.Std(), aiq.results, aiq.opts))
}
