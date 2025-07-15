package handlers

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"

	"github.com/enetx/tg/core"
	"github.com/enetx/tg/types/chat"
	"github.com/enetx/tg/types/entity"

	. "github.com/enetx/g"
)

type MessageHandler struct {
	bot           core.BotAPI
	filter        filters.Message
	handler       Handler
	name          string
	allowEdited   bool
	allowChannel  bool
	allowBusiness bool
}

func (h *MessageHandler) AllowEdited() *MessageHandler {
	h.allowEdited = true
	return h
}

func (h *MessageHandler) AllowChannel() *MessageHandler {
	h.allowChannel = true
	return h
}

func (h *MessageHandler) AllowBusiness() *MessageHandler {
	h.allowBusiness = true
	return h
}

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

type MessageHandlers struct{ Bot core.BotAPI }

func (h *MessageHandlers) handleMessage(f filters.Message, fn Handler) *MessageHandler {
	return (&MessageHandler{
		bot:     h.Bot,
		filter:  f,
		handler: fn,
		name:    fmt.Sprintf("message_%x_%x", reflect.ValueOf(f).Pointer(), reflect.ValueOf(fn).Pointer()),
	}).Register()
}

func (h *MessageHandlers) Any(fn Handler) *MessageHandler { return h.handleMessage(nil, fn) }

func (h *MessageHandlers) Text(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Text != "" }, fn)
}

func (h *MessageHandlers) Location(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Location != nil }, fn)
}

func (h *MessageHandlers) Contact(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Contact != nil }, fn)
}

func (h *MessageHandlers) Poll(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Poll != nil }, fn)
}

func (h *MessageHandlers) Photo(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return len(msg.Photo) > 0 }, fn)
}

func (h *MessageHandlers) Voice(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Voice != nil }, fn)
}

func (h *MessageHandlers) Video(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Video != nil }, fn)
}

func (h *MessageHandlers) Audio(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Audio != nil }, fn)
}

func (h *MessageHandlers) Sticker(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Sticker != nil }, fn)
}

func (h *MessageHandlers) Document(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Document != nil }, fn)
}

func (h *MessageHandlers) Reply(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.ReplyToMessage != nil }, fn)
}

func (h *MessageHandlers) Animation(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Animation != nil }, fn)
}

func (h *MessageHandlers) VideoNote(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.VideoNote != nil }, fn)
}

func (h *MessageHandlers) Dice(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Dice != nil }, fn)
}

func (h *MessageHandlers) Game(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Game != nil }, fn)
}

func (h *MessageHandlers) Venue(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Venue != nil }, fn)
}

func (h *MessageHandlers) NewChatMembers(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.NewChatMembers != nil }, fn)
}

func (h *MessageHandlers) LeftChatMember(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.LeftChatMember != nil }, fn)
}

func (h *MessageHandlers) PinnedMessage(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.PinnedMessage != nil }, fn)
}

func (h *MessageHandlers) ViaBot(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.ViaBot != nil }, fn)
}

func (h *MessageHandlers) Entities(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return len(msg.Entities) > 0 }, fn)
}

func (h *MessageHandlers) Caption(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Caption != "" }, fn)
}

func (h *MessageHandlers) CaptionEntities(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return len(msg.CaptionEntities) > 0 }, fn)
}

func (h *MessageHandlers) Migrate(fn Handler) *MessageHandler {
	return h.handleMessage(
		func(msg *gotgbot.Message) bool { return msg.MigrateFromChatId != 0 || msg.MigrateToChatId != 0 }, fn)
}

func (h *MessageHandlers) MediaGroup(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.MediaGroupId != "" }, fn)
}

func (h *MessageHandlers) IsAutomaticForward(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.IsAutomaticForward }, fn)
}

func (h *MessageHandlers) UsersShared(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.UsersShared != nil }, fn)
}

func (h *MessageHandlers) ChatShared(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.ChatShared != nil }, fn)
}

func (h *MessageHandlers) Story(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Story != nil }, fn)
}

func (h *MessageHandlers) TopicCreated(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.ForumTopicCreated != nil }, fn)
}

func (h *MessageHandlers) TopicEdited(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.ForumTopicEdited != nil }, fn)
}

func (h *MessageHandlers) TopicClosed(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.ForumTopicClosed != nil }, fn)
}

func (h *MessageHandlers) TopicReopened(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.ForumTopicReopened != nil }, fn)
}

func (h *MessageHandlers) SuccessfulPayment(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.SuccessfulPayment != nil }, fn)
}

func (h *MessageHandlers) RefundedPayment(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.RefundedPayment != nil }, fn)
}

func (h *MessageHandlers) Checklist(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Checklist != nil }, fn)
}

func (h *MessageHandlers) FromUser(id int64, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.From != nil && msg.From.Id == id }, fn)
}

func (h *MessageHandlers) FromUsername(name String, fn Handler) *MessageHandler {
	return h.handleMessage(
		func(msg *gotgbot.Message) bool { return msg.From != nil && msg.From.Username == name.Std() },
		fn,
	)
}

func (h *MessageHandlers) Forwarded(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.ForwardOrigin != nil }, fn)
}

func (h *MessageHandlers) ForwardFromUser(id int64, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool {
		if msg.ForwardOrigin == nil {
			return false
		}

		u := msg.ForwardOrigin.MergeMessageOrigin().SenderUser

		return u != nil && u.Id == id
	}, fn)
}

func (h *MessageHandlers) ForwardFromChat(id int64, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool {
		if msg.ForwardOrigin == nil {
			return false
		}

		c := msg.ForwardOrigin.MergeMessageOrigin().Chat

		return c != nil && c.Id == id
	}, fn)
}

func (h *MessageHandlers) Prefix(prefix String, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return String(msg.GetText()).StartsWith(prefix) }, fn)
}

func (h *MessageHandlers) Suffix(suffix String, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return String(msg.GetText()).EndsWith(suffix) }, fn)
}

func (h *MessageHandlers) Contains(substr String, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return String(msg.GetText()).Contains(substr) }, fn)
}

func (h *MessageHandlers) Equal(str String, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return String(msg.GetText()).Eq(str) }, fn)
}

func (h *MessageHandlers) MatchRegex(pattern *regexp.Regexp, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return pattern.MatchString(msg.GetText()) }, fn)
}

func (h *MessageHandlers) ChatUsername(name String, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Chat.Username == name.Std() }, fn)
}

func (h *MessageHandlers) ChatID(id int64, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Chat.Id == id }, fn)
}

func (h *MessageHandlers) ChatType(chattype chat.ChatType, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Chat.Type == chattype.String() }, fn)
}

func (h *MessageHandlers) Private(fn Handler) *MessageHandler {
	return h.ChatType(chat.Private, fn)
}

func (h *MessageHandlers) Group(fn Handler) *MessageHandler {
	return h.ChatType(chat.Group, fn)
}

func (h *MessageHandlers) Supergroup(fn Handler) *MessageHandler {
	return h.ChatType(chat.Supergroup, fn)
}

func (h *MessageHandlers) Channel(fn Handler) *MessageHandler {
	return h.ChatType(chat.Channel, fn)
}

func (h *MessageHandlers) Business(fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.BusinessConnectionId != "" }, fn)
}

func (h *MessageHandlers) Entity(entType entity.EntityType, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool {
		return Slice[gotgbot.MessageEntity](msg.Entities).Iter().
			Any(func(ent gotgbot.MessageEntity) bool { return ent.Type == entType.String() })
	}, fn)
}

func (h *MessageHandlers) CaptionEntity(entType entity.EntityType, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool {
		return Slice[gotgbot.MessageEntity](msg.CaptionEntities).Iter().
			Any(func(ent gotgbot.MessageEntity) bool { return ent.Type == entType.String() })
	}, fn)
}

func (h *MessageHandlers) DiceValue(val int64, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool { return msg.Dice != nil && msg.Dice.Value == val }, fn)
}

func (h *MessageHandlers) SuccessfulPaymentPrefix(prefix String, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool {
		return msg.SuccessfulPayment != nil && String(msg.SuccessfulPayment.InvoicePayload).StartsWith(prefix)
	}, fn)
}

func (h *MessageHandlers) RefundedPaymentPrefix(prefix String, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool {
		return msg.RefundedPayment != nil && String(msg.RefundedPayment.InvoicePayload).StartsWith(prefix)
	}, fn)
}

func (h *MessageHandlers) ChecklistTitleContains(substr String, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool {
		return msg.Checklist != nil && String(msg.Checklist.Title).Contains(substr)
	}, fn)
}

func (h *MessageHandlers) ChecklistTitleEquals(title String, fn Handler) *MessageHandler {
	return h.handleMessage(func(msg *gotgbot.Message) bool {
		return msg.Checklist != nil && msg.Checklist.Title == title.Std()
	}, fn)
}
