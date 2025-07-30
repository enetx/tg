package bot

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// SetMyShortDescription represents a request to set the bot's short description.
type SetMyShortDescription struct {
	bot  *Bot
	opts *gotgbot.SetMyShortDescriptionOpts
}

// Description sets the bot's short description text (0-120 characters).
func (smsd *SetMyShortDescription) Description(desc g.String) *SetMyShortDescription {
	smsd.opts.ShortDescription = desc.Std()
	return smsd
}

// Language sets the language code for the short description.
func (smsd *SetMyShortDescription) Language(code g.String) *SetMyShortDescription {
	smsd.opts.LanguageCode = code.Std()
	return smsd
}

// Remove removes the short description for the given language by setting empty string.
func (smsd *SetMyShortDescription) Remove() *SetMyShortDescription {
	smsd.opts.ShortDescription = ""
	return smsd
}

// Timeout sets a custom timeout for this request.
func (smsd *SetMyShortDescription) Timeout(duration time.Duration) *SetMyShortDescription {
	if smsd.opts.RequestOpts == nil {
		smsd.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	smsd.opts.RequestOpts.Timeout = duration

	return smsd
}

// APIURL sets a custom API URL for this request.
func (smsd *SetMyShortDescription) APIURL(url g.String) *SetMyShortDescription {
	if smsd.opts.RequestOpts == nil {
		smsd.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	smsd.opts.RequestOpts.APIURL = url.Std()

	return smsd
}

// Send sets the bot's short description and returns the result.
func (smsd *SetMyShortDescription) Send() g.Result[bool] {
	return g.ResultOf(smsd.bot.raw.SetMyShortDescription(smsd.opts))
}
