package handlers_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/handlers"
	"github.com/enetx/tg/types/entity"
)

func TestMessageHandlers_EntityWithMatch(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for messages with hashtag entities
	handlers.NewHandlers(bot).Message.Entity(entity.Hashtag, handler)

	// Create update with message containing hashtag entity
	update := &gotgbot.Update{
		UpdateId: 1,
		Message: &gotgbot.Message{
			MessageId: 123,
			Date:      1234567890,
			Text:      "Check out this #awesome hashtag!",
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
			Entities: []gotgbot.MessageEntity{
				{
					Type:   "hashtag", // This should match
					Offset: 15,
					Length: 8,
				},
				{
					Type:   "mention",
					Offset: 0,
					Length: 5,
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
		t.Error("Handler should have been executed for message with matching entity")
	}
}

func TestMessageHandlers_EntityNoMatch(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for messages with URL entities
	handlers.NewHandlers(bot).Message.Entity(entity.URL, handler)

	// Create update with message containing other entities but not URL
	update := &gotgbot.Update{
		UpdateId: 1,
		Message: &gotgbot.Message{
			MessageId: 123,
			Date:      1234567890,
			Text:      "Check out this #awesome hashtag @user!",
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
			Entities: []gotgbot.MessageEntity{
				{
					Type:   "hashtag", // This should NOT match (looking for URL)
					Offset: 15,
					Length: 8,
				},
				{
					Type:   "mention", // This should NOT match (looking for URL)
					Offset: 24,
					Length: 5,
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
		t.Error("Handler should NOT have been executed for message without matching entity")
	}
}

func TestMessageHandlers_EntityNoEntities(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for messages with mention entities
	handlers.NewHandlers(bot).Message.Entity(entity.Mention, handler)

	// Create update with message without any entities
	update := &gotgbot.Update{
		UpdateId: 1,
		Message: &gotgbot.Message{
			MessageId: 123,
			Date:      1234567890,
			Text:      "Just plain text without any entities",
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
			Entities: nil, // No entities
		},
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if handlerExecuted {
		t.Error("Handler should NOT have been executed for message without any entities")
	}
}

func TestMessageHandlers_CaptionEntityWithMatch(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for messages with bold entities in caption
	handlers.NewHandlers(bot).Message.CaptionEntity(entity.Bold, handler)

	// Create update with photo message containing bold entity in caption
	update := &gotgbot.Update{
		UpdateId: 1,
		Message: &gotgbot.Message{
			MessageId: 123,
			Date:      1234567890,
			Caption:   "This is *bold* text in caption",
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
			Photo: []gotgbot.PhotoSize{
				{
					FileId:       "photo123",
					FileUniqueId: "unique123",
					Width:        800,
					Height:       600,
					FileSize:     50000,
				},
			},
			CaptionEntities: []gotgbot.MessageEntity{
				{
					Type:   "bold", // This should match
					Offset: 8,
					Length: 4,
				},
				{
					Type:   "italic",
					Offset: 20,
					Length: 4,
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
		t.Error("Handler should have been executed for message with matching caption entity")
	}
}

func TestMessageHandlers_CaptionEntityNoMatch(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for messages with code entities in caption
	handlers.NewHandlers(bot).Message.CaptionEntity(entity.Code, handler)

	// Create update with photo message containing other entities but not code
	update := &gotgbot.Update{
		UpdateId: 1,
		Message: &gotgbot.Message{
			MessageId: 123,
			Date:      1234567890,
			Caption:   "This is *bold* and _italic_ text",
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
			Photo: []gotgbot.PhotoSize{
				{
					FileId:       "photo123",
					FileUniqueId: "unique123",
					Width:        800,
					Height:       600,
					FileSize:     50000,
				},
			},
			CaptionEntities: []gotgbot.MessageEntity{
				{
					Type:   "bold", // This should NOT match (looking for code)
					Offset: 8,
					Length: 4,
				},
				{
					Type:   "italic", // This should NOT match (looking for code)
					Offset: 17,
					Length: 6,
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
		t.Error("Handler should NOT have been executed for message without matching caption entity")
	}
}

func TestMessageHandlers_CaptionEntityNoCaptionEntities(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for messages with underline entities in caption
	handlers.NewHandlers(bot).Message.CaptionEntity(entity.Underline, handler)

	// Create update with photo message without caption entities
	update := &gotgbot.Update{
		UpdateId: 1,
		Message: &gotgbot.Message{
			MessageId: 123,
			Date:      1234567890,
			Caption:   "Plain text caption without formatting",
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
			Photo: []gotgbot.PhotoSize{
				{
					FileId:       "photo123",
					FileUniqueId: "unique123",
					Width:        800,
					Height:       600,
					FileSize:     50000,
				},
			},
			CaptionEntities: nil, // No caption entities
		},
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if handlerExecuted {
		t.Error("Handler should NOT have been executed for message without any caption entities")
	}
}

func TestMessageHandlers_EntityMultipleEntitiesWithMatch(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for messages with bot command entities
	handlers.NewHandlers(bot).Message.Entity(entity.BotCommand, handler)

	// Create update with message containing bot command among other entities
	update := &gotgbot.Update{
		UpdateId: 1,
		Message: &gotgbot.Message{
			MessageId: 123,
			Date:      1234567890,
			Text:      "/start @username #hashtag https://example.com",
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
			Entities: []gotgbot.MessageEntity{
				{
					Type:   "bot_command", // This should match
					Offset: 0,
					Length: 6,
				},
				{
					Type:   "mention",
					Offset: 7,
					Length: 9,
				},
				{
					Type:   "hashtag",
					Offset: 17,
					Length: 8,
				},
				{
					Type:   "url",
					Offset: 26,
					Length: 19,
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
		t.Error("Handler should have been executed for message with bot command entity among others")
	}
}
