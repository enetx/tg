package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_GetGameHighScores(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)

	result := ctx.GetGameHighScores(userID)

	if result == nil {
		t.Error("Expected GetGameHighScores builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(123)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}

// Tests for methods with 0% coverage

func TestGetGameHighScores_UserID(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test UserID method with various user IDs
	userIDs := []int64{123, 456, 789, 0, -1, 999999999}

	for _, userID := range userIDs {
		result := ctx.GetGameHighScores(111)  // Initial userID
		userIDResult := result.UserID(userID) // Override with new userID
		if userIDResult == nil {
			t.Errorf("UserID method should return GetGameHighScores for chaining with userID %d", userID)
		}

		// Test that UserID can be chained
		chainedResult := userIDResult.UserID(userID + 1)
		if chainedResult == nil {
			t.Errorf("UserID method should support chaining with userID %d", userID+1)
		}
	}
}

func TestGetGameHighScores_MessageID(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test MessageID method with various message IDs
	messageIDs := []int64{100, 200, 300, 0, -1, 999999}

	for _, messageID := range messageIDs {
		result := ctx.GetGameHighScores(456)
		messageIDResult := result.MessageID(messageID)
		if messageIDResult == nil {
			t.Errorf("MessageID method should return GetGameHighScores for chaining with messageID %d", messageID)
		}

		// Test method chaining with other methods
		chainedResult := messageIDResult.ChatID(123).UserID(789)
		if chainedResult == nil {
			t.Errorf("MessageID method should support chaining with messageID %d", messageID)
		}
	}
}

func TestGetGameHighScores_InlineMessageID(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test InlineMessageID method with various inline message IDs
	inlineMessageIDs := []string{
		"inline_msg_123",
		"inline_msg_456",
		"inline_msg_789",
		"",
		"very_long_inline_message_id_with_special_chars_!@#$%^&*()",
	}

	for _, inlineMessageID := range inlineMessageIDs {
		result := ctx.GetGameHighScores(456)
		inlineMessageIDResult := result.InlineMessageID(g.String(inlineMessageID))
		if inlineMessageIDResult == nil {
			t.Errorf("InlineMessageID method should return GetGameHighScores for chaining with inlineMessageID '%s'", inlineMessageID)
		}

		// Test method chaining with other methods
		chainedResult := inlineMessageIDResult.UserID(789).Timeout(30 * time.Second)
		if chainedResult == nil {
			t.Errorf("InlineMessageID method should support chaining with inlineMessageID '%s'", inlineMessageID)
		}
	}
}

func TestGetGameHighScores_Timeout(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Timeout method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.GetGameHighScores(456)
	timeoutResultNil := freshResult.Timeout(30 * time.Second)
	if timeoutResultNil == nil {
		t.Error("Timeout method should return GetGameHighScores for chaining with nil RequestOpts")
	}

	// Test Timeout method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.GetGameHighScores(456)
	timeoutFirst := anotherFreshResult.Timeout(15 * time.Second) // This creates RequestOpts
	timeoutSecond := timeoutFirst.Timeout(25 * time.Second)      // This uses existing RequestOpts
	if timeoutSecond == nil {
		t.Error("Timeout method should return GetGameHighScores for chaining with existing RequestOpts")
	}

	// Test various timeout durations
	timeouts := []time.Duration{
		1 * time.Second,
		5 * time.Second,
		15 * time.Second,
		30 * time.Second,
		60 * time.Second,
		2 * time.Minute,
		0 * time.Second, // Zero timeout
	}

	for _, timeout := range timeouts {
		timeoutResult := ctx.GetGameHighScores(456).
			UserID(789).
			ChatID(123).
			MessageID(100).
			Timeout(timeout).
			Send()

		if timeoutResult.IsErr() {
			t.Logf("GetGameHighScores with timeout %v Send failed as expected: %v", timeout, timeoutResult.Err())
		}
	}
}

func TestGetGameHighScores_APIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test APIURL method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.GetGameHighScores(456)
	apiURLResultNil := freshResult.APIURL(g.String("https://custom-game-api.telegram.org"))
	if apiURLResultNil == nil {
		t.Error("APIURL method should return GetGameHighScores for chaining with nil RequestOpts")
	}

	// Test APIURL method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.GetGameHighScores(456)
	apiURLFirst := anotherFreshResult.APIURL(g.String("https://first-game-api.telegram.org")) // This creates RequestOpts
	apiURLSecond := apiURLFirst.APIURL(g.String("https://second-game-api.telegram.org"))      // This uses existing RequestOpts
	if apiURLSecond == nil {
		t.Error("APIURL method should return GetGameHighScores for chaining with existing RequestOpts")
	}

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://game-api.example.com",
		"https://custom-game.telegram.org",
		"https://regional-game-api.telegram.org",
		"https://backup-game-api.telegram.org",
		"", // Empty URL
	}

	for _, apiURL := range apiURLs {
		apiResult := ctx.GetGameHighScores(456).
			UserID(789).
			ChatID(123).
			MessageID(100).
			Timeout(30 * time.Second).
			APIURL(g.String(apiURL)).
			Send()

		if apiResult.IsErr() {
			t.Logf("GetGameHighScores with API URL '%s' Send failed as expected: %v", apiURL, apiResult.Err())
		}
	}
}

func TestGetGameHighScores_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.GetGameHighScores(456).Send()

	if sendResult.IsErr() {
		t.Logf("GetGameHighScores Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with various configurations
	testScenarios := []struct {
		userID          int64
		chatID          int64
		messageID       int64
		inlineMessageID string
		description     string
	}{
		{123, 456, 789, "", "Regular game message"},
		{456, 789, 123, "", "Different user/chat/message"},
		{789, 0, 0, "inline_msg_123", "Inline game with inline message ID"},
		{111, 222, 333, "inline_msg_456", "Mixed configuration"},
		{0, 123, 456, "", "Zero user ID"},
		{999, -1, -1, "", "Negative IDs"},
	}

	for _, scenario := range testScenarios {
		result := ctx.GetGameHighScores(scenario.userID).
			UserID(scenario.userID). // Test UserID override
			ChatID(scenario.chatID).
			MessageID(scenario.messageID)

		if scenario.inlineMessageID != "" {
			result = result.InlineMessageID(g.String(scenario.inlineMessageID))
		}

		sendResult := result.
			Timeout(30 * time.Second).
			APIURL(g.String("https://api.example.com")).
			Send()

		if sendResult.IsErr() {
			t.Logf("GetGameHighScores Send with %s failed as expected: %v", scenario.description, sendResult.Err())
		}
	}

	// Test comprehensive workflow with all methods
	comprehensiveResult := ctx.GetGameHighScores(456).
		UserID(789). // Override initial userID
		ChatID(123).
		MessageID(100).
		InlineMessageID(g.String("inline_comprehensive_test")).
		Timeout(90 * time.Second).
		APIURL(g.String("https://comprehensive-game-api.telegram.org")).
		Send()

	if comprehensiveResult.IsErr() {
		t.Logf("GetGameHighScores comprehensive workflow Send failed as expected: %v", comprehensiveResult.Err())
	}
}
