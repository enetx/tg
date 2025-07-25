package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// GetCustomEmojiStickers represents a request to get custom emoji stickers.
type GetCustomEmojiStickers struct {
	ctx            *Context
	customEmojiIDs Slice[String]
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
func (gces *GetCustomEmojiStickers) APIURL(url String) *GetCustomEmojiStickers {
	if gces.opts.RequestOpts == nil {
		gces.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gces.opts.RequestOpts.APIURL = url.Std()

	return gces
}

// Send retrieves the custom emoji stickers.
func (gces *GetCustomEmojiStickers) Send() Result[Slice[gotgbot.Sticker]] {
	stickers, err := gces.ctx.Bot.Raw().GetCustomEmojiStickers(gces.customEmojiIDs.ToStringSlice(), gces.opts)
	return ResultOf[Slice[gotgbot.Sticker]](stickers, err)
}
