package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// SetStickerPositionInSet represents a request to set sticker position in set.
type SetStickerPositionInSet struct {
	ctx      *Context
	sticker  String
	position int64
	opts     *gotgbot.SetStickerPositionInSetOpts
}

// Timeout sets a custom timeout for this request.
func (sspis *SetStickerPositionInSet) Timeout(duration time.Duration) *SetStickerPositionInSet {
	if sspis.opts.RequestOpts == nil {
		sspis.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sspis.opts.RequestOpts.Timeout = duration

	return sspis
}

// APIURL sets a custom API URL for this request.
func (sspis *SetStickerPositionInSet) APIURL(url String) *SetStickerPositionInSet {
	if sspis.opts.RequestOpts == nil {
		sspis.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sspis.opts.RequestOpts.APIURL = url.Std()

	return sspis
}

// Send sets the sticker position in the set.
func (sspis *SetStickerPositionInSet) Send() Result[bool] {
	return ResultOf(sspis.ctx.Bot.Raw().SetStickerPositionInSet(sspis.sticker.Std(), int64(sspis.position), sspis.opts))
}
