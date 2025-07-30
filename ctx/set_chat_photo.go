package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// SetChatPhoto represents a request to set the chat photo.
type SetChatPhoto struct {
	ctx    *Context
	opts   *gotgbot.SetChatPhotoOpts
	doc    gotgbot.InputFile
	file   *g.File
	chatID g.Option[int64]
	err    error
}

// ChatID sets the target chat ID for this request.
func (scp *SetChatPhoto) ChatID(id int64) *SetChatPhoto {
	scp.chatID = g.Some(id)
	return scp
}

// Timeout sets a custom timeout for this request.
func (scp *SetChatPhoto) Timeout(duration time.Duration) *SetChatPhoto {
	if scp.opts.RequestOpts == nil {
		scp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	scp.opts.RequestOpts.Timeout = duration

	return scp
}

// APIURL sets a custom API URL for this request.
func (scp *SetChatPhoto) APIURL(url g.String) *SetChatPhoto {
	if scp.opts.RequestOpts == nil {
		scp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	scp.opts.RequestOpts.APIURL = url.Std()

	return scp
}

// Send executes the SetChatPhoto request.
func (scp *SetChatPhoto) Send() g.Result[bool] {
	if scp.err != nil {
		return g.Err[bool](scp.err)
	}

	if scp.file != nil {
		defer scp.file.Close()
	}

	chatID := scp.chatID.UnwrapOr(scp.ctx.EffectiveChat.Id)

	return g.ResultOf(scp.ctx.Bot.Raw().SetChatPhoto(chatID, scp.doc, scp.opts))
}
