package bot

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// GetMyDescription represents a request to get the bot's description.
type GetMyDescription struct {
	bot  *Bot
	opts *gotgbot.GetMyDescriptionOpts
}

// Language sets the language code for getting the description.
func (gmd *GetMyDescription) Language(code String) *GetMyDescription {
	gmd.opts.LanguageCode = code.Std()
	return gmd
}

// Timeout sets a custom timeout for this request.
func (gmd *GetMyDescription) Timeout(duration time.Duration) *GetMyDescription {
	if gmd.opts.RequestOpts == nil {
		gmd.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gmd.opts.RequestOpts.Timeout = duration

	return gmd
}

// APIURL sets a custom API URL for this request.
func (gmd *GetMyDescription) APIURL(url String) *GetMyDescription {
	if gmd.opts.RequestOpts == nil {
		gmd.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gmd.opts.RequestOpts.APIURL = url.Std()

	return gmd
}

// Send gets the bot's description and returns the result.
func (gmd *GetMyDescription) Send() Result[*gotgbot.BotDescription] {
	return ResultOf(gmd.bot.raw.GetMyDescription(gmd.opts))
}
