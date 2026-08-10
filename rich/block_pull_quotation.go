package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/richtext"
)

// BlockPullQuotation represents a rich message pull quotation block builder.
type BlockPullQuotation struct {
	block *gotgbot.InputRichBlockPullQuotation
}

// PullQuotation creates a new BlockPullQuotation builder with the required text.
func PullQuotation(text *richtext.Text) *BlockPullQuotation {
	return &BlockPullQuotation{
		block: &gotgbot.InputRichBlockPullQuotation{
			Text: text.Std(),
		},
	}
}

// Text sets the pull quotation text.
func (bpq *BlockPullQuotation) Text(text *richtext.Text) *BlockPullQuotation {
	bpq.block.Text = text.Std()
	return bpq
}

// Credit sets the pull quotation credit, usually the author or the source.
func (bpq *BlockPullQuotation) Credit(credit *richtext.Text) *BlockPullQuotation {
	bpq.block.Credit = credit.Std()
	return bpq
}

// Build creates the gotgbot.InputRichBlock.
func (bpq *BlockPullQuotation) Build() gotgbot.InputRichBlock {
	return *bpq.block
}
