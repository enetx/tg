package handlers_test

import (
	"regexp"
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/handlers"
	"github.com/enetx/tg/types/chat"
	"github.com/enetx/tg/types/entity"
)

func TestMessageHandlers_Any(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Any(MockHandler)

	if handler == nil {
		t.Error("Any should return a MessageHandler")
	}

	// Test that handler is properly configured
	if handler == nil {
		t.Error("Handler should not be nil")
	}
}

func TestMessageHandlers_Text(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Text(MockHandler)

	if handler == nil {
		t.Error("Text should return a MessageHandler")
	}
}

func TestMessageHandlers_Photo(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Photo(MockHandler)

	if handler == nil {
		t.Error("Photo should return a MessageHandler")
	}
}

func TestMessageHandlers_Voice(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Voice(MockHandler)

	if handler == nil {
		t.Error("Voice should return a MessageHandler")
	}
}

func TestMessageHandlers_Video(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Video(MockHandler)

	if handler == nil {
		t.Error("Video should return a MessageHandler")
	}
}

func TestMessageHandlers_Audio(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Audio(MockHandler)

	if handler == nil {
		t.Error("Audio should return a MessageHandler")
	}
}

func TestMessageHandlers_Sticker(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Sticker(MockHandler)

	if handler == nil {
		t.Error("Sticker should return a MessageHandler")
	}
}

func TestMessageHandlers_Document(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Document(MockHandler)

	if handler == nil {
		t.Error("Document should return a MessageHandler")
	}
}

func TestMessageHandlers_Location(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Location(MockHandler)

	if handler == nil {
		t.Error("Location should return a MessageHandler")
	}
}

func TestMessageHandlers_Contact(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Contact(MockHandler)

	if handler == nil {
		t.Error("Contact should return a MessageHandler")
	}
}

func TestMessageHandlers_Poll(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Poll(MockHandler)

	if handler == nil {
		t.Error("Poll should return a MessageHandler")
	}
}

func TestMessageHandlers_Reply(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Reply(MockHandler)

	if handler == nil {
		t.Error("Reply should return a MessageHandler")
	}
}

func TestMessageHandlers_Animation(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Animation(MockHandler)

	if handler == nil {
		t.Error("Animation should return a MessageHandler")
	}
}

func TestMessageHandlers_VideoNote(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.VideoNote(MockHandler)

	if handler == nil {
		t.Error("VideoNote should return a MessageHandler")
	}
}

func TestMessageHandlers_Dice(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Dice(MockHandler)

	if handler == nil {
		t.Error("Dice should return a MessageHandler")
	}
}

func TestMessageHandlers_Game(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Game(MockHandler)

	if handler == nil {
		t.Error("Game should return a MessageHandler")
	}
}

func TestMessageHandlers_Venue(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Venue(MockHandler)

	if handler == nil {
		t.Error("Venue should return a MessageHandler")
	}
}

func TestMessageHandlers_FromUser(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.FromUser(987654321, MockHandler)

	if handler == nil {
		t.Error("FromUser should return a MessageHandler")
	}
}

func TestMessageHandlers_FromUsername(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.FromUsername(g.String("testuser"), MockHandler)

	if handler == nil {
		t.Error("FromUsername should return a MessageHandler")
	}
}

func TestMessageHandlers_Forwarded(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Forwarded(MockHandler)

	if handler == nil {
		t.Error("Forwarded should return a MessageHandler")
	}
}

func TestMessageHandlers_ForwardFromUser(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.ForwardFromUser(123456789, MockHandler)

	if handler == nil {
		t.Error("ForwardFromUser should return a MessageHandler")
	}
}

func TestMessageHandlers_ForwardFromChat(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.ForwardFromChat(-1001234567890, MockHandler)

	if handler == nil {
		t.Error("ForwardFromChat should return a MessageHandler")
	}
}

func TestMessageHandlers_Prefix(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Prefix(g.String("/"), MockHandler)

	if handler == nil {
		t.Error("Prefix should return a MessageHandler")
	}
}

func TestMessageHandlers_Suffix(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Suffix(g.String("!"), MockHandler)

	if handler == nil {
		t.Error("Suffix should return a MessageHandler")
	}
}

func TestMessageHandlers_Contains(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Contains(g.String("hello"), MockHandler)

	if handler == nil {
		t.Error("Contains should return a MessageHandler")
	}
}

func TestMessageHandlers_Equal(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Equal(g.String("exact match"), MockHandler)

	if handler == nil {
		t.Error("Equal should return a MessageHandler")
	}
}

func TestMessageHandlers_MatchRegex(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	pattern := regexp.MustCompile(`^\d+$`)
	handler := msgHandlers.MatchRegex(pattern, MockHandler)

	if handler == nil {
		t.Error("MatchRegex should return a MessageHandler")
	}
}

func TestMessageHandlers_ChatUsername(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.ChatUsername(g.String("testchat"), MockHandler)

	if handler == nil {
		t.Error("ChatUsername should return a MessageHandler")
	}
}

func TestMessageHandlers_ChatID(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.ChatID(-1001234567890, MockHandler)

	if handler == nil {
		t.Error("ChatID should return a MessageHandler")
	}
}

func TestMessageHandlers_ChatType(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.ChatType(chat.Private, MockHandler)

	if handler == nil {
		t.Error("ChatType should return a MessageHandler")
	}
}

func TestMessageHandlers_Private(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Private(MockHandler)

	if handler == nil {
		t.Error("Private should return a MessageHandler")
	}
}

func TestMessageHandlers_Group(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Group(MockHandler)

	if handler == nil {
		t.Error("Group should return a MessageHandler")
	}
}

func TestMessageHandlers_Supergroup(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Supergroup(MockHandler)

	if handler == nil {
		t.Error("Supergroup should return a MessageHandler")
	}
}

func TestMessageHandlers_Channel(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Channel(MockHandler)

	if handler == nil {
		t.Error("Channel should return a MessageHandler")
	}
}

func TestMessageHandlers_Business(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Business(MockHandler)

	if handler == nil {
		t.Error("Business should return a MessageHandler")
	}
}

func TestMessageHandlers_Entity(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Entity(entity.Bold, MockHandler)

	if handler == nil {
		t.Error("Entity should return a MessageHandler")
	}
}

func TestMessageHandlers_CaptionEntity(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.CaptionEntity(entity.Italic, MockHandler)

	if handler == nil {
		t.Error("CaptionEntity should return a MessageHandler")
	}
}

func TestMessageHandlers_DiceValue(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.DiceValue(6, MockHandler)

	if handler == nil {
		t.Error("DiceValue should return a MessageHandler")
	}
}

func TestMessageHandlers_Checklist(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Checklist(MockHandler)

	if handler == nil {
		t.Error("Checklist should return a MessageHandler")
	}
}

func TestMessageHandlers_ChecklistTitleContains(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.ChecklistTitleContains(g.String("task"), MockHandler)

	if handler == nil {
		t.Error("ChecklistTitleContains should return a MessageHandler")
	}
}

func TestMessageHandlers_ChecklistTitleEquals(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.ChecklistTitleEquals(g.String("Shopping List"), MockHandler)

	if handler == nil {
		t.Error("ChecklistTitleEquals should return a MessageHandler")
	}
}

func TestMessageHandler_AllowEdited(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Text(MockHandler).AllowEdited()

	if handler == nil {
		t.Error("AllowEdited should return the same MessageHandler")
	}
}

func TestMessageHandler_AllowChannel(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Text(MockHandler).AllowChannel()

	if handler == nil {
		t.Error("AllowChannel should return the same MessageHandler")
	}
}

func TestMessageHandler_AllowBusiness(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Text(MockHandler).AllowBusiness()

	if handler == nil {
		t.Error("AllowBusiness should return the same MessageHandler")
	}
}

func TestMessageHandler_ChainedMethods(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Text(MockHandler).
		AllowEdited().
		AllowChannel().
		AllowBusiness()

	if handler == nil {
		t.Error("Chained methods should return the same MessageHandler")
	}
}

func TestMessageHandler_Register(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.Text(MockHandler)

	// Register should return the same handler
	registered := handler.Register()
	if registered != handler {
		t.Error("Register should return the same handler instance")
	}
}

func TestMessageHandlers_SpecialMessageTypes(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	// Test all special message type handlers
	tests := []struct {
		name    string
		handler func() *handlers.MessageHandler
	}{
		{"NewChatMembers", func() *handlers.MessageHandler { return msgHandlers.NewChatMembers(MockHandler) }},
		{"LeftChatMember", func() *handlers.MessageHandler { return msgHandlers.LeftChatMember(MockHandler) }},
		{"PinnedMessage", func() *handlers.MessageHandler { return msgHandlers.PinnedMessage(MockHandler) }},
		{"ViaBot", func() *handlers.MessageHandler { return msgHandlers.ViaBot(MockHandler) }},
		{"Entities", func() *handlers.MessageHandler { return msgHandlers.Entities(MockHandler) }},
		{"Caption", func() *handlers.MessageHandler { return msgHandlers.Caption(MockHandler) }},
		{"CaptionEntities", func() *handlers.MessageHandler { return msgHandlers.CaptionEntities(MockHandler) }},
		{"Migrate", func() *handlers.MessageHandler { return msgHandlers.Migrate(MockHandler) }},
		{"MediaGroup", func() *handlers.MessageHandler { return msgHandlers.MediaGroup(MockHandler) }},
		{"IsAutomaticForward", func() *handlers.MessageHandler { return msgHandlers.IsAutomaticForward(MockHandler) }},
		{"UsersShared", func() *handlers.MessageHandler { return msgHandlers.UsersShared(MockHandler) }},
		{"ChatShared", func() *handlers.MessageHandler { return msgHandlers.ChatShared(MockHandler) }},
		{"Story", func() *handlers.MessageHandler { return msgHandlers.Story(MockHandler) }},
		{"WebAppData", func() *handlers.MessageHandler { return msgHandlers.WebAppData(MockHandler) }},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := test.handler()
			if handler == nil {
				t.Errorf("%s should return a MessageHandler", test.name)
			}
		})
	}
}

func TestMessageHandlers_ForumSpecificHandlers(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	// Test forum-specific handlers
	tests := []struct {
		name    string
		handler func() *handlers.MessageHandler
	}{
		{"TopicCreated", func() *handlers.MessageHandler { return msgHandlers.TopicCreated(MockHandler) }},
		{"TopicEdited", func() *handlers.MessageHandler { return msgHandlers.TopicEdited(MockHandler) }},
		{"TopicClosed", func() *handlers.MessageHandler { return msgHandlers.TopicClosed(MockHandler) }},
		{"TopicReopened", func() *handlers.MessageHandler { return msgHandlers.TopicReopened(MockHandler) }},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := test.handler()
			if handler == nil {
				t.Errorf("%s should return a MessageHandler", test.name)
			}
		})
	}
}

func TestMessageHandlers_PaymentHandlers(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	// Test payment-related handlers
	tests := []struct {
		name    string
		handler func() *handlers.MessageHandler
	}{
		{"SuccessfulPayment", func() *handlers.MessageHandler { return msgHandlers.SuccessfulPayment(MockHandler) }},
		{"RefundedPayment", func() *handlers.MessageHandler { return msgHandlers.RefundedPayment(MockHandler) }},
		{"SuccessfulPaymentPrefix", func() *handlers.MessageHandler {
			return msgHandlers.SuccessfulPaymentPrefix(g.String("order_"), MockHandler)
		}},
		{"RefundedPaymentPrefix", func() *handlers.MessageHandler {
			return msgHandlers.RefundedPaymentPrefix(g.String("refund_"), MockHandler)
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := test.handler()
			if handler == nil {
				t.Errorf("%s should return a MessageHandler", test.name)
			}
		})
	}
}
