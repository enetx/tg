package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/richtext"
)

// BlockParagraph represents a rich message paragraph block builder.
type BlockParagraph struct {
	block *gotgbot.InputRichBlockParagraph
}

// Paragraph creates a new BlockParagraph builder with the required text.
func Paragraph(text *richtext.Text) *BlockParagraph {
	return &BlockParagraph{
		block: &gotgbot.InputRichBlockParagraph{
			Text: text.Std(),
		},
	}
}

// Text sets the paragraph text.
func (bp *BlockParagraph) Text(text *richtext.Text) *BlockParagraph {
	bp.block.Text = text.Std()
	return bp
}

// Build creates the gotgbot.InputRichBlock.
func (bp *BlockParagraph) Build() gotgbot.InputRichBlock {
	return *bp.block
}
