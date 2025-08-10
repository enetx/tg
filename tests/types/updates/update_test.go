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
		{"Poll", updates.Poll, "poll"},
		{"PollAnswer", updates.PollAnswer, "poll_answer"},
		{"MyChatMember", updates.MyChatMember, "my_chat_member"},
		{"ChatMember", updates.ChatMember, "chat_member"},
		{"ChatJoinRequest", updates.ChatJoinRequest, "chat_join_request"},
		{"MessageReaction", updates.MessageReaction, "message_reaction"},
		{"MessageReactionCount", updates.MessageReactionCount, "message_reaction_count"},
		{"BusinessConnection", updates.BusinessConnection, "business_connection"},
		{"BusinessMessage", updates.BusinessMessage, "business_message"},
		{"EditedBusinessMessage", updates.EditedBusinessMessage, "edited_business_message"},
		{"DeletedBusinessMessages", updates.DeletedBusinessMessages, "deleted_business_messages"},
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
