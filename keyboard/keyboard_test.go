package keyboard

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

func TestInlineKeyboardBuilder(t *testing.T) {
	keyboard := Inline().
		Row().
		Text("Button 1", "btn1").
		Text("Button 2", "btn2")

	rawMarkup := keyboard.Markup()
	markup, ok := rawMarkup.(gotgbot.InlineKeyboardMarkup)
	if !ok {
		t.Fatal("Expected InlineKeyboardMarkup")
	}

	if len(markup.InlineKeyboard) != 1 {
		t.Errorf("Expected 1 row, got %d", len(markup.InlineKeyboard))
	}

	// Test first row
	firstRow := markup.InlineKeyboard[0]
	if len(firstRow) != 2 {
		t.Errorf("Expected 2 buttons in first row, got %d", len(firstRow))
	}

	if firstRow[0].Text != "Button 1" {
		t.Errorf("Expected first button text 'Button 1', got '%s'", firstRow[0].Text)
	}

	if firstRow[0].CallbackData != "btn1" {
		t.Errorf("Expected first button callback data 'btn1', got '%s'", firstRow[0].CallbackData)
	}

	if firstRow[1].Text != "Button 2" {
		t.Errorf("Expected second button text 'Button 2', got '%s'", firstRow[1].Text)
	}
}

func TestReplyKeyboardBuilder(t *testing.T) {
	keyboard := Reply().
		Row().
		Text("Simple Button").
		Contact("Share Contact")

	rawMarkup := keyboard.Markup()
	markup, ok := rawMarkup.(gotgbot.ReplyKeyboardMarkup)
	if !ok {
		t.Fatal("Expected ReplyKeyboardMarkup")
	}

	if len(markup.Keyboard) != 1 {
		t.Errorf("Expected 1 row, got %d", len(markup.Keyboard))
	}

	// Test first row
	firstRow := markup.Keyboard[0]
	if len(firstRow) != 2 {
		t.Errorf("Expected 2 buttons in first row, got %d", len(firstRow))
	}

	if firstRow[0].Text != "Simple Button" {
		t.Errorf("Expected first button text 'Simple Button', got '%s'", firstRow[0].Text)
	}

	if firstRow[1].Text != "Share Contact" {
		t.Errorf("Expected second button text 'Share Contact', got '%s'", firstRow[1].Text)
	}

	if !firstRow[1].RequestContact {
		t.Error("Expected second button to request contact")
	}
}

func TestInlineKeyboardURL(t *testing.T) {
	keyboard := Inline().
		Row().
		URL("Visit", "https://example.com")

	rawMarkup := keyboard.Markup()
	markup, ok := rawMarkup.(gotgbot.InlineKeyboardMarkup)
	if !ok {
		t.Fatal("Expected InlineKeyboardMarkup")
	}

	button := markup.InlineKeyboard[0][0]
	if button.Text != "Visit" {
		t.Errorf("Expected button text 'Visit', got '%s'", button.Text)
	}

	if button.Url != "https://example.com" {
		t.Errorf("Expected URL 'https://example.com', got '%s'", button.Url)
	}
}
