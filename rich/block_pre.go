package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/richtext"
)

// BlockPre represents a rich message preformatted block builder.
type BlockPre struct {
	block *gotgbot.InputRichBlockPreformatted
}

// Pre creates a new BlockPre builder with the required text.
func Pre(text *richtext.Text) *BlockPre {
	return &BlockPre{
		block: &gotgbot.InputRichBlockPreformatted{
			Text: text.Std(),
		},
	}
}

// Text sets the preformatted text.
func (bp *BlockPre) Text(text *richtext.Text) *BlockPre {
	bp.block.Text = text.Std()
	return bp
}

// Language sets the programming language of the block content.
func (bp *BlockPre) Language(language g.String) *BlockPre {
	bp.block.Language = language.Std()
	return bp
}

// Build creates the gotgbot.InputRichBlock.
func (bp *BlockPre) Build() gotgbot.InputRichBlock {
	return *bp.block
}
