package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// DeleteChatPhoto represents a request to delete the chat photo.
type DeleteChatPhoto struct {
	ctx    *Context
	opts   *gotgbot.DeleteChatPhotoOpts
	chatID g.Option[int64]
}

// ChatID sets the target chat ID for this request.
func (dcp *DeleteChatPhoto) ChatID(id int64) *DeleteChatPhoto {
	dcp.chatID = g.Some(id)
	return dcp
}

// Timeout sets a custom timeout for this request.
func (dcp *DeleteChatPhoto) Timeout(duration time.Duration) *DeleteChatPhoto {
	if dcp.opts.RequestOpts == nil {
		dcp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dcp.opts.RequestOpts.Timeout = duration

	return dcp
}

// APIURL sets a custom API URL for this request.
func (dcp *DeleteChatPhoto) APIURL(url g.String) *DeleteChatPhoto {
	if dcp.opts.RequestOpts == nil {
		dcp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dcp.opts.RequestOpts.APIURL = url.Std()

	return dcp
}

// Send executes the DeleteChatPhoto request.
func (dcp *DeleteChatPhoto) Send() g.Result[bool] {
	chatID := dcp.chatID.UnwrapOr(dcp.ctx.EffectiveChat.Id)
	return g.ResultOf(dcp.ctx.Bot.Raw().DeleteChatPhoto(chatID, dcp.opts))
}
