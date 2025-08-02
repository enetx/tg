package handlers_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/core"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/handlers"
)

// MockBot implements the core.BotAPI interface for testing
type MockBot struct {
	dispatcher   *ext.Dispatcher
	updater      *ext.Updater
	rawBot       *gotgbot.Bot
	middlewares  g.Slice[handlers.Handler]
	addedHandler any
	handlerGroup int
	removedName  string
}

// NewMockBot creates a new mock bot instance
func NewMockBot() *MockBot {
	return &MockBot{
		dispatcher:  &ext.Dispatcher{},
		updater:     &ext.Updater{},
		rawBot:      &gotgbot.Bot{},
		middlewares: g.NewSlice[handlers.Handler](),
	}
}

// Raw returns the raw gotgbot.Bot instance
func (m *MockBot) Raw() *gotgbot.Bot {
	return m.rawBot
}

// Dispatcher returns the ext.Dispatcher
func (m *MockBot) Dispatcher() *ext.Dispatcher {
	// Mock dispatcher that tracks handler operations
	return &ext.Dispatcher{}
}

// Updater returns the ext.Updater
func (m *MockBot) Updater() *ext.Updater {
	return m.updater
}

// Middlewares returns the middleware handlers (implements the middleware interface)
func (m *MockBot) Middlewares() g.Slice[handlers.Handler] {
	return m.middlewares
}

// SetMiddlewares sets the middleware handlers for testing
func (m *MockBot) SetMiddlewares(mw g.Slice[handlers.Handler]) {
	m.middlewares = mw
}

// MockContext creates a mock context for testing
func NewMockContext() *ctx.Context {
	bot := NewMockBot()

	update := &gotgbot.Update{
		UpdateId: 1,
		Message: &gotgbot.Message{
			MessageId: 123,
			Date:      1234567890,
			Chat: gotgbot.Chat{
				Id:   -1001234567890,
				Type: "supergroup",
			},
			From: &gotgbot.User{
				Id:           987654321,
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
			Text: "/test hello world",
		},
	}

	extCtx := &ext.Context{
		Update:           update,
		EffectiveMessage: update.Message,
		EffectiveUser:    update.Message.From,
		EffectiveChat:    &update.Message.Chat,
		EffectiveSender:  &gotgbot.Sender{User: update.Message.From},
	}

	return ctx.New(bot, extCtx)
}

// MockContextWithCallback creates a mock context with callback query
func NewMockContextWithCallback() *ctx.Context {
	bot := NewMockBot()

	update := &gotgbot.Update{
		UpdateId: 1,
		CallbackQuery: &gotgbot.CallbackQuery{
			Id: "callback123",
			From: gotgbot.User{
				Id:           987654321,
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
			Data: "test_callback_data",
		},
	}

	extCtx := &ext.Context{
		Update:        update,
		EffectiveUser: &update.CallbackQuery.From,
	}

	return ctx.New(bot, extCtx)
}

// MockContextWithInlineQuery creates a mock context with inline query
func NewMockContextWithInlineQuery() *ctx.Context {
	bot := NewMockBot()

	update := &gotgbot.Update{
		UpdateId: 1,
		InlineQuery: &gotgbot.InlineQuery{
			Id: "inline123",
			From: gotgbot.User{
				Id:           987654321,
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
			Query:    "test query",
			Offset:   "",
			ChatType: "private",
		},
	}

	extCtx := &ext.Context{
		Update:        update,
		EffectiveUser: &update.InlineQuery.From,
	}

	return ctx.New(bot, extCtx)
}

// MockContextWithPoll creates a mock context with poll
func NewMockContextWithPoll() *ctx.Context {
	bot := NewMockBot()

	update := &gotgbot.Update{
		UpdateId: 1,
		Poll: &gotgbot.Poll{
			Id:       "poll123",
			Question: "Test question?",
			Options: []gotgbot.PollOption{
				{Text: "Option 1", VoterCount: 5},
				{Text: "Option 2", VoterCount: 3},
			},
			TotalVoterCount:       8,
			IsClosed:              false,
			IsAnonymous:           true,
			Type:                  "regular",
			AllowsMultipleAnswers: false,
		},
	}

	extCtx := &ext.Context{
		Update: update,
	}

	return ctx.New(bot, extCtx)
}

// MockContextWithPollAnswer creates a mock context with poll answer
func NewMockContextWithPollAnswer() *ctx.Context {
	bot := NewMockBot()

	update := &gotgbot.Update{
		UpdateId: 1,
		PollAnswer: &gotgbot.PollAnswer{
			PollId: "poll123",
			VoterChat: &gotgbot.Chat{
				Id:   -1001234567890,
				Type: "supergroup",
			},
			User: &gotgbot.User{
				Id:           987654321,
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
			OptionIds: []int64{0, 1},
		},
	}

	extCtx := &ext.Context{
		Update: update,
	}

	return ctx.New(bot, extCtx)
}

// MockContextWithChatMemberUpdate creates a mock context with chat member update
func NewMockContextWithChatMemberUpdate() *ctx.Context {
	bot := NewMockBot()

	update := &gotgbot.Update{
		UpdateId: 1,
		ChatMember: &gotgbot.ChatMemberUpdated{
			Chat: gotgbot.Chat{
				Id:   -1001234567890,
				Type: "supergroup",
			},
			From: gotgbot.User{
				Id:           987654321,
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
			Date: 1234567890,
		},
	}

	extCtx := &ext.Context{
		Update: update,
	}

	return ctx.New(bot, extCtx)
}

// MockHandler is a simple mock handler function
func MockHandler(c *ctx.Context) error {
	return nil
}

// MockMiddleware is a simple mock middleware function
func MockMiddleware(c *ctx.Context) error {
	return nil
}

// AssertHandlerInterface checks if a handler implements the required interfaces
func AssertHandlerInterface(t *testing.T, handler any) {
	t.Helper()

	// Check if handler implements ext.Handler interface
	if _, ok := handler.(ext.Handler); !ok {
		t.Errorf("Handler %T does not implement ext.Handler interface", handler)
	}
}

// AssertHandlerFunctionality tests basic handler functionality
func AssertHandlerFunctionality(t *testing.T, handler ext.Handler, update *gotgbot.Update) {
	t.Helper()

	// Test CheckUpdate method
	bot := &gotgbot.Bot{}
	extCtx := &ext.Context{Update: update}

	if !handler.CheckUpdate(bot, extCtx) {
		t.Error("Handler.CheckUpdate should return true for valid update")
	}

	// Test HandleUpdate method

	err := handler.HandleUpdate(bot, extCtx)
	if err != nil {
		t.Errorf("Handler.HandleUpdate returned error: %v", err)
	}
}

// CreateTestUpdate creates a basic test update
func CreateTestUpdate() *gotgbot.Update {
	return &gotgbot.Update{
		UpdateId: 1,
		Message: &gotgbot.Message{
			MessageId: 123,
			Date:      1234567890,
			Chat: gotgbot.Chat{
				Id:   -1001234567890,
				Type: "supergroup",
			},
			From: &gotgbot.User{
				Id:           987654321,
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
			Text: "/test hello world",
		},
	}
}

// CreateCallbackUpdate creates a callback query update
func CreateCallbackUpdate() *gotgbot.Update {
	return &gotgbot.Update{
		UpdateId: 1,
		CallbackQuery: &gotgbot.CallbackQuery{
			Id: "callback123",
			From: gotgbot.User{
				Id:           987654321,
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
			Data: "test_data",
		},
	}
}

// CreateInlineQueryUpdate creates an inline query update
func CreateInlineQueryUpdate() *gotgbot.Update {
	return &gotgbot.Update{
		UpdateId: 1,
		InlineQuery: &gotgbot.InlineQuery{
			Id: "inline123",
			From: gotgbot.User{
				Id:           987654321,
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
			Query:  "test query",
			Offset: "",
		},
	}
}

// TestValidBot ensures the mock bot implements the required interface
func TestValidBot(t *testing.T) {
	bot := NewMockBot()

	// Ensure bot implements core.BotAPI
	var _ core.BotAPI = bot

	if bot.Raw() == nil {
		t.Error("Mock bot Raw() should not return nil")
	}

	if bot.Dispatcher() == nil {
		t.Error("Mock bot Dispatcher() should not return nil")
	}

	if bot.Updater() == nil {
		t.Error("Mock bot Updater() should not return nil")
	}
}

// TestValidContext ensures the mock context is properly constructed
func TestValidContext(t *testing.T) {
	c := NewMockContext()

	if c == nil {
		t.Error("Mock context should not be nil")
	}

	if c.Bot == nil {
		t.Error("Context Bot should not be nil")
	}

	if c.EffectiveMessage == nil {
		t.Error("Context EffectiveMessage should not be nil")
	}

	if c.EffectiveUser == nil {
		t.Error("Context EffectiveUser should not be nil")
	}

	if c.EffectiveChat == nil {
		t.Error("Context EffectiveChat should not be nil")
	}

	if c.Update == nil {
		t.Error("Context Update should not be nil")
	}

	if c.Raw == nil {
		t.Error("Context Raw should not be nil")
	}
}
