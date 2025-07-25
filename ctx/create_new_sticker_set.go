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

// StickerBuilder represents a builder for individual stickers.
type StickerBuilder struct {
	parent  *CreateNewStickerSet
	sticker gotgbot.InputSticker
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

// Sticker creates a new sticker builder for configuring individual sticker properties.
func (cns *CreateNewStickerSet) Sticker(filename, format String, emojiList Slice[String]) *StickerBuilder {
	sticker := gotgbot.InputSticker{
		Sticker:   filename.Std(),
		Format:    format.Std(),
		EmojiList: emojiList.ToStringSlice(),
	}

	return &StickerBuilder{
		parent:  cns,
		sticker: sticker,
	}
}

// Keywords sets search keywords for the sticker.
func (sb *StickerBuilder) Keywords(keywords Slice[String]) *StickerBuilder {
	sb.sticker.Keywords = keywords.ToStringSlice()
	return sb
}

// MaskPosition sets the mask position for mask stickers.
func (sb *StickerBuilder) MaskPosition(point String, xShift, yShift, scale float64) *StickerBuilder {
	sb.sticker.MaskPosition = &gotgbot.MaskPosition{
		Point:  point.Std(),
		XShift: xShift,
		YShift: yShift,
		Scale:  scale,
	}

	return sb
}

// Add completes the sticker configuration and adds it to the sticker set.
func (sb *StickerBuilder) Add() *CreateNewStickerSet {
	sb.parent.stickers.Push(sb.sticker)
	return sb.parent
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
