package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/richtext"
)

// BlockFooter represents a rich message footer block builder.
type BlockFooter struct {
	block *gotgbot.InputRichBlockFooter
}

// Footer creates a new BlockFooter builder with the required text.
func Footer(text *richtext.Text) *BlockFooter {
	return &BlockFooter{
		block: &gotgbot.InputRichBlockFooter{
			Text: text.Std(),
		},
	}
}

// Text sets the footer text.
func (bf *BlockFooter) Text(text *richtext.Text) *BlockFooter {
	bf.block.Text = text.Std()
	return bf
}

// Build creates the gotgbot.InputRichBlock.
func (bf *BlockFooter) Build() gotgbot.InputRichBlock {
	return *bf.block
}
