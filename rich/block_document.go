package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/input"
)

// BlockDocument represents a rich message document block builder.
type BlockDocument struct {
	block *gotgbot.InputRichBlockDocument
}

// Document creates a new BlockDocument builder with the required document.
func Document(document *input.MediaDocument) *BlockDocument {
	bd := &BlockDocument{block: new(gotgbot.InputRichBlockDocument)}
	return bd.Document(document)
}

// Document sets the block document.
func (bd *BlockDocument) Document(document *input.MediaDocument) *BlockDocument {
	if document != nil {
		if media, ok := document.Build().(gotgbot.InputMediaDocument); ok {
			bd.block.Document = media
		}
	}

	return bd
}

// Caption sets the block caption.
func (bd *BlockDocument) Caption(caption *Caption) *BlockDocument {
	bd.block.Caption = caption.Std()
	return bd
}

// Build creates the gotgbot.InputRichBlock.
func (bd *BlockDocument) Build() gotgbot.InputRichBlock {
	return *bd.block
}
