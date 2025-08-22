package bot_test

import (
	"testing"
	"time"

	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
)

func TestBot_GetMyShortDescription(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	getShortDesc := bot.GetMyShortDescription()

	if getShortDesc == nil {
		t.Error("Expected GetMyShortDescription to return a builder")
	}
}

func TestGetMyShortDescription_ChainedMethods(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	req := bot.GetMyShortDescription()

	// Test all chained methods for GetMyShortDescription
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
	req2 := bot.GetMyShortDescription().APIURL(g.String(""))
	if req2 == nil {
		t.Error("Expected APIURL with empty string to return request")
	}
}

func TestGetMyShortDescription_Send(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	req := bot.GetMyShortDescription()

	// Test Send method - expect it to fail with invalid token but increase coverage
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
