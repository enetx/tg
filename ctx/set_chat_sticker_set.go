package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// SetChatStickerSet represents a request to set a chat's sticker set.
type SetChatStickerSet struct {
	ctx            *Context
	stickerSetName String
	opts           *gotgbot.SetChatStickerSetOpts
	chatID         Option[int64]
}

// ChatID sets the target chat ID.
func (scss *SetChatStickerSet) ChatID(chatID int64) *SetChatStickerSet {
	scss.chatID = Some(chatID)
	return scss
}

// Timeout sets a custom timeout for this request.
func (scss *SetChatStickerSet) Timeout(duration time.Duration) *SetChatStickerSet {
	if scss.opts.RequestOpts == nil {
		scss.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	scss.opts.RequestOpts.Timeout = duration

	return scss
}

// APIURL sets a custom API URL for this request.
func (scss *SetChatStickerSet) APIURL(url String) *SetChatStickerSet {
	if scss.opts.RequestOpts == nil {
		scss.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	scss.opts.RequestOpts.APIURL = url.Std()

	return scss
}

// Send sets the chat sticker set and returns the result.
func (scss *SetChatStickerSet) Send() Result[bool] {
	chatID := scss.chatID.UnwrapOr(scss.ctx.EffectiveChat.Id)
	return ResultOf(scss.ctx.Bot.Raw().SetChatStickerSet(chatID, scss.stickerSetName.Std(), scss.opts))
}
