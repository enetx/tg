package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SetChatDescription(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	description := g.String("New chat description")

	result := ctx.SetChatDescription(description)

	if result == nil {
		t.Error("Expected SetChatDescription builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(-1001234567890)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}

func TestContext_SetChatDescriptionChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	description := g.String("New chat description")

	result := ctx.SetChatDescription(description).
		ChatID(-1001234567890)

	if result == nil {
		t.Error("Expected SetChatDescription builder to be created")
	}

	// Test that builder is functional
	_ = result
}

func TestSetChatDescription_Timeout(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	description := g.String("New chat description")
	if ctx.SetChatDescription(description).Timeout(time.Minute) == nil { t.Error("Timeout should return builder") }
}

func TestSetChatDescription_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	description := g.String("New chat description")
	if ctx.SetChatDescription(description).APIURL(g.String("https://api.example.com")) == nil { t.Error("APIURL should return builder") }
}

func TestSetChatDescription_Send(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	description := g.String("New chat description")
	
	sendResult := ctx.SetChatDescription(description).Send()
	
	if sendResult.IsErr() {
		t.Logf("SetChatDescription Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
