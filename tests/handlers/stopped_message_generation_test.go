package handlers_test

import (
	"encoding/json"
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/handlers"
)

func TestStoppedMessageGenerationDispatch(t *testing.T) {
	tests := []struct {
		name     string
		register func(*handlers.Handlers, handlers.Handler)
		matches  bool
	}{
		{"Any", func(h *handlers.Handlers, fn handlers.Handler) { h.StoppedMessageGeneration.Any(fn) }, true},
		{"AllEvents", func(h *handlers.Handlers, fn handlers.Handler) { h.Any(fn) }, true},
		{"ChatID", func(h *handlers.Handlers, fn handlers.Handler) { h.StoppedMessageGeneration.ChatID(456, fn) }, true},
		{"OtherChat", func(h *handlers.Handlers, fn handlers.Handler) { h.StoppedMessageGeneration.ChatID(999, fn) }, false},
		{"Thread", func(h *handlers.Handlers, fn handlers.Handler) { h.StoppedMessageGeneration.Thread(7, fn) }, true},
		{"OtherThread", func(h *handlers.Handlers, fn handlers.Handler) { h.StoppedMessageGeneration.Thread(99, fn) }, false},
		{"DraftID", func(h *handlers.Handlers, fn handlers.Handler) { h.StoppedMessageGeneration.DraftID(42, fn) }, true},
		{"OtherDraft", func(h *handlers.Handlers, fn handlers.Handler) { h.StoppedMessageGeneration.DraftID(99, fn) }, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bot := NewMockBot()
			called, middlewareCalled := false, false
			bot.SetMiddlewares([]handlers.Handler{func(c *ctx.Context) error {
				middlewareCalled = true
				return nil
			}})
			tt.register(handlers.NewHandlers(bot), func(c *ctx.Context) error {
				called = true
				if c.EffectiveChat == nil || c.EffectiveChat.Id != 456 || c.Raw.EffectiveChat != c.EffectiveChat {
					t.Error("Expected the stopped draft chat in both contexts")
				}
				if c.Update.StoppedMessageGeneration.DraftId != 42 {
					t.Error("Expected stopped draft identifier")
				}
				return nil
			})
			var update gotgbot.Update
			if err := json.Unmarshal([]byte(`{"update_id":1,"stopped_message_generation":{"chat":{"id":456,"type":"private"},"message_thread_id":7,"draft_id":42}}`), &update); err != nil {
				t.Fatal(err)
			}
			if err := bot.Dispatcher().ProcessUpdate(bot.Raw(), &update, nil); err != nil {
				t.Fatal(err)
			}
			if called != tt.matches || middlewareCalled != tt.matches {
				t.Errorf("Handler called=%v, middleware called=%v; expected %v", called, middlewareCalled, tt.matches)
			}
			called = false
			if err := bot.Dispatcher().ProcessUpdate(bot.Raw(), &gotgbot.Update{}, nil); err != nil {
				t.Fatal(err)
			}
			if called {
				t.Error("Stopped message generation handler must ignore unrelated updates")
			}
		})
	}
}

func TestCommunityChatJoinedDispatch(t *testing.T) {
	bot := NewMockBot()
	called := false
	handlers.NewHandlers(bot).Message.CommunityChatJoined(func(c *ctx.Context) error {
		called = true
		if c.EffectiveMessage.CommunityChatJoined == nil {
			t.Error("Expected community chat joined information")
		}
		return nil
	})
	var update gotgbot.Update
	if err := json.Unmarshal([]byte(`{"update_id":1,"message":{"message_id":1,"date":1,"chat":{"id":456,"type":"supergroup"},"community_chat_joined":{"community":{"id":123,"name":"Community"}}}}`), &update); err != nil {
		t.Fatal(err)
	}
	if err := bot.Dispatcher().ProcessUpdate(bot.Raw(), &update, nil); err != nil {
		t.Fatal(err)
	}
	if !called {
		t.Error("Expected the community chat joined handler to run")
	}
}
