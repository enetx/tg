package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/file"
)

// InputSticker represents an input sticker builder.
type InputSticker struct {
	input *gotgbot.InputSticker
}

// Sticker creates a new Sticker builder.
func Sticker(media file.InputFile, format g.String, emojiList g.Slice[g.String]) *InputSticker {
	return &InputSticker{
		input: &gotgbot.InputSticker{
			Sticker:   media.Doc,
			Format:    format.Std(),
			EmojiList: emojiList.ToStringSlice(),
		},
	}
}

// MaskPosition sets the position where the mask should be placed on faces.
func (s *InputSticker) MaskPosition(maskPosition *gotgbot.MaskPosition) *InputSticker {
	s.input.MaskPosition = maskPosition
	return s
}

// Keywords sets keywords for the sticker.
func (s *InputSticker) Keywords(keywords g.Slice[g.String]) *InputSticker {
	s.input.Keywords = keywords.ToStringSlice()
	return s
}

// Build returns the gotgbot.InputSticker directly as it's not an interface.
func (s *InputSticker) Build() gotgbot.InputSticker {
	return *s.input
}
