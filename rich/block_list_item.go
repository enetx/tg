package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// BlockListItem represents a rich message list item builder.
type BlockListItem struct {
	item   *gotgbot.InputRichBlockListItem
	blocks g.Slice[Block]
}

// ListItem creates a new BlockListItem builder with the item content blocks.
func ListItem(blocks ...Block) *BlockListItem {
	return (&BlockListItem{
		item:   new(gotgbot.InputRichBlockListItem),
		blocks: g.NewSlice[Block](),
	}).Block(blocks...)
}

// Block appends blocks to the item content.
func (bli *BlockListItem) Block(blocks ...Block) *BlockListItem {
	g.Slice[Block](blocks).
		Iter().
		Exclude(func(block Block) bool { return block == nil }).
		ForEach(func(block Block) { bli.blocks.Push(block) })

	return bli
}

// Checkbox adds an unchecked checkbox to the item.
func (bli *BlockListItem) Checkbox() *BlockListItem {
	bli.item.HasCheckbox = true
	return bli
}

// Checked adds a checked checkbox to the item.
func (bli *BlockListItem) Checked() *BlockListItem {
	bli.item.HasCheckbox = true
	bli.item.IsChecked = true

	return bli
}

// Value sets the numeric value of the item label for ordered lists.
func (bli *BlockListItem) Value(value int64) *BlockListItem {
	bli.item.Value = value
	return bli
}

// Label sets the type of the item label for ordered lists.
func (bli *BlockListItem) Label(label ListLabel) *BlockListItem {
	bli.item.Type = label.String()
	return bli
}

// Std returns the underlying gotgbot.InputRichBlockListItem.
func (bli *BlockListItem) Std() gotgbot.InputRichBlockListItem {
	item := *bli.item
	if !bli.blocks.IsEmpty() {
		item.Blocks = g.TransformSlice(bli.blocks, Block.Build)
	}

	return item
}
