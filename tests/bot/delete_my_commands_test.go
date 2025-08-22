package bot_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
)

func TestBot_DeleteMyCommands(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	deleteCommands := bot.DeleteMyCommands()

	if deleteCommands == nil {
		t.Error("Expected DeleteMyCommands to return a builder")
	}
}

// Test remaining API methods for comprehensive coverage
func TestDeleteMyCommands_AllMethods(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	req := bot.DeleteMyCommands()

	// Test all scope methods
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

	req = req.APIURL(g.String("https://api.telegram.org"))
	if req == nil {
		t.Error("Expected APIURL method to return request")
	}

	// Test APIURL with empty string for coverage
	req2 := bot.DeleteMyCommands().APIURL(g.String(""))
	if req2 == nil {
		t.Error("Expected APIURL with empty string to return request")
	}
}

func TestDeleteMyCommands_Send(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	botInstance := result.Ok()
	req := botInstance.DeleteMyCommands()

	// Test Send method
	result2 := req.Send()
	if result2.IsOk() {
		success := result2.Ok()
		_ = success
	} else {
		// Error expected in test environment
		err := result2.Err()
		_ = err
	}
}
