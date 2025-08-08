package business_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/ctx/business"
)

func TestGetConnection(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_123")
	account := business.NewAccount(bot, connectionID)

	result := account.GetConnection()

	if result == nil {
		t.Error("Expected GetConnection builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}

func TestGetConnection_Send(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_send_123")
	account := business.NewAccount(bot, connectionID)

	// Test Send method - will fail with mock but covers the method
	sendResult := account.GetConnection().Send()

	if sendResult.IsErr() {
		t.Logf("GetConnection Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
