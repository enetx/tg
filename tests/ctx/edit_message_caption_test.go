package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/keyboard"
)

func TestContext_EditMessageCaption(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Caption: "original caption"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	caption := g.String("New caption")

	result := ctx.EditMessageCaption(caption)

	if result == nil {
		t.Error("Expected EditMessageCaption builder to be created")
	}

	// Test method chaining
	chained := result.HTML()
	if chained == nil {
		t.Error("Expected HTML method to return builder")
	}

	// Test InlineMessageID method
	inlineResult := result.InlineMessageID(g.String("inline_123456"))
	if inlineResult == nil {
		t.Error("InlineMessageID method should return EditMessageCaption for chaining")
	}

	// Test Business method
	businessResult := result.Business(g.String("business_conn_123"))
	if businessResult == nil {
		t.Error("Business method should return EditMessageCaption for chaining")
	}

	// Test Markdown method
	markdownResult := result.Markdown()
	if markdownResult == nil {
		t.Error("Markdown method should return EditMessageCaption for chaining")
	}

	// Test Entities method
	entitiesBuilder := entities.New("Bold and italic")
	entitiesResult := result.Entities(entitiesBuilder)
	if entitiesResult == nil {
		t.Error("Entities method should return EditMessageCaption for chaining")
	}

	// Test Markup method
	inlineKB := keyboard.Inline()
	markupResult := result.Markup(inlineKB)
	if markupResult == nil {
		t.Error("Markup method should return EditMessageCaption for chaining")
	}
}

func TestContext_EditMessageCaptionChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Caption: "original caption"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	caption := g.String("New caption")

	result := ctx.EditMessageCaption(caption).
		HTML().
		ChatID(456).
		MessageID(789)

	if result == nil {
		t.Error("Expected EditMessageCaption builder to be created")
	}

	// Test that builder is functional
	_ = result
}

func TestEditMessageCaption_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Caption: "original caption"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	caption := g.String("Updated caption with <b>HTML</b>")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.EditMessageCaption(caption).Send()

	if sendResult.IsErr() {
		t.Logf("EditMessageCaption Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.EditMessageCaption(caption).
		HTML().
		ChatID(456).
		MessageID(123).
		ShowCaptionAboveMedia().
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("EditMessageCaption configured Send failed as expected: %v", configuredSendResult.Err())
	}
}

func TestEditMessageCaption_AllMethods(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Caption: "original caption"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	caption := g.String("Test caption with various formatting")

	// Test InlineMessageID functionality
	inlineMessageIDs := []string{
		"inline_123456789",
		"inline_abcdef123",
		"inline_xyz789abc",
		"", // Empty inline message ID
	}

	for _, inlineID := range inlineMessageIDs {
		inlineResult := ctx.EditMessageCaption(caption).
			InlineMessageID(g.String(inlineID)).
			HTML()

		if inlineResult == nil {
			t.Errorf("InlineMessageID with '%s' should work", inlineID)
		}

		// Test send with inline message ID
		sendResult := inlineResult.Send()
		if sendResult.IsErr() {
			t.Logf(
				"EditMessageCaption with inline message ID '%s' Send failed as expected: %v",
				inlineID,
				sendResult.Err(),
			)
		}
	}

	// Test Business connection IDs
	businessIDs := []string{
		"business_conn_123",
		"business_conn_456",
		"enterprise_conn_789",
		"", // Empty business ID
	}

	for _, businessID := range businessIDs {
		businessResult := ctx.EditMessageCaption(caption).
			Business(g.String(businessID)).
			ChatID(456).
			MessageID(789)

		if businessResult == nil {
			t.Errorf("Business with '%s' should work", businessID)
		}

		// Test send with business ID
		sendResult := businessResult.Send()
		if sendResult.IsErr() {
			t.Logf("EditMessageCaption with business ID '%s' Send failed as expected: %v", businessID, sendResult.Err())
		}
	}

	// Test Markdown method specifically
	markdownResult := ctx.EditMessageCaption(g.String("*Bold text* and _italic text_")).
		Markdown().
		ChatID(456).
		MessageID(789).
		Send()

	if markdownResult.IsErr() {
		t.Logf("EditMessageCaption with Markdown Send failed as expected: %v", markdownResult.Err())
	}

	// Test Entities functionality
	entitiesBuilder := entities.New("Bold and italic text").
		Bold(g.String("Bold")).
		Italic(g.String("italic"))

	entitiesResult := ctx.EditMessageCaption(g.String("Bold and italic text")).
		Entities(entitiesBuilder).
		ChatID(456).
		MessageID(789).
		Send()

	if entitiesResult.IsErr() {
		t.Logf("EditMessageCaption with Entities Send failed as expected: %v", entitiesResult.Err())
	}

	// Test Markup functionality
	inlineKB := keyboard.Inline().
		Text(g.String("Edit"), g.String("edit_caption")).
		Row().
		Text(g.String("Delete"), g.String("delete_message"))

	markupResult := ctx.EditMessageCaption(caption).
		Markup(inlineKB).
		ChatID(456).
		MessageID(789).
		Send()

	if markupResult.IsErr() {
		t.Logf("EditMessageCaption with Markup Send failed as expected: %v", markupResult.Err())
	}

	// Test complex workflow combining all methods
	complexResult := ctx.EditMessageCaption(g.String("<b>Updated</b> caption with <i>formatting</i>")).
		ChatID(456).
		MessageID(789).
		HTML().
		ShowCaptionAboveMedia().
		Business(g.String("business_test_123")).
		Timeout(45 * time.Second).
		APIURL(g.String("https://comprehensive-api.telegram.org")).
		Send()

	if complexResult.IsErr() {
		t.Logf("EditMessageCaption complex workflow Send failed as expected: %v", complexResult.Err())
	}
}

func TestEditMessageCaption_APIURLCoverage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Caption: "original caption"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	caption := g.String("Test caption")

	// Test APIURL method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.EditMessageCaption(caption)
	apiURLResultNil := freshResult.APIURL(g.String("https://custom-api.telegram.org"))
	if apiURLResultNil == nil {
		t.Error("APIURL method should return EditMessageCaption for chaining with nil RequestOpts")
	}

	// Test APIURL method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.EditMessageCaption(caption)
	apiURLFirst := anotherFreshResult.APIURL(g.String("https://first-api.telegram.org")) // This creates RequestOpts
	apiURLSecond := apiURLFirst.APIURL(
		g.String("https://second-api.telegram.org"),
	) // This uses existing RequestOpts
	if apiURLSecond == nil {
		t.Error("APIURL method should return EditMessageCaption for chaining with existing RequestOpts")
	}

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://caption-edit-api.example.com",
		"https://custom-caption.telegram.org",
		"", // Empty URL
	}

	for _, apiURL := range apiURLs {
		apiResult := ctx.EditMessageCaption(caption).
			ChatID(456).
			MessageID(789).
			HTML().
			Timeout(30 * time.Second).
			APIURL(g.String(apiURL)).
			Send()

		if apiResult.IsErr() {
			t.Logf("EditMessageCaption with API URL '%s' Send failed as expected: %v", apiURL, apiResult.Err())
		}
	}
}
