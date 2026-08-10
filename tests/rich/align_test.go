package rich_test

import (
	"testing"

	"github.com/enetx/tg/rich"
)

func TestCellAlign_String(t *testing.T) {
	tests := []struct {
		align rich.CellAlign
		want  string
	}{
		{rich.AlignLeft, "left"},
		{rich.AlignCenter, "center"},
		{rich.AlignRight, "right"},
		{rich.CellAlign(99), "left"},
	}

	for _, tt := range tests {
		if got := tt.align.String(); got != tt.want {
			t.Errorf("Expected %q, got %q", tt.want, got)
		}
	}
}

func TestCellValign_String(t *testing.T) {
	tests := []struct {
		valign rich.CellValign
		want   string
	}{
		{rich.ValignTop, "top"},
		{rich.ValignMiddle, "middle"},
		{rich.ValignBottom, "bottom"},
		{rich.CellValign(99), "top"},
	}

	for _, tt := range tests {
		if got := tt.valign.String(); got != tt.want {
			t.Errorf("Expected %q, got %q", tt.want, got)
		}
	}
}
