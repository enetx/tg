// Package chataction defines supported chat actions for Telegram bots.
//
// These actions are used with the sendChatAction method to indicate that the bot is performing an action,
// such as typing, uploading media, or recording audio.
//
// See: https://core.telegram.org/bots/api#sendchataction
package chataction

import "github.com/PaulSonOfLars/gotgbot/v2"

// ChatAction enumerates all supported Telegram chat actions.
type ChatAction int

const (
	Typing          ChatAction = iota // Typing a text message
	UploadPhoto                       // Uploading a photo
	RecordVideo                       // Recording a video
	UploadVideo                       // Uploading a video
	RecordVoice                       // Recording a voice message
	UploadVoice                       // Uploading a voice message
	UploadDocument                    // Uploading a document
	ChooseSticker                     // Choosing a sticker
	FindLocation                      // Finding a location
	RecordVideoNote                   // Recording a video note
	UploadVideoNote                   // Uploading a video note
)

// String returns the canonical string representation used by Telegram Bot API.
func (c ChatAction) String() string {
	switch c {
	case Typing:
		return gotgbot.ChatActionTyping
	case UploadPhoto:
		return gotgbot.ChatActionUploadPhoto
	case RecordVideo:
		return gotgbot.ChatActionRecordVideo
	case UploadVideo:
		return gotgbot.ChatActionUploadVideo
	case RecordVoice:
		return gotgbot.ChatActionRecordVoice
	case UploadVoice:
		return gotgbot.ChatActionUploadVoice
	case UploadDocument:
		return gotgbot.ChatActionUploadDocument
	case ChooseSticker:
		return gotgbot.ChatActionChooseSticker
	case FindLocation:
		return gotgbot.ChatActionFindLocation
	case RecordVideoNote:
		return gotgbot.ChatActionRecordVideoNote
	case UploadVideoNote:
		return gotgbot.ChatActionUploadVideoNote
	default:
		return "unknown"
	}
}
