package suggested

import (
	"testing"
	"time"

	"github.com/enetx/tg/suggested"
)

func TestNew(t *testing.T) {
	params := suggested.New()

	if params == nil {
		t.Error("New should return non-nil PostParameters")
	}

	// Empty parameters should return nil from Build
	result := params.Build()
	if result != nil {
		t.Error("Build should return nil for empty parameters")
	}

	// Std should behave the same as Build
	result = params.Std()
	if result != nil {
		t.Error("Std should return nil for empty parameters")
	}
}

func TestPostParameters_PriceStars(t *testing.T) {
	amount := int64(100)
	params := suggested.New().PriceStars(amount)

	if params == nil {
		t.Error("PriceStars should return non-nil PostParameters for chaining")
	}

	result := params.Build()
	if result == nil {
		t.Error("Build should return non-nil SuggestedPostParameters when price is set")
	}

	if result.Price == nil {
		t.Error("Price should be set")
	}

	if result.Price.Currency != "XTR" {
		t.Errorf("Expected currency XTR, got %s", result.Price.Currency)
	}

	if result.Price.Amount != amount {
		t.Errorf("Expected amount %d, got %d", amount, result.Price.Amount)
	}
}

func TestPostParameters_PriceTon(t *testing.T) {
	nanotons := int64(1000000000)
	params := suggested.New().PriceTon(nanotons)

	if params == nil {
		t.Error("PriceTon should return non-nil PostParameters for chaining")
	}

	result := params.Build()
	if result == nil {
		t.Error("Build should return non-nil SuggestedPostParameters when price is set")
	}

	if result.Price == nil {
		t.Error("Price should be set")
	}

	if result.Price.Currency != "TON" {
		t.Errorf("Expected currency TON, got %s", result.Price.Currency)
	}

	if result.Price.Amount != nanotons {
		t.Errorf("Expected amount %d, got %d", nanotons, result.Price.Amount)
	}
}

func TestPostParameters_SendDate(t *testing.T) {
	sendTime := time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)
	params := suggested.New().SendDate(sendTime)

	if params == nil {
		t.Error("SendDate should return non-nil PostParameters for chaining")
	}

	result := params.Build()
	if result == nil {
		t.Error("Build should return non-nil SuggestedPostParameters when send date is set")
	}

	expectedUnix := sendTime.Unix()
	if result.SendDate != expectedUnix {
		t.Errorf("Expected SendDate %d, got %d", expectedUnix, result.SendDate)
	}
}

func TestPostParameters_SendAfter(t *testing.T) {
	duration := 24 * time.Hour
	beforeCall := time.Now()
	params := suggested.New().SendAfter(duration)
	afterCall := time.Now()

	if params == nil {
		t.Error("SendAfter should return non-nil PostParameters for chaining")
	}

	result := params.Build()
	if result == nil {
		t.Error("Build should return non-nil SuggestedPostParameters when send date is set")
	}

	expectedMinUnix := beforeCall.Add(duration).Unix()
	expectedMaxUnix := afterCall.Add(duration).Unix()

	if result.SendDate < expectedMinUnix || result.SendDate > expectedMaxUnix {
		t.Errorf("SendDate %d should be between %d and %d", result.SendDate, expectedMinUnix, expectedMaxUnix)
	}
}

func TestPostParameters_Chaining(t *testing.T) {
	amount := int64(500)
	sendTime := time.Date(2025, 6, 1, 15, 30, 0, 0, time.UTC)

	// Test method chaining with both price and send date
	params := suggested.New().PriceStars(amount).SendDate(sendTime)

	if params == nil {
		t.Error("Method chaining should work")
	}

	result := params.Build()
	if result == nil {
		t.Error("Build should return non-nil SuggestedPostParameters")
	}

	// Check price
	if result.Price == nil {
		t.Error("Price should be set")
	}
	if result.Price.Currency != "XTR" {
		t.Errorf("Expected currency XTR, got %s", result.Price.Currency)
	}
	if result.Price.Amount != amount {
		t.Errorf("Expected amount %d, got %d", amount, result.Price.Amount)
	}

	// Check send date
	expectedUnix := sendTime.Unix()
	if result.SendDate != expectedUnix {
		t.Errorf("Expected SendDate %d, got %d", expectedUnix, result.SendDate)
	}
}

func TestPostParameters_PriceOverride(t *testing.T) {
	starsAmount := int64(100)
	tonAmount := int64(2000000000)

	// Test that later price calls override previous ones
	params := suggested.New().PriceStars(starsAmount).PriceTon(tonAmount)

	result := params.Build()
	if result == nil {
		t.Error("Build should return non-nil SuggestedPostParameters")
	}

	if result.Price == nil {
		t.Error("Price should be set")
	}

	// Should have TON price (the last one set)
	if result.Price.Currency != "TON" {
		t.Errorf("Expected final currency TON, got %s", result.Price.Currency)
	}
	if result.Price.Amount != tonAmount {
		t.Errorf("Expected final amount %d, got %d", tonAmount, result.Price.Amount)
	}
}

func TestPostParameters_EdgeCases(t *testing.T) {
	// Test minimum valid Stars amount
	params := suggested.New().PriceStars(5)
	result := params.Build()
	if result == nil || result.Price == nil {
		t.Error("Should handle minimum Stars amount")
	}
	if result.Price.Amount != 5 {
		t.Error("Should preserve minimum Stars amount")
	}

	// Test maximum valid Stars amount
	params = suggested.New().PriceStars(100000)
	result = params.Build()
	if result == nil || result.Price == nil {
		t.Error("Should handle maximum Stars amount")
	}
	if result.Price.Amount != 100000 {
		t.Error("Should preserve maximum Stars amount")
	}

	// Test minimum valid TON amount
	params = suggested.New().PriceTon(10000000)
	result = params.Build()
	if result == nil || result.Price == nil {
		t.Error("Should handle minimum TON amount")
	}
	if result.Price.Amount != 10000000 {
		t.Error("Should preserve minimum TON amount")
	}

	// Test zero values
	params = suggested.New().PriceStars(0)
	result = params.Build()
	if result == nil || result.Price == nil {
		t.Error("Should handle zero Stars amount")
	}
	if result.Price.Amount != 0 {
		t.Error("Should preserve zero Stars amount")
	}

	// Test with zero time - time.Time{}.Unix() returns -62135596800, not 0
	params = suggested.New().SendDate(time.Time{})
	result = params.Build()
	if result == nil {
		t.Error("Build should return non-nil when SendDate is set (even with zero time)")
	}
	zeroTimeUnix := time.Time{}.Unix()
	if result != nil && result.SendDate != zeroTimeUnix {
		t.Errorf("Should preserve zero time SendDate, got %d, expected %d", result.SendDate, zeroTimeUnix)
	}

	// Test with non-zero SendDate but nil price
	params = suggested.New().SendDate(time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC))
	result = params.Build()
	if result == nil {
		t.Error("Build should return non-nil when SendDate is non-zero")
	}
	if result != nil && result.SendDate == 0 {
		t.Error("Should preserve non-zero SendDate")
	}
}

func TestPostParameters_BuildIdempotent(t *testing.T) {
	amount := int64(250)
	sendTime := time.Date(2025, 3, 15, 10, 0, 0, 0, time.UTC)

	params := suggested.New().PriceStars(amount).SendDate(sendTime)

	// Call Build multiple times
	result1 := params.Build()
	result2 := params.Build()

	if result1.Price.Amount != result2.Price.Amount {
		t.Error("Build should be idempotent for Price.Amount")
	}

	if result1.Price.Currency != result2.Price.Currency {
		t.Error("Build should be idempotent for Price.Currency")
	}

	if result1.SendDate != result2.SendDate {
		t.Error("Build should be idempotent for SendDate")
	}
}

func TestPostParameters_StdAlias(t *testing.T) {
	amount := int64(75)
	params := suggested.New().PriceStars(amount)

	// Test that Std() and Build() return equivalent results
	buildResult := params.Build()
	stdResult := params.Std()

	if buildResult == nil || stdResult == nil {
		t.Error("Both Build and Std should return non-nil results")
	}

	if buildResult.Price.Amount != stdResult.Price.Amount {
		t.Error("Build and Std should return equivalent Price.Amount")
	}

	if buildResult.Price.Currency != stdResult.Price.Currency {
		t.Error("Build and Std should return equivalent Price.Currency")
	}

	if buildResult.SendDate != stdResult.SendDate {
		t.Error("Build and Std should return equivalent SendDate")
	}
}
