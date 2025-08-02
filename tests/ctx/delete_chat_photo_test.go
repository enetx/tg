package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_DeleteChatPhoto(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test basic creation
	result := testCtx.DeleteChatPhoto()
	if result == nil {
		t.Error("Expected DeleteChatPhoto builder to be created")
	}

	// Test ChatID method
	result = result.ChatID(-1001234567890)
	if result == nil {
		t.Error("ChatID method should return DeleteChatPhoto for chaining")
	}

	// Test Timeout method
	result = testCtx.DeleteChatPhoto().Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return DeleteChatPhoto for chaining")
	}

	// Test APIURL method
	result = testCtx.DeleteChatPhoto().APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return DeleteChatPhoto for chaining")
	}
}

func TestDeleteChatPhoto_CompleteChain(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test complete method chaining
	result := testCtx.DeleteChatPhoto().
		ChatID(-1001987654321).
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return DeleteChatPhoto")
	}
}

func TestDeleteChatPhoto_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test with zero ChatID (should use effective chat)
	result := testCtx.DeleteChatPhoto().ChatID(0)
	if result == nil {
		t.Error("DeleteChatPhoto should handle zero ChatID")
	}

	// Test with negative ChatID (supergroup/channel)
	result = testCtx.DeleteChatPhoto().ChatID(-1001111111111)
	if result == nil {
		t.Error("DeleteChatPhoto should handle negative ChatID")
	}

	// Test with positive ChatID (user/group)
	result = testCtx.DeleteChatPhoto().ChatID(123456789)
	if result == nil {
		t.Error("DeleteChatPhoto should handle positive ChatID")
	}

	// Test with zero timeout
	result = testCtx.DeleteChatPhoto().Timeout(0 * time.Second)
	if result == nil {
		t.Error("DeleteChatPhoto should handle zero timeout")
	}

	// Test with empty API URL
	result = testCtx.DeleteChatPhoto().APIURL(g.String(""))
	if result == nil {
		t.Error("DeleteChatPhoto should handle empty API URL")
	}
}

func TestDeleteChatPhoto_DefaultChatID(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test without setting ChatID (should use EffectiveChat.Id)
	result := testCtx.DeleteChatPhoto()
	if result == nil {
		t.Error("DeleteChatPhoto should work without explicit ChatID")
	}

	// The Send() method should use EffectiveChat.Id when no ChatID is set
	// We can't test the actual API call with mockBot, but we can ensure the builder works
}
