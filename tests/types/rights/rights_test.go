package rights_test

import (
	"testing"

	"github.com/enetx/tg/types/rights"
)

func TestRights_Single(t *testing.T) {
	result := rights.Rights(rights.ManageChat)
	if result == nil {
		t.Error("Expected non-nil ChatAdministratorRights")
	}
	if !result.CanManageChat {
		t.Error("Expected CanManageChat to be true")
	}
}

func TestRights_Multiple(t *testing.T) {
	result := rights.Rights(
		rights.ManageChat,
		rights.DeleteMessages,
		rights.InviteUsers,
	)
	if result == nil {
		t.Error("Expected non-nil ChatAdministratorRights")
	}
	if !result.CanManageChat {
		t.Error("Expected CanManageChat to be true")
	}
	if !result.CanDeleteMessages {
		t.Error("Expected CanDeleteMessages to be true")
	}
	if !result.CanInviteUsers {
		t.Error("Expected CanInviteUsers to be true")
	}
}

func TestRights_Empty(t *testing.T) {
	result := rights.Rights()
	if result == nil {
		t.Error("Expected non-nil ChatAdministratorRights")
	}
	// All rights should be false for empty list
	if result.CanManageChat {
		t.Error("Expected CanManageChat to be false for empty rights")
	}
}

// Test all individual rights
func TestRights_Anonymous(t *testing.T) {
	result := rights.Rights(rights.Anonymous)
	if !result.IsAnonymous {
		t.Error("Expected IsAnonymous to be true")
	}
}

func TestRights_DeleteMessages(t *testing.T) {
	result := rights.Rights(rights.DeleteMessages)
	if !result.CanDeleteMessages {
		t.Error("Expected CanDeleteMessages to be true")
	}
}

func TestRights_ManageVideoChats(t *testing.T) {
	result := rights.Rights(rights.ManageVideoChats)
	if !result.CanManageVideoChats {
		t.Error("Expected CanManageVideoChats to be true")
	}
}

func TestRights_RestrictMembers(t *testing.T) {
	result := rights.Rights(rights.RestrictMembers)
	if !result.CanRestrictMembers {
		t.Error("Expected CanRestrictMembers to be true")
	}
}

func TestRights_PromoteMembers(t *testing.T) {
	result := rights.Rights(rights.PromoteMembers)
	if !result.CanPromoteMembers {
		t.Error("Expected CanPromoteMembers to be true")
	}
}

func TestRights_ChangeInfo(t *testing.T) {
	result := rights.Rights(rights.ChangeInfo)
	if !result.CanChangeInfo {
		t.Error("Expected CanChangeInfo to be true")
	}
}

func TestRights_InviteUsers(t *testing.T) {
	result := rights.Rights(rights.InviteUsers)
	if !result.CanInviteUsers {
		t.Error("Expected CanInviteUsers to be true")
	}
}

func TestRights_PostMessages(t *testing.T) {
	result := rights.Rights(rights.PostMessages)
	if !result.CanPostMessages {
		t.Error("Expected CanPostMessages to be true")
	}
}

func TestRights_EditMessages(t *testing.T) {
	result := rights.Rights(rights.EditMessages)
	if !result.CanEditMessages {
		t.Error("Expected CanEditMessages to be true")
	}
}

func TestRights_PinMessages(t *testing.T) {
	result := rights.Rights(rights.PinMessages)
	if !result.CanPinMessages {
		t.Error("Expected CanPinMessages to be true")
	}
}

func TestRights_PostStories(t *testing.T) {
	result := rights.Rights(rights.PostStories)
	if !result.CanPostStories {
		t.Error("Expected CanPostStories to be true")
	}
}

func TestRights_EditStories(t *testing.T) {
	result := rights.Rights(rights.EditStories)
	if !result.CanEditStories {
		t.Error("Expected CanEditStories to be true")
	}
}

func TestRights_DeleteStories(t *testing.T) {
	result := rights.Rights(rights.DeleteStories)
	if !result.CanDeleteStories {
		t.Error("Expected CanDeleteStories to be true")
	}
}

func TestRights_ManageTopics(t *testing.T) {
	result := rights.Rights(rights.ManageTopics)
	if !result.CanManageTopics {
		t.Error("Expected CanManageTopics to be true")
	}
}

// Test all rights combined
func TestRights_AllRights(t *testing.T) {
	result := rights.Rights(
		rights.Anonymous,
		rights.ManageChat,
		rights.DeleteMessages,
		rights.ManageVideoChats,
		rights.RestrictMembers,
		rights.PromoteMembers,
		rights.ChangeInfo,
		rights.InviteUsers,
		rights.PostMessages,
		rights.EditMessages,
		rights.PinMessages,
		rights.PostStories,
		rights.EditStories,
		rights.DeleteStories,
		rights.ManageTopics,
	)

	if !result.IsAnonymous {
		t.Error("Expected IsAnonymous to be true")
	}
	if !result.CanManageChat {
		t.Error("Expected CanManageChat to be true")
	}
	if !result.CanDeleteMessages {
		t.Error("Expected CanDeleteMessages to be true")
	}
	if !result.CanManageVideoChats {
		t.Error("Expected CanManageVideoChats to be true")
	}
	if !result.CanRestrictMembers {
		t.Error("Expected CanRestrictMembers to be true")
	}
	if !result.CanPromoteMembers {
		t.Error("Expected CanPromoteMembers to be true")
	}
	if !result.CanChangeInfo {
		t.Error("Expected CanChangeInfo to be true")
	}
	if !result.CanInviteUsers {
		t.Error("Expected CanInviteUsers to be true")
	}
	if !result.CanPostMessages {
		t.Error("Expected CanPostMessages to be true")
	}
	if !result.CanEditMessages {
		t.Error("Expected CanEditMessages to be true")
	}
	if !result.CanPinMessages {
		t.Error("Expected CanPinMessages to be true")
	}
	if !result.CanPostStories {
		t.Error("Expected CanPostStories to be true")
	}
	if !result.CanEditStories {
		t.Error("Expected CanEditStories to be true")
	}
	if !result.CanDeleteStories {
		t.Error("Expected CanDeleteStories to be true")
	}
	if !result.CanManageTopics {
		t.Error("Expected CanManageTopics to be true")
	}
}
