package inline

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/inline/content"
	"github.com/enetx/tg/keyboard"
)

// Article represents an inline query result article builder.
type Article struct {
	inline *gotgbot.InlineQueryResultArticle
}

// NewArticle creates a new Article builder with ContentBuilder.
func NewArticle(id, title g.String, message content.Content) *Article {
	return &Article{
		inline: &gotgbot.InlineQueryResultArticle{
			Id:                  id.Std(),
			Title:               title.Std(),
			InputMessageContent: message.Build(),
		},
	}
}

// URL sets the URL associated with the article.
func (a *Article) URL(url g.String) *Article {
	a.inline.Url = url.Std()
	return a
}

// Description sets the short description of the article.
func (a *Article) Description(desc g.String) *Article {
	a.inline.Description = desc.Std()
	return a
}

// ThumbnailURL sets the URL of the thumbnail for the article.
func (a *Article) ThumbnailURL(url g.String) *Article {
	a.inline.ThumbnailUrl = url.Std()
	return a
}

// ThumbnailSize sets the thumbnail width and height.
func (a *Article) ThumbnailSize(width, height int64) *Article {
	a.inline.ThumbnailWidth = width
	a.inline.ThumbnailHeight = height

	return a
}

// Markup sets the inline keyboard attached to the message.
func (a *Article) Markup(kb keyboard.Keyboard) *Article {
	if markup := kb.Markup(); markup != nil {
		if ikm, ok := markup.(gotgbot.InlineKeyboardMarkup); ok {
			a.inline.ReplyMarkup = &ikm
		}
	}

	return a
}

// Build creates the gotgbot.InlineQueryResultArticle.
func (a *Article) Build() gotgbot.InlineQueryResult {
	return *a.inline
}
