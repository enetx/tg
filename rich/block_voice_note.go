package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/input"
)

// BlockVoiceNote represents a rich message voice note block builder.
type BlockVoiceNote struct {
	block *gotgbot.InputRichBlockVoiceNote
}

// VoiceNote creates a new BlockVoiceNote builder with the required voice note.
func VoiceNote(voiceNote *input.MediaVoiceNote) *BlockVoiceNote {
	bvn := &BlockVoiceNote{block: new(gotgbot.InputRichBlockVoiceNote)}
	return bvn.VoiceNote(voiceNote)
}

// VoiceNote sets the block voice note. Its caption is ignored by Telegram.
func (bvn *BlockVoiceNote) VoiceNote(voiceNote *input.MediaVoiceNote) *BlockVoiceNote {
	if voiceNote != nil {
		bvn.block.VoiceNote = voiceNote.Build()
	}

	return bvn
}

// Caption sets the block caption.
func (bvn *BlockVoiceNote) Caption(caption *Caption) *BlockVoiceNote {
	bvn.block.Caption = caption.Std()
	return bvn
}

// Build creates the gotgbot.InputRichBlock.
func (bvn *BlockVoiceNote) Build() gotgbot.InputRichBlock {
	return *bvn.block
}
