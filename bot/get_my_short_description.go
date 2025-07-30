package bot

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// GetMyShortDescription represents a request to get the bot's short description.
type GetMyShortDescription struct {
	bot  *Bot
	opts *gotgbot.GetMyShortDescriptionOpts
}

// Language sets the language code for getting the short description.
func (gmsd *GetMyShortDescription) Language(code g.String) *GetMyShortDescription {
	gmsd.opts.LanguageCode = code.Std()
	return gmsd
}

// Timeout sets a custom timeout for this request.
func (gmsd *GetMyShortDescription) Timeout(duration time.Duration) *GetMyShortDescription {
	if gmsd.opts.RequestOpts == nil {
		gmsd.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gmsd.opts.RequestOpts.Timeout = duration

	return gmsd
}

// APIURL sets a custom API URL for this request.
func (gmsd *GetMyShortDescription) APIURL(url g.String) *GetMyShortDescription {
	if gmsd.opts.RequestOpts == nil {
		gmsd.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gmsd.opts.RequestOpts.APIURL = url.Std()

	return gmsd
}

// Send gets the bot's short description and returns the result.
func (gmsd *GetMyShortDescription) Send() g.Result[*gotgbot.BotShortDescription] {
	return g.ResultOf(gmsd.bot.raw.GetMyShortDescription(gmsd.opts))
}
