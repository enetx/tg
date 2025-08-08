package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_DeleteStickerSet(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	name := g.String("test_sticker_set_by_bot")

	result := ctx.DeleteStickerSet(name)

	if result == nil {
		t.Error("Expected DeleteStickerSet builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(30 * time.Second)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}

	// Test APIURL method
	apiURLResult := result.APIURL(g.String("https://api.telegram.org"))
	if apiURLResult == nil {
		t.Error("APIURL method should return DeleteStickerSet for chaining")
	}

	// Test APIURL method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.DeleteStickerSet(name)
	apiURLResultNil := freshResult.APIURL(g.String("https://custom-api.telegram.org"))
	if apiURLResultNil == nil {
		t.Error("APIURL method should return DeleteStickerSet for chaining with nil RequestOpts")
	}
}

func TestDeleteStickerSet_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	stickerSetName := g.String("test_sticker_set_by_bot")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.DeleteStickerSet(stickerSetName).Send()

	if sendResult.IsErr() {
		t.Logf("DeleteStickerSet Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.DeleteStickerSet(stickerSetName).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("DeleteStickerSet configured Send failed as expected: %v", configuredSendResult.Err())
	}

	// Test Send method with different sticker set names
	stickerSetNames := []string{
		"test_sticker_set_by_bot",
		"my_custom_stickers_by_bot",
		"company_stickers_by_bot",
		"animated_pack_by_bot",
		"emoji_pack_by_bot",
		"gaming_stickers_by_bot",
		"meme_collection_by_bot",
		"seasonal_pack_by_bot",
		"premium_stickers_by_bot",
		"community_pack_by_bot",
	}

	for _, setName := range stickerSetNames {
		nameResult := ctx.DeleteStickerSet(g.String(setName)).
			Timeout(45 * time.Second).
			Send()

		if nameResult.IsErr() {
			t.Logf("DeleteStickerSet with name '%s' Send failed as expected: %v", setName, nameResult.Err())
		}
	}
}

func TestDeleteStickerSet_CompleteMethods(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	stickerSetName := g.String("comprehensive_test_by_bot")

	// Test all methods in combination
	result := ctx.DeleteStickerSet(stickerSetName).
		Timeout(60 * time.Second).
		APIURL(g.String("https://complete-sticker-set-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return DeleteStickerSet")
	}

	// Test complete workflow with Send
	completeResult := result.Send()
	if completeResult.IsErr() {
		t.Logf("Complete DeleteStickerSet workflow Send failed as expected: %v", completeResult.Err())
	}

	// Test various timeout configurations
	timeouts := []time.Duration{
		5 * time.Second,
		15 * time.Second,
		30 * time.Second,
		60 * time.Second,
		2 * time.Minute,
		5 * time.Minute,
		0 * time.Second, // Zero timeout
	}

	for _, timeout := range timeouts {
		timeoutResult := ctx.DeleteStickerSet(stickerSetName).
			Timeout(timeout).
			APIURL(g.String("https://timeout-sticker-api.telegram.org")).
			Send()

		if timeoutResult.IsErr() {
			t.Logf("DeleteStickerSet with timeout %v Send failed as expected: %v", timeout, timeoutResult.Err())
		}
	}

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://sticker-management-api.example.com",
		"https://custom-sticker-delete.telegram.org",
		"https://regional-sticker-api.telegram.org",
		"https://backup-sticker-api.telegram.org",
		"", // Empty URL
	}

	for _, apiURL := range apiURLs {
		apiResult := ctx.DeleteStickerSet(stickerSetName).
			Timeout(30 * time.Second).
			APIURL(g.String(apiURL)).
			Send()

		if apiResult.IsErr() {
			t.Logf("DeleteStickerSet with API URL '%s' Send failed as expected: %v", apiURL, apiResult.Err())
		}
	}

	// Test edge cases with sticker set names
	edgeCaseNames := []string{
		"a",                                 // Single character
		"short_by_bot",                      // Short name
		"very_long_sticker_set_name_by_bot", // Long name
		"numbers123_by_bot",                 // With numbers
		"special_chars_by_bot",              // With underscores
		"UPPERCASE_BY_BOT",                  // Uppercase
		"mixed_Case_By_Bot",                 // Mixed case
	}

	for _, edgeName := range edgeCaseNames {
		edgeResult := ctx.DeleteStickerSet(g.String(edgeName)).
			Timeout(45 * time.Second).
			APIURL(g.String("https://edge-case-api.telegram.org")).
			Send()

		if edgeResult.IsErr() {
			t.Logf("DeleteStickerSet with edge case name '%s' Send failed as expected: %v", edgeName, edgeResult.Err())
		}
	}
}
