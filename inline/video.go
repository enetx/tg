package inline

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/inline/content"
	"github.com/enetx/tg/keyboard"
)

// Video represents an inline query result video builder.
type Video struct {
	inline *gotgbot.InlineQueryResultVideo
}

// NewVideo creates a new Video builder with the required fields.
func NewVideo(id, videoURL, mimeType, thumbnailURL, title g.String) *Video {
	return &Video{
		inline: &gotgbot.InlineQueryResultVideo{
			Id:           id.Std(),
			VideoUrl:     videoURL.Std(),
			MimeType:     mimeType.Std(),
			ThumbnailUrl: thumbnailURL.Std(),
			Title:        title.Std(),
		},
	}
}

// Caption sets the caption for the video.
func (v *Video) Caption(caption g.String) *Video {
	v.inline.Caption = caption.Std()
	return v
}

// HTML sets parse mode to HTML for the caption.
func (v *Video) HTML() *Video {
	v.inline.ParseMode = "HTML"
	return v
}

// Markdown sets parse mode to MarkdownV2 for the caption.
func (v *Video) Markdown() *Video {
	v.inline.ParseMode = "MarkdownV2"
	return v
}

// CaptionEntities sets the message entities for the caption.
func (v *Video) CaptionEntities(e entities.Entities) *Video {
	v.inline.CaptionEntities = e.Std()
	return v
}

// ShowCaptionAboveMedia sets whether to show the caption above the media.
func (v *Video) ShowCaptionAboveMedia() *Video {
	v.inline.ShowCaptionAboveMedia = true
	return v
}

// Size sets the video width and height.
func (v *Video) Size(width, height int64) *Video {
	v.inline.VideoWidth = width
	v.inline.VideoHeight = height

	return v
}

// Duration sets the video duration.
func (v *Video) Duration(duration time.Duration) *Video {
	v.inline.VideoDuration = int64(duration.Seconds())
	return v
}

// Description sets the short description of the result.
func (v *Video) Description(desc g.String) *Video {
	v.inline.Description = desc.Std()
	return v
}

// Markup sets the inline keyboard attached to the message.
func (v *Video) Markup(kb keyboard.Keyboard) *Video {
	if markup := kb.Markup(); markup != nil {
		if ikm, ok := markup.(gotgbot.InlineKeyboardMarkup); ok {
			v.inline.ReplyMarkup = &ikm
		}
	}

	return v
}

// InputMessageContent sets the content of the message to be sent instead of the video.
func (v *Video) InputMessageContent(message content.Content) *Video {
	v.inline.InputMessageContent = message.Build()
	return v
}

// Build creates the gotgbot.InlineQueryResultVideo.
func (v *Video) Build() gotgbot.InlineQueryResult {
	return *v.inline
}
