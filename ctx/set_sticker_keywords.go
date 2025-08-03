package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// SetStickerKeywords represents a request to set sticker keywords.
type SetStickerKeywords struct {
	ctx      *Context
	sticker  gotgbot.InputFileOrString
	keywords g.Slice[g.String]
	opts     *gotgbot.SetStickerKeywordsOpts
}

// Keywords sets the keywords for the sticker.
func (ssk *SetStickerKeywords) Keywords(keywords g.Slice[g.String]) *SetStickerKeywords {
	ssk.keywords = keywords
	return ssk
}

// Timeout sets a custom timeout for this request.
func (ssk *SetStickerKeywords) Timeout(duration time.Duration) *SetStickerKeywords {
	if ssk.opts.RequestOpts == nil {
		ssk.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ssk.opts.RequestOpts.Timeout = duration

	return ssk
}

// APIURL sets a custom API URL for this request.
func (ssk *SetStickerKeywords) APIURL(url g.String) *SetStickerKeywords {
	if ssk.opts.RequestOpts == nil {
		ssk.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ssk.opts.RequestOpts.APIURL = url.Std()

	return ssk
}

// Send sets the sticker keywords.
func (ssk *SetStickerKeywords) Send() g.Result[bool] {
	ssk.opts.Keywords = ssk.keywords.ToStringSlice()
	return g.ResultOf(ssk.ctx.Bot.Raw().SetStickerKeywords(ssk.sticker, ssk.opts))
}
