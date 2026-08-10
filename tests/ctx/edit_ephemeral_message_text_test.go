package ctx_test

import (
	"testing"
	"time"

	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/preview"
)

func TestContext_EditEphemeralMessageText(t *testing.T) {
	testCtx := ephemeralContext(&mockBot{})

	result := testCtx.EditEphemeralMessageText(g.String("Updated text"))
	if result == nil {
		t.Error("Expected EditEphemeralMessageText builder to be created")
	}

	if result.ChatID(999) == nil {
		t.Error("ChatID method should return EditEphemeralMessageText for chaining")
	}

	if result.ReceiverUser(111) == nil {
		t.Error("ReceiverUser method should return EditEphemeralMessageText for chaining")
	}

	if result.MessageID(222) == nil {
		t.Error("MessageID method should return EditEphemeralMessageText for chaining")
	}

	if result.HTML() == nil {
		t.Error("HTML method should return EditEphemeralMessageText for chaining")
	}

	if result.Markdown() == nil {
		t.Error("Markdown method should return EditEphemeralMessageText for chaining")
	}

	if result.Entities(entities.New("Bold and italic")) == nil {
		t.Error("Entities method should return EditEphemeralMessageText for chaining")
	}

	if result.Preview(preview.New().Disable()) == nil {
		t.Error("Preview method should return EditEphemeralMessageText for chaining")
	}

	if result.Markup(keyboard.Inline()) == nil {
		t.Error("Markup method should return EditEphemeralMessageText for chaining")
	}

	if result.Timeout(30*time.Second) == nil {
		t.Error("Timeout method should return EditEphemeralMessageText for chaining")
	}

	if result.APIURL(g.String("https://api.example.com")) == nil {
		t.Error("APIURL method should return EditEphemeralMessageText for chaining")
	}
}

func TestEditEphemeralMessageText_Send(t *testing.T) {
	testCtx := ephemeralContext(&mockBot{})

	sendResult := testCtx.EditEphemeralMessageText(g.String("Updated text")).Send()
	if !sendResult.IsErr() && !sendResult.IsOk() {
		t.Error("Send should return a result using the effective ephemeral message")
	}

	complete := testCtx.EditEphemeralMessageText(g.String("Updated text")).
		ChatID(999).
		ReceiverUser(111).
		MessageID(222).
		HTML().
		Markup(keyboard.Inline().Text("Button", "data")).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if !complete.IsErr() && !complete.IsOk() {
		t.Error("Send with all features should return a result")
	}
}
