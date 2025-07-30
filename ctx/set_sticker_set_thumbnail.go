package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// SetStickerSetThumbnail represents a request to set sticker set thumbnail.
type SetStickerSetThumbnail struct {
	ctx    *Context
	name   g.String
	userID int64
	format g.String
	opts   *gotgbot.SetStickerSetThumbnailOpts
	thumb  *g.File
	err    error
}

// Thumbnail sets the thumbnail file for the sticker set.
func (ssst *SetStickerSetThumbnail) Thumbnail(filename g.String) *SetStickerSetThumbnail {
	ssst.thumb = g.NewFile(filename)

	reader := ssst.thumb.Open()
	if reader.IsErr() {
		ssst.err = reader.Err()
		return ssst
	}

	ssst.opts.Thumbnail = gotgbot.InputFileByReader(ssst.thumb.Name().Std(), reader.Ok().Std())
	return ssst
}

// Format sets the thumbnail format.
// format of the thumbnail, must be one of "static" for a .WEBP or .PNG image,
// "animated" for a .TGS animation, or "video" for a .WEBM video.
func (ssst *SetStickerSetThumbnail) Format(format g.String) *SetStickerSetThumbnail {
	ssst.format = format
	return ssst
}

// Timeout sets a custom timeout for this request.
func (ssst *SetStickerSetThumbnail) Timeout(duration time.Duration) *SetStickerSetThumbnail {
	if ssst.opts.RequestOpts == nil {
		ssst.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ssst.opts.RequestOpts.Timeout = duration

	return ssst
}

// APIURL sets a custom API URL for this request.
func (ssst *SetStickerSetThumbnail) APIURL(url g.String) *SetStickerSetThumbnail {
	if ssst.opts.RequestOpts == nil {
		ssst.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ssst.opts.RequestOpts.APIURL = url.Std()

	return ssst
}

// Send sets the sticker set thumbnail.
func (ssst *SetStickerSetThumbnail) Send() g.Result[bool] {
	if ssst.err != nil {
		return g.Err[bool](ssst.err)
	}

	if ssst.thumb != nil {
		defer ssst.thumb.Close()
	}

	return g.ResultOf(ssst.ctx.Bot.Raw().
		SetStickerSetThumbnail(ssst.name.Std(), ssst.userID, ssst.format.Std(), ssst.opts),
	)
}
