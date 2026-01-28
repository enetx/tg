package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// GetCustomEmojiStickers represents a request to get custom emoji stickers.
type GetCustomEmojiStickers struct {
	ctx            *Context
	customEmojiIDs g.Slice[g.String]
	opts           *gotgbot.GetCustomEmojiStickersOpts
}

// Timeout sets a custom timeout for this request.
func (gces *GetCustomEmojiStickers) Timeout(duration time.Duration) *GetCustomEmojiStickers {
	if gces.opts.RequestOpts == nil {
		gces.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gces.opts.RequestOpts.Timeout = duration

	return gces
}

// APIURL sets a custom API URL for this request.
func (gces *GetCustomEmojiStickers) APIURL(url g.String) *GetCustomEmojiStickers {
	if gces.opts.RequestOpts == nil {
		gces.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gces.opts.RequestOpts.APIURL = url.Std()

	return gces
}

// Send retrieves the custom emoji stickers.
func (gces *GetCustomEmojiStickers) Send() g.Result[g.Slice[gotgbot.Sticker]] {
	stickers, err := gces.ctx.Bot.Raw().GetCustomEmojiStickers(g.TransformSlice(gces.customEmojiIDs, g.String.Std), gces.opts)
	return g.ResultOf[g.Slice[gotgbot.Sticker]](stickers, err)
}
