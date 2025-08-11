package ctx_test

import (
	"os"
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SetStickerSetThumbnail(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	name := g.String("test_sticker_set")
	userID := int64(456)

	result := ctx.SetStickerSetThumbnail(name, userID)

	if result == nil {
		t.Error("Expected SetStickerSetThumbnail builder to be created")
	}

	// Test method chaining
	withFormat := result.Format(g.String("static"))
	if withFormat == nil {
		t.Error("Expected Format method to return builder")
	}
}

func TestSetStickerSetThumbnail_Thumbnail(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	name := g.String("test_sticker_set")
	userID := int64(456)
	if ctx.SetStickerSetThumbnail(name, userID).Thumbnail(g.String("thumbnail.png")) == nil {
		t.Error("Thumbnail should return builder")
	}
}

func TestSetStickerSetThumbnail_Timeout(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	name := g.String("test_sticker_set")
	userID := int64(456)
	if ctx.SetStickerSetThumbnail(name, userID).Timeout(time.Minute) == nil {
		t.Error("Timeout should return builder")
	}
}

func TestSetStickerSetThumbnail_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	name := g.String("test_sticker_set")
	userID := int64(456)
	if ctx.SetStickerSetThumbnail(name, userID).APIURL(g.String("https://api.example.com")) == nil {
		t.Error("APIURL should return builder")
	}
}

func TestSetStickerSetThumbnail_Send(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	name := g.String("test_sticker_set")
	userID := int64(456)

	sendResult := ctx.SetStickerSetThumbnail(name, userID).Format(g.String("static")).Send()

	if sendResult.IsErr() {
		t.Logf("SetStickerSetThumbnail Send failed as expected with mock bot: %v", sendResult.Err())
	}
}

// Test Send method with error conditions for complete coverage
func TestSetStickerSetThumbnail_SendErrorConditions(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	name := g.String("test_sticker_set")
	userID := int64(456)

	// Test 1: Send with thumbnail file error (invalid file)
	result := ctx.SetStickerSetThumbnail(name, userID).
		Thumbnail(g.String("/nonexistent/file.png")).
		Format(g.String("static"))

	sendResult := result.Send()
	if !sendResult.IsErr() {
		t.Error("Expected error for invalid thumbnail file")
	} else {
		t.Logf("Send failed as expected with invalid thumbnail: %v", sendResult.Err())
	}

	// Test 2: Send without error (should reach the API call)
	result2 := ctx.SetStickerSetThumbnail(name, userID).
		Format(g.String("static")).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com"))

	sendResult2 := result2.Send()
	if sendResult2.IsOk() {
		t.Log("Send succeeded (unexpected with mock bot)")
	} else {
		t.Logf("Send failed as expected with mock bot: %v", sendResult2.Err())
	}

	// Test 3: Send with thumbnail file (covers file closing logic)
	tempFile := "/tmp/test_thumbnail.png"
	os.WriteFile(tempFile, []byte("test thumbnail content"), 0644)
	defer os.Remove(tempFile)

	result3 := ctx.SetStickerSetThumbnail(name, userID).
		Thumbnail(g.String(tempFile)).
		Format(g.String("static"))

	sendResult3 := result3.Send()
	if sendResult3.IsOk() {
		t.Log("Send with thumbnail file succeeded (unexpected with mock bot)")
	} else {
		t.Logf("Send with thumbnail file failed as expected: %v", sendResult3.Err())
	}
}
