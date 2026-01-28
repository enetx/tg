package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/file"
)

// AddStickerToSet represents a request to add a sticker to an existing set.
type AddStickerToSet struct {
	ctx     *Context
	userID  int64
	name    g.String
	sticker gotgbot.InputSticker
	opts    *gotgbot.AddStickerToSetOpts
}

// File sets the sticker file.
func (ats *AddStickerToSet) File(filename file.InputFile) *AddStickerToSet {
	ats.sticker.Sticker = filename.Doc
	return ats
}

// Format sets the sticker format.
func (ats *AddStickerToSet) Format(format g.String) *AddStickerToSet {
	ats.sticker.Format = format.Std()
	return ats
}

// EmojiList sets the emoji list for the sticker.
func (ats *AddStickerToSet) EmojiList(emojis g.Slice[g.String]) *AddStickerToSet {
	ats.sticker.EmojiList = g.TransformSlice(emojis, g.String.Std)
	return ats
}

// Keywords sets keywords for the sticker.
func (ats *AddStickerToSet) Keywords(keywords g.Slice[g.String]) *AddStickerToSet {
	ats.sticker.Keywords = g.TransformSlice(keywords, g.String.Std)
	return ats
}

// MaskPosition sets the mask position for the sticker.
func (ats *AddStickerToSet) MaskPosition(point g.String, xShift, yShift, scale float64) *AddStickerToSet {
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
func (ats *AddStickerToSet) APIURL(url g.String) *AddStickerToSet {
	if ats.opts.RequestOpts == nil {
		ats.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ats.opts.RequestOpts.APIURL = url.Std()

	return ats
}

// Send adds the sticker to the set and returns the result.
func (ats *AddStickerToSet) Send() g.Result[bool] {
	return g.ResultOf(ats.ctx.Bot.Raw().AddStickerToSet(ats.userID, ats.name.Std(), ats.sticker, ats.opts))
}
