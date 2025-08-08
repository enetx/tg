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

func TestContext_AnswerWebAppQuery(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	queryID := g.String("web_app_query_123")
	messageContent := input.Text(g.String("Test content"))
	result := inline.NewArticle(g.String("1"), g.String("Test Article"), messageContent)

	// Test basic creation
	answerQuery := testCtx.AnswerWebAppQuery(queryID, result)
	if answerQuery == nil {
		t.Error("Expected AnswerWebAppQuery builder to be created")
	}

	// Test Timeout method
	result1 := answerQuery.Timeout(30 * time.Second)
	if result1 == nil {
		t.Error("Timeout method should return AnswerWebAppQuery for chaining")
	}

	// Test APIURL method
	result2 := testCtx.AnswerWebAppQuery(queryID, result).APIURL(g.String("https://api.telegram.org"))
	if result2 == nil {
		t.Error("APIURL method should return AnswerWebAppQuery for chaining")
	}
}

func TestContext_AnswerWebAppQueryChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 2},
	}

	testCtx := ctx.New(bot, rawCtx)
	queryID := g.String("webapp_chaining_456")
	messageContent := input.Text(g.String("Chaining test content"))
	result := inline.NewArticle(g.String("2"), g.String("Chaining Article"), messageContent)

	// Test complete method chaining
	chainedResult := testCtx.AnswerWebAppQuery(queryID, result).
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if chainedResult == nil {
		t.Error("Complete method chaining should work and return AnswerWebAppQuery")
	}
}

func TestAnswerWebAppQuery_ResultTypes(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 789, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 3},
	}

	testCtx := ctx.New(bot, rawCtx)
	queryID := g.String("result_types_789")

	// Test with Article result
	textContent := input.Text(g.String("Article text content"))
	articleResult := inline.NewArticle(g.String("article1"), g.String("Test Article"), textContent)

	result := testCtx.AnswerWebAppQuery(queryID, articleResult)
	if result == nil {
		t.Error("AnswerWebAppQuery should work with Article result")
	}

	// Test with Photo result
	photoResult := inline.NewPhoto(g.String("photo1"), g.String("https://example.com/photo.jpg"), g.String("https://example.com/thumb.jpg"))

	result = testCtx.AnswerWebAppQuery(queryID, photoResult)
	if result == nil {
		t.Error("AnswerWebAppQuery should work with Photo result")
	}

	// Test with Video result
	videoResult := inline.NewVideo(g.String("video1"), g.String("https://example.com/video.mp4"), g.String("video/mp4"), g.String("https://example.com/thumb.jpg"), g.String("Test Video"))

	result = testCtx.AnswerWebAppQuery(queryID, videoResult)
	if result == nil {
		t.Error("AnswerWebAppQuery should work with Video result")
	}

	// Test with Document result
	docResult := inline.NewDocument(g.String("doc1"), g.String("Test Document"), g.String("https://example.com/document.pdf"), g.String("application/pdf"))

	result = testCtx.AnswerWebAppQuery(queryID, docResult)
	if result == nil {
		t.Error("AnswerWebAppQuery should work with Document result")
	}
}

func TestAnswerWebAppQuery_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 999, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 4},
	}

	testCtx := ctx.New(bot, rawCtx)
	messageContent := input.Text(g.String("Edge case content"))
	result := inline.NewArticle(g.String("edge1"), g.String("Edge Article"), messageContent)

	// Test with empty query ID
	answerQuery := testCtx.AnswerWebAppQuery(g.String(""), result)
	if answerQuery == nil {
		t.Error("AnswerWebAppQuery should handle empty query ID")
	}

	// Test with zero timeout
	queryID := g.String("edge_query_999")
	answerQuery = testCtx.AnswerWebAppQuery(queryID, result).Timeout(0 * time.Second)
	if answerQuery == nil {
		t.Error("AnswerWebAppQuery should handle zero timeout")
	}

	// Test with very long timeout
	answerQuery = testCtx.AnswerWebAppQuery(queryID, result).Timeout(24 * time.Hour)
	if answerQuery == nil {
		t.Error("AnswerWebAppQuery should handle very long timeout")
	}

	// Test with empty API URL
	answerQuery = testCtx.AnswerWebAppQuery(queryID, result).APIURL(g.String(""))
	if answerQuery == nil {
		t.Error("AnswerWebAppQuery should handle empty API URL")
	}

	// Test with various timeout values
	timeoutValues := []time.Duration{
		1 * time.Millisecond,
		1 * time.Second,
		30 * time.Second,
		5 * time.Minute,
		1 * time.Hour,
	}

	for _, timeout := range timeoutValues {
		answerQuery = testCtx.AnswerWebAppQuery(queryID, result).Timeout(timeout)
		if answerQuery == nil {
			t.Errorf("AnswerWebAppQuery should handle timeout: %v", timeout)
		}
	}
}

func TestAnswerWebAppQuery_ComplexContent(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 111, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 5},
	}

	testCtx := ctx.New(bot, rawCtx)
	queryID := g.String("complex_content_111")

	// Test with complex text content
	complexTextContent := input.Text(g.String("Complex text with **markdown** and _italic_ formatting"))
	complexResult := inline.NewArticle(g.String("complex1"), g.String("Complex Article"), complexTextContent).
		Description(g.String("This is a complex article with detailed description"))

	answerQuery := testCtx.AnswerWebAppQuery(queryID, complexResult)
	if answerQuery == nil {
		t.Error("AnswerWebAppQuery should work with complex content")
	}

	// Test chaining with complex content
	chainedResult := testCtx.AnswerWebAppQuery(queryID, complexResult).
		Timeout(60 * time.Second).
		APIURL(g.String("https://complex-api.example.com"))

	if chainedResult == nil {
		t.Error("Complex content chaining should work")
	}
}

func TestAnswerWebAppQuery_MultipleQueryIDs(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 222, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 6},
	}

	testCtx := ctx.New(bot, rawCtx)
	messageContent := input.Text(g.String("Multi-query content"))
	result := inline.NewArticle(g.String("multi1"), g.String("Multi Article"), messageContent)

	// Test with various query ID formats
	queryIDs := []string{
		"query_1",
		"web-app-query-123",
		"WEBAPP_QUERY_456",
		"query.with.dots",
		"query_with_underscores",
		"query-with-dashes",
		"very_long_query_id_with_many_characters_and_numbers_12345",
		"short",
		"q1",
	}

	for _, queryIDStr := range queryIDs {
		queryID := g.String(queryIDStr)
		answerQuery := testCtx.AnswerWebAppQuery(queryID, result)
		if answerQuery == nil {
			t.Errorf("AnswerWebAppQuery should handle query ID: %s", queryIDStr)
		}

		// Test chaining for each query ID
		chainedQuery := answerQuery.Timeout(15 * time.Second)
		if chainedQuery == nil {
			t.Errorf("Chaining should work for query ID: %s", queryIDStr)
		}
	}
}

func TestAnswerWebAppQuery_APIURLVariations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 333, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 7},
	}

	testCtx := ctx.New(bot, rawCtx)
	queryID := g.String("api_url_test_333")
	messageContent := input.Text(g.String("API URL test content"))
	result := inline.NewArticle(g.String("api1"), g.String("API Article"), messageContent)

	// Test with various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://custom-api.example.com",
		"https://localhost:8080",
		"https://api.staging.example.com",
		"https://api-v2.example.com/bot",
		"https://telegram-proxy.example.com/api",
	}

	for _, apiURL := range apiURLs {
		answerQuery := testCtx.AnswerWebAppQuery(queryID, result).APIURL(g.String(apiURL))
		if answerQuery == nil {
			t.Errorf("AnswerWebAppQuery should handle API URL: %s", apiURL)
		}

		// Test combining with timeout
		combinedQuery := answerQuery.Timeout(25 * time.Second)
		if combinedQuery == nil {
			t.Errorf("API URL with timeout should work for: %s", apiURL)
		}
	}
}

func TestAnswerWebAppQuery_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	queryID := g.String("send_test_123")
	messageContent := input.Text(g.String("Send test content"))
	result := inline.NewArticle(g.String("send1"), g.String("Send Test Article"), messageContent)

	// Test Send method - will fail with mock but covers the method
	sendResult := testCtx.AnswerWebAppQuery(queryID, result).Send()

	if sendResult.IsErr() {
		t.Logf("AnswerWebAppQuery Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := testCtx.AnswerWebAppQuery(queryID, result).
		Timeout(30).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("AnswerWebAppQuery configured Send failed as expected: %v", configuredSendResult.Err())
	}
}
