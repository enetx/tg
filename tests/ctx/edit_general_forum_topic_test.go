package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_EditGeneralForumTopic(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	name := g.String("Updated General Topic")

	result := ctx.EditGeneralForumTopic(name)

	if result == nil {
		t.Error("Expected EditGeneralForumTopic builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(-1001234567890)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}

	// Test Timeout method
	timeoutResult := result.Timeout(30 * time.Second)
	if timeoutResult == nil {
		t.Error("Timeout method should return EditGeneralForumTopic for chaining")
	}

	// Test APIURL method
	apiURLResult := result.APIURL(g.String("https://api.telegram.org"))
	if apiURLResult == nil {
		t.Error("APIURL method should return EditGeneralForumTopic for chaining")
	}

	// Test methods with nil RequestOpts (covers the nil branch)
	freshResult := ctx.EditGeneralForumTopic(name)
	timeoutResultNil := freshResult.Timeout(45 * time.Second)
	if timeoutResultNil == nil {
		t.Error("Timeout method should return EditGeneralForumTopic for chaining with nil RequestOpts")
	}

	apiURLResultNil := freshResult.APIURL(g.String("https://custom-api.telegram.org"))
	if apiURLResultNil == nil {
		t.Error("APIURL method should return EditGeneralForumTopic for chaining with nil RequestOpts")
	}

	// Test APIURL method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.EditGeneralForumTopic(name)
	apiURLFirst := anotherFreshResult.APIURL(g.String("https://first-api.telegram.org")) // This creates RequestOpts
	apiURLSecond := apiURLFirst.APIURL(g.String("https://second-api.telegram.org"))      // This uses existing RequestOpts
	if apiURLSecond == nil {
		t.Error("APIURL method should return EditGeneralForumTopic for chaining with existing RequestOpts")
	}

	// Test Timeout method with existing RequestOpts (covers the non-nil branch)
	timeoutFirst := anotherFreshResult.Timeout(15 * time.Second) // This creates RequestOpts
	timeoutSecond := timeoutFirst.Timeout(25 * time.Second)      // This uses existing RequestOpts
	if timeoutSecond == nil {
		t.Error("Timeout method should return EditGeneralForumTopic for chaining with existing RequestOpts")
	}
}

func TestEditGeneralForumTopic_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	topicName := g.String("Updated General Forum Topic")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.EditGeneralForumTopic(topicName).Send()

	if sendResult.IsErr() {
		t.Logf("EditGeneralForumTopic Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.EditGeneralForumTopic(topicName).
		ChatID(-1001987654321).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("EditGeneralForumTopic configured Send failed as expected: %v", configuredSendResult.Err())
	}

	// Test Send method using EffectiveChat ID (no explicit ChatID)
	effectiveChatSendResult := ctx.EditGeneralForumTopic(topicName).
		Timeout(60 * time.Second).
		Send()

	if effectiveChatSendResult.IsErr() {
		t.Logf("EditGeneralForumTopic with effective chat Send failed as expected: %v", effectiveChatSendResult.Err())
	}
}

func TestEditGeneralForumTopic_CompleteMethods(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	topicName := g.String("Comprehensive Test General Topic")

	// Test all methods in combination
	result := ctx.EditGeneralForumTopic(topicName).
		ChatID(-1001987654321).
		Timeout(60 * time.Second).
		APIURL(g.String("https://complete-general-forum-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return EditGeneralForumTopic")
	}

	// Test complete workflow with Send
	completeResult := result.Send()
	if completeResult.IsErr() {
		t.Logf("Complete EditGeneralForumTopic workflow Send failed as expected: %v", completeResult.Err())
	}

	// Test various general forum topic names
	generalTopicNames := []string{
		"General Discussion",
		"General",
		"Main Topic",
		"Community Hub",
		"Central Discussion",
		"🌟 General Chat",
		"💬 Main Discussion",
		"📋 General Information",
		"🔘 General Topics",
		"⭐ Main Channel",
		"Short",
		"", // Empty name
		"Very Long General Forum Topic Name For Testing Extended Character Limits",
		"General Topic with Special Characters: !@#$%^&*()",
		"日本語の一般トピック",        // Japanese characters
		"Общие темы форума", // Russian characters
	}

	for _, generalName := range generalTopicNames {
		displayName := generalName
		if generalName == "" {
			displayName = "[empty]"
		}

		nameResult := ctx.EditGeneralForumTopic(g.String(generalName)).
			ChatID(-1001234567890).
			Timeout(45 * time.Second).
			Send()

		if nameResult.IsErr() {
			t.Logf("EditGeneralForumTopic with name '%s' Send failed as expected: %v", displayName, nameResult.Err())
		}
	}

	// Test various timeout configurations
	timeouts := []time.Duration{
		1 * time.Second,
		5 * time.Second,
		15 * time.Second,
		30 * time.Second,
		60 * time.Second,
		2 * time.Minute,
		5 * time.Minute,
		0 * time.Second, // Zero timeout
	}

	for _, timeout := range timeouts {
		timeoutResult := ctx.EditGeneralForumTopic(topicName).
			ChatID(-1001234567890).
			Timeout(timeout).
			APIURL(g.String("https://timeout-general-forum-api.telegram.org")).
			Send()

		if timeoutResult.IsErr() {
			t.Logf("EditGeneralForumTopic with timeout %v Send failed as expected: %v", timeout, timeoutResult.Err())
		}
	}

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://general-forum-api.example.com",
		"https://custom-general-forum.telegram.org",
		"https://regional-general-forum-api.telegram.org",
		"https://backup-general-forum-api.telegram.org",
		"", // Empty URL
	}

	for _, apiURL := range apiURLs {
		apiResult := ctx.EditGeneralForumTopic(topicName).
			ChatID(-1001234567890).
			Timeout(30 * time.Second).
			APIURL(g.String(apiURL)).
			Send()

		if apiResult.IsErr() {
			t.Logf("EditGeneralForumTopic with API URL '%s' Send failed as expected: %v", apiURL, apiResult.Err())
		}
	}

	// Test different chat ID scenarios
	chatIDScenarios := []struct {
		chatID      int64
		description string
	}{
		{-1001234567890, "Standard supergroup"},
		{-1001987654321, "Another supergroup"},
		{-1002000000000, "Large supergroup"},
		{-1003000000000, "Enterprise supergroup"},
		{-1004000000000, "Community supergroup"},
		{-1005000000000, "Educational supergroup"},
		{-1006000000000, "Business supergroup"},
		{-1007000000000, "Gaming supergroup"},
	}

	for _, scenario := range chatIDScenarios {
		chatResult := ctx.EditGeneralForumTopic(topicName).
			ChatID(scenario.chatID).
			Timeout(45 * time.Second).
			APIURL(g.String("https://chat-scenario-api.telegram.org")).
			Send()

		if chatResult.IsErr() {
			t.Logf("EditGeneralForumTopic with %s (ID: %d) Send failed as expected: %v",
				scenario.description, scenario.chatID, chatResult.Err())
		}
	}
}
