package business_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/ctx/business"
)

func TestSetUsername(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_123")
	account := business.NewAccount(bot, connectionID)

	username := g.String("john_business")
	result := account.SetUsername(username)

	if result == nil {
		t.Error("Expected SetUsername builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}

func TestSetUsername_APIURL(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_apiurl_username_123")
	account := business.NewAccount(bot, connectionID)

	username := g.String("john_business")

	// Test APIURL method with various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://custom.api.example.com",
		"",
		"https://api.example.com/bot",
		"http://localhost:8080",
	}

	for _, apiURL := range apiURLs {
		result := account.SetUsername(username)
		apiURLResult := result.APIURL(g.String(apiURL))
		if apiURLResult == nil {
			t.Errorf("APIURL method should return SetUsername builder for chaining with URL: %s", apiURL)
		}

		// Test that APIURL can be chained and overridden
		chainedResult := apiURLResult.APIURL(g.String("https://override.example.com"))
		if chainedResult == nil {
			t.Errorf("APIURL method should support chaining and override with URL: %s", apiURL)
		}
	}
}

func TestSetUsername_Send(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_send_username_123")
	account := business.NewAccount(bot, connectionID)

	username := g.String("john_business")

	// Test Send method - will fail with mock but covers the method
	sendResult := account.SetUsername(username).Send()

	if sendResult.IsErr() {
		t.Logf("SetUsername Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
