package keyboard_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/tg/keyboard"
)

// Test compile-time interface checks
func TestKeyboardInterfaceCompliance(t *testing.T) {
	// Test InlineKeyboard implements Keyboard interface
	var _ Keyboard = (*InlineKeyboard)(nil)

	// Test ReplyKeyboard implements Keyboard interface
	var _ Keyboard = (*ReplyKeyboard)(nil)
}

// Test Inline constructor with no arguments
func TestInline_NoArgs(t *testing.T) {
	keyboard := Inline()
	if keyboard == nil {
		t.Error("Expected keyboard to be created")
	}

	markup := keyboard.Markup()
	inlineMarkup, ok := markup.(gotgbot.InlineKeyboardMarkup)
	if !ok {
		t.Fatal("Expected InlineKeyboardMarkup")
	}

	if len(inlineMarkup.InlineKeyboard) != 0 {
		t.Error("Expected empty keyboard with no arguments")
	}
}

// Test Inline constructor with nil argument
func TestInline_NilArg(t *testing.T) {
	keyboard := Inline(nil)
	if keyboard == nil {
		t.Error("Expected keyboard to be created with nil argument")
	}

	markup := keyboard.Markup()
	inlineMarkup, ok := markup.(gotgbot.InlineKeyboardMarkup)
	if !ok {
		t.Fatal("Expected InlineKeyboardMarkup")
	}

	if len(inlineMarkup.InlineKeyboard) != 0 {
		t.Error("Expected empty keyboard with nil argument")
	}
}

// Test Inline constructor with InlineKeyboard argument
func TestInline_WithInlineKeyboard(t *testing.T) {
	source := Inline().Row().Text("Test", "test")

	keyboard := Inline(source)
	if keyboard == nil {
		t.Error("Expected keyboard to be created")
	}

	markup := keyboard.Markup()
	inlineMarkup, ok := markup.(gotgbot.InlineKeyboardMarkup)
	if !ok {
		t.Fatal("Expected InlineKeyboardMarkup")
	}

	if len(inlineMarkup.InlineKeyboard) != 1 {
		t.Error("Expected 1 row from source keyboard")
	}

	if inlineMarkup.InlineKeyboard[0][0].Text != "Test" {
		t.Error("Expected copied button text")
	}
}

// Test Inline constructor with gotgbot.InlineKeyboardMarkup pointer
func TestInline_WithInlineKeyboardMarkupPointer(t *testing.T) {
	original := &gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
			{{Text: "Original", CallbackData: "original"}},
		},
	}

	keyboard := Inline(original)
	if keyboard == nil {
		t.Error("Expected keyboard to be created")
	}

	markup := keyboard.Markup()
	inlineMarkup, ok := markup.(gotgbot.InlineKeyboardMarkup)
	if !ok {
		t.Fatal("Expected InlineKeyboardMarkup")
	}

	if len(inlineMarkup.InlineKeyboard) != 1 {
		t.Error("Expected 1 row from original markup")
	}

	if inlineMarkup.InlineKeyboard[0][0].Text != "Original" {
		t.Error("Expected imported button text")
	}
}

// Test Inline constructor with gotgbot.InlineKeyboardMarkup value
func TestInline_WithInlineKeyboardMarkupValue(t *testing.T) {
	original := gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
			{{Text: "Value", CallbackData: "value"}},
		},
	}

	keyboard := Inline(original)
	if keyboard == nil {
		t.Error("Expected keyboard to be created")
	}

	markup := keyboard.Markup()
	inlineMarkup, ok := markup.(gotgbot.InlineKeyboardMarkup)
	if !ok {
		t.Fatal("Expected InlineKeyboardMarkup")
	}

	if len(inlineMarkup.InlineKeyboard) != 1 {
		t.Error("Expected 1 row from original markup")
	}

	if inlineMarkup.InlineKeyboard[0][0].Text != "Value" {
		t.Error("Expected imported button text")
	}
}

// Test Reply constructor
func TestReply_NoArgs(t *testing.T) {
	keyboard := Reply()
	if keyboard == nil {
		t.Error("Expected keyboard to be created")
	}

	markup := keyboard.Markup()
	replyMarkup, ok := markup.(gotgbot.ReplyKeyboardMarkup)
	if !ok {
		t.Fatal("Expected ReplyKeyboardMarkup")
	}

	if len(replyMarkup.Keyboard) != 0 {
		t.Error("Expected empty keyboard")
	}
}
