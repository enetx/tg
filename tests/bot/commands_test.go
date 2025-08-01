package bot_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func TestBot_Command(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()

	// Test Command method
	handler := func(ctx *ctx.Context) error {
		return nil
	}

	cmd := bot.Command(g.String("start"), handler)

	if cmd == nil {
		t.Error("Expected Command to return a command handler")
	}
}
