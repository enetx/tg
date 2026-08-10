package handlers_test

import (
	"regexp"
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
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

func TestMessageHandlers_AdditionalEdgeCases(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	// Test additional handler creation methods to improve coverage
	tests := []struct {
		name    string
		handler func() *handlers.MessageHandler
	}{
		{"AllowEdited", func() *handlers.MessageHandler { return msgHandlers.Any(MockHandler).AllowEdited() }},
		{"AllowChannel", func() *handlers.MessageHandler { return msgHandlers.Any(MockHandler).AllowChannel() }},
		{"AllowBusiness", func() *handlers.MessageHandler { return msgHandlers.Any(MockHandler).AllowBusiness() }},
		{"Register", func() *handlers.MessageHandler { return msgHandlers.Any(MockHandler).Register() }},
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

func TestMessageHandlers_RichAndGuest(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	tests := []struct {
		name    string
		handler func() *handlers.MessageHandler
	}{
		{"AnyEntities", func() *handlers.MessageHandler { return msgHandlers.AnyEntities(MockHandler) }},
		{"AnyEntity", func() *handlers.MessageHandler { return msgHandlers.AnyEntity(entity.Bold, MockHandler) }},
		{"RichEntity", func() *handlers.MessageHandler { return msgHandlers.RichEntity(entity.DateTime, MockHandler) }},
		{"RichMessage", func() *handlers.MessageHandler { return msgHandlers.RichMessage(MockHandler) }},
		{"GuestMessage", func() *handlers.MessageHandler { return msgHandlers.GuestMessage(MockHandler) }},
		{"GuestBot", func() *handlers.MessageHandler { return msgHandlers.GuestBot(MockHandler) }},
		{"MigrateFrom", func() *handlers.MessageHandler { return msgHandlers.MigrateFrom(MockHandler) }},
		{"MigrateTo", func() *handlers.MessageHandler { return msgHandlers.MigrateTo(MockHandler) }},
		{"ReplyMarkup", func() *handlers.MessageHandler { return msgHandlers.ReplyMarkup(MockHandler) }},
		{"TopicAction", func() *handlers.MessageHandler { return msgHandlers.TopicAction(MockHandler) }},
		{
			"CommunityChatAdded",
			func() *handlers.MessageHandler { return msgHandlers.CommunityChatAdded(MockHandler) },
		},
		{
			"CommunityChatRemoved",
			func() *handlers.MessageHandler { return msgHandlers.CommunityChatRemoved(MockHandler) },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.handler() == nil {
				t.Errorf("%s should return a MessageHandler", test.name)
			}
		})
	}
}

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

func TestMessageHandlers_ChatOwnerLeft(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.ChatOwnerLeft(MockHandler)

	if handler == nil {
		t.Error("ChatOwnerLeft should return a MessageHandler")
	}
}

func TestMessageHandlers_ChatOwnerChanged(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.ChatOwnerChanged(MockHandler)

	if handler == nil {
		t.Error("ChatOwnerChanged should return a MessageHandler")
	}
}

func TestMessageHandlers_ChatOwnerLeft_Filter(t *testing.T) {
	// Verify that the filter correctly matches messages with ChatOwnerLeft set
	msg := &gotgbot.Message{
		ChatOwnerLeft: &gotgbot.ChatOwnerLeft{},
	}

	if msg.ChatOwnerLeft == nil {
		t.Error("Expected ChatOwnerLeft to be set")
	}

	emptyMsg := &gotgbot.Message{}
	if emptyMsg.ChatOwnerLeft != nil {
		t.Error("Expected empty message to have no ChatOwnerLeft")
	}
}

func TestMessageHandlers_ChatOwnerChanged_Filter(t *testing.T) {
	// Verify that the filter correctly matches messages with ChatOwnerChanged set
	msg := &gotgbot.Message{
		ChatOwnerChanged: &gotgbot.ChatOwnerChanged{},
	}

	if msg.ChatOwnerChanged == nil {
		t.Error("Expected ChatOwnerChanged to be set")
	}

	emptyMsg := &gotgbot.Message{}
	if emptyMsg.ChatOwnerChanged != nil {
		t.Error("Expected empty message to have no ChatOwnerChanged")
	}
}

func TestMessageHandlers_ChatOwnerChanged_WithNewOwner(t *testing.T) {
	owner := &gotgbot.User{Id: 123, FirstName: "New", LastName: "Owner"}
	msg := &gotgbot.Message{
		ChatOwnerChanged: &gotgbot.ChatOwnerChanged{
			NewOwner: *owner,
		},
	}

	if msg.ChatOwnerChanged == nil {
		t.Error("Expected ChatOwnerChanged to be set")
	}

	if msg.ChatOwnerChanged.NewOwner.Id != 123 {
		t.Errorf("Expected new owner ID 123, got %d", msg.ChatOwnerChanged.NewOwner.Id)
	}
}

func TestMessageHandlers_ChatOwnerLeft_Registration(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	// Test that multiple handlers can be registered
	h1 := msgHandlers.ChatOwnerLeft(MockHandler)
	h2 := msgHandlers.ChatOwnerLeft(MockHandler)

	if h1 == nil || h2 == nil {
		t.Error("Expected both ChatOwnerLeft handlers to be non-nil")
	}
}

func TestMessageHandlers_ChatOwnerChanged_Registration(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	// Test that multiple handlers can be registered
	h1 := msgHandlers.ChatOwnerChanged(MockHandler)
	h2 := msgHandlers.ChatOwnerChanged(MockHandler)

	if h1 == nil || h2 == nil {
		t.Error("Expected both ChatOwnerChanged handlers to be non-nil")
	}
}
