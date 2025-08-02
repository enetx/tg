package inline_test

import (
	"testing"

	"github.com/enetx/tg/inline"
)

func TestNewArticle(t *testing.T) {
	messageContent := createTestMessageContent()
	article := inline.NewArticle(testID, testTitle, messageContent)

	if article == nil {
		t.Error("Expected Article to be created")
	}

	// Test that it builds correctly
	built := article.Build()
	if built == nil {
		t.Error("Expected Article to build correctly")
	}

	// Test that it implements QueryResult interface
	if !assertQueryResult(article) {
		t.Error("Article should implement QueryResult correctly")
	}
}

func TestArticle_URL(t *testing.T) {
	messageContent := createTestMessageContent()
	article := inline.NewArticle(testID, testTitle, messageContent)

	result := article.URL(testURL)
	if result == nil {
		t.Error("Expected URL method to return Article")
	}

	// Test method chaining works
	if result != article {
		t.Error("Expected URL method to return same Article instance")
	}
}

func TestArticle_Description(t *testing.T) {
	messageContent := createTestMessageContent()
	article := inline.NewArticle(testID, testTitle, messageContent)

	result := article.Description(testDescription)
	if result == nil {
		t.Error("Expected Description method to return Article")
	}

	// Test method chaining works
	if result != article {
		t.Error("Expected Description method to return same Article instance")
	}
}

func TestArticle_ThumbnailURL(t *testing.T) {
	messageContent := createTestMessageContent()
	article := inline.NewArticle(testID, testTitle, messageContent)

	result := article.ThumbnailURL(testThumbnailURL)
	if result == nil {
		t.Error("Expected ThumbnailURL method to return Article")
	}

	// Test method chaining works
	if result != article {
		t.Error("Expected ThumbnailURL method to return same Article instance")
	}
}

func TestArticle_ThumbnailSize(t *testing.T) {
	messageContent := createTestMessageContent()
	article := inline.NewArticle(testID, testTitle, messageContent)

	width := int64(100)
	height := int64(100)
	result := article.ThumbnailSize(width, height)
	if result == nil {
		t.Error("Expected ThumbnailSize method to return Article")
	}

	// Test method chaining works
	if result != article {
		t.Error("Expected ThumbnailSize method to return same Article instance")
	}
}

func TestArticle_MethodChaining(t *testing.T) {
	messageContent := createTestMessageContent()

	result := inline.NewArticle(testID, testTitle, messageContent).
		URL(testURL).
		Description(testDescription).
		ThumbnailURL(testThumbnailURL).
		ThumbnailSize(100, 100)

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained Article to build correctly")
	}
}
