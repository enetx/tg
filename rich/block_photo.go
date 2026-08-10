package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/input"
)

// BlockPhoto represents a rich message photo block builder.
type BlockPhoto struct {
	block *gotgbot.InputRichBlockPhoto
}

// Photo creates a new BlockPhoto builder with the required photo.
func Photo(photo *input.MediaPhoto) *BlockPhoto {
	bp := &BlockPhoto{block: new(gotgbot.InputRichBlockPhoto)}
	return bp.Photo(photo)
}

// Photo sets the block photo.
func (bp *BlockPhoto) Photo(photo *input.MediaPhoto) *BlockPhoto {
	if photo != nil {
		if media, ok := photo.Build().(gotgbot.InputMediaPhoto); ok {
			bp.block.Photo = media
		}
	}

	return bp
}

// Caption sets the block caption.
func (bp *BlockPhoto) Caption(caption *Caption) *BlockPhoto {
	bp.block.Caption = caption.Std()
	return bp
}

// Build creates the gotgbot.InputRichBlock.
func (bp *BlockPhoto) Build() gotgbot.InputRichBlock {
	return *bp.block
}
