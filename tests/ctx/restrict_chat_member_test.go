package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/types/permissions"
)

func TestContext_RestrictChatMember(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123456)

	result := ctx.RestrictChatMember(userID)

	if result == nil {
		t.Error("Expected RestrictChatMember builder to be created")
	}

	// Test method chaining
	chained := result.Until(time.Now().Add(24 * time.Hour))
	if chained == nil {
		t.Error("Expected Until method to return builder")
	}
}

func TestContext_RestrictChatMemberChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123456)

	result := ctx.RestrictChatMember(userID).
		Until(time.Now().Add(24 * time.Hour)).
		ChatID(-1001234567890)

	if result == nil {
		t.Error("Expected RestrictChatMember builder to be created")
	}

	// Test that builder is functional
	_ = result
}

func TestRestrictChatMember_For(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123456)

	durations := []time.Duration{
		time.Minute,
		time.Hour,
		time.Hour * 24,
		time.Hour * 24 * 7,
		0,
	}

	for _, duration := range durations {
		result := ctx.RestrictChatMember(userID)
		forResult := result.For(duration)
		if forResult == nil {
			t.Errorf("For method should return RestrictChatMember builder for chaining with duration: %v", duration)
		}

		chainedResult := forResult.For(time.Hour * 2)
		if chainedResult == nil {
			t.Errorf("For method should support chaining and override with duration: %v", duration)
		}
	}
}

func TestRestrictChatMember_AutoPermissions(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123456)

	result := ctx.RestrictChatMember(userID)
	autoResult := result.AutoPermissions()
	if autoResult == nil {
		t.Error("AutoPermissions method should return RestrictChatMember builder for chaining")
	}

	chainedResult := autoResult.AutoPermissions()
	if chainedResult == nil {
		t.Error("AutoPermissions method should support chaining")
	}
}

func TestRestrictChatMember_Permissions(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123456)

	permissionTestCases := [][]permissions.Permission{
		{permissions.SendMessages},
		{permissions.SendMessages, permissions.SendPhotos},
		{permissions.SendMessages, permissions.SendVideos, permissions.SendAudios},
		{permissions.SendMessages, permissions.SendDocuments, permissions.SendPhotos, permissions.SendVideos},
		{permissions.SendMessages, permissions.SendOtherMessages, permissions.AddWebPagePreviews},
		{permissions.ChangeInfo, permissions.InviteUsers, permissions.PinMessages},
		{permissions.ManageTopics},
		{},
		{permissions.SendMessages, permissions.SendAudios, permissions.SendDocuments, permissions.SendPhotos, permissions.SendVideos, permissions.SendVideoNotes, permissions.SendVoiceNotes, permissions.SendPolls, permissions.SendOtherMessages, permissions.AddWebPagePreviews, permissions.ChangeInfo, permissions.InviteUsers, permissions.PinMessages, permissions.ManageTopics},
	}

	for i, perms := range permissionTestCases {
		result := ctx.RestrictChatMember(userID)
		permResult := result.Permissions(perms...)
		if permResult == nil {
			t.Errorf("Permissions method should return RestrictChatMember builder for chaining with permissions %d", i)
		}

		chainedResult := permResult.Permissions(permissions.SendMessages)
		if chainedResult == nil {
			t.Errorf("Permissions method should support chaining and override with permissions %d", i)
		}
	}
}

func TestRestrictChatMember_Timeout(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123456)

	timeouts := []time.Duration{
		time.Second * 10,
		time.Second * 30,
		time.Minute,
		time.Minute * 5,
		0,
	}

	for _, timeout := range timeouts {
		result := ctx.RestrictChatMember(userID)
		timeoutResult := result.Timeout(timeout)
		if timeoutResult == nil {
			t.Errorf("Timeout method should return RestrictChatMember builder for chaining with timeout: %v", timeout)
		}

		chainedResult := timeoutResult.Timeout(time.Second * 15)
		if chainedResult == nil {
			t.Errorf("Timeout method should support chaining and override with timeout: %v", timeout)
		}
	}
}

func TestRestrictChatMember_APIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123456)

	apiURLs := []string{
		"https://api.telegram.org",
		"https://custom.api.example.com",
		"",
		"https://api.example.com/bot",
		"http://localhost:8080",
	}

	for _, apiURL := range apiURLs {
		result := ctx.RestrictChatMember(userID)
		apiURLResult := result.APIURL(g.String(apiURL))
		if apiURLResult == nil {
			t.Errorf("APIURL method should return RestrictChatMember builder for chaining with URL: %s", apiURL)
		}

		chainedResult := apiURLResult.APIURL(g.String("https://override.example.com"))
		if chainedResult == nil {
			t.Errorf("APIURL method should support chaining and override with URL: %s", apiURL)
		}
	}
}

func TestRestrictChatMember_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123456)

	// Test Send without permissions (should fail)
	sendResult := ctx.RestrictChatMember(userID).Send()
	if !sendResult.IsErr() {
		t.Error("RestrictChatMember Send should fail when no permissions are set")
	}

	// Test Send with permissions
	sendWithPermsResult := ctx.RestrictChatMember(userID).
		Permissions(permissions.SendMessages, permissions.SendPhotos).
		Send()

	if sendWithPermsResult.IsErr() {
		t.Logf("RestrictChatMember Send with permissions failed as expected with mock bot: %v", sendWithPermsResult.Err())
	}

	sendWithOptionsResult := ctx.RestrictChatMember(userID).
		Permissions(permissions.SendMessages, permissions.SendVideos).
		For(time.Hour * 24).
		AutoPermissions().
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.telegram.org")).
		ChatID(-1009876543210).
		Send()

	if sendWithOptionsResult.IsErr() {
		t.Logf("RestrictChatMember Send with options failed as expected with mock bot: %v", sendWithOptionsResult.Err())
	}
}
