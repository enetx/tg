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
func (i *Invoice) To(chatID int64) *Invoice {
	i.chatID = Some(chatID)
	return i
}

// Price adds a labeled price item to the invoice.
func (i *Invoice) Price(label String, amount int) *Invoice {
	i.prices.Push(gotgbot.LabeledPrice{Label: label.Std(), Amount: int64(amount)})
	return i
}

// Thread sets the message thread ID for the invoice.
func (i *Invoice) Thread(id int64) *Invoice {
	i.opts.MessageThreadId = id
	return i
}

// ProviderToken sets the payment provider token.
func (i *Invoice) ProviderToken(token String) *Invoice {
	i.opts.ProviderToken = token.Std()
	return i
}

// MaxTip sets the maximum accepted tip amount.
func (i *Invoice) MaxTip(amount int64) *Invoice {
	i.opts.MaxTipAmount = amount
	return i
}

// SuggestedTips sets suggested tip amounts for the invoice.
func (i *Invoice) SuggestedTips(tips ...int64) *Invoice {
	i.opts.SuggestedTipAmounts = tips
	return i
}

// StartParameter sets the unique deep-linking parameter.
func (i *Invoice) StartParameter(param String) *Invoice {
	i.opts.StartParameter = param.Std()
	return i
}

// ProviderData sets JSON-encoded data for the payment provider.
func (i *Invoice) ProviderData(data String) *Invoice {
	i.opts.ProviderData = data.Std()
	return i
}

// Photo sets the product photo URL and dimensions.
func (i *Invoice) Photo(url String, size, width, height int64) *Invoice {
	i.opts.PhotoUrl = url.Std()
	i.opts.PhotoSize = size
	i.opts.PhotoWidth = width
	i.opts.PhotoHeight = height

	return i
}

// NeedName requests user's full name for payment.
func (i *Invoice) NeedName() *Invoice {
	i.opts.NeedName = true
	return i
}

// NeedPhone requests user's phone number for payment.
func (i *Invoice) NeedPhone() *Invoice {
	i.opts.NeedPhoneNumber = true
	return i
}

// NeedEmail requests user's email address for payment.
func (i *Invoice) NeedEmail() *Invoice {
	i.opts.NeedEmail = true
	return i
}

// NeedShipping requests user's shipping address for payment.
func (i *Invoice) NeedShipping() *Invoice {
	i.opts.NeedShippingAddress = true
	return i
}

// SendPhone sends the user's phone number to the payment provider.
func (i *Invoice) SendPhone() *Invoice {
	i.opts.SendPhoneNumberToProvider = true
	return i
}

// SendEmail sends the user's email to the payment provider.
func (i *Invoice) SendEmail() *Invoice {
	i.opts.SendEmailToProvider = true
	return i
}

// Flexible enables flexible pricing (final price depends on shipping).
func (i *Invoice) Flexible() *Invoice {
	i.opts.IsFlexible = true
	return i
}

// Silent disables notification for the invoice message.
func (i *Invoice) Silent() *Invoice {
	i.opts.DisableNotification = true
	return i
}

// Protect enables content protection for the invoice message.
func (i *Invoice) Protect() *Invoice {
	i.opts.ProtectContent = true
	return i
}

// AllowPaidBroadcast allows the invoice to be sent in paid broadcast channels.
func (i *Invoice) AllowPaidBroadcast() *Invoice {
	i.opts.AllowPaidBroadcast = true
	return i
}

// Effect sets a message effect for the invoice message.
func (i *Invoice) Effect(effect string) *Invoice {
	i.opts.MessageEffectId = effect
	return i
}

// ReplyTo sets the message ID to reply to.
func (i *Invoice) ReplyTo(messageID int64) *Invoice {
	i.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return i
}

// Markup sets the reply markup keyboard for the invoice message.
func (i *Invoice) Markup(kb keyboard.KeyboardBuilder) *Invoice {
	if markup, ok := kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
		i.opts.ReplyMarkup = markup
	}

	return i
}

// Timeout sets a custom timeout for this request.
func (i *Invoice) Timeout(duration time.Duration) *Invoice {
	if i.opts.RequestOpts == nil {
		i.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	i.opts.RequestOpts.Timeout = duration

	return i
}

// APIURL sets a custom API URL for this request.
func (i *Invoice) APIURL(url String) *Invoice {
	if i.opts.RequestOpts == nil {
		i.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	i.opts.RequestOpts.APIURL = url.Std()

	return i
}

// Send sends the invoice to Telegram and returns the result.
func (i *Invoice) Send() Result[*gotgbot.Message] {
	return ResultOf(i.ctx.Bot.Raw().SendInvoice(
		i.chatID.UnwrapOr(i.ctx.EffectiveChat.Id),
		i.title.Std(),
		i.desc.Std(),
		i.payload.Std(),
		i.currency.Std(),
		i.prices,
		i.opts,
	))
}
