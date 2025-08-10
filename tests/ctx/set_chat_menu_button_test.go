package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SetChatMenuButton(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.SetChatMenuButton()

	if result == nil {
		t.Error("Expected SetChatMenuButton builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(123)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}

func TestSetChatMenuButton_MenuButton(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	menuButton := gotgbot.MenuButtonDefault{}
	if ctx.SetChatMenuButton().MenuButton(menuButton) == nil {
		t.Error("MenuButton should return builder")
	}
}

func TestSetChatMenuButton_DefaultMenu(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	if ctx.SetChatMenuButton().DefaultMenu() == nil {
		t.Error("DefaultMenu should return builder")
	}
}

func TestSetChatMenuButton_WebAppMenu(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	if ctx.SetChatMenuButton().WebAppMenu(g.String("Menu"), "https://example.com/webapp") == nil {
		t.Error("WebAppMenu should return builder")
	}
}

func TestSetChatMenuButton_CommandsMenu(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	if ctx.SetChatMenuButton().CommandsMenu() == nil {
		t.Error("CommandsMenu should return builder")
	}
}

func TestSetChatMenuButton_Timeout(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	if ctx.SetChatMenuButton().Timeout(time.Minute) == nil {
		t.Error("Timeout should return builder")
	}
}

func TestSetChatMenuButton_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	if ctx.SetChatMenuButton().APIURL(g.String("https://api.example.com")) == nil {
		t.Error("APIURL should return builder")
	}
}

func TestSetChatMenuButton_Send(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)

	sendResult := ctx.SetChatMenuButton().Send()

	if sendResult.IsErr() {
		t.Logf("SetChatMenuButton Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
