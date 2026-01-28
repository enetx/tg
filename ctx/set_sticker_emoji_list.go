package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// SetStickerEmojiList represents a request to set sticker emoji list.
type SetStickerEmojiList struct {
	ctx       *Context
	sticker   gotgbot.InputFileOrString
	emojiList g.Slice[g.String]
	opts      *gotgbot.SetStickerEmojiListOpts
}

// EmojiList sets the emoji list for the sticker.
func (ssel *SetStickerEmojiList) EmojiList(emojis g.Slice[g.String]) *SetStickerEmojiList {
	ssel.emojiList = emojis
	return ssel
}

// Timeout sets a custom timeout for this request.
func (ssel *SetStickerEmojiList) Timeout(duration time.Duration) *SetStickerEmojiList {
	if ssel.opts.RequestOpts == nil {
		ssel.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ssel.opts.RequestOpts.Timeout = duration

	return ssel
}

// APIURL sets a custom API URL for this request.
func (ssel *SetStickerEmojiList) APIURL(url g.String) *SetStickerEmojiList {
	if ssel.opts.RequestOpts == nil {
		ssel.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ssel.opts.RequestOpts.APIURL = url.Std()

	return ssel
}

// Send sets the sticker emoji list.
func (ssel *SetStickerEmojiList) Send() g.Result[bool] {
	return g.ResultOf(ssel.ctx.Bot.Raw().
		SetStickerEmojiList(ssel.sticker, g.TransformSlice(ssel.emojiList, g.String.Std), ssel.opts),
	)
}
