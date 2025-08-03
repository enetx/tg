package input_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/input"
)

func TestNewChecklistTask(t *testing.T) {
	taskText := g.String("Complete documentation")
	task := input.NewChecklistTask(taskText)
	if task == nil {
		t.Error("Expected ChecklistTask to be created")
	}
}

func TestChecklistTask_Build(t *testing.T) {
	taskText := g.String("Complete documentation")
	task := input.NewChecklistTask(taskText)
	built := task.Build()

	if built.Text != taskText.Std() {
		t.Errorf("Expected task text to be %s, got %s", taskText.Std(), built.Text)
	}
}

func TestChecklistTask_Checked(t *testing.T) {
	taskText := g.String("Complete documentation")
	task := input.NewChecklistTask(taskText)
	result := task.Checked()

	if result == nil {
		t.Error("Expected Checked method to return ChecklistTask")
	}
	if result != task {
		t.Error("Expected Checked to return same ChecklistTask instance")
	}
}

func TestChecklistTask_BuildReturnsCorrectType(t *testing.T) {
	taskText := g.String("Complete documentation")
	task := input.NewChecklistTask(taskText)
	built := task.Build()

	// Verify that Build() returns the correct type
	if _, ok := any(built).(gotgbot.InputChecklistTask); !ok {
		t.Error("Expected Build() to return gotgbot.InputChecklistTask")
	}
}

func TestChecklistTask_MethodChaining(t *testing.T) {
	taskText := g.String("Complete documentation")
	result := input.NewChecklistTask(taskText).
		Checked()

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built.Text != taskText.Std() {
		t.Errorf("Expected chained task text to be %s, got %s", taskText.Std(), built.Text)
	}
}

func TestChecklistTask_EmptyText(t *testing.T) {
	emptyText := g.String("")
	task := input.NewChecklistTask(emptyText)
	if task == nil {
		t.Error("Expected ChecklistTask to be created with empty text")
	}

	built := task.Build()
	if built.Text != "" {
		t.Errorf("Expected empty text, got %s", built.Text)
	}
}

func TestChecklistTask_MultipleBuilds(t *testing.T) {
	taskText := g.String("Complete documentation")
	task := input.NewChecklistTask(taskText)

	// Build multiple times to ensure consistency
	built1 := task.Build()
	built2 := task.Build()

	if built1.Text != built2.Text {
		t.Error("Expected multiple builds to return consistent results")
	}
}
