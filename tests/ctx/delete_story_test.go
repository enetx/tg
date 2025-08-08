package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_DeleteStory(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_123")
	storyID := int64(456)

	result := ctx.DeleteStory(businessConnectionID, storyID)

	if result == nil {
		t.Error("Expected DeleteStory builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(30 * time.Second)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}

	// Test APIURL method
	apiURLResult := result.APIURL(g.String("https://api.telegram.org"))
	if apiURLResult == nil {
		t.Error("APIURL method should return DeleteStory for chaining")
	}

	// Test APIURL method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.DeleteStory(businessConnectionID, storyID)
	apiURLResultNil := freshResult.APIURL(g.String("https://custom-api.telegram.org"))
	if apiURLResultNil == nil {
		t.Error("APIURL method should return DeleteStory for chaining with nil RequestOpts")
	}
}

func TestDeleteStory_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_456")
	storyID := int64(789)

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.DeleteStory(businessConnectionID, storyID).Send()

	if sendResult.IsErr() {
		t.Logf("DeleteStory Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.DeleteStory(businessConnectionID, storyID).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("DeleteStory configured Send failed as expected: %v", configuredSendResult.Err())
	}

	// Test Send method with different business connection IDs and story IDs
	businessConnections := []string{
		"business_conn_123",
		"business_conn_456",
		"business_conn_789",
		"conn_premium_001",
		"enterprise_conn_002",
		"", // Empty connection ID
	}

	storyIDs := []int64{
		100, 200, 300, 456, 789, 1000, 9999, 123456,
	}

	for _, connID := range businessConnections {
		for _, storyNum := range storyIDs {
			idResult := ctx.DeleteStory(g.String(connID), storyNum).
				Timeout(45 * time.Second).
				Send()

			if idResult.IsErr() {
				t.Logf("DeleteStory with connection '%s' and story %d Send failed as expected: %v", connID, storyNum, idResult.Err())
			}
		}
	}
}

func TestDeleteStory_CompleteMethods(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("comprehensive_test_conn")
	storyID := int64(12345)

	// Test all methods in combination
	result := ctx.DeleteStory(businessConnectionID, storyID).
		Timeout(60 * time.Second).
		APIURL(g.String("https://complete-story-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return DeleteStory")
	}

	// Test complete workflow with Send
	completeResult := result.Send()
	if completeResult.IsErr() {
		t.Logf("Complete DeleteStory workflow Send failed as expected: %v", completeResult.Err())
	}

	// Test various timeout configurations
	timeouts := []time.Duration{
		1 * time.Second,
		10 * time.Second,
		30 * time.Second,
		60 * time.Second,
		2 * time.Minute,
		5 * time.Minute,
		0 * time.Second, // Zero timeout
	}

	for _, timeout := range timeouts {
		timeoutResult := ctx.DeleteStory(businessConnectionID, storyID).
			Timeout(timeout).
			APIURL(g.String("https://timeout-story-api.telegram.org")).
			Send()

		if timeoutResult.IsErr() {
			t.Logf("DeleteStory with timeout %v Send failed as expected: %v", timeout, timeoutResult.Err())
		}
	}

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://story-management-api.example.com",
		"https://business-story-api.telegram.org",
		"https://custom-story-delete.example.com",
		"https://regional-story-api.telegram.org",
		"", // Empty URL
	}

	for _, apiURL := range apiURLs {
		apiResult := ctx.DeleteStory(businessConnectionID, storyID).
			Timeout(30 * time.Second).
			APIURL(g.String(apiURL)).
			Send()

		if apiResult.IsErr() {
			t.Logf("DeleteStory with API URL '%s' Send failed as expected: %v", apiURL, apiResult.Err())
		}
	}

	// Test edge cases with story scenarios
	storyScenarios := []struct {
		connectionID string
		storyID      int64
		description  string
	}{
		{"personal_conn", 1, "Personal story deletion"},
		{"business_conn_small", 999, "Small business story"},
		{"enterprise_conn", 1000000, "Enterprise large story ID"},
		{"premium_conn", 42, "Premium account story"},
		{"test_conn", 0, "Edge case: zero story ID"},
		{"", 123, "Empty connection ID"},
		{"very_long_connection_id_name", 456, "Long connection ID name"},
		{"conn123", 789, "Alphanumeric connection ID"},
	}

	for _, scenario := range storyScenarios {
		scenarioResult := ctx.DeleteStory(g.String(scenario.connectionID), scenario.storyID).
			Timeout(45 * time.Second).
			APIURL(g.String("https://scenario-api.telegram.org")).
			Send()

		if scenarioResult.IsErr() {
			t.Logf("DeleteStory scenario '%s' (conn: %s, story: %d) Send failed as expected: %v",
				scenario.description, scenario.connectionID, scenario.storyID, scenarioResult.Err())
		}
	}
}
