package chataction_test

import (
	"testing"

	"github.com/enetx/tg/types/chataction"
)

func TestChatAction_String(t *testing.T) {
	tests := []struct {
		name     string
		action   chataction.ChatAction
		expected string
	}{
		{
			name:     "Typing action",
			action:   chataction.Typing,
			expected: "typing",
		},
		{
			name:     "UploadPhoto action",
			action:   chataction.UploadPhoto,
			expected: "upload_photo",
		},
		{
			name:     "RecordVideo action",
			action:   chataction.RecordVideo,
			expected: "record_video",
		},
		{
			name:     "UploadVideo action",
			action:   chataction.UploadVideo,
			expected: "upload_video",
		},
		{
			name:     "RecordVoice action",
			action:   chataction.RecordVoice,
			expected: "record_voice",
		},
		{
			name:     "UploadVoice action",
			action:   chataction.UploadVoice,
			expected: "upload_voice",
		},
		{
			name:     "UploadDocument action",
			action:   chataction.UploadDocument,
			expected: "upload_document",
		},
		{
			name:     "ChooseSticker action",
			action:   chataction.ChooseSticker,
			expected: "choose_sticker",
		},
		{
			name:     "FindLocation action",
			action:   chataction.FindLocation,
			expected: "find_location",
		},
		{
			name:     "RecordVideoNote action",
			action:   chataction.RecordVideoNote,
			expected: "record_video_note",
		},
		{
			name:     "UploadVideoNote action",
			action:   chataction.UploadVideoNote,
			expected: "upload_video_note",
		},
		{
			name:     "Unknown action",
			action:   chataction.ChatAction(999),
			expected: "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.action.String()
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestChatAction_AllActionsHaveUniqueStrings(t *testing.T) {
	actions := []chataction.ChatAction{
		chataction.Typing,
		chataction.UploadPhoto,
		chataction.RecordVideo,
		chataction.UploadVideo,
		chataction.RecordVoice,
		chataction.UploadVoice,
		chataction.UploadDocument,
		chataction.ChooseSticker,
		chataction.FindLocation,
		chataction.RecordVideoNote,
		chataction.UploadVideoNote,
	}

	seen := make(map[string]bool)

	for _, action := range actions {
		str := action.String()
		if seen[str] {
			t.Errorf("Duplicate action string found: %s", str)
		}
		seen[str] = true
	}
}

func TestChatAction_ValidActionsNotUnknown(t *testing.T) {
	actions := []chataction.ChatAction{
		chataction.Typing,
		chataction.UploadPhoto,
		chataction.RecordVideo,
		chataction.UploadVideo,
		chataction.RecordVoice,
		chataction.UploadVoice,
		chataction.UploadDocument,
		chataction.ChooseSticker,
		chataction.FindLocation,
		chataction.RecordVideoNote,
		chataction.UploadVideoNote,
	}

	for _, action := range actions {
		str := action.String()
		if str == "unknown" {
			t.Errorf("Valid action %d should not return 'unknown'", action)
		}
		if str == "" {
			t.Errorf("Valid action %d should not return empty string", action)
		}
	}
}

func TestChatAction_Constants(t *testing.T) {
	// Test that constants have expected values
	if chataction.Typing != 0 {
		t.Errorf("Expected Typing to be 0, got %d", chataction.Typing)
	}
	if chataction.UploadPhoto != 1 {
		t.Errorf("Expected UploadPhoto to be 1, got %d", chataction.UploadPhoto)
	}
	if chataction.RecordVideo != 2 {
		t.Errorf("Expected RecordVideo to be 2, got %d", chataction.RecordVideo)
	}
	if chataction.UploadVideo != 3 {
		t.Errorf("Expected UploadVideo to be 3, got %d", chataction.UploadVideo)
	}
	if chataction.RecordVoice != 4 {
		t.Errorf("Expected RecordVoice to be 4, got %d", chataction.RecordVoice)
	}
	if chataction.UploadVoice != 5 {
		t.Errorf("Expected UploadVoice to be 5, got %d", chataction.UploadVoice)
	}
	if chataction.UploadDocument != 6 {
		t.Errorf("Expected UploadDocument to be 6, got %d", chataction.UploadDocument)
	}
	if chataction.ChooseSticker != 7 {
		t.Errorf("Expected ChooseSticker to be 7, got %d", chataction.ChooseSticker)
	}
	if chataction.FindLocation != 8 {
		t.Errorf("Expected FindLocation to be 8, got %d", chataction.FindLocation)
	}
	if chataction.RecordVideoNote != 9 {
		t.Errorf("Expected RecordVideoNote to be 9, got %d", chataction.RecordVideoNote)
	}
	if chataction.UploadVideoNote != 10 {
		t.Errorf("Expected UploadVideoNote to be 10, got %d", chataction.UploadVideoNote)
	}
}
