package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// ReplaceStickerInSet represents a request to replace a sticker in a sticker set.
type ReplaceStickerInSet struct {
	ctx        *Context
	userID     int64
	name       g.String
	oldSticker g.String
	sticker    gotgbot.InputSticker
	opts       *gotgbot.ReplaceStickerInSetOpts
}

// Sticker sets the new sticker to replace the old one.
func (rsis *ReplaceStickerInSet) Sticker(sticker gotgbot.InputSticker) *ReplaceStickerInSet {
	rsis.sticker = sticker
	return rsis
}

// Timeout sets a custom timeout for this request.
func (rsis *ReplaceStickerInSet) Timeout(duration time.Duration) *ReplaceStickerInSet {
	if rsis.opts.RequestOpts == nil {
		rsis.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rsis.opts.RequestOpts.Timeout = duration

	return rsis
}

// APIURL sets a custom API URL for this request.
func (rsis *ReplaceStickerInSet) APIURL(url g.String) *ReplaceStickerInSet {
	if rsis.opts.RequestOpts == nil {
		rsis.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rsis.opts.RequestOpts.APIURL = url.Std()

	return rsis
}

// Send replaces the sticker in the sticker set.
func (rsis *ReplaceStickerInSet) Send() g.Result[bool] {
	return g.ResultOf(rsis.ctx.Bot.Raw().ReplaceStickerInSet(
		rsis.userID,
		rsis.name.Std(),
		rsis.oldSticker.Std(),
		rsis.sticker,
		rsis.opts,
	))
}
