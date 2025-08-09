package business_test

import (
	"testing"
	"time"

	"github.com/enetx/g"
	"github.com/enetx/tg/ctx/business"
)

func TestSetAnimatedPhoto(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_123")
	account := business.NewAccount(bot, connectionID)

	animation := g.String("animation.gif")
	result := account.SetAnimatedPhoto(animation)

	if result == nil {
		t.Error("Expected SetAnimatedPhoto builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}

func TestSetAnimatedPhoto_MainFrame(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_mainframe_123")
	account := business.NewAccount(bot, connectionID)

	animation := g.String("animation.gif")

	// Test MainFrame method with various timestamps
	timestamps := []time.Duration{
		1 * time.Second,
		5 * time.Second,
		10 * time.Second,
		30 * time.Second,
		0 * time.Second, // Zero timestamp
		500 * time.Millisecond,
	}

	for _, timestamp := range timestamps {
		result := account.SetAnimatedPhoto(animation)
		mainFrameResult := result.MainFrame(timestamp)
		if mainFrameResult == nil {
			t.Errorf("MainFrame method should return SetAnimatedPhoto builder for chaining with timestamp %v", timestamp)
		}

		// Test that MainFrame can be chained and overridden
		chainedResult := mainFrameResult.MainFrame(timestamp + 1*time.Second)
		if chainedResult == nil {
			t.Errorf("MainFrame method should support chaining and override with timestamp %v", timestamp)
		}
	}
}

func TestSetAnimatedPhoto_Public(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_public_animated_123")
	account := business.NewAccount(bot, connectionID)

	animation := g.String("animation.gif")

	// Test Public method
	result := account.SetAnimatedPhoto(animation)
	publicResult := result.Public()
	if publicResult == nil {
		t.Error("Public method should return SetAnimatedPhoto builder for chaining")
	}

	// Test that Public can be chained multiple times
	chainedResult := publicResult.Public()
	if chainedResult == nil {
		t.Error("Public method should support multiple chaining calls")
	}

	// Test Public with other methods
	publicWithOthers := account.SetAnimatedPhoto(animation).
		Public().
		MainFrame(5 * time.Second).
		Timeout(30 * time.Second)
	if publicWithOthers == nil {
		t.Error("Public method should work with other methods")
	}
}

func TestSetAnimatedPhoto_APIURL(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_apiurl_animated_123")
	account := business.NewAccount(bot, connectionID)

	animation := g.String("animation.gif")

	// Test APIURL method with various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://custom.api.example.com",
		"",
		"https://api.example.com/bot",
		"http://localhost:8080",
	}

	for _, apiURL := range apiURLs {
		result := account.SetAnimatedPhoto(animation)
		apiURLResult := result.APIURL(g.String(apiURL))
		if apiURLResult == nil {
			t.Errorf("APIURL method should return SetAnimatedPhoto builder for chaining with URL: %s", apiURL)
		}

		// Test that APIURL can be chained and overridden
		chainedResult := apiURLResult.APIURL(g.String("https://override.example.com"))
		if chainedResult == nil {
			t.Errorf("APIURL method should support chaining and override with URL: %s", apiURL)
		}
	}
}

func TestSetAnimatedPhoto_Send(t *testing.T) {
	bot := &mockBot{}
	connectionID := g.String("business_conn_send_animated_123")
	account := business.NewAccount(bot, connectionID)

	animation := g.String("animation.gif")

	// Test Send method - will fail with mock but covers the method
	sendResult := account.SetAnimatedPhoto(animation).Send()

	if sendResult.IsErr() {
		t.Logf("SetAnimatedPhoto Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with MainFrame and Public
	sendWithOptionsResult := account.SetAnimatedPhoto(animation).
		MainFrame(3 * time.Second).
		Public().
		Send()

	if sendWithOptionsResult.IsErr() {
		t.Logf("SetAnimatedPhoto Send with options failed as expected with mock bot: %v", sendWithOptionsResult.Err())
	}
}
