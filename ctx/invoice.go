package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
)

type Invoice struct {
	ctx      *Context
	title    String
	desc     String
	payload  String
	currency String
	prices   Slice[gotgbot.LabeledPrice]
	chatID   Option[int64]
	opts     *gotgbot.SendInvoiceOpts
}

// To sets the target chat ID for the invoice.
func (c *Invoice) To(chatID int64) *Invoice {
	c.chatID = Some(chatID)
	return c
}

// Price adds a labeled price item to the invoice.
func (c *Invoice) Price(label String, amount int) *Invoice {
	c.prices.Push(gotgbot.LabeledPrice{Label: label.Std(), Amount: int64(amount)})
	return c
}

// Thread sets the message thread ID for the invoice.
func (c *Invoice) Thread(id int64) *Invoice {
	c.opts.MessageThreadId = id
	return c
}

// ProviderToken sets the payment provider token.
func (c *Invoice) ProviderToken(token String) *Invoice {
	c.opts.ProviderToken = token.Std()
	return c
}

// MaxTip sets the maximum accepted tip amount.
func (c *Invoice) MaxTip(amount int64) *Invoice {
	c.opts.MaxTipAmount = amount
	return c
}

// SuggestedTips sets suggested tip amounts for the invoice.
func (c *Invoice) SuggestedTips(tips ...int64) *Invoice {
	c.opts.SuggestedTipAmounts = tips
	return c
}

// StartParameter sets the unique deep-linking parameter.
func (c *Invoice) StartParameter(param String) *Invoice {
	c.opts.StartParameter = param.Std()
	return c
}

// ProviderData sets JSON-encoded data for the payment provider.
func (c *Invoice) ProviderData(data String) *Invoice {
	c.opts.ProviderData = data.Std()
	return c
}

// Photo sets the product photo URL and dimensions.
func (c *Invoice) Photo(url String, size, width, height int64) *Invoice {
	c.opts.PhotoUrl = url.Std()
	c.opts.PhotoSize = size
	c.opts.PhotoWidth = width
	c.opts.PhotoHeight = height

	return c
}

// NeedName requests user's full name for payment.
func (c *Invoice) NeedName() *Invoice {
	c.opts.NeedName = true
	return c
}

// NeedPhone requests user's phone number for payment.
func (c *Invoice) NeedPhone() *Invoice {
	c.opts.NeedPhoneNumber = true
	return c
}

// NeedEmail requests user's email address for payment.
func (c *Invoice) NeedEmail() *Invoice {
	c.opts.NeedEmail = true
	return c
}

// NeedShipping requests user's shipping address for payment.
func (c *Invoice) NeedShipping() *Invoice {
	c.opts.NeedShippingAddress = true
	return c
}

// SendPhone sends the user's phone number to the payment provider.
func (c *Invoice) SendPhone() *Invoice {
	c.opts.SendPhoneNumberToProvider = true
	return c
}

// SendEmail sends the user's email to the payment provider.
func (c *Invoice) SendEmail() *Invoice {
	c.opts.SendEmailToProvider = true
	return c
}

// Flexible enables flexible pricing (final price depends on shipping).
func (c *Invoice) Flexible() *Invoice {
	c.opts.IsFlexible = true
	return c
}

// Silent disables notification for the invoice message.
func (c *Invoice) Silent() *Invoice {
	c.opts.DisableNotification = true
	return c
}

// Protect enables content protection for the invoice message.
func (c *Invoice) Protect() *Invoice {
	c.opts.ProtectContent = true
	return c
}

// AllowPaidBroadcast allows the invoice to be sent in paid broadcast channels.
func (c *Invoice) AllowPaidBroadcast() *Invoice {
	c.opts.AllowPaidBroadcast = true
	return c
}

// Effect sets a message effect for the invoice message.
func (c *Invoice) Effect(effect string) *Invoice {
	c.opts.MessageEffectId = effect
	return c
}

// ReplyTo sets the message ID to reply to.
func (c *Invoice) ReplyTo(messageID int64) *Invoice {
	c.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return c
}

// Markup sets the reply markup keyboard for the invoice message.
func (c *Invoice) Markup(kb keyboard.KeyboardBuilder) *Invoice {
	if markup, ok := kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
		c.opts.ReplyMarkup = markup
	}

	return c
}

// Timeout sets a custom timeout for this request.
func (c *Invoice) Timeout(duration time.Duration) *Invoice {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *Invoice) APIURL(url String) *Invoice {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Send sends the invoice to Telegram and returns the result.
func (c *Invoice) Send() Result[*gotgbot.Message] {
	return ResultOf(c.ctx.Bot.Raw().SendInvoice(
		c.chatID.UnwrapOr(c.ctx.EffectiveChat.Id),
		c.title.Std(),
		c.desc.Std(),
		c.payload.Std(),
		c.currency.Std(),
		c.prices,
		c.opts,
	))
}
