package inline

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/keyboard"
)

// Mpeg4Gif represents an inline query result MPEG4 GIF builder.
type Mpeg4Gif struct {
	inline *gotgbot.InlineQueryResultMpeg4Gif
}

// NewMpeg4Gif creates a new Mpeg4Gif builder with the required fields.
func NewMpeg4Gif(id, mpeg4URL, thumbnailURL g.String) *Mpeg4Gif {
	return &Mpeg4Gif{
		inline: &gotgbot.InlineQueryResultMpeg4Gif{
			Id:           id.Std(),
			Mpeg4Url:     mpeg4URL.Std(),
			ThumbnailUrl: thumbnailURL.Std(),
		},
	}
}

// Size sets the MPEG4 GIF width and height.
func (m *Mpeg4Gif) Size(width, height int64) *Mpeg4Gif {
	m.inline.Mpeg4Width = width
	m.inline.Mpeg4Height = height

	return m
}

// Duration sets the MPEG4 GIF duration.
func (m *Mpeg4Gif) Duration(duration time.Duration) *Mpeg4Gif {
	m.inline.Mpeg4Duration = int64(duration.Seconds())
	return m
}

// ThumbnailMimeType sets the MIME type of the thumbnail.
func (m *Mpeg4Gif) ThumbnailMimeType(mimeType g.String) *Mpeg4Gif {
	m.inline.ThumbnailMimeType = mimeType.Std()
	return m
}

// Title sets the title for the result.
func (m *Mpeg4Gif) Title(title g.String) *Mpeg4Gif {
	m.inline.Title = title.Std()
	return m
}

// Caption sets the caption for the MPEG4 GIF.
func (m *Mpeg4Gif) Caption(caption g.String) *Mpeg4Gif {
	m.inline.Caption = caption.Std()
	return m
}

// HTML sets parse mode to HTML for the caption.
func (m *Mpeg4Gif) HTML() *Mpeg4Gif {
	m.inline.ParseMode = "HTML"
	return m
}

// Markdown sets parse mode to MarkdownV2 for the caption.
func (m *Mpeg4Gif) Markdown() *Mpeg4Gif {
	m.inline.ParseMode = "MarkdownV2"
	return m
}

// CaptionEntities sets the message entities for the caption.
func (m *Mpeg4Gif) CaptionEntities(e entities.Entities) *Mpeg4Gif {
	m.inline.CaptionEntities = e.Std()
	return m
}

// ShowCaptionAboveMedia sets whether to show the caption above the media.
func (m *Mpeg4Gif) ShowCaptionAboveMedia() *Mpeg4Gif {
	m.inline.ShowCaptionAboveMedia = true
	return m
}

// Markup sets the inline keyboard attached to the message.
func (m *Mpeg4Gif) Markup(kb keyboard.Keyboard) *Mpeg4Gif {
	if markup := kb.Markup(); markup != nil {
		if ikm, ok := markup.(gotgbot.InlineKeyboardMarkup); ok {
			m.inline.ReplyMarkup = &ikm
		}
	}

	return m
}

// InputMessageContent sets the content of the message to be sent instead of the MPEG4 GIF.
func (m *Mpeg4Gif) InputMessageContent(message input.MessageContent) *Mpeg4Gif {
	m.inline.InputMessageContent = message.Build()
	return m
}

// Build creates the gotgbot.InlineQueryResultMpeg4Gif.
func (m *Mpeg4Gif) Build() gotgbot.InlineQueryResult {
	return *m.inline
}
