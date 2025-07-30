package business

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// RemovePhoto is a request builder for removing the business account profile photo.
type RemovePhoto struct {
	account *Account
	opts    *gotgbot.RemoveBusinessAccountProfilePhotoOpts
}

// Public removes the public profile photo if present.
func (rp *RemovePhoto) Public() *RemovePhoto {
	rp.opts.IsPublic = true
	return rp
}

// Timeout sets a custom timeout for this request.
func (rp *RemovePhoto) Timeout(duration time.Duration) *RemovePhoto {
	if rp.opts.RequestOpts == nil {
		rp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rp.opts.RequestOpts.Timeout = duration

	return rp
}

// APIURL sets a custom API URL for this request.
func (rp *RemovePhoto) APIURL(url g.String) *RemovePhoto {
	if rp.opts.RequestOpts == nil {
		rp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rp.opts.RequestOpts.APIURL = url.Std()

	return rp
}

// Send executes the RemovePhoto request.
func (rp *RemovePhoto) Send() g.Result[bool] {
	return g.ResultOf(rp.account.bot.Raw().RemoveBusinessAccountProfilePhoto(
		rp.account.connID.Std(),
		rp.opts,
	))
}
