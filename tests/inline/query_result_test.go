package inline_test

import (
	"testing"

	"github.com/enetx/tg/inline"
)

func TestQueryResult_Interface(t *testing.T) {
	// Test that all concrete types implement QueryResult interface
	var _ inline.QueryResult = &inline.Article{}
	var _ inline.QueryResult = &inline.Audio{}
	var _ inline.QueryResult = &inline.CachedAudio{}
	var _ inline.QueryResult = &inline.CachedDocument{}
	var _ inline.QueryResult = &inline.CachedGif{}
	var _ inline.QueryResult = &inline.CachedMpeg4Gif{}
	var _ inline.QueryResult = &inline.CachedPhoto{}
	var _ inline.QueryResult = &inline.CachedSticker{}
	var _ inline.QueryResult = &inline.CachedVideo{}
	var _ inline.QueryResult = &inline.CachedVoice{}
	var _ inline.QueryResult = &inline.Contact{}
	var _ inline.QueryResult = &inline.Document{}
	var _ inline.QueryResult = &inline.Game{}
	var _ inline.QueryResult = &inline.Gif{}
	var _ inline.QueryResult = &inline.Location{}
	var _ inline.QueryResult = &inline.Mpeg4Gif{}
	var _ inline.QueryResult = &inline.Photo{}
	var _ inline.QueryResult = &inline.Venue{}
	var _ inline.QueryResult = &inline.Video{}
	var _ inline.QueryResult = &inline.Voice{}
}

func TestQueryResult_CommonBehavior(t *testing.T) {
	// Test that QueryResult implementations behave correctly
	messageContent := createTestMessageContent()

	// Test Article result
	article := inline.NewArticle(testID, testTitle, messageContent)
	if !assertQueryResult(article) {
		t.Error("Article should implement QueryResult correctly")
	}

	// Test Photo result
	photo := inline.NewPhoto(testID, testURL, testThumbnailURL)
	if !assertQueryResult(photo) {
		t.Error("Photo should implement QueryResult correctly")
	}

	// Test Contact result
	contact := inline.NewContact(testID, testTitle, testTitle)
	if !assertQueryResult(contact) {
		t.Error("Contact should implement QueryResult correctly")
	}
}
