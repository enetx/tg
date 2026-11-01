package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/richtext"
)

// BlockExpandableQuotation represents a rich message expandable block quotation builder.
type BlockExpandableQuotation struct {
	block *gotgbot.InputRichBlockExpandableBlockQuotation
}

// ExpandableQuotation creates a new BlockExpandableQuotation builder with the required text.
func ExpandableQuotation(text *richtext.Text) *BlockExpandableQuotation {
	return &BlockExpandableQuotation{
		block: &gotgbot.InputRichBlockExpandableBlockQuotation{
			Text: text.Std(),
		},
	}
}

// Text sets the expandable quotation text.
func (beq *BlockExpandableQuotation) Text(text *richtext.Text) *BlockExpandableQuotation {
	beq.block.Text = text.Std()
	return beq
}

// Credit sets the expandable quotation credit, usually the author or the source.
func (beq *BlockExpandableQuotation) Credit(credit *richtext.Text) *BlockExpandableQuotation {
	beq.block.Credit = credit.Std()
	return beq
}

// Build creates the gotgbot.InputRichBlock.
func (beq *BlockExpandableQuotation) Build() gotgbot.InputRichBlock {
	return *beq.block
}
