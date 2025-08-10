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
}

func TestRights_Empty(t *testing.T) {
	result := rights.Rights()
	if result == nil {
		t.Error("Expected non-nil ChatAdministratorRights")
	}
}
