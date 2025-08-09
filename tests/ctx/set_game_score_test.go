package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SetGameScore(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)
	score := int64(1000)

	result := ctx.SetGameScore(userID, score)

	if result == nil {
		t.Error("Expected SetGameScore builder to be created")
	}

	// Test method chaining
	withForce := result.Force()
	if withForce == nil {
		t.Error("Expected Force method to return builder")
	}
}

func TestSetGameScore_UserID(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(456)
	score := int64(1000)
	if ctx.SetGameScore(userID, score).UserID(789) == nil { t.Error("UserID should return builder") }
}

func TestSetGameScore_Score(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(456)
	score := int64(1000)
	if ctx.SetGameScore(userID, score).Score(2000) == nil { t.Error("Score should return builder") }
}

func TestSetGameScore_DisableEditMessage(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(456)
	score := int64(1000)
	if ctx.SetGameScore(userID, score).DisableEditMessage() == nil { t.Error("DisableEditMessage should return builder") }
}

func TestSetGameScore_ChatID(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(456)
	score := int64(1000)
	if ctx.SetGameScore(userID, score).ChatID(789) == nil { t.Error("ChatID should return builder") }
}

func TestSetGameScore_MessageID(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(456)
	score := int64(1000)
	if ctx.SetGameScore(userID, score).MessageID(987) == nil { t.Error("MessageID should return builder") }
}

func TestSetGameScore_InlineMessageID(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(456)
	score := int64(1000)
	if ctx.SetGameScore(userID, score).InlineMessageID(g.String("inline_123")) == nil { t.Error("InlineMessageID should return builder") }
}

func TestSetGameScore_Timeout(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(456)
	score := int64(1000)
	if ctx.SetGameScore(userID, score).Timeout(time.Minute) == nil { t.Error("Timeout should return builder") }
}

func TestSetGameScore_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(456)
	score := int64(1000)
	if ctx.SetGameScore(userID, score).APIURL(g.String("https://api.example.com")) == nil { t.Error("APIURL should return builder") }
}

func TestSetGameScore_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 456},
		Update: &gotgbot.Update{UpdateId: 1},
	}
	ctx := ctx.New(bot, rawCtx)
	userID := int64(789)
	score := int64(1000)
	
	sendResult := ctx.SetGameScore(userID, score).Send()
	
	if sendResult.IsErr() {
		t.Logf("SetGameScore Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
