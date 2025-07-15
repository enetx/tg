package permissions

import "github.com/PaulSonOfLars/gotgbot/v2"

type Permission int

const (
	SendMessages Permission = iota
	SendAudios
	SendDocuments
	SendPhotos
	SendVideos
	SendVideoNotes
	SendVoiceNotes
	SendPolls
	SendOtherMessages
	AddWebPagePreviews
	ChangeInfo
	InviteUsers
	PinMessages
	ManageTopics
)

func Permissions(perms ...Permission) *gotgbot.ChatPermissions {
	cp := new(gotgbot.ChatPermissions)
	for _, p := range perms {
		switch p {
		case SendMessages:
			cp.CanSendMessages = true
		case SendAudios:
			cp.CanSendAudios = true
		case SendDocuments:
			cp.CanSendDocuments = true
		case SendPhotos:
			cp.CanSendPhotos = true
		case SendVideos:
			cp.CanSendVideos = true
		case SendVideoNotes:
			cp.CanSendVideoNotes = true
		case SendVoiceNotes:
			cp.CanSendVoiceNotes = true
		case SendPolls:
			cp.CanSendPolls = true
		case SendOtherMessages:
			cp.CanSendOtherMessages = true
		case AddWebPagePreviews:
			cp.CanAddWebPagePreviews = true
		case ChangeInfo:
			cp.CanChangeInfo = true
		case InviteUsers:
			cp.CanInviteUsers = true
		case PinMessages:
			cp.CanPinMessages = true
		case ManageTopics:
			cp.CanManageTopics = true
		}
	}

	return cp
}
