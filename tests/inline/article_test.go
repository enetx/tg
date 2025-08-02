package inline_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/inline"
)

func TestNewArticle(t *testing.T) {
	msg := createTestMessageContent()
	article := inline.NewArticle(testID, testTitle, msg)
	if article == nil {
		t.Error("Expected Article to be created")
	}

	built := article.Build()
	if built == nil {
		t.Error("Expected Article to build correctly")
	}

	if built.GetType() != "article" {
		t.Error("Expected Article type to be 'article'")
	}
}

func TestArticle_URL(t *testing.T) {
	url := g.String("https://example.com")
	msg := createTestMessageContent()
	result := inline.NewArticle(testID, testTitle, msg).URL(url)
	if result == nil {
		t.Error("Expected URL method to return Article")
	}

	built := result.Build()
	if a, ok := built.(gotgbot.InlineQueryResultArticle); ok {
		if a.Url != url.Std() {
			t.Error("Expected Url to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultArticle")
	}
}

func TestArticle_Description(t *testing.T) {
	desc := g.String("description")
	msg := createTestMessageContent()
	result := inline.NewArticle(testID, testTitle, msg).Description(desc)
	if result == nil {
		t.Error("Expected Description method to return Article")
	}

	built := result.Build()
	if a, ok := built.(gotgbot.InlineQueryResultArticle); ok {
		if a.Description != desc.Std() {
			t.Error("Expected Description to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultArticle")
	}
}

func TestArticle_Thumbnail(t *testing.T) {
	thumb := g.String("https://example.com/thumb.jpg")
	msg := createTestMessageContent()
	result := inline.NewArticle(testID, testTitle, msg).ThumbnailURL(thumb).ThumbnailSize(100, 200)
	if result == nil {
		t.Error("Expected Thumbnail methods to return Article")
	}

	built := result.Build()
	if a, ok := built.(gotgbot.InlineQueryResultArticle); ok {
		if a.ThumbnailUrl != thumb.Std() {
			t.Error("Expected ThumbnailUrl to be set correctly")
		}
		if a.ThumbnailWidth != 100 || a.ThumbnailHeight != 200 {
			t.Error("Expected Thumbnail size to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultArticle")
	}
}

func TestArticle_Markup(t *testing.T) {
	msg := createTestMessageContent()
	kb := createTestKeyboard()
	result := inline.NewArticle(testID, testTitle, msg).Markup(kb)
	if result == nil {
		t.Error("Expected Markup method to return Article")
	}

	built := result.Build()
	if a, ok := built.(gotgbot.InlineQueryResultArticle); ok {
		if a.ReplyMarkup == nil {
			t.Error("Expected ReplyMarkup to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultArticle")
	}
}

func TestArticle_MethodChaining(t *testing.T) {
	msg := createTestMessageContent()
	url := g.String("https://example.com")
	desc := g.String("description")
	thumb := g.String("https://example.com/thumb.jpg")
	kb := createTestKeyboard()

	result := inline.NewArticle(testID, testTitle, msg).
		URL(url).
		Description(desc).
		ThumbnailURL(thumb).
		ThumbnailSize(100, 200).
		Markup(kb)

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained Article to build correctly")
	}

	if !assertQueryResult(result) {
		t.Error("Expected result to implement QueryResult interface")
	}
}
