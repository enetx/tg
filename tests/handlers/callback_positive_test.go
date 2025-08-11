package handlers_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/handlers"
)

func TestCallbackHandlers_EqualPositiveMatch(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for exact data match
	handlers.NewHandlers(bot).Callback.Equal(g.String("exact_match"), handler)

	// Create update with CallbackQuery that matches exactly
	update := &gotgbot.Update{
		UpdateId: 1,
		CallbackQuery: &gotgbot.CallbackQuery{
			Id:   "callback123",
			Data: "exact_match", // This should match
			From: gotgbot.User{
				Id:           987654321,
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
		},
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if !handlerExecuted {
		t.Error("Handler should have been executed for exact match")
	}
}

func TestCallbackHandlers_EqualNoMatch(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for specific data
	handlers.NewHandlers(bot).Callback.Equal(g.String("specific_data"), handler)

	// Create update with CallbackQuery that doesn't match
	update := &gotgbot.Update{
		UpdateId: 1,
		CallbackQuery: &gotgbot.CallbackQuery{
			Id:   "callback123",
			Data: "different_data", // This should NOT match
			From: gotgbot.User{
				Id:           987654321,
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
		},
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if handlerExecuted {
		t.Error("Handler should NOT have been executed for non-matching data")
	}
}

func TestCallbackHandlers_PrefixPositiveMatch(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for prefix match
	handlers.NewHandlers(bot).Callback.Prefix(g.String("cmd_"), handler)

	// Create update with CallbackQuery that has the prefix
	update := &gotgbot.Update{
		UpdateId: 1,
		CallbackQuery: &gotgbot.CallbackQuery{
			Id:   "callback123",
			Data: "cmd_start_game", // This should match the prefix
			From: gotgbot.User{
				Id:           987654321,
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
		},
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if !handlerExecuted {
		t.Error("Handler should have been executed for prefix match")
	}
}

func TestCallbackHandlers_PrefixNoMatch(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for specific prefix
	handlers.NewHandlers(bot).Callback.Prefix(g.String("admin_"), handler)

	// Create update with CallbackQuery without the prefix
	update := &gotgbot.Update{
		UpdateId: 1,
		CallbackQuery: &gotgbot.CallbackQuery{
			Id:   "callback123",
			Data: "user_action", // This should NOT match the prefix
			From: gotgbot.User{
				Id:           987654321,
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
		},
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if handlerExecuted {
		t.Error("Handler should NOT have been executed for non-matching prefix")
	}
}

func TestCallbackHandlers_SuffixPositiveMatch(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for suffix match
	handlers.NewHandlers(bot).Callback.Suffix(g.String("_confirm"), handler)

	// Create update with CallbackQuery that has the suffix
	update := &gotgbot.Update{
		UpdateId: 1,
		CallbackQuery: &gotgbot.CallbackQuery{
			Id:   "callback123",
			Data: "delete_confirm", // This should match the suffix
			From: gotgbot.User{
				Id:           987654321,
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
		},
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if !handlerExecuted {
		t.Error("Handler should have been executed for suffix match")
	}
}

func TestCallbackHandlers_SuffixNoMatch(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for specific suffix
	handlers.NewHandlers(bot).Callback.Suffix(g.String("_cancel"), handler)

	// Create update with CallbackQuery without the suffix
	update := &gotgbot.Update{
		UpdateId: 1,
		CallbackQuery: &gotgbot.CallbackQuery{
			Id:   "callback123",
			Data: "action_confirm", // This should NOT match the suffix
			From: gotgbot.User{
				Id:           987654321,
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
		},
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if handlerExecuted {
		t.Error("Handler should NOT have been executed for non-matching suffix")
	}
}

func TestCallbackHandlers_FromUserIDPositiveMatch(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for specific user ID
	handlers.NewHandlers(bot).Callback.FromUserID(123456789, handler)

	// Create update with CallbackQuery from the matching user
	update := &gotgbot.Update{
		UpdateId: 1,
		CallbackQuery: &gotgbot.CallbackQuery{
			Id:   "callback123",
			Data: "test_data",
			From: gotgbot.User{
				Id:           123456789, // This should match
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
		},
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if !handlerExecuted {
		t.Error("Handler should have been executed for matching user ID")
	}
}

func TestCallbackHandlers_FromUserIDNoMatch(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for specific user ID
	handlers.NewHandlers(bot).Callback.FromUserID(555555555, handler)

	// Create update with CallbackQuery from different user
	update := &gotgbot.Update{
		UpdateId: 1,
		CallbackQuery: &gotgbot.CallbackQuery{
			Id:   "callback123",
			Data: "test_data",
			From: gotgbot.User{
				Id:           999999999, // This should NOT match
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
		},
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if handlerExecuted {
		t.Error("Handler should NOT have been executed for non-matching user ID")
	}
}

func TestCallbackHandlers_GameNamePositiveMatch(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for specific game name
	handlers.NewHandlers(bot).Callback.GameName(g.String("snake"), handler)

	// Create update with CallbackQuery with matching game name
	update := &gotgbot.Update{
		UpdateId: 1,
		CallbackQuery: &gotgbot.CallbackQuery{
			Id:   "callback123",
			Data: "game_action",
			From: gotgbot.User{
				Id:           987654321,
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
			GameShortName: "snake", // This should match
		},
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if !handlerExecuted {
		t.Error("Handler should have been executed for matching game name")
	}
}

func TestCallbackHandlers_GameNameNoMatch(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for specific game name
	handlers.NewHandlers(bot).Callback.GameName(g.String("tetris"), handler)

	// Create update with CallbackQuery with different game name
	update := &gotgbot.Update{
		UpdateId: 1,
		CallbackQuery: &gotgbot.CallbackQuery{
			Id:   "callback123",
			Data: "game_action",
			From: gotgbot.User{
				Id:           987654321,
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
			GameShortName: "pong", // This should NOT match
		},
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if handlerExecuted {
		t.Error("Handler should NOT have been executed for non-matching game name")
	}
}

func TestCallbackHandlers_ChatInstancePositiveMatch(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for specific chat instance
	handlers.NewHandlers(bot).Callback.ChatInstance(g.String("chat_abc123"), handler)

	// Create update with CallbackQuery with matching chat instance
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
			ChatInstance: "chat_abc123", // This should match
		},
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if !handlerExecuted {
		t.Error("Handler should have been executed for matching chat instance")
	}
}

func TestCallbackHandlers_ChatInstanceNoMatch(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for specific chat instance
	handlers.NewHandlers(bot).Callback.ChatInstance(g.String("chat_xyz789"), handler)

	// Create update with CallbackQuery with different chat instance
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
			ChatInstance: "chat_different", // This should NOT match
		},
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if handlerExecuted {
		t.Error("Handler should NOT have been executed for non-matching chat instance")
	}
}
