package bot_test

import (
	"testing"
	"time"

	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/input"
)

func TestBot_SetMyProfilePhoto(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	b := result.Ok()
	req := b.SetMyProfilePhoto(input.StaticPhoto("file_id_123"))

	if req == nil {
		t.Error("Expected SetMyProfilePhoto to return a builder")
	}
}

func TestSetMyProfilePhoto_ChainedMethods(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	b := result.Ok()

	req := b.SetMyProfilePhoto(input.StaticPhoto("file_id_123"))

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

	req2 := b.SetMyProfilePhoto(input.StaticPhoto("file_id_456")).APIURL(g.String(""))
	if req2 == nil {
		t.Error("Expected APIURL with empty string to return builder")
	}
}

func TestSetMyProfilePhoto_WithAnimatedPhoto(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	b := result.Ok()

	req := b.SetMyProfilePhoto(input.AnimatedPhoto("animation_file_id"))
	if req == nil {
		t.Error("Expected SetMyProfilePhoto with animated photo to return builder")
	}

	req2 := b.SetMyProfilePhoto(input.AnimatedPhoto("anim_id").MainFrameTimestamp(0.5))
	if req2 == nil {
		t.Error("Expected SetMyProfilePhoto with animated photo (timestamp) to return builder")
	}
}

func TestSetMyProfilePhoto_Send(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	b := result.Ok()
	sendResult := b.SetMyProfilePhoto(input.StaticPhoto("file_id_123")).Send()

	if sendResult.IsOk() {
		t.Error("Expected Send to fail with invalid token")
	}

	if sendResult.IsErr() {
		t.Logf("Send failed as expected: %v", sendResult.Err())
	}
}

func TestSetMyProfilePhoto_Send_WithOptions(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	b := result.Ok()

	sendResult := b.SetMyProfilePhoto(input.StaticPhoto("file_id")).
		Timeout(5 * time.Second).
		APIURL(g.String("https://api.telegram.org")).
		Send()

	if sendResult.IsErr() {
		t.Logf("Send failed as expected: %v", sendResult.Err())
	}
}
