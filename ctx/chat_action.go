package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/types/chataction"
)

type ChatAction struct {
	ctx    *Context
	action string
	opts   *gotgbot.SendChatActionOpts
	chatID Option[int64]
}

// To sets the target chat ID for the chat action.
func (c *ChatAction) To(chatID int64) *ChatAction {
	c.chatID = Some(chatID)
	return c
}

// Thread sets the message thread ID for the chat action.
func (c *ChatAction) Thread(id int64) *ChatAction {
	c.opts.MessageThreadId = id
	return c
}

// Business sets the business connection ID for the chat action.
func (c *ChatAction) Business(id String) *ChatAction {
	c.opts.BusinessConnectionId = id.Std()
	return c
}

// Action sets a custom chat action type.
func (c *ChatAction) Action(action chataction.ChatAction) *ChatAction {
	c.action = action.String()
	return c
}

// Typing sets the chat action to typing.
func (c *ChatAction) Typing() *ChatAction {
	return c.Action(chataction.Typing)
}

// UploadPhoto sets the chat action to uploading photo.
func (c *ChatAction) UploadPhoto() *ChatAction {
	return c.Action(chataction.UploadPhoto)
}

// RecordVideo sets the chat action to recording video.
func (c *ChatAction) RecordVideo() *ChatAction {
	return c.Action(chataction.RecordVideo)
}

// UploadVideo sets the chat action to uploading video.
func (c *ChatAction) UploadVideo() *ChatAction {
	return c.Action(chataction.UploadVideo)
}

// RecordVoice sets the chat action to recording voice message.
func (c *ChatAction) RecordVoice() *ChatAction {
	return c.Action(chataction.RecordVoice)
}

// UploadVoice sets the chat action to uploading voice message.
func (c *ChatAction) UploadVoice() *ChatAction {
	return c.Action(chataction.UploadVoice)
}

// UploadDocument sets the chat action to uploading document.
func (c *ChatAction) UploadDocument() *ChatAction {
	return c.Action(chataction.UploadDocument)
}

// ChooseSticker sets the chat action to choosing sticker.
func (c *ChatAction) ChooseSticker() *ChatAction {
	return c.Action(chataction.ChooseSticker)
}

// FindLocation sets the chat action to finding location.
func (c *ChatAction) FindLocation() *ChatAction {
	return c.Action(chataction.FindLocation)
}

// RecordVideoNote sets the chat action to recording video note.
func (c *ChatAction) RecordVideoNote() *ChatAction {
	return c.Action(chataction.RecordVideoNote)
}

// UploadVideoNote sets the chat action to uploading video note.
func (c *ChatAction) UploadVideoNote() *ChatAction {
	return c.Action(chataction.UploadVideoNote)
}

// Timeout sets a custom timeout for this request.
func (c *ChatAction) Timeout(duration time.Duration) *ChatAction {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *ChatAction) APIURL(url String) *ChatAction {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Send sends the chat action to Telegram and returns the result.
func (c *ChatAction) Send() Result[bool] {
	chatID := c.chatID.UnwrapOr(c.ctx.EffectiveChat.Id)
	return ResultOf(c.ctx.Bot.Raw().SendChatAction(chatID, c.action, c.opts))
}
