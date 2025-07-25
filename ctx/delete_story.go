package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

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
