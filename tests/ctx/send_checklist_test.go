package ctx_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
)

func TestContext_SendChecklist(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	title := g.String("My Checklist")
	intro := g.String("This is my checklist")

	result := ctx.SendChecklist(title, intro)

	if result == nil {
		t.Error("Expected SendChecklist builder to be created")
	}

	// Test method chaining
	chained := result.Silent()
	if chained == nil {
		t.Error("Expected Silent method to return builder")
	}
}

func TestContext_SendChecklistChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	title := g.String("My Checklist")
	intro := g.String("This is my checklist")

	result := ctx.SendChecklist(title, intro).
		Silent().
		Protect().
		To(123)

	if result == nil {
		t.Error("Expected SendChecklist builder to be created")
	}

	// Test that builder is functional
	_ = result
}

func TestSendChecklist_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	title := g.String("Test Checklist")
	intro := g.String("Test checklist description")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendChecklist(title, intro).Send()

	if sendResult.IsErr() {
		t.Logf("SendChecklist Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendChecklist(title, intro).
		Silent().
		Protect().
		To(123).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendChecklist configured Send failed as expected: %v", configuredSendResult.Err())
	}
}

func TestSendChecklist_After(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	title := g.String("Test Checklist")
	intro := g.String("Test intro")

	durations := []time.Duration{
		time.Second,
		time.Minute,
		time.Hour,
		0,
	}

	for _, duration := range durations {
		result := ctx.SendChecklist(title, intro).After(duration)
		if result == nil {
			t.Errorf("After method should return SendChecklist builder for chaining with duration: %v", duration)
		}

		chainedResult := result.After(time.Second * 30)
		if chainedResult == nil {
			t.Errorf("After method should support chaining and override with duration: %v", duration)
		}
	}
}

func TestSendChecklist_DeleteAfter(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	title := g.String("Test Checklist")
	intro := g.String("Test intro")

	durations := []time.Duration{
		time.Second * 30,
		time.Minute * 5,
		time.Hour,
		0,
	}

	for _, duration := range durations {
		result := ctx.SendChecklist(title, intro).DeleteAfter(duration)
		if result == nil {
			t.Errorf("DeleteAfter method should return SendChecklist builder for chaining with duration: %v", duration)
		}

		chainedResult := result.DeleteAfter(time.Minute * 10)
		if chainedResult == nil {
			t.Errorf("DeleteAfter method should support chaining and override with duration: %v", duration)
		}
	}
}

func TestSendChecklist_OthersCanAddTasks(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	title := g.String("Test Checklist")
	intro := g.String("Test intro")

	result := ctx.SendChecklist(title, intro).OthersCanAddTasks()
	if result == nil {
		t.Error("OthersCanAddTasks method should return SendChecklist builder for chaining")
	}

	chainedResult := result.OthersCanAddTasks()
	if chainedResult == nil {
		t.Error("OthersCanAddTasks method should support chaining")
	}
}

func TestSendChecklist_OthersCanMarkTasksAsDone(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	title := g.String("Test Checklist")
	intro := g.String("Test intro")

	result := ctx.SendChecklist(title, intro).OthersCanMarkTasksAsDone()
	if result == nil {
		t.Error("OthersCanMarkTasksAsDone method should return SendChecklist builder for chaining")
	}

	chainedResult := result.OthersCanMarkTasksAsDone()
	if chainedResult == nil {
		t.Error("OthersCanMarkTasksAsDone method should support chaining")
	}
}

func TestSendChecklist_Markup(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	title := g.String("Test Checklist")
	intro := g.String("Test intro")

	btn1 := keyboard.NewButton().Text(g.String("Test Button")).Callback(g.String("test_data"))
	inlineKeyboard := keyboard.Inline().Button(btn1)

	result := ctx.SendChecklist(title, intro).Markup(inlineKeyboard)
	if result == nil {
		t.Error("Markup method should return SendChecklist builder for chaining with inline keyboard")
	}

	replyKeyboard := keyboard.Reply()
	chainedResult := result.Markup(replyKeyboard)
	if chainedResult == nil {
		t.Error("Markup method should support chaining and override with reply keyboard")
	}
}

func TestSendChecklist_ReplyTo(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	title := g.String("Test Checklist")
	intro := g.String("Test intro")

	messageIDs := []int64{1, 123, 456, 999}

	for _, messageID := range messageIDs {
		result := ctx.SendChecklist(title, intro).ReplyTo(messageID)
		if result == nil {
			t.Errorf("ReplyTo method should return SendChecklist builder for chaining with messageID: %d", messageID)
		}

		chainedResult := result.ReplyTo(messageID + 100)
		if chainedResult == nil {
			t.Errorf("ReplyTo method should support chaining and override with messageID: %d", messageID)
		}
	}
}

func TestSendChecklist_SendWithNoTasks(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	title := g.String("Empty Checklist")
	intro := g.String("Test intro")

	// Test Send with no tasks - should return error
	sendResult := ctx.SendChecklist(title, intro).Send()

	if !sendResult.IsErr() {
		t.Error("Send should return error for checklist with no tasks")
	} else {
		t.Logf("SendChecklist Send with no tasks returned error as expected: %v", sendResult.Err())
	}
}

func TestSendChecklist_SendWithValidTasks(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	title := g.String("Valid Checklist")
	intro := g.String("Test intro")

	// Add a valid task
	sendResult := ctx.SendChecklist(title, intro).
		Task(g.String("Task 1")).Add().
		Send()

	// This will fail with mock bot, but covers the valid tasks path
	if sendResult.IsErr() {
		t.Logf("SendChecklist Send with valid tasks failed as expected: %v", sendResult.Err())
	}
}

func TestSendChecklist_SendWithTooManyTasks(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	title := g.String("Overloaded Checklist")
	intro := g.String("Test intro")

	builder := ctx.SendChecklist(title, intro)

	// Add 101 tasks to exceed the limit of 100
	for i := 0; i < 101; i++ {
		taskText := g.String(fmt.Sprintf("Task %d", i+1))
		builder = builder.Task(taskText).Add()
	}

	sendResult := builder.Send()

	if !sendResult.IsErr() {
		t.Error("Send should return error for too many tasks")
	} else {
		t.Logf("SendChecklist Send with too many tasks returned error as expected: %v", sendResult.Err())
	}
}

func TestSendChecklist_APIURLWithExistingRequestOpts(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	title := g.String("Test Checklist")
	intro := g.String("Test intro")

	// First set Timeout to create RequestOpts, then test APIURL
	result := ctx.SendChecklist(title, intro).
		Timeout(15 * time.Second).                         // This creates RequestOpts
		APIURL(g.String("https://custom.api.example.com")) // This should use existing RequestOpts

	if result == nil {
		t.Error("APIURL with existing RequestOpts should return builder")
	}
}

func TestSendChecklist_TimersIntegration(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	title := g.String("Timed Checklist")
	intro := g.String("Test intro")

	// Test with After and DeleteAfter to cover timer integration
	result := ctx.SendChecklist(title, intro).
		Task(g.String("Timed task")).Add().
		After(time.Second).
		DeleteAfter(time.Minute)

	if result == nil {
		t.Error("Checklist with timers should return builder")
	}

	sendResult := result.Send()

	// This will fail with mock bot, but covers the timer integration path
	if sendResult.IsErr() {
		t.Logf("SendChecklist Send with timers failed as expected: %v", sendResult.Err())
	}
}
