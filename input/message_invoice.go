package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// MessageInvoice represents an input invoice message content builder.
type MessageInvoice struct {
	input *gotgbot.InputInvoiceMessageContent
}

// NewMessageInvoice creates a new MessageInvoice builder with the required fields.
func NewMessageInvoice(title, description, payload, currency String, prices []gotgbot.LabeledPrice) *MessageInvoice {
	return &MessageInvoice{
		input: &gotgbot.InputInvoiceMessageContent{
			Title:       title.Std(),
			Description: description.Std(),
			Payload:     payload.Std(),
			Currency:    currency.Std(),
			Prices:      prices,
		},
	}
}

// ProviderToken sets the payment provider token.
func (mi *MessageInvoice) ProviderToken(token String) *MessageInvoice {
	mi.input.ProviderToken = token.Std()
	return mi
}

// MaxTipAmount sets the maximum accepted amount for tips in the smallest currency unit.
func (mi *MessageInvoice) MaxTipAmount(amount int64) *MessageInvoice {
	mi.input.MaxTipAmount = amount
	return mi
}

// SuggestedTipAmounts sets suggested amounts of tip in the smallest currency unit.
func (mi *MessageInvoice) SuggestedTipAmounts(amounts []int64) *MessageInvoice {
	mi.input.SuggestedTipAmounts = amounts
	return mi
}

// ProviderData sets JSON-encoded data about the invoice.
func (mi *MessageInvoice) ProviderData(data String) *MessageInvoice {
	mi.input.ProviderData = data.Std()
	return mi
}

// PhotoURL sets the URL of the product photo for the invoice.
func (mi *MessageInvoice) PhotoURL(url String) *MessageInvoice {
	mi.input.PhotoUrl = url.Std()
	return mi
}

// PhotoSize sets the photo size in bytes.
func (mi *MessageInvoice) PhotoSize(size int64) *MessageInvoice {
	mi.input.PhotoSize = size
	return mi
}

// PhotoWidth sets the photo width.
func (mi *MessageInvoice) PhotoWidth(width int64) *MessageInvoice {
	mi.input.PhotoWidth = width
	return mi
}

// PhotoHeight sets the photo height.
func (mi *MessageInvoice) PhotoHeight(height int64) *MessageInvoice {
	mi.input.PhotoHeight = height
	return mi
}

// NeedName requests the user's full name.
func (mi *MessageInvoice) NeedName() *MessageInvoice {
	mi.input.NeedName = true
	return mi
}

// NeedPhoneNumber requests the user's phone number.
func (mi *MessageInvoice) NeedPhoneNumber() *MessageInvoice {
	mi.input.NeedPhoneNumber = true
	return mi
}

// NeedEmail requests the user's email address.
func (mi *MessageInvoice) NeedEmail() *MessageInvoice {
	mi.input.NeedEmail = true
	return mi
}

// NeedShippingAddress requests the user's shipping address.
func (mi *MessageInvoice) NeedShippingAddress() *MessageInvoice {
	mi.input.NeedShippingAddress = true
	return mi
}

// SendPhoneNumberToProvider sends the user's phone number to the provider.
func (mi *MessageInvoice) SendPhoneNumberToProvider() *MessageInvoice {
	mi.input.SendPhoneNumberToProvider = true
	return mi
}

// SendEmailToProvider sends the user's email to the provider.
func (mi *MessageInvoice) SendEmailToProvider() *MessageInvoice {
	mi.input.SendEmailToProvider = true
	return mi
}

// IsFlexible makes prices flexible.
func (mi *MessageInvoice) IsFlexible() *MessageInvoice {
	mi.input.IsFlexible = true
	return mi
}

// Build creates the gotgbot.InputInvoiceMessageContent.
func (mi *MessageInvoice) Build() gotgbot.InputMessageContent {
	return *mi.input
}