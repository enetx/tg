package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/input"
)

// BlockAnimation represents a rich message animation block builder.
type BlockAnimation struct {
	block *gotgbot.InputRichBlockAnimation
}

// Animation creates a new BlockAnimation builder with the required animation.
func Animation(animation *input.MediaAnimation) *BlockAnimation {
	ba := &BlockAnimation{block: new(gotgbot.InputRichBlockAnimation)}
	return ba.Animation(animation)
}

// Animation sets the block animation.
func (ba *BlockAnimation) Animation(animation *input.MediaAnimation) *BlockAnimation {
	if animation != nil {
		if media, ok := animation.Build().(gotgbot.InputMediaAnimation); ok {
			ba.block.Animation = media
		}
	}

	return ba
}

// Caption sets the block caption.
func (ba *BlockAnimation) Caption(caption *Caption) *BlockAnimation {
	ba.block.Caption = caption.Std()
	return ba
}

// Build creates the gotgbot.InputRichBlock.
func (ba *BlockAnimation) Build() gotgbot.InputRichBlock {
	return *ba.block
}
