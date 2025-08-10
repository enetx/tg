package roles_test

import (
	"testing"

	"github.com/enetx/tg/types/roles"
)

func TestRoles_Single(t *testing.T) {
	result := roles.Roles(roles.ManageChat)
	if result == nil {
		t.Error("Expected non-nil PromoteChatMemberOpts")
	}
	if !result.CanManageChat {
		t.Error("Expected CanManageChat to be true")
	}
}

func TestRoles_Multiple(t *testing.T) {
	result := roles.Roles(
		roles.ManageChat,
		roles.DeleteMessages,
		roles.InviteUsers,
	)
	if result == nil {
		t.Error("Expected non-nil PromoteChatMemberOpts")
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

func TestRoles_Empty(t *testing.T) {
	result := roles.Roles()
	if result == nil {
		t.Error("Expected non-nil PromoteChatMemberOpts")
	}
	// All roles should be false for empty list
	if result.CanManageChat {
		t.Error("Expected CanManageChat to be false for empty roles")
	}
}

// Test all individual roles
func TestRoles_ManageChat(t *testing.T) {
	result := roles.Roles(roles.ManageChat)
	if !result.CanManageChat {
		t.Error("Expected CanManageChat to be true")
	}
}

func TestRoles_DeleteMessages(t *testing.T) {
	result := roles.Roles(roles.DeleteMessages)
	if !result.CanDeleteMessages {
		t.Error("Expected CanDeleteMessages to be true")
	}
}

func TestRoles_ManageVideoChats(t *testing.T) {
	result := roles.Roles(roles.ManageVideoChats)
	if !result.CanManageVideoChats {
		t.Error("Expected CanManageVideoChats to be true")
	}
}

func TestRoles_RestrictMembers(t *testing.T) {
	result := roles.Roles(roles.RestrictMembers)
	if !result.CanRestrictMembers {
		t.Error("Expected CanRestrictMembers to be true")
	}
}

func TestRoles_PromoteMembers(t *testing.T) {
	result := roles.Roles(roles.PromoteMembers)
	if !result.CanPromoteMembers {
		t.Error("Expected CanPromoteMembers to be true")
	}
}

func TestRoles_ChangeInfo(t *testing.T) {
	result := roles.Roles(roles.ChangeInfo)
	if !result.CanChangeInfo {
		t.Error("Expected CanChangeInfo to be true")
	}
}

func TestRoles_InviteUsers(t *testing.T) {
	result := roles.Roles(roles.InviteUsers)
	if !result.CanInviteUsers {
		t.Error("Expected CanInviteUsers to be true")
	}
}

func TestRoles_PostMessages(t *testing.T) {
	result := roles.Roles(roles.PostMessages)
	if !result.CanPostMessages {
		t.Error("Expected CanPostMessages to be true")
	}
}

func TestRoles_EditMessages(t *testing.T) {
	result := roles.Roles(roles.EditMessages)
	if !result.CanEditMessages {
		t.Error("Expected CanEditMessages to be true")
	}
}

func TestRoles_PinMessages(t *testing.T) {
	result := roles.Roles(roles.PinMessages)
	if !result.CanPinMessages {
		t.Error("Expected CanPinMessages to be true")
	}
}

func TestRoles_ManageTopics(t *testing.T) {
	result := roles.Roles(roles.ManageTopics)
	if !result.CanManageTopics {
		t.Error("Expected CanManageTopics to be true")
	}
}

// Test all roles combined
func TestRoles_AllRoles(t *testing.T) {
	result := roles.Roles(
		roles.ManageChat,
		roles.DeleteMessages,
		roles.ManageVideoChats,
		roles.RestrictMembers,
		roles.PromoteMembers,
		roles.ChangeInfo,
		roles.InviteUsers,
		roles.PostMessages,
		roles.EditMessages,
		roles.PinMessages,
		roles.ManageTopics,
	)

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
	if !result.CanManageTopics {
		t.Error("Expected CanManageTopics to be true")
	}
}
