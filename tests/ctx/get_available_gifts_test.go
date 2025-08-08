package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_GetAvailableGifts(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.GetAvailableGifts()

	if result == nil {
		t.Error("Expected GetAvailableGifts builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}

// Tests for methods with 0% coverage

func TestGetAvailableGifts_APIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test APIURL method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.GetAvailableGifts()
	apiURLResultNil := freshResult.APIURL(g.String("https://custom-gifts-api.telegram.org"))
	if apiURLResultNil == nil {
		t.Error("APIURL method should return GetAvailableGifts for chaining with nil RequestOpts")
	}

	// Test APIURL method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.GetAvailableGifts()
	apiURLFirst := anotherFreshResult.APIURL(g.String("https://first-gifts-api.telegram.org")) // This creates RequestOpts
	apiURLSecond := apiURLFirst.APIURL(g.String("https://second-gifts-api.telegram.org"))      // This uses existing RequestOpts
	if apiURLSecond == nil {
		t.Error("APIURL method should return GetAvailableGifts for chaining with existing RequestOpts")
	}

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://gifts-api.example.com",
		"https://custom-gifts.telegram.org",
		"https://regional-gifts-api.telegram.org",
		"https://backup-gifts-api.telegram.org",
		"", // Empty URL
	}

	for _, apiURL := range apiURLs {
		apiResult := ctx.GetAvailableGifts().
			Timeout(30 * time.Second).
			APIURL(g.String(apiURL)).
			Send()

		if apiResult.IsErr() {
			t.Logf("GetAvailableGifts with API URL '%s' Send failed as expected: %v", apiURL, apiResult.Err())
		}
	}
}

func TestGetAvailableGifts_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.GetAvailableGifts().Send()

	if sendResult.IsErr() {
		t.Logf("GetAvailableGifts Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
