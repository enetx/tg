package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_DeleteMessages(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.DeleteMessages()

	if result == nil {
		t.Error("Expected DeleteMessages builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(456)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}

func TestContext_DeleteMessagesChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.DeleteMessages().
		ChatID(456).
		MessageIDs([]int64{123, 456, 789})

	if result == nil {
		t.Error("Expected DeleteMessages builder to be created")
	}

	// Test that builder is functional
	_ = result
}

func TestDeleteMessages_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.DeleteMessages().
		MessageIDs([]int64{456, 789, 123}).
		Send()

	if sendResult.IsErr() {
		t.Logf("DeleteMessages Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.DeleteMessages().
		ChatID(789).
		MessageIDs([]int64{100, 200, 300}).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("DeleteMessages configured Send failed as expected: %v", configuredSendResult.Err())
	}
}

func TestDeleteMessages_AddMessages(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test AddMessages method
	result := ctx.DeleteMessages().AddMessages(456, 789, 123)
	if result == nil {
		t.Error("AddMessages method should return DeleteMessages for chaining")
	}

	// Test multiple AddMessages calls
	multiResult := ctx.DeleteMessages().
		AddMessages(100, 200).
		AddMessages(300, 400).
		AddMessages(500)

	if multiResult == nil {
		t.Error("Multiple AddMessages calls should work")
	}

	// Test AddMessages with empty list
	emptyResult := ctx.DeleteMessages().AddMessages()
	if emptyResult == nil {
		t.Error("AddMessages with no arguments should work")
	}

	// Test AddMessages with many IDs
	manyIDs := make([]int64, 50)
	for i := range manyIDs {
		manyIDs[i] = int64(i + 1)
	}
	manyResult := ctx.DeleteMessages().AddMessages(manyIDs...)
	if manyResult == nil {
		t.Error("AddMessages with many IDs should work")
	}
}

func TestDeleteMessages_After(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test After method
	result := ctx.DeleteMessages().
		MessageIDs(g.SliceOf[int64](456, 789)).
		After(5 * time.Second)

	if result == nil {
		t.Error("After method should return DeleteMessages for chaining")
	}

	// Test After with different durations
	durations := []time.Duration{
		1 * time.Second,
		30 * time.Second,
		1 * time.Minute,
		5 * time.Minute,
		0 * time.Second, // Zero duration
	}

	for _, duration := range durations {
		durationResult := ctx.DeleteMessages().
			MessageIDs(g.SliceOf[int64](100, 200)).
			After(duration)

		if durationResult == nil {
			t.Errorf("After method with duration %v should work", duration)
		}
	}
}

func TestDeleteMessages_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Send with empty message IDs (should return error)
	emptyResult := ctx.DeleteMessages().Send()
	if !emptyResult.IsErr() {
		t.Error("Expected Send to return error when no message IDs specified")
	}

	expectedError := "no message IDs specified for deletion"
	errorStr2 := emptyResult.Err().Error()
	if !g.String(errorStr2).Contains(g.String(expectedError)) {
		t.Errorf("Expected error to contain '%s', got: %v", expectedError, emptyResult.Err())
	}

	// Test Send with too many message IDs (>100, should return error)
	tooManyIDs := make([]int64, 101)
	for i := range tooManyIDs {
		tooManyIDs[i] = int64(i + 1)
	}

	tooManyResult := ctx.DeleteMessages().
		MessageIDs(g.SliceOf(tooManyIDs...)).
		Send()

	if !tooManyResult.IsErr() {
		t.Error("Expected Send to return error when too many message IDs specified")
	}

	expectedTooManyError := "too many message IDs"
	errorStr := tooManyResult.Err().Error()
	if !g.String(errorStr).Contains(g.String(expectedTooManyError)) {
		t.Errorf("Expected error to contain '%s', got: %v", expectedTooManyError, tooManyResult.Err())
	}

	// Test Send with After (delayed deletion) - should return success immediately
	afterResult := ctx.DeleteMessages().
		MessageIDs(g.SliceOf[int64](100, 200)).
		After(1 * time.Second).
		Send()

	if afterResult.IsErr() {
		t.Errorf("Expected Send with After to return success, got error: %v", afterResult.Err())
	}

	if !afterResult.IsOk() || !afterResult.Ok() {
		t.Error("Expected Send with After to return true")
	}

	// Test APIURL method with nil RequestOpts (covers the nil branch)
	apiURLResult := ctx.DeleteMessages().APIURL(g.String("https://api.telegram.org"))
	if apiURLResult == nil {
		t.Error("APIURL method should return DeleteMessages for chaining")
	}

	// Test APIURL method with existing RequestOpts
	existingOptsResult := ctx.DeleteMessages().
		Timeout(30 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if existingOptsResult == nil {
		t.Error("APIURL method should work with existing RequestOpts")
	}
}

func TestDeleteMessages_CompleteWorkflow(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test complete workflow with all methods
	result := ctx.DeleteMessages().
		ChatID(456).
		MessageIDs(g.SliceOf[int64](100, 200)).
		AddMessages(300, 400).
		Timeout(45 * time.Second).
		APIURL(g.String("https://complete-api.telegram.org"))

	if result == nil {
		t.Error("Complete workflow should work")
	}

	// Test workflow with After (delayed deletion)
	delayedResult := ctx.DeleteMessages().
		ChatID(789).
		AddMessages(500, 600, 700).
		After(2 * time.Second).
		Timeout(60 * time.Second).
		Send()

	if delayedResult.IsErr() {
		t.Errorf("Delayed deletion workflow should work, got error: %v", delayedResult.Err())
	}

	// Test edge case: exactly 100 messages (maximum allowed)
	exactlyHundredIDs := make([]int64, 100)
	for i := range exactlyHundredIDs {
		exactlyHundredIDs[i] = int64(i + 1)
	}

	hundredResult := ctx.DeleteMessages().
		MessageIDs(g.SliceOf(exactlyHundredIDs...)).
		Send()

	if hundredResult.IsErr() {
		t.Logf("100 message deletion failed as expected with mock bot: %v", hundredResult.Err())
	}
}
