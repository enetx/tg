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
		t.Error("Expected DeleteMessage builder to be created")
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
