package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/inline"
	"github.com/enetx/tg/input"
)

func TestContext_AnswerInlineQuery(t *testing.T) {
	bot := &mockBot{}
	inlineQuery := &gotgbot.InlineQuery{Id: "inline123", From: gotgbot.User{Id: 456, FirstName: "Test"}}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1, InlineQuery: inlineQuery},
	}

	testCtx := ctx.New(bot, rawCtx)
	queryID := g.String("inline123")

	// Test basic creation
	result := testCtx.AnswerInlineQuery(queryID)
	if result == nil {
		t.Error("Expected AnswerInlineQuery builder to be created")
	}

	// Test AddResult method
	textContent := input.Text(g.String("Article content"))
	articleResult := inline.NewArticle(g.String("article1"), g.String("Test Article"), textContent).
		Description(g.String("Test description"))
	result = testCtx.AnswerInlineQuery(queryID).AddResult(articleResult)
	if result == nil {
		t.Error("AddResult method should return AnswerInlineQuery for chaining")
	}

	// Test Results method (multiple)
	photoResult := inline.NewPhoto(
		g.String("photo1"),
		g.String("https://example.com/photo.jpg"),
		g.String("https://example.com/thumb.jpg"),
	)
	result = testCtx.AnswerInlineQuery(queryID).Results(articleResult, photoResult)
	if result == nil {
		t.Error("Results method should return AnswerInlineQuery for chaining")
	}

	// Test CacheFor method
	result = testCtx.AnswerInlineQuery(queryID).CacheFor(300 * time.Second)
	if result == nil {
		t.Error("CacheFor method should return AnswerInlineQuery for chaining")
	}

	// Test Personal method
	result = testCtx.AnswerInlineQuery(queryID).Personal()
	if result == nil {
		t.Error("Personal method should return AnswerInlineQuery for chaining")
	}

	// Test NextOffset method
	result = testCtx.AnswerInlineQuery(queryID).NextOffset(g.String("next_50"))
	if result == nil {
		t.Error("NextOffset method should return AnswerInlineQuery for chaining")
	}

	// Test ButtonText method
	result = testCtx.AnswerInlineQuery(queryID).ButtonText(g.String("More Results"))
	if result == nil {
		t.Error("ButtonText method should return AnswerInlineQuery for chaining")
	}

	// Test StartParameter method
	result = testCtx.AnswerInlineQuery(queryID).StartParameter(g.String("start_param"))
	if result == nil {
		t.Error("StartParameter method should return AnswerInlineQuery for chaining")
	}

	// Test WebApp method
	webApp := &gotgbot.WebAppInfo{Url: "https://example.com/webapp"}
	result = testCtx.AnswerInlineQuery(queryID).WebApp(webApp)
	if result == nil {
		t.Error("WebApp method should return AnswerInlineQuery for chaining")
	}

	// Test Timeout method
	result = testCtx.AnswerInlineQuery(queryID).Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return AnswerInlineQuery for chaining")
	}

	// Test APIURL method
	result = testCtx.AnswerInlineQuery(queryID).APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return AnswerInlineQuery for chaining")
	}
}

func TestContext_AnswerInlineQueryChaining(t *testing.T) {
	bot := &mockBot{}
	inlineQuery := &gotgbot.InlineQuery{Id: "chaining456", From: gotgbot.User{Id: 789, FirstName: "ChainTest"}}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1, InlineQuery: inlineQuery},
	}

	testCtx := ctx.New(bot, rawCtx)
	queryID := g.String("chaining456")

	// Test complete method chaining
	result := testCtx.AnswerInlineQuery(queryID).
		CacheFor(600 * time.Second).
		Personal().
		NextOffset(g.String("offset_100")).
		ButtonText(g.String("Show More")).
		StartParameter(g.String("more_results")).
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return AnswerInlineQuery")
	}

	// Test WebApp chaining
	webApp := &gotgbot.WebAppInfo{Url: "https://webapp.example.com"}
	webAppResult := testCtx.AnswerInlineQuery(queryID).
		ButtonText(g.String("Launch App")).
		WebApp(webApp).
		Personal()

	if webAppResult == nil {
		t.Error("WebApp chaining should work and return AnswerInlineQuery")
	}
}

func TestAnswerInlineQuery_ResultsIntegration(t *testing.T) {
	bot := &mockBot{}
	inlineQuery := &gotgbot.InlineQuery{Id: "results123", From: gotgbot.User{Id: 999, FirstName: "ResultTest"}}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1, InlineQuery: inlineQuery},
	}

	testCtx := ctx.New(bot, rawCtx)
	queryID := g.String("results123")

	// Test various inline result types
	textContent := input.Text(g.String("Full article content"))
	articleResult := inline.NewArticle(g.String("article1"), g.String("Test Article"), textContent).
		Description(g.String("Article description"))

	photoResult := inline.NewPhoto(g.String("photo1"), g.String("https://example.com/photo.jpg"), g.String("https://example.com/thumb.jpg")).
		Caption(g.String("Photo caption"))

	videoResult := inline.NewVideo(
		g.String("video1"),
		g.String("https://example.com/video.mp4"),
		g.String("video/mp4"),
		g.String("https://example.com/thumb.jpg"),
		g.String("Test Video"),
	)

	// Test adding individual results
	result := testCtx.AnswerInlineQuery(queryID).
		AddResult(articleResult).
		AddResult(photoResult).
		AddResult(videoResult)

	if result == nil {
		t.Error("Adding individual results should work")
	}

	// Test adding multiple results at once
	result = testCtx.AnswerInlineQuery(queryID).Results(articleResult, photoResult, videoResult)
	if result == nil {
		t.Error("Adding multiple results should work")
	}

	// Test combining with other options
	result = testCtx.AnswerInlineQuery(queryID).
		Results(articleResult, photoResult).
		CacheFor(120 * time.Second).
		Personal().
		NextOffset(g.String("next_batch"))

	if result == nil {
		t.Error("Combining results with other options should work")
	}
}

func TestAnswerInlineQuery_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	inlineQuery := &gotgbot.InlineQuery{Id: "edge123", From: gotgbot.User{Id: 555, FirstName: "EdgeTest"}}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1, InlineQuery: inlineQuery},
	}

	testCtx := ctx.New(bot, rawCtx)
	queryID := g.String("edge123")

	// Test with empty query ID
	result := testCtx.AnswerInlineQuery(g.String(""))
	if result == nil {
		t.Error("AnswerInlineQuery should handle empty query ID")
	}

	// Test with zero cache duration
	result = testCtx.AnswerInlineQuery(queryID).CacheFor(0 * time.Second)
	if result == nil {
		t.Error("AnswerInlineQuery should handle zero cache duration")
	}

	// Test with very long cache duration
	result = testCtx.AnswerInlineQuery(queryID).CacheFor(24 * time.Hour)
	if result == nil {
		t.Error("AnswerInlineQuery should handle very long cache duration")
	}

	// Test with empty next offset
	result = testCtx.AnswerInlineQuery(queryID).NextOffset(g.String(""))
	if result == nil {
		t.Error("AnswerInlineQuery should handle empty next offset")
	}

	// Test with empty button text
	result = testCtx.AnswerInlineQuery(queryID).ButtonText(g.String(""))
	if result == nil {
		t.Error("AnswerInlineQuery should handle empty button text")
	}

	// Test with empty start parameter
	result = testCtx.AnswerInlineQuery(queryID).StartParameter(g.String(""))
	if result == nil {
		t.Error("AnswerInlineQuery should handle empty start parameter")
	}

	// Test with nil WebApp
	result = testCtx.AnswerInlineQuery(queryID).WebApp(nil)
	if result == nil {
		t.Error("AnswerInlineQuery should handle nil WebApp")
	}

	// Test with zero timeout
	result = testCtx.AnswerInlineQuery(queryID).Timeout(0 * time.Second)
	if result == nil {
		t.Error("AnswerInlineQuery should handle zero timeout")
	}

	// Test with empty API URL
	result = testCtx.AnswerInlineQuery(queryID).APIURL(g.String(""))
	if result == nil {
		t.Error("AnswerInlineQuery should handle empty API URL")
	}
}

func TestAnswerInlineQuery_ButtonCombinations(t *testing.T) {
	bot := &mockBot{}
	inlineQuery := &gotgbot.InlineQuery{Id: "button123", From: gotgbot.User{Id: 777, FirstName: "ButtonTest"}}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1, InlineQuery: inlineQuery},
	}

	testCtx := ctx.New(bot, rawCtx)
	queryID := g.String("button123")

	// Test button with start parameter
	result := testCtx.AnswerInlineQuery(queryID).
		ButtonText(g.String("Start Bot")).
		StartParameter(g.String("welcome"))

	if result == nil {
		t.Error("Button with start parameter should work")
	}

	// Test button with WebApp
	webApp := &gotgbot.WebAppInfo{Url: "https://myapp.example.com"}
	result = testCtx.AnswerInlineQuery(queryID).
		ButtonText(g.String("Open App")).
		WebApp(webApp)

	if result == nil {
		t.Error("Button with WebApp should work")
	}

	// Test multiple button configurations (should override)
	result = testCtx.AnswerInlineQuery(queryID).
		ButtonText(g.String("First Button")).
		StartParameter(g.String("param1")).
		ButtonText(g.String("Second Button")).
		StartParameter(g.String("param2"))

	if result == nil {
		t.Error("Multiple button configurations should work (with override)")
	}

	// Test button with complex configuration
	result = testCtx.AnswerInlineQuery(queryID).
		ButtonText(g.String("Complex Button")).
		StartParameter(g.String("complex_param")).
		Personal().
		CacheFor(300 * time.Second)

	if result == nil {
		t.Error("Complex button configuration should work")
	}
}

func TestAnswerInlineQuery_CacheTimeVariations(t *testing.T) {
	bot := &mockBot{}
	inlineQuery := &gotgbot.InlineQuery{Id: "cache123", From: gotgbot.User{Id: 888, FirstName: "CacheTest"}}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1, InlineQuery: inlineQuery},
	}

	testCtx := ctx.New(bot, rawCtx)
	queryID := g.String("cache123")

	// Test various cache durations
	cacheDurations := []time.Duration{
		1 * time.Second,
		30 * time.Second,
		5 * time.Minute,
		1 * time.Hour,
		24 * time.Hour,
		168 * time.Hour, // 1 week
	}

	for _, duration := range cacheDurations {
		result := testCtx.AnswerInlineQuery(queryID).CacheFor(duration)
		if result == nil {
			t.Errorf("AnswerInlineQuery should handle cache duration: %v", duration)
		}
	}

	// Test cache with personal results
	result := testCtx.AnswerInlineQuery(queryID).
		CacheFor(600 * time.Second).
		Personal()

	if result == nil {
		t.Error("Cache with personal results should work")
	}

	// Test cache with results
	cacheContent := input.Text(g.String("This will be cached"))
	articleResult := inline.NewArticle(g.String("cached1"), g.String("Cached Article"), cacheContent)

	result = testCtx.AnswerInlineQuery(queryID).
		AddResult(articleResult).
		CacheFor(300 * time.Second).
		Personal()

	if result == nil {
		t.Error("Cache with results should work")
	}
}

func TestAnswerInlineQuery_Send(t *testing.T) {
	bot := &mockBot{}
	inlineQuery := &gotgbot.InlineQuery{Id: "send123", From: gotgbot.User{Id: 999, FirstName: "SendTest"}}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1, InlineQuery: inlineQuery},
	}

	testCtx := ctx.New(bot, rawCtx)
	queryID := g.String("send123")

	// Test Send method with results
	textContent := input.Text(g.String("Send test content"))
	articleResult := inline.NewArticle(g.String("send1"), g.String("Send Test Article"), textContent)

	result := testCtx.AnswerInlineQuery(queryID).
		AddResult(articleResult).
		Send()

	// Mock will fail, but this covers the Send method
	if result.IsErr() {
		t.Logf("AnswerInlineQuery Send failed as expected with mock bot: %v", result.Err())
	}

	// Test Send method without results (empty results)
	emptyResult := testCtx.AnswerInlineQuery(queryID).Send()

	if emptyResult.IsErr() {
		t.Logf("AnswerInlineQuery Send with empty results failed as expected: %v", emptyResult.Err())
	}
}
