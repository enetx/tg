package rich_test

import (
	"testing"

	"github.com/enetx/tg/rich"
)

func TestListLabel_String(t *testing.T) {
	tests := []struct {
		label rich.ListLabel
		want  string
	}{
		{rich.LabelNone, ""},
		{rich.LabelLowerAlpha, "a"},
		{rich.LabelUpperAlpha, "A"},
		{rich.LabelLowerRoman, "i"},
		{rich.LabelUpperRoman, "I"},
		{rich.LabelDecimal, "1"},
		{rich.ListLabel(99), ""},
	}

	for _, tt := range tests {
		if got := tt.label.String(); got != tt.want {
			t.Errorf("Expected %q, got %q", tt.want, got)
		}
	}
}
