package bot_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/types/updates"
)

func TestBot_Polling(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	polling := bot.Polling()

	if polling == nil {
		t.Error("Expected polling to be non-nil")
	}
}

func TestPolling_AllowedUpdates(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()

	// Note: This test may fail if updates package is not available
	// commenting it out for now to avoid import issues
	/*
		polling := bot.Polling().AllowedUpdates(
			updates.UpdateTypeMessage,
			updates.UpdateTypeCallbackQuery,
			updates.UpdateTypeInlineQuery,
		)

		if polling == nil {
			t.Error("Expected polling to be non-nil")
		}
	*/

	// Test without allowed updates for now
	polling := bot.Polling()
	if polling == nil {
		t.Error("Expected polling to be non-nil")
	}
}

// Test additional polling methods that need more coverage
func TestPolling_AllowedUpdatesExtended(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	polling := bot.Polling()

	// Import needed for update types
	updateTypes := []updates.UpdateType{updates.Message, updates.CallbackQuery}
	result2 := polling.AllowedUpdates(updateTypes...)

	if result2 != polling {
		t.Error("Expected AllowedUpdates to return the same polling instance")
	}
}

func TestPolling_DropPendingUpdates(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	polling := bot.Polling().DropPendingUpdates()

	if polling == nil {
		t.Error("Expected polling to be non-nil")
	}

	// Check that the option was set
	if !polling.Opts().DropPendingUpdates {
		t.Error("Expected DropPendingUpdates to be true")
	}
}

func TestPolling_EnableWebhookDeletion(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	polling := bot.Polling().EnableWebhookDeletion()

	if polling == nil {
		t.Error("Expected polling to be non-nil")
	}

	// Check that the option was set
	if !polling.Opts().EnableWebhookDeletion {
		t.Error("Expected EnableWebhookDeletion to be true")
	}
}

func TestPolling_Timeout(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	timeout := int64(30)
	polling := bot.Polling().Timeout(timeout)

	if polling == nil {
		t.Error("Expected polling to be non-nil")
	}

	// Check that the timeout was set
	if polling.Opts().GetUpdatesOpts.Timeout != timeout {
		t.Errorf("Expected timeout %d, got %d", timeout, polling.Opts().GetUpdatesOpts.Timeout)
	}
}

func TestPolling_Limit(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	limit := int64(100)
	polling := bot.Polling().Limit(limit)

	if polling == nil {
		t.Error("Expected polling to be non-nil")
	}

	// Check that the limit was set
	if polling.Opts().GetUpdatesOpts.Limit != limit {
		t.Errorf("Expected limit %d, got %d", limit, polling.Opts().GetUpdatesOpts.Limit)
	}
}

func TestPolling_Offset(t *testing.T) {
	token := g.String("123456:ABCDEF-test-token-here")
	result := bot.New(token).DisableTokenCheck().Build()

	if result.IsErr() {
		t.Errorf("Failed to create bot: %v", result.Err())
		return
	}

	bot := result.Ok()
	offset := int64(123)
	polling := bot.Polling().Offset(offset)

	if polling == nil {
		t.Error("Expected polling to be non-nil")
	}

	// Check that the offset was set
	if polling.Opts().GetUpdatesOpts.Offset != offset {
		t.Errorf("Expected offset %d, got %d", offset, polling.Opts().GetUpdatesOpts.Offset)
	}
}
