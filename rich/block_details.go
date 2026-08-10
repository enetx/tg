package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/richtext"
)

// BlockDetails represents a rich message collapsible details block builder.
type BlockDetails struct {
	block  *gotgbot.InputRichBlockDetails
	blocks g.Slice[Block]
}

// Details creates a new BlockDetails builder with the required summary.
func Details(summary *richtext.Text) *BlockDetails {
	return &BlockDetails{
		block:  &gotgbot.InputRichBlockDetails{Summary: summary.Std()},
		blocks: g.NewSlice[Block](),
	}
}

// Summary sets the always visible summary of the block.
func (bd *BlockDetails) Summary(summary *richtext.Text) *BlockDetails {
	bd.block.Summary = summary.Std()
	return bd
}

// Block appends blocks to the collapsible content.
func (bd *BlockDetails) Block(blocks ...Block) *BlockDetails {
	g.Slice[Block](blocks).
		Iter().
		Exclude(func(block Block) bool { return block == nil }).
		ForEach(func(block Block) { bd.blocks.Push(block) })

	return bd
}

// Open marks the block as expanded by default.
func (bd *BlockDetails) Open() *BlockDetails {
	bd.block.IsOpen = true
	return bd
}

// Build creates the gotgbot.InputRichBlock.
func (bd *BlockDetails) Build() gotgbot.InputRichBlock {
	block := *bd.block
	if !bd.blocks.IsEmpty() {
		block.Blocks = g.TransformSlice(bd.blocks, Block.Build)
	}

	return block
}
