package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/areas"
	"github.com/enetx/tg/entities"
)

// PostStory represents a request to post a story to a business account.
type PostStory struct {
	ctx                  *Context
	businessConnectionID String
	content              gotgbot.InputStoryContent
	activePeriod         int64
	opts                 *gotgbot.PostStoryOpts
	storyType            string
}

// Caption sets the story caption text.
func (ps *PostStory) Caption(caption String) *PostStory {
	ps.opts.Caption = caption.Std()
	return ps
}

// HTML sets the story caption parse mode to HTML.
func (ps *PostStory) HTML() *PostStory {
	ps.opts.ParseMode = "HTML"
	return ps
}

// Markdown sets the story caption parse mode to MarkdownV2.
func (ps *PostStory) Markdown() *PostStory {
	ps.opts.ParseMode = "MarkdownV2"
	return ps
}

// CaptionEntities sets custom formatting entities for the caption.
func (ps *PostStory) CaptionEntities(e *entities.Entities) *PostStory {
	ps.opts.CaptionEntities = e.Std()
	return ps
}

// Areas adds clickable areas to the story using Areas builder.
func (ps *PostStory) Areas(a *areas.Areas) *PostStory {
	ps.opts.Areas = a.Std()
	return ps
}

// ActivePeriod sets how long the story will be active before being archived.
// Valid values: 21600 (6h), 43200 (12h), 86400 (24h), 172800 (48h)
func (ps *PostStory) ActivePeriod(seconds int64) *PostStory {
	ps.activePeriod = seconds
	return ps
}

// PostToChatPage determines if the story should be posted to the chat page as well.
func (ps *PostStory) PostToChatPage() *PostStory {
	ps.opts.PostToChatPage = true
	return ps
}

// Protect protects the story content from forwarding and saving.
func (ps *PostStory) Protect() *PostStory {
	ps.opts.ProtectContent = true
	return ps
}

// CoverFrame sets the cover frame timestamp for video stories (0-60 seconds).
func (ps *PostStory) CoverFrame(timestamp float64) *PostStory {
	if ps.storyType == "video" {
		if videoContent, ok := ps.content.(*gotgbot.InputStoryContentVideo); ok {
			videoContent.CoverFrameTimestamp = timestamp
		}
	}

	return ps
}

// Timeout sets a custom timeout for this request.
func (ps *PostStory) Timeout(duration time.Duration) *PostStory {
	if ps.opts.RequestOpts == nil {
		ps.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ps.opts.RequestOpts.Timeout = duration

	return ps
}

// APIURL sets a custom API URL for this request.
func (ps *PostStory) APIURL(url String) *PostStory {
	if ps.opts.RequestOpts == nil {
		ps.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ps.opts.RequestOpts.APIURL = url.Std()

	return ps
}

// Send executes the PostStory request.
func (ps *PostStory) Send() Result[*gotgbot.Story] {
	return ResultOf(ps.ctx.Bot.Raw().PostStory(
		ps.businessConnectionID.Std(),
		ps.content,
		ps.activePeriod,
		ps.opts,
	))
}

// EditStory represents a request to edit an existing story.
type EditStory struct {
	ctx                  *Context
	businessConnectionID String
	storyID              int64
	content              gotgbot.InputStoryContent
	opts                 *gotgbot.EditStoryOpts
	storyType            string
}

// Caption sets the new caption text for the story.
func (es *EditStory) Caption(caption String) *EditStory {
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
func (es *EditStory) ParseMode(mode String) *EditStory {
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

// CoverFrame sets the cover frame timestamp for video stories (0-60 seconds).
func (es *EditStory) CoverFrame(timestamp float64) *EditStory {
	if es.storyType == "video" {
		if videoContent, ok := es.content.(*gotgbot.InputStoryContentVideo); ok {
			videoContent.CoverFrameTimestamp = timestamp
		}
	}

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
func (es *EditStory) APIURL(url String) *EditStory {
	if es.opts.RequestOpts == nil {
		es.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	es.opts.RequestOpts.APIURL = url.Std()

	return es
}

// Send executes the EditStory request.
func (es *EditStory) Send() Result[*gotgbot.Story] {
	return ResultOf(es.ctx.Bot.Raw().EditStory(
		es.businessConnectionID.Std(),
		es.storyID,
		es.content,
		es.opts,
	))
}

// DeleteStory represents a request to delete a story.
type DeleteStory struct {
	ctx                  *Context
	businessConnectionID String
	storyID              int64
	opts                 *gotgbot.DeleteStoryOpts
}

// Timeout sets a custom timeout for this request.
func (ds *DeleteStory) Timeout(duration time.Duration) *DeleteStory {
	if ds.opts.RequestOpts == nil {
		ds.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ds.opts.RequestOpts.Timeout = duration

	return ds
}

// APIURL sets a custom API URL for this request.
func (ds *DeleteStory) APIURL(url String) *DeleteStory {
	if ds.opts.RequestOpts == nil {
		ds.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ds.opts.RequestOpts.APIURL = url.Std()

	return ds
}

// Send executes the DeleteStory request.
func (ds *DeleteStory) Send() Result[bool] {
	return ResultOf(ds.ctx.Bot.Raw().DeleteStory(
		ds.businessConnectionID.Std(),
		ds.storyID,
		ds.opts,
	))
}
