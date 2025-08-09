package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SetChatStickerSet(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	stickerSetName := g.String("test_sticker_set")

	result := ctx.SetChatStickerSet(stickerSetName)

	if result == nil {
		t.Error("Expected SetChatStickerSet builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(-1001234567890)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}

func TestSetChatStickerSet_Timeout(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	stickerSetName := g.String("test_sticker_set")
	if ctx.SetChatStickerSet(stickerSetName).Timeout(time.Minute) == nil { t.Error("Timeout should return builder") }
}

func TestSetChatStickerSet_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	stickerSetName := g.String("test_sticker_set")
	if ctx.SetChatStickerSet(stickerSetName).APIURL(g.String("https://api.example.com")) == nil { t.Error("APIURL should return builder") }
}

func TestSetChatStickerSet_Send(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	stickerSetName := g.String("test_sticker_set")
	
	sendResult := ctx.SetChatStickerSet(stickerSetName).Send()
	
	if sendResult.IsErr() {
		t.Logf("SetChatStickerSet Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
