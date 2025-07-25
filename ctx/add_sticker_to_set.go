package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// AddStickerToSet represents a request to add a sticker to an existing set.
type AddStickerToSet struct {
	ctx     *Context
	userID  int64
	name    String
	sticker gotgbot.InputSticker
	opts    *gotgbot.AddStickerToSetOpts
}

// File sets the sticker file.
func (ats *AddStickerToSet) File(filename String) *AddStickerToSet {
	ats.sticker.Sticker = string(filename)
	return ats
}

// Format sets the sticker format.
func (ats *AddStickerToSet) Format(format String) *AddStickerToSet {
	ats.sticker.Format = format.Std()
	return ats
}

// EmojiList sets the emoji list for the sticker.
func (ats *AddStickerToSet) EmojiList(emojis Slice[String]) *AddStickerToSet {
	ats.sticker.EmojiList = emojis.ToStringSlice()
	return ats
}

// Keywords sets keywords for the sticker.
func (ats *AddStickerToSet) Keywords(keywords Slice[String]) *AddStickerToSet {
	ats.sticker.Keywords = keywords.ToStringSlice()
	return ats
}

// MaskPosition sets the mask position for the sticker.
func (ats *AddStickerToSet) MaskPosition(point String, xShift, yShift, scale float64) *AddStickerToSet {
	ats.sticker.MaskPosition = &gotgbot.MaskPosition{
		Point:  point.Std(),
		XShift: xShift,
		YShift: yShift,
		Scale:  scale,
	}

	return ats
}

// Timeout sets a custom timeout for this request.
func (ats *AddStickerToSet) Timeout(duration time.Duration) *AddStickerToSet {
	if ats.opts.RequestOpts == nil {
		ats.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ats.opts.RequestOpts.Timeout = duration

	return ats
}

// APIURL sets a custom API URL for this request.
func (ats *AddStickerToSet) APIURL(url String) *AddStickerToSet {
	if ats.opts.RequestOpts == nil {
		ats.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ats.opts.RequestOpts.APIURL = url.Std()

	return ats
}

// Send adds the sticker to the set and returns the result.
func (ats *AddStickerToSet) Send() Result[bool] {
	return ResultOf(ats.ctx.Bot.Raw().AddStickerToSet(ats.userID, ats.name.Std(), ats.sticker, ats.opts))
}
