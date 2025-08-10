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

func TestInlineKeyboard_ButtonWithParent(t *testing.T) {
	// Test Button method and parent attachment
	keyboard := Inline()
	btn := NewButton().Text("Toggle").Callback("toggle").On("ON").Off("OFF")

	// Add button to keyboard, which should attach parent
	keyboard.Button(btn)

	// Test SetActive with parent (should trigger update method)
	btn.SetActive(true)
	if !btn.Get.IsActive() {
		t.Error("Expected button to be active after SetActive(true)")
	}

	// Test Flip with parent (should trigger update method)
	btn.Flip()
	if btn.Get.IsActive() {
		t.Error("Expected button to be inactive after Flip()")
	}

	// Test update method indirectly by checking keyboard state
	markup := keyboard.Markup().(gotgbot.InlineKeyboardMarkup)
	if len(markup.InlineKeyboard) == 0 || len(markup.InlineKeyboard[0]) == 0 {
		t.Error("Expected at least one button in keyboard")
	}

	// Verify button text shows toggle state
	btnText := markup.InlineKeyboard[0][0].Text
	if btnText != "OFF" {
		t.Errorf("Expected button text to show 'OFF', got '%s'", btnText)
	}
}

func TestInlineKeyboard_ButtonUpdateExisting(t *testing.T) {
	// Test update method by updating existing button
	keyboard := Inline().Row().Text("Original", "same_callback")

	// Create new button with same callback data
	newBtn := NewButton().Text("Updated").Callback("same_callback")

	// Add button - should update existing one
	keyboard.Button(newBtn)

	markup := keyboard.Markup().(gotgbot.InlineKeyboardMarkup)
	if len(markup.InlineKeyboard) != 1 || len(markup.InlineKeyboard[0]) != 1 {
		t.Error("Expected exactly one button after update")
	}

	if markup.InlineKeyboard[0][0].Text != "Updated" {
		t.Errorf("Expected updated button text 'Updated', got '%s'", markup.InlineKeyboard[0][0].Text)
	}
}

func TestInlineKeyboard_ButtonWithPayment(t *testing.T) {
	// Test Button method with Pay button (different code path)
	keyboard := Inline()
	btn := NewButton().Text("Pay Now").Pay()

	// Add payment button to keyboard
	keyboard.Button(btn)

	markup := keyboard.Markup().(gotgbot.InlineKeyboardMarkup)
	if len(markup.InlineKeyboard) == 0 || len(markup.InlineKeyboard[0]) == 0 {
		t.Error("Expected at least one button in keyboard")
	}

	if !markup.InlineKeyboard[0][0].Pay {
		t.Error("Expected button to have Pay flag set")
	}
}

func TestInlineKeyboard_ButtonNilHandling(t *testing.T) {
	// Test Button method with nil button
	keyboard := Inline()
	result := keyboard.Button(nil)

	if result != keyboard {
		t.Error("Expected Button(nil) to return same keyboard instance")
	}

	// Test Button method with button that has nil raw
	btn := &Button{} // Empty button without raw
	result2 := keyboard.Button(btn)

	if result2 != keyboard {
		t.Error("Expected Button with nil raw to return same keyboard instance")
	}
}

func TestInlineKeyboard_UpdateMethodDirectAccess(t *testing.T) {
	// Test update method by creating button with callback and then using SetActive
	keyboard := Inline()
	btn := NewButton().Text("Test").Callback("test_cb").On("Active").Off("Inactive")

	// Add to keyboard to attach parent
	keyboard.Button(btn)

	// Initial state - should be inactive
	btn.SetActive(false)
	markup1 := keyboard.Markup().(gotgbot.InlineKeyboardMarkup)
	if markup1.InlineKeyboard[0][0].Text != "Inactive" {
		t.Error("Expected initial state to be 'Inactive'")
	}

	// Change state - should trigger update
	btn.SetActive(true)
	markup2 := keyboard.Markup().(gotgbot.InlineKeyboardMarkup)
	if markup2.InlineKeyboard[0][0].Text != "Active" {
		t.Error("Expected state to change to 'Active' after SetActive(true)")
	}

	// Test Flip - should trigger update again
	btn.Flip()
	markup3 := keyboard.Markup().(gotgbot.InlineKeyboardMarkup)
	if markup3.InlineKeyboard[0][0].Text != "Inactive" {
		t.Error("Expected state to change to 'Inactive' after Flip()")
	}
}

// Additional tests for missing coverage
func TestInlineKeyboard_FromMarkupPointer(t *testing.T) {
	// Test fromMarkup with pointer type (covers pointer case)
	original := &gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
			{{Text: "Test Pointer", CallbackData: "test_ptr"}},
		},
	}

	keyboard := Inline(original)
	result := keyboard.Markup().(gotgbot.InlineKeyboardMarkup)

	if len(result.InlineKeyboard) != 1 {
		t.Errorf("Expected 1 row, got %d", len(result.InlineKeyboard))
	}

	if result.InlineKeyboard[0][0].Text != "Test Pointer" {
		t.Errorf("Expected button text 'Test Pointer', got '%s'", result.InlineKeyboard[0][0].Text)
	}
}

func TestInlineKeyboard_FromNilKeyboard(t *testing.T) {
	// Test fromKeyboard with nil keyboard (covers nil check)
	keyboard := Inline()
	result := keyboard.Button(nil) // This should use fromKeyboard with nil
	if result != keyboard {
		t.Error("Expected fromKeyboard with nil to return same keyboard")
	}
}

func TestInlineKeyboard_UpdateWithNilOrEmptyCallback(t *testing.T) {
	// Test update method with button without callback data
	keyboard := Inline().Row().Text("Original", "original_cb")

	// Test update with button that has nil raw
	btn := &Button{} // Empty button
	result := keyboard.Button(btn)
	if result != keyboard {
		t.Error("Expected update with nil raw to return same keyboard")
	}

	// Test update with button that has empty callback
	btn2 := NewButton().Text("No Callback")
	result2 := keyboard.Button(btn2)
	if result2 != keyboard {
		t.Error("Expected update with empty callback to return same keyboard")
	}
}

func TestInlineKeyboard_UpdateAddToNewRow(t *testing.T) {
	// Test update method when button is not found - should add to new row
	keyboard := Inline().Row().Text("Existing", "existing_cb")

	// Create button with different callback that doesn't exist
	newBtn := NewButton().Text("New Button").Callback("new_cb")
	keyboard.Button(newBtn)

	// Should add to last row since button not found
	markup := keyboard.Markup().(gotgbot.InlineKeyboardMarkup)
	if len(markup.InlineKeyboard) != 1 {
		t.Error("Expected 1 row after adding new button")
	}
	if len(markup.InlineKeyboard[0]) != 2 {
		t.Error("Expected 2 buttons in row after adding new button")
	}
}

func TestInlineKeyboard_ConstructorWithNilMarkup(t *testing.T) {
	// Test Inline constructor with various nil inputs
	keyboard1 := Inline(nil)
	if keyboard1 == nil {
		t.Error("Expected keyboard to be created with nil input")
	}

	// Test with ReplyMarkup that's not InlineKeyboardMarkup
	var wrongMarkup gotgbot.ReplyMarkup = gotgbot.ReplyKeyboardMarkup{}
	keyboard2 := Inline(wrongMarkup)
	if keyboard2 == nil {
		t.Error("Expected keyboard to be created with wrong markup type")
	}

	markup2 := keyboard2.Markup().(gotgbot.InlineKeyboardMarkup)
	if len(markup2.InlineKeyboard) != 0 {
		t.Error("Expected empty keyboard with wrong markup type")
	}
}

func TestInlineKeyboard_FromKeyboardValueType(t *testing.T) {
	// Test fromKeyboard with gotgbot.InlineKeyboardMarkup value type (not pointer)
	// Create a keyboard that will return value type from Markup()
	original := gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
			{{Text: "Value Type", CallbackData: "value_cb"}},
		},
	}

	// Test fromMarkup with value type
	target := Inline(original)
	resultMarkup := target.Markup().(gotgbot.InlineKeyboardMarkup)

	if len(resultMarkup.InlineKeyboard) != 1 {
		t.Error("Expected 1 row after copying value type markup")
	}

	if resultMarkup.InlineKeyboard[0][0].Text != "Value Type" {
		t.Error("Expected copied value type button text to match")
	}
}

func TestInlineKeyboard_FromKeyboardDefaultCase(t *testing.T) {
	// Create a mock keyboard that returns an unexpected markup type
	// This tests the default case in fromKeyboard switch statement
	// Test with non-keyboard markup (this should hit default case and return same keyboard)
	var nonKeyboardMarkup gotgbot.ReplyMarkup = gotgbot.ReplyKeyboardRemove{}
	result := Inline(nonKeyboardMarkup)

	if result == nil {
		t.Error("Expected keyboard to be created even with unsupported markup type")
	}

	markup := result.Markup().(gotgbot.InlineKeyboardMarkup)
	if len(markup.InlineKeyboard) != 0 {
		t.Error("Expected empty keyboard with unsupported markup type")
	}
}
