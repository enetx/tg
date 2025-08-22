package bot_test

import (
	"testing"
	"time"

	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/types/rights"
)

func TestBot_SetMyDefaultAdministratorRights(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	setRights := bot.SetMyDefaultAdministratorRights()

	if setRights == nil {
		t.Error("Expected SetMyDefaultAdministratorRights to return a builder")
	}
}

func TestSetMyDefaultAdministratorRights_AllMethods(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	req := bot.SetMyDefaultAdministratorRights()

	// Test Rights method with some rights
	req = req.Rights(rights.ManageChat, rights.DeleteMessages)
	if req == nil {
		t.Error("Expected Rights method to return request")
	}

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
	req2 := bot.SetMyDefaultAdministratorRights().APIURL(g.String(""))
	if req2 == nil {
		t.Error("Expected APIURL with empty string to return request")
	}
}

func TestSetMyDefaultAdministratorRights_Send(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	botInstance := result.Ok()
	req := botInstance.SetMyDefaultAdministratorRights().Rights(rights.ManageChat)

	result2 := req.Send()
	if result2.IsOk() {
		success := result2.Ok()
		_ = success
	} else {
		err := result2.Err()
		_ = err
	}
}
