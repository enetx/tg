package business_test

import (
	"testing"
	"time"

	"github.com/enetx/g"
	"github.com/enetx/tg/ctx/business"
)

func TestSetGiftSettings(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_gift_settings_123")
	account := business.NewAccount(bot, connectionID)

	result := account.SetGiftSettings()

	if result == nil {
		t.Error("Expected SetGiftSettings builder to be created")
	}
}

func TestSetGiftSettings_ShowGiftButton(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_show_gift_button_123")
	account := business.NewAccount(bot, connectionID)

	result := account.SetGiftSettings().ShowGiftButton()
	if result == nil {
		t.Error("ShowGiftButton method should return builder for chaining")
	}
}

func TestSetGiftSettings_AcceptUnlimitedGifts(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_accept_unlimited_123")
	account := business.NewAccount(bot, connectionID)

	result := account.SetGiftSettings().AcceptUnlimitedGifts()
	if result == nil {
		t.Error("AcceptUnlimitedGifts method should return builder for chaining")
	}

	chained := result.AcceptLimitedGifts()
	if chained == nil {
		t.Error("AcceptUnlimitedGifts should support chaining with other methods")
	}
}

func TestSetGiftSettings_AcceptLimitedGifts(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_accept_limited_123")
	account := business.NewAccount(bot, connectionID)

	result := account.SetGiftSettings().AcceptLimitedGifts()
	if result == nil {
		t.Error("AcceptLimitedGifts method should return builder for chaining")
	}

	chained := result.AcceptUniqueGifts()
	if chained == nil {
		t.Error("AcceptLimitedGifts should support chaining with other methods")
	}
}

func TestSetGiftSettings_AcceptUniqueGifts(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_accept_unique_123")
	account := business.NewAccount(bot, connectionID)

	result := account.SetGiftSettings().AcceptUniqueGifts()
	if result == nil {
		t.Error("AcceptUniqueGifts method should return builder for chaining")
	}

	chained := result.AcceptPremiumSubscription()
	if chained == nil {
		t.Error("AcceptUniqueGifts should support chaining with other methods")
	}
}

func TestSetGiftSettings_AcceptPremiumSubscription(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_accept_premium_123")
	account := business.NewAccount(bot, connectionID)

	result := account.SetGiftSettings().AcceptPremiumSubscription()
	if result == nil {
		t.Error("AcceptPremiumSubscription method should return builder for chaining")
	}

	chained := result.AcceptGiftsFromChannels()
	if chained == nil {
		t.Error("AcceptPremiumSubscription should support chaining with other methods")
	}
}

func TestSetGiftSettings_AcceptGiftsFromChannels(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_accept_from_channels_123")
	account := business.NewAccount(bot, connectionID)

	result := account.SetGiftSettings().AcceptGiftsFromChannels()
	if result == nil {
		t.Error("AcceptGiftsFromChannels method should return builder for chaining")
	}

	chained := result.ShowGiftButton()
	if chained == nil {
		t.Error("AcceptGiftsFromChannels should support chaining with other methods")
	}
}

func TestSetGiftSettings_AcceptAllGifts(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_accept_all_123")
	account := business.NewAccount(bot, connectionID)

	result := account.SetGiftSettings().AcceptAllGifts()
	if result == nil {
		t.Error("AcceptAllGifts method should return builder for chaining")
	}

	withButton := result.ShowGiftButton()
	if withButton == nil {
		t.Error("AcceptAllGifts should work with ShowGiftButton")
	}
}

func TestSetGiftSettings_AcceptNoGifts(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_accept_no_123")
	account := business.NewAccount(bot, connectionID)

	result := account.SetGiftSettings().AcceptNoGifts()
	if result == nil {
		t.Error("AcceptNoGifts method should return builder for chaining")
	}

	// Can still enable specific gifts after AcceptNoGifts
	withUnlimited := result.AcceptUnlimitedGifts()
	if withUnlimited == nil {
		t.Error("AcceptNoGifts should allow chaining with specific accept methods")
	}
}

func TestSetGiftSettings_Timeout(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_timeout_gift_settings_123")
	account := business.NewAccount(bot, connectionID)

	timeouts := []time.Duration{
		time.Second,
		5 * time.Second,
		30 * time.Second,
		time.Minute,
	}

	for _, timeout := range timeouts {
		result := account.SetGiftSettings().Timeout(timeout)
		if result == nil {
			t.Errorf("Timeout method should return builder for duration: %v", timeout)
		}
	}
}

func TestSetGiftSettings_APIURL(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_apiurl_gift_settings_123")
	account := business.NewAccount(bot, connectionID)

	apiURLs := []string{
		"https://api.telegram.org",
		"https://custom.api.example.com",
		"http://localhost:8080",
	}

	for _, apiURL := range apiURLs {
		result := account.SetGiftSettings().APIURL(g.String(apiURL))
		if result == nil {
			t.Errorf("APIURL method should return builder for URL: %s", apiURL)
		}
	}
}

func TestSetGiftSettings_FullChaining(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_full_chain_123")
	account := business.NewAccount(bot, connectionID)

	result := account.SetGiftSettings().
		ShowGiftButton().
		AcceptUnlimitedGifts().
		AcceptLimitedGifts().
		AcceptPremiumSubscription().
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.telegram.org"))

	if result == nil {
		t.Error("Full method chaining should work")
	}
}

func TestSetGiftSettings_Send(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_send_gift_settings_123")
	account := business.NewAccount(bot, connectionID)

	sendResult := account.SetGiftSettings().
		ShowGiftButton().
		AcceptAllGifts().
		Send()

	if sendResult.IsErr() {
		t.Logf("SetGiftSettings Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
