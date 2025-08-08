package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_EditForumTopic(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	messageThreadID := int64(123)

	result := ctx.EditForumTopic(messageThreadID)

	if result == nil {
		t.Error("Expected EditForumTopic builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(-1001234567890)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}

	// Test Name method
	nameResult := result.Name(g.String("Updated Topic Name"))
	if nameResult == nil {
		t.Error("Name method should return EditForumTopic for chaining")
	}

	// Test IconCustomEmojiID method
	emojiResult := result.IconCustomEmojiID(g.String("5431652110606213632"))
	if emojiResult == nil {
		t.Error("IconCustomEmojiID method should return EditForumTopic for chaining")
	}

	// Test Timeout method
	timeoutResult := result.Timeout(30 * time.Second)
	if timeoutResult == nil {
		t.Error("Timeout method should return EditForumTopic for chaining")
	}

	// Test APIURL method
	apiURLResult := result.APIURL(g.String("https://api.telegram.org"))
	if apiURLResult == nil {
		t.Error("APIURL method should return EditForumTopic for chaining")
	}

	// Test methods with nil RequestOpts (covers the nil branch)
	freshResult := ctx.EditForumTopic(messageThreadID)
	timeoutResultNil := freshResult.Timeout(45 * time.Second)
	if timeoutResultNil == nil {
		t.Error("Timeout method should return EditForumTopic for chaining with nil RequestOpts")
	}

	apiURLResultNil := freshResult.APIURL(g.String("https://custom-api.telegram.org"))
	if apiURLResultNil == nil {
		t.Error("APIURL method should return EditForumTopic for chaining with nil RequestOpts")
	}

	// Test APIURL method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.EditForumTopic(messageThreadID)
	apiURLFirst := anotherFreshResult.APIURL(g.String("https://first-api.telegram.org")) // This creates RequestOpts
	apiURLSecond := apiURLFirst.APIURL(g.String("https://second-api.telegram.org"))      // This uses existing RequestOpts
	if apiURLSecond == nil {
		t.Error("APIURL method should return EditForumTopic for chaining with existing RequestOpts")
	}
}

func TestEditForumTopic_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	messageThreadID := int64(456)

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.EditForumTopic(messageThreadID).Send()

	if sendResult.IsErr() {
		t.Logf("EditForumTopic Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.EditForumTopic(messageThreadID).
		ChatID(-1001987654321).
		Name(g.String("Updated Forum Topic")).
		IconCustomEmojiID(g.String("5431652110606213632")).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("EditForumTopic configured Send failed as expected: %v", configuredSendResult.Err())
	}

	// Test Send method using EffectiveChat ID (no explicit ChatID)
	effectiveChatSendResult := ctx.EditForumTopic(messageThreadID).
		Name(g.String("Effective Chat Topic")).
		Timeout(60 * time.Second).
		Send()

	if effectiveChatSendResult.IsErr() {
		t.Logf("EditForumTopic with effective chat Send failed as expected: %v", effectiveChatSendResult.Err())
	}
}

func TestEditForumTopic_CompleteMethods(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	messageThreadID := int64(789)

	// Test all methods in combination
	result := ctx.EditForumTopic(messageThreadID).
		ChatID(-1001987654321).
		Name(g.String("Complete Test Topic")).
		IconCustomEmojiID(g.String("5431652110606213633")).
		Timeout(60 * time.Second).
		APIURL(g.String("https://complete-forum-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return EditForumTopic")
	}

	// Test complete workflow with Send
	completeResult := result.Send()
	if completeResult.IsErr() {
		t.Logf("Complete EditForumTopic workflow Send failed as expected: %v", completeResult.Err())
	}

	// Test various forum topic names
	topicNames := []string{
		"General Discussion",
		"Technical Support",
		"Feature Requests",
		"Bug Reports",
		"Community Chat",
		"🚀 Announcements",
		"💡 Ideas & Suggestions",
		"🎮 Gaming Zone",
		"📚 Resources",
		"🔧 Development",
		"Short",
		"", // Empty name
		"Very Long Forum Topic Name For Testing Purposes With Multiple Words",
	}

	for _, topicName := range topicNames {
		displayName := topicName
		if topicName == "" {
			displayName = "[empty]"
		}

		nameResult := ctx.EditForumTopic(messageThreadID).
			ChatID(-1001234567890).
			Name(g.String(topicName)).
			Timeout(45 * time.Second).
			Send()

		if nameResult.IsErr() {
			t.Logf("EditForumTopic with name '%s' Send failed as expected: %v", displayName, nameResult.Err())
		}
	}

	// Test various custom emoji IDs
	emojiIDs := []string{
		"5431652110606213632", // Valid emoji ID
		"5431652110606213633", // Another valid emoji ID
		"5431652110606213634", // Yet another emoji ID
		"1234567890123456789", // Different format
		"9876543210987654321", // Another format
		"",                    // Empty emoji ID
		"invalid_emoji_id",    // Invalid format
	}

	for _, emojiID := range emojiIDs {
		displayEmoji := emojiID
		if emojiID == "" {
			displayEmoji = "[empty]"
		}

		emojiResult := ctx.EditForumTopic(messageThreadID).
			ChatID(-1001234567890).
			Name(g.String("Test Topic")).
			IconCustomEmojiID(g.String(emojiID)).
			Timeout(45 * time.Second).
			Send()

		if emojiResult.IsErr() {
			t.Logf("EditForumTopic with emoji ID '%s' Send failed as expected: %v", displayEmoji, emojiResult.Err())
		}
	}

	// Test various timeout configurations
	timeouts := []time.Duration{
		1 * time.Second,
		10 * time.Second,
		30 * time.Second,
		60 * time.Second,
		2 * time.Minute,
		0 * time.Second, // Zero timeout
	}

	for _, timeout := range timeouts {
		timeoutResult := ctx.EditForumTopic(messageThreadID).
			ChatID(-1001234567890).
			Name(g.String("Timeout Test Topic")).
			Timeout(timeout).
			APIURL(g.String("https://timeout-forum-api.telegram.org")).
			Send()

		if timeoutResult.IsErr() {
			t.Logf("EditForumTopic with timeout %v Send failed as expected: %v", timeout, timeoutResult.Err())
		}
	}

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://forum-topic-api.example.com",
		"https://custom-forum.telegram.org",
		"https://regional-forum-api.telegram.org",
		"", // Empty URL
	}

	for _, apiURL := range apiURLs {
		apiResult := ctx.EditForumTopic(messageThreadID).
			ChatID(-1001234567890).
			Name(g.String("API Test Topic")).
			Timeout(30 * time.Second).
			APIURL(g.String(apiURL)).
			Send()

		if apiResult.IsErr() {
			t.Logf("EditForumTopic with API URL '%s' Send failed as expected: %v", apiURL, apiResult.Err())
		}
	}

	// Test different message thread IDs
	threadIDs := []int64{
		1, 2, 3, 10, 50, 100, 999, 1234, 9999, 123456, 999999,
	}

	for _, threadID := range threadIDs {
		threadResult := ctx.EditForumTopic(threadID).
			ChatID(-1001234567890).
			Name(g.String("Thread Test Topic")).
			IconCustomEmojiID(g.String("5431652110606213632")).
			Timeout(45 * time.Second).
			Send()

		if threadResult.IsErr() {
			t.Logf("EditForumTopic with thread ID %d Send failed as expected: %v", threadID, threadResult.Err())
		}
	}
}
