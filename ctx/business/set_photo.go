package business

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/input"
)

// SetPhoto is a request builder for setting the business account profile photo.
type SetPhoto struct {
	account *Account
	photo   g.String
	opts    *gotgbot.SetBusinessAccountProfilePhotoOpts
}

// Public marks the profile photo as publicly visible.
func (sp *SetPhoto) Public() *SetPhoto {
	sp.opts.IsPublic = true
	return sp
}

// Timeout sets a custom timeout for this request.
func (sp *SetPhoto) Timeout(duration time.Duration) *SetPhoto {
	if sp.opts.RequestOpts == nil {
		sp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sp.opts.RequestOpts.Timeout = duration

	return sp
}

// APIURL sets a custom API URL for this request.
func (sp *SetPhoto) APIURL(url g.String) *SetPhoto {
	if sp.opts.RequestOpts == nil {
		sp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sp.opts.RequestOpts.APIURL = url.Std()

	return sp
}

// Send executes the SetPhoto request.
func (sp *SetPhoto) Send() g.Result[bool] {
	return g.ResultOf(sp.account.bot.Raw().SetBusinessAccountProfilePhoto(
		sp.account.connID.Std(),
		input.StaticPhoto(sp.photo).Build(),
		sp.opts,
	))
}
