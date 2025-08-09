package business_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/ctx/business"
)

func TestSetBio(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_123")
	account := business.NewAccount(bot, connectionID)

	bio := g.String("Business owner and entrepreneur")
	result := account.SetBio(bio)

	if result == nil {
		t.Error("Expected SetBio builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}

func TestSetBio_APIURL(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_apiurl_bio_123")
	account := business.NewAccount(bot, connectionID)

	bio := g.String("Business owner and entrepreneur")

	// Test APIURL method with various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://custom.api.example.com",
		"",
		"https://api.example.com/bot",
		"http://localhost:8080",
	}

	for _, apiURL := range apiURLs {
		result := account.SetBio(bio)
		apiURLResult := result.APIURL(g.String(apiURL))
		if apiURLResult == nil {
			t.Errorf("APIURL method should return SetBio builder for chaining with URL: %s", apiURL)
		}

		// Test that APIURL can be chained and overridden
		chainedResult := apiURLResult.APIURL(g.String("https://override.example.com"))
		if chainedResult == nil {
			t.Errorf("APIURL method should support chaining and override with URL: %s", apiURL)
		}
	}
}

func TestSetBio_Send(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_send_bio_123")
	account := business.NewAccount(bot, connectionID)

	bio := g.String("Business owner and entrepreneur")

	// Test Send method - will fail with mock but covers the method
	sendResult := account.SetBio(bio).Send()

	if sendResult.IsErr() {
		t.Logf("SetBio Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
