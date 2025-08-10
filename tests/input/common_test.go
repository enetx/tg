package input_test

import (
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/input"
)

// Common test data
var (
	testText         = g.String("Test message text")
	testURL          = g.String("https://example.com/test.jpg")
	testThumbnailURL = g.String("https://example.com/thumbnail.jpg")
	testFileID       = g.String("BAADBAADrwADBREAAYdaWKOKDj8X")
	testCaption      = g.String("Test caption")
	testLatitude     = 40.7128
	testLongitude    = -74.0060
	testTitle        = g.String("Test Title")
	testAddress      = g.String("New York, NY")
	testPhoneNumber  = g.String("+1234567890")
	testFirstName    = g.String("John")
	testLastName     = g.String("Doe")
)

// Helper functions to test interface implementations
func assertMedia(m input.Media) bool {
	return m != nil && m.Build() != nil
}

func assertMessageContent(mc input.MessageContent) bool {
	return mc != nil && mc.Build() != nil
}

func assertPaidMedia(pm input.PaidMedia) bool {
	return pm != nil && pm.Build() != nil
}

func assertProfilePhoto(pp input.ProfilePhoto) bool {
	return pp != nil && pp.Build() != nil
}

func assertStoryContent(sc input.StoryContent) bool {
	return sc != nil && sc.Build() != nil
}

func assertPollOption(po input.PollOption) bool {
	if po == nil {
		return false
	}
	built := po.Build()
	// Check if the built result is not empty (text should be set)
	return built.Text != ""
}

func createTestEntities() entities.Entities {
	return *entities.New(g.String("Test text with bold")).Bold("bold")
}
