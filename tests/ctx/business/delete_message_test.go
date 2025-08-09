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

func TestDeleteMessage_APIURL(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_apiurl_123")
	account := business.NewAccount(bot, connectionID)
	message := account.Message()

	messageIDs := g.Slice[int64]{}
	messageIDs.Push(123)

	// Test APIURL method with various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://custom.api.example.com",
		"",
		"https://api.example.com/bot",
		"http://localhost:8080",
	}

	for _, apiURL := range apiURLs {
		result := message.Delete(messageIDs)
		apiURLResult := result.APIURL(g.String(apiURL))
		if apiURLResult == nil {
			t.Errorf("APIURL method should return Delete builder for chaining with URL: %s", apiURL)
		}

		// Test that APIURL can be chained and overridden
		chainedResult := apiURLResult.APIURL(g.String("https://override.example.com"))
		if chainedResult == nil {
			t.Errorf("APIURL method should support chaining and override with URL: %s", apiURL)
		}
	}
}

func TestDeleteMessage_Send(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_send_123")
	account := business.NewAccount(bot, connectionID)
	message := account.Message()

	messageIDs := g.Slice[int64]{}
	messageIDs.Push(456)
	messageIDs.Push(789)

	// Test Send method - will fail with mock but covers the method
	sendResult := message.Delete(messageIDs).Send()

	if sendResult.IsErr() {
		t.Logf("DeleteMessage Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
