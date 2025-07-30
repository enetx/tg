// Package chataction defines supported chat actions for Telegram bots.
//
// These actions are used with the sendChatAction method to indicate that the bot is performing an action,
// such as typing, uploading media, or recording audio.
//
// See: https://core.telegram.org/bots/api#sendchataction
package chataction

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

// g.String returns the canonical string representation used by Telegram Bot API.
func (c ChatAction) String() string {
	switch c {
	case Typing:
		return "typing"
	case UploadPhoto:
		return "upload_photo"
	case RecordVideo:
		return "record_video"
	case UploadVideo:
		return "upload_video"
	case RecordVoice:
		return "record_voice"
	case UploadVoice:
		return "upload_voice"
	case UploadDocument:
		return "upload_document"
	case ChooseSticker:
		return "choose_sticker"
	case FindLocation:
		return "find_location"
	case RecordVideoNote:
		return "record_video_note"
	case UploadVideoNote:
		return "upload_video_note"
	default:
		return "unknown"
	}
}
