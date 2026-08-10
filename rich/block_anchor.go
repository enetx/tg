package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// BlockAnchor represents a rich message anchor block builder.
type BlockAnchor struct {
	block *gotgbot.InputRichBlockAnchor
}

// Anchor creates a new BlockAnchor builder with the required anchor name.
func Anchor(name g.String) *BlockAnchor {
	return &BlockAnchor{
		block: &gotgbot.InputRichBlockAnchor{
			Name: name.Std(),
		},
	}
}

// Name sets the anchor name.
func (ba *BlockAnchor) Name(name g.String) *BlockAnchor {
	ba.block.Name = name.Std()
	return ba
}

// Build creates the gotgbot.InputRichBlock.
func (ba *BlockAnchor) Build() gotgbot.InputRichBlock {
	return *ba.block
}
