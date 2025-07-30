package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/areas"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/input"
)

// EditStory represents a request to edit an existing story.
type EditStory struct {
	ctx                  *Context
	businessConnectionID g.String
	storyID              int64
	content              input.StoryContent
	opts                 *gotgbot.EditStoryOpts
}

// Caption sets the new caption text for the story.
func (es *EditStory) Caption(caption g.String) *EditStory {
	es.opts.Caption = caption.Std()
	return es
}

// HTML sets the story caption parse mode to HTML.
func (es *EditStory) HTML() *EditStory {
	es.opts.ParseMode = "HTML"
	return es
}

// Markdown sets the story caption parse mode to MarkdownV2.
func (es *EditStory) Markdown() *EditStory {
	es.opts.ParseMode = "MarkdownV2"
	return es
}

// ParseMode sets the parse mode for the caption.
func (es *EditStory) ParseMode(mode g.String) *EditStory {
	es.opts.ParseMode = mode.Std()
	return es
}

// CaptionEntities sets custom formatting entities for the caption.
func (es *EditStory) CaptionEntities(e *entities.Entities) *EditStory {
	es.opts.CaptionEntities = e.Std()
	return es
}

// Areas updates the clickable areas on the story using Areas builder.
func (es *EditStory) Areas(a *areas.Areas) *EditStory {
	es.opts.Areas = a.Std()
	return es
}

// Timeout sets a custom timeout for this request.
func (es *EditStory) Timeout(duration time.Duration) *EditStory {
	if es.opts.RequestOpts == nil {
		es.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	es.opts.RequestOpts.Timeout = duration

	return es
}

// APIURL sets a custom API URL for this request.
func (es *EditStory) APIURL(url g.String) *EditStory {
	if es.opts.RequestOpts == nil {
		es.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	es.opts.RequestOpts.APIURL = url.Std()

	return es
}

// Send executes the EditStory request.
func (es *EditStory) Send() g.Result[*gotgbot.Story] {
	return g.ResultOf(es.ctx.Bot.Raw().EditStory(
		es.businessConnectionID.Std(),
		es.storyID,
		es.content.Build(),
		es.opts,
	))
}
