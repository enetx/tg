package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/richtext"
)

// BlockThinking represents a rich message thinking block builder.
type BlockThinking struct {
	block *gotgbot.InputRichBlockThinking
}

// Thinking creates a new BlockThinking builder with the required text.
func Thinking(text *richtext.Text) *BlockThinking {
	return &BlockThinking{
		block: &gotgbot.InputRichBlockThinking{
			Text: text.Std(),
		},
	}
}

// Text sets the thinking text.
func (bt *BlockThinking) Text(text *richtext.Text) *BlockThinking {
	bt.block.Text = text.Std()
	return bt
}

// Build creates the gotgbot.InputRichBlock.
func (bt *BlockThinking) Build() gotgbot.InputRichBlock {
	return *bt.block
}
