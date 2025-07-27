package inline

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/keyboard"
)

// Photo represents an inline query result photo builder.
type Photo struct {
	inline *gotgbot.InlineQueryResultPhoto
}

// NewPhoto creates a new Photo builder with the required fields.
func NewPhoto(id, photoURL, thumbnailURL g.String) *Photo {
	return &Photo{
		inline: &gotgbot.InlineQueryResultPhoto{
			Id:           id.Std(),
			PhotoUrl:     photoURL.Std(),
			ThumbnailUrl: thumbnailURL.Std(),
		},
	}
}

// Size sets the photo width and height.
func (p *Photo) Size(width, height int64) *Photo {
	p.inline.PhotoWidth = width
	p.inline.PhotoHeight = height

	return p
}

// Title sets the title for the result.
func (p *Photo) Title(title g.String) *Photo {
	p.inline.Title = title.Std()
	return p
}

// Description sets the short description of the result.
func (p *Photo) Description(desc g.String) *Photo {
	p.inline.Description = desc.Std()
	return p
}

// Caption sets the caption for the photo.
func (p *Photo) Caption(caption g.String) *Photo {
	p.inline.Caption = caption.Std()
	return p
}

// HTML sets parse mode to HTML for the caption.
func (p *Photo) HTML() *Photo {
	p.inline.ParseMode = "HTML"
	return p
}

// Markdown sets parse mode to MarkdownV2 for the caption.
func (p *Photo) Markdown() *Photo {
	p.inline.ParseMode = "MarkdownV2"
	return p
}

// CaptionEntities sets the message entities for the caption.
func (p *Photo) CaptionEntities(e entities.Entities) *Photo {
	p.inline.CaptionEntities = e.Std()
	return p
}

// ShowCaptionAboveMedia sets whether to show the caption above the media.
func (p *Photo) ShowCaptionAboveMedia() *Photo {
	p.inline.ShowCaptionAboveMedia = true
	return p
}

// Markup sets the inline keyboard attached to the message.
func (p *Photo) Markup(kb keyboard.Keyboard) *Photo {
	if markup := kb.Markup(); markup != nil {
		if ikm, ok := markup.(gotgbot.InlineKeyboardMarkup); ok {
			p.inline.ReplyMarkup = &ikm
		}
	}

	return p
}

// InputMessageContent sets the content of the message to be sent instead of the photo.
func (p *Photo) InputMessageContent(message input.MessageContent) *Photo {
	p.inline.InputMessageContent = message.Build()
	return p
}

// Build creates the gotgbot.InlineQueryResultPhoto.
func (p *Photo) Build() gotgbot.InlineQueryResult {
	return *p.inline
}
