package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/keyboard"
)

func TestContext_SendVideo(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("video.mp4")

	result := ctx.SendVideo(filename)

	if result == nil {
		t.Error("Expected SendVideo builder to be created")
	}

	// Test method chaining
	chained := result.Caption(g.String("Video caption"))
	if chained == nil {
		t.Error("Expected caption method to return builder")
	}
}

func TestContext_SendVideoChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("video.mp4")

	result := ctx.SendVideo(filename).
		Caption(g.String("Test video")).
		HTML().
		Silent()

	if result == nil {
		t.Error("Expected SendVideo builder to be created")
	}

	// Test continued chaining
	final := result.To(123)
	if final == nil {
		t.Error("Expected To method to return builder")
	}
}

func TestSendVideo_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("test_video.mp4")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendVideo(filename).Send()

	if sendResult.IsErr() {
		t.Logf("SendVideo Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendVideo(filename).
		Caption(g.String("Test <b>video</b> with HTML")).
		HTML().
		Width(640).
		Height(480).
		Duration(120).
		Silent().
		Protect().
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendVideo configured Send failed as expected: %v", configuredSendResult.Err())
	}
}

func TestSendVideo_CaptionEntities(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("video.mp4")
	ent := entities.New(g.String("Bold text")).Bold(g.String("Bold"))
	if ctx.SendVideo(filename).CaptionEntities(ent) == nil { t.Error("CaptionEntities should return builder") }
}

func TestSendVideo_After(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("video.mp4")
	if ctx.SendVideo(filename).After(time.Minute) == nil { t.Error("After should return builder") }
}

func TestSendVideo_DeleteAfter(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("video.mp4")
	if ctx.SendVideo(filename).DeleteAfter(time.Hour) == nil { t.Error("DeleteAfter should return builder") }
}

func TestSendVideo_Thumbnail(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("video.mp4")
	if ctx.SendVideo(filename).Thumbnail(g.String("thumb.jpg")) == nil { t.Error("Thumbnail should return builder") }
}

func TestSendVideo_Spoiler(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("video.mp4")
	if ctx.SendVideo(filename).Spoiler() == nil { t.Error("Spoiler should return builder") }
}

func TestSendVideo_Streamable(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("video.mp4")
	if ctx.SendVideo(filename).Streamable() == nil { t.Error("Streamable should return builder") }
}

func TestSendVideo_Markdown(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("video.mp4")
	if ctx.SendVideo(filename).Markdown() == nil { t.Error("Markdown should return builder") }
}

func TestSendVideo_Markup(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("video.mp4")
	btn1 := keyboard.NewButton().Text(g.String("Watch Video")).Callback(g.String("watch_video"))
	if ctx.SendVideo(filename).Markup(keyboard.Inline().Button(btn1)) == nil { t.Error("Markup should return builder") }
}

func TestSendVideo_ReplyTo(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("video.mp4")
	if ctx.SendVideo(filename).ReplyTo(123) == nil { t.Error("ReplyTo should return builder") }
}

func TestSendVideo_Business(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("video.mp4")
	if ctx.SendVideo(filename).Business(g.String("biz_123")) == nil { t.Error("Business should return builder") }
}

func TestSendVideo_Thread(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("video.mp4")
	if ctx.SendVideo(filename).Thread(456) == nil { t.Error("Thread should return builder") }
}

func TestSendVideo_ShowCaptionAboveMedia(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("video.mp4")
	if ctx.SendVideo(filename).ShowCaptionAboveMedia() == nil { t.Error("ShowCaptionAboveMedia should return builder") }
}

func TestSendVideo_Cover(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("video.mp4")
	if ctx.SendVideo(filename).Cover(g.String("cover.jpg")) == nil { t.Error("Cover should return builder") }
}

func TestSendVideo_StartAt(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("video.mp4")
	if ctx.SendVideo(filename).StartAt(30*time.Second) == nil { t.Error("StartAt should return builder") }
}

func TestSendVideo_ApplyMetadata(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("video.mp4")
	if ctx.SendVideo(filename).ApplyMetadata() == nil { t.Error("ApplyMetadata should return builder") }
}

func TestSendVideo_GenerateThumbnail(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("video.mp4")
	if ctx.SendVideo(filename).GenerateThumbnail() == nil { t.Error("GenerateThumbnail should return builder") }
}

func TestSendVideo_ErrorHandling(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	
	// Test with invalid filename that should cause file.Input to fail
	invalidFilename := g.String("")  // Empty filename should cause an error
	result := ctx.SendVideo(invalidFilename)
	
	// The builder should still be created even with error
	if result == nil {
		t.Error("SendVideo should return builder even with invalid filename")
	}
	
	// Test that Send() properly handles the error
	sendResult := result.Send()
	if !sendResult.IsErr() {
		t.Error("Send should fail with empty filename")
	} else {
		t.Logf("Send failed as expected with empty filename: %v", sendResult.Err())
	}
	
	// Test with nonexistent file
	nonexistentFile := g.String("/nonexistent/path/to/video.mp4")
	result2 := ctx.SendVideo(nonexistentFile)
	if result2 == nil {
		t.Error("SendVideo should return builder even with nonexistent file")
	}
	
	sendResult2 := result2.Send()
	if !sendResult2.IsErr() {
		t.Error("Send should fail with nonexistent file")
	} else {
		t.Logf("Send failed as expected with nonexistent file: %v", sendResult2.Err())
	}
}
