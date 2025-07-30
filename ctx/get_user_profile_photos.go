package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// GetUserProfilePhotos represents a request to get user profile photos.
type GetUserProfilePhotos struct {
	ctx    *Context
	userID int64
	opts   *gotgbot.GetUserProfilePhotosOpts
	err    error
}

// Offset sets the sequential number of the first photo to be returned.
func (gupp *GetUserProfilePhotos) Offset(offset int64) *GetUserProfilePhotos {
	gupp.opts.Offset = offset
	return gupp
}

// Limit sets the number of photos to fetch (1-100).
func (gupp *GetUserProfilePhotos) Limit(limit int64) *GetUserProfilePhotos {
	gupp.opts.Limit = limit
	return gupp
}

// Timeout sets a custom timeout for this request.
func (gupp *GetUserProfilePhotos) Timeout(duration time.Duration) *GetUserProfilePhotos {
	if gupp.opts.RequestOpts == nil {
		gupp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gupp.opts.RequestOpts.Timeout = duration

	return gupp
}

// APIURL sets a custom API URL for this request.
func (gupp *GetUserProfilePhotos) APIURL(url g.String) *GetUserProfilePhotos {
	if gupp.opts.RequestOpts == nil {
		gupp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gupp.opts.RequestOpts.APIURL = url.Std()

	return gupp
}

// Send gets user profile photos and returns the result.
func (gupp *GetUserProfilePhotos) Send() g.Result[*gotgbot.UserProfilePhotos] {
	return g.ResultOf(gupp.ctx.Bot.Raw().GetUserProfilePhotos(gupp.userID, gupp.opts))
}
