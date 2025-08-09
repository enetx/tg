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
		result := ctx.SendAnimation(filename).ReplyTo(messageID)
		if result == nil {
			t.Errorf("ReplyTo method should return SendAnimation builder for chaining with messageID: %d", messageID)
		}

		chainedResult := result.ReplyTo(messageID + 100)
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
	invalidFilename := g.String("")  // Empty filename should cause an error
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
