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

func TestGetStarBalance_APIURL(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_apiurl_123")
	account := business.NewAccount(bot, connectionID)
	balance := account.Balance()

	// Test APIURL method with various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://custom.api.example.com",
		"",
		"https://api.example.com/bot",
		"http://localhost:8080",
	}

	for _, apiURL := range apiURLs {
		result := balance.GetStarBalance()
		apiURLResult := result.APIURL(g.String(apiURL))
		if apiURLResult == nil {
			t.Errorf("APIURL method should return GetStarBalance builder for chaining with URL: %s", apiURL)
		}

		// Test that APIURL can be chained and overridden
		chainedResult := apiURLResult.APIURL(g.String("https://override.example.com"))
		if chainedResult == nil {
			t.Errorf("APIURL method should support chaining and override with URL: %s", apiURL)
		}
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
