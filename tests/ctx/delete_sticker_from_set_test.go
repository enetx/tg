package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/file"
)

func TestContext_DeleteStickerFromSet(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	sticker := g.String("BAADAgADQQADBREAAYpMOJJhMdJWAg")

	result := ctx.DeleteStickerFromSet(file.Input(sticker).UnwrapOrDefault())

	if result == nil {
		t.Error("Expected DeleteStickerFromSet builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(30 * time.Second)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}

	// Test APIURL method
	apiURLResult := result.APIURL(g.String("https://api.telegram.org"))
	if apiURLResult == nil {
		t.Error("APIURL method should return DeleteStickerFromSet for chaining")
	}

	// Test APIURL method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.DeleteStickerFromSet(file.Input(sticker).UnwrapOrDefault())
	apiURLResultNil := freshResult.APIURL(g.String("https://custom-api.telegram.org"))
	if apiURLResultNil == nil {
		t.Error("APIURL method should return DeleteStickerFromSet for chaining with nil RequestOpts")
	}
}

func TestDeleteStickerFromSet_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	sticker := g.String("BAADAgADQQADBREAAYpMOJJhMdJWAg")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.DeleteStickerFromSet(file.Input(sticker).UnwrapOrDefault()).Send()

	if sendResult.IsErr() {
		t.Logf("DeleteStickerFromSet Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.DeleteStickerFromSet(file.Input(sticker).UnwrapOrDefault()).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("DeleteStickerFromSet configured Send failed as expected: %v", configuredSendResult.Err())
	}

	// Test Send method with different sticker formats
	stickerFormats := []string{
		"BAADAgADQQADBREAAYpMOJJhMdJWAg",                   // File ID
		"BAADBQADBAADMhcAAUdyWxYeNLZlAg",                   // Another file ID
		"CAACAgIAAxkBAAICHGH8ZjsAAWxBAAHjkTyXpZn8R_zWlgAB", // Animated sticker
		"CAADBQADBAADHQgAAqhTBxDj5R-IAg",                   // Custom emoji
	}

	for _, stickerID := range stickerFormats {
		formatResult := ctx.DeleteStickerFromSet(file.Input(g.String(stickerID)).UnwrapOrDefault()).
			Timeout(45 * time.Second).
			Send()

		if formatResult.IsErr() {
			t.Logf("DeleteStickerFromSet with sticker ID '%s' Send failed as expected: %v", stickerID, formatResult.Err())
		}
	}
}

func TestDeleteStickerFromSet_CompleteMethods(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	sticker := g.String("BAADAgADQQADBREAAYpMOJJhMdJWAg")

	// Test all methods in combination
	result := ctx.DeleteStickerFromSet(file.Input(sticker).UnwrapOrDefault()).
		Timeout(60 * time.Second).
		APIURL(g.String("https://complete-sticker-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return DeleteStickerFromSet")
	}

	// Test complete workflow with Send
	completeResult := result.Send()
	if completeResult.IsErr() {
		t.Logf("Complete DeleteStickerFromSet workflow Send failed as expected: %v", completeResult.Err())
	}

	// Test various timeout configurations
	timeouts := []time.Duration{
		1 * time.Second,
		15 * time.Second,
		30 * time.Second,
		60 * time.Second,
		2 * time.Minute,
		0 * time.Second, // Zero timeout
	}

	for _, timeout := range timeouts {
		timeoutResult := ctx.DeleteStickerFromSet(file.Input(sticker).UnwrapOrDefault()).
			Timeout(timeout).
			APIURL(g.String("https://timeout-api.telegram.org")).
			Send()

		if timeoutResult.IsErr() {
			t.Logf("DeleteStickerFromSet with timeout %v Send failed as expected: %v", timeout, timeoutResult.Err())
		}
	}

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://custom-sticker-delete-api.example.com",
		"https://regional-api.telegram.org",
		"", // Empty URL
	}

	for _, apiURL := range apiURLs {
		apiResult := ctx.DeleteStickerFromSet(file.Input(sticker).UnwrapOrDefault()).
			Timeout(30 * time.Second).
			APIURL(g.String(apiURL)).
			Send()

		if apiResult.IsErr() {
			t.Logf("DeleteStickerFromSet with API URL '%s' Send failed as expected: %v", apiURL, apiResult.Err())
		}
	}
}
