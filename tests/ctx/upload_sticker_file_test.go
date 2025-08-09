package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_UploadStickerFile(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)
	stickerFormat := g.String("static")

	result := ctx.UploadStickerFile(userID, stickerFormat)

	if result == nil {
		t.Error("Expected UploadStickerFile builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(time.Second * 30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}

func TestUploadStickerFile_File(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(456)
	stickerFormat := g.String("static")
	if ctx.UploadStickerFile(userID, stickerFormat).File(g.String("sticker.png")) == nil { t.Error("File should return builder") }
}

func TestUploadStickerFile_Format(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(456)
	stickerFormat := g.String("static")
	if ctx.UploadStickerFile(userID, stickerFormat).Format(g.String("animated")) == nil { t.Error("Format should return builder") }
}

func TestUploadStickerFile_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(456)
	stickerFormat := g.String("static")
	if ctx.UploadStickerFile(userID, stickerFormat).APIURL(g.String("https://api.example.com")) == nil { t.Error("APIURL should return builder") }
}

func TestUploadStickerFile_Send(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(456)
	stickerFormat := g.String("static")
	
	sendResult := ctx.UploadStickerFile(userID, stickerFormat).Format(g.String("animated")).Send()
	
	if sendResult.IsErr() {
		t.Logf("UploadStickerFile Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
