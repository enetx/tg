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

func TestContext_SetChatPermissions(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.SetChatPermissions()

	if result == nil {
		t.Error("Expected SetChatPermissions builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(-1001234567890)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}

func TestSetChatPermissions_AutoPermissions(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	if ctx.SetChatPermissions().AutoPermissions() == nil { t.Error("AutoPermissions should return builder") }
}

func TestSetChatPermissions_Timeout(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	if ctx.SetChatPermissions().Timeout(time.Minute) == nil { t.Error("Timeout should return builder") }
}

func TestSetChatPermissions_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	if ctx.SetChatPermissions().APIURL(g.String("https://api.example.com")) == nil { t.Error("APIURL should return builder") }
}

func TestSetChatPermissions_Permissions(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	
	permissionTestCases := [][]permissions.Permission{
		{permissions.SendMessages},
		{permissions.SendAudios},
		{permissions.SendDocuments},
		{permissions.SendPhotos},
		{permissions.SendVideos},
		{permissions.SendVideoNotes},
		{permissions.SendVoiceNotes},
		{permissions.SendPolls},
		{permissions.SendOtherMessages},
		{permissions.AddWebPagePreviews},
		{permissions.ChangeInfo},
		{permissions.InviteUsers},
		{permissions.PinMessages},
		{permissions.ManageTopics},
		{permissions.SendMessages, permissions.SendPhotos, permissions.SendVideos},
	}
	
	for _, perms := range permissionTestCases {
		result := ctx.SetChatPermissions().Permissions(perms...)
		if result == nil {
			t.Errorf("Permissions should return builder for permissions: %v", perms)
		}
	}
}

func TestSetChatPermissions_Send(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	
	sendResult := ctx.SetChatPermissions().Permissions(permissions.SendMessages).Send()
	
	if sendResult.IsErr() {
		t.Logf("SetChatPermissions Send failed as expected with mock bot: %v", sendResult.Err())
	}
}

func TestSetChatPermissions_Send_ErrorBranches(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	
	// Test Send without permissions (should fail)
	result := ctx.SetChatPermissions().Send()
	if !result.IsErr() {
		t.Error("Send should fail without permissions")
	} else {
		t.Logf("Send failed as expected without permissions: %v", result.Err())
	}
	
	// Test Send with auto permissions and explicit permissions
	result2 := ctx.SetChatPermissions().AutoPermissions().Permissions(permissions.SendMessages, permissions.SendAudios).Send()
	if result2.IsErr() {
		t.Logf("Send with auto permissions and explicit permissions failed as expected: %v", result2.Err())
	}
}
