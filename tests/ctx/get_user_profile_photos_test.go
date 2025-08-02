package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/tg/ctx"
)

func TestContext_GetUserProfilePhotos(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123)

	result := ctx.GetUserProfilePhotos(userID)

	if result == nil {
		t.Error("Expected GetUserProfilePhotos builder to be created")
	}

	// Test method chaining
	chained := result.Offset(0)
	if chained == nil {
		t.Error("Expected Offset method to return builder")
	}
}

func TestContext_GetUserProfilePhotosChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123)

	result := ctx.GetUserProfilePhotos(userID).
		Offset(0).
		Limit(10)

	if result == nil {
		t.Error("Expected GetUserProfilePhotos builder to be created")
	}

	// Test that builder is functional
	_ = result
}
