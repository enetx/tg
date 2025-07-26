package handlers

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	. "github.com/enetx/g"
	"github.com/enetx/tg/core"
	"github.com/enetx/tg/types/chat"
	"github.com/enetx/tg/types/entity"
)

// MessageHandler handles message events with specific filters.
type MessageHandler struct {
	bot           core.BotAPI
	filter        filters.Message
	handler       Handler
	name          string
	allowEdited   bool
	allowChannel  bool
	allowBusiness bool
}

// AllowEdited configures the handler to process edited messages.
func (h *MessageHandler) AllowEdited() *MessageHandler {
	h.allowEdited = true
	return h
}

// AllowChannel configures the handler to process messages from channels.
func (h *MessageHandler) AllowChannel() *MessageHandler {
	h.allowChannel = true
	return h
}

// AllowBusiness configures the handler to process business messages.
func (h *MessageHandler) AllowBusiness() *MessageHandler {
	h.allowBusiness = true
	return h
}

// Register registers the message handler with the bot dispatcher.
func (h *MessageHandler) Register() *MessageHandler {
	h.bot.Dispatcher().RemoveHandlerFromGroup(h.name, 0)

	m := handlers.Message{
		AllowEdited:   h.allowEdited,
		AllowChannel:  h.allowChannel,
		AllowBusiness: h.allowBusiness,
		Filter:        h.filter,
		Response:      wrap(h.bot, middlewares(h.bot), h.handler),
	}

	h.bot.Dispatcher().AddHandlerToGroup(namedHandler{h.name, m}, 0)

	return h
}

// MessageHandlers provides methods to handle message events with various filters.
type MessageHandlers struct{ Bot core.BotAPI }

// handleMessage creates and registers a new message handler with the specified filter.
func (h *MessageHandlers) handleMessage(f filters.Message, fn Handler) *MessageHandler {
	return (&MessageHandler{
		bot:     h.Bot,
		filter:  f,
		handler: fn,
		name:    fmt.Sprintf("message_%x_%x", reflect.ValueOf(f).Pointer(), reflect.ValueOf(fn).Pointer()),
	}).Register()
}

// Any handles all messages.
func (h *MessageHandlers) Any(fn Handler) *MessageHandler { return h.handleMessage(nil, fn) }

// Text handles messages that contain text.
func (h *MessageHandlers) Text(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Text != "" }, fn)
}

// Location handles messages that contain location data.
func (h *MessageHandlers) Location(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Location != nil }, fn)
}

// Contact handles messages that contain contact information.
func (h *MessageHandlers) Contact(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Contact != nil }, fn)
}

// Poll handles messages that contain polls.
func (h *MessageHandlers) Poll(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Poll != nil }, fn)
}

// Photo handles messages that contain photos.
func (h *MessageHandlers) Photo(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return len(msg.Photo) > 0 }, fn)
}

// Voice handles messages that contain voice recordings.
func (h *MessageHandlers) Voice(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Voice != nil }, fn)
}

// Video handles messages that contain videos.
func (h *MessageHandlers) Video(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Video != nil }, fn)
}

// Audio handles messages that contain audio files.
func (h *MessageHandlers) Audio(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Audio != nil }, fn)
}

// Sticker handles messages that contain stickers.
func (h *MessageHandlers) Sticker(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Sticker != nil }, fn)
}

// Document handles messages that contain documents.
func (h *MessageHandlers) Document(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Document != nil }, fn)
}

// Reply handles messages that are replies to other messages.
func (h *MessageHandlers) Reply(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.ReplyToMessage != nil }, fn)
}

// Animation handles messages that contain animations (GIFs).
func (h *MessageHandlers) Animation(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Animation != nil }, fn)
}

// VideoNote handles messages that contain video notes (voice messages).
func (h *MessageHandlers) VideoNote(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.VideoNote != nil }, fn)
}

// Dice handles messages that contain dice.
func (h *MessageHandlers) Dice(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Dice != nil }, fn)
}

// Game handles messages that contain games.
func (h *MessageHandlers) Game(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Game != nil }, fn)
}

// Venue handles messages that contain venue information.
func (h *MessageHandlers) Venue(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Venue != nil }, fn)
}

// NewChatMembers handles messages about new chat members joining.
func (h *MessageHandlers) NewChatMembers(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.NewChatMembers != nil }, fn)
}

// LeftChatMember handles messages about chat members leaving.
func (h *MessageHandlers) LeftChatMember(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.LeftChatMember != nil }, fn)
}

// PinnedMessage handles messages about pinned messages.
func (h *MessageHandlers) PinnedMessage(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.PinnedMessage != nil }, fn)
}

// ViaBot handles messages sent via bots.
func (h *MessageHandlers) ViaBot(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.ViaBot != nil }, fn)
}

// Entities handles messages that contain text entities.
func (h *MessageHandlers) Entities(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return len(msg.Entities) > 0 }, fn)
}

// Caption handles messages that contain captions.
func (h *MessageHandlers) Caption(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Caption != "" }, fn)
}

// CaptionEntities handles messages that contain caption entities.
func (h *MessageHandlers) CaptionEntities(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return len(msg.CaptionEntities) > 0 }, fn)
}

// Migrate handles messages about chat migrations.
func (h *MessageHandlers) Migrate(fn Handler) *MessageHandler {
	return h.handleMessage(
		func(msg *gotgbot.Message) bool { return msg.MigrateFromChatId != 0 || msg.MigrateToChatId != 0 }, fn)
}

// MediaGroup handles messages that are part of a media group.
func (h *MessageHandlers) MediaGroup(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.MediaGroupId != "" }, fn)
}

// IsAutomaticForward handles messages that are automatic forwards.
func (h *MessageHandlers) IsAutomaticForward(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.IsAutomaticForward }, fn)
}

// UsersShared handles messages about shared users.
func (h *MessageHandlers) UsersShared(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.UsersShared != nil }, fn)
}

// ChatShared handles messages about shared chats.
func (h *MessageHandlers) ChatShared(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.ChatShared != nil }, fn)
}

// Story handles messages that contain stories.
func (h *MessageHandlers) Story(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Story != nil }, fn)
}

// TopicCreated handles messages about forum topic creation.
func (h *MessageHandlers) TopicCreated(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.ForumTopicCreated != nil }, fn)
}

// TopicEdited handles messages about forum topic edits.
func (h *MessageHandlers) TopicEdited(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.ForumTopicEdited != nil }, fn)
}

// TopicClosed handles messages about forum topic closure.
func (h *MessageHandlers) TopicClosed(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.ForumTopicClosed != nil }, fn)
}

// TopicReopened handles messages about forum topic reopening.
func (h *MessageHandlers) TopicReopened(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.ForumTopicReopened != nil }, fn)
}

// SuccessfulPayment handles messages about successful payments.
func (h *MessageHandlers) SuccessfulPayment(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.SuccessfulPayment != nil }, fn)
}

// RefundedPayment handles messages about refunded payments.
func (h *MessageHandlers) RefundedPayment(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.RefundedPayment != nil }, fn)
}

// Checklist handles messages that contain checklists.
func (h *MessageHandlers) Checklist(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Checklist != nil }, fn)
}

// FromUser handles messages from a specific user ID.
func (h *MessageHandlers) FromUser(id int64, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.From != nil && msg.From.Id == id }, fn)
}

// FromUsername handles messages from a specific username.
func (h *MessageHandlers) FromUsername(name String, fn Handler) *MessageHandler {
	return h.handleMessage(
		func(msg *gotgbot.Message) bool { return msg.From != nil && msg.From.Username == name.Std() },
		fn,
	)
}

// Forwarded handles messages that are forwarded from other chats.
func (h *MessageHandlers) Forwarded(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.ForwardOrigin != nil }, fn)
}

// ForwardFromUser handles messages forwarded from a specific user.
func (h *MessageHandlers) ForwardFromUser(id int64, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool {
		if msg.ForwardOrigin == nil {
			return false
		}

		u := msg.ForwardOrigin.MergeMessageOrigin().SenderUser

		return u != nil && u.Id == id
	}, fn)
}

// ForwardFromChat handles messages forwarded from a specific chat.
func (h *MessageHandlers) ForwardFromChat(id int64, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool {
		if msg.ForwardOrigin == nil {
			return false
		}

		c := msg.ForwardOrigin.MergeMessageOrigin().Chat

		return c != nil && c.Id == id
	}, fn)
}

// Prefix handles messages where text starts with the specified prefix.
func (h *MessageHandlers) Prefix(prefix String, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return String(msg.GetText()).StartsWith(prefix) }, fn)
}

// Suffix handles messages where text ends with the specified suffix.
func (h *MessageHandlers) Suffix(suffix String, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return String(msg.GetText()).EndsWith(suffix) }, fn)
}

// Contains handles messages where text contains the specified substring.
func (h *MessageHandlers) Contains(substr String, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return String(msg.GetText()).Contains(substr) }, fn)
}

// Equal handles messages where text exactly matches the specified string.
func (h *MessageHandlers) Equal(str String, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return String(msg.GetText()).Eq(str) }, fn)
}

// MatchRegex handles messages where text matches the specified regular expression.
func (h *MessageHandlers) MatchRegex(pattern *regexp.Regexp, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return pattern.MatchString(msg.GetText()) }, fn)
}

// ChatUsername handles messages from a chat with the specified username.
func (h *MessageHandlers) ChatUsername(name String, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Chat.Username == name.Std() }, fn)
}

// ChatID handles messages from a specific chat ID.
func (h *MessageHandlers) ChatID(id int64, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Chat.Id == id }, fn)
}

// ChatType handles messages from chats of the specified type.
func (h *MessageHandlers) ChatType(chattype chat.ChatType, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Chat.Type == chattype.String() }, fn)
}

// Private handles messages from private chats.
func (h *MessageHandlers) Private(fn Handler) *MessageHandler {
	return h.ChatType(chat.Private, fn)
}

// Group handles messages from group chats.
func (h *MessageHandlers) Group(fn Handler) *MessageHandler {
	return h.ChatType(chat.Group, fn)
}

// Supergroup handles messages from supergroup chats.
func (h *MessageHandlers) Supergroup(fn Handler) *MessageHandler {
	return h.ChatType(chat.Supergroup, fn)
}

// Channel handles messages from channels.
func (h *MessageHandlers) Channel(fn Handler) *MessageHandler {
	return h.ChatType(chat.Channel, fn)
}

// Business handles messages from business connections.
func (h *MessageHandlers) Business(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.BusinessConnectionId != "" }, fn)
}

// WebAppData handles messages that contain web app data.
func (h *MessageHandlers) WebAppData(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.WebAppData != nil }, fn)
}

// Entity handles messages that contain the specified entity type.
func (h *MessageHandlers) Entity(entType entity.EntityType, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool {
		return Slice[gotgbot.MessageEntity](msg.Entities).Iter().
			Any(func(ent gotgbot.MessageEntity) bool { return ent.Type == entType.String() })
	}, fn)
}

// CaptionEntity handles messages that contain the specified entity type in captions.
func (h *MessageHandlers) CaptionEntity(entType entity.EntityType, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool {
		return Slice[gotgbot.MessageEntity](msg.CaptionEntities).Iter().
			Any(func(ent gotgbot.MessageEntity) bool { return ent.Type == entType.String() })
	}, fn)
}

// DiceValue handles messages with dice that have the specified value.
func (h *MessageHandlers) DiceValue(val int64, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Dice != nil && msg.Dice.Value == val }, fn)
}

// SuccessfulPaymentPrefix handles successful payment messages with payload starting with the specified prefix.
func (h *MessageHandlers) SuccessfulPaymentPrefix(prefix String, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool {
		return msg.SuccessfulPayment != nil && String(msg.SuccessfulPayment.InvoicePayload).StartsWith(prefix)
	}, fn)
}

// RefundedPaymentPrefix handles refunded payment messages with payload starting with the specified prefix.
func (h *MessageHandlers) RefundedPaymentPrefix(prefix String, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool {
		return msg.RefundedPayment != nil && String(msg.RefundedPayment.InvoicePayload).StartsWith(prefix)
	}, fn)
}

// ChecklistTitleContains handles checklist messages where title contains the specified substring.
func (h *MessageHandlers) ChecklistTitleContains(substr String, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool {
		return msg.Checklist != nil && String(msg.Checklist.Title).Contains(substr)
	}, fn)
}

// ChecklistTitleEquals handles checklist messages where title exactly matches the specified string.
func (h *MessageHandlers) ChecklistTitleEquals(title String, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool {
		return msg.Checklist != nil && msg.Checklist.Title == title.Std()
	}, fn)
}
