package bot

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// SetMyName represents a request to set the bot's name.
type SetMyName struct {
	bot  *Bot
	opts *gotgbot.SetMyNameOpts
}

// Name sets the bot's name (0-64 characters).
func (smn *SetMyName) Name(name g.String) *SetMyName {
	smn.opts.Name = name.Std()
	return smn
}

// Language sets the language code for the name.
func (smn *SetMyName) Language(code g.String) *SetMyName {
	smn.opts.LanguageCode = code.Std()
	return smn
}

// Remove removes the name for the given language by setting empty string.
func (smn *SetMyName) Remove() *SetMyName {
	smn.opts.Name = ""
	return smn
}

// Timeout sets a custom timeout for this request.
func (smn *SetMyName) Timeout(duration time.Duration) *SetMyName {
	if smn.opts.RequestOpts == nil {
		smn.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	smn.opts.RequestOpts.Timeout = duration

	return smn
}

// APIURL sets a custom API URL for this request.
func (smn *SetMyName) APIURL(url g.String) *SetMyName {
	if smn.opts.RequestOpts == nil {
		smn.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	smn.opts.RequestOpts.APIURL = url.Std()

	return smn
}

// Send sets the bot's name and returns the result.
func (smn *SetMyName) Send() g.Result[bool] {
	return g.ResultOf(smn.bot.raw.SetMyName(smn.opts))
}
