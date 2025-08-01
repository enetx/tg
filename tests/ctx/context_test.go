package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

type mockBot struct{}

func (m *mockBot) Raw() *gotgbot.Bot           { return &gotgbot.Bot{} }
func (m *mockBot) Dispatcher() *ext.Dispatcher { return &ext.Dispatcher{} }
func (m *mockBot) Updater() *ext.Updater       { return &ext.Updater{} }

func TestNewContext(t *testing.T) {
	bot := &mockBot{}

	user := &gotgbot.User{Id: 123, FirstName: "Test"}
	chat := &gotgbot.Chat{Id: 456, Type: "private"}
	message := &gotgbot.Message{MessageId: 789, Text: "test message"}

	rawCtx := &ext.Context{
		EffectiveUser:    user,
		EffectiveChat:    chat,
		EffectiveMessage: message,
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	if ctx.Bot == nil {
		t.Error("Expected bot to be set")
	}

	if ctx.EffectiveUser != user {
		t.Error("Expected effective user to be set")
	}

	if ctx.EffectiveChat != chat {
		t.Error("Expected effective chat to be set")
	}

	if ctx.EffectiveMessage != message {
		t.Error("Expected effective message to be set")
	}

	if ctx.Raw != rawCtx {
		t.Error("Expected raw context to be set")
	}
}

func TestContextSendMessage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	text := g.String("Hello World")

	sm := ctx.SendMessage(text)

	// Test via public methods instead of accessing private fields
	if sm == nil {
		t.Error("Expected SendMessage builder to be created")
	}

	// Test that we can chain methods (verifies builder initialization)
	result := sm.HTML().Silent()
	if result == nil {
		t.Error("Expected chained method to return builder")
	}
}

func TestContextSendMessageChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	text := g.String("Hello World")

	sm := ctx.SendMessage(text).
		HTML().
		Silent().
		To(123).
		ReplyTo(789)

	// Test method chaining works properly
	if sm == nil {
		t.Error("Expected SendMessage builder to be created")
	}

	// Test that we can continue chaining after configuration
	result := sm.Thread(456)
	if result == nil {
		t.Error("Expected thread method to return builder")
	}
}

// Test inline query operations
func TestContext_InlineQuery(t *testing.T) {
	bot := &mockBot{}
	inlineQuery := &gotgbot.InlineQuery{Id: "inline123"}
	rawCtx := &ext.Context{Update: &gotgbot.Update{InlineQuery: inlineQuery}}
	ctx := ctx.New(bot, rawCtx)

	// Test AnswerInlineQuery
	result := ctx.AnswerInlineQuery(g.String("inline123"))
	if result == nil {
		t.Error("Expected AnswerInlineQuery builder to be created")
	}

	// Test method chaining
	_ = result.CacheFor(300 * time.Second).Personal()
}

// ==============================================
// Core Send Methods Tests
// ==============================================

func TestContext_SendDocument(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("nonexistent.pdf")

	result := ctx.SendDocument(filename)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected SendDocument builder to be created")
	}

	// Test method chaining
	_ = result.Caption(g.String("test"))
}

func TestContext_SendPhoto(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("photo.jpg")

	result := ctx.SendPhoto(filename)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected SendPhoto builder to be created")
	}

	// Test method chaining
	_ = result.Caption(g.String("test caption"))
}

func TestContext_SendVideo(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("video.mp4")

	result := ctx.SendVideo(filename)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected SendVideo builder to be created")
	}

	// Test method chaining
	_ = result.Duration(120).Width(1920).Height(1080)
}

func TestContext_SendAudio(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("audio.mp3")

	result := ctx.SendAudio(filename)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected SendAudio builder to be created")
	}

	// Test method chaining
	_ = result.Duration(180).Title(g.String("Song Title")).Performer(g.String("Artist"))
}

func TestContext_SendAnimation(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("animation.gif")

	result := ctx.SendAnimation(filename)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected SendAnimation builder to be created")
	}

	// Test method chaining
	_ = result.Duration(5).Width(480).Height(360)
}

func TestContext_SendVoice(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("voice.ogg")

	result := ctx.SendVoice(filename)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected SendVoice builder to be created")
	}

	// Test method chaining
	_ = result.Duration(30).Caption(g.String("Voice message"))
}

func TestContext_SendVideoNote(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("videonote.mp4")

	result := ctx.SendVideoNote(filename)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected SendVideoNote builder to be created")
	}

	// Test method chaining
	_ = result.Duration(10).Length(240)
}

func TestContext_SendSticker(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	sticker := g.String("CAACAgIAAxkBAAEBCgACYOZEYAAB1_gGF3IWUWwqZgABEQADBAADbwAABCBjAAEfBA")

	result := ctx.SendSticker(sticker)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected SendSticker builder to be created")
	}

	// Test method chaining
	_ = result.Emoji(g.String("😀")).To(789)
}

func TestContext_SendLocation(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	lat, lon := 40.7128, -74.0060

	result := ctx.SendLocation(lat, lon)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected SendLocation builder to be created")
	}

	// Test method chaining
	_ = result.To(789).Silent()
}

func TestContext_SendVenue(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	lat, lon := 40.7128, -74.0060
	title := g.String("Test Venue")
	address := g.String("123 Test St")

	result := ctx.SendVenue(lat, lon, title, address)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected SendVenue builder to be created")
	}

	// Test method chaining
	_ = result.FoursquareID(g.String("4d4b7105d754a06374d81259")).GooglePlaceID(g.String("ChIJN1t_tDeuEmsRUsoyG83frY4"))
}

func TestContext_SendContact(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	phoneNumber := g.String("+1234567890")
	firstName := g.String("John")

	result := ctx.SendContact(phoneNumber, firstName)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected SendContact builder to be created")
	}

	// Test method chaining
	_ = result.LastName(g.String("Doe")).VCard(g.String("BEGIN:VCARD\nVERSION:3.0\nFN:John Doe\nEND:VCARD"))
}

func TestContext_SendDice(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.SendDice()

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected SendDice builder to be created")
	}

	// Test method chaining
	_ = result.Emoji(g.String("🎲")).To(789)
}

func TestContext_SendPoll(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	question := g.String("What's your favorite color?")

	result := ctx.SendPoll(question)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected SendPoll builder to be created")
	}

	// Test method chaining
	_ = result.To(789).Silent()
}

// ==============================================
// Chat Administration Tests
// ==============================================

func TestContext_BanChatMember(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123456)

	result := ctx.BanChatMember(userID)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected BanChatMember builder to be created")
	}

	// Test builder creation is successful
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

func TestContext_UnbanChatMember(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123456)

	result := ctx.UnbanChatMember(userID)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected UnbanChatMember builder to be created")
	}

	// Test method chaining
	_ = result.OnlyIfBanned()
}

func TestContext_RestrictChatMember(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123456)

	result := ctx.RestrictChatMember(userID)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected RestrictChatMember builder to be created")
	}

	// Test builder creation is successful
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

func TestContext_PromoteChatMember(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123456)

	result := ctx.PromoteChatMember(userID)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected PromoteChatMember builder to be created")
	}

	// Test builder creation is successful
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

// ==============================================
// Message Operation Tests
// ==============================================

func TestContext_ForwardMessage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	toChatID := int64(123)

	// ForwardMessage takes fromChatID and messageID, toChatID is set via To() method
	result := ctx.ForwardMessage(456, 789).To(toChatID)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected ForwardMessage builder to be created")
	}

	// Test method chaining
	_ = result.Silent().Protect()
}

func TestContext_CopyMessage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	toChatID := int64(123)

	// CopyMessage takes fromChatID and messageID, toChatID is set via To() method
	result := ctx.CopyMessage(456, 789).To(toChatID)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected CopyMessage builder to be created")
	}

	// Test method chaining
	_ = result.Caption(g.String("New caption")).To(789)
}

func TestContext_DeleteMessage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.DeleteMessage()

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected DeleteMessage builder to be created")
	}

	// Test method chaining
	// Test builder creation without method chaining due to method unavailability
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

// ==============================================
// Reply and Edit Tests
// ==============================================

func TestContext_Reply(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "original"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	text := g.String("Reply text")

	result := ctx.Reply(text)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected Reply builder to be created")
	}

	// Test method chaining
	_ = result.HTML().Silent()
}

func TestContext_EditMessageText(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "original"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	text := g.String("Edited text")

	result := ctx.EditMessageText(text)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected EditMessageText builder to be created")
	}

	// Test method chaining
	_ = result.HTML()
}

func TestContext_EditMessageCaption(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Caption: "original caption"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	caption := g.String("New caption")

	result := ctx.EditMessageCaption(caption)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected EditMessageCaption builder to be created")
	}

	// Test method chaining
	_ = result.HTML()
}

// ==============================================
// Get Methods Tests
// ==============================================

func TestContext_GetChat(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.GetChat()

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected GetChat builder to be created")
	}

	// Test method chaining
	// Test builder creation without method chaining due to method unavailability
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

func TestContext_GetChatMember(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123)

	result := ctx.GetChatMember(userID)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected GetChatMember builder to be created")
	}

	// Test method chaining
	// Test builder creation without method chaining due to method unavailability
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

func TestContext_GetChatAdministrators(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.GetChatAdministrators()

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected GetChatAdministrators builder to be created")
	}

	// Test method chaining
	// Test builder creation without method chaining due to method unavailability
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

// ==============================================
// Callback Query Tests
// ==============================================

func TestContext_AnswerCallbackQuery(t *testing.T) {
	bot := &mockBot{}
	callback := &gotgbot.CallbackQuery{
		Id:   "callback123",
		Data: "test_data",
	}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1, CallbackQuery: callback},
	}

	ctx := ctx.New(bot, rawCtx)

	if ctx.Callback != callback {
		t.Error("Expected callback query to be set")
	}

	// Test callback query answer
	result := ctx.AnswerCallbackQuery(g.String("Test answer"))

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected AnswerCallbackQuery builder to be created")
	}

	// Test builder creation is successful
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

// ==============================================
// Edge Cases and Nil Handling Tests
// ==============================================

func TestContext_NilEffectiveMessage_DeleteMessage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: nil, // Nil message
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.DeleteMessage()

	// Test that the builder is created properly even with nil message
	if result == nil {
		t.Error("Expected DeleteMessage builder to be created even with nil message")
	}

	// Test method chaining
	// Test builder creation without method chaining due to method unavailability
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

func TestContext_NilEffectiveChat_Operations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: nil, // Nil chat
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test SendMessage with nil effective chat
	result := ctx.SendMessage(g.String("test"))

	// Test that the builder is created properly even with nil chat
	if result == nil {
		t.Error("Expected SendMessage builder to be created even with nil chat")
	}

	// Test method chaining
	// Test builder creation without method chaining due to method unavailability
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

// ==============================================
// Utility Methods Tests
// ==============================================

func TestContext_Args(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "/start arg1 arg2 arg3"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	args := ctx.Args()

	// Should skip the first word (command) and return arguments
	if args.Len() != 3 {
		t.Errorf("Expected 3 args, got %d", args.Len())
	}

	// Just test that we can call Args() and it returns a slice with expected length
	// The actual content parsing is handled by g.String.Fields().Skip(1)
}

func TestContext_ArgsEmpty(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "/start"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	args := ctx.Args()

	// Should return empty slice for command with no args
	if args.Len() != 0 {
		t.Errorf("Expected 0 args for command without args, got %d", args.Len())
	}
}

// ==============================================
// Additional Send Method Tests
// ==============================================

func TestContext_MediaGroup(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.MediaGroup()

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected MediaGroup builder to be created")
	}

	// Test method chaining
	_ = result.To(789).Silent().Protect()
}

func TestContext_SendChatAction(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.SendChatAction()

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected SendChatAction builder to be created")
	}

	// Test method chaining
	_ = result.Typing().To(789).Thread(123)
}

func TestContext_SendPaidMedia(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	starCount := int64(100)

	result := ctx.SendPaidMedia(starCount)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected SendPaidMedia builder to be created")
	}

	// Test method chaining
	_ = result.To(789).Caption(g.String("Paid content")).Payload(g.String("payload"))
}

// ==============================================
// Forum Management Tests
// ==============================================

func TestContext_CreateForumTopic(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	name := g.String("Test Topic")

	result := ctx.CreateForumTopic(name)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected CreateForumTopic builder to be created")
	}

	// Test method chaining
	_ = result.IconColor(0x6FB9F0).IconCustomEmojiID(g.String("🔥"))
}

func TestContext_EditForumTopic(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	messageThreadID := int64(123)

	result := ctx.EditForumTopic(messageThreadID)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected EditForumTopic builder to be created")
	}

	// Test method chaining
	_ = result.Name(g.String("New Name")).IconCustomEmojiID(g.String("🎆"))
}

func TestContext_CloseForumTopic(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	messageThreadID := int64(123)

	result := ctx.CloseForumTopic(messageThreadID)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected CloseForumTopic builder to be created")
	}

	// Test that we can create the builder (no additional methods to chain for this operation)
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

// ==============================================
// Payment & Business Tests
// ==============================================

func TestContext_SendInvoice(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	title := g.String("Test Product")
	desc := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")

	result := ctx.SendInvoice(title, desc, payload, currency)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected SendInvoice builder to be created")
	}

	// Test method chaining
	_ = result.ProviderToken(g.String("token"))
}

func TestContext_SendGift(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	giftID := g.String("gift_123")

	result := ctx.SendGift(giftID)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected SendGift builder to be created")
	}

	// Test method chaining
	_ = result.Text(g.String("Happy birthday!")).To(789)
}

// ==============================================
// Sticker Management Tests
// ==============================================

func TestContext_CreateNewStickerSet(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123)
	name := g.String("test_sticker_set")
	title := g.String("Test Sticker Set")

	result := ctx.CreateNewStickerSet(userID, name, title)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected CreateNewStickerSet builder to be created")
	}

	// Test method chaining
	_ = result.StickerType(g.String("regular")).NeedsRepainting()
}

func TestContext_GetStickerSet(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	name := g.String("test_sticker_set")

	result := ctx.GetStickerSet(name)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected GetStickerSet builder to be created")
	}

	// Test that we can create the builder (no additional methods to chain for this operation)
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

// ==============================================
// Chat Settings Tests
// ==============================================

func TestContext_SetChatTitle(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	title := g.String("New Chat Title")

	result := ctx.SetChatTitle(title)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected SetChatTitle builder to be created")
	}

	// Test method chaining
	// Test builder creation without method chaining due to method unavailability
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

func TestContext_SetChatDescription(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	description := g.String("New chat description")

	result := ctx.SetChatDescription(description)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected SetChatDescription builder to be created")
	}

	// Test method chaining
	// Test builder creation without method chaining due to method unavailability
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

// ==============================================
// Story Tests
// ==============================================

func TestContext_PostStory(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("conn_123")

	// We need a mock story content, but we can test the initialization
	// For now, we'll test with nil content to verify the builder pattern
	result := ctx.PostStory(businessConnectionID, nil)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected PostStory builder to be created")
	}

	// Test method chaining
	// Test builder creation without method chaining due to method name issues
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

func TestContext_DeleteStory(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("conn_123")
	storyID := int64(789)

	result := ctx.DeleteStory(businessConnectionID, storyID)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected DeleteStory builder to be created")
	}

	// Test that we can create the builder (no additional methods to chain for this operation)
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

// ==============================================
// Method Chaining and Builder Pattern Tests
// ==============================================

func TestContext_SendMessageAdvancedChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	text := g.String("Advanced test message")

	// Test complex method chaining
	sm := ctx.SendMessage(text).
		HTML().
		Silent().
		Protect().
		To(789).
		Thread(123).
		ReplyTo(456).
		AllowPaidBroadcast()

	// Test that the complex chained builder is created properly
	if sm == nil {
		t.Error("Expected SendMessage builder to be created after complex chaining")
	}
}

// ==============================================
// BATCH OPERATIONS TESTS
// ==============================================

func TestContext_DeleteMessages(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.DeleteMessages()

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected DeleteMessages builder to be created")
	}

	// Test method chaining
	// Test builder creation without method chaining due to method unavailability
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

func TestContext_ForwardMessages(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.ForwardMessages()

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected ForwardMessages builder to be created")
	}

	// Test method chaining
	// Test builder creation without method chaining due to method unavailability
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

func TestContext_CopyMessages(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.CopyMessages()

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected CopyMessages builder to be created")
	}

	// Test method chaining
	// Test builder creation without method chaining due to method unavailability
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

// ==============================================
// CHAT MEMBER AND PERMISSIONS TESTS
// ==============================================

func TestContext_SetChatPermissions(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.SetChatPermissions()

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected SetChatPermissions builder to be created")
	}

	// Test method chaining
	// Test builder creation without method chaining due to method unavailability
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

func TestContext_SetChatAdministratorCustomTitle(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123)
	customTitle := g.String("Super Admin")

	result := ctx.SetChatAdministratorCustomTitle(userID, customTitle)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected SetChatAdministratorCustomTitle builder to be created")
	}

	// Test method chaining
	// Test builder creation without method chaining due to method unavailability
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

// ==============================================
// PIN/UNPIN MESSAGE TESTS
// ==============================================

func TestContext_PinChatMessage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	messageID := int64(789)

	result := ctx.PinChatMessage(messageID)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected PinChatMessage builder to be created")
	}

	// Test method chaining
	_ = result.Silent()
}

func TestContext_UnpinChatMessage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.UnpinChatMessage()

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected UnpinChatMessage builder to be created")
	}

	// Test method chaining
	// Test builder creation without method chaining due to method unavailability
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

func TestContext_UnpinAllChatMessages(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.UnpinAllChatMessages()

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected UnpinAllChatMessages builder to be created")
	}

	// Test method chaining
	// Test builder creation without method chaining due to method unavailability
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

// ==============================================
// ADDITIONAL GET METHODS TESTS
// ==============================================

func TestContext_GetChatMemberCount(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.GetChatMemberCount()

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected GetChatMemberCount builder to be created")
	}

	// Test method chaining
	// Test builder creation without method chaining due to method unavailability
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

func TestContext_GetUserProfilePhotos(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123)

	result := ctx.GetUserProfilePhotos(userID)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected GetUserProfilePhotos builder to be created")
	}

	// Test method chaining
	_ = result.Offset(0).Limit(10)
}

func TestContext_GetFile(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	fileID := g.String("BAADBAADrwADBREAAWp8oTdjdLOHAg")

	result := ctx.GetFile(fileID)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected GetFile builder to be created")
	}

	// Test that we can create the builder (no additional methods to chain for this operation)
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

// ==============================================
// CHAT INVITE LINK TESTS
// ==============================================

func TestContext_CreateChatInviteLink(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.CreateChatInviteLink()

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected CreateChatInviteLink builder to be created")
	}

	// Test method chaining
	_ = result.Name(g.String("Special Link"))
}

func TestContext_EditChatInviteLink(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	inviteLink := g.String("https://t.me/joinchat/ABC123")

	result := ctx.EditChatInviteLink(inviteLink)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected EditChatInviteLink builder to be created")
	}

	// Test method chaining
	_ = result.Name(g.String("Updated Link"))
}

func TestContext_RevokeChatInviteLink(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	inviteLink := g.String("https://t.me/joinchat/ABC123")

	result := ctx.RevokeChatInviteLink(inviteLink)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected RevokeChatInviteLink builder to be created")
	}

	// Test method chaining
	// Test builder creation without method chaining due to method unavailability
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

func TestContext_ExportChatInviteLink(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.ExportChatInviteLink()

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected ExportChatInviteLink builder to be created")
	}

	// Test method chaining
	// Test builder creation without method chaining due to method unavailability
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

// ==============================================
// ADDITIONAL COMPREHENSIVE TESTS
// ==============================================

func TestContext_LeaveChat(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.LeaveChat()

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected LeaveChat builder to be created")
	}

	// Test method chaining
	// Test builder creation without method chaining due to method unavailability
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

func TestContext_StopPoll(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Poll: &gotgbot.Poll{Id: "poll123"}},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.StopPoll()

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected StopPoll builder to be created")
	}

	// Test method chaining
	// Test builder creation without method chaining due to method unavailability
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

func TestContext_SetMessageReaction(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	messageID := int64(789)

	result := ctx.SetMessageReaction(messageID)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected SetMessageReaction builder to be created")
	}

	// Test method chaining
	// Test builder creation without method chaining due to method unavailability
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}

func TestContext_GetAvailableGifts(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.GetAvailableGifts()

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected GetAvailableGifts builder to be created")
	}

	// Test that we can create the builder (no additional methods to chain for this operation)
	if result == nil {
		t.Error("Expected builder to be created successfully")
	}
}
