package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/entities"
)

// SendGift is a request builder for sending gifts.
type SendGift struct {
	ctx    *Context
	giftID String
	userID Option[int64]
	chatID Option[int64]
	opts   *gotgbot.SendGiftOpts
}

// To sets the target user ID for the gift.
func (sg *SendGift) To(userID int64) *SendGift {
	sg.userID = Some(userID)
	return sg
}

// ToChat sets the target chat ID for the gift.
func (sg *SendGift) ToChat(chatID int64) *SendGift {
	sg.chatID = Some(chatID)
	return sg
}

// PayForUpgrade enables paying for upgrade from bot's balance.
func (sg *SendGift) PayForUpgrade() *SendGift {
	sg.opts.PayForUpgrade = true
	return sg
}

// Text sets the text shown with the gift (0-128 characters).
func (sg *SendGift) Text(text String) *SendGift {
	sg.opts.Text = text.Std()
	return sg
}

// HTML sets the gift text parse mode to HTML.
func (sg *SendGift) HTML() *SendGift {
	sg.opts.TextParseMode = "HTML"
	return sg
}

// Markdown sets the gift text parse mode to MarkdownV2.
func (sg *SendGift) Markdown() *SendGift {
	sg.opts.TextParseMode = "MarkdownV2"
	return sg
}

// TextEntities sets special entities in the gift text using Entities builder.
func (sg *SendGift) TextEntities(e *entities.Entities) *SendGift {
	sg.opts.TextEntities = e.Std()
	return sg
}

// Timeout sets a custom timeout for this request.
func (sg *SendGift) Timeout(duration time.Duration) *SendGift {
	if sg.opts.RequestOpts == nil {
		sg.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sg.opts.RequestOpts.Timeout = duration

	return sg
}

// APIURL sets a custom API URL for this request.
func (sg *SendGift) APIURL(url String) *SendGift {
	if sg.opts.RequestOpts == nil {
		sg.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sg.opts.RequestOpts.APIURL = url.Std()

	return sg
}

// Send executes the SendGift request.
func (sg *SendGift) Send() Result[bool] {
	if sg.userID.IsSome() {
		sg.opts.UserId = sg.userID.Some()
	} else if sg.chatID.IsSome() {
		sg.opts.ChatId = sg.chatID.Some()
	} else {
		sg.opts.UserId = sg.ctx.EffectiveUser.Id
	}

	return ResultOf(sg.ctx.Bot.Raw().SendGift(sg.giftID.Std(), sg.opts))
}

// GetAvailableGifts is a request builder for getting available gifts.
type GetAvailableGifts struct {
	ctx  *Context
	opts *gotgbot.GetAvailableGiftsOpts
}

// Timeout sets a custom timeout for this request.
func (gags *GetAvailableGifts) Timeout(duration time.Duration) *GetAvailableGifts {
	if gags.opts.RequestOpts == nil {
		gags.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gags.opts.RequestOpts.Timeout = duration

	return gags
}

// APIURL sets a custom API URL for this request.
func (gags *GetAvailableGifts) APIURL(url String) *GetAvailableGifts {
	if gags.opts.RequestOpts == nil {
		gags.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gags.opts.RequestOpts.APIURL = url.Std()

	return gags
}

// Send executes the GetAvailableGifts request.
func (gags *GetAvailableGifts) Send() Result[*gotgbot.Gifts] {
	return ResultOf(gags.ctx.Bot.Raw().GetAvailableGifts(gags.opts))
}

// ConvertGiftToStars is a request builder for converting gifts to stars.
type ConvertGiftToStars struct {
	ctx                  *Context
	businessConnectionID String
	ownedGiftID          String
	opts                 *gotgbot.ConvertGiftToStarsOpts
}

// Timeout sets a custom timeout for this request.
func (cgts *ConvertGiftToStars) Timeout(duration time.Duration) *ConvertGiftToStars {
	if cgts.opts.RequestOpts == nil {
		cgts.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	cgts.opts.RequestOpts.Timeout = duration

	return cgts
}

// APIURL sets a custom API URL for this request.
func (cgts *ConvertGiftToStars) APIURL(url String) *ConvertGiftToStars {
	if cgts.opts.RequestOpts == nil {
		cgts.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	cgts.opts.RequestOpts.APIURL = url.Std()

	return cgts
}

// Send executes the ConvertGiftToStars request.
func (cgts *ConvertGiftToStars) Send() Result[bool] {
	return ResultOf(cgts.ctx.Bot.Raw().ConvertGiftToStars(
		cgts.businessConnectionID.Std(),
		cgts.ownedGiftID.Std(),
		cgts.opts,
	))
}

// TransferGift is a request builder for transferring gifts.
type TransferGift struct {
	ctx                  *Context
	businessConnectionID String
	ownedGiftID          String
	newOwnerChatID       int64
	opts                 *gotgbot.TransferGiftOpts
}

// StarCount sets the amount of stars to pay for transfer from business balance.
func (tg *TransferGift) StarCount(count int64) *TransferGift {
	tg.opts.StarCount = count
	return tg
}

// Timeout sets a custom timeout for this request.
func (tg *TransferGift) Timeout(duration time.Duration) *TransferGift {
	if tg.opts.RequestOpts == nil {
		tg.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	tg.opts.RequestOpts.Timeout = duration

	return tg
}

// APIURL sets a custom API URL for this request.
func (tg *TransferGift) APIURL(url String) *TransferGift {
	if tg.opts.RequestOpts == nil {
		tg.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	tg.opts.RequestOpts.APIURL = url.Std()

	return tg
}

// Send executes the TransferGift request.
func (tg *TransferGift) Send() Result[bool] {
	return ResultOf(tg.ctx.Bot.Raw().TransferGift(
		tg.businessConnectionID.Std(),
		tg.ownedGiftID.Std(),
		tg.newOwnerChatID,
		tg.opts,
	))
}

// UpgradeGift is a request builder for upgrading gifts.
type UpgradeGift struct {
	ctx                  *Context
	businessConnectionID String
	ownedGiftID          String
	opts                 *gotgbot.UpgradeGiftOpts
}

// KeepOriginalDetails preserves original gift text, sender, and receiver.
func (ug *UpgradeGift) KeepOriginalDetails() *UpgradeGift {
	ug.opts.KeepOriginalDetails = true
	return ug
}

// StarCount sets the amount of stars to pay for upgrade from business balance.
func (ug *UpgradeGift) StarCount(count int64) *UpgradeGift {
	ug.opts.StarCount = count
	return ug
}

// Timeout sets a custom timeout for this request.
func (ug *UpgradeGift) Timeout(duration time.Duration) *UpgradeGift {
	if ug.opts.RequestOpts == nil {
		ug.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ug.opts.RequestOpts.Timeout = duration

	return ug
}

// APIURL sets a custom API URL for this request.
func (ug *UpgradeGift) APIURL(url String) *UpgradeGift {
	if ug.opts.RequestOpts == nil {
		ug.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ug.opts.RequestOpts.APIURL = url.Std()

	return ug
}

// Send executes the UpgradeGift request.
func (ug *UpgradeGift) Send() Result[bool] {
	return ResultOf(ug.ctx.Bot.Raw().UpgradeGift(
		ug.businessConnectionID.Std(),
		ug.ownedGiftID.Std(),
		ug.opts,
	))
}
