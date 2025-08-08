package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/entities"
)

func createSendChecklist(title string) *ctx.SendChecklist {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	context := ctx.New(bot, rawCtx)
	return context.SendChecklist(g.String("business_conn"), g.String(title))
}

func TestTaskBuilder_HTML(t *testing.T) {
	sc := createSendChecklist("Test Checklist")

	// Test HTML method on TaskBuilder
	task := sc.Task(g.String("Test task")).HTML()

	// Add the task and check if SendChecklist was returned
	returned := task.Add()

	if returned != sc {
		t.Error("Expected Add to return the SendChecklist instance")
	}
}

func TestTaskBuilder_Markdown(t *testing.T) {
	sc := createSendChecklist("Test Checklist")

	// Test Markdown method on TaskBuilder
	task := sc.Task(g.String("Test task")).Markdown()

	// Add the task and check if SendChecklist was returned
	returned := task.Add()

	if returned != sc {
		t.Error("Expected Add to return the SendChecklist instance")
	}
}

func TestTaskBuilder_Entities(t *testing.T) {
	sc := createSendChecklist("Test Checklist")

	// Create entities
	entity := entities.New("Test task").Bold("Test")

	// Test Entities method on TaskBuilder
	task := sc.Task(g.String("Test task")).Entities(entity)

	// Add the task and check if SendChecklist was returned
	returned := task.Add()

	if returned != sc {
		t.Error("Expected Add to return the SendChecklist instance")
	}
}

func TestTaskBuilder_Add(t *testing.T) {
	sc := createSendChecklist("Test Checklist")

	// Add a task
	returned := sc.Task(g.String("Test task content")).Add()

	// Should return the SendChecklist
	if returned != sc {
		t.Error("Expected Add to return the SendChecklist instance")
	}
}

func TestTaskBuilder_ChainedMethods(t *testing.T) {
	sc := createSendChecklist("Test Checklist")

	// Test chaining methods
	returned := sc.Task(g.String("Test task")).HTML().Markdown().Add()

	// Should return the SendChecklist
	if returned != sc {
		t.Error("Expected chained methods to return the SendChecklist instance")
	}
}

func TestTaskBuilder_MultipleTasksWithDifferentFormatting(t *testing.T) {
	sc := createSendChecklist("Test Checklist")

	// Add multiple tasks with different formatting
	entity := entities.New("Entity Task").Bold("Entity")

	sc.Task(g.String("HTML Task")).HTML().Add()
	sc.Task(g.String("Markdown Task")).Markdown().Add()
	sc.Task(g.String("Entity Task")).Entities(entity).Add()
	sc.Task(g.String("Plain Task")).Add()
}

func TestTaskBuilder_EntitiesMethodChaining(t *testing.T) {
	sc := createSendChecklist("Test Checklist")

	// Create a proper entities object
	entity := entities.New("Test task with entities").Bold("Test")

	// Test method chaining with entities
	returned := sc.Task(g.String("Test task")).HTML().Entities(entity).Add()

	if returned != sc {
		t.Error("Expected Add to return the SendChecklist instance")
	}
}

func TestTaskBuilder_EmptyText(t *testing.T) {
	sc := createSendChecklist("Test Checklist")

	// Test with empty text
	returned := sc.Task(g.String("")).Add()

	if returned != sc {
		t.Error("Expected Add to return the SendChecklist instance with empty text")
	}
}
