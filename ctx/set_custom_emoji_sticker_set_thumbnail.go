package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// SetCustomEmojiStickerSetThumbnail represents a request to set the thumbnail of a custom emoji sticker set.
type SetCustomEmojiStickerSetThumbnail struct {
	ctx  *Context
	name g.String
	opts *gotgbot.SetCustomEmojiStickerSetThumbnailOpts
}

// CustomEmojiID sets the custom emoji identifier for the thumbnail.
func (scesst *SetCustomEmojiStickerSetThumbnail) CustomEmojiID(emojiID g.String) *SetCustomEmojiStickerSetThumbnail {
	scesst.opts.CustomEmojiId = emojiID.Std()
	return scesst
}

// DropThumbnail removes the thumbnail and uses the first sticker as the thumbnail.
func (scesst *SetCustomEmojiStickerSetThumbnail) DropThumbnail() *SetCustomEmojiStickerSetThumbnail {
	scesst.opts.CustomEmojiId = ""
	return scesst
}

// Timeout sets a custom timeout for this request.
func (scesst *SetCustomEmojiStickerSetThumbnail) Timeout(duration time.Duration) *SetCustomEmojiStickerSetThumbnail {
	if scesst.opts.RequestOpts == nil {
		scesst.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	scesst.opts.RequestOpts.Timeout = duration

	return scesst
}

// APIURL sets a custom API URL for this request.
func (scesst *SetCustomEmojiStickerSetThumbnail) APIURL(url g.String) *SetCustomEmojiStickerSetThumbnail {
	if scesst.opts.RequestOpts == nil {
		scesst.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	scesst.opts.RequestOpts.APIURL = url.Std()

	return scesst
}

// Send sets the custom emoji sticker set thumbnail.
func (scesst *SetCustomEmojiStickerSetThumbnail) Send() g.Result[bool] {
	return g.ResultOf(scesst.ctx.Bot.Raw().SetCustomEmojiStickerSetThumbnail(scesst.name.Std(), scesst.opts))
}
