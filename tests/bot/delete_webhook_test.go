package bot_test

import (
	"testing"
	"time"

	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
)

func TestBot_DeleteWebhook(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	deleteWebhook := bot.DeleteWebhook()

	if deleteWebhook == nil {
		t.Error("Expected DeleteWebhook to return a builder")
	}
}

func TestDeleteWebhook_AllMethods(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	req := bot.DeleteWebhook()

	// Test all methods
	req = req.DropPendingUpdates()
	if req == nil {
		t.Error("Expected DropPendingUpdates method to return request")
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

func TestDeleteWebhook_Send(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	botInstance := result.Ok()
	req := botInstance.DeleteWebhook()

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
