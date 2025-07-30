package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// SetStickerSetTitle represents a request to set the title of a sticker set.
type SetStickerSetTitle struct {
	ctx   *Context
	name  g.String
	title g.String
	opts  *gotgbot.SetStickerSetTitleOpts
}

// Timeout sets a custom timeout for this request.
func (ssst *SetStickerSetTitle) Timeout(duration time.Duration) *SetStickerSetTitle {
	if ssst.opts.RequestOpts == nil {
		ssst.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ssst.opts.RequestOpts.Timeout = duration

	return ssst
}

// APIURL sets a custom API URL for this request.
func (ssst *SetStickerSetTitle) APIURL(url g.String) *SetStickerSetTitle {
	if ssst.opts.RequestOpts == nil {
		ssst.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ssst.opts.RequestOpts.APIURL = url.Std()

	return ssst
}

// Send sets the sticker set title.
func (ssst *SetStickerSetTitle) Send() g.Result[bool] {
	return g.ResultOf(ssst.ctx.Bot.Raw().SetStickerSetTitle(ssst.name.Std(), ssst.title.Std(), ssst.opts))
}
