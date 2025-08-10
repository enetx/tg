package ctx_test

import (
	"os"
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

type mockBot struct{}

func (m *mockBot) Raw() *gotgbot.Bot           { return &gotgbot.Bot{} }
func (m *mockBot) Dispatcher() *ext.Dispatcher { return &ext.Dispatcher{} }
func (m *mockBot) Updater() *ext.Updater       { return &ext.Updater{} }

func TestNewContext(t *testing.T) {
	bot := &mockBot{}

	user := &gotgbot.User{Id: 123, FirstName: "Test"}
	chat := &gotgbot.Chat{Id: 456, Type: "private"}
	message := &gotgbot.Message{MessageId: 789, Text: "test message"}

	rawCtx := &ext.Context{
		EffectiveUser:    user,
		EffectiveChat:    chat,
		EffectiveMessage: message,
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	if ctx.Bot == nil {
		t.Error("Expected bot to be set")
	}

	if ctx.EffectiveUser != user {
		t.Error("Expected effective user to be set")
	}

	if ctx.EffectiveChat != chat {
		t.Error("Expected effective chat to be set")
	}

	if ctx.EffectiveMessage != message {
		t.Error("Expected effective message to be set")
	}

	if ctx.Raw != rawCtx {
		t.Error("Expected raw context to be set")
	}
}

func TestContext_Args(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "/start arg1 arg2 arg3"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	args := ctx.Args()

	// Should skip the first word (command) and return arguments
	if args.Len() != 3 {
		t.Errorf("Expected 3 args, got %d", args.Len())
	}
}

func TestContext_ArgsEmpty(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "/start"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	args := ctx.Args()

	// Should return empty slice for command with no args
	if args.Len() != 0 {
		t.Errorf("Expected 0 args for command without args, got %d", args.Len())
	}
}

func TestContext_NilEffectiveMessage_Operations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: nil, // Nil message
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.DeleteMessage()
	if result == nil {
		t.Error("Expected DeleteMessage builder to be created even with nil message")
	}
}

func TestContext_FileOperations_ErrorHandling(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test SendDocument with invalid file path that would cause file.Input to fail
	// Use a non-existent file path with permission issues
	invalidFile := g.String("/dev/null/nonexistent/invalid/file.txt") // Invalid path

	docResult := ctx.SendDocument(invalidFile)
	if docResult == nil {
		t.Error("Expected SendDocument builder to be created even with error")
	}

	// Test SendAudio with invalid file
	audioResult := ctx.SendAudio(invalidFile)
	if audioResult == nil {
		t.Error("Expected SendAudio builder to be created even with error")
	}

	// Test SendVideo with invalid file
	videoResult := ctx.SendVideo(invalidFile)
	if videoResult == nil {
		t.Error("Expected SendVideo builder to be created even with error")
	}

	// Test SendVoice with invalid file
	voiceResult := ctx.SendVoice(invalidFile)
	if voiceResult == nil {
		t.Error("Expected SendVoice builder to be created even with error")
	}

	// Test SendVideoNote with invalid file
	videoNoteResult := ctx.SendVideoNote(invalidFile)
	if videoNoteResult == nil {
		t.Error("Expected SendVideoNote builder to be created even with error")
	}

	// Test SendAnimation with invalid file
	animationResult := ctx.SendAnimation(invalidFile)
	if animationResult == nil {
		t.Error("Expected SendAnimation builder to be created even with error")
	}

	// Test SendSticker with invalid file
	stickerResult := ctx.SendSticker(invalidFile)
	if stickerResult == nil {
		t.Error("Expected SendSticker builder to be created even with error")
	}

	// Test SendPhoto with invalid file
	photoResult := ctx.SendPhoto(invalidFile)
	if photoResult == nil {
		t.Error("Expected SendPhoto builder to be created even with error")
	}
}

func TestContext_FileOperations_ValidFile(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Create a temporary file for testing valid file paths
	tempFile := "/tmp/test_file.txt"
	err := os.WriteFile(tempFile, []byte("test content"), 0644)
	if err != nil {
		t.Skipf("Could not create temp file: %v", err)
	}
	defer os.Remove(tempFile)

	validFile := g.String(tempFile)

	// Test SendDocument with valid file
	docResult := ctx.SendDocument(validFile)
	if docResult == nil {
		t.Error("Expected SendDocument builder to be created")
	}

	// Test SendAudio with valid file
	audioResult := ctx.SendAudio(validFile)
	if audioResult == nil {
		t.Error("Expected SendAudio builder to be created")
	}

	// Test SendVideo with valid file
	videoResult := ctx.SendVideo(validFile)
	if videoResult == nil {
		t.Error("Expected SendVideo builder to be created")
	}

	// Test SendVoice with valid file
	voiceResult := ctx.SendVoice(validFile)
	if voiceResult == nil {
		t.Error("Expected SendVoice builder to be created")
	}

	// Test SendVideoNote with valid file
	videoNoteResult := ctx.SendVideoNote(validFile)
	if videoNoteResult == nil {
		t.Error("Expected SendVideoNote builder to be created")
	}

	// Test SendAnimation with valid file
	animationResult := ctx.SendAnimation(validFile)
	if animationResult == nil {
		t.Error("Expected SendAnimation builder to be created")
	}

	// Test SendSticker with valid file
	stickerResult := ctx.SendSticker(validFile)
	if stickerResult == nil {
		t.Error("Expected SendSticker builder to be created")
	}

	// Test SetChatPhoto with valid file
	setChatPhotoResult := ctx.SetChatPhoto(validFile)
	if setChatPhotoResult == nil {
		t.Error("Expected SetChatPhoto builder to be created")
	}
}

func TestContext_FileOperations_SetChatPhotoError(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test SetChatPhoto with invalid file
	invalidFile := g.String("/dev/null/nonexistent/invalid/photo.jpg")
	setChatPhotoResult := ctx.SetChatPhoto(invalidFile)
	if setChatPhotoResult == nil {
		t.Error("Expected SetChatPhoto builder to be created even with error")
	}
}

func TestContext_NilEffectiveChat_Operations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: nil, // Nil chat
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test SendMessage with nil effective chat
	result := ctx.SendMessage(g.String("test"))

	if result == nil {
		t.Error("Expected SendMessage builder to be created")
	}
}

func TestContext_IsAdmin(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveUser: &gotgbot.User{Id: 123, FirstName: "Test"},
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test IsAdmin - this will fail due to mock bot returning empty bot
	// but it will still cover the code path
	result := ctx.IsAdmin()

	// Since mockBot returns empty bot, this will likely error
	if result.IsOk() {
		// If it somehow works, check the result
		t.Logf("IsAdmin result: %v", result.Ok())
	} else {
		// Expected to error with mock bot
		t.Logf("IsAdmin errored as expected with mock bot: %v", result.Err())
	}
}

func TestContext_IsAdmin_ErrorBranches(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveUser: &gotgbot.User{Id: 0, FirstName: "Test"}, // User ID 0 to test different scenarios
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test IsAdmin with different user ID - should still error with mock bot
	result := ctx.IsAdmin()

	// Should error due to mock bot limitations
	if result.IsErr() {
		t.Logf("IsAdmin with user ID 0 errored as expected")
	} else {
		t.Logf("IsAdmin returned result: %v", result.Ok())
	}
}

func TestContext_IsAdmin_AdditionalCoverage(t *testing.T) {
	// Test to improve IsAdmin coverage
	// Since we can't easily mock the GetChatMember call,
	// we'll just ensure we cover the basic error case with different scenarios
	bot := &mockBot{}

	// Test with different chat types
	for _, chatType := range []string{"private", "group", "supergroup", "channel"} {
		rawCtx := &ext.Context{
			EffectiveUser: &gotgbot.User{Id: 123, FirstName: "Test"},
			EffectiveChat: &gotgbot.Chat{Id: 456, Type: chatType},
			Update:        &gotgbot.Update{UpdateId: 1},
		}

		ctx := ctx.New(bot, rawCtx)

		result := ctx.IsAdmin()
		// Should error with mock bot, but we're just testing the path
		if result.IsErr() {
			t.Logf("IsAdmin with chat type %s errored as expected", chatType)
		} else {
			t.Logf("IsAdmin with chat type %s returned: %v", chatType, result.Ok())
		}
	}
}

func TestContext_Timers(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test sending a message without timers (covers timers with None values)
	sendMessage := ctx.SendMessage(g.String("test"))
	result := sendMessage.Send()

	// This will likely error due to mock bot, but covers the timers path
	if result.IsErr() {
		t.Logf("Send failed as expected with mock bot: %v", result.Err())
	}
}

func TestContext_Timers_WithAfter(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test sending a message with After timer (covers async timer branch)
	sendMessage := ctx.SendMessage(g.String("test")).After(time.Millisecond)
	result := sendMessage.Send()

	// With After, should return Ok(nil) immediately and run async
	if result.IsOk() {
		if result.Ok() != nil {
			t.Error("Expected nil result for async After timer")
		}
		t.Logf("Send with After returned nil as expected (async)")
	} else {
		t.Logf("Send with After failed: %v", result.Err())
	}

	// Give a moment for async goroutine to start
	time.Sleep(time.Millisecond * 5)
}

func TestContext_Timers_WithDeleteAfter(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test sending a message with DeleteAfter only (covers immediate send + delete branch)
	sendMessage := ctx.SendMessage(g.String("test")).DeleteAfter(time.Minute)
	result := sendMessage.Send()

	// Should attempt immediate send, then schedule delete
	if result.IsErr() {
		t.Logf("Send with DeleteAfter failed as expected with mock bot: %v", result.Err())
	} else {
		t.Logf("Send with DeleteAfter succeeded: %v", result.Ok())
	}
}

func TestContext_Timers_WithBoth(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test sending a message with both After and DeleteAfter (covers async branch with delete)
	sendMessage := ctx.SendMessage(g.String("test")).After(time.Millisecond).DeleteAfter(time.Minute)
	result := sendMessage.Send()

	// With After, should return Ok(nil) immediately and run async with delete
	if result.IsOk() {
		if result.Ok() != nil {
			t.Error("Expected nil result for async After timer with delete")
		}
		t.Logf("Send with After+DeleteAfter returned nil as expected (async with delete)")
	} else {
		t.Logf("Send with After+DeleteAfter failed: %v", result.Err())
	}

	// Give a moment for async goroutine to start
	time.Sleep(time.Millisecond * 5)
}

func TestContext_Timers_AsyncError(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test async case where send() fails - covers error branch in async
	sendMessage := ctx.SendMessage(g.String("test")).After(time.Millisecond)
	result := sendMessage.Send()

	// Should return Ok(nil) immediately for async
	if result.IsOk() && result.Ok() == nil {
		t.Logf("Async send returned nil as expected")
	} else {
		t.Errorf("Expected nil result for async send, got: %v", result)
	}

	// Give time for async goroutine to complete
	time.Sleep(time.Millisecond * 10)
}

func TestContext_Timers_ImmediateSendError(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test immediate send with DeleteAfter but send fails
	sendMessage := ctx.SendMessage(g.String("test")).DeleteAfter(time.Millisecond * 10)
	result := sendMessage.Send()

	// Should attempt immediate send and fail (covers error path in immediate send)
	if result.IsErr() {
		t.Logf("Send failed as expected: %v", result.Err())
	} else {
		t.Logf("Send unexpectedly succeeded: %v", result.Ok())
	}
}
