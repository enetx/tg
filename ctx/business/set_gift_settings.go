package business

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// SetGiftSettings is a request builder for changing gift privacy settings in a business account.
// Use fluent methods to configure options before calling Send().
type SetGiftSettings struct {
	account           *Account
	showGiftButton    bool
	acceptedGiftTypes gotgbot.AcceptedGiftTypes
	opts              *gotgbot.SetBusinessAccountGiftSettingsOpts
}

// ShowGiftButton sets whether a button for sending a gift should always be shown in the input field.
func (sgs *SetGiftSettings) ShowGiftButton() *SetGiftSettings {
	sgs.showGiftButton = true
	return sgs
}

// AcceptUnlimitedGifts sets whether unlimited regular gifts are accepted.
func (sgs *SetGiftSettings) AcceptUnlimitedGifts() *SetGiftSettings {
	sgs.acceptedGiftTypes.UnlimitedGifts = true
	return sgs
}

// AcceptLimitedGifts sets whether limited regular gifts are accepted.
func (sgs *SetGiftSettings) AcceptLimitedGifts() *SetGiftSettings {
	sgs.acceptedGiftTypes.LimitedGifts = true
	return sgs
}

// AcceptUniqueGifts sets whether unique gifts or gifts that can be upgraded to unique are accepted.
func (sgs *SetGiftSettings) AcceptUniqueGifts() *SetGiftSettings {
	sgs.acceptedGiftTypes.UniqueGifts = true
	return sgs
}

// AcceptPremiumSubscription sets whether Telegram Premium subscriptions are accepted.
func (sgs *SetGiftSettings) AcceptPremiumSubscription() *SetGiftSettings {
	sgs.acceptedGiftTypes.PremiumSubscription = true
	return sgs
}

// AcceptGiftsFromChannels sets whether transfers of unique gifts from channels are accepted.
func (sgs *SetGiftSettings) AcceptGiftsFromChannels() *SetGiftSettings {
	sgs.acceptedGiftTypes.GiftsFromChannels = true
	return sgs
}

// AcceptAllGifts is a convenience method to accept all gift types.
func (sgs *SetGiftSettings) AcceptAllGifts() *SetGiftSettings {
	sgs.acceptedGiftTypes.UnlimitedGifts = true
	sgs.acceptedGiftTypes.LimitedGifts = true
	sgs.acceptedGiftTypes.UniqueGifts = true
	sgs.acceptedGiftTypes.PremiumSubscription = true
	sgs.acceptedGiftTypes.GiftsFromChannels = true
	return sgs
}

// AcceptNoGifts is a convenience method to reject all gift types.
func (sgs *SetGiftSettings) AcceptNoGifts() *SetGiftSettings {
	sgs.acceptedGiftTypes.UnlimitedGifts = false
	sgs.acceptedGiftTypes.LimitedGifts = false
	sgs.acceptedGiftTypes.UniqueGifts = false
	sgs.acceptedGiftTypes.PremiumSubscription = false
	sgs.acceptedGiftTypes.GiftsFromChannels = false
	return sgs
}

// Timeout sets a custom timeout for this request.
func (sgs *SetGiftSettings) Timeout(duration time.Duration) *SetGiftSettings {
	if sgs.opts.RequestOpts == nil {
		sgs.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sgs.opts.RequestOpts.Timeout = duration
	return sgs
}

// APIURL sets a custom Telegram Bot API URL for this request.
func (sgs *SetGiftSettings) APIURL(url g.String) *SetGiftSettings {
	if sgs.opts.RequestOpts == nil {
		sgs.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sgs.opts.RequestOpts.APIURL = url.Std()
	return sgs
}

// Send executes the request to change gift settings.
// Returns true on success wrapped in g.Result.
func (sgs *SetGiftSettings) Send() g.Result[bool] {
	return g.ResultOf(sgs.account.bot.Raw().SetBusinessAccountGiftSettings(
		sgs.account.connID.Std(),
		sgs.showGiftButton,
		sgs.acceptedGiftTypes,
		sgs.opts,
	))
}
