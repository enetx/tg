package rich

// ListLabel represents the type of the item label in an ordered list.
//
// See: https://core.telegram.org/bots/api#inputrichblocklistitem
type ListLabel int

const (
	// LabelNone leaves the item unlabeled, producing an unordered list item.
	LabelNone ListLabel = iota
	// LabelLowerAlpha labels items with lowercase letters.
	LabelLowerAlpha
	// LabelUpperAlpha labels items with uppercase letters.
	LabelUpperAlpha
	// LabelLowerRoman labels items with lowercase Roman numerals.
	LabelLowerRoman
	// LabelUpperRoman labels items with uppercase Roman numerals.
	LabelUpperRoman
	// LabelDecimal labels items with decimal numbers.
	LabelDecimal
)

// String returns the Telegram Bot API string value for the list label.
func (l ListLabel) String() string {
	switch l {
	case LabelLowerAlpha:
		return "a"
	case LabelUpperAlpha:
		return "A"
	case LabelLowerRoman:
		return "i"
	case LabelUpperRoman:
		return "I"
	case LabelDecimal:
		return "1"
	default:
		return ""
	}
}
