package business

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// SetUsername is a request builder for setting the account's username.
type SetUsername struct {
	account *Account
	opts    *gotgbot.SetBusinessAccountUsernameOpts
}

// Timeout sets a custom timeout for this request.
func (su *SetUsername) Timeout(duration time.Duration) *SetUsername {
	if su.opts.RequestOpts == nil {
		su.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	su.opts.RequestOpts.Timeout = duration

	return su
}

// APIURL sets a custom API URL for this request.
func (su *SetUsername) APIURL(url String) *SetUsername {
	if su.opts.RequestOpts == nil {
		su.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	su.opts.RequestOpts.APIURL = url.Std()

	return su
}

// Send executes the SetUsername request.
func (su *SetUsername) Send() Result[bool] {
	return ResultOf(su.account.bot.Raw().SetBusinessAccountUsername(
		su.account.connID.Std(),
		su.opts,
	))
}
