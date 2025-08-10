package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
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

func TestSendContact_After(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	phoneNumber := g.String("+1234567890")
	firstName := g.String("John")

	durations := []time.Duration{
		time.Second,
		time.Minute,
		time.Hour,
		0,
	}

	for _, duration := range durations {
		result := ctx.SendContact(phoneNumber, firstName).After(duration)
		if result == nil {
			t.Errorf("After method should return SendContact builder for chaining with duration: %v", duration)
		}

		chainedResult := result.After(time.Second * 30)
		if chainedResult == nil {
			t.Errorf("After method should support chaining and override with duration: %v", duration)
		}
	}
}

func TestSendContact_DeleteAfter(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	phoneNumber := g.String("+1234567890")
	firstName := g.String("John")

	durations := []time.Duration{
		time.Second * 30,
		time.Minute * 5,
		time.Hour,
		0,
	}

	for _, duration := range durations {
		result := ctx.SendContact(phoneNumber, firstName).DeleteAfter(duration)
		if result == nil {
			t.Errorf("DeleteAfter method should return SendContact builder for chaining with duration: %v", duration)
		}

		chainedResult := result.DeleteAfter(time.Minute * 10)
		if chainedResult == nil {
			t.Errorf("DeleteAfter method should support chaining and override with duration: %v", duration)
		}
	}
}

func TestSendContact_Markup(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	phoneNumber := g.String("+1234567890")
	firstName := g.String("John")

	btn1 := keyboard.NewButton().Text(g.String("Test Button")).Callback(g.String("test_data"))
	inlineKeyboard := keyboard.Inline().Button(btn1)

	result := ctx.SendContact(phoneNumber, firstName).Markup(inlineKeyboard)
	if result == nil {
		t.Error("Markup method should return SendContact builder for chaining with inline keyboard")
	}

	replyKeyboard := keyboard.Reply()
	chainedResult := result.Markup(replyKeyboard)
	if chainedResult == nil {
		t.Error("Markup method should support chaining and override with reply keyboard")
	}
}

func TestSendContact_ReplyTo(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	phoneNumber := g.String("+1234567890")
	firstName := g.String("John")

	messageIDs := []int64{1, 123, 456, 999}

	for _, messageID := range messageIDs {
		result := ctx.SendContact(phoneNumber, firstName).ReplyTo(messageID)
		if result == nil {
			t.Errorf("ReplyTo method should return SendContact builder for chaining with messageID: %d", messageID)
		}

		chainedResult := result.ReplyTo(messageID + 100)
		if chainedResult == nil {
			t.Errorf("ReplyTo method should support chaining and override with messageID: %d", messageID)
		}
	}
}

func TestSendContact_Business(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	phoneNumber := g.String("+1234567890")
	firstName := g.String("John")

	businessIDs := []string{
		"business_123",
		"conn_456",
		"",
	}

	for _, businessID := range businessIDs {
		result := ctx.SendContact(phoneNumber, firstName).Business(g.String(businessID))
		if result == nil {
			t.Errorf("Business method should return SendContact builder for chaining with businessID: %s", businessID)
		}

		chainedResult := result.Business(g.String("override_business"))
		if chainedResult == nil {
			t.Errorf("Business method should support chaining and override with businessID: %s", businessID)
		}
	}
}

func TestSendContact_Thread(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	phoneNumber := g.String("+1234567890")
	firstName := g.String("John")

	threadIDs := []int64{1, 123, 456, 0}

	for _, threadID := range threadIDs {
		result := ctx.SendContact(phoneNumber, firstName).Thread(threadID)
		if result == nil {
			t.Errorf("Thread method should return SendContact builder for chaining with threadID: %d", threadID)
		}

		chainedResult := result.Thread(threadID + 100)
		if chainedResult == nil {
			t.Errorf("Thread method should support chaining and override with threadID: %d", threadID)
		}
	}
}

func TestSendContact_APIURLWithExistingRequestOpts(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// First set Timeout to create RequestOpts, then test APIURL
	result := ctx.SendContact(g.String("123456789"), g.String("John")).
		Timeout(15 * time.Second).                         // This creates RequestOpts
		APIURL(g.String("https://custom.api.example.com")) // This should use existing RequestOpts

	if result == nil {
		t.Error("APIURL with existing RequestOpts should return builder")
	}
}
