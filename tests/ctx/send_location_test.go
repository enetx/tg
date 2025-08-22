package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/reply"
	"github.com/enetx/tg/types/effects"
)

func TestContext_SendLocation(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	lat, lon := 40.7128, -74.0060

	result := ctx.SendLocation(lat, lon)

	if result == nil {
		t.Error("Expected SendLocation builder to be created")
	}

	// Test method chaining
	chained := result.Silent()
	if chained == nil {
		t.Error("Expected Silent method to return builder")
	}
}

func TestContext_SendLocationChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	lat, lon := 40.7128, -74.0060

	result := ctx.SendLocation(lat, lon).
		Silent().
		To(123)

	if result == nil {
		t.Error("Expected SendLocation builder to be created")
	}

	// Test continued chaining
	final := result.Protect()
	if final == nil {
		t.Error("Expected Protect method to return builder")
	}

	// Test AllowPaidBroadcast method
	result = ctx.SendLocation(lat, lon).AllowPaidBroadcast()
	if result == nil {
		t.Error("AllowPaidBroadcast method should return SendLocation for chaining")
	}

	// Test Effect method
	result = ctx.SendLocation(lat, lon).Effect(effects.Fire)
	if result == nil {
		t.Error("Effect method should return SendLocation for chaining")
	}
}

func TestSendLocation_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	lat, lon := 40.7128, -74.0060

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendLocation(lat, lon).Send()

	if sendResult.IsErr() {
		t.Logf("SendLocation Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendLocation(lat, lon).
		Heading(90).
		ProximityAlertRadius(100).
		Silent().
		Protect().
		To(123).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendLocation configured Send failed as expected: %v", configuredSendResult.Err())
	}
}

func TestSendLocation_After(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	lat, lon := 40.7128, -74.0060
	if ctx.SendLocation(lat, lon).After(time.Minute) == nil {
		t.Error("After should return builder")
	}
}

func TestSendLocation_DeleteAfter(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	lat, lon := 40.7128, -74.0060
	if ctx.SendLocation(lat, lon).DeleteAfter(time.Hour) == nil {
		t.Error("DeleteAfter should return builder")
	}
}

func TestSendLocation_Markup(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	lat, lon := 40.7128, -74.0060
	btn1 := keyboard.NewButton().Text(g.String("View Map")).URL(g.String("https://maps.google.com"))
	if ctx.SendLocation(lat, lon).Markup(keyboard.Inline().Button(btn1)) == nil {
		t.Error("Markup should return builder")
	}
}

func TestSendLocation_LiveFor(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	lat, lon := 40.7128, -74.0060
	if ctx.SendLocation(lat, lon).LiveFor(time.Hour) == nil {
		t.Error("LiveFor should return builder")
	}
}

func TestSendLocation_HorizontalAccuracy(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	lat, lon := 40.7128, -74.0060
	if ctx.SendLocation(lat, lon).HorizontalAccuracy(1.5) == nil {
		t.Error("HorizontalAccuracy should return builder")
	}
}

func TestSendLocation_ReplyTo(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	lat, lon := 40.7128, -74.0060
	if ctx.SendLocation(lat, lon).Reply(reply.New(123)) == nil {
		t.Error("ReplyTo should return builder")
	}
}

func TestSendLocation_Business(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	lat, lon := 40.7128, -74.0060
	if ctx.SendLocation(lat, lon).Business(g.String("biz_123")) == nil {
		t.Error("Business should return builder")
	}
}

func TestSendLocation_Thread(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	lat, lon := 40.7128, -74.0060
	if ctx.SendLocation(lat, lon).Thread(456) == nil {
		t.Error("Thread should return builder")
	}
}

func TestSendLocation_DirectMessagesTopic(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	latitude := 40.7128
	longitude := -74.0060

	topicIDs := []int64{123, 456, 789, 0, -1}

	for _, topicID := range topicIDs {
		result := ctx.SendLocation(latitude, longitude).DirectMessagesTopic(topicID)
		if result == nil {
			t.Errorf("DirectMessagesTopic method should return SendLocation builder for chaining with topicID: %d", topicID)
		}

		chainedResult := result.DirectMessagesTopic(topicID + 100)
		if chainedResult == nil {
			t.Errorf("DirectMessagesTopic method should support chaining and override with topicID: %d", topicID)
		}
	}
}

func TestSendLocation_SuggestedPost(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	latitude := 40.7128
	longitude := -74.0060

	// Test with nil params
	result := ctx.SendLocation(latitude, longitude).SuggestedPost(nil)
	if result == nil {
		t.Error("SuggestedPost method should return SendLocation builder for chaining with nil params")
	}

	// Test chaining
	chainedResult := result.SuggestedPost(nil)
	if chainedResult == nil {
		t.Error("SuggestedPost method should support chaining")
	}
}

func TestSendLocation_APIURLWithExistingRequestOpts(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// First set Timeout to create RequestOpts, then test APIURL
	result := ctx.SendLocation(40.7128, -74.0060).
		Timeout(15 * time.Second).                         // This creates RequestOpts
		APIURL(g.String("https://custom.api.example.com")) // This should use existing RequestOpts

	if result == nil {
		t.Error("APIURL with existing RequestOpts should return builder")
	}
}
