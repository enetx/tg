package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_RemoveChatVerification(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	chatID := int64(-1001234567890)

	result := ctx.RemoveChatVerification(chatID)

	if result == nil {
		t.Error("Expected RemoveChatVerification builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}

func TestRemoveChatVerification_APIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	chatID := int64(-1001234567890)

	apiURLs := []string{
		"https://api.telegram.org",
		"https://custom.api.example.com",
		"",
		"https://api.example.com/bot",
		"http://localhost:8080",
	}

	for _, apiURL := range apiURLs {
		result := ctx.RemoveChatVerification(chatID)
		apiURLResult := result.APIURL(g.String(apiURL))
		if apiURLResult == nil {
			t.Errorf("APIURL method should return RemoveChatVerification builder for chaining with URL: %s", apiURL)
		}

		chainedResult := apiURLResult.APIURL(g.String("https://override.example.com"))
		if chainedResult == nil {
			t.Errorf("APIURL method should support chaining and override with URL: %s", apiURL)
		}
	}
}

func TestRemoveChatVerification_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	chatID := int64(-1001234567890)

	sendResult := ctx.RemoveChatVerification(chatID).Send()

	if sendResult.IsErr() {
		t.Logf("RemoveChatVerification Send failed as expected with mock bot: %v", sendResult.Err())
	}

	sendWithOptionsResult := ctx.RemoveChatVerification(chatID).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.telegram.org")).
		Send()

	if sendWithOptionsResult.IsErr() {
		t.Logf("RemoveChatVerification Send with options failed as expected with mock bot: %v", sendWithOptionsResult.Err())
	}
}
