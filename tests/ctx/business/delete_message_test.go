package business_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/ctx/business"
)

func TestDeleteMessage(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_123")
	account := business.NewAccount(bot, connectionID)
	message := account.Message()

	messageIDs := g.Slice[int64]{}
	messageIDs.Push(123)
	messageIDs.Push(124)
	result := message.Delete(messageIDs)

	if result == nil {
		t.Error("Expected DeleteMessage builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}
