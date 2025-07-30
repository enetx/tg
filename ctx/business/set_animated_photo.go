package business

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/input"
)

// SetAnimatedPhoto is a request builder for setting the business account animated profile photo.
type SetAnimatedPhoto struct {
	account            *Account
	animation          g.String
	mainFrameTimestamp g.Option[float64]
	opts               *gotgbot.SetBusinessAccountProfilePhotoOpts
}

// MainFrame sets the timestamp in seconds of the frame that will be used as the static profile photo.
func (sap *SetAnimatedPhoto) MainFrame(timestamp time.Duration) *SetAnimatedPhoto {
	sap.mainFrameTimestamp = g.Some(timestamp.Seconds())
	return sap
}

// Public marks the profile photo as publicly visible.
func (sap *SetAnimatedPhoto) Public() *SetAnimatedPhoto {
	sap.opts.IsPublic = true
	return sap
}

// Timeout sets a custom timeout for this request.
func (sap *SetAnimatedPhoto) Timeout(duration time.Duration) *SetAnimatedPhoto {
	if sap.opts.RequestOpts == nil {
		sap.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sap.opts.RequestOpts.Timeout = duration

	return sap
}

// APIURL sets a custom API URL for this request.
func (sap *SetAnimatedPhoto) APIURL(url g.String) *SetAnimatedPhoto {
	if sap.opts.RequestOpts == nil {
		sap.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sap.opts.RequestOpts.APIURL = url.Std()

	return sap
}

// Send executes the SetAnimatedPhoto request.
func (sap *SetAnimatedPhoto) Send() g.Result[bool] {
	animated := input.AnimatedPhoto(sap.animation)

	if sap.mainFrameTimestamp.IsSome() {
		animated.MainFrameTimestamp(sap.mainFrameTimestamp.Some())
	}

	return g.ResultOf(sap.account.bot.Raw().SetBusinessAccountProfilePhoto(
		sap.account.connID.Std(),
		animated.Build(),
		sap.opts,
	))
}
