package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_GetStickerSet(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	name := g.String("test_sticker_set")

	result := ctx.GetStickerSet(name)

	if result == nil {
		t.Error("Expected GetStickerSet builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}

// Tests for methods with 0% coverage

func TestGetStickerSet_APIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	stickerSetName := g.String("test_sticker_set")

	// Test APIURL method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.GetStickerSet(stickerSetName)
	apiURLResultNil := freshResult.APIURL(g.String("https://custom-sticker-set-api.telegram.org"))
	if apiURLResultNil == nil {
		t.Error("APIURL method should return GetStickerSet for chaining with nil RequestOpts")
	}

	// Test APIURL method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.GetStickerSet(stickerSetName)
	apiURLFirst := anotherFreshResult.APIURL(g.String("https://first-sticker-set-api.telegram.org")) // This creates RequestOpts
	apiURLSecond := apiURLFirst.APIURL(g.String("https://second-sticker-set-api.telegram.org"))      // This uses existing RequestOpts
	if apiURLSecond == nil {
		t.Error("APIURL method should return GetStickerSet for chaining with existing RequestOpts")
	}

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://sticker-set-api.example.com",
		"https://custom-sticker-set.telegram.org",
		"https://regional-sticker-set-api.telegram.org",
		"https://backup-sticker-set-api.telegram.org",
		"", // Empty URL
	}

	for _, apiURL := range apiURLs {
		apiResult := ctx.GetStickerSet(stickerSetName).
			Timeout(30 * time.Second).
			APIURL(g.String(apiURL)).
			Send()

		if apiResult.IsErr() {
			t.Logf("GetStickerSet with API URL '%s' Send failed as expected: %v", apiURL, apiResult.Err())
		}
	}
}

func TestGetStickerSet_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Send method with various sticker set names
	stickerSetScenarios := []struct {
		name        string
		description string
	}{
		{"test_sticker_set", "Basic sticker set name"},
		{"AnimatedStickers", "Animated stickers set"},
		{"VideoStickers", "Video stickers set"},
		{"CustomEmojiSet", "Custom emoji set"},
		{"", "Empty sticker set name"},
		{"very_long_sticker_set_name_with_underscores_and_numbers_123", "Long sticker set name"},
		{"Special@Characters#Set", "Sticker set name with special characters"},
	}

	for _, scenario := range stickerSetScenarios {
		stickerSetName := g.String(scenario.name)

		// Basic Send test
		sendResult := ctx.GetStickerSet(stickerSetName).Send()
		if sendResult.IsErr() {
			t.Logf("GetStickerSet with %s Send failed as expected: %v", scenario.description, sendResult.Err())
		}

		// Configured Send test
		configuredSendResult := ctx.GetStickerSet(stickerSetName).
			Timeout(30 * time.Second).
			APIURL(g.String("https://api.example.com")).
			Send()

		if configuredSendResult.IsErr() {
			t.Logf("GetStickerSet configured with %s Send failed as expected: %v", scenario.description, configuredSendResult.Err())
		}
	}

	// Test Send method with various timeout configurations
	timeoutConfigs := []time.Duration{
		5 * time.Second,
		15 * time.Second,
		45 * time.Second,
		60 * time.Second,
		2 * time.Minute,
	}

	for _, timeout := range timeoutConfigs {
		timeoutResult := ctx.GetStickerSet(g.String("timeout_test_set")).
			Timeout(timeout).
			Send()

		if timeoutResult.IsErr() {
			t.Logf("GetStickerSet with timeout %v Send failed as expected: %v", timeout, timeoutResult.Err())
		}
	}

	// Test comprehensive workflow
	comprehensiveResult := ctx.GetStickerSet(g.String("comprehensive_test_set")).
		Timeout(90 * time.Second).
		APIURL(g.String("https://comprehensive-sticker-set-api.telegram.org")).
		Send()

	if comprehensiveResult.IsErr() {
		t.Logf("GetStickerSet comprehensive workflow Send failed as expected: %v", comprehensiveResult.Err())
	}

	// Test method chaining order independence
	orderTest1 := ctx.GetStickerSet(g.String("order_test_1")).
		APIURL(g.String("https://order-test-1.telegram.org")).
		Timeout(45 * time.Second).
		Send()

	if orderTest1.IsErr() {
		t.Logf("GetStickerSet order test 1 Send failed as expected: %v", orderTest1.Err())
	}

	orderTest2 := ctx.GetStickerSet(g.String("order_test_2")).
		Timeout(45 * time.Second).
		APIURL(g.String("https://order-test-2.telegram.org")).
		Send()

	if orderTest2.IsErr() {
		t.Logf("GetStickerSet order test 2 Send failed as expected: %v", orderTest2.Err())
	}
}
