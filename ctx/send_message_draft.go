package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
)

// SendMessageDraft represents a request to send a message draft.
// Message drafts allow partial messages to be streamed to a user while being generated.
type SendMessageDraft struct {
	ctx     *Context
	chatID  g.Option[int64]
	draftID int64
	text    g.String
	opts    *gotgbot.SendMessageDraftOpts
}

// To sets the target chat ID for the message draft.
func (smd *SendMessageDraft) To(chatID int64) *SendMessageDraft {
	smd.chatID = g.Some(chatID)
	return smd
}

// Thread sets the message thread ID for the message draft.
func (smd *SendMessageDraft) Thread(id int64) *SendMessageDraft {
	smd.opts.MessageThreadId = id
	return smd
}

// HTML sets the message parse mode to HTML.
func (smd *SendMessageDraft) HTML() *SendMessageDraft {
	smd.opts.ParseMode = "HTML"
	return smd
}

// Markdown sets the message parse mode to MarkdownV2.
func (smd *SendMessageDraft) Markdown() *SendMessageDraft {
	smd.opts.ParseMode = "MarkdownV2"
	return smd
}

// Entities sets special entities in the message text using Entities builder.
func (smd *SendMessageDraft) Entities(e *entities.Entities) *SendMessageDraft {
	smd.opts.Entities = e.Std()
	return smd
}

// Timeout sets a custom timeout for this request.
func (smd *SendMessageDraft) Timeout(duration time.Duration) *SendMessageDraft {
	if smd.opts.RequestOpts == nil {
		smd.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	smd.opts.RequestOpts.Timeout = duration

	return smd
}

// APIURL sets a custom API URL for this request.
func (smd *SendMessageDraft) APIURL(url g.String) *SendMessageDraft {
	if smd.opts.RequestOpts == nil {
		smd.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	smd.opts.RequestOpts.APIURL = url.Std()

	return smd
}

// Send sends the message draft to Telegram and returns the result.
func (smd *SendMessageDraft) Send() g.Result[bool] {
	chatID := smd.chatID.UnwrapOr(smd.ctx.EffectiveChat.Id)
	return g.ResultOf(smd.ctx.Bot.Raw().SendMessageDraft(
		chatID,
		smd.draftID,
		smd.text.Std(),
		smd.opts,
	))
}
