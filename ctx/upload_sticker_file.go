package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// UploadStickerFile represents a request to upload a sticker file.
type UploadStickerFile struct {
	ctx           *Context
	userID        int64
	sticker       gotgbot.InputFile
	stickerFormat String
	opts          *gotgbot.UploadStickerFileOpts
	file          *File
	err           error
}

// File sets the sticker file to upload.
func (usf *UploadStickerFile) File(filename String) *UploadStickerFile {
	usf.file = NewFile(filename)

	reader := usf.file.Open()
	if reader.IsErr() {
		usf.err = reader.Err()
		return usf
	}

	usf.sticker = gotgbot.InputFileByReader(usf.file.Name().Std(), reader.Ok().Std())
	return usf
}

// Format sets the sticker format.
func (usf *UploadStickerFile) Format(format String) *UploadStickerFile {
	usf.stickerFormat = format
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
func (usf *UploadStickerFile) APIURL(url String) *UploadStickerFile {
	if usf.opts.RequestOpts == nil {
		usf.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	usf.opts.RequestOpts.APIURL = url.Std()

	return usf
}

// Send uploads the sticker file.
func (usf *UploadStickerFile) Send() Result[*gotgbot.File] {
	if usf.err != nil {
		return Err[*gotgbot.File](usf.err)
	}

	if usf.file != nil {
		defer usf.file.Close()
	}

	return ResultOf(usf.ctx.Bot.Raw().UploadStickerFile(usf.userID, usf.sticker, usf.stickerFormat.Std(), usf.opts))
}
