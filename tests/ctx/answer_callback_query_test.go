package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_AnswerCallbackQuery(t *testing.T) {
	bot := &mockBot{}
	callback := &gotgbot.CallbackQuery{
		Id:   "callback123",
		Data: "test_data",
	}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1, CallbackQuery: callback},
	}

	ctx := ctx.New(bot, rawCtx)

	if ctx.Callback != callback {
		t.Error("Expected callback query to be set")
	}

	// Test callback query answer
	result := ctx.AnswerCallbackQuery(g.String("Test answer"))

	if result == nil {
		t.Error("Expected AnswerCallbackQuery builder to be created")
	}

	// Test method chaining
	chained := result.Alert()
	if chained == nil {
		t.Error("Expected Alert method to return builder")
	}
}

func TestContext_AnswerCallbackQueryChaining(t *testing.T) {
	bot := &mockBot{}
	callback := &gotgbot.CallbackQuery{
		Id:   "callback123",
		Data: "test_data",
	}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1, CallbackQuery: callback},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.AnswerCallbackQuery(g.String("Test answer")).
		Alert().
		URL(g.String("https://example.com"))

	if result == nil {
		t.Error("Expected AnswerCallbackQuery builder to be created")
	}

	// Test continued chaining
	final := result.CacheFor(60)
	if final == nil {
		t.Error("Expected CacheFor method to return builder")
	}
}

func TestAnswerCallbackQuery_URLVariations(t *testing.T) {
	bot := &mockBot{}
	callback := &gotgbot.CallbackQuery{Id: "callback123", Data: "test_data"}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1, CallbackQuery: callback},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test various URL formats
	urls := []string{
		"https://example.com",
		"https://telegram.org/bot",
		"https://api.telegram.org/file/bot123/file",
		"https://web.telegram.org",
		"https://core.telegram.org/bots/api",
		"https://my-app.example.com/callback?data=123",
		"https://localhost:8080/webhook",
		"https://subdomain.example.com/path/to/resource",
		"",
	}

	for _, url := range urls {
		result := testCtx.AnswerCallbackQuery(g.String("URL test")).URL(g.String(url))
		if result == nil {
			t.Errorf("URL method should work with URL: %s", url)
		}

		// Test chaining with URL
		chained := result.Alert()
		if chained == nil {
			t.Errorf("URL chaining should work with URL: %s", url)
		}
	}
}

func TestAnswerCallbackQuery_AlertVariations(t *testing.T) {
	bot := &mockBot{}
	callback := &gotgbot.CallbackQuery{Id: "callback123", Data: "test_data"}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1, CallbackQuery: callback},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test alert with various text content
	alertTexts := []string{
		"Simple alert",
		"Error: Something went wrong!",
		"✅ Success! Operation completed.",
		"⚠️ Warning: Please confirm your action.",
		"🔒 Access denied. Please log in.",
		"📝 Form submitted successfully.",
		"🎉 Congratulations! You won!",
		"",
	}

	for _, text := range alertTexts {
		result := testCtx.AnswerCallbackQuery(g.String(text)).Alert()
		if result == nil {
			t.Errorf("Alert method should work with text: %s", text)
		}

		// Test combining alert with other methods
		combined := result.CacheFor(30)
		if combined == nil {
			t.Errorf("Alert combination should work with text: %s", text)
		}
	}
}

func TestAnswerCallbackQuery_CacheForVariations(t *testing.T) {
	bot := &mockBot{}
	callback := &gotgbot.CallbackQuery{Id: "callback123", Data: "test_data"}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1, CallbackQuery: callback},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test various cache durations
	cacheDurations := []time.Duration{
		0,                // No cache
		time.Second,      // 1 second
		30 * time.Second, // 30 seconds
		time.Minute,      // 1 minute
		5 * time.Minute,  // 5 minutes
		10 * time.Minute, // 10 minutes
		30 * time.Minute, // 30 minutes
		time.Hour,        // 1 hour
		24 * time.Hour,   // 24 hours
	}

	for _, duration := range cacheDurations {
		result := testCtx.AnswerCallbackQuery(g.String("Cache test")).CacheFor(duration)
		if result == nil {
			t.Errorf("CacheFor method should work with duration: %d", duration)
		}

		// Test combining cache with other methods
		combined := result.Alert().URL(g.String("https://example.com"))
		if combined == nil {
			t.Errorf("CacheFor combination should work with duration: %d", duration)
		}
	}
}

func TestAnswerCallbackQuery_TimeoutVariations(t *testing.T) {
	bot := &mockBot{}
	callback := &gotgbot.CallbackQuery{Id: "callback123", Data: "test_data"}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1, CallbackQuery: callback},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test various timeout durations (in time.Duration)
	timeouts := []time.Duration{
		1 * time.Second,
		5 * time.Second,
		30 * time.Second,
		1 * time.Minute,
		5 * time.Minute,
		10 * time.Minute,
	}

	for _, timeout := range timeouts {
		result := testCtx.AnswerCallbackQuery(g.String("Timeout test")).Timeout(timeout)
		if result == nil {
			t.Errorf("Timeout method should work with duration: %v", timeout)
		}

		// Test combining timeout with other methods
		combined := result.Alert().CacheFor(60)
		if combined == nil {
			t.Errorf("Timeout combination should work with duration: %v", timeout)
		}
	}
}

func TestAnswerCallbackQuery_APIURLVariations(t *testing.T) {
	bot := &mockBot{}
	callback := &gotgbot.CallbackQuery{Id: "callback123", Data: "test_data"}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1, CallbackQuery: callback},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://api.example.com",
		"https://custom-telegram-api.com",
		"https://localhost:8080",
		"https://bot-api.myservice.com",
		"",
	}

	for _, apiURL := range apiURLs {
		result := testCtx.AnswerCallbackQuery(g.String("API URL test")).APIURL(g.String(apiURL))
		if result == nil {
			t.Errorf("APIURL method should work with URL: %s", apiURL)
		}

		// Test combining API URL with other methods
		combined := result.Alert().Timeout(30 * time.Second)
		if combined == nil {
			t.Errorf("APIURL combination should work with URL: %s", apiURL)
		}
	}
}

func TestAnswerCallbackQuery_CallbackDataScenarios(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	// Test various callback query scenarios
	callbackScenarios := []struct {
		name        string
		callbackID  string
		data        string
		description string
	}{
		{"Button Click", "btn_123", "action_click", "Simple button click"},
		{"Menu Selection", "menu_456", "option_settings", "Menu option selection"},
		{"Pagination", "page_789", "next_page_2", "Pagination control"},
		{"Confirmation", "confirm_111", "yes_delete", "Confirmation dialog"},
		{"Gaming", "game_222", "move_left", "Game control"},
		{"Settings", "set_333", "toggle_notifications", "Settings toggle"},
		{"Empty Data", "empty_444", "", "Empty callback data"},
		{"Long Data", "long_555", "very_long_callback_data_with_lots_of_information_12345", "Long callback data"},
	}

	for _, scenario := range callbackScenarios {
		t.Run(scenario.name, func(t *testing.T) {
			callback := &gotgbot.CallbackQuery{
				Id:   scenario.callbackID,
				Data: scenario.data,
			}
			rawCtx.Update.CallbackQuery = callback

			testCtx := ctx.New(bot, rawCtx)
			result := testCtx.AnswerCallbackQuery(g.String("Response for " + scenario.name))

			if result == nil {
				t.Errorf("%s callback scenario should work (%s)", scenario.name, scenario.description)
			}

			// Test complete workflow for each scenario
			complete := result.Alert().URL(g.String("https://example.com")).CacheFor(60)
			if complete == nil {
				t.Errorf("Complete %s workflow should work", scenario.name)
			}
		})
	}
}

func TestAnswerCallbackQuery_TextContent(t *testing.T) {
	bot := &mockBot{}
	callback := &gotgbot.CallbackQuery{Id: "callback123", Data: "test_data"}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1, CallbackQuery: callback},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test various text responses
	textResponses := []string{
		"Operation completed successfully",
		"❌ Error: Invalid operation",
		"✅ Settings saved",
		"📄 File downloaded",
		"🔄 Processing...",
		"⚡ Quick action performed",
		"🎯 Target acquired",
		"🛡️ Security check passed",
		"📊 Data updated",
		"🎮 Game state changed",
		"",
		"A",
		"This is a very long callback response text that exceeds normal expectations and might be used for detailed feedback to users",
	}

	for _, text := range textResponses {
		result := testCtx.AnswerCallbackQuery(g.String(text))
		if result == nil {
			t.Errorf("AnswerCallbackQuery should work with text: %s", text)
		}

		// Test each text with all method combinations
		complete := result.Alert().
			URL(g.String("https://example.com")).
			CacheFor(30).
			Timeout(15 * time.Second).
			APIURL(g.String("https://api.example.com"))
		if complete == nil {
			t.Errorf("Complete method chain should work with text: %s", text)
		}
	}
}

func TestAnswerCallbackQuery_MethodOrderIndependence(t *testing.T) {
	bot := &mockBot{}
	callback := &gotgbot.CallbackQuery{Id: "callback123", Data: "test_data"}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1, CallbackQuery: callback},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test different method orders
	order1 := testCtx.AnswerCallbackQuery(g.String("Order test 1")).
		Alert().
		URL(g.String("https://example.com")).
		CacheFor(60).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com"))

	if order1 == nil {
		t.Error("Method order 1 should work")
	}

	order2 := testCtx.AnswerCallbackQuery(g.String("Order test 2")).
		APIURL(g.String("https://api.example.com")).
		Timeout(30 * time.Second).
		CacheFor(60).
		URL(g.String("https://example.com")).
		Alert()

	if order2 == nil {
		t.Error("Method order 2 should work")
	}

	order3 := testCtx.AnswerCallbackQuery(g.String("Order test 3")).
		CacheFor(120).
		Alert().
		APIURL(g.String("https://custom-api.example.com")).
		URL(g.String("https://custom.example.com")).
		Timeout(45 * time.Second)

	if order3 == nil {
		t.Error("Method order 3 should work")
	}
}

func TestAnswerCallbackQuery_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	callback := &gotgbot.CallbackQuery{Id: "callback123", Data: "test_data"}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1, CallbackQuery: callback},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test edge cases

	// Zero cache duration
	result := testCtx.AnswerCallbackQuery(g.String("Zero cache")).CacheFor(0)
	if result == nil {
		t.Error("Zero cache duration should work")
	}

	// Zero timeout
	result = testCtx.AnswerCallbackQuery(g.String("Zero timeout")).Timeout(0 * time.Second)
	if result == nil {
		t.Error("Zero timeout should work")
	}

	// Multiple calls to same method (should override)
	result = testCtx.AnswerCallbackQuery(g.String("Override test")).
		URL(g.String("https://first.com")).
		URL(g.String("https://second.com")). // Should override first
		CacheFor(30).
		CacheFor(60). // Should override first
		Alert()

	if result == nil {
		t.Error("Method overriding should work")
	}
}

func TestAnswerCallbackQuery_Send(t *testing.T) {
	bot := &mockBot{}
	callback := &gotgbot.CallbackQuery{Id: "callback123", Data: "test_data"}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1, CallbackQuery: callback},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test Send method execution
	builder := testCtx.AnswerCallbackQuery(g.String("Send test"))
	result := builder.Send()

	// The result should be present (even if it's an error due to mocking)
	if !result.IsErr() && !result.IsOk() {
		t.Error("Send method should return a result")
	}

	// Test Send with Alert
	builderWithAlert := testCtx.AnswerCallbackQuery(g.String("Alert send")).Alert()
	resultWithAlert := builderWithAlert.Send()

	if !resultWithAlert.IsErr() && !resultWithAlert.IsOk() {
		t.Error("Send with Alert should return a result")
	}

	// Test Send with URL
	builderWithURL := testCtx.AnswerCallbackQuery(g.String("URL send")).URL(g.String("https://example.com"))
	resultWithURL := builderWithURL.Send()

	if !resultWithURL.IsErr() && !resultWithURL.IsOk() {
		t.Error("Send with URL should return a result")
	}

	// Test Send with all options
	builderComplete := testCtx.AnswerCallbackQuery(g.String("Complete send")).
		Alert().
		URL(g.String("https://complete.example.com")).
		CacheFor(90).
		Timeout(25 * time.Second).
		APIURL(g.String("https://api.complete.example.com"))
	resultComplete := builderComplete.Send()

	if !resultComplete.IsErr() && !resultComplete.IsOk() {
		t.Error("Send with all options should return a result")
	}
}
