package chat_test

import (
	"testing"

	"github.com/enetx/tg/types/chat"
)

func TestChatType_String(t *testing.T) {
	tests := []struct {
		name     string
		chatType chat.ChatType
		expected string
	}{
		{
			name:     "Private chat",
			chatType: chat.Private,
			expected: "private",
		},
		{
			name:     "Group chat",
			chatType: chat.Group,
			expected: "group",
		},
		{
			name:     "Supergroup chat",
			chatType: chat.Supergroup,
			expected: "supergroup",
		},
		{
			name:     "Channel chat",
			chatType: chat.Channel,
			expected: "channel",
		},
		{
			name:     "Sender chat",
			chatType: chat.Sender,
			expected: "sender",
		},
		{
			name:     "Unknown chat type",
			chatType: chat.ChatType(999),
			expected: "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.chatType.String()
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestChatType_Constants(t *testing.T) {
	// Test that constants have expected values for ordering
	if chat.Private != 0 {
		t.Errorf("Expected Private to be 0, got %d", int(chat.Private))
	}
	if chat.Group != 1 {
		t.Errorf("Expected Group to be 1, got %d", int(chat.Group))
	}
	if chat.Supergroup != 2 {
		t.Errorf("Expected Supergroup to be 2, got %d", int(chat.Supergroup))
	}
	if chat.Channel != 3 {
		t.Errorf("Expected Channel to be 3, got %d", int(chat.Channel))
	}
	if chat.Sender != 4 {
		t.Errorf("Expected Sender to be 4, got %d", int(chat.Sender))
	}
}
