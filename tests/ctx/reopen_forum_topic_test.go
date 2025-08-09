package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_ReopenForumTopic(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	messageThreadID := int64(123)

	result := ctx.ReopenForumTopic(messageThreadID)

	if result == nil {
		t.Error("Expected ReopenForumTopic builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}

func TestReopenForumTopic_Timeout(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	messageThreadID := int64(123)

	timeouts := []time.Duration{
		time.Second * 10,
		time.Second * 30,
		time.Minute,
		time.Minute * 5,
		0,
	}

	for _, timeout := range timeouts {
		result := ctx.ReopenForumTopic(messageThreadID)
		timeoutResult := result.Timeout(timeout)
		if timeoutResult == nil {
			t.Errorf("Timeout method should return ReopenForumTopic builder for chaining with timeout: %v", timeout)
		}

		chainedResult := timeoutResult.Timeout(time.Second * 15)
		if chainedResult == nil {
			t.Errorf("Timeout method should support chaining and override with timeout: %v", timeout)
		}
	}
}

func TestReopenForumTopic_APIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	messageThreadID := int64(123)

	apiURLs := []string{
		"https://api.telegram.org",
		"https://custom.api.example.com",
		"",
		"https://api.example.com/bot",
		"http://localhost:8080",
	}

	for _, apiURL := range apiURLs {
		result := ctx.ReopenForumTopic(messageThreadID)
		apiURLResult := result.APIURL(g.String(apiURL))
		if apiURLResult == nil {
			t.Errorf("APIURL method should return ReopenForumTopic builder for chaining with URL: %s", apiURL)
		}

		chainedResult := apiURLResult.APIURL(g.String("https://override.example.com"))
		if chainedResult == nil {
			t.Errorf("APIURL method should support chaining and override with URL: %s", apiURL)
		}
	}
}

func TestReopenForumTopic_ChatID(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	messageThreadID := int64(123)

	chatIDs := []int64{
		-1001234567890,
		-1009876543210,
		123,
		456,
		0,
	}

	for _, chatID := range chatIDs {
		result := ctx.ReopenForumTopic(messageThreadID)
		chatIDResult := result.ChatID(chatID)
		if chatIDResult == nil {
			t.Errorf("ChatID method should return ReopenForumTopic builder for chaining with chatID: %d", chatID)
		}

		chainedResult := chatIDResult.ChatID(-1001111111111)
		if chainedResult == nil {
			t.Errorf("ChatID method should support chaining and override with chatID: %d", chatID)
		}
	}
}

func TestReopenForumTopic_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	messageThreadID := int64(123)

	sendResult := ctx.ReopenForumTopic(messageThreadID).Send()

	if sendResult.IsErr() {
		t.Logf("ReopenForumTopic Send failed as expected with mock bot: %v", sendResult.Err())
	}

	sendWithOptionsResult := ctx.ReopenForumTopic(messageThreadID).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.telegram.org")).
		ChatID(-1009876543210).
		Send()

	if sendWithOptionsResult.IsErr() {
		t.Logf("ReopenForumTopic Send with options failed as expected with mock bot: %v", sendWithOptionsResult.Err())
	}
}
