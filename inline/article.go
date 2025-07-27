package inline

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/keyboard"
)

// InlineArticle represents an inline query result article builder.
type InlineArticle struct {
	inline *gotgbot.InlineQueryResultArticle
}

// Article creates a new Article builder with ContentBuilder.
func Article(id, title g.String, message input.MessageContent) *InlineArticle {
	return &InlineArticle{
		inline: &gotgbot.InlineQueryResultArticle{
			Id:                  id.Std(),
			Title:               title.Std(),
			InputMessageContent: message.Build(),
		},
	}
}

// URL sets the URL associated with the article.
func (a *InlineArticle) URL(url g.String) *InlineArticle {
	a.inline.Url = url.Std()
	return a
}

// Description sets the short description of the article.
func (a *InlineArticle) Description(desc g.String) *InlineArticle {
	a.inline.Description = desc.Std()
	return a
}

// ThumbnailURL sets the URL of the thumbnail for the article.
func (a *InlineArticle) ThumbnailURL(url g.String) *InlineArticle {
	a.inline.ThumbnailUrl = url.Std()
	return a
}

// ThumbnailSize sets the thumbnail width and height.
func (a *InlineArticle) ThumbnailSize(width, height int64) *InlineArticle {
	a.inline.ThumbnailWidth = width
	a.inline.ThumbnailHeight = height

	return a
}

// Markup sets the inline keyboard attached to the message.
func (a *InlineArticle) Markup(kb keyboard.Keyboard) *InlineArticle {
	if markup := kb.Markup(); markup != nil {
		if ikm, ok := markup.(gotgbot.InlineKeyboardMarkup); ok {
			a.inline.ReplyMarkup = &ikm
		}
	}

	return a
}

// Build creates the gotgbot.InlineQueryResultArticle.
func (a *InlineArticle) Build() gotgbot.InlineQueryResult {
	return *a.inline
}
