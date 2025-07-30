package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
)

// SendGift is a request builder for sending gifts.
type SendGift struct {
	ctx    *Context
	giftID g.String
	userID g.Option[int64]
	chatID g.Option[int64]
	opts   *gotgbot.SendGiftOpts
}

// To sets the target user ID for the gift.
func (sg *SendGift) To(userID int64) *SendGift {
	sg.userID = g.Some(userID)
	return sg
}

// ToChat sets the target chat ID for the gift.
func (sg *SendGift) ToChat(chatID int64) *SendGift {
	sg.chatID = g.Some(chatID)
	return sg
}

// PayForUpgrade enables paying for upgrade from bot's balance.
func (sg *SendGift) PayForUpgrade() *SendGift {
	sg.opts.PayForUpgrade = true
	return sg
}

// Text sets the text shown with the gift (0-128 characters).
func (sg *SendGift) Text(text g.String) *SendGift {
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
func (sg *SendGift) APIURL(url g.String) *SendGift {
	if sg.opts.RequestOpts == nil {
		sg.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sg.opts.RequestOpts.APIURL = url.Std()

	return sg
}

// Send executes the SendGift request.
func (sg *SendGift) Send() g.Result[bool] {
	if sg.userID.IsSome() {
		sg.opts.UserId = sg.userID.Some()
	} else if sg.chatID.IsSome() {
		sg.opts.ChatId = sg.chatID.Some()
	} else {
		sg.opts.UserId = sg.ctx.EffectiveUser.Id
	}

	return g.ResultOf(sg.ctx.Bot.Raw().SendGift(sg.giftID.Std(), sg.opts))
}
