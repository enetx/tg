package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
)

type SendInvoice struct {
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
func (si *SendInvoice) To(chatID int64) *SendInvoice {
	si.chatID = Some(chatID)
	return si
}

// Price adds a labeled price item to the invoice.
func (si *SendInvoice) Price(label String, amount int64) *SendInvoice {
	si.prices.Push(gotgbot.LabeledPrice{Label: label.Std(), Amount: amount})
	return si
}

// Thread sets the message thread ID for the invoice.
func (si *SendInvoice) Thread(id int64) *SendInvoice {
	si.opts.MessageThreadId = id
	return si
}

// ProviderToken sets the payment provider token.
func (si *SendInvoice) ProviderToken(token String) *SendInvoice {
	si.opts.ProviderToken = token.Std()
	return si
}

// MaxTip sets the maximum accepted tip amount.
func (si *SendInvoice) MaxTip(amount int64) *SendInvoice {
	si.opts.MaxTipAmount = amount
	return si
}

// SuggestedTips sets suggested tip amounts for the invoice.
func (si *SendInvoice) SuggestedTips(tips ...int64) *SendInvoice {
	si.opts.SuggestedTipAmounts = tips
	return si
}

// StartParameter sets the unique deep-linking parameter.
func (si *SendInvoice) StartParameter(param String) *SendInvoice {
	si.opts.StartParameter = param.Std()
	return si
}

// ProviderData sets JSON-encoded data for the payment provider.
func (si *SendInvoice) ProviderData(data String) *SendInvoice {
	si.opts.ProviderData = data.Std()
	return si
}

// Photo sets the product photo URL and dimensions.
func (si *SendInvoice) Photo(url String, size, width, height int64) *SendInvoice {
	si.opts.PhotoUrl = url.Std()
	si.opts.PhotoSize = size
	si.opts.PhotoWidth = width
	si.opts.PhotoHeight = height

	return si
}

// NeedName requests user's full name for payment.
func (si *SendInvoice) NeedName() *SendInvoice {
	si.opts.NeedName = true
	return si
}

// NeedPhone requests user's phone number for payment.
func (si *SendInvoice) NeedPhone() *SendInvoice {
	si.opts.NeedPhoneNumber = true
	return si
}

// NeedEmail requests user's email address for payment.
func (si *SendInvoice) NeedEmail() *SendInvoice {
	si.opts.NeedEmail = true
	return si
}

// NeedShipping requests user's shipping address for payment.
func (si *SendInvoice) NeedShipping() *SendInvoice {
	si.opts.NeedShippingAddress = true
	return si
}

// SendPhone sends the user's phone number to the payment provider.
func (si *SendInvoice) SendPhone() *SendInvoice {
	si.opts.SendPhoneNumberToProvider = true
	return si
}

// SendEmail sends the user's email to the payment provider.
func (si *SendInvoice) SendEmail() *SendInvoice {
	si.opts.SendEmailToProvider = true
	return si
}

// Flexible enables flexible pricing (final price depends on shipping).
func (si *SendInvoice) Flexible() *SendInvoice {
	si.opts.IsFlexible = true
	return si
}

// Silent disables notification for the invoice message.
func (si *SendInvoice) Silent() *SendInvoice {
	si.opts.DisableNotification = true
	return si
}

// Protect enables content protection for the invoice message.
func (si *SendInvoice) Protect() *SendInvoice {
	si.opts.ProtectContent = true
	return si
}

// AllowPaidBroadcast allows the invoice to be sent in paid broadcast channels.
func (si *SendInvoice) AllowPaidBroadcast() *SendInvoice {
	si.opts.AllowPaidBroadcast = true
	return si
}

// Effect sets a message effect for the invoice message.
func (si *SendInvoice) Effect(effect string) *SendInvoice {
	si.opts.MessageEffectId = effect
	return si
}

// ReplyTo sets the message ID to reply to.
func (si *SendInvoice) ReplyTo(messageID int64) *SendInvoice {
	si.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return si
}

// Markup sets the reply markup keyboard for the invoice message.
func (si *SendInvoice) Markup(kb keyboard.Keyboard) *SendInvoice {
	if markup, ok := kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
		si.opts.ReplyMarkup = markup
	}

	return si
}

// Timeout sets a custom timeout for this request.
func (si *SendInvoice) Timeout(duration time.Duration) *SendInvoice {
	if si.opts.RequestOpts == nil {
		si.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	si.opts.RequestOpts.Timeout = duration

	return si
}

// APIURL sets a custom API URL for this request.
func (si *SendInvoice) APIURL(url String) *SendInvoice {
	if si.opts.RequestOpts == nil {
		si.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	si.opts.RequestOpts.APIURL = url.Std()

	return si
}

// Send sends the invoice to Telegram and returns the result.
func (si *SendInvoice) Send() Result[*gotgbot.Message] {
	return ResultOf(si.ctx.Bot.Raw().SendInvoice(
		si.chatID.UnwrapOr(si.ctx.EffectiveChat.Id),
		si.title.Std(),
		si.desc.Std(),
		si.payload.Std(),
		si.currency.Std(),
		si.prices,
		si.opts,
	))
}
