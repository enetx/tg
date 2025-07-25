package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// GetForumTopicIconStickers represents a request to get forum topic icon stickers.
type GetForumTopicIconStickers struct {
	ctx  *Context
	opts *gotgbot.GetForumTopicIconStickersOpts
}

// Timeout sets a custom timeout for this request.
func (gftis *GetForumTopicIconStickers) Timeout(duration time.Duration) *GetForumTopicIconStickers {
	if gftis.opts.RequestOpts == nil {
		gftis.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gftis.opts.RequestOpts.Timeout = duration

	return gftis
}

// APIURL sets a custom API URL for this request.
func (gftis *GetForumTopicIconStickers) APIURL(url String) *GetForumTopicIconStickers {
	if gftis.opts.RequestOpts == nil {
		gftis.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gftis.opts.RequestOpts.APIURL = url.Std()

	return gftis
}

// Send gets the custom emoji stickers that can be used as forum topic icons.
func (gftis *GetForumTopicIconStickers) Send() Result[Slice[gotgbot.Sticker]] {
	return ResultOf[Slice[gotgbot.Sticker]](gftis.ctx.Bot.Raw().GetForumTopicIconStickers(gftis.opts))
}
