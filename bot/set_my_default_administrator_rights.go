package bot

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/types/rights"
)

// SetMyDefaultAdministratorRights represents a request to set the bot's default administrator rights.
type SetMyDefaultAdministratorRights struct {
	bot  *Bot
	opts *gotgbot.SetMyDefaultAdministratorRightsOpts
}

// Rights sets the administrator rights.
func (smdar *SetMyDefaultAdministratorRights) Rights(r ...rights.Right) *SetMyDefaultAdministratorRights {
	smdar.opts.Rights = rights.Rights(r...)
	return smdar
}

// ForChannels sets whether these rights are for channels.
func (smdar *SetMyDefaultAdministratorRights) ForChannels() *SetMyDefaultAdministratorRights {
	smdar.opts.ForChannels = true
	return smdar
}

// Timeout sets a custom timeout for this request.
func (smdar *SetMyDefaultAdministratorRights) Timeout(duration time.Duration) *SetMyDefaultAdministratorRights {
	if smdar.opts.RequestOpts == nil {
		smdar.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	smdar.opts.RequestOpts.Timeout = duration

	return smdar
}

// APIURL sets a custom API URL for this request.
func (smdar *SetMyDefaultAdministratorRights) APIURL(url String) *SetMyDefaultAdministratorRights {
	if smdar.opts.RequestOpts == nil {
		smdar.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	smdar.opts.RequestOpts.APIURL = url.Std()

	return smdar
}

// Send sets the bot's default administrator rights.
func (smdar *SetMyDefaultAdministratorRights) Send() Result[bool] {
	return ResultOf(smdar.bot.Raw().SetMyDefaultAdministratorRights(smdar.opts))
}
