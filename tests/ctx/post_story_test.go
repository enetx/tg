package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/input"
)

func TestContext_PostStory(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_123")
	content := input.StoryPhoto(g.String("photo.jpg"))

	result := ctx.PostStory(businessConnectionID, content)

	if result == nil {
		t.Error("Expected PostStory builder to be created")
	}

	// Test method chaining
	withCaption := result.Caption(g.String("Story caption"))
	if withCaption == nil {
		t.Error("Expected Caption method to return builder")
	}
}
