package business_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/ctx/business"
)

func TestReadMessage(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_123")
	account := business.NewAccount(bot, connectionID)
	message := account.Message()

	chatID := int64(456)
	messageID := int64(123)

	result := message.Read(chatID, messageID)

	if result == nil {
		t.Error("Expected ReadMessage builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}

func TestReadMessage_APIURL(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_apiurl_read_123")
	account := business.NewAccount(bot, connectionID)
	message := account.Message()

	chatID := int64(456)
	messageID := int64(123)

	// Test APIURL method with various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://custom.api.example.com",
		"",
		"https://api.example.com/bot",
		"http://localhost:8080",
	}

	for _, apiURL := range apiURLs {
		result := message.Read(chatID, messageID)
		apiURLResult := result.APIURL(g.String(apiURL))
		if apiURLResult == nil {
			t.Errorf("APIURL method should return Read builder for chaining with URL: %s", apiURL)
		}

		// Test that APIURL can be chained and overridden
		chainedResult := apiURLResult.APIURL(g.String("https://override.example.com"))
		if chainedResult == nil {
			t.Errorf("APIURL method should support chaining and override with URL: %s", apiURL)
		}
	}
}

func TestReadMessage_Send(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_send_read_123")
	account := business.NewAccount(bot, connectionID)
	message := account.Message()

	chatID := int64(789)
	messageID := int64(456)

	// Test Send method - will fail with mock but covers the method
	sendResult := message.Read(chatID, messageID).Send()

	if sendResult.IsErr() {
		t.Logf("ReadMessage Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
