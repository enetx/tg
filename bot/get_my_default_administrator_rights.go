package bot

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// GetMyDefaultAdministratorRights represents a request to get the bot's default administrator rights.
type GetMyDefaultAdministratorRights struct {
	bot  *Bot
	opts *gotgbot.GetMyDefaultAdministratorRightsOpts
}

// ForChannels sets whether to get rights for channels.
func (gmdar *GetMyDefaultAdministratorRights) ForChannels() *GetMyDefaultAdministratorRights {
	gmdar.opts.ForChannels = true
	return gmdar
}

// Timeout sets a custom timeout for this request.
func (gmdar *GetMyDefaultAdministratorRights) Timeout(duration time.Duration) *GetMyDefaultAdministratorRights {
	if gmdar.opts.RequestOpts == nil {
		gmdar.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gmdar.opts.RequestOpts.Timeout = duration

	return gmdar
}

// APIURL sets a custom API URL for this request.
func (gmdar *GetMyDefaultAdministratorRights) APIURL(url g.String) *GetMyDefaultAdministratorRights {
	if gmdar.opts.RequestOpts == nil {
		gmdar.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gmdar.opts.RequestOpts.APIURL = url.Std()

	return gmdar
}

// Send gets the bot's default administrator rights.
func (gmdar *GetMyDefaultAdministratorRights) Send() g.Result[*gotgbot.ChatAdministratorRights] {
	return g.ResultOf(gmdar.bot.Raw().GetMyDefaultAdministratorRights(gmdar.opts))
}
