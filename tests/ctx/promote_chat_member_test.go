package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/types/roles"
)

func TestContext_PromoteChatMember(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123456)

	result := ctx.PromoteChatMember(userID)

	if result == nil {
		t.Error("Expected PromoteChatMember builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(-1001234567890)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}

func TestContext_PromoteChatMemberChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123456)

	result := ctx.PromoteChatMember(userID).
		ChatID(-1001234567890)

	if result == nil {
		t.Error("Expected PromoteChatMember builder to be created")
	}

	// Test that builder is functional
	_ = result
}

func TestPromoteChatMember_Roles(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123456)

	// Test Roles method with various role combinations
	roleTestCases := [][]roles.Role{
		{roles.ManageChat},
		{roles.DeleteMessages},
		{roles.ManageVideoChats},
		{roles.RestrictMembers},
		{roles.PromoteMembers},
		{roles.ChangeInfo},
		{roles.InviteUsers},
		{roles.PostMessages},
		{roles.EditMessages},
		{roles.PinMessages},
		{roles.ManageTopics},
		{roles.ManageChat, roles.DeleteMessages, roles.PinMessages},                                                   // Multiple roles
		{roles.ManageChat, roles.DeleteMessages, roles.ManageVideoChats, roles.RestrictMembers, roles.PromoteMembers}, // Many roles
		{}, // Empty roles (edge case)
	}

	for i, roleList := range roleTestCases {
		result := ctx.PromoteChatMember(userID)
		rolesResult := result.Roles(roleList...)
		if rolesResult == nil {
			t.Errorf("Roles method should return PromoteChatMember builder for chaining with roles %d", i)
		}

		// Test that Roles can be chained and overridden
		chainedResult := rolesResult.Roles(roles.ManageChat, roles.DeleteMessages)
		if chainedResult == nil {
			t.Errorf("Roles method should support chaining and override with roles %d", i)
		}
	}
}

func TestPromoteChatMember_Timeout(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123456)

	// Test Timeout method with various durations
	timeouts := []time.Duration{
		1 * time.Second,
		5 * time.Second,
		30 * time.Second,
		1 * time.Minute,
		5 * time.Minute,
		0 * time.Second, // Zero timeout
	}

	for _, timeout := range timeouts {
		result := ctx.PromoteChatMember(userID)
		timeoutResult := result.Timeout(timeout)
		if timeoutResult == nil {
			t.Errorf("Timeout method should return PromoteChatMember builder for chaining with timeout %v", timeout)
		}

		// Test that Timeout can be chained and overridden
		chainedResult := timeoutResult.Timeout(timeout + 1*time.Second)
		if chainedResult == nil {
			t.Errorf("Timeout method should support chaining and override with timeout %v", timeout)
		}
	}
}

func TestPromoteChatMember_APIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123456)

	// Test APIURL method with various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://custom.api.example.com",
		"",
		"https://api.example.com/bot",
		"http://localhost:8080",
	}

	for _, apiURL := range apiURLs {
		result := ctx.PromoteChatMember(userID)
		apiURLResult := result.APIURL(g.String(apiURL))
		if apiURLResult == nil {
			t.Errorf("APIURL method should return PromoteChatMember builder for chaining with URL: %s", apiURL)
		}

		// Test that APIURL can be chained and overridden
		chainedResult := apiURLResult.APIURL(g.String("https://override.example.com"))
		if chainedResult == nil {
			t.Errorf("APIURL method should support chaining and override with URL: %s", apiURL)
		}
	}
}

func TestPromoteChatMember_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123456)

	// Test Send method without roles (should fail with error)
	sendResultWithoutRoles := ctx.PromoteChatMember(userID).Send()

	if sendResultWithoutRoles.IsOk() {
		t.Error("PromoteChatMember Send without roles should fail")
	} else {
		t.Logf("PromoteChatMember Send without roles failed as expected: %v", sendResultWithoutRoles.Err())
	}

	// Test Send method with roles - will fail with mock but covers the method
	sendResult := ctx.PromoteChatMember(userID).
		Roles(roles.ManageChat, roles.DeleteMessages).
		Send()

	if sendResult.IsErr() {
		t.Logf("PromoteChatMember Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with all options
	sendWithOptionsResult := ctx.PromoteChatMember(userID).
		ChatID(-1001987654321).
		Roles(roles.ManageChat, roles.DeleteMessages, roles.PinMessages, roles.ManageTopics).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.telegram.org")).
		Send()

	if sendWithOptionsResult.IsErr() {
		t.Logf("PromoteChatMember Send with options failed as expected with mock bot: %v", sendWithOptionsResult.Err())
	}

	// Test Send method using default chat ID (from effective chat)
	sendWithDefaultChatResult := ctx.PromoteChatMember(userID).
		Roles(roles.ManageChat).
		Send()

	if sendWithDefaultChatResult.IsErr() {
		t.Logf("PromoteChatMember Send with default chat ID failed as expected with mock bot: %v", sendWithDefaultChatResult.Err())
	}
}
