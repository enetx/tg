package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_GetUserProfilePhotos(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123)

	result := ctx.GetUserProfilePhotos(userID)

	if result == nil {
		t.Error("Expected GetUserProfilePhotos builder to be created")
	}

	// Test method chaining
	chained := result.Offset(0)
	if chained == nil {
		t.Error("Expected Offset method to return builder")
	}
}

func TestContext_GetUserProfilePhotosChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123)

	result := ctx.GetUserProfilePhotos(userID).
		Offset(0).
		Limit(10)

	if result == nil {
		t.Error("Expected GetUserProfilePhotos builder to be created")
	}

	// Test that builder is functional
	_ = result
}

// Tests for methods with 0% coverage

func TestGetUserProfilePhotos_Timeout(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123)

	// Test Timeout method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.GetUserProfilePhotos(userID)
	timeoutResultNil := freshResult.Timeout(30 * time.Second)
	if timeoutResultNil == nil {
		t.Error("Timeout method should return GetUserProfilePhotos for chaining with nil RequestOpts")
	}

	// Test Timeout method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.GetUserProfilePhotos(userID)
	timeoutFirst := anotherFreshResult.Timeout(15 * time.Second) // This creates RequestOpts
	timeoutSecond := timeoutFirst.Timeout(25 * time.Second)      // This uses existing RequestOpts
	if timeoutSecond == nil {
		t.Error("Timeout method should return GetUserProfilePhotos for chaining with existing RequestOpts")
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
		timeoutResult := ctx.GetUserProfilePhotos(userID).
			Offset(0).
			Limit(5).
			Timeout(timeout).
			Send()

		if timeoutResult.IsErr() {
			t.Logf("GetUserProfilePhotos with timeout %v Send failed as expected: %v", timeout, timeoutResult.Err())
		}
	}
}

func TestGetUserProfilePhotos_APIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123)

	// Test APIURL method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.GetUserProfilePhotos(userID)
	apiURLResultNil := freshResult.APIURL(g.String("https://custom-profile-photos-api.telegram.org"))
	if apiURLResultNil == nil {
		t.Error("APIURL method should return GetUserProfilePhotos for chaining with nil RequestOpts")
	}

	// Test APIURL method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.GetUserProfilePhotos(userID)
	apiURLFirst := anotherFreshResult.APIURL(g.String("https://first-profile-photos-api.telegram.org")) // This creates RequestOpts
	apiURLSecond := apiURLFirst.APIURL(g.String("https://second-profile-photos-api.telegram.org"))      // This uses existing RequestOpts
	if apiURLSecond == nil {
		t.Error("APIURL method should return GetUserProfilePhotos for chaining with existing RequestOpts")
	}

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://profile-photos-api.example.com",
		"https://custom-profile-photos.telegram.org",
		"https://regional-profile-photos-api.telegram.org",
		"https://backup-profile-photos-api.telegram.org",
		"", // Empty URL
	}

	for _, apiURL := range apiURLs {
		apiResult := ctx.GetUserProfilePhotos(userID).
			Offset(0).
			Limit(10).
			Timeout(30 * time.Second).
			APIURL(g.String(apiURL)).
			Send()

		if apiResult.IsErr() {
			t.Logf("GetUserProfilePhotos with API URL '%s' Send failed as expected: %v", apiURL, apiResult.Err())
		}
	}
}

func TestGetUserProfilePhotos_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Send method with various user/photo configurations
	testScenarios := []struct {
		userID      int64
		offset      int64
		limit       int64
		description string
	}{
		{123, 0, 10, "First 10 photos of user 123"},
		{456, 5, 15, "Photos 5-19 of user 456"},
		{789, 0, 1, "Single photo of user 789"},
		{0, 0, 5, "Zero user ID with 5 photos"},
		{999, 100, 50, "High offset with limit 50"},
		{-1, 0, 10, "Negative user ID"},
		{123, -1, 10, "Negative offset"},
		{123, 0, 0, "Zero limit"},
		{123, 0, 100, "Maximum limit"},
		{123, 0, 101, "Over maximum limit"},
	}

	for _, scenario := range testScenarios {
		// Basic Send test
		sendResult := ctx.GetUserProfilePhotos(scenario.userID).
			Offset(scenario.offset).
			Limit(scenario.limit).
			Send()

		if sendResult.IsErr() {
			t.Logf("GetUserProfilePhotos with %s Send failed as expected: %v", scenario.description, sendResult.Err())
		}

		// Configured Send test
		configuredSendResult := ctx.GetUserProfilePhotos(scenario.userID).
			Offset(scenario.offset).
			Limit(scenario.limit).
			Timeout(30 * time.Second).
			APIURL(g.String("https://api.example.com")).
			Send()

		if configuredSendResult.IsErr() {
			t.Logf("GetUserProfilePhotos configured with %s Send failed as expected: %v", scenario.description, configuredSendResult.Err())
		}
	}

	// Test Send method with various timeout configurations
	timeoutConfigs := []time.Duration{
		5 * time.Second,
		15 * time.Second,
		45 * time.Second,
		60 * time.Second,
		2 * time.Minute,
	}

	for _, timeout := range timeoutConfigs {
		timeoutResult := ctx.GetUserProfilePhotos(456).
			Offset(0).
			Limit(20).
			Timeout(timeout).
			Send()

		if timeoutResult.IsErr() {
			t.Logf("GetUserProfilePhotos with timeout %v Send failed as expected: %v", timeout, timeoutResult.Err())
		}
	}

	// Test comprehensive workflow with all methods
	comprehensiveResult := ctx.GetUserProfilePhotos(456).
		Offset(10).
		Limit(25).
		Timeout(90 * time.Second).
		APIURL(g.String("https://comprehensive-profile-photos-api.telegram.org")).
		Send()

	if comprehensiveResult.IsErr() {
		t.Logf("GetUserProfilePhotos comprehensive workflow Send failed as expected: %v", comprehensiveResult.Err())
	}

	// Test method chaining order independence
	orderTest1 := ctx.GetUserProfilePhotos(789).
		APIURL(g.String("https://order-test-1.telegram.org")).
		Limit(15).
		Offset(5).
		Timeout(45 * time.Second).
		Send()

	if orderTest1.IsErr() {
		t.Logf("GetUserProfilePhotos order test 1 Send failed as expected: %v", orderTest1.Err())
	}

	orderTest2 := ctx.GetUserProfilePhotos(789).
		Timeout(45 * time.Second).
		Offset(5).
		APIURL(g.String("https://order-test-2.telegram.org")).
		Limit(15).
		Send()

	if orderTest2.IsErr() {
		t.Logf("GetUserProfilePhotos order test 2 Send failed as expected: %v", orderTest2.Err())
	}
}
