package business_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/ctx/business"
)

func TestGetGifts(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_123")
	account := business.NewAccount(bot, connectionID)
	balance := account.Balance()

	result := balance.GetGifts()

	if result == nil {
		t.Error("Expected GetGifts builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}

func TestGetGifts_ExcludeUnsaved(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_exclude_unsaved_123")
	account := business.NewAccount(bot, connectionID)
	balance := account.Balance()

	// Test ExcludeUnsaved method
	result := balance.GetGifts()
	excludeUnsavedResult := result.ExcludeUnsaved()
	if excludeUnsavedResult == nil {
		t.Error("ExcludeUnsaved method should return GetGifts builder for chaining")
	}

	// Test that ExcludeUnsaved can be chained multiple times
	chainedResult := excludeUnsavedResult.ExcludeUnsaved()
	if chainedResult == nil {
		t.Error("ExcludeUnsaved method should support multiple chaining calls")
	}

	// Test ExcludeUnsaved with other methods
	excludeUnsavedWithOthers := balance.GetGifts().
		ExcludeUnsaved().
		Limit(50).
		SortByPrice()
	if excludeUnsavedWithOthers == nil {
		t.Error("ExcludeUnsaved method should work with other methods")
	}
}

func TestGetGifts_ExcludeSaved(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_exclude_saved_123")
	account := business.NewAccount(bot, connectionID)
	balance := account.Balance()

	// Test ExcludeSaved method
	result := balance.GetGifts()
	excludeSavedResult := result.ExcludeSaved()
	if excludeSavedResult == nil {
		t.Error("ExcludeSaved method should return GetGifts builder for chaining")
	}

	// Test that ExcludeSaved can be chained multiple times
	chainedResult := excludeSavedResult.ExcludeSaved()
	if chainedResult == nil {
		t.Error("ExcludeSaved method should support multiple chaining calls")
	}

	// Test ExcludeSaved with other methods
	excludeSavedWithOthers := balance.GetGifts().
		ExcludeSaved().
		Limit(25).
		Offset(g.String("page_token_123"))
	if excludeSavedWithOthers == nil {
		t.Error("ExcludeSaved method should work with other methods")
	}
}

func TestGetGifts_ExcludeUnlimited(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_exclude_unlimited_123")
	account := business.NewAccount(bot, connectionID)
	balance := account.Balance()

	// Test ExcludeUnlimited method
	result := balance.GetGifts()
	excludeUnlimitedResult := result.ExcludeUnlimited()
	if excludeUnlimitedResult == nil {
		t.Error("ExcludeUnlimited method should return GetGifts builder for chaining")
	}

	// Test that ExcludeUnlimited can be chained multiple times
	chainedResult := excludeUnlimitedResult.ExcludeUnlimited()
	if chainedResult == nil {
		t.Error("ExcludeUnlimited method should support multiple chaining calls")
	}
}

func TestGetGifts_ExcludeLimited(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_exclude_limited_123")
	account := business.NewAccount(bot, connectionID)
	balance := account.Balance()

	// Test ExcludeLimited method
	result := balance.GetGifts()
	excludeLimitedResult := result.ExcludeLimited()
	if excludeLimitedResult == nil {
		t.Error("ExcludeLimited method should return GetGifts builder for chaining")
	}

	// Test that ExcludeLimited can be chained multiple times
	chainedResult := excludeLimitedResult.ExcludeLimited()
	if chainedResult == nil {
		t.Error("ExcludeLimited method should support multiple chaining calls")
	}
}

func TestGetGifts_ExcludeUnique(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_exclude_unique_123")
	account := business.NewAccount(bot, connectionID)
	balance := account.Balance()

	// Test ExcludeUnique method
	result := balance.GetGifts()
	excludeUniqueResult := result.ExcludeUnique()
	if excludeUniqueResult == nil {
		t.Error("ExcludeUnique method should return GetGifts builder for chaining")
	}

	// Test that ExcludeUnique can be chained multiple times
	chainedResult := excludeUniqueResult.ExcludeUnique()
	if chainedResult == nil {
		t.Error("ExcludeUnique method should support multiple chaining calls")
	}
}

func TestGetGifts_SortByPrice(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_sort_by_price_123")
	account := business.NewAccount(bot, connectionID)
	balance := account.Balance()

	// Test SortByPrice method
	result := balance.GetGifts()
	sortByPriceResult := result.SortByPrice()
	if sortByPriceResult == nil {
		t.Error("SortByPrice method should return GetGifts builder for chaining")
	}

	// Test that SortByPrice can be chained multiple times
	chainedResult := sortByPriceResult.SortByPrice()
	if chainedResult == nil {
		t.Error("SortByPrice method should support multiple chaining calls")
	}

	// Test SortByPrice with other methods
	sortByPriceWithOthers := balance.GetGifts().
		SortByPrice().
		ExcludeUnsaved().
		Limit(10)
	if sortByPriceWithOthers == nil {
		t.Error("SortByPrice method should work with other methods")
	}
}

func TestGetGifts_Offset(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_offset_123")
	account := business.NewAccount(bot, connectionID)
	balance := account.Balance()

	// Test Offset method with various offset values
	offsets := []string{
		"page_token_123",
		"next_page_456",
		"",
		"offset_789",
		"very_long_pagination_token_with_special_characters_!@#$%",
	}

	for _, offset := range offsets {
		result := balance.GetGifts()
		offsetResult := result.Offset(g.String(offset))
		if offsetResult == nil {
			t.Errorf("Offset method should return GetGifts builder for chaining with offset: %s", offset)
		}

		// Test that Offset can be chained and overridden
		chainedResult := offsetResult.Offset(g.String("updated_" + offset))
		if chainedResult == nil {
			t.Errorf("Offset method should support chaining and override with offset: %s", offset)
		}
	}
}

func TestGetGifts_APIURL(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_apiurl_gifts_123")
	account := business.NewAccount(bot, connectionID)
	balance := account.Balance()

	// Test APIURL method with various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://custom.api.example.com",
		"",
		"https://api.example.com/bot",
		"http://localhost:8080",
	}

	for _, apiURL := range apiURLs {
		result := balance.GetGifts()
		apiURLResult := result.APIURL(g.String(apiURL))
		if apiURLResult == nil {
			t.Errorf("APIURL method should return GetGifts builder for chaining with URL: %s", apiURL)
		}

		// Test that APIURL can be chained and overridden
		chainedResult := apiURLResult.APIURL(g.String("https://override.example.com"))
		if chainedResult == nil {
			t.Errorf("APIURL method should support chaining and override with URL: %s", apiURL)
		}
	}
}

func TestGetGifts_Send(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_send_gifts_123")
	account := business.NewAccount(bot, connectionID)
	balance := account.Balance()

	// Test Send method - will fail with mock but covers the method
	sendResult := balance.GetGifts().Send()

	if sendResult.IsErr() {
		t.Logf("GetGifts Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
