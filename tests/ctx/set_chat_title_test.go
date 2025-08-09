package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SetChatTitle(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	title := g.String("New Chat Title")

	result := ctx.SetChatTitle(title)

	if result == nil {
		t.Error("Expected SetChatTitle builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(-1001234567890)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}

func TestContext_SetChatTitleChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	title := g.String("New Chat Title")

	result := ctx.SetChatTitle(title).
		ChatID(-1001234567890)

	if result == nil {
		t.Error("Expected SetChatTitle builder to be created")
	}

	// Test that builder is functional
	_ = result
}

func TestSetChatTitle_Timeout(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	title := g.String("New Chat Title")
	if ctx.SetChatTitle(title).Timeout(time.Minute) == nil { t.Error("Timeout should return builder") }
}

func TestSetChatTitle_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	title := g.String("New Chat Title")
	if ctx.SetChatTitle(title).APIURL(g.String("https://api.example.com")) == nil { t.Error("APIURL should return builder") }
}

func TestSetChatTitle_Send(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	title := g.String("New Chat Title")
	
	sendResult := ctx.SetChatTitle(title).Send()
	
	if sendResult.IsErr() {
		t.Logf("SetChatTitle Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
