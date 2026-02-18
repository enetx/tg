package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_GetUserProfileAudios(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	c := ctx.New(bot, rawCtx)
	userID := int64(123)

	result := c.GetUserProfileAudios(userID)

	if result == nil {
		t.Error("Expected GetUserProfileAudios builder to be created")
	}
}

func TestContext_GetUserProfileAudiosChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	c := ctx.New(bot, rawCtx)
	userID := int64(123)

	result := c.GetUserProfileAudios(userID).
		Offset(0).
		Limit(10)

	if result == nil {
		t.Error("Expected GetUserProfileAudios builder with chained methods to be non-nil")
	}
}

func TestGetUserProfileAudios_Timeout(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	c := ctx.New(bot, rawCtx)
	userID := int64(123)

	// covers nil RequestOpts branch
	freshResult := c.GetUserProfileAudios(userID)
	timeoutResultNil := freshResult.Timeout(30 * time.Second)
	if timeoutResultNil == nil {
		t.Error("Timeout should return builder with nil RequestOpts")
	}

	// covers non-nil RequestOpts branch
	timeoutFirst := c.GetUserProfileAudios(userID).Timeout(15 * time.Second)
	timeoutSecond := timeoutFirst.Timeout(25 * time.Second)
	if timeoutSecond == nil {
		t.Error("Timeout should return builder with existing RequestOpts")
	}

	timeouts := []time.Duration{
		1 * time.Second,
		5 * time.Second,
		15 * time.Second,
		30 * time.Second,
		60 * time.Second,
		2 * time.Minute,
		0 * time.Second,
	}

	for _, timeout := range timeouts {
		timeoutResult := c.GetUserProfileAudios(userID).
			Offset(0).
			Limit(5).
			Timeout(timeout).
			Send()

		if timeoutResult.IsErr() {
			t.Logf("GetUserProfileAudios with timeout %v Send failed as expected: %v", timeout, timeoutResult.Err())
		}
	}
}

func TestGetUserProfileAudios_APIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	c := ctx.New(bot, rawCtx)
	userID := int64(123)

	// covers nil RequestOpts branch
	freshResult := c.GetUserProfileAudios(userID)
	apiURLResultNil := freshResult.APIURL(g.String("https://api.telegram.org"))
	if apiURLResultNil == nil {
		t.Error("APIURL should return builder with nil RequestOpts")
	}

	// covers non-nil RequestOpts branch
	apiURLFirst := c.GetUserProfileAudios(userID).APIURL(g.String("https://first.telegram.org"))
	apiURLSecond := apiURLFirst.APIURL(g.String("https://second.telegram.org"))
	if apiURLSecond == nil {
		t.Error("APIURL should return builder with existing RequestOpts")
	}

	apiURLs := []string{
		"https://api.telegram.org",
		"https://custom.example.com",
		"https://regional.telegram.org",
		"",
	}

	for _, apiURL := range apiURLs {
		apiResult := c.GetUserProfileAudios(userID).
			Offset(0).
			Limit(10).
			Timeout(30 * time.Second).
			APIURL(g.String(apiURL)).
			Send()

		if apiResult.IsErr() {
			t.Logf("GetUserProfileAudios with API URL '%s' Send failed as expected: %v", apiURL, apiResult.Err())
		}
	}
}

func TestGetUserProfileAudios_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	c := ctx.New(bot, rawCtx)

	testScenarios := []struct {
		userID      int64
		offset      int64
		limit       int64
		description string
	}{
		{123, 0, 10, "First 10 audios of user 123"},
		{456, 5, 15, "Audios 5-19 of user 456"},
		{789, 0, 1, "Single audio of user 789"},
		{0, 0, 5, "Zero user ID"},
		{999, 100, 50, "High offset with limit 50"},
		{-1, 0, 10, "Negative user ID"},
		{123, -1, 10, "Negative offset"},
		{123, 0, 0, "Zero limit"},
		{123, 0, 100, "Maximum limit"},
	}

	for _, scenario := range testScenarios {
		sendResult := c.GetUserProfileAudios(scenario.userID).
			Offset(scenario.offset).
			Limit(scenario.limit).
			Send()

		if sendResult.IsErr() {
			t.Logf("GetUserProfileAudios with %s Send failed as expected: %v", scenario.description, sendResult.Err())
		}

		configuredSendResult := c.GetUserProfileAudios(scenario.userID).
			Offset(scenario.offset).
			Limit(scenario.limit).
			Timeout(30 * time.Second).
			APIURL(g.String("https://api.example.com")).
			Send()

		if configuredSendResult.IsErr() {
			t.Logf(
				"GetUserProfileAudios configured with %s Send failed as expected: %v",
				scenario.description,
				configuredSendResult.Err(),
			)
		}
	}

	// Test comprehensive workflow
	comprehensiveResult := c.GetUserProfileAudios(456).
		Offset(10).
		Limit(25).
		Timeout(90 * time.Second).
		APIURL(g.String("https://api.telegram.org")).
		Send()

	if comprehensiveResult.IsErr() {
		t.Logf("GetUserProfileAudios comprehensive workflow Send failed as expected: %v", comprehensiveResult.Err())
	}

	// Test method chaining order independence
	orderTest1 := c.GetUserProfileAudios(789).
		APIURL(g.String("https://order-test-1.telegram.org")).
		Limit(15).
		Offset(5).
		Timeout(45 * time.Second).
		Send()

	if orderTest1.IsErr() {
		t.Logf("GetUserProfileAudios order test 1 Send failed as expected: %v", orderTest1.Err())
	}

	orderTest2 := c.GetUserProfileAudios(789).
		Timeout(45 * time.Second).
		Offset(5).
		APIURL(g.String("https://order-test-2.telegram.org")).
		Limit(15).
		Send()

	if orderTest2.IsErr() {
		t.Logf("GetUserProfileAudios order test 2 Send failed as expected: %v", orderTest2.Err())
	}
}
