package rich_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/rich"
)

func TestNewCaption(t *testing.T) {
	caption := rich.NewCaption(testText(testCaption))
	if caption == nil {
		t.Fatal("Expected NewCaption to create a Caption builder")
	}

	built := caption.Std()
	if built == nil {
		t.Fatal("Expected Std to return the caption")
	}

	if string(built.Text.(gotgbot.RichTextString)) != testCaption.Std() {
		t.Error("Expected Text to be set correctly")
	}
}

func TestCaption_Text(t *testing.T) {
	caption := rich.NewCaption(testText("old"))

	result := caption.Text(testText(testCaption))
	if result != caption {
		t.Error("Expected Text to return same Caption instance")
	}

	if string(result.Std().Text.(gotgbot.RichTextString)) != testCaption.Std() {
		t.Error("Expected Text to be set correctly")
	}
}

func TestCaption_Credit(t *testing.T) {
	caption := rich.NewCaption(testText(testCaption))

	result := caption.Credit(testText(testCredit))
	if result != caption {
		t.Error("Expected Credit to return same Caption instance")
	}

	if string(result.Std().Credit.(gotgbot.RichTextString)) != testCredit.Std() {
		t.Error("Expected Credit to be set correctly")
	}
}

func TestCaption_StdNil(t *testing.T) {
	var caption *rich.Caption
	if caption.Std() != nil {
		t.Error("Expected Std on a nil Caption to return nil")
	}
}
