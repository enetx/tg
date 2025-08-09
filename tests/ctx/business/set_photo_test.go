package business_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/ctx/business"
)

func TestSetPhoto(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_123")
	account := business.NewAccount(bot, connectionID)

	photo := g.String("photo.jpg")
	result := account.SetPhoto(photo)

	if result == nil {
		t.Error("Expected SetPhoto builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}

func TestSetPhoto_Public(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_public_photo_123")
	account := business.NewAccount(bot, connectionID)

	photo := g.String("photo.jpg")

	// Test Public method
	result := account.SetPhoto(photo)
	publicResult := result.Public()
	if publicResult == nil {
		t.Error("Public method should return SetPhoto builder for chaining")
	}

	// Test that Public can be chained multiple times
	chainedResult := publicResult.Public()
	if chainedResult == nil {
		t.Error("Public method should support multiple chaining calls")
	}

	// Test Public with other methods
	publicWithOthers := account.SetPhoto(photo).
		Public().
		Timeout(30)
	if publicWithOthers == nil {
		t.Error("Public method should work with other methods")
	}
}

func TestSetPhoto_APIURL(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_apiurl_photo_123")
	account := business.NewAccount(bot, connectionID)

	photo := g.String("photo.jpg")

	// Test APIURL method with various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://custom.api.example.com",
		"",
		"https://api.example.com/bot",
		"http://localhost:8080",
	}

	for _, apiURL := range apiURLs {
		result := account.SetPhoto(photo)
		apiURLResult := result.APIURL(g.String(apiURL))
		if apiURLResult == nil {
			t.Errorf("APIURL method should return SetPhoto builder for chaining with URL: %s", apiURL)
		}

		// Test that APIURL can be chained and overridden
		chainedResult := apiURLResult.APIURL(g.String("https://override.example.com"))
		if chainedResult == nil {
			t.Errorf("APIURL method should support chaining and override with URL: %s", apiURL)
		}
	}
}

func TestSetPhoto_Send(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_send_photo_123")
	account := business.NewAccount(bot, connectionID)

	photo := g.String("photo.jpg")

	// Test Send method - will fail with mock but covers the method
	sendResult := account.SetPhoto(photo).Send()

	if sendResult.IsErr() {
		t.Logf("SetPhoto Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with Public
	sendWithPublicResult := account.SetPhoto(photo).
		Public().
		Send()

	if sendWithPublicResult.IsErr() {
		t.Logf("SetPhoto Send with Public failed as expected with mock bot: %v", sendWithPublicResult.Err())
	}
}
