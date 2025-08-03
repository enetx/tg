package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// SetStickerMaskPosition represents a request to set sticker mask position.
type SetStickerMaskPosition struct {
	ctx          *Context
	sticker      gotgbot.InputFileOrString
	maskPosition *gotgbot.MaskPosition
	opts         *gotgbot.SetStickerMaskPositionOpts
}

// MaskPosition sets the mask position for the sticker.
func (ssmp *SetStickerMaskPosition) MaskPosition(
	point g.String,
	xShift, yShift, scale float64,
) *SetStickerMaskPosition {
	ssmp.maskPosition = &gotgbot.MaskPosition{
		Point:  point.Std(),
		XShift: xShift,
		YShift: yShift,
		Scale:  scale,
	}

	return ssmp
}

// Timeout sets a custom timeout for this request.
func (ssmp *SetStickerMaskPosition) Timeout(duration time.Duration) *SetStickerMaskPosition {
	if ssmp.opts.RequestOpts == nil {
		ssmp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ssmp.opts.RequestOpts.Timeout = duration

	return ssmp
}

// APIURL sets a custom API URL for this request.
func (ssmp *SetStickerMaskPosition) APIURL(url g.String) *SetStickerMaskPosition {
	if ssmp.opts.RequestOpts == nil {
		ssmp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ssmp.opts.RequestOpts.APIURL = url.Std()

	return ssmp
}

// Send sets the sticker mask position.
func (ssmp *SetStickerMaskPosition) Send() g.Result[bool] {
	ssmp.opts.MaskPosition = ssmp.maskPosition
	return g.ResultOf(ssmp.ctx.Bot.Raw().SetStickerMaskPosition(ssmp.sticker, ssmp.opts))
}
