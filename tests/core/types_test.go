package core_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	. "github.com/enetx/tg/core"
)

// mockBotAPI implements BotAPI interface for testing
type mockBotAPI struct {
	dispatcher *ext.Dispatcher
	updater    *ext.Updater
	bot        *gotgbot.Bot
}

func (m *mockBotAPI) Dispatcher() *ext.Dispatcher {
	return m.dispatcher
}

func (m *mockBotAPI) Updater() *ext.Updater {
	return m.updater
}

func (m *mockBotAPI) Raw() *gotgbot.Bot {
	return m.bot
}

func TestBotAPIInterface(t *testing.T) {
	// Test that our mock implements BotAPI interface
	var _ BotAPI = (*mockBotAPI)(nil)

	// Create mock instance
	mock := &mockBotAPI{
		dispatcher: &ext.Dispatcher{},
		updater:    &ext.Updater{},
		bot:        &gotgbot.Bot{},
	}

	// Test Dispatcher method
	if mock.Dispatcher() == nil {
		t.Error("Expected Dispatcher to return non-nil dispatcher")
	}

	// Test Updater method
	if mock.Updater() == nil {
		t.Error("Expected Updater to return non-nil updater")
	}

	// Test Raw method
	if mock.Raw() == nil {
		t.Error("Expected Raw to return non-nil bot")
	}
}

func TestBotAPIInterfaceCompliance(t *testing.T) {
	// Test that the interface defines expected methods
	mock := &mockBotAPI{
		dispatcher: &ext.Dispatcher{},
		updater:    &ext.Updater{},
		bot:        &gotgbot.Bot{},
	}

	// Verify interface methods work correctly
	dispatcher := mock.Dispatcher()
	if dispatcher != mock.dispatcher {
		t.Error("Expected Dispatcher to return the correct dispatcher instance")
	}

	updater := mock.Updater()
	if updater != mock.updater {
		t.Error("Expected Updater to return the correct updater instance")
	}

	bot := mock.Raw()
	if bot != mock.bot {
		t.Error("Expected Raw to return the correct bot instance")
	}

	// Test that bot is properly set
	if bot == nil {
		t.Error("Expected bot to be non-nil")
	}
}
