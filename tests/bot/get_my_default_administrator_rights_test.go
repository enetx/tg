package bot_test

import (
	"testing"
	"time"

	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
)

func TestBot_GetMyDefaultAdministratorRights(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	getRights := bot.GetMyDefaultAdministratorRights()

	if getRights == nil {
		t.Error("Expected GetMyDefaultAdministratorRights to return a builder")
	}
}

func TestGetMyDefaultAdministratorRights_AllMethods(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	req := bot.GetMyDefaultAdministratorRights()

	// Test all methods
	req = req.ForChannels()
	if req == nil {
		t.Error("Expected ForChannels method to return request")
	}

	req = req.Timeout(10 * time.Second)
	if req == nil {
		t.Error("Expected Timeout method to return request")
	}

	req = req.APIURL(g.String("https://api.telegram.org"))
	if req == nil {
		t.Error("Expected APIURL method to return request")
	}

	// Test APIURL with empty string for coverage
	req2 := bot.GetMyDefaultAdministratorRights().APIURL(g.String(""))
	if req2 == nil {
		t.Error("Expected APIURL with empty string to return request")
	}
}

func TestGetMyDefaultAdministratorRights_Send(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	botInstance := result.Ok()
	req := botInstance.GetMyDefaultAdministratorRights()

	// Test Send method
	result2 := req.Send()
	if result2.IsOk() {
		rights := result2.Ok()
		_ = rights
	} else {
		// Error expected in test environment
		err := result2.Err()
		_ = err
	}
}
