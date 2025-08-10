package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
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

func TestSendVenue_After(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	lat, lon := 40.7128, -74.0060
	title := g.String("Test Venue")
	address := g.String("123 Test St")
	if ctx.SendVenue(lat, lon, title, address).After(time.Minute) == nil {
		t.Error("After should return builder")
	}
}

func TestSendVenue_DeleteAfter(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	lat, lon := 40.7128, -74.0060
	title := g.String("Test Venue")
	address := g.String("123 Test St")
	if ctx.SendVenue(lat, lon, title, address).DeleteAfter(time.Hour) == nil {
		t.Error("DeleteAfter should return builder")
	}
}

func TestSendVenue_Markup(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	lat, lon := 40.7128, -74.0060
	title := g.String("Test Venue")
	address := g.String("123 Test St")
	btn1 := keyboard.NewButton().Text(g.String("View Venue")).URL(g.String("https://maps.google.com"))
	if ctx.SendVenue(lat, lon, title, address).Markup(keyboard.Inline().Button(btn1)) == nil {
		t.Error("Markup should return builder")
	}
}

func TestSendVenue_ReplyTo(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	lat, lon := 40.7128, -74.0060
	title := g.String("Test Venue")
	address := g.String("123 Test St")
	if ctx.SendVenue(lat, lon, title, address).ReplyTo(123) == nil {
		t.Error("ReplyTo should return builder")
	}
}

func TestSendVenue_Business(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	lat, lon := 40.7128, -74.0060
	title := g.String("Test Venue")
	address := g.String("123 Test St")
	if ctx.SendVenue(lat, lon, title, address).Business(g.String("biz_123")) == nil {
		t.Error("Business should return builder")
	}
}

func TestSendVenue_Thread(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	lat, lon := 40.7128, -74.0060
	title := g.String("Test Venue")
	address := g.String("123 Test St")
	if ctx.SendVenue(lat, lon, title, address).Thread(456) == nil {
		t.Error("Thread should return builder")
	}
}

func TestSendVenue_APIURLWithExistingRequestOpts(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// First set Timeout to create RequestOpts, then test APIURL
	result := ctx.SendVenue(40.7128, -74.0060, g.String("Test Venue"), g.String("123 Test St")).
		Timeout(15 * time.Second).                         // This creates RequestOpts
		APIURL(g.String("https://custom.api.example.com")) // This should use existing RequestOpts

	if result == nil {
		t.Error("APIURL with existing RequestOpts should return builder")
	}
}
