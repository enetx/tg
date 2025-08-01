package handlers_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/handlers"
)

type mockBot struct{}

func (m *mockBot) Raw() *gotgbot.Bot                      { return &gotgbot.Bot{} }
func (m *mockBot) Dispatcher() *ext.Dispatcher            { return &ext.Dispatcher{} }
func (m *mockBot) Updater() *ext.Updater                  { return &ext.Updater{} }
func (m *mockBot) Middlewares() g.Slice[handlers.Handler] { return g.Slice[handlers.Handler]{} }

func TestNewHandlers(t *testing.T) {
	bot := &mockBot{}
	h := handlers.NewHandlers(bot)

	if h == nil {
		t.Error("Expected NewHandlers to return non-nil Handlers")
	}

	if h.Message == nil {
		t.Error("Expected Message handlers to be initialized")
	}

	if h.Callback == nil {
		t.Error("Expected Callback handlers to be initialized")
	}

	if h.Inline == nil {
		t.Error("Expected Inline handlers to be initialized")
	}

	if h.Poll == nil {
		t.Error("Expected Poll handlers to be initialized")
	}

	if h.ChatMember == nil {
		t.Error("Expected ChatMember handlers to be initialized")
	}

	if h.PollAnswer == nil {
		t.Error("Expected PollAnswer handlers to be initialized")
	}

	if h.MyChatMember == nil {
		t.Error("Expected MyChatMember handlers to be initialized")
	}

	if h.ChatJoinRequest == nil {
		t.Error("Expected ChatJoinRequest handlers to be initialized")
	}

	if h.ChosenInlineResult == nil {
		t.Error("Expected ChosenInlineResult handlers to be initialized")
	}

	if h.Shipping == nil {
		t.Error("Expected Shipping handlers to be initialized")
	}

	if h.PreCheckout == nil {
		t.Error("Expected PreCheckout handlers to be initialized")
	}

	if h.Reaction == nil {
		t.Error("Expected Reaction handlers to be initialized")
	}

	if h.PaidMedia == nil {
		t.Error("Expected PaidMedia handlers to be initialized")
	}

	if h.BusinessConnection == nil {
		t.Error("Expected BusinessConnection handlers to be initialized")
	}

	if h.BusinessMessagesDeleted == nil {
		t.Error("Expected BusinessMessagesDeleted handlers to be initialized")
	}
}

func TestHandlersStructure(t *testing.T) {
	bot := &mockBot{}
	h := handlers.NewHandlers(bot)

	// Test that we can access all handler types without panicking
	_ = h.Message
	_ = h.Callback
	_ = h.Inline
	_ = h.Poll
	_ = h.PollAnswer
	_ = h.ChatMember
	_ = h.MyChatMember
	_ = h.ChatJoinRequest
	_ = h.ChosenInlineResult
	_ = h.Shipping
	_ = h.PreCheckout
	_ = h.Reaction
	_ = h.PaidMedia
	_ = h.BusinessConnection
	_ = h.BusinessMessagesDeleted

	// If we get here without panicking, the structure is properly initialized
	t.Log("All handler types are properly accessible")
}

func TestHandlersNilBot(t *testing.T) {
	// Test that NewHandlers can handle nil bot (if the implementation allows it)
	h := handlers.NewHandlers(nil)

	if h == nil {
		t.Error("Expected NewHandlers to return non-nil Handlers even with nil bot")
	}

	// Test that handler types are still initialized
	if h.Message == nil {
		t.Error("Expected Message handlers to be initialized even with nil bot")
	}

	if h.Callback == nil {
		t.Error("Expected Callback handlers to be initialized even with nil bot")
	}
}

func TestHandlersTypes(t *testing.T) {
	bot := &mockBot{}
	h := handlers.NewHandlers(bot)

	// Test type assertions to ensure correct types
	if _, ok := any(h.Message).(*handlers.MessageHandlers); !ok {
		t.Error("Expected Message to be of type *MessageHandlers")
	}

	if _, ok := any(h.Callback).(*handlers.CallbackHandlers); !ok {
		t.Error("Expected Callback to be of type *CallbackHandlers")
	}

	if _, ok := any(h.Inline).(*handlers.InlineQueryHandlers); !ok {
		t.Error("Expected Inline to be of type *InlineQueryHandlers")
	}

	if _, ok := any(h.Poll).(*handlers.PollHandlers); !ok {
		t.Error("Expected Poll to be of type *PollHandlers")
	}

	if _, ok := any(h.PollAnswer).(*handlers.PollAnswerHandlers); !ok {
		t.Error("Expected PollAnswer to be of type *PollAnswerHandlers")
	}

	if _, ok := any(h.ChatMember).(*handlers.ChatMemberHandlers); !ok {
		t.Error("Expected ChatMember to be of type *ChatMemberHandlers")
	}

	if _, ok := any(h.MyChatMember).(*handlers.MyChatMemberHandlers); !ok {
		t.Error("Expected MyChatMember to be of type *MyChatMemberHandlers")
	}

	if _, ok := any(h.ChatJoinRequest).(*handlers.ChatJoinRequestHandlers); !ok {
		t.Error("Expected ChatJoinRequest to be of type *ChatJoinRequestHandlers")
	}

	if _, ok := any(h.ChosenInlineResult).(*handlers.ChosenInlineResultHandlers); !ok {
		t.Error("Expected ChosenInlineResult to be of type *ChosenInlineResultHandlers")
	}

	if _, ok := any(h.Shipping).(*handlers.ShippingHandlers); !ok {
		t.Error("Expected Shipping to be of type *ShippingHandlers")
	}

	if _, ok := any(h.PreCheckout).(*handlers.PreCheckoutHandlers); !ok {
		t.Error("Expected PreCheckout to be of type *PreCheckoutHandlers")
	}

	if _, ok := any(h.Reaction).(*handlers.ReactionHandlers); !ok {
		t.Error("Expected Reaction to be of type *ReactionHandlers")
	}

	if _, ok := any(h.PaidMedia).(*handlers.PaidMediaHandlers); !ok {
		t.Error("Expected PaidMedia to be of type *PaidMediaHandlers")
	}

	if _, ok := any(h.BusinessConnection).(*handlers.BusinessConnection); !ok {
		t.Error("Expected BusinessConnection to be of type *BusinessConnection")
	}

	if _, ok := any(h.BusinessMessagesDeleted).(*handlers.BusinessMessagesDeleted); !ok {
		t.Error("Expected BusinessMessagesDeleted to be of type *BusinessMessagesDeleted")
	}
}
