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
}

func TestRoles_Empty(t *testing.T) {
	result := roles.Roles()
	if result == nil {
		t.Error("Expected non-nil PromoteChatMemberOpts")
	}
}
