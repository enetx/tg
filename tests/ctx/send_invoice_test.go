package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/types/effects"
)

func TestContext_SendInvoice(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	title := g.String("Test Product")
	desc := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")

	result := ctx.SendInvoice(title, desc, payload, currency)

	if result == nil {
		t.Error("Expected SendInvoice builder to be created")
	}

	// Test method chaining
	chained := result.Silent()
	if chained == nil {
		t.Error("Expected Silent method to return builder")
	}
}

func TestContext_SendInvoiceChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	title := g.String("Test Product")
	desc := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")

	result := ctx.SendInvoice(title, desc, payload, currency).
		StartParameter(g.String("start_param")).
		Silent().
		To(123)

	if result == nil {
		t.Error("Expected SendInvoice builder to be created")
	}

	// Test continued chaining
	final := result.Protect()
	if final == nil {
		t.Error("Expected Protect method to return builder")
	}
}

func TestSendInvoice_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	title := g.String("Test Product")
	desc := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendInvoice(title, desc, payload, currency).Send()

	if sendResult.IsErr() {
		t.Logf("SendInvoice Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendInvoice(title, desc, payload, currency).
		StartParameter(g.String("start_param")).
		ProviderToken(g.String("test_token")).
		Silent().
		Protect().
		To(123).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendInvoice configured Send failed as expected: %v", configuredSendResult.Err())
	}
}

func TestSendInvoice_Price(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	title := g.String("Test Product")
	desc := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")
	if ctx.SendInvoice(title, desc, payload, currency).Price(g.String("Item 1"), 1000) == nil {
		t.Error("Price should return builder")
	}
}

func TestSendInvoice_Thread(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	title := g.String("Test Product")
	desc := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")
	if ctx.SendInvoice(title, desc, payload, currency).Thread(123) == nil {
		t.Error("Thread should return builder")
	}
}

func TestSendInvoice_MaxTip(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	title := g.String("Test Product")
	desc := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")
	if ctx.SendInvoice(title, desc, payload, currency).MaxTip(5000) == nil {
		t.Error("MaxTip should return builder")
	}
}

func TestSendInvoice_SuggestedTips(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	title := g.String("Test Product")
	desc := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")
	if ctx.SendInvoice(title, desc, payload, currency).SuggestedTips(100, 500, 1000, 2000) == nil {
		t.Error("SuggestedTips should return builder")
	}
}

func TestSendInvoice_ProviderData(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	title := g.String("Test Product")
	desc := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")
	if ctx.SendInvoice(title, desc, payload, currency).ProviderData(g.String("{\"key\": \"value\"}")) == nil {
		t.Error("ProviderData should return builder")
	}
}

func TestSendInvoice_Photo(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	title := g.String("Test Product")
	desc := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")
	if ctx.SendInvoice(title, desc, payload, currency).
		Photo(g.String("https://example.com/photo.jpg"), 1024, 300, 200) ==
		nil {
		t.Error("Photo should return builder")
	}
}

func TestSendInvoice_NeedName(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	title := g.String("Test Product")
	desc := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")
	if ctx.SendInvoice(title, desc, payload, currency).NeedName() == nil {
		t.Error("NeedName should return builder")
	}
}

func TestSendInvoice_NeedPhone(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	title := g.String("Test Product")
	desc := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")
	if ctx.SendInvoice(title, desc, payload, currency).NeedPhone() == nil {
		t.Error("NeedPhone should return builder")
	}
}

func TestSendInvoice_NeedEmail(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	title := g.String("Test Product")
	desc := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")
	if ctx.SendInvoice(title, desc, payload, currency).NeedEmail() == nil {
		t.Error("NeedEmail should return builder")
	}
}

func TestSendInvoice_NeedShipping(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	title := g.String("Test Product")
	desc := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")
	if ctx.SendInvoice(title, desc, payload, currency).NeedShipping() == nil {
		t.Error("NeedShipping should return builder")
	}
}

func TestSendInvoice_SendPhone(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	title := g.String("Test Product")
	desc := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")
	if ctx.SendInvoice(title, desc, payload, currency).SendPhone() == nil {
		t.Error("SendPhone should return builder")
	}
}

func TestSendInvoice_SendEmail(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	title := g.String("Test Product")
	desc := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")
	if ctx.SendInvoice(title, desc, payload, currency).SendEmail() == nil {
		t.Error("SendEmail should return builder")
	}
}

func TestSendInvoice_Flexible(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	title := g.String("Test Product")
	desc := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")
	if ctx.SendInvoice(title, desc, payload, currency).Flexible() == nil {
		t.Error("Flexible should return builder")
	}
}

func TestSendInvoice_AllowPaidBroadcast(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	title := g.String("Test Product")
	desc := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")
	if ctx.SendInvoice(title, desc, payload, currency).AllowPaidBroadcast() == nil {
		t.Error("AllowPaidBroadcast should return builder")
	}
}

func TestSendInvoice_Effect(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	title := g.String("Test Product")
	desc := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")
	if ctx.SendInvoice(title, desc, payload, currency).Effect(effects.Fire) == nil {
		t.Error("Effect should return builder")
	}
}

func TestSendInvoice_ReplyTo(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	title := g.String("Test Product")
	desc := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")
	if ctx.SendInvoice(title, desc, payload, currency).ReplyTo(123) == nil {
		t.Error("ReplyTo should return builder")
	}
}

func TestSendInvoice_Markup(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	title := g.String("Test Product")
	desc := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")
	btn1 := keyboard.NewButton().Text(g.String("Pay Now")).Pay()
	if ctx.SendInvoice(title, desc, payload, currency).Markup(keyboard.Inline().Button(btn1)) == nil {
		t.Error("Markup should return builder")
	}
}
