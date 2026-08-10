package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/richtext"
)

// BlockQuotation represents a rich message block quotation builder.
type BlockQuotation struct {
	block  *gotgbot.InputRichBlockBlockQuotation
	blocks g.Slice[Block]
}

// Quotation creates a new BlockQuotation builder with the quoted blocks.
func Quotation(blocks ...Block) *BlockQuotation {
	return (&BlockQuotation{
		block:  new(gotgbot.InputRichBlockBlockQuotation),
		blocks: g.NewSlice[Block](),
	}).Block(blocks...)
}

// Block appends blocks to the quotation content.
func (bq *BlockQuotation) Block(blocks ...Block) *BlockQuotation {
	g.Slice[Block](blocks).
		Iter().
		Exclude(func(block Block) bool { return block == nil }).
		ForEach(func(block Block) { bq.blocks.Push(block) })

	return bq
}

// Credit sets the quotation credit, usually the author or the source.
func (bq *BlockQuotation) Credit(credit *richtext.Text) *BlockQuotation {
	bq.block.Credit = credit.Std()
	return bq
}

// Build creates the gotgbot.InputRichBlock.
func (bq *BlockQuotation) Build() gotgbot.InputRichBlock {
	block := *bq.block
	if !bq.blocks.IsEmpty() {
		block.Blocks = g.TransformSlice(bq.blocks, Block.Build)
	}

	return block
}
