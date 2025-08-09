package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/areas"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/input"
)

func TestContext_PostStory(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_123")
	content := input.StoryPhoto(g.String("photo.jpg"))

	result := ctx.PostStory(businessConnectionID, content)

	if result == nil {
		t.Error("Expected PostStory builder to be created")
	}

	// Test method chaining
	withCaption := result.Caption(g.String("Story caption"))
	if withCaption == nil {
		t.Error("Expected Caption method to return builder")
	}
}

func TestPostStory_HTML(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_html_123")
	content := input.StoryPhoto(g.String("photo.jpg"))

	// Test HTML method
	result := ctx.PostStory(businessConnectionID, content)
	htmlResult := result.HTML()
	if htmlResult == nil {
		t.Error("HTML method should return PostStory builder for chaining")
	}

	// Test that HTML can be chained multiple times
	chainedResult := htmlResult.HTML()
	if chainedResult == nil {
		t.Error("HTML method should support multiple chaining calls")
	}

	// Test HTML with other methods
	htmlWithOthers := ctx.PostStory(businessConnectionID, content).
		Caption(g.String("<b>Bold story caption</b>")).
		HTML().
		Protect()
	if htmlWithOthers == nil {
		t.Error("HTML method should work with other methods")
	}
}

func TestPostStory_Markdown(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_markdown_123")
	content := input.StoryPhoto(g.String("photo.jpg"))

	// Test Markdown method
	result := ctx.PostStory(businessConnectionID, content)
	markdownResult := result.Markdown()
	if markdownResult == nil {
		t.Error("Markdown method should return PostStory builder for chaining")
	}

	// Test that Markdown can be chained multiple times
	chainedResult := markdownResult.Markdown()
	if chainedResult == nil {
		t.Error("Markdown method should support multiple chaining calls")
	}

	// Test Markdown with other methods
	markdownWithOthers := ctx.PostStory(businessConnectionID, content).
		Caption(g.String("*Italic story caption*")).
		Markdown().
		PostToChatPage()
	if markdownWithOthers == nil {
		t.Error("Markdown method should work with other methods")
	}
}

func TestPostStory_CaptionEntities(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_entities_123")
	content := input.StoryPhoto(g.String("photo.jpg"))

	// Test CaptionEntities method with various entity configurations
	entityBuilders := []*entities.Entities{
		entities.New(g.String("Bold text")).Bold(g.String("Bold text")),
		entities.New(g.String("Italic text")).Italic(g.String("Italic text")),
		entities.New(g.String("Code text")).Code(g.String("Code text")),
		entities.New(g.String("Link text")).URL(g.String("Link text"), g.String("https://example.com")),
		entities.New(g.String("Mention")).Mention(g.String("Mention"), 123456),
	}

	for i, entityBuilder := range entityBuilders {
		result := ctx.PostStory(businessConnectionID, content)
		entitiesResult := result.CaptionEntities(entityBuilder)
		if entitiesResult == nil {
			t.Errorf("CaptionEntities method should return PostStory builder for chaining with entities %d", i)
		}

		// Test that CaptionEntities can be chained and overridden
		chainedResult := entitiesResult.CaptionEntities(entities.New(g.String("Another entity")).Bold(g.String("Another entity")))
		if chainedResult == nil {
			t.Errorf("CaptionEntities method should support chaining and override with entities %d", i)
		}
	}
}

func TestPostStory_Areas(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_areas_123")
	content := input.StoryPhoto(g.String("photo.jpg"))

	// Test Areas method with various area configurations
	areaBuilders := []*areas.Areas{
		areas.New().Position(0.1, 0.1).Size(0.2, 0.2).Link(g.String("https://example.com")),
		areas.New().Position(0.3, 0.3).Size(0.4, 0.4).Reaction(g.String("👍")),
		areas.New().Position(0.5, 0.5).Size(0.6, 0.6).Location(),
		areas.New().Position(0.0, 0.0).Size(1.0, 1.0).Link(g.String("https://full-area.com")),
	}

	for i, areaBuilder := range areaBuilders {
		result := ctx.PostStory(businessConnectionID, content)
		areasResult := result.Areas(areaBuilder)
		if areasResult == nil {
			t.Errorf("Areas method should return PostStory builder for chaining with areas %d", i)
		}

		// Test that Areas can be chained and overridden
		chainedResult := areasResult.Areas(areas.New().Position(0.7, 0.7).Size(0.8, 0.8).Link(g.String("https://override.com")))
		if chainedResult == nil {
			t.Errorf("Areas method should support chaining and override with areas %d", i)
		}
	}
}

func TestPostStory_ActiveFor(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_active_123")
	content := input.StoryPhoto(g.String("photo.jpg"))

	// Test ActiveFor method with various durations
	durations := []time.Duration{
		1 * time.Hour,
		6 * time.Hour,
		12 * time.Hour,
		24 * time.Hour,   // 1 day
		48 * time.Hour,   // 2 days
		168 * time.Hour,  // 7 days
		0 * time.Second,  // Zero duration
		30 * time.Second, // Short duration
	}

	for _, duration := range durations {
		result := ctx.PostStory(businessConnectionID, content)
		activeForResult := result.ActiveFor(duration)
		if activeForResult == nil {
			t.Errorf("ActiveFor method should return PostStory builder for chaining with duration %v", duration)
		}

		// Test that ActiveFor can be chained and overridden
		chainedResult := activeForResult.ActiveFor(duration + 1*time.Hour)
		if chainedResult == nil {
			t.Errorf("ActiveFor method should support chaining and override with duration %v", duration)
		}
	}
}

func TestPostStory_PostToChatPage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_chat_page_123")
	content := input.StoryPhoto(g.String("photo.jpg"))

	// Test PostToChatPage method
	result := ctx.PostStory(businessConnectionID, content)
	postToChatPageResult := result.PostToChatPage()
	if postToChatPageResult == nil {
		t.Error("PostToChatPage method should return PostStory builder for chaining")
	}

	// Test that PostToChatPage can be chained multiple times
	chainedResult := postToChatPageResult.PostToChatPage()
	if chainedResult == nil {
		t.Error("PostToChatPage method should support multiple chaining calls")
	}

	// Test PostToChatPage with other methods
	postToChatPageWithOthers := ctx.PostStory(businessConnectionID, content).
		Caption(g.String("Story that will also appear on chat page")).
		PostToChatPage().
		ActiveFor(24 * time.Hour)
	if postToChatPageWithOthers == nil {
		t.Error("PostToChatPage method should work with other methods")
	}
}

func TestPostStory_Protect(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_protect_123")
	content := input.StoryPhoto(g.String("photo.jpg"))

	// Test Protect method
	result := ctx.PostStory(businessConnectionID, content)
	protectResult := result.Protect()
	if protectResult == nil {
		t.Error("Protect method should return PostStory builder for chaining")
	}

	// Test that Protect can be chained multiple times
	chainedResult := protectResult.Protect()
	if chainedResult == nil {
		t.Error("Protect method should support multiple chaining calls")
	}

	// Test Protect with other methods
	protectWithOthers := ctx.PostStory(businessConnectionID, content).
		Caption(g.String("Protected story content")).
		Protect().
		HTML().
		PostToChatPage()
	if protectWithOthers == nil {
		t.Error("Protect method should work with other methods")
	}
}

func TestPostStory_Timeout(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_timeout_123")
	content := input.StoryPhoto(g.String("photo.jpg"))

	// Test Timeout method with various durations
	timeouts := []time.Duration{
		1 * time.Second,
		5 * time.Second,
		30 * time.Second,
		1 * time.Minute,
		5 * time.Minute,
		0 * time.Second, // Zero timeout
	}

	for _, timeout := range timeouts {
		result := ctx.PostStory(businessConnectionID, content)
		timeoutResult := result.Timeout(timeout)
		if timeoutResult == nil {
			t.Errorf("Timeout method should return PostStory builder for chaining with timeout %v", timeout)
		}

		// Test that Timeout can be chained and overridden
		chainedResult := timeoutResult.Timeout(timeout + 1*time.Second)
		if chainedResult == nil {
			t.Errorf("Timeout method should support chaining and override with timeout %v", timeout)
		}
	}
}

func TestPostStory_APIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_apiurl_123")
	content := input.StoryPhoto(g.String("photo.jpg"))

	// Test APIURL method with various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://custom.api.example.com",
		"",
		"https://api.example.com/bot",
		"http://localhost:8080",
	}

	for _, apiURL := range apiURLs {
		result := ctx.PostStory(businessConnectionID, content)
		apiURLResult := result.APIURL(g.String(apiURL))
		if apiURLResult == nil {
			t.Errorf("APIURL method should return PostStory builder for chaining with URL: %s", apiURL)
		}

		// Test that APIURL can be chained and overridden
		chainedResult := apiURLResult.APIURL(g.String("https://override.example.com"))
		if chainedResult == nil {
			t.Errorf("APIURL method should support chaining and override with URL: %s", apiURL)
		}
	}
}

func TestPostStory_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_send_123")
	content := input.StoryPhoto(g.String("photo.jpg"))

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.PostStory(businessConnectionID, content).Send()

	if sendResult.IsErr() {
		t.Logf("PostStory Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with all options
	sendWithOptionsResult := ctx.PostStory(businessConnectionID, content).
		Caption(g.String("Complete story with all options")).
		HTML().
		CaptionEntities(entities.New(g.String("Bold text")).Bold(g.String("Bold text"))).
		Areas(areas.New().Position(0.1, 0.1).Size(0.2, 0.2).Link(g.String("https://example.com"))).
		ActiveFor(24 * time.Hour).
		PostToChatPage().
		Protect().
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.telegram.org")).
		Send()

	if sendWithOptionsResult.IsErr() {
		t.Logf("PostStory Send with options failed as expected with mock bot: %v", sendWithOptionsResult.Err())
	}
}
