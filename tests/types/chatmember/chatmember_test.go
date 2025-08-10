package chatmember_test

import (
	"testing"

	"github.com/enetx/tg/types/chatmember"
)

func TestChatMemberStatus_String(t *testing.T) {
	tests := []struct {
		name     string
		status   chatmember.ChatMemberStatus
		expected string
	}{
		{"Creator", chatmember.Creator, "creator"},
		{"Administrator", chatmember.Administrator, "administrator"},
		{"Member", chatmember.Member, "member"},
		{"Restricted", chatmember.Restricted, "restricted"},
		{"Left", chatmember.Left, "left"},
		{"Kicked", chatmember.Kicked, "kicked"},
		{"Unknown", chatmember.ChatMemberStatus(999), "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.status.String()
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}
