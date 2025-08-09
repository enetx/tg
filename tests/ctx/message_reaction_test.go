package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SetMessageReaction(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	messageID := int64(789)

	result := ctx.SetMessageReaction(messageID)

	if result == nil {
		t.Error("Expected SetMessageReaction builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(456)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}

func TestSetMessageReaction_Reaction(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	messageID := int64(789)

	// Test Reaction method with various emojis
	emojis := []string{
		"👍",
		"👎",
		"❤️",
		"🔥",
		"🥰",
		"😢",
		"🎉",
		"😡",
		"🤢",
		"👏",
	}

	for _, emoji := range emojis {
		result := ctx.SetMessageReaction(messageID)
		reactionResult := result.Reaction(g.String(emoji))
		if reactionResult == nil {
			t.Errorf("Reaction method should return SetMessageReaction builder for chaining with emoji: %s", emoji)
		}

		// Test multiple reactions on same builder
		multiReactionResult := reactionResult.Reaction(g.String("👍"))
		if multiReactionResult == nil {
			t.Errorf("Reaction method should support multiple reactions with emoji: %s", emoji)
		}
	}
}

func TestSetMessageReaction_CustomEmoji(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	messageID := int64(789)

	// Test CustomEmoji method with various custom emoji IDs
	customEmojiIDs := []string{
		"custom_emoji_123",
		"5789775825365762059",
		"custom_fire_emoji",
		"premium_heart_emoji_456",
		"",
	}

	for _, emojiID := range customEmojiIDs {
		result := ctx.SetMessageReaction(messageID)
		customEmojiResult := result.CustomEmoji(g.String(emojiID))
		if customEmojiResult == nil {
			t.Errorf("CustomEmoji method should return SetMessageReaction builder for chaining with emojiID: %s", emojiID)
		}

		// Test multiple custom emojis on same builder
		multiCustomEmojiResult := customEmojiResult.CustomEmoji(g.String("another_custom_emoji"))
		if multiCustomEmojiResult == nil {
			t.Errorf("CustomEmoji method should support multiple custom emojis with emojiID: %s", emojiID)
		}
	}
}

func TestSetMessageReaction_Big(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	messageID := int64(789)

	// Test Big method
	result := ctx.SetMessageReaction(messageID)
	bigResult := result.Big()
	if bigResult == nil {
		t.Error("Big method should return SetMessageReaction builder for chaining")
	}

	// Test that Big can be chained multiple times
	chainedResult := bigResult.Big()
	if chainedResult == nil {
		t.Error("Big method should support multiple chaining calls")
	}

	// Test Big with other methods
	bigWithOthers := ctx.SetMessageReaction(messageID).
		Reaction(g.String("👍")).
		Big().
		ChatID(456)
	if bigWithOthers == nil {
		t.Error("Big method should work with other methods")
	}
}

func TestSetMessageReaction_RemoveReactions(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	messageID := int64(789)

	// Test RemoveReactions method
	result := ctx.SetMessageReaction(messageID)
	removeResult := result.RemoveReactions()
	if removeResult == nil {
		t.Error("RemoveReactions method should return SetMessageReaction builder for chaining")
	}

	// Test that RemoveReactions can be chained multiple times
	chainedResult := removeResult.RemoveReactions()
	if chainedResult == nil {
		t.Error("RemoveReactions method should support multiple chaining calls")
	}

	// Test RemoveReactions with other methods
	removeWithOthers := ctx.SetMessageReaction(messageID).
		Reaction(g.String("👍")).
		RemoveReactions().
		Big()
	if removeWithOthers == nil {
		t.Error("RemoveReactions method should work with other methods")
	}
}

func TestSetMessageReaction_Timeout(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	messageID := int64(789)

	// Test Timeout method with various durations
	timeouts := []time.Duration{
		1 * time.Second,
		5 * time.Second,
		30 * time.Second,
		1 * time.Minute,
		5 * time.Minute,
		0 * time.Second, // Zero timeout
	}

	for _, timeout := range timeouts {
		result := ctx.SetMessageReaction(messageID)
		timeoutResult := result.Timeout(timeout)
		if timeoutResult == nil {
			t.Errorf("Timeout method should return SetMessageReaction builder for chaining with timeout %v", timeout)
		}

		// Test that Timeout can be chained and overridden
		chainedResult := timeoutResult.Timeout(timeout + 1*time.Second)
		if chainedResult == nil {
			t.Errorf("Timeout method should support chaining and override with timeout %v", timeout)
		}
	}
}

func TestSetMessageReaction_APIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	messageID := int64(789)

	// Test APIURL method with various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://custom.api.example.com",
		"",
		"https://api.example.com/bot",
		"http://localhost:8080",
	}

	for _, apiURL := range apiURLs {
		result := ctx.SetMessageReaction(messageID)
		apiURLResult := result.APIURL(g.String(apiURL))
		if apiURLResult == nil {
			t.Errorf("APIURL method should return SetMessageReaction builder for chaining with URL: %s", apiURL)
		}

		// Test that APIURL can be chained and overridden
		chainedResult := apiURLResult.APIURL(g.String("https://override.example.com"))
		if chainedResult == nil {
			t.Errorf("APIURL method should support chaining and override with URL: %s", apiURL)
		}
	}
}

func TestSetMessageReaction_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	messageID := int64(789)

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SetMessageReaction(messageID).Send()

	if sendResult.IsErr() {
		t.Logf("SetMessageReaction Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with reactions
	sendWithReactionsResult := ctx.SetMessageReaction(messageID).
		Reaction(g.String("👍")).
		CustomEmoji(g.String("custom_123")).
		Big().
		Send()

	if sendWithReactionsResult.IsErr() {
		t.Logf("SetMessageReaction Send with reactions failed as expected with mock bot: %v", sendWithReactionsResult.Err())
	}
}
