package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// RepostStory represents a request to repost a story from another business account.
type RepostStory struct {
	ctx                  *Context
	businessConnectionID g.String
	fromChatID           int64
	fromStoryID          int64
	activePeriod         int64
	opts                 *gotgbot.RepostStoryOpts
}

// PostToChatPage determines if the story should be posted to the chat page as well.
func (rs *RepostStory) PostToChatPage() *RepostStory {
	rs.opts.PostToChatPage = true
	return rs
}

// Protect protects the story content from forwarding and saving.
func (rs *RepostStory) Protect() *RepostStory {
	rs.opts.ProtectContent = true
	return rs
}

// ActiveFor sets how long the story will be active before being archived.
func (rs *RepostStory) ActiveFor(duration time.Duration) *RepostStory {
	rs.activePeriod = int64(duration.Seconds())
	return rs
}

// Timeout sets a custom timeout for this request.
func (rs *RepostStory) Timeout(duration time.Duration) *RepostStory {
	if rs.opts.RequestOpts == nil {
		rs.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rs.opts.RequestOpts.Timeout = duration

	return rs
}

// APIURL sets a custom API URL for this request.
func (rs *RepostStory) APIURL(url g.String) *RepostStory {
	if rs.opts.RequestOpts == nil {
		rs.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rs.opts.RequestOpts.APIURL = url.Std()

	return rs
}

// Send executes the RepostStory request and returns the reposted story.
func (rs *RepostStory) Send() g.Result[*gotgbot.Story] {
	return g.ResultOf(rs.ctx.Bot.Raw().RepostStory(
		rs.businessConnectionID.Std(),
		rs.fromChatID,
		rs.fromStoryID,
		rs.activePeriod,
		rs.opts,
	))
}
