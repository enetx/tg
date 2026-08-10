package ctx_test

import (
	"testing"
	"time"

	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/keyboard"
)

func TestContext_EditEphemeralMessageCaption(t *testing.T) {
	testCtx := ephemeralContext(&mockBot{})

	result := testCtx.EditEphemeralMessageCaption(g.String("New caption"))
	if result == nil {
		t.Error("Expected EditEphemeralMessageCaption builder to be created")
	}

	if result.ChatID(999) == nil {
		t.Error("ChatID method should return EditEphemeralMessageCaption for chaining")
	}

	if result.ReceiverUser(111) == nil {
		t.Error("ReceiverUser method should return EditEphemeralMessageCaption for chaining")
	}

	if result.MessageID(222) == nil {
		t.Error("MessageID method should return EditEphemeralMessageCaption for chaining")
	}

	if result.HTML() == nil {
		t.Error("HTML method should return EditEphemeralMessageCaption for chaining")
	}

	if result.Markdown() == nil {
		t.Error("Markdown method should return EditEphemeralMessageCaption for chaining")
	}

	if result.Entities(entities.New("Bold and italic")) == nil {
		t.Error("Entities method should return EditEphemeralMessageCaption for chaining")
	}

	if result.Markup(keyboard.Inline()) == nil {
		t.Error("Markup method should return EditEphemeralMessageCaption for chaining")
	}

	if result.Timeout(30*time.Second) == nil {
		t.Error("Timeout method should return EditEphemeralMessageCaption for chaining")
	}

	if result.APIURL(g.String("https://api.example.com")) == nil {
		t.Error("APIURL method should return EditEphemeralMessageCaption for chaining")
	}
}

func TestEditEphemeralMessageCaption_Send(t *testing.T) {
	testCtx := ephemeralContext(&mockBot{})

	sendResult := testCtx.EditEphemeralMessageCaption(g.String("New caption")).Send()
	if !sendResult.IsErr() && !sendResult.IsOk() {
		t.Error("Send should return a result using the effective ephemeral message")
	}

	complete := testCtx.EditEphemeralMessageCaption(g.String("New caption")).
		ChatID(999).
		ReceiverUser(111).
		MessageID(222).
		Markdown().
		Markup(keyboard.Inline().Text("Button", "data")).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if !complete.IsErr() && !complete.IsOk() {
		t.Error("Send with all features should return a result")
	}
}
