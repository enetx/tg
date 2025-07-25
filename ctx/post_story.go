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

// ActiveFor sets how long the story will be active before being archived.
func (ps *PostStory) ActiveFor(duration time.Duration) *PostStory {
	ps.activePeriod = int64(duration.Seconds())
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
