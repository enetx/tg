package ctx

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/entities"
)

type TaskBuilder[T any] struct {
	text     String
	parsmode Option[string]
	entities Option[*entities.Entities]

	target T
	add    func(T, gotgbot.InputChecklistTask)
	next   func(T) int64
}

// HTML sets the task's text to be parsed as HTML.
// Ignored if Entities are set.
func (tb *TaskBuilder[T]) HTML() *TaskBuilder[T] {
	tb.parsmode = Some("HTML")
	return tb
}

// Markdown sets the task's text to be parsed as MarkdownV2.
// Ignored if Entities are set.
func (tb *TaskBuilder[T]) Markdown() *TaskBuilder[T] {
	tb.parsmode = Some("MarkdownV2")
	return tb
}

// Entities sets custom message entities for the task.
// Overrides any ParseMode if set.
func (tb *TaskBuilder[T]) Entities(e *entities.Entities) *TaskBuilder[T] {
	tb.entities = Some(e)
	return tb
}

// Add finalizes and adds the task to the checklist.
// Returns the parent builder, allowing the chain to continue.
func (tb *TaskBuilder[T]) Add() T {
	task := gotgbot.InputChecklistTask{
		Id:   tb.next(tb.target),
		Text: tb.text.Std(),
	}

	if tb.entities.IsSome() {
		task.TextEntities = tb.entities.Some().Std()
	} else if tb.parsmode.IsSome() {
		task.ParseMode = tb.parsmode.Some()
	}

	tb.add(tb.target, task)

	return tb.target
}
