package inline_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/inline"
)

func TestNewContact(t *testing.T) {
	contact := inline.NewContact(testID, g.String("+1234567890"), g.String("John"))

	if contact == nil {
		t.Error("Expected Contact to be created")
	}

	built := contact.Build()
	if built == nil {
		t.Error("Expected Contact to build correctly")
	}

	if result, ok := built.(gotgbot.InlineQueryResultContact); ok {
		if result.GetType() != "contact" {
			t.Error("Expected type to be 'contact'")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultContact")
	}
}

func TestContact_LastName(t *testing.T) {
	contact := inline.NewContact(testID, g.String("+1234567890"), g.String("John"))

	result := contact.LastName(g.String("Doe"))
	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultContact); ok {
		if v.LastName != "Doe" {
			t.Error("Expected LastName to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultContact")
	}
}

func TestContact_VCard(t *testing.T) {
	vcard := g.String("BEGIN:VCARD\nVERSION:3.0\nFN:John Doe\nEND:VCARD")
	contact := inline.NewContact(testID, g.String("+1234567890"), g.String("John"))

	result := contact.VCard(vcard)
	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultContact); ok {
		if v.Vcard != vcard.Std() {
			t.Error("Expected VCard to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultContact")
	}
}

func TestContact_Markup(t *testing.T) {
	contact := inline.NewContact(testID, g.String("+1234567890"), g.String("John"))
	keyboard := createTestKeyboard()

	result := contact.Markup(keyboard)
	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultContact); ok {
		if v.ReplyMarkup == nil {
			t.Error("Expected ReplyMarkup to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultContact")
	}
}

func TestContact_ThumbnailURL(t *testing.T) {
	contact := inline.NewContact(testID, g.String("+1234567890"), g.String("John"))

	result := contact.ThumbnailURL(testThumbnailURL)
	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultContact); ok {
		if v.ThumbnailUrl != testThumbnailURL.Std() {
			t.Error("Expected ThumbnailUrl to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultContact")
	}
}

func TestContact_ThumbnailSize(t *testing.T) {
	contact := inline.NewContact(testID, g.String("+1234567890"), g.String("John"))

	result := contact.ThumbnailSize(100, 100)
	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultContact); ok {
		if v.ThumbnailWidth != 100 {
			t.Error("Expected ThumbnailWidth to be set correctly")
		}
		if v.ThumbnailHeight != 100 {
			t.Error("Expected ThumbnailHeight to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultContact")
	}
}

func TestContact_MethodChaining(t *testing.T) {
	result := inline.NewContact(testID, g.String("+1234567890"), g.String("John")).
		LastName(g.String("Doe")).
		VCard(g.String("BEGIN:VCARD\nVERSION:3.0\nFN:John Doe\nEND:VCARD")).
		ThumbnailURL(testThumbnailURL).
		ThumbnailSize(100, 100).
		Markup(createTestKeyboard())

	built := result.Build()
	if built == nil {
		t.Error("Expected chained Contact to build correctly")
	}

	if _, ok := built.(gotgbot.InlineQueryResultContact); !ok {
		t.Error("Expected result to be InlineQueryResultContact")
	}

	if !assertQueryResult(result) {
		t.Error("Expected result to implement QueryResult interface")
	}
}
