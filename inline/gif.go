package inline

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/inline/content"
	"github.com/enetx/tg/keyboard"
)

// Gif represents an inline query result GIF builder.
type Gif struct {
	inline *gotgbot.InlineQueryResultGif
}

// NewGif creates a new Gif builder with the required fields.
func NewGif(id, gifURL, thumbnailURL g.String) *Gif {
	return &Gif{
		inline: &gotgbot.InlineQueryResultGif{
			Id:           id.Std(),
			GifUrl:       gifURL.Std(),
			ThumbnailUrl: thumbnailURL.Std(),
		},
	}
}

// Size sets the GIF width and height.
func (gf *Gif) Size(width, height int64) *Gif {
	gf.inline.GifWidth = width
	gf.inline.GifHeight = height

	return gf
}

// Duration sets the GIF duration.
func (gf *Gif) Duration(duration time.Duration) *Gif {
	gf.inline.GifDuration = int64(duration.Seconds())
	return gf
}

// Title sets the title for the result.
func (gf *Gif) Title(title g.String) *Gif {
	gf.inline.Title = title.Std()
	return gf
}

func (gf *Gif) ThumbnailMimeType(mimetype g.String) *Gif {
	gf.inline.ThumbnailMimeType = mimetype.Std()
	return gf
}

// Caption sets the caption for the GIF.
func (gf *Gif) Caption(caption g.String) *Gif {
	gf.inline.Caption = caption.Std()
	return gf
}

// HTML sets parse mode to HTML for the caption.
func (gf *Gif) HTML() *Gif {
	gf.inline.ParseMode = "HTML"
	return gf
}

// Markdown sets parse mode to MarkdownV2 for the caption.
func (gf *Gif) Markdown() *Gif {
	gf.inline.ParseMode = "MarkdownV2"
	return gf
}

// CaptionEntities sets the message entities for the caption.
func (gf *Gif) CaptionEntities(e entities.Entities) *Gif {
	gf.inline.CaptionEntities = e.Std()
	return gf
}

// ShowCaptionAboveMedia sets whether to show the caption above the media.
func (gf *Gif) ShowCaptionAboveMedia() *Gif {
	gf.inline.ShowCaptionAboveMedia = true
	return gf
}

// Markup sets the inline keyboard attached to the message.
func (gf *Gif) Markup(kb keyboard.Keyboard) *Gif {
	if markup := kb.Markup(); markup != nil {
		if ikm, ok := markup.(gotgbot.InlineKeyboardMarkup); ok {
			gf.inline.ReplyMarkup = &ikm
		}
	}

	return gf
}

// InputMessageContent sets the content of the message to be sent instead of the GIF.
func (gf *Gif) InputMessageContent(message content.Content) *Gif {
	gf.inline.InputMessageContent = message.Build()
	return gf
}

// Build creates the gotgbot.InlineQueryResultGif.
func (gf *Gif) Build() gotgbot.InlineQueryResult {
	return *gf.inline
}
