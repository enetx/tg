package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// GetFile represents a request to get a file.
type GetFile struct {
	ctx    *Context
	fileID g.String
	opts   *gotgbot.GetFileOpts
}

// Timeout sets a custom timeout for this request.
func (gf *GetFile) Timeout(duration time.Duration) *GetFile {
	if gf.opts.RequestOpts == nil {
		gf.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gf.opts.RequestOpts.Timeout = duration

	return gf
}

// APIURL sets a custom API URL for this request.
func (gf *GetFile) APIURL(url g.String) *GetFile {
	if gf.opts.RequestOpts == nil {
		gf.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gf.opts.RequestOpts.APIURL = url.Std()

	return gf
}

// Send gets the file and returns the result.
func (gf *GetFile) Send() g.Result[*gotgbot.File] {
	return g.ResultOf(gf.ctx.Bot.Raw().GetFile(gf.fileID.Std(), gf.opts))
}
