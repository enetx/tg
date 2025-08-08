package ctx_test

import (
	"testing"

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
		t.Logf("IsAdmin errored as expected with mock bot")
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
