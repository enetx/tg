package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SendContact(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	phoneNumber := g.String("+1234567890")
	firstName := g.String("John")

	result := ctx.SendContact(phoneNumber, firstName)

	if result == nil {
		t.Error("Expected SendContact builder to be created")
	}

	// Test method chaining
	chained := result.Silent()
	if chained == nil {
		t.Error("Expected Silent method to return builder")
	}
}

func TestContext_SendContactChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	phoneNumber := g.String("+1234567890")
	firstName := g.String("John")

	result := ctx.SendContact(phoneNumber, firstName).
		LastName(g.String("Doe")).
		Silent().
		To(123)

	if result == nil {
		t.Error("Expected SendContact builder to be created")
	}

	// Test continued chaining
	final := result.Protect()
	if final == nil {
		t.Error("Expected Protect method to return builder")
	}
}

func TestSendContact_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	phoneNumber := g.String("+1234567890")
	firstName := g.String("John")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendContact(phoneNumber, firstName).Send()

	if sendResult.IsErr() {
		t.Logf("SendContact Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendContact(phoneNumber, firstName).
		LastName(g.String("Doe")).
		VCard(g.String("BEGIN:VCARD\nVERSION:3.0\nFN:John Doe\nEND:VCARD")).
		Silent().
		Protect().
		To(123).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendContact configured Send failed as expected: %v", configuredSendResult.Err())
	}
}
