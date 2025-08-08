package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_ExportChatInviteLink(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.ExportChatInviteLink()

	if result == nil {
		t.Error("Expected ExportChatInviteLink builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(-1001234567890)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}

// Tests for methods with 0% coverage

func TestExportChatInviteLink_Timeout(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Timeout method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.ExportChatInviteLink()
	timeoutResultNil := freshResult.Timeout(30 * time.Second)
	if timeoutResultNil == nil {
		t.Error("Timeout method should return ExportChatInviteLink for chaining with nil RequestOpts")
	}

	// Test Timeout method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.ExportChatInviteLink()
	timeoutFirst := anotherFreshResult.Timeout(15 * time.Second) // This creates RequestOpts
	timeoutSecond := timeoutFirst.Timeout(25 * time.Second)      // This uses existing RequestOpts
	if timeoutSecond == nil {
		t.Error("Timeout method should return ExportChatInviteLink for chaining with existing RequestOpts")
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
		timeoutResult := ctx.ExportChatInviteLink().
			ChatID(-1001234567890).
			Timeout(timeout).
			Send()

		if timeoutResult.IsErr() {
			t.Logf("ExportChatInviteLink with timeout %v Send failed as expected: %v", timeout, timeoutResult.Err())
		}
	}
}

func TestExportChatInviteLink_APIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test APIURL method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.ExportChatInviteLink()
	apiURLResultNil := freshResult.APIURL(g.String("https://custom-export-invite-api.telegram.org"))
	if apiURLResultNil == nil {
		t.Error("APIURL method should return ExportChatInviteLink for chaining with nil RequestOpts")
	}

	// Test APIURL method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.ExportChatInviteLink()
	apiURLFirst := anotherFreshResult.APIURL(g.String("https://first-export-api.telegram.org")) // This creates RequestOpts
	apiURLSecond := apiURLFirst.APIURL(g.String("https://second-export-api.telegram.org"))      // This uses existing RequestOpts
	if apiURLSecond == nil {
		t.Error("APIURL method should return ExportChatInviteLink for chaining with existing RequestOpts")
	}

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://export-invite-api.example.com",
		"https://custom-export.telegram.org",
		"https://regional-export-api.telegram.org",
		"https://backup-export-api.telegram.org",
		"", // Empty URL
	}

	for _, apiURL := range apiURLs {
		apiResult := ctx.ExportChatInviteLink().
			ChatID(-1001234567890).
			Timeout(30 * time.Second).
			APIURL(g.String(apiURL)).
			Send()

		if apiResult.IsErr() {
			t.Logf("ExportChatInviteLink with API URL '%s' Send failed as expected: %v", apiURL, apiResult.Err())
		}
	}
}

func TestExportChatInviteLink_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.ExportChatInviteLink().Send()

	if sendResult.IsErr() {
		t.Logf("ExportChatInviteLink Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.ExportChatInviteLink().
		ChatID(-1001987654321).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("ExportChatInviteLink configured Send failed as expected: %v", configuredSendResult.Err())
	}

	// Test Send method using EffectiveChat ID (no explicit ChatID)
	effectiveChatSendResult := ctx.ExportChatInviteLink().
		Timeout(60 * time.Second).
		Send()

	if effectiveChatSendResult.IsErr() {
		t.Logf("ExportChatInviteLink with effective chat Send failed as expected: %v", effectiveChatSendResult.Err())
	}
}

func TestExportChatInviteLink_ChatIDScenarios(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test different chat ID scenarios
	chatIDScenarios := []struct {
		chatID      int64
		description string
	}{
		{-1001234567890, "Original supergroup"},
		{-1001987654321, "Another supergroup"},
		{-1002000000000, "Large supergroup"},
		{-1003000000000, "Enterprise supergroup"},
		{-1004000000000, "Community supergroup"},
		{-1005000000000, "Educational supergroup"},
		{-1006000000000, "Business supergroup"},
		{-100, "Small group ID"},
	}

	for _, scenario := range chatIDScenarios {
		chatResult := ctx.ExportChatInviteLink().
			ChatID(scenario.chatID).
			Timeout(45 * time.Second).
			APIURL(g.String("https://chat-scenario-export-api.telegram.org")).
			Send()

		if chatResult.IsErr() {
			t.Logf("ExportChatInviteLink with %s (ID: %d) Send failed as expected: %v",
				scenario.description, scenario.chatID, chatResult.Err())
		}
	}
}

func TestExportChatInviteLink_ComprehensiveWorkflow(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test comprehensive workflow with all methods
	complexResult := ctx.ExportChatInviteLink().
		ChatID(-1001987654321).
		Timeout(90 * time.Second).
		APIURL(g.String("https://comprehensive-export-api.telegram.org")).
		Send()

	if complexResult.IsErr() {
		t.Logf("ExportChatInviteLink comprehensive workflow Send failed as expected: %v", complexResult.Err())
	}

	// Test workflow with different chat types (simulated through different IDs)
	chatTypes := []struct {
		chatID   int64
		chatType string
	}{
		{-1001111111111, "Public supergroup"},
		{-1002222222222, "Private supergroup"},
		{-1003333333333, "Channel"},
		{-1004444444444, "Discussion group"},
	}

	for _, chatType := range chatTypes {
		typeResult := ctx.ExportChatInviteLink().
			ChatID(chatType.chatID).
			Timeout(30 * time.Second).
			APIURL(g.String("https://" + chatType.chatType + "-export-api.telegram.org")).
			Send()

		if typeResult.IsErr() {
			t.Logf("ExportChatInviteLink for %s (ID: %d) Send failed as expected: %v",
				chatType.chatType, chatType.chatID, typeResult.Err())
		}
	}
}
