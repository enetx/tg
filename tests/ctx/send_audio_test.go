package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SendAudio(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("audio.mp3")

	result := ctx.SendAudio(filename)

	if result == nil {
		t.Error("Expected SendAudio builder to be created")
	}

	// Test method chaining
	chained := result.Caption(g.String("Audio caption"))
	if chained == nil {
		t.Error("Expected caption method to return builder")
	}
}

func TestContext_SendAudioChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("audio.mp3")

	result := ctx.SendAudio(filename).
		Caption(g.String("Test audio")).
		HTML().
		Silent().
		Duration(180)

	if result == nil {
		t.Error("Expected SendAudio builder to be created")
	}

	// Test continued chaining
	final := result.To(123)
	if final == nil {
		t.Error("Expected To method to return builder")
	}
}

func TestSendAudio_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("test_audio.mp3")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendAudio(filename).Send()

	if sendResult.IsErr() {
		t.Logf("SendAudio Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendAudio(filename).
		Caption(g.String("Test <b>audio</b> with HTML")).
		HTML().
		Duration(180).
		Performer(g.String("Test Artist")).
		Title(g.String("Test Song")).
		Silent().
		Protect().
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendAudio configured Send failed as expected: %v", configuredSendResult.Err())
	}
}
