package bot

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// SetMyProfilePhoto represents a request to set the bot's profile photo.
type SetMyProfilePhoto struct {
	bot   *Bot
	photo gotgbot.InputProfilePhoto
	opts  *gotgbot.SetMyProfilePhotoOpts
}

// Timeout sets a custom timeout for this request.
func (smpp *SetMyProfilePhoto) Timeout(duration time.Duration) *SetMyProfilePhoto {
	if smpp.opts.RequestOpts == nil {
		smpp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	smpp.opts.RequestOpts.Timeout = duration

	return smpp
}

// APIURL sets a custom API URL for this request.
func (smpp *SetMyProfilePhoto) APIURL(url g.String) *SetMyProfilePhoto {
	if smpp.opts.RequestOpts == nil {
		smpp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	smpp.opts.RequestOpts.APIURL = url.Std()

	return smpp
}

// Send sets the bot's profile photo and returns the result.
func (smpp *SetMyProfilePhoto) Send() g.Result[bool] {
	return g.ResultOf(smpp.bot.raw.SetMyProfilePhoto(smpp.photo, smpp.opts))
}
