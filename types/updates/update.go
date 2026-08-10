package updates

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// UpdateType enumerates all supported Telegram update types.
//
// See: https://core.telegram.org/bots/api#update
type UpdateType int

const (
	Message                 UpdateType = iota // Incoming message
	EditedMessage                             // Edited incoming message
	ChannelPost                               // New channel post
	EditedChannelPost                         // Edited channel post
	InlineQuery                               // Inline query (e.g. @your_bot query)
	ChosenInlineResult                        // Result chosen from inline query
	CallbackQuery                             // Button callback with callback_data
	ShippingQuery                             // Shipping address query for payments
	PreCheckoutQuery                          // Pre-checkout query before confirming a payment
	PurchasedPaidMedia                        // Paid media purchased by a user
	Poll                                      // Poll creation
	PollAnswer                                // User’s answer to a non-anonymous poll
	MyChatMember                              // Bot's status updated in a chat
	ChatMember                                // Chat member's status updated
	ChatJoinRequest                           // Incoming join request to a chat
	ChatBoost                                 // Chat boost added or changed
	RemovedChatBoost                          // Chat boost removed
	MessageReaction                           // New or removed reaction to a message
	MessageReactionCount                      // Updated reaction count on a message
	BusinessConnection                        // Bot connected/disconnected from business account
	BusinessMessage                           // New message from connected business account
	EditedBusinessMessage                     // Edited message from connected business account
	DeletedBusinessMessages                   // Messages deleted from connected business account
	GuestMessage                              // Message from a guest bot user
	ManagedBot                                // Managed bot state updated
	Subscription                              // Bot subscription of a user created or changed
)

// String returns the raw string name for UpdateType, as used in Telegram Bot API updates.
func (u UpdateType) String() string {
	switch u {
	case Message:
		return gotgbot.UpdateTypeMessage
	case EditedMessage:
		return gotgbot.UpdateTypeEditedMessage
	case ChannelPost:
		return gotgbot.UpdateTypeChannelPost
	case EditedChannelPost:
		return gotgbot.UpdateTypeEditedChannelPost
	case InlineQuery:
		return gotgbot.UpdateTypeInlineQuery
	case ChosenInlineResult:
		return gotgbot.UpdateTypeChosenInlineResult
	case CallbackQuery:
		return gotgbot.UpdateTypeCallbackQuery
	case ShippingQuery:
		return gotgbot.UpdateTypeShippingQuery
	case PreCheckoutQuery:
		return gotgbot.UpdateTypePreCheckoutQuery
	case PurchasedPaidMedia:
		return gotgbot.UpdateTypePurchasedPaidMedia
	case Poll:
		return gotgbot.UpdateTypePoll
	case PollAnswer:
		return gotgbot.UpdateTypePollAnswer
	case MyChatMember:
		return gotgbot.UpdateTypeMyChatMember
	case ChatMember:
		return gotgbot.UpdateTypeChatMember
	case ChatJoinRequest:
		return gotgbot.UpdateTypeChatJoinRequest
	case ChatBoost:
		return gotgbot.UpdateTypeChatBoost
	case RemovedChatBoost:
		return gotgbot.UpdateTypeRemovedChatBoost
	case MessageReaction:
		return gotgbot.UpdateTypeMessageReaction
	case MessageReactionCount:
		return gotgbot.UpdateTypeMessageReactionCount
	case BusinessConnection:
		return gotgbot.UpdateTypeBusinessConnection
	case BusinessMessage:
		return gotgbot.UpdateTypeBusinessMessage
	case EditedBusinessMessage:
		return gotgbot.UpdateTypeEditedBusinessMessage
	case DeletedBusinessMessages:
		return gotgbot.UpdateTypeDeletedBusinessMessages
	case GuestMessage:
		return gotgbot.UpdateTypeGuestMessage
	case ManagedBot:
		return gotgbot.UpdateTypeManagedBot
	case Subscription:
		return gotgbot.UpdateTypeSubscription
	default:
		return "unknown"
	}
}

var All = g.Slice[UpdateType]{
	Message,
	EditedMessage,
	ChannelPost,
	EditedChannelPost,
	InlineQuery,
	ChosenInlineResult,
	CallbackQuery,
	ShippingQuery,
	PreCheckoutQuery,
	PurchasedPaidMedia,
	Poll,
	PollAnswer,
	MyChatMember,
	ChatMember,
	ChatJoinRequest,
	ChatBoost,
	RemovedChatBoost,
	MessageReaction,
	MessageReactionCount,
	BusinessConnection,
	BusinessMessage,
	EditedBusinessMessage,
	DeletedBusinessMessages,
	GuestMessage,
	ManagedBot,
	Subscription,
}
