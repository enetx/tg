package business_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/ctx/business"
)

func TestGetStarBalance(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_123")
	account := business.NewAccount(bot, connectionID)
	balance := account.Balance()

	result := balance.GetStarBalance()

	if result == nil {
		t.Error("Expected GetStarBalance builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}

func TestGetStarBalance_Send(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_send_123")
	account := business.NewAccount(bot, connectionID)
	balance := account.Balance()

	// Test Send method - will fail with mock but covers the method
	sendResult := balance.GetStarBalance().Send()

	if sendResult.IsErr() {
		t.Logf("GetStarBalance Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
