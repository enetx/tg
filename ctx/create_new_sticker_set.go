package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// CreateNewStickerSet represents a request to create a new sticker set.
type CreateNewStickerSet struct {
	ctx      *Context
	userID   int64
	name     String
	title    String
	stickers Slice[gotgbot.InputSticker]
	opts     *gotgbot.CreateNewStickerSetOpts
}

// StickerType sets the type of stickers in the set.
func (cns *CreateNewStickerSet) StickerType(stickerType String) *CreateNewStickerSet {
	cns.opts.StickerType = stickerType.Std()
	return cns
}

// NeedsRepainting marks stickers for repainting to custom emoji.
func (cns *CreateNewStickerSet) NeedsRepainting() *CreateNewStickerSet {
	cns.opts.NeedsRepainting = true
	return cns
}

// AddSticker adds a sticker to the new sticker set.
func (cns *CreateNewStickerSet) AddSticker(filename, format String, emojiList Slice[String]) *CreateNewStickerSet {
	sticker := gotgbot.InputSticker{
		Sticker:   filename.Std(),
		Format:    format.Std(),
		EmojiList: emojiList.ToStringSlice(),
	}

	cns.stickers.Push(sticker)

	return cns
}

// Keywords sets keywords for the last added sticker.
func (cns *CreateNewStickerSet) Keywords(keywords Slice[String]) *CreateNewStickerSet {
	if cns.stickers.NotEmpty() {
		cns.stickers[len(cns.stickers)-1].Keywords = keywords.ToStringSlice()
	}

	return cns
}

// MaskPosition sets the mask position for the last added sticker.
func (cns *CreateNewStickerSet) MaskPosition(point String, xShift, yShift, scale float64) *CreateNewStickerSet {
	if cns.stickers.NotEmpty() {
		cns.stickers[len(cns.stickers)-1].MaskPosition = &gotgbot.MaskPosition{
			Point:  point.Std(),
			XShift: xShift,
			YShift: yShift,
			Scale:  scale,
		}
	}

	return cns
}

// Timeout sets a custom timeout for this request.
func (cns *CreateNewStickerSet) Timeout(duration time.Duration) *CreateNewStickerSet {
	if cns.opts.RequestOpts == nil {
		cns.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	cns.opts.RequestOpts.Timeout = duration

	return cns
}

// APIURL sets a custom API URL for this request.
func (cns *CreateNewStickerSet) APIURL(url String) *CreateNewStickerSet {
	if cns.opts.RequestOpts == nil {
		cns.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	cns.opts.RequestOpts.APIURL = url.Std()

	return cns
}

// Send creates the new sticker set and returns the result.
func (cns *CreateNewStickerSet) Send() Result[bool] {
	if len(cns.stickers) == 0 {
		return Err[bool](Errorf("no stickers added to sticker set"))
	}

	return ResultOf(cns.ctx.Bot.Raw().
		CreateNewStickerSet(cns.userID, cns.name.Std(), cns.title.Std(), cns.stickers, cns.opts))
}
