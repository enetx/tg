package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// SetStickerEmojiList represents a request to set sticker emoji list.
type SetStickerEmojiList struct {
	ctx       *Context
	sticker   String
	emojiList Slice[String]
	opts      *gotgbot.SetStickerEmojiListOpts
}

// EmojiList sets the emoji list for the sticker.
func (ssel *SetStickerEmojiList) EmojiList(emojis Slice[String]) *SetStickerEmojiList {
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
func (ssel *SetStickerEmojiList) APIURL(url String) *SetStickerEmojiList {
	if ssel.opts.RequestOpts == nil {
		ssel.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ssel.opts.RequestOpts.APIURL = url.Std()

	return ssel
}

// Send sets the sticker emoji list.
func (ssel *SetStickerEmojiList) Send() Result[bool] {
	return ResultOf(ssel.ctx.Bot.Raw().
		SetStickerEmojiList(ssel.sticker.Std(), ssel.emojiList.ToStringSlice(), ssel.opts),
	)
}
