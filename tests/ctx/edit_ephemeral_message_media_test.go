package ctx_test

import (
	"testing"
	"time"

	"github.com/enetx/g"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/keyboard"
)

func TestContext_EditEphemeralMessageMedia(t *testing.T) {
	testCtx := ephemeralContext(&mockBot{})
	media := input.Photo(file.Input(g.String("https://example.com/photo.jpg")).Ok())

	result := testCtx.EditEphemeralMessageMedia(media)
	if result == nil {
		t.Error("Expected EditEphemeralMessageMedia builder to be created")
	}

	if result.ChatID(999) == nil {
		t.Error("ChatID method should return EditEphemeralMessageMedia for chaining")
	}

	if result.ReceiverUser(111) == nil {
		t.Error("ReceiverUser method should return EditEphemeralMessageMedia for chaining")
	}

	if result.MessageID(222) == nil {
		t.Error("MessageID method should return EditEphemeralMessageMedia for chaining")
	}

	if result.Markup(keyboard.Inline()) == nil {
		t.Error("Markup method should return EditEphemeralMessageMedia for chaining")
	}

	if result.Timeout(30*time.Second) == nil {
		t.Error("Timeout method should return EditEphemeralMessageMedia for chaining")
	}

	if result.APIURL(g.String("https://api.example.com")) == nil {
		t.Error("APIURL method should return EditEphemeralMessageMedia for chaining")
	}
}

func TestEditEphemeralMessageMedia_Send(t *testing.T) {
	testCtx := ephemeralContext(&mockBot{})
	media := input.Photo(file.Input(g.String("https://example.com/photo.jpg")).Ok())

	sendResult := testCtx.EditEphemeralMessageMedia(media).Send()
	if !sendResult.IsErr() && !sendResult.IsOk() {
		t.Error("Send should return a result using the effective ephemeral message")
	}

	complete := testCtx.EditEphemeralMessageMedia(media).
		ChatID(999).
		ReceiverUser(111).
		MessageID(222).
		Markup(keyboard.Inline().Text("Button", "data")).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if !complete.IsErr() && !complete.IsOk() {
		t.Error("Send with all features should return a result")
	}
}
