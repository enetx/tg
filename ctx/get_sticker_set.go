package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// GetStickerSet represents a request to get sticker set information.
type GetStickerSet struct {
	ctx  *Context
	name g.String
	opts *gotgbot.GetStickerSetOpts
}

// Timeout sets a custom timeout for this request.
func (gss *GetStickerSet) Timeout(duration time.Duration) *GetStickerSet {
	if gss.opts.RequestOpts == nil {
		gss.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gss.opts.RequestOpts.Timeout = duration

	return gss
}

// APIURL sets a custom API URL for this request.
func (gss *GetStickerSet) APIURL(url g.String) *GetStickerSet {
	if gss.opts.RequestOpts == nil {
		gss.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gss.opts.RequestOpts.APIURL = url.Std()

	return gss
}

// Send retrieves the sticker set information.
func (gss *GetStickerSet) Send() g.Result[*gotgbot.StickerSet] {
	return g.ResultOf(gss.ctx.Bot.Raw().GetStickerSet(gss.name.Std(), gss.opts))
}
