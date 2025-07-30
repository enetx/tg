package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// DeleteStickerSet represents a request to delete a sticker set.
type DeleteStickerSet struct {
	ctx  *Context
	name g.String
	opts *gotgbot.DeleteStickerSetOpts
}

// Timeout sets a custom timeout for this request.
func (dss *DeleteStickerSet) Timeout(duration time.Duration) *DeleteStickerSet {
	if dss.opts.RequestOpts == nil {
		dss.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dss.opts.RequestOpts.Timeout = duration

	return dss
}

// APIURL sets a custom API URL for this request.
func (dss *DeleteStickerSet) APIURL(url g.String) *DeleteStickerSet {
	if dss.opts.RequestOpts == nil {
		dss.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dss.opts.RequestOpts.APIURL = url.Std()

	return dss
}

// Send deletes the sticker set.
func (dss *DeleteStickerSet) Send() g.Result[bool] {
	return g.ResultOf(dss.ctx.Bot.Raw().DeleteStickerSet(dss.name.Std(), dss.opts))
}
