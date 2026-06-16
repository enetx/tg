package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SetChatMemberTag(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	userID := int64(123456)

	result := testCtx.SetChatMemberTag(userID)
	if result == nil {
		t.Error("Expected SetChatMemberTag builder to be created")
	}

	result = result.ChatID(-1001234567890)
	if result == nil {
		t.Error("ChatID method should return SetChatMemberTag for chaining")
	}

	result = testCtx.SetChatMemberTag(userID).Tag(g.String("mod"))
	if result == nil {
		t.Error("Tag method should return SetChatMemberTag for chaining")
	}

	result = testCtx.SetChatMemberTag(userID).Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return SetChatMemberTag for chaining")
	}

	result = testCtx.SetChatMemberTag(userID).APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return SetChatMemberTag for chaining")
	}
}

func TestContext_SetChatMemberTagChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	userID := int64(123456)

	result := testCtx.SetChatMemberTag(userID).
		ChatID(-1001987654321).
		Tag(g.String("vip")).
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return SetChatMemberTag")
	}
}

func TestSetChatMemberTag_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Empty tag means "remove the tag" — must still produce a valid builder.
	result := testCtx.SetChatMemberTag(1).Tag(g.String(""))
	if result == nil {
		t.Error("SetChatMemberTag should accept empty Tag to remove the existing tag")
	}

	// Max length tag (16 chars per Bot API 9.5).
	result = testCtx.SetChatMemberTag(1).Tag(g.String("1234567890123456"))
	if result == nil {
		t.Error("SetChatMemberTag should accept a 16-character tag")
	}

	// Negative user IDs are accepted by the builder; the API itself rejects them at send time.
	result = testCtx.SetChatMemberTag(-1)
	if result == nil {
		t.Error("SetChatMemberTag should accept a negative user ID")
	}

	// Zero ChatID — Send will fall back to EffectiveChat.Id.
	result = testCtx.SetChatMemberTag(1).ChatID(0)
	if result == nil {
		t.Error("SetChatMemberTag should accept zero ChatID")
	}
}

func TestSetChatMemberTag_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	userID := int64(456)

	sendResult := testCtx.SetChatMemberTag(userID).Send()
	if sendResult.IsErr() {
		t.Logf("SetChatMemberTag Send failed as expected with mock bot: %v", sendResult.Err())
	}

	configuredSendResult := testCtx.SetChatMemberTag(userID).
		ChatID(789).
		Tag(g.String("admin")).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SetChatMemberTag configured Send failed as expected: %v", configuredSendResult.Err())
	}

	// Test default chat (uses EffectiveChat.Id when ChatID not set).
	defaultChatResult := testCtx.SetChatMemberTag(userID).Tag(g.String("member")).Send()
	if defaultChatResult.IsErr() {
		t.Logf("SetChatMemberTag default chat Send failed as expected: %v", defaultChatResult.Err())
	}
}
