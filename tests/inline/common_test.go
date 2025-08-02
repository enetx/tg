package inline_test

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/inline"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/keyboard"
)

// Common test data for inline query results
var (
	testID           = g.String("test_id_123")
	testTitle        = g.String("Test Title")
	testDescription  = g.String("Test Description")
	testCaption      = g.String("Test Caption")
	testURL          = g.String("https://example.com/test")
	testFileID       = g.String("BAADBAADrwADBREAAYdaWKOKDj8X")
	testThumbnailURL = g.String("https://example.com/thumb.jpg")
)

// Helper function to create test message content
func createTestMessageContent() input.MessageContent {
	return input.Text(g.String("Test message content"))
}

// Helper function to create test inline keyboard
func createTestInlineKeyboard() *gotgbot.InlineKeyboardMarkup {
	return &gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
			{
				{
					Text: "Test Button",
					Url:  "https://example.com",
				},
			},
		},
	}
}

// Helper function to create test keyboard using TG framework
func createTestKeyboard() keyboard.Keyboard {
	return keyboard.Inline().
		URL(g.String("Test Button"), g.String("https://example.com"))
}

// Test the QueryResult interface implementation
func assertQueryResult(result inline.QueryResult) bool {
	return result != nil && result.Build() != nil
}
