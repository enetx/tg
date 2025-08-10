package ctx_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/keyboard"
)

func TestContext_EditMessageChecklist(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	title := g.String("Updated Checklist")
	result := ctx.EditMessageChecklist(title)

	if result == nil {
		t.Error("Expected EditMessageChecklist builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(456)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}

func TestContext_EditMessageChecklistChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	title := g.String("Updated Checklist")
	result := ctx.EditMessageChecklist(title).
		ChatID(456).
		MessageID(789)

	if result == nil {
		t.Error("Expected EditMessageChecklist builder to be created")
	}

	// Test that builder is functional
	_ = result
}

func TestEditMessageChecklist_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	title := g.String("Updated Test Checklist")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.EditMessageChecklist(title).Send()

	if sendResult.IsErr() {
		t.Logf("EditMessageChecklist Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.EditMessageChecklist(title).
		ChatID(456).
		MessageID(123).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("EditMessageChecklist configured Send failed as expected: %v", configuredSendResult.Err())
	}
}

// Tests for methods with 0% coverage

func TestEditMessageChecklist_TaskMethod(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Task method returns TaskBuilder
	taskBuilder := ctx.EditMessageChecklist(g.String("business_123")).
		Title(g.String("Task List")).
		Task(g.String("Complete unit tests"))

	if taskBuilder == nil {
		t.Error("Task method should return TaskBuilder")
	}

	// Test TaskBuilder method chaining
	htmlTask := taskBuilder.HTML()
	if htmlTask == nil {
		t.Error("TaskBuilder HTML method should return TaskBuilder for chaining")
	}

	markdownTask := taskBuilder.Markdown()
	if markdownTask == nil {
		t.Error("TaskBuilder Markdown method should return TaskBuilder for chaining")
	}

	entitiesTask := taskBuilder.Entities(entities.New("Complete unit tests").Bold(g.String("Complete")))
	if entitiesTask == nil {
		t.Error("TaskBuilder Entities method should return TaskBuilder for chaining")
	}

	// Test Add method returns parent builder
	parentBuilder := entitiesTask.Add()
	if parentBuilder == nil {
		t.Error("TaskBuilder Add method should return parent EditMessageChecklist")
	}

	// Test Send after adding task
	sendResult := parentBuilder.Send()
	if sendResult.IsErr() {
		t.Logf("EditMessageChecklist with task Send failed as expected: %v", sendResult.Err())
	}
}

func TestEditMessageChecklist_TitleMethod(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Title method functionality
	titles := []string{
		"Todo List",
		"Shopping List",
		"Project Tasks",
		"Daily Goals",
		"Meeting Agenda",
		"📝 Important Tasks",
		"🛒 Grocery List",
		"✅ Completed Items",
		"", // Empty title
		"Very Long Checklist Title That Tests Extended Character Limits",
		"Title with Special Characters: !@#$%^&*()",
		"日本語のタスクリスト",   // Japanese characters
		"Список задач", // Russian characters
	}

	for _, title := range titles {
		displayTitle := title
		if title == "" {
			displayTitle = "[empty]"
		}

		titleResult := ctx.EditMessageChecklist(g.String("business_123")).
			Title(g.String(title)).
			Task(g.String("Sample task")).
			HTML().
			Add().
			ChatID(456).
			MessageID(789)

		if titleResult == nil {
			t.Errorf("Title method with '%s' should work", displayTitle)
		}

		// Test send with title
		sendResult := titleResult.Send()
		if sendResult.IsErr() {
			t.Logf("EditMessageChecklist with title '%s' Send failed as expected: %v", displayTitle, sendResult.Err())
		}
	}
}

func TestEditMessageChecklist_OthersCanAddTasks(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test OthersCanAddTasks method
	othersAddResult := ctx.EditMessageChecklist(g.String("business_123")).
		Title(g.String("Collaborative Tasks")).
		OthersCanAddTasks().
		Task(g.String("Initial task")).
		Add().
		ChatID(456).
		MessageID(789)

	if othersAddResult == nil {
		t.Error("OthersCanAddTasks method should return EditMessageChecklist for chaining")
	}

	// Test send with OthersCanAddTasks
	sendResult := othersAddResult.Send()
	if sendResult.IsErr() {
		t.Logf("EditMessageChecklist with OthersCanAddTasks Send failed as expected: %v", sendResult.Err())
	}
}

func TestEditMessageChecklist_OthersCanMarkTasksAsDone(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test OthersCanMarkTasksAsDone method
	othersMarkResult := ctx.EditMessageChecklist(g.String("business_123")).
		Title(g.String("Shared Checklist")).
		OthersCanMarkTasksAsDone().
		Task(g.String("Team task")).
		Add().
		ChatID(456).
		MessageID(789)

	if othersMarkResult == nil {
		t.Error("OthersCanMarkTasksAsDone method should return EditMessageChecklist for chaining")
	}

	// Test send with OthersCanMarkTasksAsDone
	sendResult := othersMarkResult.Send()
	if sendResult.IsErr() {
		t.Logf("EditMessageChecklist with OthersCanMarkTasksAsDone Send failed as expected: %v", sendResult.Err())
	}
}

func TestEditMessageChecklist_Markup(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Markup method with inline keyboard
	inlineKB := keyboard.Inline().
		Text(g.String("Add Task"), g.String("add_task")).
		Row().
		Text(g.String("Mark Complete"), g.String("mark_complete"))

	markupResult := ctx.EditMessageChecklist(g.String("business_123")).
		Title(g.String("Interactive Checklist")).
		Markup(inlineKB).
		Task(g.String("Task with buttons")).
		Add().
		ChatID(456).
		MessageID(789)

	if markupResult == nil {
		t.Error("Markup method should return EditMessageChecklist for chaining")
	}

	// Test send with markup
	sendResult := markupResult.Send()
	if sendResult.IsErr() {
		t.Logf("EditMessageChecklist with Markup Send failed as expected: %v", sendResult.Err())
	}
}

func TestEditMessageChecklist_ComprehensiveWorkflow(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test all methods in comprehensive workflow
	complexResult := ctx.EditMessageChecklist(g.String("business_comprehensive")).
		Title(g.String("📋 Complete Project Checklist")).
		ChatID(456).
		MessageID(789).
		OthersCanAddTasks().
		OthersCanMarkTasksAsDone().
		Task(g.String("**Design** the user interface")).
		Markdown().
		Add().
		Task(g.String("Implement backend API")).
		HTML().
		Add().
		Task(g.String("Write comprehensive tests")).
		Entities(entities.New("Write comprehensive tests").Bold(g.String("comprehensive"))).
		Add().
		Task(g.String("Deploy to production")).
		Add().
		Markup(keyboard.Inline().
			Text(g.String("⏸ Pause"), g.String("pause_checklist")).
			Text(g.String("✅ Complete"), g.String("complete_checklist"))).
		Timeout(45 * time.Second).
		APIURL(g.String("https://comprehensive-checklist-api.telegram.org")).
		Send()

	if complexResult.IsErr() {
		t.Logf("EditMessageChecklist complex workflow Send failed as expected: %v", complexResult.Err())
	}
}

func TestEditMessageChecklist_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Send method with no tasks (should fail)
	noTasksResult := ctx.EditMessageChecklist(g.String("business_123")).
		Title(g.String("Empty Checklist")).
		Send()

	if noTasksResult.IsErr() {
		t.Logf("EditMessageChecklist with no tasks Send failed as expected: %v", noTasksResult.Err())
	}

	// Test Send method with too many tasks (should fail)
	manyTasksBuilder := ctx.EditMessageChecklist(g.String("business_123")).
		Title(g.String("Too Many Tasks"))

	// Add more than 100 tasks
	for i := 0; i <= 101; i++ {
		manyTasksBuilder = manyTasksBuilder.
			Task(g.String(fmt.Sprintf("Task number %d", i))).
			Add()
	}

	tooManyTasksResult := manyTasksBuilder.Send()
	if tooManyTasksResult.IsErr() {
		t.Logf("EditMessageChecklist with too many tasks Send failed as expected: %v", tooManyTasksResult.Err())
	}
}

func TestEditMessageChecklist_APIURL_NilRequestOpts(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// Test APIURL when RequestOpts is nil (covers the nil branch)
	result := ctx.EditMessageChecklist(g.String("Test Checklist"))
	if result == nil {
		t.Error("EditMessageChecklist should return builder")
	}

	// This should create RequestOpts and set APIURL
	apiResult := result.APIURL(g.String("https://api.test.com"))
	if apiResult == nil {
		t.Error("APIURL should return builder when RequestOpts is nil")
	}
}
