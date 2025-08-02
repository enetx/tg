package business_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/ctx/business"
)

type testBot struct {
	rawBot *gotgbot.Bot
}

func (tb *testBot) Raw() *gotgbot.Bot {
	return tb.rawBot
}

func TestBot_Interface(t *testing.T) {
	// Test that testBot implements the Bot interface
	var _ business.Bot = &testBot{}

	// Test actual usage
	bot := &testBot{rawBot: &gotgbot.Bot{}}
	rawBot := bot.Raw()

	if rawBot == nil {
		t.Error("Expected Raw() to return non-nil bot")
	}
}
