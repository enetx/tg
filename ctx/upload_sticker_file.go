package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// UploadStickerFile represents a request to upload a sticker file.
type UploadStickerFile struct {
	ctx     *Context
	userID  int64
	sticker gotgbot.InputFile
	format  g.String
	opts    *gotgbot.UploadStickerFileOpts
	file    *g.File
	err     error
}

// File sets the sticker file to upload.
func (usf *UploadStickerFile) File(filename g.String) *UploadStickerFile {
	usf.file = g.NewFile(filename)

	reader := usf.file.Open()
	if reader.IsErr() {
		usf.err = reader.Err()
		return usf
	}

	usf.sticker = gotgbot.InputFileByReader(usf.file.Name().Std(), reader.Ok().Std())
	return usf
}

// Format sets the sticker format.
func (usf *UploadStickerFile) Format(format g.String) *UploadStickerFile {
	usf.format = format
	return usf
}

// Timeout sets a custom timeout for this request.
func (usf *UploadStickerFile) Timeout(duration time.Duration) *UploadStickerFile {
	if usf.opts.RequestOpts == nil {
		usf.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	usf.opts.RequestOpts.Timeout = duration

	return usf
}

// APIURL sets a custom API URL for this request.
func (usf *UploadStickerFile) APIURL(url g.String) *UploadStickerFile {
	if usf.opts.RequestOpts == nil {
		usf.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	usf.opts.RequestOpts.APIURL = url.Std()

	return usf
}

// Send uploads the sticker file.
func (usf *UploadStickerFile) Send() g.Result[*gotgbot.File] {
	if usf.err != nil {
		return g.Err[*gotgbot.File](usf.err)
	}

	if usf.file != nil {
		defer usf.file.Close()
	}

	return g.ResultOf(usf.ctx.Bot.Raw().UploadStickerFile(usf.userID, usf.sticker, usf.format.Std(), usf.opts))
}
