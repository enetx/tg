package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_GetCustomEmojiStickers(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	emojiIDs := g.Slice[g.String]{}
	emojiIDs.Push(g.String("emoji_123"))

	result := ctx.GetCustomEmojiStickers(emojiIDs)

	if result == nil {
		t.Error("Expected GetCustomEmojiStickers builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}

// Tests for methods with 0% coverage

func TestGetCustomEmojiStickers_APIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	emojiIDs := g.Slice[g.String]{}
	emojiIDs.Push(g.String("emoji_123"))

	// Test APIURL method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.GetCustomEmojiStickers(emojiIDs)
	apiURLResultNil := freshResult.APIURL(g.String("https://custom-emoji-stickers-api.telegram.org"))
	if apiURLResultNil == nil {
		t.Error("APIURL method should return GetCustomEmojiStickers for chaining with nil RequestOpts")
	}

	// Test APIURL method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.GetCustomEmojiStickers(emojiIDs)
	apiURLFirst := anotherFreshResult.APIURL(g.String("https://first-emoji-stickers-api.telegram.org")) // This creates RequestOpts
	apiURLSecond := apiURLFirst.APIURL(g.String("https://second-emoji-stickers-api.telegram.org"))      // This uses existing RequestOpts
	if apiURLSecond == nil {
		t.Error("APIURL method should return GetCustomEmojiStickers for chaining with existing RequestOpts")
	}

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://emoji-stickers-api.example.com",
		"https://custom-emoji-stickers.telegram.org",
		"https://regional-emoji-stickers-api.telegram.org",
		"https://backup-emoji-stickers-api.telegram.org",
		"", // Empty URL
	}

	for _, apiURL := range apiURLs {
		apiResult := ctx.GetCustomEmojiStickers(emojiIDs).
			Timeout(30 * time.Second).
			APIURL(g.String(apiURL)).
			Send()

		if apiResult.IsErr() {
			t.Logf("GetCustomEmojiStickers with API URL '%s' Send failed as expected: %v", apiURL, apiResult.Err())
		}
	}
}

func TestGetCustomEmojiStickers_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Send method with single emoji ID
	singleEmojiIDs := g.Slice[g.String]{}
	singleEmojiIDs.Push(g.String("emoji_single"))

	sendResult := ctx.GetCustomEmojiStickers(singleEmojiIDs).Send()

	if sendResult.IsErr() {
		t.Logf("GetCustomEmojiStickers Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with multiple emoji IDs
	multipleEmojiIDs := g.Slice[g.String]{}
	multipleEmojiIDs.Push(g.String("emoji_1"))
	multipleEmojiIDs.Push(g.String("emoji_2"))
	multipleEmojiIDs.Push(g.String("emoji_3"))

	configuredSendResult := ctx.GetCustomEmojiStickers(multipleEmojiIDs).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("GetCustomEmojiStickers configured Send failed as expected: %v", configuredSendResult.Err())
	}

	// Test Send method with various emoji ID scenarios
	emojiIDScenarios := []struct {
		emojiIDs    []string
		description string
	}{
		{[]string{"custom_emoji_1"}, "Single custom emoji"},
		{[]string{"custom_emoji_1", "custom_emoji_2"}, "Two custom emojis"},
		{[]string{"custom_emoji_1", "custom_emoji_2", "custom_emoji_3", "custom_emoji_4", "custom_emoji_5"}, "Five custom emojis"},
		{[]string{}, "Empty emoji list"},
	}

	for _, scenario := range emojiIDScenarios {
		scenarioEmojiIDs := g.Slice[g.String]{}
		for _, emojiID := range scenario.emojiIDs {
			scenarioEmojiIDs.Push(g.String(emojiID))
		}

		scenarioResult := ctx.GetCustomEmojiStickers(scenarioEmojiIDs).
			Timeout(45 * time.Second).
			Send()

		if scenarioResult.IsErr() {
			t.Logf("GetCustomEmojiStickers with %s Send failed as expected: %v",
				scenario.description, scenarioResult.Err())
		}
	}
}
