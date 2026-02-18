package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// GetUserProfileAudios represents a request to get user profile audios.
type GetUserProfileAudios struct {
	ctx    *Context
	userID int64
	opts   *gotgbot.GetUserProfileAudiosOpts
}

// Offset sets the sequential number of the first audio to be returned.
func (gupa *GetUserProfileAudios) Offset(offset int64) *GetUserProfileAudios {
	gupa.opts.Offset = offset
	return gupa
}

// Limit sets the number of audios to fetch (1-100).
func (gupa *GetUserProfileAudios) Limit(limit int64) *GetUserProfileAudios {
	gupa.opts.Limit = limit
	return gupa
}

// Timeout sets a custom timeout for this request.
func (gupa *GetUserProfileAudios) Timeout(duration time.Duration) *GetUserProfileAudios {
	if gupa.opts.RequestOpts == nil {
		gupa.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gupa.opts.RequestOpts.Timeout = duration

	return gupa
}

// APIURL sets a custom API URL for this request.
func (gupa *GetUserProfileAudios) APIURL(url g.String) *GetUserProfileAudios {
	if gupa.opts.RequestOpts == nil {
		gupa.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gupa.opts.RequestOpts.APIURL = url.Std()

	return gupa
}

// Send gets user profile audios and returns the result.
func (gupa *GetUserProfileAudios) Send() g.Result[*gotgbot.UserProfileAudios] {
	return g.ResultOf(gupa.ctx.Bot.Raw().GetUserProfileAudios(gupa.userID, gupa.opts))
}
