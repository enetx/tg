package handlers_test

import (
	"testing"

	"github.com/enetx/g"
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
