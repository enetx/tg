package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/richtext"
)

// BlockHeading represents a rich message section heading block builder.
type BlockHeading struct {
	block *gotgbot.InputRichBlockSectionHeading
}

// Heading creates a new BlockHeading builder with the required text.
func Heading(text *richtext.Text) *BlockHeading {
	return &BlockHeading{
		block: &gotgbot.InputRichBlockSectionHeading{
			Text: text.Std(),
		},
	}
}

// Text sets the heading text.
func (bh *BlockHeading) Text(text *richtext.Text) *BlockHeading {
	bh.block.Text = text.Std()
	return bh
}

// Size sets the heading level, from 1 for the largest to 6 for the smallest.
func (bh *BlockHeading) Size(size int64) *BlockHeading {
	bh.block.Size = size
	return bh
}

// Build creates the gotgbot.InputRichBlock.
func (bh *BlockHeading) Build() gotgbot.InputRichBlock {
	return *bh.block
}
