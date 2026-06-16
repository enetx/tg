package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/file"
)

// MediaSticker represents an input media sticker builder.
// It can be attached to a poll option only.
type MediaSticker struct {
	input *gotgbot.InputMediaSticker
}

// StickerMedia creates a new MediaSticker builder with the required sticker file.
func StickerMedia(media file.InputFile) *MediaSticker {
	return &MediaSticker{
		input: &gotgbot.InputMediaSticker{
			Media: media.Doc,
		},
	}
}

// Emoji sets the emoji associated with the sticker; only for "regular" and "custom_emoji" stickers.
func (ms *MediaSticker) Emoji(emoji g.String) *MediaSticker {
	ms.input.Emoji = emoji.Std()
	return ms
}

// BuildPollOptionMedia creates the gotgbot.InputPollOptionMedia for use as poll option media.
func (ms *MediaSticker) BuildPollOptionMedia() gotgbot.InputPollOptionMedia {
	return *ms.input
}
