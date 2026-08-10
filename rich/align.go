package rich

// CellAlign represents the horizontal alignment of a table cell content.
//
// See: https://core.telegram.org/bots/api#richblocktablecell
type CellAlign int

const (
	// AlignLeft aligns the cell content to the left.
	AlignLeft CellAlign = iota
	// AlignCenter centers the cell content.
	AlignCenter
	// AlignRight aligns the cell content to the right.
	AlignRight
)

// String returns the Telegram Bot API string value for the horizontal alignment.
func (a CellAlign) String() string {
	switch a {
	case AlignCenter:
		return "center"
	case AlignRight:
		return "right"
	default:
		return "left"
	}
}

// CellValign represents the vertical alignment of a table cell content.
//
// See: https://core.telegram.org/bots/api#richblocktablecell
type CellValign int

const (
	// ValignTop aligns the cell content to the top.
	ValignTop CellValign = iota
	// ValignMiddle centers the cell content vertically.
	ValignMiddle
	// ValignBottom aligns the cell content to the bottom.
	ValignBottom
)

// String returns the Telegram Bot API string value for the vertical alignment.
func (v CellValign) String() string {
	switch v {
	case ValignMiddle:
		return "middle"
	case ValignBottom:
		return "bottom"
	default:
		return "top"
	}
}
