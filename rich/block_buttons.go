package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// BlockButtons represents a rich message button row block builder.
type BlockButtons struct {
	block   *gotgbot.InputRichBlockButtons
	buttons g.Slice[*Button]
}

// Buttons creates a new BlockButtons builder with the given buttons.
func Buttons(buttons ...*Button) *BlockButtons {
	return (&BlockButtons{
		block:   new(gotgbot.InputRichBlockButtons),
		buttons: g.NewSlice[*Button](),
	}).Button(buttons...)
}

// Button appends buttons to the row. A row must contain 1-8 buttons.
func (bb *BlockButtons) Button(buttons ...*Button) *BlockButtons {
	g.Slice[*Button](buttons).
		Iter().
		Exclude(func(button *Button) bool { return button == nil }).
		ForEach(func(button *Button) { bb.buttons.Push(button) })

	return bb
}

// Align sets the horizontal alignment of the buttons.
func (bb *BlockButtons) Align(align CellAlign) *BlockButtons {
	bb.block.Align = align.String()
	return bb
}

// Build creates the gotgbot.InputRichBlock.
func (bb *BlockButtons) Build() gotgbot.InputRichBlock {
	block := *bb.block
	if !bb.buttons.IsEmpty() {
		block.Buttons = g.TransformSlice(bb.buttons, (*Button).Std)
	}

	return block
}
