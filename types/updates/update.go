package updates

// UpdateType enumerates all supported Telegram update types.
//
// See: https://core.telegram.org/bots/api#update
type UpdateType int

const (
	Message              UpdateType = iota // Incoming message
	EditedMessage                          // Edited incoming message
	ChannelPost                            // New channel post
	EditedChannelPost                      // Edited channel post
	InlineQuery                            // Inline query (e.g. @your_bot query)
	ChosenInlineResult                     // Result chosen from inline query
	CallbackQuery                          // Button callback with callback_data
	ShippingQuery                          // Shipping address query for payments
	PreCheckoutQuery                       // Pre-checkout query before confirming a payment
	Poll                                   // Poll creation
	PollAnswer                             // User’s answer to a non-anonymous poll
	MyChatMember                           // Bot's status updated in a chat
	ChatMember                             // Chat member's status updated
	ChatJoinRequest                        // Incoming join request to a chat
	MessageReaction                        // New or removed reaction to a message
	MessageReactionCount                   // Updated reaction count on a message
)

// String returns the raw string name for UpdateType, as used in Telegram Bot API updates.
func (u UpdateType) String() string {
	switch u {
	case Message:
		return "message"
	case EditedMessage:
		return "edited_message"
	case ChannelPost:
		return "channel_post"
	case EditedChannelPost:
		return "edited_channel_post"
	case InlineQuery:
		return "inline_query"
	case ChosenInlineResult:
		return "chosen_inline_result"
	case CallbackQuery:
		return "callback_query"
	case ShippingQuery:
		return "shipping_query"
	case PreCheckoutQuery:
		return "pre_checkout_query"
	case Poll:
		return "poll"
	case PollAnswer:
		return "poll_answer"
	case MyChatMember:
		return "my_chat_member"
	case ChatMember:
		return "chat_member"
	case ChatJoinRequest:
		return "chat_join_request"
	case MessageReaction:
		return "message_reaction"
	case MessageReactionCount:
		return "message_reaction_count"
	default:
		return "unknown"
	}
}
