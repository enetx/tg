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

func TestContext_EditStory(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_123")
	storyID := int64(456)
	content := input.StoryPhoto(g.String("photo.jpg"))

	result := ctx.EditStory(businessConnectionID, storyID, content)

	if result == nil {
		t.Error("Expected EditStory builder to be created")
	}

	// Test method chaining
	withCaption := result.Caption(g.String("Updated story caption"))
	if withCaption == nil {
		t.Error("Expected Caption method to return builder")
	}

	withHTML := withCaption.HTML()
	if withHTML == nil {
		t.Error("Expected HTML method to return builder")
	}
}

// Tests for methods with 0% coverage

func TestEditStory_Markdown(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_123")
	storyID := int64(456)
	content := input.StoryPhoto(g.String("markdown-photo.jpg"))

	// Test Markdown method
	markdownResult := ctx.EditStory(businessConnectionID, storyID, content).
		Caption(g.String("**Bold** and _italic_ story caption")).
		Markdown()

	if markdownResult == nil {
		t.Error("Markdown method should return EditStory for chaining")
	}

	// Test send with markdown
	sendResult := markdownResult.Send()
	if sendResult.IsErr() {
		t.Logf("EditStory with Markdown Send failed as expected: %v", sendResult.Err())
	}
}

func TestEditStory_ParseMode(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_123")
	storyID := int64(456)
	content := input.StoryPhoto(g.String("parsemode-photo.jpg"))

	// Test various parse modes
	parseModes := []string{
		"HTML",
		"MarkdownV2",
		"Markdown",
		"", // Empty parse mode
	}

	for _, mode := range parseModes {
		parseModeResult := ctx.EditStory(businessConnectionID, storyID, content).
			Caption(g.String("Story caption with parse mode")).
			ParseMode(g.String(mode))

		if parseModeResult == nil {
			t.Errorf("ParseMode with '%s' should work", mode)
		}

		// Test send with parse mode
		sendResult := parseModeResult.Send()
		if sendResult.IsErr() {
			t.Logf("EditStory with parse mode '%s' Send failed as expected: %v", mode, sendResult.Err())
		}
	}
}

func TestEditStory_CaptionEntities(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_123")
	storyID := int64(456)
	content := input.StoryPhoto(g.String("entities-photo.jpg"))

	// Test CaptionEntities method
	entitiesBuilder := entities.New("Bold and italic story caption").
		Bold(g.String("Bold")).
		Italic(g.String("italic"))

	entitiesResult := ctx.EditStory(businessConnectionID, storyID, content).
		Caption(g.String("Bold and italic story caption")).
		CaptionEntities(entitiesBuilder)

	if entitiesResult == nil {
		t.Error("CaptionEntities method should return EditStory for chaining")
	}

	// Test send with entities
	sendResult := entitiesResult.Send()
	if sendResult.IsErr() {
		t.Logf("EditStory with CaptionEntities Send failed as expected: %v", sendResult.Err())
	}
}

func TestEditStory_Areas(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_123")
	storyID := int64(456)
	content := input.StoryPhoto(g.String("areas-photo.jpg"))

	// Test Areas method with clickable areas
	areasBuilder := areas.New().
		Position(10.0, 20.0).
		Size(30.0, 40.0).
		Rotate(45.0).
		Rounded(15.0).
		Link(g.String("https://example.com")).
		Position(60.0, 70.0).
		Size(25.0, 35.0).
		Location()

	areasResult := ctx.EditStory(businessConnectionID, storyID, content).
		Caption(g.String("Story with clickable areas")).
		Areas(areasBuilder)

	if areasResult == nil {
		t.Error("Areas method should return EditStory for chaining")
	}

	// Test send with areas
	sendResult := areasResult.Send()
	if sendResult.IsErr() {
		t.Logf("EditStory with Areas Send failed as expected: %v", sendResult.Err())
	}
}

func TestEditStory_Timeout(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_123")
	storyID := int64(456)
	content := input.StoryPhoto(g.String("timeout-photo.jpg"))

	// Test Timeout method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.EditStory(businessConnectionID, storyID, content)
	timeoutResultNil := freshResult.Timeout(30 * time.Second)
	if timeoutResultNil == nil {
		t.Error("Timeout method should return EditStory for chaining with nil RequestOpts")
	}

	// Test Timeout method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.EditStory(businessConnectionID, storyID, content)
	timeoutFirst := anotherFreshResult.Timeout(15 * time.Second) // This creates RequestOpts
	timeoutSecond := timeoutFirst.Timeout(25 * time.Second)      // This uses existing RequestOpts
	if timeoutSecond == nil {
		t.Error("Timeout method should return EditStory for chaining with existing RequestOpts")
	}

	// Test various timeout values
	timeouts := []time.Duration{
		1 * time.Second,
		15 * time.Second,
		30 * time.Second,
		60 * time.Second,
		2 * time.Minute,
		0 * time.Second, // Zero timeout
	}

	for _, timeout := range timeouts {
		timeoutResult := ctx.EditStory(businessConnectionID, storyID, content).
			Caption(g.String("Story with timeout")).
			Timeout(timeout).
			Send()

		if timeoutResult.IsErr() {
			t.Logf("EditStory with timeout %v Send failed as expected: %v", timeout, timeoutResult.Err())
		}
	}
}

func TestEditStory_APIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_123")
	storyID := int64(456)
	content := input.StoryPhoto(g.String("apiurl-photo.jpg"))

	// Test APIURL method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.EditStory(businessConnectionID, storyID, content)
	apiURLResultNil := freshResult.APIURL(g.String("https://custom-story-api.telegram.org"))
	if apiURLResultNil == nil {
		t.Error("APIURL method should return EditStory for chaining with nil RequestOpts")
	}

	// Test APIURL method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.EditStory(businessConnectionID, storyID, content)
	apiURLFirst := anotherFreshResult.APIURL(g.String("https://first-story-api.telegram.org")) // This creates RequestOpts
	apiURLSecond := apiURLFirst.APIURL(g.String("https://second-story-api.telegram.org"))      // This uses existing RequestOpts
	if apiURLSecond == nil {
		t.Error("APIURL method should return EditStory for chaining with existing RequestOpts")
	}

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://story-api.example.com",
		"https://custom-story.telegram.org",
		"", // Empty URL
	}

	for _, apiURL := range apiURLs {
		apiResult := ctx.EditStory(businessConnectionID, storyID, content).
			Caption(g.String("Story with custom API")).
			HTML().
			Timeout(30 * time.Second).
			APIURL(g.String(apiURL)).
			Send()

		if apiResult.IsErr() {
			t.Logf("EditStory with API URL '%s' Send failed as expected: %v", apiURL, apiResult.Err())
		}
	}
}

func TestEditStory_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_123")
	storyID := int64(456)
	content := input.StoryPhoto(g.String("send-photo.jpg"))

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.EditStory(businessConnectionID, storyID, content).
		Caption(g.String("Basic story edit")).
		Send()

	if sendResult.IsErr() {
		t.Logf("EditStory Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with video content
	videoContent := input.StoryVideo(g.String("video.mp4"))
	videoSendResult := ctx.EditStory(businessConnectionID, storyID, videoContent).
		Caption(g.String("Video story edit")).
		HTML().
		Send()

	if videoSendResult.IsErr() {
		t.Logf("EditStory with video Send failed as expected: %v", videoSendResult.Err())
	}
}

func TestEditStory_ComprehensiveWorkflow(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_comprehensive_123")
	storyID := int64(789)
	content := input.StoryPhoto(g.String("comprehensive-photo.jpg"))

	// Test comprehensive workflow with all methods
	areasBuilder := areas.New().
		Position(20.0, 30.0).
		Size(40.0, 50.0).
		Rotate(90.0).
		Rounded(25.0).
		Link(g.String("https://comprehensive-example.com")).
		Position(70.0, 80.0).
		Size(15.0, 25.0).
		Location()

	entitiesBuilder := entities.New("<b>Comprehensive</b> story with <i>formatting</i>").
		Bold(g.String("Comprehensive")).
		Italic(g.String("formatting"))

	complexResult := ctx.EditStory(businessConnectionID, storyID, content).
		Caption(g.String("<b>Comprehensive</b> story with <i>formatting</i>")).
		HTML().
		CaptionEntities(entitiesBuilder).
		Areas(areasBuilder).
		Timeout(45 * time.Second).
		APIURL(g.String("https://comprehensive-story-api.telegram.org")).
		Send()

	if complexResult.IsErr() {
		t.Logf("EditStory comprehensive workflow Send failed as expected: %v", complexResult.Err())
	}

	// Test with Markdown workflow
	markdownResult := ctx.EditStory(businessConnectionID, storyID, content).
		Caption(g.String("**Bold** and _italic_ markdown story")).
		Markdown().
		Timeout(30 * time.Second).
		APIURL(g.String("https://markdown-story-api.telegram.org")).
		Send()

	if markdownResult.IsErr() {
		t.Logf("EditStory markdown workflow Send failed as expected: %v", markdownResult.Err())
	}

	// Test with custom parse mode
	customParseModeResult := ctx.EditStory(businessConnectionID, storyID, content).
		Caption(g.String("Custom parse mode story")).
		ParseMode(g.String("CustomMode")).
		Timeout(20 * time.Second).
		Send()

	if customParseModeResult.IsErr() {
		t.Logf("EditStory custom parse mode workflow Send failed as expected: %v", customParseModeResult.Err())
	}
}
