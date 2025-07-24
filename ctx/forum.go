package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/g/ref"
)

// CreateForumTopic represents a request to create a forum topic.
type CreateForumTopic struct {
	ctx    *Context
	name   String
	opts   *gotgbot.CreateForumTopicOpts
	chatID Option[int64]
}

// IconColor sets the color of the topic icon in RGB format.
func (cf *CreateForumTopic) IconColor(color int64) *CreateForumTopic {
	cf.opts.IconColor = color
	return cf
}

// IconCustomEmojiID sets the unique identifier of the custom emoji.
func (cf *CreateForumTopic) IconCustomEmojiID(emojiID String) *CreateForumTopic {
	cf.opts.IconCustomEmojiId = emojiID.Std()
	return cf
}

// Timeout sets a custom timeout for this request.
func (cf *CreateForumTopic) Timeout(duration time.Duration) *CreateForumTopic {
	if cf.opts.RequestOpts == nil {
		cf.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	cf.opts.RequestOpts.Timeout = duration

	return cf
}

// APIURL sets a custom API URL for this request.
func (cf *CreateForumTopic) APIURL(url String) *CreateForumTopic {
	if cf.opts.RequestOpts == nil {
		cf.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	cf.opts.RequestOpts.APIURL = url.Std()

	return cf
}

// ChatID sets the target chat ID for this request.
func (cf *CreateForumTopic) ChatID(id int64) *CreateForumTopic {
	cf.chatID = Some(id)
	return cf
}

// Send executes the CreateForumTopic request.
func (cf *CreateForumTopic) Send() Result[*gotgbot.ForumTopic] {
	chatID := cf.chatID.UnwrapOr(cf.ctx.EffectiveChat.Id)
	return ResultOf(cf.ctx.Bot.Raw().CreateForumTopic(chatID, cf.name.Std(), cf.opts))
}

// EditForumTopic represents a request to edit a forum topic.
type EditForumTopic struct {
	ctx             *Context
	messageThreadID int64
	opts            *gotgbot.EditForumTopicOpts
	chatID          Option[int64]
}

// Name sets the new name of the topic.
func (eft *EditForumTopic) Name(name String) *EditForumTopic {
	eft.opts.Name = name.Std()
	return eft
}

// IconCustomEmojiID sets the new custom emoji identifier.
func (eft *EditForumTopic) IconCustomEmojiID(emojiID String) *EditForumTopic {
	eft.opts.IconCustomEmojiId = ref.Of(emojiID.Std())
	return eft
}

// Timeout sets a custom timeout for this request.
func (eft *EditForumTopic) Timeout(duration time.Duration) *EditForumTopic {
	if eft.opts.RequestOpts == nil {
		eft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	eft.opts.RequestOpts.Timeout = duration

	return eft
}

// APIURL sets a custom API URL for this request.
func (eft *EditForumTopic) APIURL(url String) *EditForumTopic {
	if eft.opts.RequestOpts == nil {
		eft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	eft.opts.RequestOpts.APIURL = url.Std()

	return eft
}

// ChatID sets the target chat ID for this request.
func (eft *EditForumTopic) ChatID(id int64) *EditForumTopic {
	eft.chatID = Some(id)
	return eft
}

// Send executes the EditForumTopic request.
func (eft *EditForumTopic) Send() Result[bool] {
	chatID := eft.chatID.UnwrapOr(eft.ctx.EffectiveChat.Id)
	return ResultOf(eft.ctx.Bot.Raw().EditForumTopic(chatID, eft.messageThreadID, eft.opts))
}

// CloseForumTopic represents a request to close a forum topic.
type CloseForumTopic struct {
	ctx             *Context
	messageThreadID int64
	opts            *gotgbot.CloseForumTopicOpts
	chatID          Option[int64]
}

// Timeout sets a custom timeout for this request.
func (cft *CloseForumTopic) Timeout(duration time.Duration) *CloseForumTopic {
	if cft.opts.RequestOpts == nil {
		cft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	cft.opts.RequestOpts.Timeout = duration

	return cft
}

// APIURL sets a custom API URL for this request.
func (cft *CloseForumTopic) APIURL(url String) *CloseForumTopic {
	if cft.opts.RequestOpts == nil {
		cft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	cft.opts.RequestOpts.APIURL = url.Std()

	return cft
}

// ChatID sets the target chat ID for this request.
func (cft *CloseForumTopic) ChatID(id int64) *CloseForumTopic {
	cft.chatID = Some(id)
	return cft
}

// Send executes the CloseForumTopic request.
func (cft *CloseForumTopic) Send() Result[bool] {
	chatID := cft.chatID.UnwrapOr(cft.ctx.EffectiveChat.Id)
	return ResultOf(cft.ctx.Bot.Raw().CloseForumTopic(chatID, cft.messageThreadID, cft.opts))
}

// ReopenForumTopic represents a request to reopen a forum topic.
type ReopenForumTopic struct {
	ctx             *Context
	messageThreadID int64
	opts            *gotgbot.ReopenForumTopicOpts
	chatID          Option[int64]
}

// Timeout sets a custom timeout for this request.
func (rft *ReopenForumTopic) Timeout(duration time.Duration) *ReopenForumTopic {
	if rft.opts.RequestOpts == nil {
		rft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rft.opts.RequestOpts.Timeout = duration

	return rft
}

// APIURL sets a custom API URL for this request.
func (rft *ReopenForumTopic) APIURL(url String) *ReopenForumTopic {
	if rft.opts.RequestOpts == nil {
		rft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rft.opts.RequestOpts.APIURL = url.Std()

	return rft
}

// ChatID sets the target chat ID for this request.
func (rft *ReopenForumTopic) ChatID(id int64) *ReopenForumTopic {
	rft.chatID = Some(id)
	return rft
}

// Send executes the ReopenForumTopic request.
func (rft *ReopenForumTopic) Send() Result[bool] {
	chatID := rft.chatID.UnwrapOr(rft.ctx.EffectiveChat.Id)
	return ResultOf(rft.ctx.Bot.Raw().ReopenForumTopic(chatID, rft.messageThreadID, rft.opts))
}

// DeleteForumTopic represents a request to delete a forum topic.
type DeleteForumTopic struct {
	ctx             *Context
	messageThreadID int64
	opts            *gotgbot.DeleteForumTopicOpts
	chatID          Option[int64]
}

// Timeout sets a custom timeout for this request.
func (dft *DeleteForumTopic) Timeout(duration time.Duration) *DeleteForumTopic {
	if dft.opts.RequestOpts == nil {
		dft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dft.opts.RequestOpts.Timeout = duration

	return dft
}

// APIURL sets a custom API URL for this request.
func (dft *DeleteForumTopic) APIURL(url String) *DeleteForumTopic {
	if dft.opts.RequestOpts == nil {
		dft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dft.opts.RequestOpts.APIURL = url.Std()

	return dft
}

// ChatID sets the target chat ID for this request.
func (dft *DeleteForumTopic) ChatID(id int64) *DeleteForumTopic {
	dft.chatID = Some(id)
	return dft
}

// Send executes the DeleteForumTopic request.
func (dft *DeleteForumTopic) Send() Result[bool] {
	chatID := dft.chatID.UnwrapOr(dft.ctx.EffectiveChat.Id)
	return ResultOf(dft.ctx.Bot.Raw().DeleteForumTopic(chatID, dft.messageThreadID, dft.opts))
}

// EditGeneralForumTopic represents a request to edit the general forum topic.
type EditGeneralForumTopic struct {
	ctx    *Context
	name   String
	opts   *gotgbot.EditGeneralForumTopicOpts
	chatID Option[int64]
}

// Timeout sets a custom timeout for this request.
func (egft *EditGeneralForumTopic) Timeout(duration time.Duration) *EditGeneralForumTopic {
	if egft.opts.RequestOpts == nil {
		egft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	egft.opts.RequestOpts.Timeout = duration

	return egft
}

// APIURL sets a custom API URL for this request.
func (egft *EditGeneralForumTopic) APIURL(url String) *EditGeneralForumTopic {
	if egft.opts.RequestOpts == nil {
		egft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	egft.opts.RequestOpts.APIURL = url.Std()

	return egft
}

// ChatID sets the target chat ID for this request.
func (egft *EditGeneralForumTopic) ChatID(id int64) *EditGeneralForumTopic {
	egft.chatID = Some(id)
	return egft
}

// Send executes the EditGeneralForumTopic request.
func (egft *EditGeneralForumTopic) Send() Result[bool] {
	chatID := egft.chatID.UnwrapOr(egft.ctx.EffectiveChat.Id)
	return ResultOf(egft.ctx.Bot.Raw().EditGeneralForumTopic(chatID, egft.name.Std(), egft.opts))
}

// CloseGeneralForumTopic represents a request to close the general forum topic.
type CloseGeneralForumTopic struct {
	ctx    *Context
	opts   *gotgbot.CloseGeneralForumTopicOpts
	chatID Option[int64]
}

// Timeout sets a custom timeout for this request.
func (cgft *CloseGeneralForumTopic) Timeout(duration time.Duration) *CloseGeneralForumTopic {
	if cgft.opts.RequestOpts == nil {
		cgft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	cgft.opts.RequestOpts.Timeout = duration

	return cgft
}

// APIURL sets a custom API URL for this request.
func (cgft *CloseGeneralForumTopic) APIURL(url String) *CloseGeneralForumTopic {
	if cgft.opts.RequestOpts == nil {
		cgft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	cgft.opts.RequestOpts.APIURL = url.Std()

	return cgft
}

// ChatID sets the target chat ID for this request.
func (cgft *CloseGeneralForumTopic) ChatID(id int64) *CloseGeneralForumTopic {
	cgft.chatID = Some(id)
	return cgft
}

// Send executes the CloseGeneralForumTopic request.
func (cgft *CloseGeneralForumTopic) Send() Result[bool] {
	chatID := cgft.chatID.UnwrapOr(cgft.ctx.EffectiveChat.Id)
	return ResultOf(cgft.ctx.Bot.Raw().CloseGeneralForumTopic(chatID, cgft.opts))
}
