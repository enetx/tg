package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// BlockCollage represents a rich message collage block builder.
type BlockCollage struct {
	block  *gotgbot.InputRichBlockCollage
	blocks g.Slice[Block]
}

// Collage creates a new BlockCollage builder with the collaged media blocks.
func Collage(blocks ...Block) *BlockCollage {
	return (&BlockCollage{
		block:  new(gotgbot.InputRichBlockCollage),
		blocks: g.NewSlice[Block](),
	}).Block(blocks...)
}

// Block appends blocks to the collage content.
func (bc *BlockCollage) Block(blocks ...Block) *BlockCollage {
	g.Slice[Block](blocks).
		Iter().
		Exclude(func(block Block) bool { return block == nil }).
		ForEach(func(block Block) { bc.blocks.Push(block) })

	return bc
}

// Caption sets the block caption.
func (bc *BlockCollage) Caption(caption *Caption) *BlockCollage {
	bc.block.Caption = caption.Std()
	return bc
}

// Build creates the gotgbot.InputRichBlock.
func (bc *BlockCollage) Build() gotgbot.InputRichBlock {
	block := *bc.block
	if !bc.blocks.IsEmpty() {
		block.Blocks = g.TransformSlice(bc.blocks, Block.Build)
	}

	return block
}
