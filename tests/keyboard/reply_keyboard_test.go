package keyboard_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/tg/keyboard"
)

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

// Test Reply keyboard with Location button
func TestReplyKeyboard_Location(t *testing.T) {
	keyboard := Reply().
		Row().
		Location("Share Location")

	rawMarkup := keyboard.Markup()
	markup, ok := rawMarkup.(gotgbot.ReplyKeyboardMarkup)
	if !ok {
		t.Fatal("Expected ReplyKeyboardMarkup")
	}

	button := markup.Keyboard[0][0]
	if button.Text != "Share Location" {
		t.Errorf("Expected button text 'Share Location', got '%s'", button.Text)
	}

	if !button.RequestLocation {
		t.Error("Expected RequestLocation to be true")
	}
}

// Test Reply keyboard WebApp button
func TestReplyKeyboard_WebApp(t *testing.T) {
	keyboard := Reply().
		Row().
		WebApp("Open App", "https://myapp.com")

	rawMarkup := keyboard.Markup()
	markup, ok := rawMarkup.(gotgbot.ReplyKeyboardMarkup)
	if !ok {
		t.Fatal("Expected ReplyKeyboardMarkup")
	}

	button := markup.Keyboard[0][0]
	if button.Text != "Open App" {
		t.Errorf("Expected button text 'Open App', got '%s'", button.Text)
	}

	if button.WebApp == nil {
		t.Error("Expected WebApp to be set")
	} else if button.WebApp.Url != "https://myapp.com" {
		t.Errorf("Expected WebApp URL 'https://myapp.com', got '%s'", button.WebApp.Url)
	}
}

// Test Reply keyboard Poll button
func TestReplyKeyboard_Poll(t *testing.T) {
	keyboard := Reply().
		Row().
		Poll("Create Poll")

	rawMarkup := keyboard.Markup()
	markup, ok := rawMarkup.(gotgbot.ReplyKeyboardMarkup)
	if !ok {
		t.Fatal("Expected ReplyKeyboardMarkup")
	}

	button := markup.Keyboard[0][0]
	if button.Text != "Create Poll" {
		t.Errorf("Expected button text 'Create Poll', got '%s'", button.Text)
	}

	if button.RequestPoll == nil {
		t.Error("Expected RequestPoll to be set")
	}
}

// Test Reply keyboard Chat button
func TestReplyKeyboard_Chat(t *testing.T) {
	keyboard := Reply().
		Row().
		Chat("Select Chat")

	rawMarkup := keyboard.Markup()
	markup, ok := rawMarkup.(gotgbot.ReplyKeyboardMarkup)
	if !ok {
		t.Fatal("Expected ReplyKeyboardMarkup")
	}

	button := markup.Keyboard[0][0]
	if button.Text != "Select Chat" {
		t.Errorf("Expected button text 'Select Chat', got '%s'", button.Text)
	}

	if button.RequestChat == nil {
		t.Error("Expected RequestChat to be set")
	}
}

// Test Reply keyboard Users button
func TestReplyKeyboard_Users(t *testing.T) {
	keyboard := Reply().
		Row().
		Users("Select Users")

	rawMarkup := keyboard.Markup()
	markup, ok := rawMarkup.(gotgbot.ReplyKeyboardMarkup)
	if !ok {
		t.Fatal("Expected ReplyKeyboardMarkup")
	}

	button := markup.Keyboard[0][0]
	if button.Text != "Select Users" {
		t.Errorf("Expected button text 'Select Users', got '%s'", button.Text)
	}

	if button.RequestUsers == nil {
		t.Error("Expected RequestUsers to be set")
	}
}

// Test Reply keyboard multiple rows
func TestReplyKeyboard_MultipleRows(t *testing.T) {
	keyboard := Reply().
		Row().
		Text("Row 1 Btn 1").
		Text("Row 1 Btn 2").
		Row().
		Text("Row 2 Btn 1")

	rawMarkup := keyboard.Markup()
	markup, ok := rawMarkup.(gotgbot.ReplyKeyboardMarkup)
	if !ok {
		t.Fatal("Expected ReplyKeyboardMarkup")
	}

	if len(markup.Keyboard) != 2 {
		t.Errorf("Expected 2 rows, got %d", len(markup.Keyboard))
	}

	if len(markup.Keyboard[0]) != 2 {
		t.Errorf("Expected 2 buttons in first row, got %d", len(markup.Keyboard[0]))
	}

	if len(markup.Keyboard[1]) != 1 {
		t.Errorf("Expected 1 button in second row, got %d", len(markup.Keyboard[1]))
	}

	if markup.Keyboard[1][0].Text != "Row 2 Btn 1" {
		t.Errorf("Expected second row button text 'Row 2 Btn 1', got '%s'", markup.Keyboard[1][0].Text)
	}
}

// Test Reply keyboard default options
func TestReplyKeyboard_Options(t *testing.T) {
	keyboard := Reply().
		Row().
		Text("Test Button")

	rawMarkup := keyboard.Markup()
	markup, ok := rawMarkup.(gotgbot.ReplyKeyboardMarkup)
	if !ok {
		t.Fatal("Expected ReplyKeyboardMarkup")
	}

	// ReplyKeyboard by default sets ResizeKeyboard to true
	if !markup.ResizeKeyboard {
		t.Error("Expected ResizeKeyboard to be true by default")
	}

	// Other options are not set by default
	if markup.OneTimeKeyboard {
		t.Error("Expected OneTimeKeyboard to be false by default")
	}

	if markup.Selective {
		t.Error("Expected Selective to be false by default")
	}

	if markup.InputFieldPlaceholder != "" {
		t.Error("Expected InputFieldPlaceholder to be empty by default")
	}
}

// Test Reply keyboard fluent interface
func TestReplyKeyboard_FluentInterface(t *testing.T) {
	keyboard := Reply()

	// Test that all available methods return the keyboard instance for chaining
	result := keyboard.Row()
	if result != keyboard {
		t.Error("Expected Row() to return keyboard instance")
	}

	result = keyboard.Text("Test")
	if result != keyboard {
		t.Error("Expected Text() to return keyboard instance")
	}

	result = keyboard.Contact("Contact")
	if result != keyboard {
		t.Error("Expected Contact() to return keyboard instance")
	}

	result = keyboard.Location("Location")
	if result != keyboard {
		t.Error("Expected Location() to return keyboard instance")
	}

	result = keyboard.WebApp("WebApp", "https://example.com")
	if result != keyboard {
		t.Error("Expected WebApp() to return keyboard instance")
	}
}

// Test empty Reply keyboard
func TestReplyKeyboard_Empty(t *testing.T) {
	keyboard := Reply()

	rawMarkup := keyboard.Markup()
	markup, ok := rawMarkup.(gotgbot.ReplyKeyboardMarkup)
	if !ok {
		t.Fatal("Expected ReplyKeyboardMarkup")
	}

	if len(markup.Keyboard) != 0 {
		t.Errorf("Expected 0 rows in empty keyboard, got %d", len(markup.Keyboard))
	}
}
