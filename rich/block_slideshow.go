package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// BlockSlideshow represents a rich message slideshow block builder.
type BlockSlideshow struct {
	block  *gotgbot.InputRichBlockSlideshow
	blocks g.Slice[Block]
}

// Slideshow creates a new BlockSlideshow builder with the slide media blocks.
func Slideshow(blocks ...Block) *BlockSlideshow {
	return (&BlockSlideshow{
		block:  new(gotgbot.InputRichBlockSlideshow),
		blocks: g.NewSlice[Block](),
	}).Block(blocks...)
}

// Block appends blocks to the slideshow content.
func (bs *BlockSlideshow) Block(blocks ...Block) *BlockSlideshow {
	g.Slice[Block](blocks).
		Iter().
		Exclude(func(block Block) bool { return block == nil }).
		ForEach(func(block Block) { bs.blocks.Push(block) })

	return bs
}

// Caption sets the block caption.
func (bs *BlockSlideshow) Caption(caption *Caption) *BlockSlideshow {
	bs.block.Caption = caption.Std()
	return bs
}

// Build creates the gotgbot.InputRichBlock.
func (bs *BlockSlideshow) Build() gotgbot.InputRichBlock {
	block := *bs.block
	if !bs.blocks.IsEmpty() {
		block.Blocks = g.TransformSlice(bs.blocks, Block.Build)
	}

	return block
}
