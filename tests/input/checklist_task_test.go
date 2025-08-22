package input_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/input"
)

func TestNewChecklistTask(t *testing.T) {
	taskText := g.String("Complete documentation")
	task := input.NewChecklistTask(1, taskText)
	if task == nil {
		t.Error("Expected ChecklistTask to be created")
	}
}

func TestChecklistTask_Build(t *testing.T) {
	taskText := g.String("Complete documentation")
	task := input.NewChecklistTask(1, taskText)
	built := task.Build()

	if built.Id != 1 {
		t.Errorf("Expected task ID to be 1, got %d", built.Id)
	}
	if built.Text != taskText.Std() {
		t.Errorf("Expected task text to be %s, got %s", taskText.Std(), built.Text)
	}
}

func TestChecklistTask_HTML(t *testing.T) {
	taskText := g.String("Complete documentation")
	task := input.NewChecklistTask(1, taskText)
	result := task.HTML()

	if result == nil {
		t.Error("Expected HTML method to return ChecklistTask")
	}
	if result != task {
		t.Error("Expected HTML to return same ChecklistTask instance")
	}

	built := result.Build()
	if built.ParseMode != "HTML" {
		t.Error("Expected ParseMode to be set to HTML")
	}
}

func TestChecklistTask_Markdown(t *testing.T) {
	taskText := g.String("Complete documentation")
	task := input.NewChecklistTask(1, taskText)
	result := task.Markdown()

	if result == nil {
		t.Error("Expected Markdown method to return ChecklistTask")
	}
	if result != task {
		t.Error("Expected Markdown to return same ChecklistTask instance")
	}

	built := result.Build()
	if built.ParseMode != "MarkdownV2" {
		t.Error("Expected ParseMode to be set to MarkdownV2")
	}
}

func TestChecklistTask_Entities(t *testing.T) {
	taskText := g.String("Complete documentation")
	entities := entities.New(taskText).Bold(g.String("Complete"))
	task := input.NewChecklistTask(1, taskText)
	result := task.Entities(entities)

	if result == nil {
		t.Error("Expected Entities method to return ChecklistTask")
	}
	if result != task {
		t.Error("Expected Entities to return same ChecklistTask instance")
	}

	built := result.Build()
	if len(built.TextEntities) == 0 {
		t.Error("Expected TextEntities to be set")
	}
}

func TestChecklistTask_BuildReturnsCorrectType(t *testing.T) {
	taskText := g.String("Complete documentation")
	task := input.NewChecklistTask(1, taskText)
	built := task.Build()

	// Verify that Build() returns the correct type
	if _, ok := any(built).(gotgbot.InputChecklistTask); !ok {
		t.Error("Expected Build() to return gotgbot.InputChecklistTask")
	}
}

func TestChecklistTask_MethodChaining(t *testing.T) {
	taskText := g.String("Complete documentation")
	entities := entities.New(taskText).Bold(g.String("Complete"))
	result := input.NewChecklistTask(1, taskText).
		HTML().
		Entities(entities)

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built.Text != taskText.Std() {
		t.Errorf("Expected chained task text to be %s, got %s", taskText.Std(), built.Text)
	}
	if built.ParseMode != "HTML" {
		t.Error("Expected ParseMode to be HTML")
	}
	if len(built.TextEntities) == 0 {
		t.Error("Expected TextEntities to be set")
	}
}
