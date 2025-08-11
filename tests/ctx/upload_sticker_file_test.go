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
	if ctx.UploadStickerFile(userID, stickerFormat).File(g.String("sticker.png")) == nil {
		t.Error("File should return builder")
	}
}

func TestUploadStickerFile_Format(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(456)
	stickerFormat := g.String("static")
	if ctx.UploadStickerFile(userID, stickerFormat).Format(g.String("animated")) == nil {
		t.Error("Format should return builder")
	}
}

func TestUploadStickerFile_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(456)
	stickerFormat := g.String("static")
	if ctx.UploadStickerFile(userID, stickerFormat).APIURL(g.String("https://api.example.com")) == nil {
		t.Error("APIURL should return builder")
	}
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

// Test Send method with error conditions for complete coverage
func TestUploadStickerFile_SendErrorConditions(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(456)
	stickerFormat := g.String("static")

	// Test 1: Send with file error (invalid file)
	result := ctx.UploadStickerFile(userID, stickerFormat).
		File(g.String("/nonexistent/sticker.png")).
		Format(g.String("static"))

	sendResult := result.Send()
	if !sendResult.IsErr() {
		t.Error("Expected error for invalid sticker file")
	} else {
		t.Logf("Send failed as expected with invalid file: %v", sendResult.Err())
	}

	// Test 2: Send without file error (should reach the API call)
	result2 := ctx.UploadStickerFile(userID, stickerFormat).
		Format(g.String("animated")).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com"))

	sendResult2 := result2.Send()
	if sendResult2.IsOk() {
		t.Log("Send succeeded (unexpected with mock bot)")
	} else {
		t.Logf("Send failed as expected with mock bot: %v", sendResult2.Err())
	}

	// Test 3: Send with actual file (covers file closing logic)
	tempFile := "/tmp/test_sticker.png"
	os.WriteFile(tempFile, []byte("test sticker content"), 0644)
	defer os.Remove(tempFile)

	result3 := ctx.UploadStickerFile(userID, stickerFormat).
		File(g.String(tempFile)).
		Format(g.String("static"))

	sendResult3 := result3.Send()
	if sendResult3.IsOk() {
		t.Log("Send with file succeeded (unexpected with mock bot)")
	} else {
		t.Logf("Send with file failed as expected: %v", sendResult3.Err())
	}

	// Test 4: Test different format variations
	formats := []string{"static", "animated", "video"}
	for _, format := range formats {
		result4 := ctx.UploadStickerFile(userID, g.String(format)).
			Format(g.String(format))

		sendResult4 := result4.Send()
		if sendResult4.IsOk() {
			t.Logf("Send with format %s succeeded (unexpected)", format)
		} else {
			t.Logf("Send with format %s failed as expected: %v", format, sendResult4.Err())
		}
	}
}
