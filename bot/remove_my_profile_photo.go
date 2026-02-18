package bot

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// RemoveMyProfilePhoto represents a request to remove the bot's profile photo.
type RemoveMyProfilePhoto struct {
	bot  *Bot
	opts *gotgbot.RemoveMyProfilePhotoOpts
}

// Timeout sets a custom timeout for this request.
func (rmpp *RemoveMyProfilePhoto) Timeout(duration time.Duration) *RemoveMyProfilePhoto {
	if rmpp.opts.RequestOpts == nil {
		rmpp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rmpp.opts.RequestOpts.Timeout = duration

	return rmpp
}

// APIURL sets a custom API URL for this request.
func (rmpp *RemoveMyProfilePhoto) APIURL(url g.String) *RemoveMyProfilePhoto {
	if rmpp.opts.RequestOpts == nil {
		rmpp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rmpp.opts.RequestOpts.APIURL = url.Std()

	return rmpp
}

// Send removes the bot's profile photo and returns the result.
func (rmpp *RemoveMyProfilePhoto) Send() g.Result[bool] {
	return g.ResultOf(rmpp.bot.raw.RemoveMyProfilePhoto(rmpp.opts))
}
