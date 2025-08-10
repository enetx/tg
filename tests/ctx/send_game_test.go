package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/types/effects"
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

func TestSendGame_After(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	if ctx.SendGame(g.String("test_game")).After(time.Minute) == nil {
		t.Error("After should return builder")
	}
}

func TestSendGame_DeleteAfter(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	if ctx.SendGame(g.String("test_game")).DeleteAfter(time.Hour) == nil {
		t.Error("DeleteAfter should return builder")
	}
}

func TestSendGame_AllowPaidBroadcast(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	if ctx.SendGame(g.String("test_game")).AllowPaidBroadcast() == nil {
		t.Error("AllowPaidBroadcast should return builder")
	}
}

func TestSendGame_Effect(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	if ctx.SendGame(g.String("test_game")).Effect(effects.Fire) == nil {
		t.Error("Effect should return builder")
	}
}

func TestSendGame_ReplyTo(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	if ctx.SendGame(g.String("test_game")).ReplyTo(123) == nil {
		t.Error("ReplyTo should return builder")
	}
}

func TestSendGame_Markup(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	btn1 := keyboard.NewButton().Text(g.String("Play")).Callback(g.String("play"))
	if ctx.SendGame(g.String("test_game")).Markup(keyboard.Inline().Button(btn1)) == nil {
		t.Error("Markup should return builder")
	}
}

func TestSendGame_Business(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	if ctx.SendGame(g.String("test_game")).Business(g.String("biz_123")) == nil {
		t.Error("Business should return builder")
	}
}

func TestSendGame_APIURLWithExistingRequestOpts(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// First set Timeout to create RequestOpts, then test APIURL
	result := ctx.SendGame(g.String("test_game")).
		Timeout(15 * time.Second).                         // This creates RequestOpts
		APIURL(g.String("https://custom.api.example.com")) // This should use existing RequestOpts

	if result == nil {
		t.Error("APIURL with existing RequestOpts should return builder")
	}
}
