package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// SetManagedBotAccessSettings represents a request to change the access settings of a managed bot.
type SetManagedBotAccessSettings struct {
	ctx                *Context
	userID             int64
	isAccessRestricted bool
	opts               *gotgbot.SetManagedBotAccessSettingsOpts
}

// AddedUserIDs sets the list of up to 10 identifiers of users who will have access to the bot
// in addition to its owner. Ignored if access is not restricted.
func (smbas *SetManagedBotAccessSettings) AddedUserIDs(ids ...int64) *SetManagedBotAccessSettings {
	smbas.opts.AddedUserIds = ids
	return smbas
}

// Timeout sets a custom timeout for this request.
func (smbas *SetManagedBotAccessSettings) Timeout(duration time.Duration) *SetManagedBotAccessSettings {
	if smbas.opts.RequestOpts == nil {
		smbas.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	smbas.opts.RequestOpts.Timeout = duration

	return smbas
}

// APIURL sets a custom API URL for this request.
func (smbas *SetManagedBotAccessSettings) APIURL(url g.String) *SetManagedBotAccessSettings {
	if smbas.opts.RequestOpts == nil {
		smbas.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	smbas.opts.RequestOpts.APIURL = url.Std()

	return smbas
}

// Send updates the managed bot access settings and returns the result.
func (smbas *SetManagedBotAccessSettings) Send() g.Result[bool] {
	return g.ResultOf(smbas.ctx.Bot.Raw().SetManagedBotAccessSettings(
		smbas.userID,
		smbas.isAccessRestricted,
		smbas.opts,
	))
}
