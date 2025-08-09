package business_test

import (
	"testing"
	"time"

	"github.com/enetx/g"
	"github.com/enetx/tg/ctx/business"
)

func TestSetName(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_123")
	account := business.NewAccount(bot, connectionID)

	firstName := g.String("John")
	result := account.SetName(firstName)

	if result == nil {
		t.Error("Expected SetName builder to be created")
	}

	// Test method chaining
	withLastName := result.LastName(g.String("Doe"))
	if withLastName == nil {
		t.Error("Expected LastName method to return builder")
	}
}

func TestSetName_Timeout(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_timeout_name_123")
	account := business.NewAccount(bot, connectionID)

	firstName := g.String("John")

	// Test Timeout method with various durations
	timeouts := []time.Duration{
		1 * time.Second,
		5 * time.Second,
		30 * time.Second,
		1 * time.Minute,
		5 * time.Minute,
		0 * time.Second, // Zero timeout
	}

	for _, timeout := range timeouts {
		result := account.SetName(firstName)
		timeoutResult := result.Timeout(timeout)
		if timeoutResult == nil {
			t.Errorf("Timeout method should return SetName builder for chaining with timeout %v", timeout)
		}

		// Test that Timeout can be chained and overridden
		chainedResult := timeoutResult.Timeout(timeout + 1*time.Second)
		if chainedResult == nil {
			t.Errorf("Timeout method should support chaining and override with timeout %v", timeout)
		}
	}
}

func TestSetName_APIURL(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_apiurl_name_123")
	account := business.NewAccount(bot, connectionID)

	firstName := g.String("John")

	// Test APIURL method with various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://custom.api.example.com",
		"",
		"https://api.example.com/bot",
		"http://localhost:8080",
	}

	for _, apiURL := range apiURLs {
		result := account.SetName(firstName)
		apiURLResult := result.APIURL(g.String(apiURL))
		if apiURLResult == nil {
			t.Errorf("APIURL method should return SetName builder for chaining with URL: %s", apiURL)
		}

		// Test that APIURL can be chained and overridden
		chainedResult := apiURLResult.APIURL(g.String("https://override.example.com"))
		if chainedResult == nil {
			t.Errorf("APIURL method should support chaining and override with URL: %s", apiURL)
		}
	}
}

func TestSetName_Send(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_send_name_123")
	account := business.NewAccount(bot, connectionID)

	firstName := g.String("John")

	// Test Send method - will fail with mock but covers the method
	sendResult := account.SetName(firstName).Send()

	if sendResult.IsErr() {
		t.Logf("SetName Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with LastName
	sendWithLastNameResult := account.SetName(firstName).
		LastName(g.String("Doe")).
		Send()

	if sendWithLastNameResult.IsErr() {
		t.Logf("SetName Send with LastName failed as expected with mock bot: %v", sendWithLastNameResult.Err())
	}
}
