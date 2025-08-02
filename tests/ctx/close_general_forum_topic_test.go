package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_CloseGeneralForumTopic(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test basic creation
	result := testCtx.CloseGeneralForumTopic()
	if result == nil {
		t.Error("Expected CloseGeneralForumTopic builder to be created")
	}

	// Test ChatID method
	result = result.ChatID(-1001234567890)
	if result == nil {
		t.Error("ChatID method should return CloseGeneralForumTopic for chaining")
	}

	// Test Timeout method
	result = testCtx.CloseGeneralForumTopic().Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return CloseGeneralForumTopic for chaining")
	}

	// Test APIURL method
	result = testCtx.CloseGeneralForumTopic().APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return CloseGeneralForumTopic for chaining")
	}
}

func TestContext_CloseGeneralForumTopicChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test complete method chaining
	result := testCtx.CloseGeneralForumTopic().
		ChatID(-1001987654321).
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return CloseGeneralForumTopic")
	}
}

func TestCloseGeneralForumTopic_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test with zero ChatID (should use effective chat)
	result := testCtx.CloseGeneralForumTopic().ChatID(0)
	if result == nil {
		t.Error("CloseGeneralForumTopic should handle zero ChatID")
	}

	// Test with zero timeout
	result = testCtx.CloseGeneralForumTopic().Timeout(0 * time.Second)
	if result == nil {
		t.Error("CloseGeneralForumTopic should handle zero timeout")
	}

	// Test with very long timeout
	result = testCtx.CloseGeneralForumTopic().Timeout(24 * time.Hour)
	if result == nil {
		t.Error("CloseGeneralForumTopic should handle very long timeout")
	}

	// Test with empty API URL
	result = testCtx.CloseGeneralForumTopic().APIURL(g.String(""))
	if result == nil {
		t.Error("CloseGeneralForumTopic should handle empty API URL")
	}
}

func TestCloseGeneralForumTopic_DefaultChatID(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test without setting ChatID (should use EffectiveChat.Id)
	result := testCtx.CloseGeneralForumTopic()
	if result == nil {
		t.Error("CloseGeneralForumTopic should work without explicit ChatID")
	}

	// The Send() method should use EffectiveChat.Id when no ChatID is set
	// We can't test the actual API call with mockBot, but we can ensure the builder works
}

func TestCloseGeneralForumTopic_ChatIDVariations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test various chat ID scenarios (forum groups must be supergroups)
	chatIDs := []int64{
		-1001234567890, // Supergroup with forum enabled
		-1001987654321, // Another supergroup
		-1002000000000, // Large supergroup ID
		-1009999999999, // Very large supergroup ID
	}

	for _, chatID := range chatIDs {
		result := testCtx.CloseGeneralForumTopic().ChatID(chatID)
		if result == nil {
			t.Errorf("CloseGeneralForumTopic should handle chat ID: %d", chatID)
		}

		// Test combining with other methods
		combinedResult := result.Timeout(15 * time.Second).APIURL(g.String("https://api.telegram.org"))
		if combinedResult == nil {
			t.Errorf("Combined methods should work for chat ID: %d", chatID)
		}
	}
}

func TestCloseGeneralForumTopic_TimeoutVariations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test various timeout durations
	timeouts := []time.Duration{
		1 * time.Second,
		30 * time.Second,
		5 * time.Minute,
		1 * time.Hour,
		24 * time.Hour,
	}

	for _, timeout := range timeouts {
		result := testCtx.CloseGeneralForumTopic().Timeout(timeout)
		if result == nil {
			t.Errorf("CloseGeneralForumTopic should handle timeout: %v", timeout)
		}

		// Test combining timeout with other methods
		combinedResult := result.ChatID(-1001111222333).APIURL(g.String("https://custom.api.com"))
		if combinedResult == nil {
			t.Errorf("Timeout combination should work for: %v", timeout)
		}
	}
}

func TestCloseGeneralForumTopic_APIURLVariations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://custom-api.example.com",
		"https://bot-api.mycompany.com",
		"https://localhost:8080",
		"https://api-staging.telegram.org",
		"https://proxy.telegram-api.com",
	}

	for _, apiURL := range apiURLs {
		result := testCtx.CloseGeneralForumTopic().APIURL(g.String(apiURL))
		if result == nil {
			t.Errorf("CloseGeneralForumTopic should handle API URL: %s", apiURL)
		}

		// Test combining API URL with other methods
		combinedResult := result.ChatID(-1001444555666).Timeout(35 * time.Second)
		if combinedResult == nil {
			t.Errorf("API URL combination should work for: %s", apiURL)
		}
	}
}

func TestCloseGeneralForumTopic_ForumManagement(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test general forum topic closure scenarios
	testCases := []struct {
		name        string
		chatID      int64
		description string
	}{
		{"Primary Forum", -1001234567890, "Main forum supergroup"},
		{"Secondary Forum", -1001987654321, "Secondary forum supergroup"},
		{"Large Forum", -1002000000000, "Large forum community"},
		{"Enterprise Forum", -1003333333333, "Enterprise forum"},
		{"Community Forum", -1004444444444, "Community discussion forum"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := testCtx.CloseGeneralForumTopic().ChatID(tc.chatID)
			if result == nil {
				t.Errorf("CloseGeneralForumTopic should handle %s (%s): %d", tc.name, tc.description, tc.chatID)
			}

			// Test complete workflow for each scenario
			completedResult := result.
				Timeout(30 * time.Second).
				APIURL(g.String("https://api.telegram.org"))

			if completedResult == nil {
				t.Errorf("Complete workflow should work for %s (%s)", tc.name, tc.description)
			}
		})
	}
}

func TestCloseGeneralForumTopic_AdminOperations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test admin operation scenarios
	adminScenarios := []struct {
		name        string
		timeout     time.Duration
		description string
	}{
		{"Quick Close", 5 * time.Second, "Quick administrative closure"},
		{"Standard Close", 30 * time.Second, "Standard closure with normal timeout"},
		{"Extended Close", 2 * time.Minute, "Extended closure for complex operations"},
		{"Emergency Close", 1 * time.Second, "Emergency closure with minimal timeout"},
	}

	for _, scenario := range adminScenarios {
		t.Run(scenario.name, func(t *testing.T) {
			result := testCtx.CloseGeneralForumTopic().
				ChatID(-1001234567890).
				Timeout(scenario.timeout).
				APIURL(g.String("https://api.telegram.org"))

			if result == nil {
				t.Errorf("%s (%s) should work", scenario.name, scenario.description)
			}
		})
	}
}

func TestCloseGeneralForumTopic_MethodCoverage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test all method combinations systematically
	baseBuilder := testCtx.CloseGeneralForumTopic()

	// Test ChatID variations
	chatIDBuilder := baseBuilder.ChatID(-1001111111111)
	if chatIDBuilder == nil {
		t.Error("ChatID method should work")
	}

	// Test Timeout variations
	timeoutBuilder := testCtx.CloseGeneralForumTopic().Timeout(45 * time.Second)
	if timeoutBuilder == nil {
		t.Error("Timeout method should work")
	}

	// Test APIURL variations
	apiBuilder := testCtx.CloseGeneralForumTopic().APIURL(g.String("https://custom.example.com"))
	if apiBuilder == nil {
		t.Error("APIURL method should work")
	}

	// Test all methods combined
	combinedBuilder := testCtx.CloseGeneralForumTopic().
		ChatID(-1002222222222).
		Timeout(60 * time.Second).
		APIURL(g.String("https://enterprise-api.example.com"))

	if combinedBuilder == nil {
		t.Error("All methods combined should work")
	}

	// Test method order independence
	reorderedBuilder := testCtx.CloseGeneralForumTopic().
		APIURL(g.String("https://reordered-api.example.com")).
		Timeout(75 * time.Second).
		ChatID(-1003333333333)

	if reorderedBuilder == nil {
		t.Error("Method order independence should work")
	}
}
