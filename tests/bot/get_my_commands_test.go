package bot_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
)

func TestBot_GetMyCommands(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	getCommands := bot.GetMyCommands()

	if getCommands == nil {
		t.Error("Expected GetMyCommands to return a builder")
	}
}

func TestGetMyCommands_AllScopeMethods(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	req := bot.GetMyCommands()

	// Test Scope method with default scope
	req = req.Scope(gotgbot.BotCommandScopeDefault{})
	if req == nil {
		t.Error("Expected Scope method to return request")
	}

	// Test all scope types
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

func TestGetMyCommands_Send(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	botInstance := result.Ok()
	req := botInstance.GetMyCommands()

	// Test Send method
	result2 := req.Send()
	if result2.IsOk() {
		commands := result2.Ok()
		_ = commands
	} else {
		// Error expected in test environment
		err := result2.Err()
		_ = err
	}
}
