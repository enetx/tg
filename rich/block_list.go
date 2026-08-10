package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// BlockList represents a rich message list block builder.
type BlockList struct {
	block *gotgbot.InputRichBlockList
	items g.Slice[*BlockListItem]
}

// List creates a new BlockList builder with the given items.
func List(items ...*BlockListItem) *BlockList {
	return (&BlockList{
		block: new(gotgbot.InputRichBlockList),
		items: g.NewSlice[*BlockListItem](),
	}).Item(items...)
}

// Item appends items to the list.
func (bl *BlockList) Item(items ...*BlockListItem) *BlockList {
	g.Slice[*BlockListItem](items).
		Iter().
		Exclude(func(item *BlockListItem) bool { return item == nil }).
		ForEach(func(item *BlockListItem) { bl.items.Push(item) })

	return bl
}

// Build creates the gotgbot.InputRichBlock.
func (bl *BlockList) Build() gotgbot.InputRichBlock {
	block := *bl.block
	if !bl.items.IsEmpty() {
		block.Items = g.TransformSlice(bl.items, (*BlockListItem).Std)
	}

	return block
}
