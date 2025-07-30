package bot

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// GetMyName represents a request to get the bot's name.
type GetMyName struct {
	bot  *Bot
	opts *gotgbot.GetMyNameOpts
}

// Language sets the language code for getting the name.
func (gmn *GetMyName) Language(code g.String) *GetMyName {
	gmn.opts.LanguageCode = code.Std()
	return gmn
}

// Timeout sets a custom timeout for this request.
func (gmn *GetMyName) Timeout(duration time.Duration) *GetMyName {
	if gmn.opts.RequestOpts == nil {
		gmn.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gmn.opts.RequestOpts.Timeout = duration

	return gmn
}

// APIURL sets a custom API URL for this request.
func (gmn *GetMyName) APIURL(url g.String) *GetMyName {
	if gmn.opts.RequestOpts == nil {
		gmn.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gmn.opts.RequestOpts.APIURL = url.Std()

	return gmn
}

// Send gets the bot's name and returns the result.
func (gmn *GetMyName) Send() g.Result[*gotgbot.BotName] {
	return g.ResultOf(gmn.bot.raw.GetMyName(gmn.opts))
}
