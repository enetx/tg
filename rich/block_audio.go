package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/input"
)

// BlockAudio represents a rich message audio block builder.
type BlockAudio struct {
	block *gotgbot.InputRichBlockAudio
}

// Audio creates a new BlockAudio builder with the required audio.
func Audio(audio *input.MediaAudio) *BlockAudio {
	ba := &BlockAudio{block: new(gotgbot.InputRichBlockAudio)}
	return ba.Audio(audio)
}

// Audio sets the block audio.
func (ba *BlockAudio) Audio(audio *input.MediaAudio) *BlockAudio {
	if audio != nil {
		if media, ok := audio.Build().(gotgbot.InputMediaAudio); ok {
			ba.block.Audio = media
		}
	}

	return ba
}

// Caption sets the block caption.
func (ba *BlockAudio) Caption(caption *Caption) *BlockAudio {
	ba.block.Caption = caption.Std()
	return ba
}

// Build creates the gotgbot.InputRichBlock.
func (ba *BlockAudio) Build() gotgbot.InputRichBlock {
	return *ba.block
}
