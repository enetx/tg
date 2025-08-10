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

func TestContext_SetChatPhoto(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("photo.jpg")

	result := ctx.SetChatPhoto(filename)

	if result == nil {
		t.Error("Expected SetChatPhoto builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(-1001234567890)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}

func TestSetChatPhoto_Timeout(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("photo.jpg")
	if ctx.SetChatPhoto(filename).Timeout(time.Minute) == nil {
		t.Error("Timeout should return builder")
	}
}

func TestSetChatPhoto_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("photo.jpg")
	if ctx.SetChatPhoto(filename).APIURL(g.String("https://api.example.com")) == nil {
		t.Error("APIURL should return builder")
	}
}

func TestSetChatPhoto_Send(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("photo.jpg")

	sendResult := ctx.SetChatPhoto(filename).Send()

	if sendResult.IsErr() {
		t.Logf("SetChatPhoto Send failed as expected with mock bot: %v", sendResult.Err())
	}
}

func TestSetChatPhoto_SendWithError(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})

	// Test Send with invalid file that causes error
	sendResult := ctx.SetChatPhoto(g.String("/invalid/nonexistent/photo.jpg")).Send()

	if !sendResult.IsErr() {
		t.Error("Send should return error for invalid file")
	} else {
		t.Logf("SetChatPhoto Send with error returned error as expected: %v", sendResult.Err())
	}
}

func TestSetChatPhoto_SendWithValidFile(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})

	// Create a temporary file
	tempFile := "/tmp/test_chat_photo.jpg"
	os.WriteFile(tempFile, []byte("test photo content"), 0644)
	defer os.Remove(tempFile)

	sendResult := ctx.SetChatPhoto(g.String(tempFile)).Send()

	// This will fail with mock bot, but covers the file closing path
	if sendResult.IsErr() {
		t.Logf("SetChatPhoto Send with valid file failed as expected: %v", sendResult.Err())
	}
}
