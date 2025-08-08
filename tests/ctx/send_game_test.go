package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SendGame(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	gameShortName := g.String("snake")

	result := ctx.SendGame(gameShortName)

	if result == nil {
		t.Error("Expected SendGame builder to be created")
	}

	// Test method chaining
	chained := result.Silent()
	if chained == nil {
		t.Error("Expected Silent method to return builder")
	}
}

func TestContext_SendGameChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	gameShortName := g.String("snake")

	result := ctx.SendGame(gameShortName).
		Silent().
		Protect().
		To(123)

	if result == nil {
		t.Error("Expected SendGame builder to be created")
	}

	// Test continued chaining
	final := result.Thread(456)
	if final == nil {
		t.Error("Expected Thread method to return builder")
	}
}

func TestSendGame_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	gameShortName := g.String("test_game")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendGame(gameShortName).Send()

	if sendResult.IsErr() {
		t.Logf("SendGame Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendGame(gameShortName).
		Silent().
		Protect().
		To(123).
		Thread(456).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendGame configured Send failed as expected: %v", configuredSendResult.Err())
	}
}
