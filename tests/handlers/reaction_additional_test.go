package handlers_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/handlers"
)

func TestReactionHandlers_FromPeerWithActorChat(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for reactions from a specific chat ID (actor chat)
	handlers.NewHandlers(bot).Reaction.FromPeer(123456789, handler)

	// Create update with ActorChat (channel/group reaction)
	update := &gotgbot.Update{
		UpdateId: 1,
		MessageReaction: &gotgbot.MessageReactionUpdated{
			Chat: gotgbot.Chat{
				Id:   -1001234567890,
				Type: "supergroup",
			},
			MessageId: 100,
			Date:      1234567890,
			ActorChat: &gotgbot.Chat{
				Id:   123456789, // This should match our filter
				Type: "channel",
			},
			User: nil, // No user, only ActorChat
			NewReaction: []gotgbot.ReactionType{
				gotgbot.ReactionTypeEmoji{Emoji: "👍"},
			},
			OldReaction: []gotgbot.ReactionType{},
		},
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies by dispatcher
	}

	if !handlerExecuted {
		t.Error("Handler should have been executed for ActorChat match")
	}
}

func TestReactionHandlers_FromPeerWithUser(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for reactions from a specific user ID
	handlers.NewHandlers(bot).Reaction.FromPeer(987654321, handler)

	// Create update with User reaction
	update := &gotgbot.Update{
		UpdateId: 1,
		MessageReaction: &gotgbot.MessageReactionUpdated{
			Chat: gotgbot.Chat{
				Id:   -1001234567890,
				Type: "supergroup",
			},
			MessageId: 100,
			Date:      1234567890,
			User: &gotgbot.User{
				Id:           987654321, // This should match our filter
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
			ActorChat: nil, // No ActorChat, only User
			NewReaction: []gotgbot.ReactionType{
				gotgbot.ReactionTypeEmoji{Emoji: "❤️"},
			},
			OldReaction: []gotgbot.ReactionType{},
		},
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if !handlerExecuted {
		t.Error("Handler should have been executed for User match")
	}
}

func TestReactionHandlers_FromPeerNoMatch(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for reactions from a specific ID
	handlers.NewHandlers(bot).Reaction.FromPeer(999999999, handler)

	// Create update with different User and ActorChat IDs
	update := &gotgbot.Update{
		UpdateId: 1,
		MessageReaction: &gotgbot.MessageReactionUpdated{
			Chat: gotgbot.Chat{
				Id:   -1001234567890,
				Type: "supergroup",
			},
			MessageId: 100,
			Date:      1234567890,
			User: &gotgbot.User{
				Id:           123123123, // Different from filter
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
			ActorChat: &gotgbot.Chat{
				Id:   456456456, // Different from filter
				Type: "channel",
			},
			NewReaction: []gotgbot.ReactionType{
				gotgbot.ReactionTypeEmoji{Emoji: "👎"},
			},
			OldReaction: []gotgbot.ReactionType{},
		},
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if handlerExecuted {
		t.Error("Handler should NOT have been executed for non-matching peer")
	}
}

func TestReactionHandlers_NewReactionEmojiMatch(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for specific emoji in new reactions
	handlers.NewHandlers(bot).Reaction.NewReactionEmoji(g.String("🔥"), handler)

	// Create update with matching emoji in new reactions
	update := &gotgbot.Update{
		UpdateId: 1,
		MessageReaction: &gotgbot.MessageReactionUpdated{
			Chat: gotgbot.Chat{
				Id:   -1001234567890,
				Type: "supergroup",
			},
			MessageId: 100,
			Date:      1234567890,
			User: &gotgbot.User{
				Id:           987654321,
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
			NewReaction: []gotgbot.ReactionType{
				gotgbot.ReactionTypeEmoji{Emoji: "👍"},
				gotgbot.ReactionTypeEmoji{Emoji: "🔥"}, // This should match
				gotgbot.ReactionTypeEmoji{Emoji: "❤️"},
			},
			OldReaction: []gotgbot.ReactionType{},
		},
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if !handlerExecuted {
		t.Error("Handler should have been executed for matching new reaction emoji")
	}
}

func TestReactionHandlers_NewReactionEmojiNoMatch(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for specific emoji
	handlers.NewHandlers(bot).Reaction.NewReactionEmoji(g.String("💯"), handler)

	// Create update without matching emoji
	update := &gotgbot.Update{
		UpdateId: 1,
		MessageReaction: &gotgbot.MessageReactionUpdated{
			Chat: gotgbot.Chat{
				Id:   -1001234567890,
				Type: "supergroup",
			},
			MessageId: 100,
			Date:      1234567890,
			User: &gotgbot.User{
				Id:           987654321,
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
			NewReaction: []gotgbot.ReactionType{
				gotgbot.ReactionTypeEmoji{Emoji: "👍"},
				gotgbot.ReactionTypeEmoji{Emoji: "❤️"},
			},
			OldReaction: []gotgbot.ReactionType{},
		},
	}

	// Process the update
	dispatcher := bot.Dispatcher()
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	if handlerExecuted {
		t.Error("Handler should NOT have been executed for non-matching new reaction emoji")
	}
}

func TestReactionHandlers_OldReactionEmojiMatch(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for specific emoji in old reactions
	handlers.NewHandlers(bot).Reaction.OldReactionEmoji(g.String("😢"), handler)

	// Create update with matching emoji in old reactions
	update := &gotgbot.Update{
		UpdateId: 1,
		MessageReaction: &gotgbot.MessageReactionUpdated{
			Chat: gotgbot.Chat{
				Id:   -1001234567890,
				Type: "supergroup",
			},
			MessageId: 100,
			Date:      1234567890,
			User: &gotgbot.User{
				Id:           987654321,
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
			NewReaction: []gotgbot.ReactionType{},
			OldReaction: []gotgbot.ReactionType{
				gotgbot.ReactionTypeEmoji{Emoji: "😢"}, // This should match
				gotgbot.ReactionTypeEmoji{Emoji: "😔"},
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
		t.Error("Handler should have been executed for matching old reaction emoji")
	}
}

func TestReactionHandlers_OldReactionEmojiNoMatch(t *testing.T) {
	bot := NewMockBot()
	var handlerExecuted bool

	handler := func(c *ctx.Context) error {
		handlerExecuted = true
		return nil
	}

	// Register handler for specific emoji
	handlers.NewHandlers(bot).Reaction.OldReactionEmoji(g.String("🎉"), handler)

	// Create update without matching emoji in old reactions
	update := &gotgbot.Update{
		UpdateId: 1,
		MessageReaction: &gotgbot.MessageReactionUpdated{
			Chat: gotgbot.Chat{
				Id:   -1001234567890,
				Type: "supergroup",
			},
			MessageId: 100,
			Date:      1234567890,
			User: &gotgbot.User{
				Id:           987654321,
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
			NewReaction: []gotgbot.ReactionType{},
			OldReaction: []gotgbot.ReactionType{
				gotgbot.ReactionTypeEmoji{Emoji: "😢"},
				gotgbot.ReactionTypeEmoji{Emoji: "😔"},
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
		t.Error("Handler should NOT have been executed for non-matching old reaction emoji")
	}
}
