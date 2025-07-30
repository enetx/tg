package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// MessageInvoice represents an input invoice message content builder.
type MessageInvoice struct {
	input *gotgbot.InputInvoiceMessageContent
}

// Invoice creates a new MessageInvoice builder with the required fields.
func Invoice(title, description, payload, currency g.String) *MessageInvoice {
	return &MessageInvoice{
		input: &gotgbot.InputInvoiceMessageContent{
			Title:       title.Std(),
			Description: description.Std(),
			Payload:     payload.Std(),
			Currency:    currency.Std(),
		},
	}
}

// Price adds a labeled price item to the invoice.
func (mi *MessageInvoice) Price(label g.String, amount int64) *MessageInvoice {
	mi.input.Prices = append(mi.input.Prices, gotgbot.LabeledPrice{Label: label.Std(), Amount: amount})
	return mi
}

// ProviderToken sets the payment provider token.
func (mi *MessageInvoice) ProviderToken(token g.String) *MessageInvoice {
	mi.input.ProviderToken = token.Std()
	return mi
}

// MaxTip sets the maximum accepted amount for tips in the smallest currency unit.
func (mi *MessageInvoice) MaxTip(amount int64) *MessageInvoice {
	mi.input.MaxTipAmount = amount
	return mi
}

// SuggestedTips sets suggested amounts of tip in the smallest currency unit.
func (mi *MessageInvoice) SuggestedTips(tips ...int64) *MessageInvoice {
	mi.input.SuggestedTipAmounts = tips
	return mi
}

// ProviderData sets JSON-encoded data about the invoice.
func (mi *MessageInvoice) ProviderData(data g.String) *MessageInvoice {
	mi.input.ProviderData = data.Std()
	return mi
}

// Photo sets the product photo URL and dimensions.
func (mi *MessageInvoice) Photo(url g.String, size, width, height int64) *MessageInvoice {
	mi.input.PhotoUrl = url.Std()
	mi.input.PhotoSize = size
	mi.input.PhotoWidth = width
	mi.input.PhotoHeight = height

	return mi
}

// NeedName requests the user's full name.
func (mi *MessageInvoice) NeedName() *MessageInvoice {
	mi.input.NeedName = true
	return mi
}

// NeedPhone requests the user's phone number.
func (mi *MessageInvoice) NeedPhone() *MessageInvoice {
	mi.input.NeedPhoneNumber = true
	return mi
}

// NeedEmail requests the user's email address.
func (mi *MessageInvoice) NeedEmail() *MessageInvoice {
	mi.input.NeedEmail = true
	return mi
}

// NeedShipping requests the user's shipping address.
func (mi *MessageInvoice) NeedShipping() *MessageInvoice {
	mi.input.NeedShippingAddress = true
	return mi
}

// SendPhone sends the user's phone number to the provider.
func (mi *MessageInvoice) SendPhone() *MessageInvoice {
	mi.input.SendPhoneNumberToProvider = true
	return mi
}

// SendEmail sends the user's email to the provider.
func (mi *MessageInvoice) SendEmail() *MessageInvoice {
	mi.input.SendEmailToProvider = true
	return mi
}

// Flexible makes prices flexible.
func (mi *MessageInvoice) Flexible() *MessageInvoice {
	mi.input.IsFlexible = true
	return mi
}

// Build creates the gotgbot.InputInvoiceMessageContent.
func (mi *MessageInvoice) Build() gotgbot.InputMessageContent {
	return *mi.input
}
