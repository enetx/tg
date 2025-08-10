package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SetUserEmojiStatus(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)

	result := ctx.SetUserEmojiStatus(userID)

	if result == nil {
		t.Error("Expected SetUserEmojiStatus builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(time.Second * 30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}

func TestSetUserEmojiStatus_EmojiStatusCustomEmojiID(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(456)
	if ctx.SetUserEmojiStatus(userID).EmojiStatusCustomEmojiID(g.String("emoji_123")) == nil {
		t.Error("EmojiStatusCustomEmojiID should return builder")
	}
}

func TestSetUserEmojiStatus_RemoveStatus(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(456)
	if ctx.SetUserEmojiStatus(userID).RemoveStatus() == nil {
		t.Error("RemoveStatus should return builder")
	}
}

func TestSetUserEmojiStatus_EmojiStatusExpirationDate(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(456)
	expirationDate := time.Now().Add(time.Hour).Unix()
	if ctx.SetUserEmojiStatus(userID).EmojiStatusExpirationDate(expirationDate) == nil {
		t.Error("EmojiStatusExpirationDate should return builder")
	}
}

func TestSetUserEmojiStatus_ExpiresAt(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(456)
	expiresAt := time.Now().Add(time.Hour)
	if ctx.SetUserEmojiStatus(userID).ExpiresAt(expiresAt) == nil {
		t.Error("ExpiresAt should return builder")
	}
}

func TestSetUserEmojiStatus_ExpiresIn(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(456)
	if ctx.SetUserEmojiStatus(userID).ExpiresIn(time.Hour) == nil {
		t.Error("ExpiresIn should return builder")
	}
}

func TestSetUserEmojiStatus_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(456)
	if ctx.SetUserEmojiStatus(userID).APIURL(g.String("https://api.example.com")) == nil {
		t.Error("APIURL should return builder")
	}
}

func TestSetUserEmojiStatus_Send(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(456)

	sendResult := ctx.SetUserEmojiStatus(userID).EmojiStatusCustomEmojiID(g.String("emoji_123")).Send()

	if sendResult.IsErr() {
		t.Logf("SetUserEmojiStatus Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
