// Package permissions defines Telegram chat permission types and utilities.
package permissions

import "github.com/PaulSonOfLars/gotgbot/v2"

// Permission enumerates all supported Telegram chat permission types.
type Permission int

const (
	SendMessages       Permission = iota // Send text messages
	SendAudios                           // Send audio files
	SendDocuments                        // Send documents
	SendPhotos                           // Send photos
	SendVideos                           // Send videos
	SendVideoNotes                       // Send video notes
	SendVoiceNotes                       // Send voice notes
	SendPolls                            // Send polls
	SendOtherMessages                    // Send other messages (stickers, animations)
	AddWebPagePreviews                   // Add web page previews
	ChangeInfo                           // Change chat info
	InviteUsers                          // Invite users to chat
	PinMessages                          // Pin messages
	ManageTopics                         // Manage topics in forum chats
)

// Permissions creates a ChatPermissions object with the specified permissions enabled.
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
