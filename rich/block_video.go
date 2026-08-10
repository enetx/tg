package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/input"
)

// BlockVideo represents a rich message video block builder.
type BlockVideo struct {
	block *gotgbot.InputRichBlockVideo
}

// Video creates a new BlockVideo builder with the required video.
func Video(video *input.MediaVideo) *BlockVideo {
	bv := &BlockVideo{block: new(gotgbot.InputRichBlockVideo)}
	return bv.Video(video)
}

// Video sets the block video.
func (bv *BlockVideo) Video(video *input.MediaVideo) *BlockVideo {
	if video != nil {
		if media, ok := video.Build().(gotgbot.InputMediaVideo); ok {
			bv.block.Video = media
		}
	}

	return bv
}

// Caption sets the block caption.
func (bv *BlockVideo) Caption(caption *Caption) *BlockVideo {
	bv.block.Caption = caption.Std()
	return bv
}

// Build creates the gotgbot.InputRichBlock.
func (bv *BlockVideo) Build() gotgbot.InputRichBlock {
	return *bv.block
}
