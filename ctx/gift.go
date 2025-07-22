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
func (c *SendGift) To(userID int64) *SendGift {
	c.userID = Some(userID)
	return c
}

// ToChat sets the target chat ID for the gift.
func (c *SendGift) ToChat(chatID int64) *SendGift {
	c.chatID = Some(chatID)
	return c
}

// PayForUpgrade enables paying for upgrade from bot's balance.
func (c *SendGift) PayForUpgrade() *SendGift {
	c.opts.PayForUpgrade = true
	return c
}

// Text sets the text shown with the gift (0-128 characters).
func (c *SendGift) Text(text String) *SendGift {
	c.opts.Text = text.Std()
	return c
}

// HTML sets the gift text parse mode to HTML.
func (c *SendGift) HTML() *SendGift {
	c.opts.TextParseMode = "HTML"
	return c
}

// Markdown sets the gift text parse mode to MarkdownV2.
func (c *SendGift) Markdown() *SendGift {
	c.opts.TextParseMode = "MarkdownV2"
	return c
}

// TextEntities sets special entities in the gift text using Entities builder.
func (c *SendGift) TextEntities(e *entities.Entities) *SendGift {
	c.opts.TextEntities = e.Std()
	return c
}

// Timeout sets a custom timeout for this request.
func (c *SendGift) Timeout(duration time.Duration) *SendGift {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *SendGift) APIURL(url String) *SendGift {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Send executes the SendGift request.
func (c *SendGift) Send() Result[bool] {
	if c.userID.IsSome() {
		c.opts.UserId = c.userID.Unwrap()
	} else if c.chatID.IsSome() {
		c.opts.ChatId = c.chatID.Unwrap()
	} else {
		c.opts.UserId = c.ctx.EffectiveUser.Id
	}

	return ResultOf(c.ctx.Bot.Raw().SendGift(c.giftID.Std(), c.opts))
}

// GetAvailableGifts is a request builder for getting available gifts.
type GetAvailableGifts struct {
	ctx  *Context
	opts *gotgbot.GetAvailableGiftsOpts
}

// Timeout sets a custom timeout for this request.
func (c *GetAvailableGifts) Timeout(duration time.Duration) *GetAvailableGifts {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *GetAvailableGifts) APIURL(url String) *GetAvailableGifts {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Send executes the GetAvailableGifts request.
func (c *GetAvailableGifts) Send() Result[*gotgbot.Gifts] {
	return ResultOf(c.ctx.Bot.Raw().GetAvailableGifts(c.opts))
}

// ConvertGiftToStars is a request builder for converting gifts to stars.
type ConvertGiftToStars struct {
	ctx                  *Context
	businessConnectionID String
	ownedGiftID          String
	opts                 *gotgbot.ConvertGiftToStarsOpts
}

// Timeout sets a custom timeout for this request.
func (c *ConvertGiftToStars) Timeout(duration time.Duration) *ConvertGiftToStars {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *ConvertGiftToStars) APIURL(url String) *ConvertGiftToStars {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Send executes the ConvertGiftToStars request.
func (c *ConvertGiftToStars) Send() Result[bool] {
	return ResultOf(c.ctx.Bot.Raw().ConvertGiftToStars(
		c.businessConnectionID.Std(),
		c.ownedGiftID.Std(),
		c.opts,
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
func (c *TransferGift) StarCount(count int64) *TransferGift {
	c.opts.StarCount = count
	return c
}

// Timeout sets a custom timeout for this request.
func (c *TransferGift) Timeout(duration time.Duration) *TransferGift {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *TransferGift) APIURL(url String) *TransferGift {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Send executes the TransferGift request.
func (c *TransferGift) Send() Result[bool] {
	return ResultOf(c.ctx.Bot.Raw().TransferGift(
		c.businessConnectionID.Std(),
		c.ownedGiftID.Std(),
		c.newOwnerChatID,
		c.opts,
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
func (c *UpgradeGift) KeepOriginalDetails() *UpgradeGift {
	c.opts.KeepOriginalDetails = true
	return c
}

// StarCount sets the amount of stars to pay for upgrade from business balance.
func (c *UpgradeGift) StarCount(count int64) *UpgradeGift {
	c.opts.StarCount = count
	return c
}

// Timeout sets a custom timeout for this request.
func (c *UpgradeGift) Timeout(duration time.Duration) *UpgradeGift {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *UpgradeGift) APIURL(url String) *UpgradeGift {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Send executes the UpgradeGift request.
func (c *UpgradeGift) Send() Result[bool] {
	return ResultOf(c.ctx.Bot.Raw().UpgradeGift(
		c.businessConnectionID.Std(),
		c.ownedGiftID.Std(),
		c.opts,
	))
}
