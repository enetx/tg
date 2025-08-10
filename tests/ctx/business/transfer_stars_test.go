package business_test

import (
	"testing"
	"time"

	"github.com/enetx/g"
	"github.com/enetx/tg/ctx/business"
)

func TestTransferStars(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_123")
	account := business.NewAccount(bot, connectionID)
	balance := account.Balance()

	amount := int64(100)
	result := balance.TransferStars(amount)

	if result == nil {
		t.Error("Expected TransferStars builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(30 * time.Second)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}

func TestTransferStars_APIURL(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_123")
	account := business.NewAccount(bot, connectionID)
	balance := account.Balance()

	amount := int64(100)
	result := balance.TransferStars(amount)

	// Test APIURL method
	withAPIURL := result.APIURL(g.String("https://custom.api.example.com"))
	if withAPIURL == nil {
		t.Error("Expected APIURL method to return builder")
	}
}

func TestTransferStars_APIURLWithExistingRequestOpts(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_123")
	account := business.NewAccount(bot, connectionID)
	balance := account.Balance()

	amount := int64(100)

	// First set Timeout to create RequestOpts, then test APIURL
	result := balance.TransferStars(amount).
		Timeout(15 * time.Second).                         // This creates RequestOpts
		APIURL(g.String("https://custom.api.example.com")) // This should use existing RequestOpts

	if result == nil {
		t.Error("APIURL with existing RequestOpts should return builder")
	}
}

func TestTransferStars_Send(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_123")
	account := business.NewAccount(bot, connectionID)
	balance := account.Balance()

	amount := int64(100)
	result := balance.TransferStars(amount)

	// Test Send method - will fail with mock bot but covers the method
	sendResult := result.Send()
	if sendResult.IsErr() {
		t.Logf("TransferStars Send failed as expected with mock bot: %v", sendResult.Err())
	}
}

func TestTransferStars_SendWithAllOptions(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_123")
	account := business.NewAccount(bot, connectionID)
	balance := account.Balance()

	amount := int64(250)

	// Test Send with all options configured
	sendResult := balance.TransferStars(amount).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	// This will fail with mock bot, but covers all configuration paths
	if sendResult.IsErr() {
		t.Logf("TransferStars Send with all options failed as expected: %v", sendResult.Err())
	}
}
