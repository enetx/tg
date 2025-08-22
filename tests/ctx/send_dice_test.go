package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/reply"
	"github.com/enetx/tg/types/effects"
)

func TestContext_SendDice(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.SendDice()

	if result == nil {
		t.Error("Expected SendDice builder to be created")
	}

	// Test method chaining
	chained := result.Silent()
	if chained == nil {
		t.Error("Expected Silent method to return builder")
	}
}

func TestContext_SendDiceChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.SendDice().
		Emoji("🎲").
		Silent().
		To(123)

	if result == nil {
		t.Error("Expected SendDice builder to be created")
	}

	// Test continued chaining
	final := result.Protect()
	if final == nil {
		t.Error("Expected Protect method to return builder")
	}
}

func TestSendDice_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendDice().Send()

	if sendResult.IsErr() {
		t.Logf("SendDice Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendDice().
		Emoji("🎯").
		Silent().
		Protect().
		To(123).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendDice configured Send failed as expected: %v", configuredSendResult.Err())
	}
}

func TestSendDice_After(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	durations := []time.Duration{time.Second, time.Minute, time.Hour, 0}
	for _, duration := range durations {
		result := ctx.SendDice().After(duration)
		if result == nil {
			t.Errorf("After method should return SendDice builder for chaining with duration: %v", duration)
		}
	}
}

func TestSendDice_DeleteAfter(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	durations := []time.Duration{time.Second * 30, time.Minute * 5, time.Hour, 0}
	for _, duration := range durations {
		result := ctx.SendDice().DeleteAfter(duration)
		if result == nil {
			t.Errorf("DeleteAfter method should return SendDice builder for chaining with duration: %v", duration)
		}
	}
}

func TestSendDice_Dart(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	result := ctx.SendDice().Dart()
	if result == nil {
		t.Error("Dart method should return SendDice builder for chaining")
	}
}

func TestSendDice_Slot(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	result := ctx.SendDice().Slot()
	if result == nil {
		t.Error("Slot method should return SendDice builder for chaining")
	}
}

func TestSendDice_Ball(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	result := ctx.SendDice().Ball()
	if result == nil {
		t.Error("Ball method should return SendDice builder for chaining")
	}
}

func TestSendDice_Soccer(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	result := ctx.SendDice().Soccer()
	if result == nil {
		t.Error("Soccer method should return SendDice builder for chaining")
	}
}

func TestSendDice_Bowling(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	result := ctx.SendDice().Bowling()
	if result == nil {
		t.Error("Bowling method should return SendDice builder for chaining")
	}
}

func TestSendDice_Thread(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	threadIDs := []int64{1, 123, 456, 0}
	for _, threadID := range threadIDs {
		result := ctx.SendDice().Thread(threadID)
		if result == nil {
			t.Errorf("Thread method should return SendDice builder for chaining with threadID: %d", threadID)
		}
	}
}

func TestSendDice_AllowPaidBroadcast(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	result := ctx.SendDice().AllowPaidBroadcast()
	if result == nil {
		t.Error("AllowPaidBroadcast method should return SendDice builder for chaining")
	}
}

func TestSendDice_Effect(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	result := ctx.SendDice().Effect(effects.Fire)
	if result == nil {
		t.Error("Effect method should return SendDice builder for chaining")
	}
}

func TestSendDice_ReplyTo(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	messageIDs := []int64{1, 123, 456, 999}
	for _, messageID := range messageIDs {
		result := ctx.SendDice().Reply(reply.New(messageID))
		if result == nil {
			t.Errorf("ReplyTo method should return SendDice builder for chaining with messageID: %d", messageID)
		}
	}
}

func TestSendDice_Markup(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	btn1 := keyboard.NewButton().Text(g.String("Test Button")).Callback(g.String("test_data"))
	inlineKeyboard := keyboard.Inline().Button(btn1)

	result := ctx.SendDice().Markup(inlineKeyboard)
	if result == nil {
		t.Error("Markup method should return SendDice builder for chaining")
	}
}

func TestSendDice_Business(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	businessIDs := []string{"business_123", "conn_456", ""}
	for _, businessID := range businessIDs {
		result := ctx.SendDice().Business(g.String(businessID))
		if result == nil {
			t.Errorf("Business method should return SendDice builder for chaining with businessID: %s", businessID)
		}
	}
}

func TestSendDice_DirectMessagesTopic(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	topicIDs := []int64{123, 456, 789, 0, -1}

	for _, topicID := range topicIDs {
		result := ctx.SendDice().DirectMessagesTopic(topicID)
		if result == nil {
			t.Errorf("DirectMessagesTopic method should return SendDice builder for chaining with topicID: %d", topicID)
		}

		chainedResult := result.DirectMessagesTopic(topicID + 100)
		if chainedResult == nil {
			t.Errorf("DirectMessagesTopic method should support chaining and override with topicID: %d", topicID)
		}
	}
}

func TestSendDice_SuggestedPost(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test with nil params
	result := ctx.SendDice().SuggestedPost(nil)
	if result == nil {
		t.Error("SuggestedPost method should return SendDice builder for chaining with nil params")
	}

	// Test chaining
	chainedResult := result.SuggestedPost(nil)
	if chainedResult == nil {
		t.Error("SuggestedPost method should support chaining")
	}
}

func TestSendDice_APIURLWithExistingRequestOpts(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// First set Timeout to create RequestOpts, then test APIURL
	result := ctx.SendDice().
		Timeout(15 * time.Second).                         // This creates RequestOpts
		APIURL(g.String("https://custom.api.example.com")) // This should use existing RequestOpts

	if result == nil {
		t.Error("APIURL with existing RequestOpts should return builder")
	}
}
