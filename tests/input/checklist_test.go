package input_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/input"
)

func TestNewChecklist(t *testing.T) {
	tasks := g.SliceOf(
		gotgbot.InputChecklistTask{Text: "Task 1"},
		gotgbot.InputChecklistTask{Text: "Task 2"},
	)

	checklist := input.NewChecklist(testTitle, tasks)
	if checklist == nil {
		t.Error("Expected Checklist to be created")
	}
}

func TestChecklist_Build(t *testing.T) {
	tasks := g.SliceOf(
		gotgbot.InputChecklistTask{Text: "Task 1"},
		gotgbot.InputChecklistTask{Text: "Task 2"},
	)

	checklist := input.NewChecklist(testTitle, tasks)
	built := checklist.Build()

	if built.Title != testTitle.Std() {
		t.Errorf("Expected title to be %s, got %s", testTitle.Std(), built.Title)
	}

	if len(built.Tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(built.Tasks))
	}

	if built.Tasks[0].Text != "Task 1" {
		t.Errorf("Expected first task text to be 'Task 1', got %s", built.Tasks[0].Text)
	}

	if built.Tasks[1].Text != "Task 2" {
		t.Errorf("Expected second task text to be 'Task 2', got %s", built.Tasks[1].Text)
	}
}

func TestChecklist_EmptyTasks(t *testing.T) {
	emptyTasks := g.SliceOf[gotgbot.InputChecklistTask]()

	checklist := input.NewChecklist(testTitle, emptyTasks)
	if checklist == nil {
		t.Error("Expected Checklist to be created with empty tasks")
	}

	built := checklist.Build()
	if len(built.Tasks) != 0 {
		t.Errorf("Expected 0 tasks, got %d", len(built.Tasks))
	}
}

func TestChecklist_BuildReturnsCorrectType(t *testing.T) {
	tasks := g.SliceOf(
		gotgbot.InputChecklistTask{Text: "Task 1"},
	)

	checklist := input.NewChecklist(testTitle, tasks)
	built := checklist.Build()

	// Verify that Build() returns the correct type
	if _, ok := any(built).(gotgbot.InputChecklist); !ok {
		t.Error("Expected Build() to return gotgbot.InputChecklist")
	}
}

func TestChecklist_MultipleBuilds(t *testing.T) {
	tasks := g.SliceOf(
		gotgbot.InputChecklistTask{Text: "Task 1"},
	)

	checklist := input.NewChecklist(testTitle, tasks)

	// Build multiple times to ensure consistency
	built1 := checklist.Build()
	built2 := checklist.Build()

	if built1.Title != built2.Title {
		t.Error("Expected multiple builds to return consistent results")
	}

	if len(built1.Tasks) != len(built2.Tasks) {
		t.Error("Expected multiple builds to return consistent task count")
	}
}
