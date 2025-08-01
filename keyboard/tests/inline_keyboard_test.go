package keyboard_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/tg/keyboard"
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

// Test WebApp button
func TestInlineKeyboard_WebApp(t *testing.T) {
	keyboard := Inline().
		Row().
		WebApp("Open App", "https://myapp.com")

	rawMarkup := keyboard.Markup()
	markup, ok := rawMarkup.(gotgbot.InlineKeyboardMarkup)
	if !ok {
		t.Fatal("Expected InlineKeyboardMarkup")
	}

	button := markup.InlineKeyboard[0][0]
	if button.Text != "Open App" {
		t.Errorf("Expected button text 'Open App', got '%s'", button.Text)
	}

	if button.WebApp == nil {
		t.Error("Expected WebApp to be set")
	} else if button.WebApp.Url != "https://myapp.com" {
		t.Errorf("Expected WebApp URL 'https://myapp.com', got '%s'", button.WebApp.Url)
	}
}

// Test LoginURL button
func TestInlineKeyboard_LoginURL(t *testing.T) {
	keyboard := Inline().
		Row().
		LoginURL("Login", "https://example.com/login")

	rawMarkup := keyboard.Markup()
	markup, ok := rawMarkup.(gotgbot.InlineKeyboardMarkup)
	if !ok {
		t.Fatal("Expected InlineKeyboardMarkup")
	}

	button := markup.InlineKeyboard[0][0]
	if button.Text != "Login" {
		t.Errorf("Expected button text 'Login', got '%s'", button.Text)
	}

	if button.LoginUrl == nil {
		t.Error("Expected LoginUrl to be set")
	} else if button.LoginUrl.Url != "https://example.com/login" {
		t.Errorf("Expected LoginUrl URL 'https://example.com/login', got '%s'", button.LoginUrl.Url)
	}
}

// Test CopyText button
func TestInlineKeyboard_CopyText(t *testing.T) {
	keyboard := Inline().
		Row().
		CopyText("Copy", "Hello World")

	rawMarkup := keyboard.Markup()
	markup, ok := rawMarkup.(gotgbot.InlineKeyboardMarkup)
	if !ok {
		t.Fatal("Expected InlineKeyboardMarkup")
	}

	button := markup.InlineKeyboard[0][0]
	if button.Text != "Copy" {
		t.Errorf("Expected button text 'Copy', got '%s'", button.Text)
	}

	if button.CopyText == nil {
		t.Error("Expected CopyText to be set")
	} else if button.CopyText.Text != "Hello World" {
		t.Errorf("Expected CopyText text 'Hello World', got '%s'", button.CopyText.Text)
	}
}

// Test Pay button
func TestInlineKeyboard_Pay(t *testing.T) {
	keyboard := Inline().
		Row().
		Pay("Pay Now")

	rawMarkup := keyboard.Markup()
	markup, ok := rawMarkup.(gotgbot.InlineKeyboardMarkup)
	if !ok {
		t.Fatal("Expected InlineKeyboardMarkup")
	}

	button := markup.InlineKeyboard[0][0]
	if button.Text != "Pay Now" {
		t.Errorf("Expected button text 'Pay Now', got '%s'", button.Text)
	}

	if !button.Pay {
		t.Error("Expected Pay to be true")
	}
}

// Test Game button
func TestInlineKeyboard_Game(t *testing.T) {
	keyboard := Inline().
		Row().
		Game("Play Game")

	rawMarkup := keyboard.Markup()
	markup, ok := rawMarkup.(gotgbot.InlineKeyboardMarkup)
	if !ok {
		t.Fatal("Expected InlineKeyboardMarkup")
	}

	button := markup.InlineKeyboard[0][0]
	if button.Text != "Play Game" {
		t.Errorf("Expected button text 'Play Game', got '%s'", button.Text)
	}

	if button.CallbackGame == nil {
		t.Error("Expected CallbackGame to be set")
	}
}

// Test SwitchInlineQuery button
func TestInlineKeyboard_SwitchInlineQuery(t *testing.T) {
	keyboard := Inline().
		Row().
		SwitchInlineQuery("Share", "test query")

	rawMarkup := keyboard.Markup()
	markup, ok := rawMarkup.(gotgbot.InlineKeyboardMarkup)
	if !ok {
		t.Fatal("Expected InlineKeyboardMarkup")
	}

	button := markup.InlineKeyboard[0][0]
	if button.Text != "Share" {
		t.Errorf("Expected button text 'Share', got '%s'", button.Text)
	}

	if button.SwitchInlineQuery == nil {
		t.Error("Expected SwitchInlineQuery to be set")
	} else if *button.SwitchInlineQuery != "test query" {
		t.Errorf("Expected SwitchInlineQuery 'test query', got '%s'", *button.SwitchInlineQuery)
	}
}

// Test SwitchInlineQueryCurrentChat button
func TestInlineKeyboard_SwitchInlineQueryCurrentChat(t *testing.T) {
	keyboard := Inline().
		Row().
		SwitchInlineQueryCurrentChat("Search Here", "search")

	rawMarkup := keyboard.Markup()
	markup, ok := rawMarkup.(gotgbot.InlineKeyboardMarkup)
	if !ok {
		t.Fatal("Expected InlineKeyboardMarkup")
	}

	button := markup.InlineKeyboard[0][0]
	if button.Text != "Search Here" {
		t.Errorf("Expected button text 'Search Here', got '%s'", button.Text)
	}

	if button.SwitchInlineQueryCurrentChat == nil {
		t.Error("Expected SwitchInlineQueryCurrentChat to be set")
	} else if *button.SwitchInlineQueryCurrentChat != "search" {
		t.Errorf("Expected SwitchInlineQueryCurrentChat 'search', got '%s'", *button.SwitchInlineQueryCurrentChat)
	}
}

// Test multiple rows
func TestInlineKeyboard_MultipleRows(t *testing.T) {
	keyboard := Inline().
		Row().
		Text("Row 1 Btn 1", "r1b1").
		Text("Row 1 Btn 2", "r1b2").
		Row().
		Text("Row 2 Btn 1", "r2b1")

	rawMarkup := keyboard.Markup()
	markup, ok := rawMarkup.(gotgbot.InlineKeyboardMarkup)
	if !ok {
		t.Fatal("Expected InlineKeyboardMarkup")
	}

	if len(markup.InlineKeyboard) != 2 {
		t.Errorf("Expected 2 rows, got %d", len(markup.InlineKeyboard))
	}

	if len(markup.InlineKeyboard[0]) != 2 {
		t.Errorf("Expected 2 buttons in first row, got %d", len(markup.InlineKeyboard[0]))
	}

	if len(markup.InlineKeyboard[1]) != 1 {
		t.Errorf("Expected 1 button in second row, got %d", len(markup.InlineKeyboard[1]))
	}

	if markup.InlineKeyboard[1][0].Text != "Row 2 Btn 1" {
		t.Errorf("Expected second row button text 'Row 2 Btn 1', got '%s'", markup.InlineKeyboard[1][0].Text)
	}
}

// Test keyboard from existing markup
func TestInlineKeyboard_FromMarkup(t *testing.T) {
	original := gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
			{{Text: "Test", CallbackData: "test"}},
		},
	}

	keyboard := Inline(&original)
	result := keyboard.Markup().(gotgbot.InlineKeyboardMarkup)

	if len(result.InlineKeyboard) != 1 {
		t.Errorf("Expected 1 row, got %d", len(result.InlineKeyboard))
	}

	if result.InlineKeyboard[0][0].Text != "Test" {
		t.Errorf("Expected button text 'Test', got '%s'", result.InlineKeyboard[0][0].Text)
	}
}

func TestInlineKeyboard_Button(t *testing.T) {
	keyboard := Inline().Row()

	// Test adding a nil button
	result := keyboard.Button(nil)
	if result != keyboard {
		t.Error("Expected fluent interface - method should return self")
	}

	// Test adding a Pay button
	payBtn := NewButton().Text("Pay").Pay()
	result = keyboard.Button(payBtn)
	if result != keyboard {
		t.Error("Expected fluent interface with pay button")
	}

	// Test adding a button without callback data
	noCallbackBtn := NewButton().Text("No Callback")
	result = keyboard.Button(noCallbackBtn)
	if result != keyboard {
		t.Error("Expected fluent interface with no callback button")
	}

	// Test adding a new button with callback data
	callbackBtn := NewButton().Text("Test").Callback("test_cb")
	result = keyboard.Button(callbackBtn)
	if result != keyboard {
		t.Error("Expected fluent interface with callback button")
	}

	// Test updating existing button with same callback data
	updatedBtn := NewButton().Text("Updated").Callback("test_cb")
	result = keyboard.Button(updatedBtn)
	if result != keyboard {
		t.Error("Expected fluent interface when updating existing button")
	}

	// Verify the button was updated
	markup := keyboard.Markup().(gotgbot.InlineKeyboardMarkup)
	found := false
	for _, row := range markup.InlineKeyboard {
		for _, btn := range row {
			if btn.CallbackData == "test_cb" && btn.Text == "Updated" {
				found = true
				break
			}
		}
	}
	if !found {
		t.Error("Expected button to be updated in keyboard")
	}
}

func TestInlineKeyboard_FromKeyboard(t *testing.T) {
	// Create source keyboard
	source := Inline().
		Row().
		Text("Source Button", "source_cb")

	// Test copying via constructor (this is the public way to copy keyboards)
	target := Inline(source)

	// Verify the keyboard was copied
	targetMarkup := target.Markup().(gotgbot.InlineKeyboardMarkup)
	sourceMarkup := source.Markup().(gotgbot.InlineKeyboardMarkup)

	if len(targetMarkup.InlineKeyboard) != len(sourceMarkup.InlineKeyboard) {
		t.Error("Expected same number of rows after copying")
	}

	if len(targetMarkup.InlineKeyboard) > 0 && len(targetMarkup.InlineKeyboard[0]) > 0 {
		if targetMarkup.InlineKeyboard[0][0].Text != "Source Button" {
			t.Error("Expected copied button text to match")
		}
		if targetMarkup.InlineKeyboard[0][0].CallbackData != "source_cb" {
			t.Error("Expected copied button callback to match")
		}
	}
}

func TestInlineKeyboard_Edit(t *testing.T) {
	keyboard := Inline().
		Row().
		Text("Button 1", "btn1").
		Text("Button 2", "btn2").
		Row().
		Text("Button 3", "btn3")

	// Test with nil handler
	result := keyboard.Edit(nil)
	if result != keyboard {
		t.Error("Expected fluent interface with nil handler")
	}

	// Test editing buttons - change all button texts to uppercase
	result = keyboard.Edit(func(btn *Button) {
		originalText := btn.Get.Text()
		btn.Text(originalText + " EDITED")
	})

	if result != keyboard {
		t.Error("Expected fluent interface - method should return self")
	}

	// Verify buttons were edited
	markup := keyboard.Markup().(gotgbot.InlineKeyboardMarkup)
	expectedTexts := []string{"Button 1 EDITED", "Button 2 EDITED", "Button 3 EDITED"}
	actualTexts := []string{}

	for _, row := range markup.InlineKeyboard {
		for _, btn := range row {
			actualTexts = append(actualTexts, btn.Text)
		}
	}

	if len(actualTexts) != len(expectedTexts) {
		t.Errorf("Expected %d buttons, got %d", len(expectedTexts), len(actualTexts))
	}

	for i, expected := range expectedTexts {
		if i < len(actualTexts) && actualTexts[i] != expected {
			t.Errorf("Expected button %d text '%s', got '%s'", i, expected, actualTexts[i])
		}
	}

	// Test deleting buttons during edit
	keyboard2 := Inline().
		Row().
		Text("Keep", "keep").
		Text("Delete", "delete")

	keyboard2.Edit(func(btn *Button) {
		if btn.Get.Callback() == "delete" {
			btn.Delete()
		}
	})

	markup2 := keyboard2.Markup().(gotgbot.InlineKeyboardMarkup)
	if len(markup2.InlineKeyboard) != 1 || len(markup2.InlineKeyboard[0]) != 1 {
		t.Error("Expected only one button to remain after deletion")
	}
	if markup2.InlineKeyboard[0][0].Text != "Keep" {
		t.Error("Expected 'Keep' button to remain")
	}

	// Test deleting all buttons in a row
	keyboard3 := Inline().
		Row().
		Text("Delete All", "delete_all")

	keyboard3.Edit(func(btn *Button) {
		btn.Delete()
	})

	markup3 := keyboard3.Markup().(gotgbot.InlineKeyboardMarkup)
	if len(markup3.InlineKeyboard) != 0 {
		t.Error("Expected no rows to remain after deleting all buttons")
	}
}
