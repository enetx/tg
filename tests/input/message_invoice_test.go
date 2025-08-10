package input_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/input"
)

func TestInvoice(t *testing.T) {
	title := g.String("Test Product")
	description := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")

	invoice := input.Invoice(title, description, payload, currency)
	if invoice == nil {
		t.Error("Expected MessageInvoice to be created")
	}
	if !assertMessageContent(invoice) {
		t.Error("MessageInvoice should implement MessageContent correctly")
	}
}

func TestInvoice_Price(t *testing.T) {
	title := g.String("Test Product")
	description := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")

	invoice := input.Invoice(title, description, payload, currency)
	label := g.String("Product")
	amount := int64(1000)
	result := invoice.Price(label, amount)
	if result == nil {
		t.Error("Expected Price method to return MessageInvoice")
	}
	if result != invoice {
		t.Error("Expected Price to return same MessageInvoice instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputInvoiceMessageContent); ok {
		if len(v.Prices) != 1 {
			t.Errorf("Expected 1 price, got %d", len(v.Prices))
		}
		if v.Prices[0].Label != label.Std() {
			t.Error("Expected Price label to be set correctly")
		}
		if v.Prices[0].Amount != amount {
			t.Error("Expected Price amount to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputInvoiceMessageContent")
	}
}

func TestInvoice_ProviderToken(t *testing.T) {
	title := g.String("Test Product")
	description := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")

	invoice := input.Invoice(title, description, payload, currency)
	token := g.String("test_token")
	result := invoice.ProviderToken(token)
	if result == nil {
		t.Error("Expected ProviderToken method to return MessageInvoice")
	}
	if result != invoice {
		t.Error("Expected ProviderToken to return same MessageInvoice instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputInvoiceMessageContent); ok {
		if v.ProviderToken != token.Std() {
			t.Error("Expected ProviderToken to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputInvoiceMessageContent")
	}
}

func TestInvoice_MaxTip(t *testing.T) {
	title := g.String("Test Product")
	description := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")

	invoice := input.Invoice(title, description, payload, currency)
	maxTip := int64(500)
	result := invoice.MaxTip(maxTip)
	if result == nil {
		t.Error("Expected MaxTip method to return MessageInvoice")
	}
	if result != invoice {
		t.Error("Expected MaxTip to return same MessageInvoice instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputInvoiceMessageContent); ok {
		if v.MaxTipAmount != maxTip {
			t.Error("Expected MaxTipAmount to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputInvoiceMessageContent")
	}
}

func TestInvoice_NeedName(t *testing.T) {
	title := g.String("Test Product")
	description := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")

	invoice := input.Invoice(title, description, payload, currency)
	result := invoice.NeedName()
	if result == nil {
		t.Error("Expected NeedName method to return MessageInvoice")
	}
	if result != invoice {
		t.Error("Expected NeedName to return same MessageInvoice instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputInvoiceMessageContent); ok {
		if !v.NeedName {
			t.Error("Expected NeedName to be set to true")
		}
	} else {
		t.Error("Expected result to be InputInvoiceMessageContent")
	}
}

func TestInvoice_Photo(t *testing.T) {
	title := g.String("Test Product")
	description := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")

	invoice := input.Invoice(title, description, payload, currency)
	photoURL := g.String("https://example.com/photo.jpg")
	size, width, height := int64(1024), int64(300), int64(200)
	result := invoice.Photo(photoURL, size, width, height)
	if result == nil {
		t.Error("Expected Photo method to return MessageInvoice")
	}
	if result != invoice {
		t.Error("Expected Photo to return same MessageInvoice instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputInvoiceMessageContent); ok {
		if v.PhotoUrl != photoURL.Std() {
			t.Error("Expected PhotoUrl to be set correctly")
		}
		if v.PhotoSize != size || v.PhotoWidth != width || v.PhotoHeight != height {
			t.Error("Expected Photo dimensions to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputInvoiceMessageContent")
	}
}

func TestInvoice_SuggestedTips(t *testing.T) {
	title := g.String("Test Product")
	description := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")

	invoice := input.Invoice(title, description, payload, currency)
	tips := []int64{100, 200, 500}
	result := invoice.SuggestedTips(tips...)
	if result == nil {
		t.Error("Expected SuggestedTips method to return MessageInvoice")
	}
	if result != invoice {
		t.Error("Expected SuggestedTips to return same MessageInvoice instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputInvoiceMessageContent); ok {
		if len(v.SuggestedTipAmounts) != 3 {
			t.Error("Expected SuggestedTipAmounts to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputInvoiceMessageContent")
	}
}

func TestInvoice_ProviderData(t *testing.T) {
	title := g.String("Test Product")
	description := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")

	invoice := input.Invoice(title, description, payload, currency)
	data := g.String(`{"test": "data"}`)
	result := invoice.ProviderData(data)
	if result == nil {
		t.Error("Expected ProviderData method to return MessageInvoice")
	}
	if result != invoice {
		t.Error("Expected ProviderData to return same MessageInvoice instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputInvoiceMessageContent); ok {
		if v.ProviderData != data.Std() {
			t.Error("Expected ProviderData to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputInvoiceMessageContent")
	}
}

func TestInvoice_NeedEmail(t *testing.T) {
	title := g.String("Test Product")
	description := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")

	invoice := input.Invoice(title, description, payload, currency)
	result := invoice.NeedEmail()
	if result == nil {
		t.Error("Expected NeedEmail method to return MessageInvoice")
	}
	if result != invoice {
		t.Error("Expected NeedEmail to return same MessageInvoice instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputInvoiceMessageContent); ok {
		if !v.NeedEmail {
			t.Error("Expected NeedEmail to be set to true")
		}
	} else {
		t.Error("Expected result to be InputInvoiceMessageContent")
	}
}

func TestInvoice_NeedShipping(t *testing.T) {
	title := g.String("Test Product")
	description := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")

	invoice := input.Invoice(title, description, payload, currency)
	result := invoice.NeedShipping()
	if result == nil {
		t.Error("Expected NeedShipping method to return MessageInvoice")
	}
	if result != invoice {
		t.Error("Expected NeedShipping to return same MessageInvoice instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputInvoiceMessageContent); ok {
		if !v.NeedShippingAddress {
			t.Error("Expected NeedShippingAddress to be set to true")
		}
	} else {
		t.Error("Expected result to be InputInvoiceMessageContent")
	}
}

func TestInvoice_SendPhone(t *testing.T) {
	title := g.String("Test Product")
	description := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")

	invoice := input.Invoice(title, description, payload, currency)
	result := invoice.SendPhone()
	if result == nil {
		t.Error("Expected SendPhone method to return MessageInvoice")
	}
	if result != invoice {
		t.Error("Expected SendPhone to return same MessageInvoice instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputInvoiceMessageContent); ok {
		if !v.SendPhoneNumberToProvider {
			t.Error("Expected SendPhoneNumberToProvider to be set to true")
		}
	} else {
		t.Error("Expected result to be InputInvoiceMessageContent")
	}
}

func TestInvoice_SendEmail(t *testing.T) {
	title := g.String("Test Product")
	description := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")

	invoice := input.Invoice(title, description, payload, currency)
	result := invoice.SendEmail()
	if result == nil {
		t.Error("Expected SendEmail method to return MessageInvoice")
	}
	if result != invoice {
		t.Error("Expected SendEmail to return same MessageInvoice instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputInvoiceMessageContent); ok {
		if !v.SendEmailToProvider {
			t.Error("Expected SendEmailToProvider to be set to true")
		}
	} else {
		t.Error("Expected result to be InputInvoiceMessageContent")
	}
}

func TestInvoice_MethodChaining(t *testing.T) {
	title := g.String("Test Product")
	description := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")

	result := input.Invoice(title, description, payload, currency).
		Price(g.String("Product"), 1000).
		MaxTip(500).
		NeedName().
		NeedPhone().
		Flexible()

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained Invoice to build correctly")
	}

	if _, ok := built.(gotgbot.InputInvoiceMessageContent); !ok {
		t.Error("Expected result to be InputInvoiceMessageContent")
	}

	if !assertMessageContent(result) {
		t.Error("Expected result to implement MessageContent interface")
	}
}
