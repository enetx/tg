package rich

import "github.com/PaulSonOfLars/gotgbot/v2"

// BlockDivider represents a rich message divider block builder.
type BlockDivider struct {
	block *gotgbot.InputRichBlockDivider
}

// Divider creates a new BlockDivider builder.
func Divider() *BlockDivider {
	return &BlockDivider{block: new(gotgbot.InputRichBlockDivider)}
}

// Build creates the gotgbot.InputRichBlock.
func (bd *BlockDivider) Build() gotgbot.InputRichBlock {
	return *bd.block
}
