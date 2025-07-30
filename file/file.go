package file

import (
	"errors"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/constants"
)

// InputFile contains the result of file processing for Telegram Bot API.
// It holds both the gotgbot input file representation and the original file handle.
type InputFile struct {
	Doc  gotgbot.InputFileOrString
	File *g.File
}

// Input handles the common logic for file processing across all media types.
// It supports three types of file inputs:
//   - HTTP/HTTPS URLs: processed as remote files
//   - File IDs (prefixed with constants.FileIDPrefix): processed as existing Telegram files
//   - Local file paths: opened and processed as file uploads
//
// Returns a File containing the processed file data or an error if processing fails.
func Input(filename g.String) g.Result[InputFile] {
	if filename.Empty() {
		return g.Err[InputFile](errors.New("filename is empty"))
	}

	switch {
	case filename.StartsWithAny("http://", "https://"):
		return g.Ok(InputFile{Doc: gotgbot.InputFileByURL(filename.Std())})
	case filename.StartsWith(constants.FileIDPrefix):
		return g.Ok(InputFile{Doc: gotgbot.InputFileByID(filename.StripPrefix(constants.FileIDPrefix).Std())})
	default:
		file := g.NewFile(filename).Open()
		if file.IsErr() {
			return g.Err[InputFile](file.Err())
		}

		doc := gotgbot.InputFileByReader(file.Ok().Name().Std(), file.Ok().Std())
		return g.Ok(InputFile{Doc: doc, File: file.Ok()})
	}
}
