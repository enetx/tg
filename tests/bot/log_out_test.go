package bot_test

import (
	"testing"
	"time"

	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
)

func TestBot_LogOut(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	logOut := bot.LogOut()

	if logOut == nil {
		t.Error("Expected LogOut to return a builder")
	}
}

func TestLogOut_AllMethods(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	req := bot.LogOut()

	// Test all methods
	req = req.Timeout(10 * time.Second)
	if req == nil {
		t.Error("Expected Timeout method to return request")
	}

	req = req.APIURL(g.String("https://api.telegram.org"))
	if req == nil {
		t.Error("Expected APIURL method to return request")
	}

	// Test APIURL with empty string for coverage
	req2 := bot.LogOut().APIURL(g.String(""))
	if req2 == nil {
		t.Error("Expected APIURL with empty string to return request")
	}
}

func TestLogOut_Send(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	botInstance := result.Ok()
	req := botInstance.LogOut()

	result2 := req.Send()
	if result2.IsOk() {
		success := result2.Ok()
		_ = success
	} else {
		err := result2.Err()
		_ = err
	}
}
