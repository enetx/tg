package business_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/ctx/business"
)

func TestRemovePhoto(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_123")
	account := business.NewAccount(bot, connectionID)

	result := account.RemovePhoto()

	if result == nil {
		t.Error("Expected RemovePhoto builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}

func TestRemovePhoto_Public(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_public_123")
	account := business.NewAccount(bot, connectionID)

	// Test Public method
	result := account.RemovePhoto()
	publicResult := result.Public()
	if publicResult == nil {
		t.Error("Public method should return RemovePhoto builder for chaining")
	}

	// Test that Public can be chained multiple times
	chainedResult := publicResult.Public()
	if chainedResult == nil {
		t.Error("Public method should support multiple chaining calls")
	}

	// Test Public with other methods
	publicWithOthers := account.RemovePhoto().
		Public().
		Timeout(30)
	if publicWithOthers == nil {
		t.Error("Public method should work with other methods")
	}
}

func TestRemovePhoto_APIURL(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_apiurl_remove_photo_123")
	account := business.NewAccount(bot, connectionID)

	// Test APIURL method with various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://custom.api.example.com",
		"",
		"https://api.example.com/bot",
		"http://localhost:8080",
	}

	for _, apiURL := range apiURLs {
		result := account.RemovePhoto()
		apiURLResult := result.APIURL(g.String(apiURL))
		if apiURLResult == nil {
			t.Errorf("APIURL method should return RemovePhoto builder for chaining with URL: %s", apiURL)
		}

		// Test that APIURL can be chained and overridden
		chainedResult := apiURLResult.APIURL(g.String("https://override.example.com"))
		if chainedResult == nil {
			t.Errorf("APIURL method should support chaining and override with URL: %s", apiURL)
		}
	}
}

func TestRemovePhoto_Send(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_send_remove_photo_123")
	account := business.NewAccount(bot, connectionID)

	// Test Send method - will fail with mock but covers the method
	sendResult := account.RemovePhoto().Send()

	if sendResult.IsErr() {
		t.Logf("RemovePhoto Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with Public
	sendWithPublicResult := account.RemovePhoto().Public().Send()

	if sendWithPublicResult.IsErr() {
		t.Logf("RemovePhoto Send with Public failed as expected with mock bot: %v", sendWithPublicResult.Err())
	}
}
