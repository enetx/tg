package handlers_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/handlers"
)

func TestReactionHandlers_Any(t *testing.T) {
	h := &handlers.ReactionHandlers{Bot: NewMockBot()}
	if r := h.Any(MockHandler); r != h {
		t.Error("Any should return same instance")
	}
}

func TestReactionHandlers_FromUserID(t *testing.T) {
	h := &handlers.ReactionHandlers{Bot: NewMockBot()}
	if r := h.FromPeer(987654321, MockHandler); r != h {
		t.Error("FromUserID should return same instance")
	}
}

func TestReactionHandlers_ChatID(t *testing.T) {
	h := &handlers.ReactionHandlers{Bot: NewMockBot()}
	if r := h.ChatID(-1001234567890, MockHandler); r != h {
		t.Error("ChatID should return same instance")
	}
}

func TestReactionHandlers_MessageID(t *testing.T) {
	h := &handlers.ReactionHandlers{Bot: NewMockBot()}
	if r := h.MessageID(42, MockHandler); r != h {
		t.Error("MessageID should return same instance")
	}
}

func TestReactionHandlers_Emoji(t *testing.T) {
	h := &handlers.ReactionHandlers{Bot: NewMockBot()}
	if r := h.NewReactionEmoji(g.String("👍"), MockHandler); r != h {
		t.Error("Emoji should return same instance")
	}
}

func TestReactionHandlers_ChainedMethods(t *testing.T) {
	h := &handlers.ReactionHandlers{Bot: NewMockBot()}
	r := h.Any(MockHandler).
		FromPeer(123456, MockHandler).
		ChatID(-1001234567890, MockHandler).
		MessageID(789, MockHandler).
		NewReactionEmoji(g.String("❤️"), MockHandler).
		OldReactionEmoji(g.String("👎"), MockHandler)
	if r != h {
		t.Error("Chained methods should return same instance")
	}
}

func TestReactionHandlers_VariousEmojis(t *testing.T) {
	emojis := []string{"👍", "👎", "❤️", "🔥", "🥰", "👏", "😁", "🤔", "🤯", "😱", "🤬", "😢", "🎉", "🤩", "🤮", "💩"}
	h := &handlers.ReactionHandlers{Bot: NewMockBot()}
	for _, e := range emojis {
		t.Run(e, func(t *testing.T) {
			if r := h.NewReactionEmoji(g.String(e), MockHandler); r != h {
				t.Errorf("Emoji %s should return same instance", e)
			}
		})
	}
}

func TestReactionHandlers_EmptyEmoji(t *testing.T) {
	h := &handlers.ReactionHandlers{Bot: NewMockBot()}
	if r := h.NewReactionEmoji(g.String(""), MockHandler); r != h {
		t.Error("Empty emoji should return same instance")
	}
}

func TestReactionHandlers_ZeroAndNegativeIDs(t *testing.T) {
	h := &handlers.ReactionHandlers{Bot: NewMockBot()}
	if r := h.FromPeer(0, MockHandler); r != h {
		t.Error("Zero user ID should return same instance")
	}
	if r := h.FromPeer(-123456789, MockHandler); r != h {
		t.Error("Negative user ID should return same instance")
	}
	if r := h.ChatID(0, MockHandler); r != h {
		t.Error("Zero chat ID should return same instance")
	}
	if r := h.ChatID(-987654321, MockHandler); r != h {
		t.Error("Negative chat ID should return same instance")
	}
	if r := h.MessageID(0, MockHandler); r != h {
		t.Error("Zero message ID should return same instance")
	}
	if r := h.MessageID(-1, MockHandler); r != h {
		t.Error("Negative message ID should return same instance")
	}
}

func TestReactionHandlers_LargeIDs(t *testing.T) {
	large := int64(9223372036854775807)
	h := &handlers.ReactionHandlers{Bot: NewMockBot()}
	if r := h.FromPeer(large, MockHandler); r != h {
		t.Error("Large user ID should return same instance")
	}
	if r := h.ChatID(large, MockHandler); r != h {
		t.Error("Large chat ID should return same instance")
	}
	if r := h.MessageID(large, MockHandler); r != h {
		t.Error("Large message ID should return same instance")
	}
}

func TestReactionHandlers_WithNilHandler(t *testing.T) {
	h := &handlers.ReactionHandlers{Bot: NewMockBot()}
	if r := h.Any(nil); r != h {
		t.Error("Nil handler should return same instance")
	}
}

func TestReactionHandlers_FromPeerEdgeCases(t *testing.T) {
	h := &handlers.ReactionHandlers{Bot: NewMockBot()}

	// Test with zero peer ID
	if r := h.FromPeer(0, MockHandler); r != h {
		t.Error("FromPeer with zero should return same instance")
	}

	// Test with negative peer ID
	if r := h.FromPeer(-123456789, MockHandler); r != h {
		t.Error("FromPeer with negative ID should return same instance")
	}

	// Test with large peer ID
	if r := h.FromPeer(9223372036854775807, MockHandler); r != h {
		t.Error("FromPeer with large ID should return same instance")
	}
}

func TestReactionHandlers_NewReactionEmojiEdgeCases(t *testing.T) {
	h := &handlers.ReactionHandlers{Bot: NewMockBot()}

	// Test with empty emoji
	if r := h.NewReactionEmoji(g.String(""), MockHandler); r != h {
		t.Error("NewReactionEmoji with empty string should return same instance")
	}

	// Test with single emoji
	if r := h.NewReactionEmoji(g.String("👍"), MockHandler); r != h {
		t.Error("NewReactionEmoji with thumbs up should return same instance")
	}

	// Test with multiple emojis
	if r := h.NewReactionEmoji(g.String("👍👎❤️"), MockHandler); r != h {
		t.Error("NewReactionEmoji with multiple emojis should return same instance")
	}

	// Test with text-based emoji
	if r := h.NewReactionEmoji(g.String(":thumbsup:"), MockHandler); r != h {
		t.Error("NewReactionEmoji with text emoji should return same instance")
	}
}

func TestReactionHandlers_OldReactionEmojiEdgeCases(t *testing.T) {
	h := &handlers.ReactionHandlers{Bot: NewMockBot()}

	// Test with empty emoji
	if r := h.OldReactionEmoji(g.String(""), MockHandler); r != h {
		t.Error("OldReactionEmoji with empty string should return same instance")
	}

	// Test with single emoji
	if r := h.OldReactionEmoji(g.String("👎"), MockHandler); r != h {
		t.Error("OldReactionEmoji with thumbs down should return same instance")
	}

	// Test with multiple emojis
	if r := h.OldReactionEmoji(g.String("😀😃😄"), MockHandler); r != h {
		t.Error("OldReactionEmoji with multiple emojis should return same instance")
	}

	// Test with special characters
	if r := h.OldReactionEmoji(g.String("🔥💯⚡"), MockHandler); r != h {
		t.Error("OldReactionEmoji with special emojis should return same instance")
	}
}

func TestReactionHandlers_ComprehensiveChaining(t *testing.T) {
	h := &handlers.ReactionHandlers{Bot: NewMockBot()}

	// Test comprehensive chaining with edge cases
	result := h.Any(MockHandler).
		FromPeer(0, MockHandler).
		ChatID(-1001234567890, MockHandler).
		MessageID(999999999, MockHandler).
		NewReactionEmoji(g.String("🚀"), MockHandler).
		OldReactionEmoji(g.String("⭐"), MockHandler)

	if result != h {
		t.Error("Comprehensive chaining should return same instance")
	}
}

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
