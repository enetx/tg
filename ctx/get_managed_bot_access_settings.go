package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// GetManagedBotAccessSettings represents a request to get the access settings of a managed bot.
type GetManagedBotAccessSettings struct {
	ctx    *Context
	userID int64
	opts   *gotgbot.GetManagedBotAccessSettingsOpts
}

// Timeout sets a custom timeout for this request.
func (gmbas *GetManagedBotAccessSettings) Timeout(duration time.Duration) *GetManagedBotAccessSettings {
	if gmbas.opts.RequestOpts == nil {
		gmbas.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gmbas.opts.RequestOpts.Timeout = duration

	return gmbas
}

// APIURL sets a custom API URL for this request.
func (gmbas *GetManagedBotAccessSettings) APIURL(url g.String) *GetManagedBotAccessSettings {
	if gmbas.opts.RequestOpts == nil {
		gmbas.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gmbas.opts.RequestOpts.APIURL = url.Std()

	return gmbas
}

// Send retrieves the managed bot access settings and returns the result.
func (gmbas *GetManagedBotAccessSettings) Send() g.Result[*gotgbot.BotAccessSettings] {
	return g.ResultOf(gmbas.ctx.Bot.Raw().GetManagedBotAccessSettings(gmbas.userID, gmbas.opts))
}
