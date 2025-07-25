package bot

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// SetMyDescription represents a request to set the bot's description.
type SetMyDescription struct {
	bot  *Bot
	opts *gotgbot.SetMyDescriptionOpts
}

// Description sets the bot's description text (0-512 characters).
func (smd *SetMyDescription) Description(desc String) *SetMyDescription {
	smd.opts.Description = desc.Std()
	return smd
}

// Language sets the language code for the description.
func (smd *SetMyDescription) Language(code String) *SetMyDescription {
	smd.opts.LanguageCode = code.Std()
	return smd
}

// Remove removes the description for the given language by setting empty string.
func (smd *SetMyDescription) Remove() *SetMyDescription {
	smd.opts.Description = ""
	return smd
}

// Timeout sets a custom timeout for this request.
func (smd *SetMyDescription) Timeout(duration time.Duration) *SetMyDescription {
	if smd.opts.RequestOpts == nil {
		smd.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	smd.opts.RequestOpts.Timeout = duration

	return smd
}

// APIURL sets a custom API URL for this request.
func (smd *SetMyDescription) APIURL(url String) *SetMyDescription {
	if smd.opts.RequestOpts == nil {
		smd.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	smd.opts.RequestOpts.APIURL = url.Std()

	return smd
}

// Send sets the bot's description and returns the result.
func (smd *SetMyDescription) Send() Result[bool] {
	return ResultOf(smd.bot.raw.SetMyDescription(smd.opts))
}
