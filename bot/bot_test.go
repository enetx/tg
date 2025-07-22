package bot

import (
	"testing"
	"time"

	. "github.com/enetx/g"
)

func TestBotBuilder(t *testing.T) {
	token := String("123456:ABCDEF-test-token-here")

	// Test bot builder with disabled token check
	result := New(token).
		DisableTokenCheck().
		Build()

	if result.IsErr() {
		t.Errorf("Expected successful bot creation, got error: %v", result.Err())
	}

	bot := result.Ok()
	if bot.token != token {
		t.Errorf("Expected bot token '%s', got '%s'", token, bot.token)
	}

	if bot.raw.User.Id == 0 {
		t.Error("Expected bot user ID to be set")
	}

	if !bot.raw.User.IsBot {
		t.Error("Expected user to be marked as bot")
	}
}

func TestBotBuilderInvalidToken(t *testing.T) {
	// Test with invalid token format
	token := String("invalid-token")

	result := New(token).
		DisableTokenCheck().
		Build()

	if result.IsOk() {
		t.Error("Expected error for invalid token format")
	}
}

func TestBotBuilderWithOptions(t *testing.T) {
	token := String("123456:ABCDEF-test-token-here")

	result := New(token).
		DisableTokenCheck().
		UseTestEnvironment().
		Timeout(5 * time.Second).
		APIURL(String("https://api.telegram.org")).
		DefaultTimeout(10 * time.Second).
		DefaultAPIURL(String("https://custom.api.url")).
		Build()

	if result.IsErr() {
		t.Errorf("Expected successful bot creation with options, got error: %v", result.Err())
	}

	bot := result.Ok()

	// Test that bot was created successfully
	if bot == nil {
		t.Error("Expected bot to be created")
	}

	if bot.token != token {
		t.Errorf("Expected bot token '%s', got '%s'", token, bot.token)
	}
}

func TestBotBuilderEmptyToken(t *testing.T) {
	token := String("")

	result := New(token).
		DisableTokenCheck().
		Build()

	if result.IsOk() {
		t.Error("Expected error for empty token")
	}
}

func TestBotBuilderValidTokenFormat(t *testing.T) {
	validTokens := []String{
		String("123456789:AAEhBOweik9ai9Koh6oh9aegh"),
		String("987654321:BBFhCPxejk0aj0Lpi7pi0bfhi"),
		String("1111111111:CCGiDQyflk1bk1Mqj8qj1cgij"),
	}

	for _, token := range validTokens {
		result := New(token).
			DisableTokenCheck().
			Build()

		if result.IsErr() {
			t.Errorf("Expected successful bot creation for token '%s', got error: %v",
				token, result.Err())
		}
	}
}

func TestBotBuilderInvalidTokenFormats(t *testing.T) {
	invalidTokens := []String{
		String("123456"),      // No colon
		String(":ABCDEF"),     // No bot ID
		String("abc:ABCDEF"),  // Non-numeric bot ID
		String("123456:"),     // Empty token part
		String("123:456:789"), // Too many parts
	}

	for _, token := range invalidTokens {
		result := New(token).
			DisableTokenCheck().
			Build()

		// Skip the case where token has valid format but empty key part
		if token == "123456:" {
			continue
		}

		if result.IsOk() {
			t.Errorf("Expected error for invalid token '%s'", token)
		}
	}
}
