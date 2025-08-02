package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/inline"
	"github.com/enetx/tg/input"
)

func TestContext_SavePreparedInlineMessage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)
	messageContent := input.Text(g.String("Test content"))
	result := inline.NewArticle(g.String("1"), g.String("Test Article"), messageContent)

	saveResult := ctx.SavePreparedInlineMessage(userID, result)

	if saveResult == nil {
		t.Error("Expected SavePreparedInlineMessage builder to be created")
	}

	// Test method chaining
	allowUserChats := saveResult.AllowUserChats()
	if allowUserChats == nil {
		t.Error("Expected AllowUserChats method to return builder")
	}
}
