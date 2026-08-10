package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// BlockMath represents a rich message mathematical expression block builder.
type BlockMath struct {
	block *gotgbot.InputRichBlockMathematicalExpression
}

// Math creates a new BlockMath builder with the required LaTeX expression.
func Math(expression g.String) *BlockMath {
	return &BlockMath{
		block: &gotgbot.InputRichBlockMathematicalExpression{
			Expression: expression.Std(),
		},
	}
}

// Expression sets the mathematical expression.
func (bm *BlockMath) Expression(expression g.String) *BlockMath {
	bm.block.Expression = expression.Std()
	return bm
}

// Build creates the gotgbot.InputRichBlock.
func (bm *BlockMath) Build() gotgbot.InputRichBlock {
	return *bm.block
}
