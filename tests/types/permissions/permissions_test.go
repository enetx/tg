package permissions_test

import (
	"testing"

	"github.com/enetx/tg/types/permissions"
)

func TestPermissions_Single(t *testing.T) {
	tests := []struct {
		name       string
		permission permissions.Permission
		checkFunc  func(*testing.T, any)
	}{
		{
			name:       "SendMessages",
			permission: permissions.SendMessages,
			checkFunc: func(t *testing.T, result any) {
				if result == nil {
					t.Error("Expected non-nil result")
				}
			},
		},
		{
			name:       "SendAudios",
			permission: permissions.SendAudios,
			checkFunc: func(t *testing.T, result any) {
				if result == nil {
					t.Error("Expected non-nil result")
				}
			},
		},
		{
			name:       "SendDocuments",
			permission: permissions.SendDocuments,
			checkFunc: func(t *testing.T, result any) {
				if result == nil {
					t.Error("Expected non-nil result")
				}
			},
		},
		{
			name:       "SendPhotos",
			permission: permissions.SendPhotos,
			checkFunc: func(t *testing.T, result any) {
				if result == nil {
					t.Error("Expected non-nil result")
				}
			},
		},
		{
			name:       "SendVideos",
			permission: permissions.SendVideos,
			checkFunc: func(t *testing.T, result any) {
				if result == nil {
					t.Error("Expected non-nil result")
				}
			},
		},
		{
			name:       "SendVideoNotes",
			permission: permissions.SendVideoNotes,
			checkFunc: func(t *testing.T, result any) {
				if result == nil {
					t.Error("Expected non-nil result")
				}
			},
		},
		{
			name:       "SendVoiceNotes",
			permission: permissions.SendVoiceNotes,
			checkFunc: func(t *testing.T, result any) {
				if result == nil {
					t.Error("Expected non-nil result")
				}
			},
		},
		{
			name:       "SendPolls",
			permission: permissions.SendPolls,
			checkFunc: func(t *testing.T, result any) {
				if result == nil {
					t.Error("Expected non-nil result")
				}
			},
		},
		{
			name:       "SendOtherMessages",
			permission: permissions.SendOtherMessages,
			checkFunc: func(t *testing.T, result any) {
				if result == nil {
					t.Error("Expected non-nil result")
				}
			},
		},
		{
			name:       "AddWebPagePreviews",
			permission: permissions.AddWebPagePreviews,
			checkFunc: func(t *testing.T, result any) {
				if result == nil {
					t.Error("Expected non-nil result")
				}
			},
		},
		{
			name:       "ReactToMessages",
			permission: permissions.ReactToMessages,
			checkFunc: func(t *testing.T, result any) {
				if result == nil {
					t.Error("Expected non-nil result")
				}
			},
		},
		{
			name:       "EditTag",
			permission: permissions.EditTag,
			checkFunc: func(t *testing.T, result any) {
				if result == nil {
					t.Error("Expected non-nil result")
				}
			},
		},
		{
			name:       "ChangeInfo",
			permission: permissions.ChangeInfo,
			checkFunc: func(t *testing.T, result any) {
				if result == nil {
					t.Error("Expected non-nil result")
				}
			},
		},
		{
			name:       "InviteUsers",
			permission: permissions.InviteUsers,
			checkFunc: func(t *testing.T, result any) {
				if result == nil {
					t.Error("Expected non-nil result")
				}
			},
		},
		{
			name:       "PinMessages",
			permission: permissions.PinMessages,
			checkFunc: func(t *testing.T, result any) {
				if result == nil {
					t.Error("Expected non-nil result")
				}
			},
		},
		{
			name:       "ManageTopics",
			permission: permissions.ManageTopics,
			checkFunc: func(t *testing.T, result any) {
				if result == nil {
					t.Error("Expected non-nil result")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := permissions.Permissions(tt.permission)
			if result == nil {
				t.Error("Expected non-nil ChatPermissions")
			}

			// Convert to any for the test function
			tt.checkFunc(t, any(result))
		})
	}
}

func TestPermissions_Multiple(t *testing.T) {
	result := permissions.Permissions(
		permissions.SendMessages,
		permissions.SendPhotos,
		permissions.InviteUsers,
	)

	if result == nil {
		t.Error("Expected non-nil ChatPermissions")
	}

	// Test that we can call with multiple permissions without error
	// The actual field checking would require casting to gotgbot.ChatPermissions
	// but we're testing the function execution and basic return value
}

func TestPermissions_Empty(t *testing.T) {
	result := permissions.Permissions()

	if result == nil {
		t.Error("Expected non-nil ChatPermissions even with no permissions")
	}
}

func TestPermissions_AllPermissions(t *testing.T) {
	allPerms := []permissions.Permission{
		permissions.SendMessages,
		permissions.SendAudios,
		permissions.SendDocuments,
		permissions.SendPhotos,
		permissions.SendVideos,
		permissions.SendVideoNotes,
		permissions.SendVoiceNotes,
		permissions.SendPolls,
		permissions.SendOtherMessages,
		permissions.AddWebPagePreviews,
		permissions.ReactToMessages,
		permissions.EditTag,
		permissions.ChangeInfo,
		permissions.InviteUsers,
		permissions.PinMessages,
		permissions.ManageTopics,
	}

	result := permissions.Permissions(allPerms...)

	if result == nil {
		t.Error("Expected non-nil ChatPermissions")
	}

	// Test that we can pass all permissions without error
}

func TestPermission_Constants(t *testing.T) {
	// Test that constants have expected values
	if permissions.SendMessages != 0 {
		t.Errorf("Expected SendMessages to be 0, got %d", int(permissions.SendMessages))
	}
	if permissions.SendAudios != 1 {
		t.Errorf("Expected SendAudios to be 1, got %d", int(permissions.SendAudios))
	}
	if permissions.AddWebPagePreviews != 9 {
		t.Errorf("Expected AddWebPagePreviews to be 9, got %d", int(permissions.AddWebPagePreviews))
	}
	if permissions.ReactToMessages != 10 {
		t.Errorf("Expected ReactToMessages to be 10, got %d", int(permissions.ReactToMessages))
	}
	if permissions.EditTag != 11 {
		t.Errorf("Expected EditTag to be 11, got %d", int(permissions.EditTag))
	}
	if permissions.ManageTopics != 15 {
		t.Errorf("Expected ManageTopics to be 15, got %d", int(permissions.ManageTopics))
	}
}

func TestPermissions_NewBotAPI95Fields(t *testing.T) {
	// ReactToMessages and EditTag (Bot API 9.5) are *bool in upstream — verify the
	// generated ChatPermissions encodes them as true when the corresponding flag is set.
	cp := permissions.Permissions(permissions.ReactToMessages, permissions.EditTag)
	if cp == nil {
		t.Fatal("Expected non-nil ChatPermissions")
	}

	if cp.CanReactToMessages == nil || !*cp.CanReactToMessages {
		t.Error("Expected CanReactToMessages to be set to true")
	}

	if cp.CanEditTag == nil || !*cp.CanEditTag {
		t.Error("Expected CanEditTag to be set to true")
	}

	// ManageTopics is also now *bool — sanity check this still works after the upstream type change.
	cp2 := permissions.Permissions(permissions.ManageTopics)
	if cp2.CanManageTopics == nil || !*cp2.CanManageTopics {
		t.Error("Expected CanManageTopics to be set to true")
	}
}
