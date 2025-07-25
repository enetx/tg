package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// CreateInvoiceLink represents a request to create an invoice link.
type CreateInvoiceLink struct {
	ctx      *Context
	title    String
	desc     String
	payload  String
	currency String
	prices   Slice[gotgbot.LabeledPrice]
	opts     *gotgbot.CreateInvoiceLinkOpts
}

// Price adds a labeled price item to the invoice link.
func (cil *CreateInvoiceLink) Price(label String, amount int64) *CreateInvoiceLink {
	cil.prices.Push(gotgbot.LabeledPrice{Label: label.Std(), Amount: amount})
	return cil
}

// Business sets the business connection ID for the invoice link.
func (cil *CreateInvoiceLink) Business(id String) *CreateInvoiceLink {
	cil.opts.BusinessConnectionId = id.Std()
	return cil
}

// ProviderToken sets the payment provider token.
func (cil *CreateInvoiceLink) ProviderToken(token String) *CreateInvoiceLink {
	cil.opts.ProviderToken = token.Std()
	return cil
}

// SubscriptionPeriod sets the subscription period for Telegram Stars payments.
func (cil *CreateInvoiceLink) SubscriptionPeriod(seconds int64) *CreateInvoiceLink {
	cil.opts.SubscriptionPeriod = seconds
	return cil
}

// MaxTip sets the maximum accepted tip amount.
func (cil *CreateInvoiceLink) MaxTip(amount int64) *CreateInvoiceLink {
	cil.opts.MaxTipAmount = amount
	return cil
}

// SuggestedTips sets suggested tip amounts.
func (cil *CreateInvoiceLink) SuggestedTips(tips ...int64) *CreateInvoiceLink {
	cil.opts.SuggestedTipAmounts = tips
	return cil
}

// ProviderData sets JSON-serialized data for the payment provider.
func (cil *CreateInvoiceLink) ProviderData(data String) *CreateInvoiceLink {
	cil.opts.ProviderData = data.Std()
	return cil
}

// Photo sets the product photo URL and dimensions.
func (cil *CreateInvoiceLink) Photo(url String, size, width, height int64) *CreateInvoiceLink {
	cil.opts.PhotoUrl = url.Std()
	cil.opts.PhotoSize = size
	cil.opts.PhotoWidth = width
	cil.opts.PhotoHeight = height

	return cil
}

// NeedName requires the user's full name to complete the order.
func (cil *CreateInvoiceLink) NeedName() *CreateInvoiceLink {
	cil.opts.NeedName = true
	return cil
}

// NeedPhone requires the user's phone number to complete the order.
func (cil *CreateInvoiceLink) NeedPhone() *CreateInvoiceLink {
	cil.opts.NeedPhoneNumber = true
	return cil
}

// NeedEmail requires the user's email address to complete the order.
func (cil *CreateInvoiceLink) NeedEmail() *CreateInvoiceLink {
	cil.opts.NeedEmail = true
	return cil
}

// NeedShipping requires the user's shipping address to complete the order.
func (cil *CreateInvoiceLink) NeedShipping() *CreateInvoiceLink {
	cil.opts.NeedShippingAddress = true
	return cil
}

// SendPhone sends the user's phone number to the provider.
func (cil *CreateInvoiceLink) SendPhone() *CreateInvoiceLink {
	cil.opts.SendPhoneNumberToProvider = true
	return cil
}

// SendEmail sends the user's email to the provider.
func (cil *CreateInvoiceLink) SendEmail() *CreateInvoiceLink {
	cil.opts.SendEmailToProvider = true
	return cil
}

// Flexible marks the price as flexible.
func (cil *CreateInvoiceLink) Flexible() *CreateInvoiceLink {
	cil.opts.IsFlexible = true
	return cil
}

// Timeout sets a custom timeout for this request.
func (cil *CreateInvoiceLink) Timeout(duration time.Duration) *CreateInvoiceLink {
	if cil.opts.RequestOpts == nil {
		cil.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	cil.opts.RequestOpts.Timeout = duration

	return cil
}

// APIURL sets a custom API URL for this request.
func (cil *CreateInvoiceLink) APIURL(url String) *CreateInvoiceLink {
	if cil.opts.RequestOpts == nil {
		cil.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	cil.opts.RequestOpts.APIURL = url.Std()

	return cil
}

// Send creates the invoice link and returns the result.
func (cil *CreateInvoiceLink) Send() Result[String] {
	link, err := cil.ctx.Bot.Raw().CreateInvoiceLink(
		cil.title.Std(),
		cil.desc.Std(),
		cil.payload.Std(),
		cil.currency.Std(),
		cil.prices,
		cil.opts,
	)

	return ResultOf(String(link), err)
}
