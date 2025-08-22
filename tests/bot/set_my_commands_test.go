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

	req = req.APIURL(g.String("https://api.telegram.org"))
	if req == nil {
		t.Error("Expected APIURL method to return request")
	}

	// Test APIURL with empty string for coverage
	req2 := bot.SetMyCommands().APIURL(g.String(""))
	if req2 == nil {
		t.Error("Expected APIURL with empty string to return request")
	}
}

func TestSetMyCommands_Send(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	botInstance := result.Ok()

	// Test Send with no commands - should return error
	reqEmpty := botInstance.SetMyCommands()
	resultEmpty := reqEmpty.Send()
	if resultEmpty.IsOk() {
		t.Error("Expected Send with no commands to fail, but it succeeded")
	} else {
		t.Logf("Send with no commands failed as expected: %v", resultEmpty.Err())
	}

	// Test Send with too many commands (>100) - should return error
	tooManyCommands := make([]gotgbot.BotCommand, 101)
	for i := 0; i < 101; i++ {
		tooManyCommands[i] = gotgbot.BotCommand{
			Command:     g.String("cmd" + g.Int(i).String().Std()).Std(),
			Description: "Description",
		}
	}
	reqTooMany := botInstance.SetMyCommands().Commands(g.SliceOf(tooManyCommands...))
	resultTooMany := reqTooMany.Send()
	if resultTooMany.IsOk() {
		t.Error("Expected Send with too many commands to fail, but it succeeded")
	} else {
		t.Logf("Send with too many commands failed as expected: %v", resultTooMany.Err())
	}

	// Test Send with valid commands - expect it to fail with invalid token but increase coverage
	commands := g.SliceOf(
		gotgbot.BotCommand{Command: "help", Description: "Show help"},
	)
	req := botInstance.SetMyCommands().Commands(commands)
	result2 := req.Send()
	if result2.IsOk() {
		t.Error("Expected Send to fail with invalid token, but it succeeded")
	}
	// We expect this to fail, so check that it failed
	if result2.IsErr() {
		// This is expected behavior with invalid token
		t.Logf("Send failed as expected: %v", result2.Err())
	}
}
