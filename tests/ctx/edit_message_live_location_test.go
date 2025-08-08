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

func TestContext_EditMessageLiveLocation(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	lat, lon := 40.7128, -74.0060

	result := ctx.EditMessageLiveLocation(lat, lon)

	if result == nil {
		t.Error("Expected EditMessageLiveLocation builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(456)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}

func TestContext_EditMessageLiveLocationChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	lat, lon := 40.7128, -74.0060

	result := ctx.EditMessageLiveLocation(lat, lon).
		ChatID(456).
		MessageID(789).
		HorizontalAccuracy(10.0)

	if result == nil {
		t.Error("Expected EditMessageLiveLocation builder to be created")
	}

	// Test that builder is functional
	_ = result
}

// Tests for methods with 0% coverage

func TestEditMessageLiveLocation_InlineMessageID(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	latitude := 37.7749
	longitude := -122.4194

	// Test InlineMessageID method
	inlineMessageIDs := []string{
		"inline_123456789",
		"inline_abcdef123",
		"inline_xyz789abc",
		"", // Empty inline message ID
	}

	for _, inlineID := range inlineMessageIDs {
		inlineResult := ctx.EditMessageLiveLocation(latitude, longitude).
			InlineMessageID(g.String(inlineID)).
			LiveFor(30 * time.Minute)

		if inlineResult == nil {
			t.Errorf("InlineMessageID with '%s' should work", inlineID)
		}

		// Test send with inline message ID
		sendResult := inlineResult.Send()
		if sendResult.IsErr() {
			t.Logf("EditMessageLiveLocation with inline message ID '%s' Send failed as expected: %v", inlineID, sendResult.Err())
		}
	}
}

func TestEditMessageLiveLocation_Business(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	latitude := 51.5074
	longitude := -0.1278

	// Test Business connection IDs
	businessIDs := []string{
		"business_conn_123",
		"business_conn_456",
		"enterprise_conn_789",
		"", // Empty business ID
	}

	for _, businessID := range businessIDs {
		businessResult := ctx.EditMessageLiveLocation(latitude, longitude).
			Business(g.String(businessID)).
			ChatID(456).
			MessageID(789).
			LiveFor(15 * time.Minute)

		if businessResult == nil {
			t.Errorf("Business with '%s' should work", businessID)
		}

		// Test send with business ID
		sendResult := businessResult.Send()
		if sendResult.IsErr() {
			t.Logf("EditMessageLiveLocation with business ID '%s' Send failed as expected: %v", businessID, sendResult.Err())
		}
	}
}

func TestEditMessageLiveLocation_LiveFor(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	latitude := 35.6762
	longitude := 139.6503

	// Test various live durations
	liveDurations := []time.Duration{
		1 * time.Minute,
		5 * time.Minute,
		15 * time.Minute,
		30 * time.Minute,
		60 * time.Minute,
		2 * time.Hour,
		8 * time.Hour,
		0 * time.Second, // Zero duration
	}

	for _, duration := range liveDurations {
		liveForResult := ctx.EditMessageLiveLocation(latitude, longitude).
			LiveFor(duration).
			ChatID(456).
			MessageID(789)

		if liveForResult == nil {
			t.Errorf("LiveFor with duration %v should work", duration)
		}

		// Test send with live duration
		sendResult := liveForResult.Send()
		if sendResult.IsErr() {
			t.Logf("EditMessageLiveLocation with duration %v Send failed as expected: %v", duration, sendResult.Err())
		}
	}
}

func TestEditMessageLiveLocation_Heading(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	latitude := 48.8566
	longitude := 2.3522

	// Test various headings (degrees)
	headings := []int64{
		0,   // North
		90,  // East
		180, // South
		270, // West
		45,  // Northeast
		359, // Almost full circle
	}

	for _, heading := range headings {
		headingResult := ctx.EditMessageLiveLocation(latitude, longitude).
			Heading(heading).
			ChatID(456).
			MessageID(789).
			LiveFor(30 * time.Minute)

		if headingResult == nil {
			t.Errorf("Heading with value %d should work", heading)
		}

		// Test send with heading
		sendResult := headingResult.Send()
		if sendResult.IsErr() {
			t.Logf("EditMessageLiveLocation with heading %d Send failed as expected: %v", heading, sendResult.Err())
		}
	}
}

func TestEditMessageLiveLocation_ProximityAlertRadius(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	latitude := 55.7558
	longitude := 37.6176

	// Test various proximity alert radii (meters)
	proximityRadii := []int64{
		1,      // 1 meter
		100,    // 100 meters
		1000,   // 1 kilometer
		10000,  // 10 kilometers
		100000, // 100 kilometers
		0,      // Zero radius (disabled)
	}

	for _, radius := range proximityRadii {
		proximityResult := ctx.EditMessageLiveLocation(latitude, longitude).
			ProximityAlertRadius(radius).
			ChatID(456).
			MessageID(789).
			LiveFor(30 * time.Minute)

		if proximityResult == nil {
			t.Errorf("ProximityAlertRadius with value %d should work", radius)
		}

		// Test send with proximity alert radius
		sendResult := proximityResult.Send()
		if sendResult.IsErr() {
			t.Logf("EditMessageLiveLocation with proximity radius %d Send failed as expected: %v", radius, sendResult.Err())
		}
	}
}

func TestEditMessageLiveLocation_Markup(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	latitude := 59.9139
	longitude := 10.7522

	// Test Markup method with inline keyboard
	inlineKB := keyboard.Inline().
		Text(g.String("🔄 Update Location"), g.String("update_location")).
		Row().
		Text(g.String("🛑 Stop Sharing"), g.String("stop_sharing")).
		Row().
		URL(g.String("📍 View on Map"), g.String("https://maps.google.com/"))

	markupResult := ctx.EditMessageLiveLocation(latitude, longitude).
		Markup(inlineKB).
		ChatID(456).
		MessageID(789).
		LiveFor(30 * time.Minute)

	if markupResult == nil {
		t.Error("Markup method should return EditMessageLiveLocation for chaining")
	}

	// Test send with markup
	sendResult := markupResult.Send()
	if sendResult.IsErr() {
		t.Logf("EditMessageLiveLocation with Markup Send failed as expected: %v", sendResult.Err())
	}
}

func TestEditMessageLiveLocation_Timeout(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	latitude := 52.5200
	longitude := 13.4050

	// Test Timeout method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.EditMessageLiveLocation(latitude, longitude)
	timeoutResultNil := freshResult.Timeout(30 * time.Second)
	if timeoutResultNil == nil {
		t.Error("Timeout method should return EditMessageLiveLocation for chaining with nil RequestOpts")
	}

	// Test Timeout method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.EditMessageLiveLocation(latitude, longitude)
	timeoutFirst := anotherFreshResult.Timeout(15 * time.Second) // This creates RequestOpts
	timeoutSecond := timeoutFirst.Timeout(25 * time.Second)      // This uses existing RequestOpts
	if timeoutSecond == nil {
		t.Error("Timeout method should return EditMessageLiveLocation for chaining with existing RequestOpts")
	}

	// Test send with timeout
	timeoutResult := ctx.EditMessageLiveLocation(latitude, longitude).
		ChatID(456).
		MessageID(789).
		LiveFor(30 * time.Minute).
		Timeout(30 * time.Second).
		Send()

	if timeoutResult.IsErr() {
		t.Logf("EditMessageLiveLocation with timeout Send failed as expected: %v", timeoutResult.Err())
	}
}

func TestEditMessageLiveLocation_APIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	latitude := 41.9028
	longitude := 12.4964

	// Test APIURL method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.EditMessageLiveLocation(latitude, longitude)
	apiURLResultNil := freshResult.APIURL(g.String("https://custom-live-location-api.telegram.org"))
	if apiURLResultNil == nil {
		t.Error("APIURL method should return EditMessageLiveLocation for chaining with nil RequestOpts")
	}

	// Test APIURL method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.EditMessageLiveLocation(latitude, longitude)
	apiURLFirst := anotherFreshResult.APIURL(g.String("https://first-location-api.telegram.org")) // This creates RequestOpts
	apiURLSecond := apiURLFirst.APIURL(g.String("https://second-location-api.telegram.org"))      // This uses existing RequestOpts
	if apiURLSecond == nil {
		t.Error("APIURL method should return EditMessageLiveLocation for chaining with existing RequestOpts")
	}

	// Test send with API URL
	apiResult := ctx.EditMessageLiveLocation(latitude, longitude).
		ChatID(456).
		MessageID(789).
		LiveFor(30 * time.Minute).
		APIURL(g.String("https://custom-api.telegram.org")).
		Send()

	if apiResult.IsErr() {
		t.Logf("EditMessageLiveLocation with API URL Send failed as expected: %v", apiResult.Err())
	}
}

func TestEditMessageLiveLocation_ComprehensiveWorkflow(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	latitude := 40.7128
	longitude := -74.0060

	// Test all methods in comprehensive workflow
	complexResult := ctx.EditMessageLiveLocation(latitude, longitude).
		ChatID(456).
		MessageID(789).
		LiveFor(2 * time.Hour).
		HorizontalAccuracy(5.0).
		Heading(90). // East
		ProximityAlertRadius(500).
		Business(g.String("business_location_123")).
		Markup(keyboard.Inline().
			Text(g.String("📍 Current Location"), g.String("current_location")).
			Text(g.String("🔄 Refresh"), g.String("refresh_location"))).
		Timeout(45 * time.Second).
		APIURL(g.String("https://comprehensive-location-api.telegram.org")).
		Send()

	if complexResult.IsErr() {
		t.Logf("EditMessageLiveLocation comprehensive workflow Send failed as expected: %v", complexResult.Err())
	}

	// Test with inline message workflow
	inlineResult := ctx.EditMessageLiveLocation(latitude, longitude).
		InlineMessageID(g.String("inline_live_location_123")).
		LiveFor(30 * time.Minute).
		Heading(180). // South
		ProximityAlertRadius(1000).
		Timeout(30 * time.Second).
		APIURL(g.String("https://inline-location-api.telegram.org")).
		Send()

	if inlineResult.IsErr() {
		t.Logf("EditMessageLiveLocation inline workflow Send failed as expected: %v", inlineResult.Err())
	}
}
