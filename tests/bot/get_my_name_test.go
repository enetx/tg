package bot_test

import (
	"testing"
	"time"

	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
)

func TestBot_GetMyName(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	getName := bot.GetMyName()

	if getName == nil {
		t.Error("Expected GetMyName to return a builder")
	}
}

func TestGetMyName_ChainedMethods(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	req := bot.GetMyName()

	// Test all chained methods for GetMyName
	req = req.Language("en")
	if req == nil {
		t.Error("Expected Language method to return request")
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
	req2 := bot.GetMyName().APIURL(g.String(""))
	if req2 == nil {
		t.Error("Expected APIURL with empty string to return request")
	}
}

func TestGetMyName_Send(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	botInstance := result.Ok()
	req := botInstance.GetMyName()

	result2 := req.Send()
	if result2.IsOk() {
		name := result2.Ok()
		_ = name
	} else {
		err := result2.Err()
		_ = err
	}
}
