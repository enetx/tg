package bot_test

import (
	"testing"
	"time"

	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
)

func TestBot_SetMyShortDescription(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	setShortDesc := bot.SetMyShortDescription()

	if setShortDesc == nil {
		t.Error("Expected SetMyShortDescription to return a builder")
	}
}

func TestSetMyShortDescription_ChainedMethods(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	req := bot.SetMyShortDescription()

	// Test all chained methods for SetMyShortDescription
	req = req.Description(g.String("Short desc"))
	if req == nil {
		t.Error("Expected Description method to return request")
	}

	req = req.Language("en")
	if req == nil {
		t.Error("Expected Language method to return request")
	}

	req = req.Remove()
	if req == nil {
		t.Error("Expected Remove method to return request")
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
