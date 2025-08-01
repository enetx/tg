package bot_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
)

func TestBot_GetMe(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	getMe := bot.GetMe()

	if getMe == nil {
		t.Error("Expected GetMe to return a builder")
	}
}
