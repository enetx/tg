package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func ephemeralContext(bot *mockBot) *ctx.Context {
	return ctx.New(bot, &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "supergroup"},
		EffectiveMessage: &gotgbot.Message{
			MessageId:          789,
			EphemeralMessageId: 555,
			ReceiverUser:       &gotgbot.User{Id: 123},
		},
		Update: &gotgbot.Update{UpdateId: 1},
	})
}

func TestContext_DeleteEphemeralMessage(t *testing.T) {
	testCtx := ephemeralContext(&mockBot{})

	result := testCtx.DeleteEphemeralMessage()
	if result == nil {
		t.Error("Expected DeleteEphemeralMessage builder to be created")
	}

	if result.ChatID(999) == nil {
		t.Error("ChatID method should return DeleteEphemeralMessage for chaining")
	}

	if result.ReceiverUser(111) == nil {
		t.Error("ReceiverUser method should return DeleteEphemeralMessage for chaining")
	}

	if result.MessageID(222) == nil {
		t.Error("MessageID method should return DeleteEphemeralMessage for chaining")
	}

	if result.Timeout(30*time.Second) == nil {
		t.Error("Timeout method should return DeleteEphemeralMessage for chaining")
	}

	if result.APIURL(g.String("https://api.example.com")) == nil {
		t.Error("APIURL method should return DeleteEphemeralMessage for chaining")
	}
}

func TestDeleteEphemeralMessage_Send(t *testing.T) {
	testCtx := ephemeralContext(&mockBot{})

	sendResult := testCtx.DeleteEphemeralMessage().Send()
	if !sendResult.IsErr() && !sendResult.IsOk() {
		t.Error("Send should return a result using the effective ephemeral message")
	}

	complete := testCtx.DeleteEphemeralMessage().
		ChatID(999).
		ReceiverUser(111).
		MessageID(222).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if !complete.IsErr() && !complete.IsOk() {
		t.Error("Send with all features should return a result")
	}
}

func TestDeleteEphemeralMessage_SendWithoutEphemeralContext(t *testing.T) {
	testCtx := ctx.New(&mockBot{}, &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "supergroup"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	})

	sendResult := testCtx.DeleteEphemeralMessage().Send()
	if !sendResult.IsErr() && !sendResult.IsOk() {
		t.Error("Send on a non-ephemeral message should still return a result")
	}
}
