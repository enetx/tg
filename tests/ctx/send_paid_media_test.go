package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/keyboard"
)

func TestContext_SendPaidMedia(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	starCount := int64(100)

	result := ctx.SendPaidMedia(starCount)

	if result == nil {
		t.Error("Expected SendPaidMedia builder to be created")
	}

	// Test method chaining
	chained := result.Silent()
	if chained == nil {
		t.Error("Expected Silent method to return builder")
	}
}

func TestContext_SendPaidMediaChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	starCount := int64(100)

	result := ctx.SendPaidMedia(starCount).
		Silent().
		Protect().
		To(123)

	if result == nil {
		t.Error("Expected SendPaidMedia builder to be created")
	}

	// Test that builder is functional
	_ = result
}

func TestSendPaidMedia_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	starCount := int64(100)

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendPaidMedia(starCount).Send()

	if sendResult.IsErr() {
		t.Logf("SendPaidMedia Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendPaidMedia(starCount).
		Caption(g.String("<b>Paid media content</b>")).
		HTML().
		Silent().
		Protect().
		To(123).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendPaidMedia configured Send failed as expected: %v", configuredSendResult.Err())
	}
}

func TestSendPaidMedia_Photo(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	starCount := int64(100)
	photoResult := file.Input(g.String("https://example.com/photo.jpg"))
	if photoResult.IsErr() {
		t.Skip("Unable to create photo input for testing")
	}
	photo := input.PaidPhoto(photoResult.Unwrap())
	if ctx.SendPaidMedia(starCount).Photo(photo) == nil {
		t.Error("Photo should return builder")
	}
}

func TestSendPaidMedia_Video(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	starCount := int64(100)
	videoResult := file.Input(g.String("https://example.com/video.mp4"))
	if videoResult.IsErr() {
		t.Skip("Unable to create video input for testing")
	}
	video := input.PaidVideo(videoResult.Unwrap())
	if ctx.SendPaidMedia(starCount).Video(video) == nil {
		t.Error("Video should return builder")
	}
}

func TestSendPaidMedia_Business(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	starCount := int64(100)
	if ctx.SendPaidMedia(starCount).Business(g.String("biz_123")) == nil {
		t.Error("Business should return builder")
	}
}

func TestSendPaidMedia_Payload(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	starCount := int64(100)
	if ctx.SendPaidMedia(starCount).Payload(g.String("custom_payload_123")) == nil {
		t.Error("Payload should return builder")
	}
}

func TestSendPaidMedia_Markdown(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	starCount := int64(100)
	if ctx.SendPaidMedia(starCount).Markdown() == nil {
		t.Error("Markdown should return builder")
	}
}

func TestSendPaidMedia_ShowCaptionAbove(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	starCount := int64(100)
	if ctx.SendPaidMedia(starCount).ShowCaptionAbove() == nil {
		t.Error("ShowCaptionAbove should return builder")
	}
}

func TestSendPaidMedia_AllowPaidBroadcast(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	starCount := int64(100)
	if ctx.SendPaidMedia(starCount).AllowPaidBroadcast() == nil {
		t.Error("AllowPaidBroadcast should return builder")
	}
}

func TestSendPaidMedia_ReplyTo(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	starCount := int64(100)
	if ctx.SendPaidMedia(starCount).ReplyTo(123) == nil {
		t.Error("ReplyTo should return builder")
	}
}

func TestSendPaidMedia_Markup(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	starCount := int64(100)
	btn1 := keyboard.NewButton().Text(g.String("Buy Now")).Callback(g.String("buy_now"))
	if ctx.SendPaidMedia(starCount).Markup(keyboard.Inline().Button(btn1)) == nil {
		t.Error("Markup should return builder")
	}
}

func TestSendPaidMedia_SendWithError(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	starCount := int64(100)

	// Test Send with no media - should return error
	sendResult := ctx.SendPaidMedia(starCount).Send()

	if !sendResult.IsErr() {
		t.Error("Send should return error for no media specified")
	} else {
		t.Logf("SendPaidMedia Send with no media returned error as expected: %v", sendResult.Err())
	}
}

func TestSendPaidMedia_SendWithValidMedia(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	starCount := int64(100)

	// Add valid media
	photoResult := file.Input(g.String("https://example.com/photo.jpg"))
	if photoResult.IsErr() {
		t.Skip("Unable to create photo input for testing")
	}
	photo := input.PaidPhoto(photoResult.Unwrap())

	sendResult := ctx.SendPaidMedia(starCount).Photo(photo).Send()

	// This will fail with mock bot, but covers the valid media path
	if sendResult.IsErr() {
		t.Logf("SendPaidMedia Send with valid media failed as expected: %v", sendResult.Err())
	}
}

func TestSendPaidMedia_SendWithTooManyMedia(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	starCount := int64(100)

	// Add 11 media items (over the limit of 10)
	photoResult := file.Input(g.String("https://example.com/photo.jpg"))
	if photoResult.IsErr() {
		t.Skip("Unable to create photo input for testing")
	}

	builder := ctx.SendPaidMedia(starCount)

	// Add 11 photos to exceed the limit
	for i := 0; i < 11; i++ {
		photo := input.PaidPhoto(photoResult.Unwrap())
		builder = builder.Photo(photo)
	}

	sendResult := builder.Send()

	if !sendResult.IsErr() {
		t.Error("Send should return error for too many media items")
	} else {
		t.Logf("SendPaidMedia Send with too many media returned error as expected: %v", sendResult.Err())
	}
}

func TestSendPaidMedia_SendWithInvalidStarCount(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// Test with star count = 0 (below minimum)
	photoResult := file.Input(g.String("https://example.com/photo.jpg"))
	if photoResult.IsErr() {
		t.Skip("Unable to create photo input for testing")
	}
	photo := input.PaidPhoto(photoResult.Unwrap())

	sendResult := ctx.SendPaidMedia(0).Photo(photo).Send()

	if !sendResult.IsErr() {
		t.Error("Send should return error for invalid star count (0)")
	} else {
		t.Logf("SendPaidMedia Send with invalid star count (0) returned error as expected: %v", sendResult.Err())
	}

	// Test with star count = 10001 (above maximum)
	sendResult2 := ctx.SendPaidMedia(10001).Photo(photo).Send()

	if !sendResult2.IsErr() {
		t.Error("Send should return error for invalid star count (10001)")
	} else {
		t.Logf("SendPaidMedia Send with invalid star count (10001) returned error as expected: %v", sendResult2.Err())
	}
}

func TestSendPaidMedia_APIURLWithExistingRequestOpts(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	starCount := int64(100)

	// First set Timeout to create RequestOpts, then test APIURL
	result := ctx.SendPaidMedia(starCount).
		Timeout(15 * time.Second).                         // This creates RequestOpts
		APIURL(g.String("https://custom.api.example.com")) // This should use existing RequestOpts

	if result == nil {
		t.Error("APIURL with existing RequestOpts should return builder")
	}
}
