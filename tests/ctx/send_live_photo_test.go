package ctx_test

import (
	"errors"
	"io/fs"
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

func TestContext_SendLivePhoto(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Use file_id-style inputs so the builder doesn't attempt to open real local files.
	live := g.String("file_id:live_photo_video_id")
	cover := g.String("file_id:cover_photo_id")

	result := testCtx.SendLivePhoto(live, cover)
	if result == nil {
		t.Error("Expected SendLivePhoto builder to be created")
	}

	chained := result.Caption(g.String("Live photo caption"))
	if chained == nil {
		t.Error("Expected Caption method to return builder")
	}
}

func TestContext_SendLivePhotoChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	live := g.String("file_id:live_video")
	cover := g.String("file_id:cover_photo")

	result := testCtx.SendLivePhoto(live, cover).
		Caption(g.String("Live photo")).
		HTML().
		Silent().
		Spoiler().
		ShowCaptionAboveMedia().
		To(123)
	if result == nil {
		t.Error("Expected SendLivePhoto builder after chaining")
	}

	final := result.Protect()
	if final == nil {
		t.Error("Expected Protect to return SendLivePhoto for chaining")
	}
}

func TestSendLivePhoto_AllMethods(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	live := g.String("file_id:test_live")
	cover := g.String("file_id:test_cover")

	if r := testCtx.SendLivePhoto(live, cover).CaptionEntities(entities.New("Bold").Bold("Bold")); r == nil {
		t.Error("CaptionEntities method should return SendLivePhoto for chaining")
	}

	if r := testCtx.SendLivePhoto(live, cover).After(5 * time.Second); r == nil {
		t.Error("After method should return SendLivePhoto for chaining")
	}

	if r := testCtx.SendLivePhoto(live, cover).DeleteAfter(60 * time.Second); r == nil {
		t.Error("DeleteAfter method should return SendLivePhoto for chaining")
	}

	if r := testCtx.SendLivePhoto(live, cover).Spoiler(); r == nil {
		t.Error("Spoiler method should return SendLivePhoto for chaining")
	}

	if r := testCtx.SendLivePhoto(live, cover).Markdown(); r == nil {
		t.Error("Markdown method should return SendLivePhoto for chaining")
	}

	if r := testCtx.SendLivePhoto(live, cover).AllowPaidBroadcast(); r == nil {
		t.Error("AllowPaidBroadcast method should return SendLivePhoto for chaining")
	}

	if r := testCtx.SendLivePhoto(live, cover).Effect(effects.Fire); r == nil {
		t.Error("Effect method should return SendLivePhoto for chaining")
	}

	if r := testCtx.SendLivePhoto(live, cover).Markup(keyboard.Inline()); r == nil {
		t.Error("Markup method should return SendLivePhoto for chaining")
	}

	if r := testCtx.SendLivePhoto(live, cover).Reply(reply.New(99)); r == nil {
		t.Error("Reply method should return SendLivePhoto for chaining")
	}

	if r := testCtx.SendLivePhoto(live, cover).Timeout(30 * time.Second); r == nil {
		t.Error("Timeout method should return SendLivePhoto for chaining")
	}

	if r := testCtx.SendLivePhoto(live, cover).APIURL(g.String("https://api.example.com")); r == nil {
		t.Error("APIURL method should return SendLivePhoto for chaining")
	}

	if r := testCtx.SendLivePhoto(live, cover).Business(g.String("biz_conn")); r == nil {
		t.Error("Business method should return SendLivePhoto for chaining")
	}

	if r := testCtx.SendLivePhoto(live, cover).Thread(7); r == nil {
		t.Error("Thread method should return SendLivePhoto for chaining")
	}

	if r := testCtx.SendLivePhoto(live, cover).DirectMessagesTopic(42); r == nil {
		t.Error("DirectMessagesTopic method should return SendLivePhoto for chaining")
	}

	// Nil reply should be a no-op but still return a builder.
	if r := testCtx.SendLivePhoto(live, cover).Reply(nil); r == nil {
		t.Error("Reply(nil) should still return SendLivePhoto for chaining")
	}
}

func TestSendLivePhoto_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	live := g.String("file_id:send_live")
	cover := g.String("file_id:send_cover")

	sendResult := testCtx.SendLivePhoto(live, cover).Send()
	if sendResult.IsErr() {
		t.Logf("SendLivePhoto Send failed as expected with mock bot: %v", sendResult.Err())
	}

	configured := testCtx.SendLivePhoto(live, cover).
		Caption(g.String("hi")).
		Markdown().
		Silent().
		To(456).
		Send()
	if configured.IsErr() {
		t.Logf("SendLivePhoto configured Send failed as expected: %v", configured.Err())
	}
}

func TestSendLivePhoto_NonexistentLocalFile(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Missing local files should be surfaced as a Send error, not a panic.
	res := testCtx.SendLivePhoto(g.String("definitely_missing_live.mp4"), g.String("definitely_missing_photo.jpg")).Send()
	if res.IsOk() {
		t.Error("Expected error for missing local files")
	} else if !errors.Is(res.Err(), fs.ErrNotExist) {
		t.Logf("Expected fs.ErrNotExist, got: %v", res.Err())
	}
}
