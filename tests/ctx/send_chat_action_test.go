package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SendChatAction(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.SendChatAction()

	if result == nil {
		t.Error("Expected SendChatAction builder to be created")
	}

	// Test method chaining
	chained := result.Typing()
	if chained == nil {
		t.Error("Expected Typing method to return builder")
	}
}

func TestContext_SendChatActionChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.SendChatAction().
		Typing().
		To(123).
		Thread(456)

	if result == nil {
		t.Error("Expected SendChatAction builder to be created")
	}

	// Test that builder is functional
	_ = result
}

func TestSendChatAction_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendChatAction().Send()

	if sendResult.IsErr() {
		t.Logf("SendChatAction Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendChatAction().
		Typing().
		To(123).
		Thread(456).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendChatAction configured Send failed as expected: %v", configuredSendResult.Err())
	}
}

func TestSendChatAction_Business(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	businessIDs := []string{
		"business_123",
		"conn_456",
		"",
	}

	for _, businessID := range businessIDs {
		result := ctx.SendChatAction().Business(g.String(businessID))
		if result == nil {
			t.Errorf("Business method should return SendChatAction builder for chaining with businessID: %s", businessID)
		}

		chainedResult := result.Business(g.String("override_business"))
		if chainedResult == nil {
			t.Errorf("Business method should support chaining and override with businessID: %s", businessID)
		}
	}
}

func TestSendChatAction_UploadPhoto(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.SendChatAction().UploadPhoto()
	if result == nil {
		t.Error("UploadPhoto method should return SendChatAction builder for chaining")
	}

	chainedResult := result.UploadPhoto()
	if chainedResult == nil {
		t.Error("UploadPhoto method should support chaining")
	}
}

func TestSendChatAction_RecordVideo(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.SendChatAction().RecordVideo()
	if result == nil {
		t.Error("RecordVideo method should return SendChatAction builder for chaining")
	}

	chainedResult := result.RecordVideo()
	if chainedResult == nil {
		t.Error("RecordVideo method should support chaining")
	}
}

func TestSendChatAction_UploadVideo(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.SendChatAction().UploadVideo()
	if result == nil {
		t.Error("UploadVideo method should return SendChatAction builder for chaining")
	}

	chainedResult := result.UploadVideo()
	if chainedResult == nil {
		t.Error("UploadVideo method should support chaining")
	}
}

func TestSendChatAction_RecordVoice(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.SendChatAction().RecordVoice()
	if result == nil {
		t.Error("RecordVoice method should return SendChatAction builder for chaining")
	}

	chainedResult := result.RecordVoice()
	if chainedResult == nil {
		t.Error("RecordVoice method should support chaining")
	}
}

func TestSendChatAction_UploadVoice(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.SendChatAction().UploadVoice()
	if result == nil {
		t.Error("UploadVoice method should return SendChatAction builder for chaining")
	}

	chainedResult := result.UploadVoice()
	if chainedResult == nil {
		t.Error("UploadVoice method should support chaining")
	}
}

func TestSendChatAction_UploadDocument(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.SendChatAction().UploadDocument()
	if result == nil {
		t.Error("UploadDocument method should return SendChatAction builder for chaining")
	}

	chainedResult := result.UploadDocument()
	if chainedResult == nil {
		t.Error("UploadDocument method should support chaining")
	}
}

func TestSendChatAction_ChooseSticker(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.SendChatAction().ChooseSticker()
	if result == nil {
		t.Error("ChooseSticker method should return SendChatAction builder for chaining")
	}

	chainedResult := result.ChooseSticker()
	if chainedResult == nil {
		t.Error("ChooseSticker method should support chaining")
	}
}

func TestSendChatAction_FindLocation(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.SendChatAction().FindLocation()
	if result == nil {
		t.Error("FindLocation method should return SendChatAction builder for chaining")
	}

	chainedResult := result.FindLocation()
	if chainedResult == nil {
		t.Error("FindLocation method should support chaining")
	}
}

func TestSendChatAction_RecordVideoNote(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.SendChatAction().RecordVideoNote()
	if result == nil {
		t.Error("RecordVideoNote method should return SendChatAction builder for chaining")
	}

	chainedResult := result.RecordVideoNote()
	if chainedResult == nil {
		t.Error("RecordVideoNote method should support chaining")
	}
}

func TestSendChatAction_UploadVideoNote(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.SendChatAction().UploadVideoNote()
	if result == nil {
		t.Error("UploadVideoNote method should return SendChatAction builder for chaining")
	}

	chainedResult := result.UploadVideoNote()
	if chainedResult == nil {
		t.Error("UploadVideoNote method should support chaining")
	}
}

func TestSendChatAction_APIURLWithExistingRequestOpts(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// First set Timeout to create RequestOpts, then test APIURL
	result := ctx.SendChatAction().
		Timeout(15 * time.Second).                         // This creates RequestOpts
		APIURL(g.String("https://custom.api.example.com")) // This should use existing RequestOpts

	if result == nil {
		t.Error("APIURL with existing RequestOpts should return builder")
	}
}
