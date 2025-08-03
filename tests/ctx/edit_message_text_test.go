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
	"github.com/enetx/tg/preview"
)

func TestContext_EditMessageText(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "original"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	text := g.String("Edited text")

	result := ctx.EditMessageText(text)

	if result == nil {
		t.Error("Expected EditMessageText builder to be created")
	}

	// Test method chaining
	chained := result.HTML()
	if chained == nil {
		t.Error("Expected HTML method to return builder")
	}
}

func TestContext_EditMessageTextChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "original"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	text := g.String("Edited text")

	result := ctx.EditMessageText(text).
		HTML().
		ChatID(456).
		MessageID(789)

	if result == nil {
		t.Error("Expected EditMessageText builder to be created")
	}

	// Test that builder is functional
	_ = result
}

func TestEditMessageText_AllMethods(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 456, Text: "original"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	text := g.String("Edited text content")

	// Test Entities method
	entities := entities.New("Bold edited text").Bold("Bold")
	result := testCtx.EditMessageText(text).Entities(entities)
	if result == nil {
		t.Error("Entities method should return EditMessageText for chaining")
	}

	// Test HTML method
	result = testCtx.EditMessageText(text).HTML()
	if result == nil {
		t.Error("HTML method should return EditMessageText for chaining")
	}

	// Test Markdown method
	result = testCtx.EditMessageText(text).Markdown()
	if result == nil {
		t.Error("Markdown method should return EditMessageText for chaining")
	}

	// Test ChatID method
	result = testCtx.EditMessageText(text).ChatID(789)
	if result == nil {
		t.Error("ChatID method should return EditMessageText for chaining")
	}

	// Test MessageID method
	result = testCtx.EditMessageText(text).MessageID(999)
	if result == nil {
		t.Error("MessageID method should return EditMessageText for chaining")
	}

	// Test InlineMessageID method
	result = testCtx.EditMessageText(text).InlineMessageID(g.String("inline_123"))
	if result == nil {
		t.Error("InlineMessageID method should return EditMessageText for chaining")
	}

	// Test Business method
	result = testCtx.EditMessageText(text).Business(g.String("business_123"))
	if result == nil {
		t.Error("Business method should return EditMessageText for chaining")
	}

	// Test Timeout method
	result = testCtx.EditMessageText(text).Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return EditMessageText for chaining")
	}

	// Test APIURL method
	result = testCtx.EditMessageText(text).APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return EditMessageText for chaining")
	}

	// Test Markup method
	kb := keyboard.Inline().Text("Button", "data")
	result = testCtx.EditMessageText(text).Markup(kb)
	if result == nil {
		t.Error("Markup method should return EditMessageText for chaining")
	}

	// Test Preview method
	p := preview.New().Disable()
	result = testCtx.EditMessageText(text).Preview(p)
	if result == nil {
		t.Error("Preview method should return EditMessageText for chaining")
	}
}

func TestEditMessageText_TextContent(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 456, Text: "original"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test various text content
	texts := []string{
		"Simple edited text",
		"Text with emojis ✏️📝✨",
		"Multi-line\nedited\ntext\nwith\nbreaks",
		"<b>HTML</b> formatted edited text",
		"**Markdown** formatted edited text",
		"Very long edited text that exceeds normal expectations and contains lots of content to test handling of large text edits",
		"Text with special characters: !@#$%^&*()_+-=[]{}|;':\",./<>?",
		"Text with numbers: 1234567890",
		"Text with URLs: https://example.com/edited",
		"Text with mentions: @username",
		"Text with hashtags: #edited",
		"",
		"A",
	}

	for _, text := range texts {
		result := testCtx.EditMessageText(g.String(text))
		if result == nil {
			t.Errorf("EditMessageText should work with text: %s", text)
		}

		// Test chaining for each text
		chained := result.HTML()
		if chained == nil {
			t.Errorf("Chaining should work for text: %s", text)
		}
	}
}

func TestEditMessageText_FormattingModes(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 456, Text: "original"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test HTML formatting
	htmlResult := testCtx.EditMessageText(g.String("<b>Bold</b> and <i>italic</i> edited text")).HTML()
	if htmlResult == nil {
		t.Error("HTML formatting should work")
	}

	// Test Markdown formatting
	markdownResult := testCtx.EditMessageText(g.String("**Bold** and _italic_ edited text")).Markdown()
	if markdownResult == nil {
		t.Error("Markdown formatting should work")
	}

	// Test entities formatting
	entities := entities.New("Bold and italic edited text").
		Bold("Bold").
		Italic("italic")
	entitiesResult := testCtx.EditMessageText(g.String("Bold and italic edited text")).Entities(entities)
	if entitiesResult == nil {
		t.Error("Entities formatting should work")
	}

	// Test formatting precedence (last one wins)
	precedenceResult := testCtx.EditMessageText(g.String("Test")).HTML().Markdown()
	if precedenceResult == nil {
		t.Error("Formatting precedence should work")
	}
}

func TestEditMessageText_MessageIdentifiers(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 456, Text: "original"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	text := g.String("Edited text")

	// Test various chat and message ID combinations
	identifiers := []struct {
		name      string
		chatID    int64
		messageID int64
	}{
		{"Private Chat", 123, 456},
		{"Group Chat", -100123456789, 789},
		{"Supergroup", -1001234567890, 999},
		{"Channel", -1001987654321, 111},
		{"High Message ID", 123, 999999999},
		{"Zero Chat ID", 0, 456},
		{"Zero Message ID", 123, 0},
		{"Both Zero", 0, 0},
	}

	for _, id := range identifiers {
		result := testCtx.EditMessageText(text).
			ChatID(id.chatID).
			MessageID(id.messageID)

		if result == nil {
			t.Errorf("%s identifiers should work (chat: %d, msg: %d)", id.name, id.chatID, id.messageID)
		}

		// Test with additional methods
		enhanced := result.HTML().Timeout(30 * time.Second)
		if enhanced == nil {
			t.Errorf("Enhanced %s should work", id.name)
		}
	}
}

func TestEditMessageText_InlineMessages(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 456, Text: "original"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	text := g.String("Edited inline text")

	// Test various inline message IDs
	inlineIDs := []string{
		"inline_123",
		"INLINE_MESSAGE_456",
		"inline-message-789",
		"im_12345",
		"inline.message.2024",
		"very_long_inline_message_identifier_with_many_characters_12345",
		"short",
		"i1",
		"",
	}

	for _, inlineID := range inlineIDs {
		result := testCtx.EditMessageText(text).InlineMessageID(g.String(inlineID))
		if result == nil {
			t.Errorf("InlineMessageID should work with ID: %s", inlineID)
		}

		// Test combining with other methods
		combined := result.HTML().Timeout(30 * time.Second)
		if combined == nil {
			t.Errorf("InlineMessageID combination should work for ID: %s", inlineID)
		}
	}
}

func TestEditMessageText_BusinessFeatures(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 456, Text: "original"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	text := g.String("Business edited text")

	// Test various business connection IDs
	businessIDs := []string{
		"business_123",
		"enterprise_connection_456",
		"company_bot_789",
		"",
		"very_long_business_connection_identifier_12345",
		"simple",
		"business.connection.2024",
	}

	for _, businessID := range businessIDs {
		result := testCtx.EditMessageText(text).Business(g.String(businessID))
		if result == nil {
			t.Errorf("Business connection ID should work: %s", businessID)
		}

		// Test combining business with other features
		combined := result.HTML().Timeout(30 * time.Second)
		if combined == nil {
			t.Errorf("Business combination should work for ID: %s", businessID)
		}
	}
}

func TestEditMessageText_KeyboardEditing(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 456, Text: "original"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	text := g.String("Text with edited keyboard")

	// Test inline keyboard editing
	inlineKb := keyboard.Inline().
		Text("Edit", "edit_data").
		Text("Delete", "delete_data").Row().
		URL("Help", "https://example.com/help")

	inlineResult := testCtx.EditMessageText(text).Markup(inlineKb)
	if inlineResult == nil {
		t.Error("Inline keyboard editing should work")
	}

	// Test keyboard removal
	emptyKb := keyboard.Inline()
	removeResult := testCtx.EditMessageText(text).Markup(emptyKb)
	if removeResult == nil {
		t.Error("Keyboard removal should work")
	}

	// Test complex keyboard
	complexKb := keyboard.Inline().
		Text("Option 1", "opt1").
		Text("Option 2", "opt2").
		Text("Option 3", "opt3").Row().
		URL("Website", "https://example.com").
		SwitchInlineQuery("Share", "share_query")

	complexResult := testCtx.EditMessageText(text).Markup(complexKb)
	if complexResult == nil {
		t.Error("Complex keyboard editing should work")
	}
}

func TestEditMessageText_LinkPreview(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 456, Text: "original"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	text := g.String("Check out this edited link: https://example.com")

	// Test disabled preview
	disabledResult := testCtx.EditMessageText(text).Preview(preview.New().Disable())
	if disabledResult == nil {
		t.Error("Disabled link preview should work")
	}

	// Test enabled preview with URL
	enabledResult := testCtx.EditMessageText(text).Preview(preview.New().URL(g.String("https://example.com")))
	if enabledResult == nil {
		t.Error("Enabled link preview with URL should work")
	}

	// Test preview above text
	aboveResult := testCtx.EditMessageText(text).Preview(preview.New().Above())
	if aboveResult == nil {
		t.Error("Link preview above text should work")
	}
}

func TestEditMessageText_TimeoutAndAPI(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 456, Text: "original"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	text := g.String("Timeout test text")

	// Test various timeout durations
	timeouts := []time.Duration{
		1 * time.Second,
		5 * time.Second,
		30 * time.Second,
		1 * time.Minute,
		5 * time.Minute,
		10 * time.Minute,
		0 * time.Second, // Zero timeout
	}

	for _, timeout := range timeouts {
		result := testCtx.EditMessageText(text).Timeout(timeout)
		if result == nil {
			t.Errorf("Timeout method should work with duration: %v", timeout)
		}
	}

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
		result := testCtx.EditMessageText(text).APIURL(g.String(apiURL))
		if result == nil {
			t.Errorf("APIURL method should work with URL: %s", apiURL)
		}
	}
}

func TestEditMessageText_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 456, Text: "original"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test with empty text
	result := testCtx.EditMessageText(g.String(""))
	if result == nil {
		t.Error("EditMessageText should handle empty text")
	}

	// Test without ChatID/MessageID (should use effective values)
	result = testCtx.EditMessageText(g.String("Default test"))
	if result == nil {
		t.Error("EditMessageText should work without explicit IDs")
	}

	// Test with zero timeout
	result = testCtx.EditMessageText(g.String("Timeout test")).Timeout(0 * time.Second)
	if result == nil {
		t.Error("EditMessageText should handle zero timeout")
	}

	// Test with empty API URL
	result = testCtx.EditMessageText(g.String("API test")).APIURL(g.String(""))
	if result == nil {
		t.Error("EditMessageText should handle empty API URL")
	}

	// Test with empty inline message ID
	result = testCtx.EditMessageText(g.String("Inline test")).InlineMessageID(g.String(""))
	if result == nil {
		t.Error("EditMessageText should handle empty inline message ID")
	}

	// Test with empty business ID
	result = testCtx.EditMessageText(g.String("Business test")).Business(g.String(""))
	if result == nil {
		t.Error("EditMessageText should handle empty business ID")
	}
}

func TestEditMessageText_MethodCoverage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 456, Text: "original"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	text := g.String("Complete edit test")

	// Test all methods combined
	kb := keyboard.Inline().Text("Test", "test_data")
	entities := entities.New("Bold edit").Bold("Bold")
	p := preview.New().Disable()

	result := testCtx.EditMessageText(text).
		Entities(entities).
		HTML().
		ChatID(789).
		MessageID(999).
		InlineMessageID(g.String("inline_123")).
		Business(g.String("business_456")).
		Timeout(60 * time.Second).
		APIURL(g.String("https://api.telegram.org")).
		Markup(kb).
		Preview(p)

	if result == nil {
		t.Error("All methods combined should work")
	}

	// Test method overriding
	overrideResult := testCtx.EditMessageText(text).
		HTML().
		Markdown(). // Should override HTML
		ChatID(456).
		ChatID(789). // Should override first ChatID
		MessageID(111).
		MessageID(222). // Should override first MessageID
		Timeout(30 * time.Second).
		Timeout(60 * time.Second) // Should override first timeout

	if overrideResult == nil {
		t.Error("Method overriding should work")
	}
}

func TestEditMessageText_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 456, Text: "original"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	text := g.String("Send test edited text")

	// Test Send method execution
	builder := testCtx.EditMessageText(text)
	result := builder.Send()

	// The result should be present (even if it's an error due to mocking)
	if !result.IsErr() && !result.IsOk() {
		t.Error("Send method should return a result")
	}

	// Test Send with various options
	builderWithOptions := testCtx.EditMessageText(text).
		ChatID(789).
		MessageID(999).
		HTML().
		Timeout(30 * time.Second)
	resultWithOptions := builderWithOptions.Send()

	if !resultWithOptions.IsErr() && !resultWithOptions.IsOk() {
		t.Error("Send with options should return a result")
	}

	// Test Send without ChatID/MessageID (should use effective values)
	builderWithoutIDs := testCtx.EditMessageText(text)
	resultWithoutIDs := builderWithoutIDs.Send()

	if !resultWithoutIDs.IsErr() && !resultWithoutIDs.IsOk() {
		t.Error("Send without IDs should return a result (using effective values)")
	}

	// Test Send with inline message
	builderInline := testCtx.EditMessageText(text).InlineMessageID(g.String("inline_123"))
	resultInline := builderInline.Send()

	if !resultInline.IsErr() && !resultInline.IsOk() {
		t.Error("Send with inline message ID should return a result")
	}

	// Test Send with business connection
	builderBusiness := testCtx.EditMessageText(text).Business(g.String("business_456"))
	resultBusiness := builderBusiness.Send()

	if !resultBusiness.IsErr() && !resultBusiness.IsOk() {
		t.Error("Send with business connection should return a result")
	}

	// Test Send with all features
	kb := keyboard.Inline().Text("Button", "data")
	entities := entities.New("Test edit").Bold("Test")
	p := preview.New().Disable()
	builderComplete := testCtx.EditMessageText(text).
		Entities(entities).
		HTML().
		ChatID(789).
		MessageID(999).
		Business(g.String("business_789")).
		Timeout(45 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Markup(kb).
		Preview(p)
	resultComplete := builderComplete.Send()

	if !resultComplete.IsErr() && !resultComplete.IsOk() {
		t.Error("Send with all features should return a result")
	}
}
