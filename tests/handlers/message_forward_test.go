package handlers_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/handlers"
)

func TestMessageHandlers_ForwardFromUserNilOrigin(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for messages forwarded from specific user
	handlers.NewHandlers(bot).Message.ForwardFromUser(123456789, handler)

	// Create update with message that has no ForwardOrigin
	update := &gotgbot.Update{
		UpdateId: 1,
		Message: &gotgbot.Message{
			MessageId: 123,
			Date:      1234567890,
			Text:      "Hello world",
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
			ForwardOrigin: nil, // No forward origin - should not match
		},
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if handlerExecuted {
		t.Error("Handler should NOT have been executed for message without ForwardOrigin")
	}
}

func TestMessageHandlers_ForwardFromUserWithHiddenUser(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for messages forwarded from specific user
	handlers.NewHandlers(bot).Message.ForwardFromUser(123456789, handler)

	// Create update with message forwarded from hidden user (no SenderUser)
	update := &gotgbot.Update{
		UpdateId: 1,
		Message: &gotgbot.Message{
			MessageId: 123,
			Date:      1234567890,
			Text:      "Forwarded message",
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
			ForwardOrigin: &gotgbot.MessageOriginHiddenUser{
				Date:           1234567890,
				SenderUserName: "HiddenUser",
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
		t.Error("Handler should NOT have been executed for message from hidden user (no SenderUser)")
	}
}

func TestMessageHandlers_ForwardFromUserPositiveMatch(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for messages forwarded from specific user
	handlers.NewHandlers(bot).Message.ForwardFromUser(555555555, handler)

	// Create update with message forwarded from matching user
	update := &gotgbot.Update{
		UpdateId: 1,
		Message: &gotgbot.Message{
			MessageId: 123,
			Date:      1234567890,
			Text:      "Forwarded message",
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
			ForwardOrigin: &gotgbot.MessageOriginUser{
				Date: 1234567890,
				SenderUser: gotgbot.User{
					Id:        555555555, // This should match
					IsBot:     false,
					FirstName: "Forwarded",
					Username:  "forwardeduser",
				},
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
		t.Error("Handler should have been executed for message forwarded from matching user")
	}
}

func TestMessageHandlers_ForwardFromUserNoMatch(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for messages forwarded from specific user
	handlers.NewHandlers(bot).Message.ForwardFromUser(999999999, handler)

	// Create update with message forwarded from different user
	update := &gotgbot.Update{
		UpdateId: 1,
		Message: &gotgbot.Message{
			MessageId: 123,
			Date:      1234567890,
			Text:      "Forwarded message",
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
			ForwardOrigin: &gotgbot.MessageOriginUser{
				Date: 1234567890,
				SenderUser: gotgbot.User{
					Id:        111111111, // This should NOT match
					IsBot:     false,
					FirstName: "Forwarded",
					Username:  "forwardeduser",
				},
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
		t.Error("Handler should NOT have been executed for message forwarded from different user")
	}
}

func TestMessageHandlers_ForwardFromChatNilOrigin(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for messages forwarded from specific chat
	handlers.NewHandlers(bot).Message.ForwardFromChat(-1001111111111, handler)

	// Create update with message that has no ForwardOrigin
	update := &gotgbot.Update{
		UpdateId: 1,
		Message: &gotgbot.Message{
			MessageId: 123,
			Date:      1234567890,
			Text:      "Hello world",
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
			ForwardOrigin: nil, // No forward origin - should not match
		},
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if handlerExecuted {
		t.Error("Handler should NOT have been executed for message without ForwardOrigin")
	}
}

func TestMessageHandlers_ForwardFromChatWithUserOrigin(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for messages forwarded from specific chat
	handlers.NewHandlers(bot).Message.ForwardFromChat(-1001111111111, handler)

	// Create update with message forwarded from user (no Chat)
	update := &gotgbot.Update{
		UpdateId: 1,
		Message: &gotgbot.Message{
			MessageId: 123,
			Date:      1234567890,
			Text:      "Forwarded message",
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
			ForwardOrigin: &gotgbot.MessageOriginUser{
				Date: 1234567890,
				SenderUser: gotgbot.User{
					Id:        555555555,
					IsBot:     false,
					FirstName: "Forwarded",
					Username:  "forwardeduser",
				},
			}, // No Chat in origin - should not match
		},
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if handlerExecuted {
		t.Error("Handler should NOT have been executed for message forwarded from user (no Chat)")
	}
}

func TestMessageHandlers_ForwardFromChatPositiveMatch(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for messages forwarded from specific chat
	handlers.NewHandlers(bot).Message.ForwardFromChat(-1002222222222, handler)

	// Create update with message forwarded from matching chat
	update := &gotgbot.Update{
		UpdateId: 1,
		Message: &gotgbot.Message{
			MessageId: 123,
			Date:      1234567890,
			Text:      "Forwarded message",
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
			ForwardOrigin: &gotgbot.MessageOriginChannel{
				Date: 1234567890,
				Chat: gotgbot.Chat{
					Id:    -1002222222222, // This should match
					Type:  "channel",
					Title: "Source Channel",
				},
				MessageId: 456,
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
		t.Error("Handler should have been executed for message forwarded from matching chat")
	}
}

func TestMessageHandlers_ForwardFromChatNoMatch(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for messages forwarded from specific chat
	handlers.NewHandlers(bot).Message.ForwardFromChat(-1009999999999, handler)

	// Create update with message forwarded from different chat
	update := &gotgbot.Update{
		UpdateId: 1,
		Message: &gotgbot.Message{
			MessageId: 123,
			Date:      1234567890,
			Text:      "Forwarded message",
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
			ForwardOrigin: &gotgbot.MessageOriginChannel{
				Date: 1234567890,
				Chat: gotgbot.Chat{
					Id:    -1003333333333, // This should NOT match
					Type:  "channel",
					Title: "Different Channel",
				},
				MessageId: 789,
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
		t.Error("Handler should NOT have been executed for message forwarded from different chat")
	}
}
