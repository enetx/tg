package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SendVenue(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	lat, lon := 40.7128, -74.0060
	title := g.String("Test Venue")
	address := g.String("123 Test St")

	result := ctx.SendVenue(lat, lon, title, address)

	if result == nil {
		t.Error("Expected SendVenue builder to be created")
	}

	// Test method chaining
	chained := result.Silent()
	if chained == nil {
		t.Error("Expected Silent method to return builder")
	}
}

func TestContext_SendVenueChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	lat, lon := 40.7128, -74.0060
	title := g.String("Test Venue")
	address := g.String("123 Test St")

	result := ctx.SendVenue(lat, lon, title, address).
		FoursquareID(g.String("123456")).
		Silent().
		To(123)

	if result == nil {
		t.Error("Expected SendVenue builder to be created")
	}

	// Test continued chaining
	final := result.Protect()
	if final == nil {
		t.Error("Expected Protect method to return builder")
	}
}

func TestSendVenue_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	lat, lon := 40.7128, -74.0060
	title := g.String("Test Venue")
	address := g.String("123 Test St")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendVenue(lat, lon, title, address).Send()

	if sendResult.IsErr() {
		t.Logf("SendVenue Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendVenue(lat, lon, title, address).
		FoursquareID(g.String("test_foursquare_id")).
		FoursquareType(g.String("restaurant")).
		GooglePlaceID(g.String("test_google_id")).
		GooglePlaceType(g.String("establishment")).
		Silent().
		Protect().
		To(123).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendVenue configured Send failed as expected: %v", configuredSendResult.Err())
	}
}
