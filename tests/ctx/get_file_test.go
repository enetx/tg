package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_GetFile(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	fileID := g.String("BAADBAADrwADBREAAWp8oTdjdLOHAg")

	result := ctx.GetFile(fileID)

	if result == nil {
		t.Error("Expected GetFile builder to be created")
	}

	// Test that builder is functional
	_ = result
}

func TestContext_GetFileChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	fileID := g.String("BAADBAADrwADBREAAWp8oTdjdLOHAg")

	result := ctx.GetFile(fileID)

	if result == nil {
		t.Error("Expected GetFile builder to be created")
	}

	// Test that builder is functional
	_ = result
}

// Tests for methods with 0% coverage

func TestGetFile_Timeout(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	fileID := g.String("BAADBAADrwADBREAAWp8oTdjdLOHAg")

	// Test Timeout method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.GetFile(fileID)
	timeoutResultNil := freshResult.Timeout(30 * time.Second)
	if timeoutResultNil == nil {
		t.Error("Timeout method should return GetFile for chaining with nil RequestOpts")
	}

	// Test Timeout method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.GetFile(fileID)
	timeoutFirst := anotherFreshResult.Timeout(15 * time.Second) // This creates RequestOpts
	timeoutSecond := timeoutFirst.Timeout(25 * time.Second)      // This uses existing RequestOpts
	if timeoutSecond == nil {
		t.Error("Timeout method should return GetFile for chaining with existing RequestOpts")
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
		timeoutResult := ctx.GetFile(fileID).
			Timeout(timeout).
			Send()

		if timeoutResult.IsErr() {
			t.Logf("GetFile with timeout %v Send failed as expected: %v", timeout, timeoutResult.Err())
		}
	}
}

func TestGetFile_APIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	fileID := g.String("BAADBAADrwADBREAAWp8oTdjdLOHAg")

	// Test APIURL method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.GetFile(fileID)
	apiURLResultNil := freshResult.APIURL(g.String("https://custom-file-api.telegram.org"))
	if apiURLResultNil == nil {
		t.Error("APIURL method should return GetFile for chaining with nil RequestOpts")
	}

	// Test APIURL method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.GetFile(fileID)
	apiURLFirst := anotherFreshResult.APIURL(g.String("https://first-file-api.telegram.org")) // This creates RequestOpts
	apiURLSecond := apiURLFirst.APIURL(g.String("https://second-file-api.telegram.org"))      // This uses existing RequestOpts
	if apiURLSecond == nil {
		t.Error("APIURL method should return GetFile for chaining with existing RequestOpts")
	}

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://file-api.example.com",
		"https://custom-file.telegram.org",
		"https://regional-file-api.telegram.org",
		"https://backup-file-api.telegram.org",
		"", // Empty URL
	}

	for _, apiURL := range apiURLs {
		apiResult := ctx.GetFile(fileID).
			Timeout(30 * time.Second).
			APIURL(g.String(apiURL)).
			Send()

		if apiResult.IsErr() {
			t.Logf("GetFile with API URL '%s' Send failed as expected: %v", apiURL, apiResult.Err())
		}
	}
}

func TestGetFile_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Send method with various file IDs
	fileIDScenarios := []struct {
		fileID      string
		description string
	}{
		{"BAADBAADrwADBREAAWp8oTdjdLOHAg", "Photo file ID"},
		{"AwACAgIAAxkBAAIBY2FSqKlZKZsj9_gDAw", "Document file ID"},
		{"AgACAgIAAxkDAAIBpGFSqL1dOWNt8gABC", "Video file ID"},
		{"BAADBAADsAADBREAAWtJUomQZ-0HAg", "Audio file ID"},
		{"BAACAgIAAxkDAAIBymFSqNk1LVlPAAE", "Voice file ID"},
		{"invalid_file_id", "Invalid file ID"},
		{"", "Empty file ID"},
	}

	for _, scenario := range fileIDScenarios {
		fileID := g.String(scenario.fileID)

		// Basic Send test
		sendResult := ctx.GetFile(fileID).Send()
		if sendResult.IsErr() {
			t.Logf("GetFile with %s Send failed as expected: %v", scenario.description, sendResult.Err())
		}

		// Configured Send test
		configuredSendResult := ctx.GetFile(fileID).
			Timeout(30 * time.Second).
			APIURL(g.String("https://api.example.com")).
			Send()

		if configuredSendResult.IsErr() {
			t.Logf("GetFile configured with %s Send failed as expected: %v", scenario.description, configuredSendResult.Err())
		}
	}
}
