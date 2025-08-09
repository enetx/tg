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

func TestContext_SendAudio(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("audio.mp3")

	result := ctx.SendAudio(filename)

	if result == nil {
		t.Error("Expected SendAudio builder to be created")
	}

	// Test method chaining
	chained := result.Caption(g.String("Audio caption"))
	if chained == nil {
		t.Error("Expected caption method to return builder")
	}
}

func TestContext_SendAudioChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("audio.mp3")

	result := ctx.SendAudio(filename).
		Caption(g.String("Test audio")).
		HTML().
		Silent().
		Duration(180)

	if result == nil {
		t.Error("Expected SendAudio builder to be created")
	}

	// Test continued chaining
	final := result.To(123)
	if final == nil {
		t.Error("Expected To method to return builder")
	}
}

func TestSendAudio_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("test_audio.mp3")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendAudio(filename).Send()

	if sendResult.IsErr() {
		t.Logf("SendAudio Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendAudio(filename).
		Caption(g.String("Test <b>audio</b> with HTML")).
		HTML().
		Duration(180).
		Performer(g.String("Test Artist")).
		Title(g.String("Test Song")).
		Silent().
		Protect().
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendAudio configured Send failed as expected: %v", configuredSendResult.Err())
	}
}

func TestSendAudio_CaptionEntities(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("audio.mp3")

	ent := entities.New(g.String("Bold text")).Bold(g.String("Bold"))
	result := ctx.SendAudio(filename).CaptionEntities(ent)
	if result == nil {
		t.Error("CaptionEntities method should return SendAudio builder for chaining")
	}

	entItalic := entities.New(g.String("Italic text")).Italic(g.String("Italic"))
	chainedResult := result.CaptionEntities(entItalic)
	if chainedResult == nil {
		t.Error("CaptionEntities method should support chaining and override")
	}
}

func TestSendAudio_After(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("audio.mp3")

	durations := []time.Duration{
		time.Second,
		time.Minute,
		time.Hour,
		0,
	}

	for _, duration := range durations {
		result := ctx.SendAudio(filename).After(duration)
		if result == nil {
			t.Errorf("After method should return SendAudio builder for chaining with duration: %v", duration)
		}

		chainedResult := result.After(time.Second * 30)
		if chainedResult == nil {
			t.Errorf("After method should support chaining and override with duration: %v", duration)
		}
	}
}

func TestSendAudio_DeleteAfter(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("audio.mp3")

	durations := []time.Duration{
		time.Second * 30,
		time.Minute * 5,
		time.Hour,
		0,
	}

	for _, duration := range durations {
		result := ctx.SendAudio(filename).DeleteAfter(duration)
		if result == nil {
			t.Errorf("DeleteAfter method should return SendAudio builder for chaining with duration: %v", duration)
		}

		chainedResult := result.DeleteAfter(time.Minute * 10)
		if chainedResult == nil {
			t.Errorf("DeleteAfter method should support chaining and override with duration: %v", duration)
		}
	}
}

func TestSendAudio_Markdown(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("audio.mp3")

	result := ctx.SendAudio(filename).Markdown()
	if result == nil {
		t.Error("Markdown method should return SendAudio builder for chaining")
	}

	chainedResult := result.Markdown()
	if chainedResult == nil {
		t.Error("Markdown method should support chaining")
	}
}

func TestSendAudio_Markup(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("audio.mp3")

	btn1 := keyboard.NewButton().Text(g.String("Test Button")).Callback(g.String("test_data"))
	inlineKeyboard := keyboard.Inline().Button(btn1)

	result := ctx.SendAudio(filename).Markup(inlineKeyboard)
	if result == nil {
		t.Error("Markup method should return SendAudio builder for chaining with inline keyboard")
	}

	replyKeyboard := keyboard.Reply()
	chainedResult := result.Markup(replyKeyboard)
	if chainedResult == nil {
		t.Error("Markup method should support chaining and override with reply keyboard")
	}
}

func TestSendAudio_Thumbnail(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("audio.mp3")

	thumbnails := []string{
		"thumb.jpg",
		"album_cover.png",
		"preview.webp",
	}

	for _, thumb := range thumbnails {
		result := ctx.SendAudio(filename).Thumbnail(g.String(thumb))
		if result == nil {
			t.Errorf("Thumbnail method should return SendAudio builder for chaining with thumbnail: %s", thumb)
		}

		chainedResult := result.Thumbnail(g.String("override.jpg"))
		if chainedResult == nil {
			t.Errorf("Thumbnail method should support chaining and override with thumbnail: %s", thumb)
		}
	}
}

func TestSendAudio_ReplyTo(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("audio.mp3")

	messageIDs := []int64{1, 123, 456, 999}

	for _, messageID := range messageIDs {
		result := ctx.SendAudio(filename).ReplyTo(messageID)
		if result == nil {
			t.Errorf("ReplyTo method should return SendAudio builder for chaining with messageID: %d", messageID)
		}

		chainedResult := result.ReplyTo(messageID + 100)
		if chainedResult == nil {
			t.Errorf("ReplyTo method should support chaining and override with messageID: %d", messageID)
		}
	}
}

func TestSendAudio_Business(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("audio.mp3")

	businessIDs := []string{
		"business_123",
		"conn_456",
		"",
	}

	for _, businessID := range businessIDs {
		result := ctx.SendAudio(filename).Business(g.String(businessID))
		if result == nil {
			t.Errorf("Business method should return SendAudio builder for chaining with businessID: %s", businessID)
		}

		chainedResult := result.Business(g.String("override_business"))
		if chainedResult == nil {
			t.Errorf("Business method should support chaining and override with businessID: %s", businessID)
		}
	}
}

func TestSendAudio_Thread(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("audio.mp3")

	threadIDs := []int64{1, 123, 456, 0}

	for _, threadID := range threadIDs {
		result := ctx.SendAudio(filename).Thread(threadID)
		if result == nil {
			t.Errorf("Thread method should return SendAudio builder for chaining with threadID: %d", threadID)
		}

		chainedResult := result.Thread(threadID + 100)
		if chainedResult == nil {
			t.Errorf("Thread method should support chaining and override with threadID: %d", threadID)
		}
	}
}
