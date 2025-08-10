package poll_test

import (
	"testing"

	"github.com/enetx/tg/types/poll"
)

func TestPollType_String(t *testing.T) {
	tests := []struct {
		name     string
		pollType poll.PollType
		expected string
	}{
		{
			name:     "Regular poll",
			pollType: poll.Regular,
			expected: "regular",
		},
		{
			name:     "Quiz poll",
			pollType: poll.Quiz,
			expected: "quiz",
		},
		{
			name:     "Unknown poll type",
			pollType: poll.PollType(999),
			expected: "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.pollType.String()
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestPollType_Constants(t *testing.T) {
	if poll.Regular != 0 {
		t.Errorf("Expected Regular to be 0, got %d", int(poll.Regular))
	}
	if poll.Quiz != 1 {
		t.Errorf("Expected Quiz to be 1, got %d", int(poll.Quiz))
	}
}
