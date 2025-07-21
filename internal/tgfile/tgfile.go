package tgfile

import (
	"errors"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/constants"
)

// tgFile contains the result of file processing for Telegram Bot API.
// It holds both the gotgbot input file representation and the original file handle.
type tgFile struct {
	Doc  gotgbot.InputFileOrString
	File *File
}

// ProcessFile handles the common logic for file processing across all media types.
// It supports three types of file inputs:
//   - HTTP/HTTPS URLs: processed as remote files
//   - File IDs (prefixed with constants.FileIDPrefix): processed as existing Telegram files
//   - Local file paths: opened and processed as file uploads
//
// Returns a tgFile containing the processed file data or an error if processing fails.
func ProcessFile(filename String) Result[tgFile] {
	if filename.Empty() {
		return Err[tgFile](errors.New("filename is empty"))
	}

	switch {
	case filename.StartsWithAny("http://", "https://"):
		return Ok(tgFile{Doc: gotgbot.InputFileByURL(filename.Std())})
	case filename.StartsWith(constants.FileIDPrefix):
		return Ok(tgFile{Doc: gotgbot.InputFileByID(filename.StripPrefix(constants.FileIDPrefix).Std())})
	default:
		file := NewFile(filename).Open()
		if file.IsErr() {
			return Err[tgFile](file.Err())
		}

		doc := gotgbot.InputFileByReader(file.Ok().Name().Std(), file.Ok().Std())
		return Ok(tgFile{Doc: doc, File: file.Ok()})
	}
}
