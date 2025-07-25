package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/types/chataction"
)

type SendChatAction struct {
	ctx    *Context
	action string
	opts   *gotgbot.SendChatActionOpts
	chatID Option[int64]
}

// To sets the target chat ID for the chat action.
func (sca *SendChatAction) To(chatID int64) *SendChatAction {
	sca.chatID = Some(chatID)
	return sca
}

// Thread sets the message thread ID for the chat action.
func (sca *SendChatAction) Thread(id int64) *SendChatAction {
	sca.opts.MessageThreadId = id
	return sca
}

// Business sets the business connection ID for the chat action.
func (sca *SendChatAction) Business(id String) *SendChatAction {
	sca.opts.BusinessConnectionId = id.Std()
	return sca
}

// Action sets a custom chat action type.
func (sca *SendChatAction) Action(action chataction.ChatAction) *SendChatAction {
	sca.action = action.String()
	return sca
}

// Typing sets the chat action to typing.
func (sca *SendChatAction) Typing() *SendChatAction {
	return sca.Action(chataction.Typing)
}

// UploadPhoto sets the chat action to uploading photo.
func (sca *SendChatAction) UploadPhoto() *SendChatAction {
	return sca.Action(chataction.UploadPhoto)
}

// RecordVideo sets the chat action to recording video.
func (sca *SendChatAction) RecordVideo() *SendChatAction {
	return sca.Action(chataction.RecordVideo)
}

// UploadVideo sets the chat action to uploading video.
func (sca *SendChatAction) UploadVideo() *SendChatAction {
	return sca.Action(chataction.UploadVideo)
}

// RecordVoice sets the chat action to recording voice message.
func (sca *SendChatAction) RecordVoice() *SendChatAction {
	return sca.Action(chataction.RecordVoice)
}

// UploadVoice sets the chat action to uploading voice message.
func (sca *SendChatAction) UploadVoice() *SendChatAction {
	return sca.Action(chataction.UploadVoice)
}

// UploadDocument sets the chat action to uploading document.
func (sca *SendChatAction) UploadDocument() *SendChatAction {
	return sca.Action(chataction.UploadDocument)
}

// ChooseSticker sets the chat action to choosing sticker.
func (sca *SendChatAction) ChooseSticker() *SendChatAction {
	return sca.Action(chataction.ChooseSticker)
}

// FindLocation sets the chat action to finding location.
func (sca *SendChatAction) FindLocation() *SendChatAction {
	return sca.Action(chataction.FindLocation)
}

// RecordVideoNote sets the chat action to recording video note.
func (sca *SendChatAction) RecordVideoNote() *SendChatAction {
	return sca.Action(chataction.RecordVideoNote)
}

// UploadVideoNote sets the chat action to uploading video note.
func (sca *SendChatAction) UploadVideoNote() *SendChatAction {
	return sca.Action(chataction.UploadVideoNote)
}

// Timeout sets a custom timeout for this request.
func (sca *SendChatAction) Timeout(duration time.Duration) *SendChatAction {
	if sca.opts.RequestOpts == nil {
		sca.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sca.opts.RequestOpts.Timeout = duration

	return sca
}

// APIURL sets a custom API URL for this request.
func (sca *SendChatAction) APIURL(url String) *SendChatAction {
	if sca.opts.RequestOpts == nil {
		sca.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sca.opts.RequestOpts.APIURL = url.Std()

	return sca
}

// Send sends the chat action to Telegram and returns the result.
func (sca *SendChatAction) Send() Result[bool] {
	chatID := sca.chatID.UnwrapOr(sca.ctx.EffectiveChat.Id)
	return ResultOf(sca.ctx.Bot.Raw().SendChatAction(chatID, sca.action, sca.opts))
}
