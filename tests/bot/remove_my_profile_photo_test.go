package bot_test

import (
	"testing"
	"time"

	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
)

func TestBot_RemoveMyProfilePhoto(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	b := result.Ok()
	req := b.RemoveMyProfilePhoto()

	if req == nil {
		t.Error("Expected RemoveMyProfilePhoto to return a builder")
	}
}

func TestRemoveMyProfilePhoto_ChainedMethods(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	b := result.Ok()
	req := b.RemoveMyProfilePhoto()

	req = req.Timeout(10 * time.Second)
	if req == nil {
		t.Error("Expected Timeout to return builder")
	}

	req = req.Timeout(20 * time.Second) // second call - covers non-nil RequestOpts branch
	if req == nil {
		t.Error("Expected Timeout (second) to return builder")
	}

	req = req.APIURL(g.String("https://api.telegram.org"))
	if req == nil {
		t.Error("Expected APIURL to return builder")
	}

	req = req.APIURL(g.String("https://api.telegram.org")) // second call - covers non-nil branch
	if req == nil {
		t.Error("Expected APIURL (second) to return builder")
	}

	req2 := b.RemoveMyProfilePhoto().APIURL(g.String(""))
	if req2 == nil {
		t.Error("Expected APIURL with empty string to return builder")
	}
}

func TestRemoveMyProfilePhoto_Send(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	b := result.Ok()
	sendResult := b.RemoveMyProfilePhoto().Send()

	if sendResult.IsOk() {
		t.Error("Expected Send to fail with invalid token")
	}

	if sendResult.IsErr() {
		t.Logf("Send failed as expected: %v", sendResult.Err())
	}
}

func TestRemoveMyProfilePhoto_Send_WithOptions(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	b := result.Ok()

	timeouts := []time.Duration{
		1 * time.Second,
		5 * time.Second,
		15 * time.Second,
		30 * time.Second,
	}

	for _, timeout := range timeouts {
		sendResult := b.RemoveMyProfilePhoto().
			Timeout(timeout).
			APIURL(g.String("https://api.telegram.org")).
			Send()

		if sendResult.IsErr() {
			t.Logf("Send with timeout %v failed as expected: %v", timeout, sendResult.Err())
		}
	}
}
