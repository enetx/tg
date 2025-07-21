package ctx

import (
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
func (ca *ChatAction) To(chatID int64) *ChatAction {
	ca.chatID = Some(chatID)
	return ca
}

// Thread sets the message thread ID for the chat action.
func (ca *ChatAction) Thread(id int64) *ChatAction {
	ca.opts.MessageThreadId = id
	return ca
}

// Business sets the business connection ID for the chat action.
func (ca *ChatAction) Business(id String) *ChatAction {
	ca.opts.BusinessConnectionId = id.Std()
	return ca
}

// Action sets a custom chat action type.
func (ca *ChatAction) Action(action chataction.ChatAction) *ChatAction {
	ca.action = action.String()
	return ca
}

// Typing sets the chat action to typing.
func (ca *ChatAction) Typing() *ChatAction {
	return ca.Action(chataction.Typing)
}

// UploadPhoto sets the chat action to uploading photo.
func (ca *ChatAction) UploadPhoto() *ChatAction {
	return ca.Action(chataction.UploadPhoto)
}

// RecordVideo sets the chat action to recording video.
func (ca *ChatAction) RecordVideo() *ChatAction {
	return ca.Action(chataction.RecordVideo)
}

// UploadVideo sets the chat action to uploading video.
func (ca *ChatAction) UploadVideo() *ChatAction {
	return ca.Action(chataction.UploadVideo)
}

// RecordVoice sets the chat action to recording voice message.
func (ca *ChatAction) RecordVoice() *ChatAction {
	return ca.Action(chataction.RecordVoice)
}

// UploadVoice sets the chat action to uploading voice message.
func (ca *ChatAction) UploadVoice() *ChatAction {
	return ca.Action(chataction.UploadVoice)
}

// UploadDocument sets the chat action to uploading document.
func (ca *ChatAction) UploadDocument() *ChatAction {
	return ca.Action(chataction.UploadDocument)
}

// ChooseSticker sets the chat action to choosing sticker.
func (ca *ChatAction) ChooseSticker() *ChatAction {
	return ca.Action(chataction.ChooseSticker)
}

// FindLocation sets the chat action to finding location.
func (ca *ChatAction) FindLocation() *ChatAction {
	return ca.Action(chataction.FindLocation)
}

// RecordVideoNote sets the chat action to recording video note.
func (ca *ChatAction) RecordVideoNote() *ChatAction {
	return ca.Action(chataction.RecordVideoNote)
}

// UploadVideoNote sets the chat action to uploading video note.
func (ca *ChatAction) UploadVideoNote() *ChatAction {
	return ca.Action(chataction.UploadVideoNote)
}

// Send sends the chat action to Telegram and returns the result.
func (ca *ChatAction) Send() Result[bool] {
	chatID := ca.chatID.UnwrapOr(ca.ctx.EffectiveChat.Id)
	return ResultOf(ca.ctx.Bot.Raw().SendChatAction(chatID, ca.action, ca.opts))
}
