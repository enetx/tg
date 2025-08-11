package handlers_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/handlers"
)

func TestCallbackHandlers_EqualNilQuery(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for exact data match
	handlers.NewHandlers(bot).Callback.Equal(g.String("test_data"), handler)

	// Create update with nil CallbackQuery
	update := &gotgbot.Update{
		UpdateId:      1,
		CallbackQuery: nil, // This should not match
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if handlerExecuted {
		t.Error("Handler should NOT have been executed for nil CallbackQuery")
	}
}

func TestCallbackHandlers_PrefixNilQuery(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for prefix match
	handlers.NewHandlers(bot).Callback.Prefix(g.String("cmd_"), handler)

	// Create update with nil CallbackQuery
	update := &gotgbot.Update{
		UpdateId:      1,
		CallbackQuery: nil, // This should not match
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if handlerExecuted {
		t.Error("Handler should NOT have been executed for nil CallbackQuery")
	}
}

func TestCallbackHandlers_SuffixNilQuery(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for suffix match
	handlers.NewHandlers(bot).Callback.Suffix(g.String("_end"), handler)

	// Create update with nil CallbackQuery
	update := &gotgbot.Update{
		UpdateId:      1,
		CallbackQuery: nil, // This should not match
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if handlerExecuted {
		t.Error("Handler should NOT have been executed for nil CallbackQuery")
	}
}

func TestCallbackHandlers_FromUserIDNilQuery(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for specific user ID
	handlers.NewHandlers(bot).Callback.FromUserID(123456789, handler)

	// Create update with nil CallbackQuery
	update := &gotgbot.Update{
		UpdateId:      1,
		CallbackQuery: nil, // This should not match
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if handlerExecuted {
		t.Error("Handler should NOT have been executed for nil CallbackQuery")
	}
}

func TestCallbackHandlers_GameNameNilQuery(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for game name
	handlers.NewHandlers(bot).Callback.GameName(g.String("tetris"), handler)

	// Create update with nil CallbackQuery
	update := &gotgbot.Update{
		UpdateId:      1,
		CallbackQuery: nil, // This should not match
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if handlerExecuted {
		t.Error("Handler should NOT have been executed for nil CallbackQuery")
	}
}

func TestCallbackHandlers_InlineNilQuery(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for inline callback
	handlers.NewHandlers(bot).Callback.Inline(handler)

	// Create update with nil CallbackQuery
	update := &gotgbot.Update{
		UpdateId:      1,
		CallbackQuery: nil, // This should not match
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if handlerExecuted {
		t.Error("Handler should NOT have been executed for nil CallbackQuery")
	}
}

func TestCallbackHandlers_ChatInstanceNilQuery(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for specific chat instance
	handlers.NewHandlers(bot).Callback.ChatInstance(g.String("instance123"), handler)

	// Create update with nil CallbackQuery
	update := &gotgbot.Update{
		UpdateId:      1,
		CallbackQuery: nil, // This should not match
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if handlerExecuted {
		t.Error("Handler should NOT have been executed for nil CallbackQuery")
	}
}

func TestCallbackHandlers_InlineWithEmptyMessageId(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for inline callback
	handlers.NewHandlers(bot).Callback.Inline(handler)

	// Create update with CallbackQuery but empty InlineMessageId
	update := &gotgbot.Update{
		UpdateId: 1,
		CallbackQuery: &gotgbot.CallbackQuery{
			Id:   "callback123",
			Data: "test_data",
			From: gotgbot.User{
				Id:           987654321,
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
			InlineMessageId: "", // Empty inline message ID should not match
		},
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if handlerExecuted {
		t.Error("Handler should NOT have been executed for empty InlineMessageId")
	}
}

func TestCallbackHandlers_InlineWithNonEmptyMessageId(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for inline callback
	handlers.NewHandlers(bot).Callback.Inline(handler)

	// Create update with CallbackQuery and non-empty InlineMessageId
	update := &gotgbot.Update{
		UpdateId: 1,
		CallbackQuery: &gotgbot.CallbackQuery{
			Id:   "callback123",
			Data: "test_data",
			From: gotgbot.User{
				Id:           987654321,
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
			InlineMessageId: "inline123", // Non-empty inline message ID should match
		},
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if !handlerExecuted {
		t.Error("Handler should have been executed for non-empty InlineMessageId")
	}
}
