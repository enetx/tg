package rich_test

import (
	"github.com/enetx/g"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/richtext"
)

// Common test data
var (
	testURL     = g.String("https://example.com/media.jpg")
	testCaption = g.String("Test caption")
	testCredit  = g.String("Test credit")
)

// Helper functions to build common builders used across the rich tests.
func testText(text g.String) *richtext.Text { return richtext.New(text) }

func testPhoto() *input.MediaPhoto { return input.Photo(file.Input(testURL).Ok()) }

func testVideo() *input.MediaVideo { return input.Video(file.Input(testURL).Ok()) }

func testAudio() *input.MediaAudio { return input.Audio(file.Input(testURL).Ok()) }

func testAnimation() *input.MediaAnimation { return input.Animation(file.Input(testURL).Ok()) }

func testVoiceNote() *input.MediaVoiceNote { return input.VoiceNote(file.Input(testURL).Ok()) }
