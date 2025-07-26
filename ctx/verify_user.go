package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// VerifyUser represents a request to verify a user.
type VerifyUser struct {
	ctx    *Context
	userID int64
	opts   *gotgbot.VerifyUserOpts
}

// CustomDescription for the verification; 0-70 characters.
// Must be empty if the organization isn't allowed to provide a custom verification description.
func (vu *VerifyUser) CustomDescription(description String) *VerifyUser {
	vu.opts.CustomDescription = description.Std()
	return vu
}

// Timeout sets a custom timeout for this request.
func (vu *VerifyUser) Timeout(duration time.Duration) *VerifyUser {
	if vu.opts.RequestOpts == nil {
		vu.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	vu.opts.RequestOpts.Timeout = duration

	return vu
}

// APIURL sets a custom API URL for this request.
func (vu *VerifyUser) APIURL(url String) *VerifyUser {
	if vu.opts.RequestOpts == nil {
		vu.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	vu.opts.RequestOpts.APIURL = url.Std()

	return vu
}

// Send verifies the user.
func (vu *VerifyUser) Send() Result[bool] {
	return ResultOf(vu.ctx.Bot.Raw().VerifyUser(vu.userID, vu.opts))
}
