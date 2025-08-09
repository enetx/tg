package ctx_test

import (
	"testing"
	"time"

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

func TestSavePreparedInlineMessage_AllowBotChats(t *testing.T) {
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
	allowBotsResult := saveResult.AllowBotChats()
	if allowBotsResult == nil {
		t.Error("AllowBotChats method should return SavePreparedInlineMessage builder for chaining")
	}

	chainedResult := allowBotsResult.AllowBotChats()
	if chainedResult == nil {
		t.Error("AllowBotChats method should support chaining")
	}
}

func TestSavePreparedInlineMessage_AllowGroupChats(t *testing.T) {
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
	allowGroupsResult := saveResult.AllowGroupChats()
	if allowGroupsResult == nil {
		t.Error("AllowGroupChats method should return SavePreparedInlineMessage builder for chaining")
	}

	chainedResult := allowGroupsResult.AllowGroupChats()
	if chainedResult == nil {
		t.Error("AllowGroupChats method should support chaining")
	}
}

func TestSavePreparedInlineMessage_AllowChannelChats(t *testing.T) {
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
	allowChannelsResult := saveResult.AllowChannelChats()
	if allowChannelsResult == nil {
		t.Error("AllowChannelChats method should return SavePreparedInlineMessage builder for chaining")
	}

	chainedResult := allowChannelsResult.AllowChannelChats()
	if chainedResult == nil {
		t.Error("AllowChannelChats method should support chaining")
	}
}

func TestSavePreparedInlineMessage_Timeout(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)
	messageContent := input.Text(g.String("Test content"))
	result := inline.NewArticle(g.String("1"), g.String("Test Article"), messageContent)

	timeouts := []time.Duration{
		time.Second * 10,
		time.Second * 30,
		time.Minute,
		time.Minute * 5,
		0,
	}

	for _, timeout := range timeouts {
		saveResult := ctx.SavePreparedInlineMessage(userID, result)
		timeoutResult := saveResult.Timeout(timeout)
		if timeoutResult == nil {
			t.Errorf("Timeout method should return SavePreparedInlineMessage builder for chaining with timeout: %v", timeout)
		}

		chainedResult := timeoutResult.Timeout(time.Second * 15)
		if chainedResult == nil {
			t.Errorf("Timeout method should support chaining and override with timeout: %v", timeout)
		}
	}
}

func TestSavePreparedInlineMessage_APIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)
	messageContent := input.Text(g.String("Test content"))
	result := inline.NewArticle(g.String("1"), g.String("Test Article"), messageContent)

	apiURLs := []string{
		"https://api.telegram.org",
		"https://custom.api.example.com",
		"",
		"https://api.example.com/bot",
		"http://localhost:8080",
	}

	for _, apiURL := range apiURLs {
		saveResult := ctx.SavePreparedInlineMessage(userID, result)
		apiURLResult := saveResult.APIURL(g.String(apiURL))
		if apiURLResult == nil {
			t.Errorf("APIURL method should return SavePreparedInlineMessage builder for chaining with URL: %s", apiURL)
		}

		chainedResult := apiURLResult.APIURL(g.String("https://override.example.com"))
		if chainedResult == nil {
			t.Errorf("APIURL method should support chaining and override with URL: %s", apiURL)
		}
	}
}

func TestSavePreparedInlineMessage_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)
	messageContent := input.Text(g.String("Test content"))
	result := inline.NewArticle(g.String("1"), g.String("Test Article"), messageContent)

	sendResult := ctx.SavePreparedInlineMessage(userID, result).Send()

	if sendResult.IsErr() {
		t.Logf("SavePreparedInlineMessage Send failed as expected with mock bot: %v", sendResult.Err())
	}

	sendWithOptionsResult := ctx.SavePreparedInlineMessage(userID, result).
		AllowUserChats().
		AllowBotChats().
		AllowGroupChats().
		AllowChannelChats().
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.telegram.org")).
		Send()

	if sendWithOptionsResult.IsErr() {
		t.Logf("SavePreparedInlineMessage Send with options failed as expected with mock bot: %v", sendWithOptionsResult.Err())
	}
}
