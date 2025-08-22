package reply

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/reply"
)

func TestNew(t *testing.T) {
	messageID := int64(123)
	params := reply.New(messageID)

	if params == nil {
		t.Error("New should return non-nil Parameters")
	}

	result := params.Std()
	if result == nil {
		t.Error("Std should return non-nil ReplyParameters")
	}

	if result.MessageId != messageID {
		t.Errorf("Expected MessageId %d, got %d", messageID, result.MessageId)
	}
}

func TestParameters_ChecklistTask(t *testing.T) {
	messageID := int64(123)
	taskID := int64(456)

	params := reply.New(messageID).ChecklistTask(taskID)

	if params == nil {
		t.Error("ChecklistTask should return non-nil Parameters for chaining")
	}

	result := params.Std()
	if result.ChecklistTaskId != taskID {
		t.Errorf("Expected ChecklistTaskId %d, got %d", taskID, result.ChecklistTaskId)
	}

	if result.MessageId != messageID {
		t.Errorf("Expected MessageId %d, got %d", messageID, result.MessageId)
	}
}

func TestParameters_ChecklistTaskChaining(t *testing.T) {
	messageID := int64(789)
	taskID := int64(999)

	// Test method chaining
	params := reply.New(messageID).ChecklistTask(taskID)

	// Should return same instance for chaining
	if params == nil {
		t.Error("ChecklistTask should support method chaining")
	}

	result := params.Std()
	if result.MessageId != messageID {
		t.Errorf("Chaining failed: expected MessageId %d, got %d", messageID, result.MessageId)
	}

	if result.ChecklistTaskId != taskID {
		t.Errorf("Chaining failed: expected ChecklistTaskId %d, got %d", taskID, result.ChecklistTaskId)
	}
}

func TestParameters_EdgeCases(t *testing.T) {
	// Test with zero values
	params := reply.New(0)
	if params == nil {
		t.Error("New should handle zero messageID")
	}

	result := params.Std()
	if result.MessageId != 0 {
		t.Error("Should preserve zero MessageId")
	}

	// Test with negative values
	params = reply.New(-1)
	if params == nil {
		t.Error("New should handle negative messageID")
	}

	result = params.Std()
	if result.MessageId != -1 {
		t.Error("Should preserve negative MessageId")
	}

	// Test ChecklistTask with zero
	params = reply.New(123).ChecklistTask(0)
	result = params.Std()
	if result.ChecklistTaskId != 0 {
		t.Error("Should preserve zero ChecklistTaskId")
	}
}

func TestParameters_MultipleChecklistTasks(t *testing.T) {
	messageID := int64(123)
	firstTaskID := int64(456)
	secondTaskID := int64(789)

	// Test that later calls override previous ones
	params := reply.New(messageID).ChecklistTask(firstTaskID).ChecklistTask(secondTaskID)

	result := params.Std()
	if result.ChecklistTaskId != secondTaskID {
		t.Errorf("Expected final ChecklistTaskId %d, got %d", secondTaskID, result.ChecklistTaskId)
	}
}

func TestParameters_Std_Idempotent(t *testing.T) {
	messageID := int64(123)
	taskID := int64(456)

	params := reply.New(messageID).ChecklistTask(taskID)

	// Call Std() multiple times
	result1 := params.Std()
	result2 := params.Std()

	if result1.MessageId != result2.MessageId {
		t.Error("Std() should be idempotent for MessageId")
	}

	if result1.ChecklistTaskId != result2.ChecklistTaskId {
		t.Error("Std() should be idempotent for ChecklistTaskId")
	}
}

func TestParameters_AllMethods(t *testing.T) {
	messageID := int64(123)
	params := reply.New(messageID)

	// Test ChatID method
	result := params.ChatID(456)
	if result == nil {
		t.Error("ChatID should return Parameters for chaining")
	}
	if result.Std().ChatId != 456 {
		t.Error("ChatID not set correctly")
	}

	// Test AllowSendingWithoutReply method
	result = params.AllowSendingWithoutReply()
	if result == nil {
		t.Error("AllowSendingWithoutReply should return Parameters for chaining")
	}
	if !result.Std().AllowSendingWithoutReply {
		t.Error("AllowSendingWithoutReply not set correctly")
	}

	// Test Quote method
	quote := "Test quote"
	result = params.Quote(g.String(quote))
	if result == nil {
		t.Error("Quote should return Parameters for chaining")
	}
	if result.Std().Quote != quote {
		t.Error("Quote not set correctly")
	}

	// Test QuoteHTML method
	result = params.QuoteHTML()
	if result == nil {
		t.Error("QuoteHTML should return Parameters for chaining")
	}
	if result.Std().QuoteParseMode != "HTML" {
		t.Error("QuoteParseMode not set to HTML")
	}

	// Test QuoteMarkdown method
	result = params.QuoteMarkdown()
	if result == nil {
		t.Error("QuoteMarkdown should return Parameters for chaining")
	}
	if result.Std().QuoteParseMode != "MarkdownV2" {
		t.Error("QuoteParseMode not set to MarkdownV2")
	}

	// Test QuotePosition method
	position := int64(10)
	result = params.QuotePosition(position)
	if result == nil {
		t.Error("QuotePosition should return Parameters for chaining")
	}
	if result.Std().QuotePosition != position {
		t.Error("QuotePosition not set correctly")
	}

	// Test Build method
	buildResult := params.Build()
	if buildResult == nil {
		t.Error("Build should return ReplyParameters")
	}
	if buildResult.MessageId != messageID {
		t.Error("Build should preserve MessageId")
	}
}
