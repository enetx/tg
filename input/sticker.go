package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// Sticker represents an input sticker builder.
type Sticker struct {
	input *gotgbot.InputSticker
}

// NewSticker creates a new Sticker builder.
func NewSticker(sticker, format g.String, emojiList g.Slice[g.String]) *Sticker {
	return &Sticker{
		input: &gotgbot.InputSticker{
			Sticker:   sticker.Std(),
			Format:    format.Std(),
			EmojiList: emojiList.ToStringSlice(),
		},
	}
}

// MaskPosition sets the position where the mask should be placed on faces.
func (s *Sticker) MaskPosition(maskPosition *gotgbot.MaskPosition) *Sticker {
	s.input.MaskPosition = maskPosition
	return s
}

// Keywords sets keywords for the sticker.
func (s *Sticker) Keywords(keywords g.Slice[g.String]) *Sticker {
	s.input.Keywords = keywords.ToStringSlice()
	return s
}

// Build returns the gotgbot.InputSticker directly as it's not an interface.
func (s *Sticker) Build() gotgbot.InputSticker {
	return *s.input
}
