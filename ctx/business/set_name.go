package business

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// SetName is a request builder for setting the account's name.
type SetName struct {
	account   *Account
	firstName String
	opts      *gotgbot.SetBusinessAccountNameOpts
}

// LastName sets the optional last name.
func (sn *SetName) LastName(lastName String) *SetName {
	sn.opts.LastName = lastName.Std()
	return sn
}

// Timeout sets a custom timeout for this request.
func (sn *SetName) Timeout(duration time.Duration) *SetName {
	if sn.opts.RequestOpts == nil {
		sn.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sn.opts.RequestOpts.Timeout = duration

	return sn
}

// APIURL sets a custom API URL for this request.
func (sn *SetName) APIURL(url String) *SetName {
	if sn.opts.RequestOpts == nil {
		sn.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sn.opts.RequestOpts.APIURL = url.Std()

	return sn
}

// Send executes the SetName request.
func (sn *SetName) Send() Result[bool] {
	return ResultOf(sn.account.bot.Raw().SetBusinessAccountName(
		sn.account.connID.Std(),
		sn.firstName.Std(),
		sn.opts,
	))
}
