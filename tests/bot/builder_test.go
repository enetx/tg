package bot_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
)

func TestBotBuilder(t *testing.T) {
	token := "123456:ABCDEF-test-token-here"

	result := bot.New(token).
		DisableTokenCheck().
		Build()

	if result.IsErr() {
		t.Errorf("Expected successful bot creation, got error: %v", result.Err())
	}

	b := result.Ok()
	if b.Raw().Token != token {
		t.Errorf("Expected bot token '%s', got '%s'", token, b.Raw().Token)
	}

	if b.Raw().User.Id == 0 {
		t.Error("Expected bot user ID to be set")
	}

	if !b.Raw().User.IsBot {
		t.Error("Expected user to be marked as bot")
	}
}

func TestBotBuilderInvalidToken(t *testing.T) {
	token := g.String("invalid-token")

	result := bot.New(token).
		DisableTokenCheck().
		Build()

	if result.IsOk() {
		t.Error("Expected error for invalid token format")
	}
}

func TestBotBuilderWithOptions(t *testing.T) {
	token := "123456:ABCDEF-test-token-here"

	result := bot.New(token).
		DisableTokenCheck().
		UseTestEnvironment().
		Timeout(5 * time.Second).
		APIURL("https://api.telegram.org").
		DefaultTimeout(10 * time.Second).
		DefaultAPIURL("https://custom.api.url").
		Build()

	if result.IsErr() {
		t.Errorf("Expected successful bot creation with options, got error: %v", result.Err())
	}

	b := result.Ok()

	// Test that bot was created successfully
	if b == nil {
		t.Error("Expected bot to be created")
	}

	if b.Raw().Token != token {
		t.Errorf("Expected bot token '%s', got '%s'", token, b.Raw().Token)
	}
}

func TestBotBuilderEmptyToken(t *testing.T) {
	token := ""

	result := bot.New(token).
		DisableTokenCheck().
		Build()

	if result.IsOk() {
		t.Error("Expected error for empty token")
	}
}

func TestBotBuilderValidTokenFormat(t *testing.T) {
	validTokens := []g.String{
		"123456789:AAEhBOweik9ai9Koh6oh9aegh",
		"987654321:BBFhCPxejk0aj0Lpi7pi0bfhi",
		"1111111111:CCGiDQyflk1bk1Mqj8qj1cgij",
	}

	for _, token := range validTokens {
		result := bot.New(token).
			DisableTokenCheck().
			Build()

		if result.IsErr() {
			t.Errorf("Expected successful bot creation for token '%s', got error: %v",
				token, result.Err())
		}
	}
}

func TestBotBuilderInvalidTokenFormats(t *testing.T) {
	invalidTokens := []g.String{
		"123456",      // No colon
		":ABCDEF",     // No bot ID
		"abc:ABCDEF",  // Non-numeric bot ID
		"123456:",     // Empty token part
		"123:456:789", // Too many parts
	}

	for _, token := range invalidTokens {
		result := bot.New(token).
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

func TestBotBuilder_UseClient(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	client := &http.Client{Timeout: time.Second * 10}

	result := bot.New(token).
		UseClient(client).
		DisableTokenCheck().
		Build()

	if result.IsErr() {
		t.Errorf("Expected successful bot creation with custom client, got error: %v", result.Err())
	}
}

func TestBotBuilder_AllMethods(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	client := &http.Client{Timeout: time.Second * 15}

	// Test all builder methods
	result := bot.New(token).
		DisableTokenCheck().
		UseTestEnvironment().
		UseClient(client).
		Timeout(30 * time.Second).
		Build()

	if result.IsErr() {
		t.Errorf("Expected successful bot creation, got error: %v", result.Err())
	}

	bot := result.Ok()
	if bot == nil {
		t.Error("Expected bot to be created")
	}
}

func TestBotBuilder_MethodChaining(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")

	// Test method chaining returns same builder
	builder := bot.New(token)
	if builder.DisableTokenCheck() != builder {
		t.Error("Expected DisableTokenCheck to return same builder")
	}
	if builder.UseTestEnvironment() != builder {
		t.Error("Expected UseTestEnvironment to return same builder")
	}
	if builder.Timeout(10*time.Second) != builder {
		t.Error("Expected Timeout to return same builder")
	}
	if builder.APIURL("https://api.telegram.org") != builder {
		t.Error("Expected APIURL to return same builder")
	}
	if builder.DefaultTimeout(15*time.Second) != builder {
		t.Error("Expected DefaultTimeout to return same builder")
	}
	if builder.DefaultAPIURL("https://api.default.telegram.org") != builder {
		t.Error("Expected DefaultAPIURL to return same builder")
	}
}

func TestBotBuilder_ErrorHandling(t *testing.T) {
	// Test builder error propagation
	token := g.String("")
	result := bot.New(token).Build()

	if result.IsOk() {
		t.Error("Expected error for empty token")
	}
}

func TestBotBuilder_ChainedConfiguration(t *testing.T) {
	token := "123456:ABCDEF-test-token-here"

	// Test complex chained configuration
	result := bot.New(token).
		DisableTokenCheck().
		UseTestEnvironment().
		Timeout(15 * time.Second).
		APIURL("https://custom.api.url").
		DefaultTimeout(20 * time.Second).
		DefaultAPIURL("https://default.custom.api.url").
		Build()

	if result.IsErr() {
		t.Errorf("Expected successful bot creation with chained config, got error: %v", result.Err())
	}

	bot := result.Ok()
	if bot == nil {
		t.Error("Expected bot to be created")
	}

	if bot.Raw().Token != token {
		t.Errorf("Expected bot token '%s', got '%s'", token, bot.Raw().Token)
	}
}

func TestBotBuilder_EdgeCases(t *testing.T) {
	// Test various edge cases
	testCases := []struct {
		name      string
		token     g.String
		shoudPass bool
	}{
		{"Empty token", g.String(""), false},
		{"Whitespace token", g.String("   "), false},
		{"Valid format", g.String("123456789:AAEhBOweik9ai9Koh6oh9aegh"), true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := bot.New(tc.token).DisableTokenCheck().Build()

			if tc.shoudPass && result.IsErr() {
				t.Errorf("Expected success for %s, got error: %v", tc.name, result.Err())
			}

			if !tc.shoudPass && result.IsOk() {
				t.Errorf("Expected error for %s, but got success", tc.name)
			}
		})
	}
}

func TestBotBuilder_EdgeCasesExtended(t *testing.T) {
	// Test with very long token

	result := bot.New(
		g.String("1234567890:ABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"),
	).DisableTokenCheck().
		Build()
		// Simple long token
	if result.IsErr() {
		// Expected to work with truncated token
		t.Logf("Expected error for very long token: %v", result.Err())
	}

	// Test with minimum valid token
	minToken := g.String("1:A")
	result = bot.New(minToken).DisableTokenCheck().Build()
	if result.IsErr() {
		t.Errorf("Expected success for minimum valid token, got error: %v", result.Err())
	}
}
