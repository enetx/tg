package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// DeleteChatStickerSet represents a request to delete a chat's sticker set.
type DeleteChatStickerSet struct {
	ctx    *Context
	opts   *gotgbot.DeleteChatStickerSetOpts
	chatID g.Option[int64]
}

// ChatID sets the target chat ID.
func (dcss *DeleteChatStickerSet) ChatID(chatID int64) *DeleteChatStickerSet {
	dcss.chatID = g.Some(chatID)
	return dcss
}

// Timeout sets a custom timeout for this request.
func (dcss *DeleteChatStickerSet) Timeout(duration time.Duration) *DeleteChatStickerSet {
	if dcss.opts.RequestOpts == nil {
		dcss.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dcss.opts.RequestOpts.Timeout = duration

	return dcss
}

// APIURL sets a custom API URL for this request.
func (dcss *DeleteChatStickerSet) APIURL(url g.String) *DeleteChatStickerSet {
	if dcss.opts.RequestOpts == nil {
		dcss.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dcss.opts.RequestOpts.APIURL = url.Std()

	return dcss
}

// Send deletes the chat sticker set and returns the result.
func (dcss *DeleteChatStickerSet) Send() g.Result[bool] {
	chatID := dcss.chatID.UnwrapOr(dcss.ctx.EffectiveChat.Id)
	return g.ResultOf(dcss.ctx.Bot.Raw().DeleteChatStickerSet(chatID, dcss.opts))
}
