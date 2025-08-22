package ctx_test

import (
	"os"
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/reply"
	"github.com/enetx/tg/types/effects"
)

func TestContext_SendAnimation(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("animation.gif")

	result := ctx.SendAnimation(filename)

	if result == nil {
		t.Error("Expected SendAnimation builder to be created")
	}

	// Test method chaining
	chained := result.Caption(g.String("Animation caption"))
	if chained == nil {
		t.Error("Expected caption method to return builder")
	}
}

func TestContext_SendAnimationChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("animation.gif")

	result := ctx.SendAnimation(filename).
		Caption(g.String("Test animation")).
		HTML().
		Width(400).
		Height(300)

	if result == nil {
		t.Error("Expected SendAnimation builder to be created")
	}

	// Test continued chaining
	final := result.Silent()
	if final == nil {
		t.Error("Expected Silent method to return builder")
	}

	// Test AllowPaidBroadcast method
	result = ctx.SendAnimation(filename).AllowPaidBroadcast()
	if result == nil {
		t.Error("AllowPaidBroadcast method should return SendAnimation for chaining")
	}

	// Test Effect method
	result = ctx.SendAnimation(filename).Effect(effects.Fire)
	if result == nil {
		t.Error("Effect method should return SendAnimation for chaining")
	}
}

func TestSendAnimation_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("test_animation.gif")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendAnimation(filename).Send()

	if sendResult.IsErr() {
		t.Logf("SendAnimation Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendAnimation(filename).
		Caption(g.String("Test <b>animation</b> with HTML")).
		HTML().
		Width(640).
		Height(480).
		Duration(10).
		Silent().
		Protect().
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendAnimation configured Send failed as expected: %v", configuredSendResult.Err())
	}
}

func TestSendAnimation_CaptionEntities(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("animation.gif")

	// Test with bold entities
	ent := entities.New(g.String("Bold text")).Bold(g.String("Bold"))
	result := ctx.SendAnimation(filename).CaptionEntities(ent)
	if result == nil {
		t.Error("CaptionEntities method should return SendAnimation builder for chaining")
	}

	// Test with italic entities
	entItalic := entities.New(g.String("Italic text")).Italic(g.String("Italic"))
	chainedResult := result.CaptionEntities(entItalic)
	if chainedResult == nil {
		t.Error("CaptionEntities method should support chaining and override")
	}
}

func TestSendAnimation_After(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("animation.gif")

	durations := []time.Duration{
		time.Second,
		time.Minute,
		time.Hour,
		0,
	}

	for _, duration := range durations {
		result := ctx.SendAnimation(filename).After(duration)
		if result == nil {
			t.Errorf("After method should return SendAnimation builder for chaining with duration: %v", duration)
		}

		chainedResult := result.After(time.Second * 30)
		if chainedResult == nil {
			t.Errorf("After method should support chaining and override with duration: %v", duration)
		}
	}
}

func TestSendAnimation_DeleteAfter(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("animation.gif")

	durations := []time.Duration{
		time.Second * 30,
		time.Minute * 5,
		time.Hour,
		0,
	}

	for _, duration := range durations {
		result := ctx.SendAnimation(filename).DeleteAfter(duration)
		if result == nil {
			t.Errorf("DeleteAfter method should return SendAnimation builder for chaining with duration: %v", duration)
		}

		chainedResult := result.DeleteAfter(time.Minute * 10)
		if chainedResult == nil {
			t.Errorf("DeleteAfter method should support chaining and override with duration: %v", duration)
		}
	}
}

func TestSendAnimation_Markdown(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("animation.gif")

	result := ctx.SendAnimation(filename).Markdown()
	if result == nil {
		t.Error("Markdown method should return SendAnimation builder for chaining")
	}

	chainedResult := result.Markdown()
	if chainedResult == nil {
		t.Error("Markdown method should support chaining")
	}
}

func TestSendAnimation_Markup(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("animation.gif")

	// Test with inline keyboard
	btn1 := keyboard.NewButton().Text(g.String("Test Button")).Callback(g.String("test_data"))
	btn2 := keyboard.NewButton().Text(g.String("Another Button")).Callback(g.String("another_data"))
	inlineKeyboard := keyboard.Inline().
		Button(btn1).
		Row().
		Button(btn2)

	result := ctx.SendAnimation(filename).Markup(inlineKeyboard)
	if result == nil {
		t.Error("Markup method should return SendAnimation builder for chaining with inline keyboard")
	}

	// Test with reply keyboard
	replyKeyboard := keyboard.Reply()

	chainedResult := result.Markup(replyKeyboard)
	if chainedResult == nil {
		t.Error("Markup method should support chaining and override with reply keyboard")
	}
}

func TestSendAnimation_Thumbnail(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("animation.gif")

	thumbnails := []string{
		"thumb.jpg",
		"thumbnail.png",
		"preview.webp",
	}

	for _, thumb := range thumbnails {
		result := ctx.SendAnimation(filename).Thumbnail(g.String(thumb))
		if result == nil {
			t.Errorf("Thumbnail method should return SendAnimation builder for chaining with thumbnail: %s", thumb)
		}

		chainedResult := result.Thumbnail(g.String("override.jpg"))
		if chainedResult == nil {
			t.Errorf("Thumbnail method should support chaining and override with thumbnail: %s", thumb)
		}
	}
}

func TestSendAnimation_ReplyTo(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("animation.gif")

	messageIDs := []int64{1, 123, 456, 999}

	for _, messageID := range messageIDs {
		result := ctx.SendAnimation(filename).Reply(reply.New(messageID))
		if result == nil {
			t.Errorf("ReplyTo method should return SendAnimation builder for chaining with messageID: %d", messageID)
		}

		chainedResult := result.Reply(reply.New(messageID + 100))
		if chainedResult == nil {
			t.Errorf("ReplyTo method should support chaining and override with messageID: %d", messageID)
		}
	}
}

func TestSendAnimation_Business(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("animation.gif")

	businessIDs := []string{
		"business_123",
		"conn_456",
		"",
	}

	for _, businessID := range businessIDs {
		result := ctx.SendAnimation(filename).Business(g.String(businessID))
		if result == nil {
			t.Errorf("Business method should return SendAnimation builder for chaining with businessID: %s", businessID)
		}

		chainedResult := result.Business(g.String("override_business"))
		if chainedResult == nil {
			t.Errorf("Business method should support chaining and override with businessID: %s", businessID)
		}
	}
}

func TestSendAnimation_Thread(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("animation.gif")

	threadIDs := []int64{1, 123, 456, 0}

	for _, threadID := range threadIDs {
		result := ctx.SendAnimation(filename).Thread(threadID)
		if result == nil {
			t.Errorf("Thread method should return SendAnimation builder for chaining with threadID: %d", threadID)
		}

		chainedResult := result.Thread(threadID + 100)
		if chainedResult == nil {
			t.Errorf("Thread method should support chaining and override with threadID: %d", threadID)
		}
	}
}

func TestSendAnimation_DirectMessagesTopic(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("animation.gif")

	topicIDs := []int64{123, 456, 789, 0, -1}

	for _, topicID := range topicIDs {
		result := ctx.SendAnimation(filename).DirectMessagesTopic(topicID)
		if result == nil {
			t.Errorf("DirectMessagesTopic method should return SendAnimation builder for chaining with topicID: %d", topicID)
		}

		chainedResult := result.DirectMessagesTopic(topicID + 100)
		if chainedResult == nil {
			t.Errorf("DirectMessagesTopic method should support chaining and override with topicID: %d", topicID)
		}
	}
}

func TestSendAnimation_SuggestedPost(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("animation.gif")

	// Test with nil params
	result := ctx.SendAnimation(filename).SuggestedPost(nil)
	if result == nil {
		t.Error("SuggestedPost method should return SendAnimation builder for chaining with nil params")
	}

	// Test chaining
	chainedResult := result.SuggestedPost(nil)
	if chainedResult == nil {
		t.Error("SuggestedPost method should support chaining")
	}
}

func TestSendAnimation_ShowCaptionAboveMedia(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("animation.gif")

	result := ctx.SendAnimation(filename).ShowCaptionAboveMedia()
	if result == nil {
		t.Error("ShowCaptionAboveMedia method should return SendAnimation builder for chaining")
	}

	chainedResult := result.ShowCaptionAboveMedia()
	if chainedResult == nil {
		t.Error("ShowCaptionAboveMedia method should support chaining")
	}
}

func TestSendAnimation_Spoiler(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("animation.gif")

	result := ctx.SendAnimation(filename).Spoiler()
	if result == nil {
		t.Error("Spoiler method should return SendAnimation builder for chaining")
	}

	chainedResult := result.Spoiler()
	if chainedResult == nil {
		t.Error("Spoiler method should support chaining")
	}
}

func TestSendAnimation_To(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("animation.gif")

	chatIDs := []int64{123, 456, -1001234567890, 0}

	for _, chatID := range chatIDs {
		result := ctx.SendAnimation(filename).To(chatID)
		if result == nil {
			t.Errorf("To method should return SendAnimation builder for chaining with chatID: %d", chatID)
		}

		chainedResult := result.To(chatID + 100)
		if chainedResult == nil {
			t.Errorf("To method should support chaining and override with chatID: %d", chatID)
		}
	}
}

func TestSendAnimation_ErrorHandling(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// Test with invalid filename that should cause file.Input to fail
	invalidFilename := g.String("") // Empty filename should cause an error
	result := ctx.SendAnimation(invalidFilename)

	// The builder should still be created even with error
	if result == nil {
		t.Error("SendAnimation should return builder even with invalid filename")
	}

	// Test that Send() properly handles the error
	sendResult := result.Send()
	if !sendResult.IsErr() {
		t.Error("Send should fail with empty filename")
	} else {
		t.Logf("Send failed as expected with empty filename: %v", sendResult.Err())
	}
}

func TestSendAnimation_ThumbnailErrorHandling(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("animation.gif")

	// Test with invalid thumbnail file
	result := ctx.SendAnimation(filename).Thumbnail(g.String("/invalid/path/thumb.jpg"))
	if result == nil {
		t.Error("Thumbnail with invalid file should still return builder")
	}

	// Test that Send() handles thumbnail error properly
	sendResult := result.Send()
	if sendResult.IsOk() {
		t.Logf("Send succeeded despite invalid thumbnail: %v", sendResult.Ok())
	} else {
		t.Logf("Send failed as expected with invalid thumbnail: %v", sendResult.Err())
	}
}

func TestSendAnimation_FileClosing(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// Create a temporary file to test file closing
	tempFile := "/tmp/test_animation.gif"
	os.WriteFile(tempFile, []byte("test animation content"), 0644)
	defer os.Remove(tempFile)

	result := ctx.SendAnimation(g.String(tempFile))
	if result == nil {
		t.Error("SendAnimation with valid file should return builder")
	}

	// Call Send to trigger file closing logic
	sendResult := result.Send()
	if sendResult.IsErr() {
		t.Logf("Send failed as expected with mock bot: %v", sendResult.Err())
	}
}

func TestSendAnimation_ThumbnailWithValidFile(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// Create temporary files for animation and thumbnail
	animFile := "/tmp/test_animation.gif"
	thumbFile := "/tmp/test_thumb.jpg"
	os.WriteFile(animFile, []byte("test animation content"), 0644)
	os.WriteFile(thumbFile, []byte("test thumb content"), 0644)
	defer os.Remove(animFile)
	defer os.Remove(thumbFile)

	result := ctx.SendAnimation(g.String(animFile)).Thumbnail(g.String(thumbFile))
	if result == nil {
		t.Error("SendAnimation with valid thumbnail should return builder")
	}

	// Call Send to trigger both file and thumbnail closing logic
	sendResult := result.Send()
	if sendResult.IsErr() {
		t.Logf("Send failed as expected with mock bot: %v", sendResult.Err())
	}
}

func TestSendAnimation_TimeoutWithNilRequestOpts(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("animation.gif")

	// Test Timeout when RequestOpts is nil
	result := ctx.SendAnimation(filename).Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout should return builder")
	}
}

func TestSendAnimation_APIURLWithNilRequestOpts(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("animation.gif")

	// Test APIURL when RequestOpts is nil
	result := ctx.SendAnimation(filename).APIURL(g.String("https://api.example.com"))
	if result == nil {
		t.Error("APIURL should return builder")
	}
}

func TestSendAnimation_APIURLWithExistingRequestOpts(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("animation.gif")

	// First set Timeout to create RequestOpts, then test APIURL
	result := ctx.SendAnimation(filename).
		Timeout(15 * time.Second).                         // This creates RequestOpts
		APIURL(g.String("https://custom.api.example.com")) // This should use existing RequestOpts

	if result == nil {
		t.Error("APIURL with existing RequestOpts should return builder")
	}
}

func TestSendAnimation_SendWithExistingError(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// First create a SendAnimation with an error (invalid filename)
	result := ctx.SendAnimation(g.String("/invalid/nonexistent/animation.gif"))
	if result == nil {
		t.Error("SendAnimation with invalid file should return builder")
	}

	// Test that Send() returns the error immediately without calling timers
	sendResult := result.Send()
	if !sendResult.IsErr() {
		t.Error("Send should return error for invalid file")
	}
}
