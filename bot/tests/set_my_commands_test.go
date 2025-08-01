package bot_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
)

func TestBot_SetMyCommands(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	setCommands := bot.SetMyCommands()

	if setCommands == nil {
		t.Error("Expected SetMyCommands to return a builder")
	}
}

func TestSetMyCommands_AllMethods(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	req := bot.SetMyCommands()

	// Test AddCommand method
	req = req.AddCommand("start", "Start the bot")
	if req == nil {
		t.Error("Expected AddCommand method to return request")
	}

	// Test Commands method
	commands := g.SliceOf(
		gotgbot.BotCommand{Command: "help", Description: "Show help"},
		gotgbot.BotCommand{Command: "settings", Description: "Show settings"},
	)
	req = req.Commands(commands)
	if req == nil {
		t.Error("Expected Commands method to return request")
	}

	// Test all scope methods (same as GetMyCommands)
	req = req.Scope(gotgbot.BotCommandScopeDefault{})
	if req == nil {
		t.Error("Expected Scope method to return request")
	}

	req = req.ScopeDefault()
	if req == nil {
		t.Error("Expected ScopeDefault method to return request")
	}

	req = req.ScopeAllPrivateChats()
	if req == nil {
		t.Error("Expected ScopeAllPrivateChats method to return request")
	}

	req = req.ScopeAllGroupChats()
	if req == nil {
		t.Error("Expected ScopeAllGroupChats method to return request")
	}

	req = req.ScopeAllChatAdministrators()
	if req == nil {
		t.Error("Expected ScopeAllChatAdministrators method to return request")
	}

	req = req.ScopeChat(123)
	if req == nil {
		t.Error("Expected ScopeChat method to return request")
	}

	req = req.ScopeChatAdministrators(456)
	if req == nil {
		t.Error("Expected ScopeChatAdministrators method to return request")
	}

	req = req.ScopeChatMember(789, 101)
	if req == nil {
		t.Error("Expected ScopeChatMember method to return request")
	}

	req = req.LanguageCode("en")
	if req == nil {
		t.Error("Expected LanguageCode method to return request")
	}

	req = req.Timeout(10 * time.Second)
	if req == nil {
		t.Error("Expected Timeout method to return request")
	}

	req = req.APIURL("https://api.telegram.org")
	if req == nil {
		t.Error("Expected APIURL method to return request")
	}
}
