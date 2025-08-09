package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_VerifyChat(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	chatID := int64(-1001234567890)

	result := ctx.VerifyChat(chatID)

	if result == nil {
		t.Error("Expected VerifyChat builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(time.Second * 30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}

func TestVerifyChat_CustomDescription(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	chatID := int64(456)
	if ctx.VerifyChat(chatID).CustomDescription(g.String("Verified chat")) == nil { t.Error("CustomDescription should return builder") }
}

func TestVerifyChat_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	chatID := int64(456)
	if ctx.VerifyChat(chatID).APIURL(g.String("https://api.example.com")) == nil { t.Error("APIURL should return builder") }
}

func TestVerifyChat_Send(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	chatID := int64(456)
	
	sendResult := ctx.VerifyChat(chatID).CustomDescription(g.String("Test verification")).Send()
	
	if sendResult.IsErr() {
		t.Logf("VerifyChat Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
