package updates_test

import (
	"testing"

	"github.com/enetx/tg/types/updates"
)

func TestUpdateType_String(t *testing.T) {
	tests := []struct {
		name       string
		updateType updates.UpdateType
		expected   string
	}{
		{"Message", updates.Message, "message"},
		{"EditedMessage", updates.EditedMessage, "edited_message"},
		{"ChannelPost", updates.ChannelPost, "channel_post"},
		{"EditedChannelPost", updates.EditedChannelPost, "edited_channel_post"},
		{"InlineQuery", updates.InlineQuery, "inline_query"},
		{"ChosenInlineResult", updates.ChosenInlineResult, "chosen_inline_result"},
		{"CallbackQuery", updates.CallbackQuery, "callback_query"},
		{"ShippingQuery", updates.ShippingQuery, "shipping_query"},
		{"PreCheckoutQuery", updates.PreCheckoutQuery, "pre_checkout_query"},
		{"PurchasedPaidMedia", updates.PurchasedPaidMedia, "purchased_paid_media"},
		{"Poll", updates.Poll, "poll"},
		{"PollAnswer", updates.PollAnswer, "poll_answer"},
		{"MyChatMember", updates.MyChatMember, "my_chat_member"},
		{"ChatMember", updates.ChatMember, "chat_member"},
		{"ChatJoinRequest", updates.ChatJoinRequest, "chat_join_request"},
		{"ChatBoost", updates.ChatBoost, "chat_boost"},
		{"RemovedChatBoost", updates.RemovedChatBoost, "removed_chat_boost"},
		{"MessageReaction", updates.MessageReaction, "message_reaction"},
		{"MessageReactionCount", updates.MessageReactionCount, "message_reaction_count"},
		{"BusinessConnection", updates.BusinessConnection, "business_connection"},
		{"BusinessMessage", updates.BusinessMessage, "business_message"},
		{"EditedBusinessMessage", updates.EditedBusinessMessage, "edited_business_message"},
		{"DeletedBusinessMessages", updates.DeletedBusinessMessages, "deleted_business_messages"},
		{"GuestMessage", updates.GuestMessage, "guest_message"},
		{"ManagedBot", updates.ManagedBot, "managed_bot"},
		{"Subscription", updates.Subscription, "subscription"},
		{"Unknown", updates.UpdateType(999), "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.updateType.String()
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestAll(t *testing.T) {
	if updates.All.Len() != 26 {
		t.Errorf("Expected All to list 26 update types, got %d", updates.All.Len())
	}

	seen := make(map[string]bool, updates.All.Len())

	updates.All.Iter().ForEach(func(u updates.UpdateType) {
		name := u.String()
		if name == "unknown" {
			t.Error("Expected every update type in All to have a known name")
		}

		if seen[name] {
			t.Errorf("Expected update type %s to appear only once in All", name)
		}

		seen[name] = true
	})
}
