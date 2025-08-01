package filters_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/tg/handlers/filters"
)

func TestBusinessMessagesDeleted_TypeDefinition(t *testing.T) {
	// Test that BusinessMessagesDeleted is a function type that can be defined
	var filter BusinessMessagesDeleted

	// Define a sample filter function
	filter = func(deletedMessages *gotgbot.BusinessMessagesDeleted) bool {
		return deletedMessages != nil
	}

	if filter == nil {
		t.Error("Expected filter to be defined")
	}
}

func TestBusinessMessagesDeleted_FilterFunction(t *testing.T) {
	// Test actual filter functionality

	// Create a filter that accepts all non-nil messages
	acceptAllFilter := BusinessMessagesDeleted(func(deletedMessages *gotgbot.BusinessMessagesDeleted) bool {
		return deletedMessages != nil
	})

	// Test with nil message
	if acceptAllFilter(nil) {
		t.Error("Expected filter to reject nil messages")
	}

	// Test with valid message
	validMessage := &gotgbot.BusinessMessagesDeleted{
		BusinessConnectionId: "test_connection",
		Chat:                 gotgbot.Chat{Id: 123},
		MessageIds:           []int64{1, 2, 3},
	}

	if !acceptAllFilter(validMessage) {
		t.Error("Expected filter to accept valid messages")
	}
}

func TestBusinessMessagesDeleted_ConditionalFilter(t *testing.T) {
	// Test a more complex filter based on message content

	// Filter that only accepts messages from specific chat
	specificChatFilter := BusinessMessagesDeleted(func(deletedMessages *gotgbot.BusinessMessagesDeleted) bool {
		return deletedMessages != nil && deletedMessages.Chat.Id == 123
	})

	// Test with matching chat
	matchingMessage := &gotgbot.BusinessMessagesDeleted{
		Chat: gotgbot.Chat{Id: 123},
	}

	if !specificChatFilter(matchingMessage) {
		t.Error("Expected filter to accept message from specific chat")
	}

	// Test with non-matching chat
	nonMatchingMessage := &gotgbot.BusinessMessagesDeleted{
		Chat: gotgbot.Chat{Id: 456},
	}

	if specificChatFilter(nonMatchingMessage) {
		t.Error("Expected filter to reject message from different chat")
	}
}

func TestBusinessMessagesDeleted_MessageCountFilter(t *testing.T) {
	// Filter that only accepts deletions with multiple messages
	multiMessageFilter := BusinessMessagesDeleted(func(deletedMessages *gotgbot.BusinessMessagesDeleted) bool {
		return deletedMessages != nil && len(deletedMessages.MessageIds) > 1
	})

	// Test with single message
	singleMessage := &gotgbot.BusinessMessagesDeleted{
		MessageIds: []int64{1},
	}

	if multiMessageFilter(singleMessage) {
		t.Error("Expected filter to reject single message deletion")
	}

	// Test with multiple messages
	multipleMessages := &gotgbot.BusinessMessagesDeleted{
		MessageIds: []int64{1, 2, 3},
	}

	if !multiMessageFilter(multipleMessages) {
		t.Error("Expected filter to accept multiple message deletion")
	}

	// Test with empty message list
	emptyMessages := &gotgbot.BusinessMessagesDeleted{
		MessageIds: []int64{},
	}

	if multiMessageFilter(emptyMessages) {
		t.Error("Expected filter to reject empty message deletion")
	}
}

func TestBusinessMessagesDeleted_BusinessConnectionFilter(t *testing.T) {
	// Filter based on business connection ID
	connectionFilter := BusinessMessagesDeleted(func(deletedMessages *gotgbot.BusinessMessagesDeleted) bool {
		return deletedMessages != nil && deletedMessages.BusinessConnectionId == "important_connection"
	})

	// Test with matching connection
	matchingConnection := &gotgbot.BusinessMessagesDeleted{
		BusinessConnectionId: "important_connection",
	}

	if !connectionFilter(matchingConnection) {
		t.Error("Expected filter to accept matching business connection")
	}

	// Test with different connection
	differentConnection := &gotgbot.BusinessMessagesDeleted{
		BusinessConnectionId: "other_connection",
	}

	if connectionFilter(differentConnection) {
		t.Error("Expected filter to reject different business connection")
	}
}

func TestBusinessMessagesDeleted_CombinedFilters(t *testing.T) {
	// Test combining multiple filter conditions
	combinedFilter := BusinessMessagesDeleted(func(deletedMessages *gotgbot.BusinessMessagesDeleted) bool {
		return deletedMessages != nil &&
			deletedMessages.Chat.Id == 123 &&
			len(deletedMessages.MessageIds) >= 2 &&
			deletedMessages.BusinessConnectionId != ""
	})

	// Test message that meets all criteria
	validMessage := &gotgbot.BusinessMessagesDeleted{
		BusinessConnectionId: "test_connection",
		Chat:                 gotgbot.Chat{Id: 123},
		MessageIds:           []int64{1, 2},
	}

	if !combinedFilter(validMessage) {
		t.Error("Expected combined filter to accept message meeting all criteria")
	}

	// Test message that fails one criterion (wrong chat)
	wrongChatMessage := &gotgbot.BusinessMessagesDeleted{
		BusinessConnectionId: "test_connection",
		Chat:                 gotgbot.Chat{Id: 456},
		MessageIds:           []int64{1, 2},
	}

	if combinedFilter(wrongChatMessage) {
		t.Error("Expected combined filter to reject message with wrong chat")
	}

	// Test message that fails one criterion (too few messages)
	fewMessagesMessage := &gotgbot.BusinessMessagesDeleted{
		BusinessConnectionId: "test_connection",
		Chat:                 gotgbot.Chat{Id: 123},
		MessageIds:           []int64{1},
	}

	if combinedFilter(fewMessagesMessage) {
		t.Error("Expected combined filter to reject message with too few messages")
	}
}
