package tg

import (
	"regexp"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	. "github.com/enetx/g"
	"github.com/enetx/tg/types/chat"
	"github.com/enetx/tg/types/entity"
)

type MessageHandler struct {
	bot           *Bot
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
	h.bot.dispatcher.RemoveHandlerFromGroup(h.name, 0)

	m := handlers.Message{
		AllowEdited:   h.allowEdited,
		AllowChannel:  h.allowChannel,
		AllowBusiness: h.allowBusiness,
		Filter:        h.filter,
		Response:      wrap(h.bot, h.handler),
	}

	h.bot.dispatcher.AddHandlerToGroup(namedHandler{h.name, m}, 0)

	return h
}

type MessageHandlers struct{ bot *Bot }

func (m *MessageHandlers) Any(fn Handler) *MessageHandler { return m.bot.handleMessage(nil, fn) }

func (m *MessageHandlers) Text(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.Text != "" }, fn)
}

func (m *MessageHandlers) Location(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.Location != nil }, fn)
}

func (m *MessageHandlers) Contact(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.Contact != nil }, fn)
}

func (m *MessageHandlers) Poll(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.Poll != nil }, fn)
}

func (m *MessageHandlers) Photo(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return len(msg.Photo) > 0 }, fn)
}

func (m *MessageHandlers) Voice(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.Voice != nil }, fn)
}

func (m *MessageHandlers) Video(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.Video != nil }, fn)
}

func (m *MessageHandlers) Audio(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.Audio != nil }, fn)
}

func (m *MessageHandlers) Sticker(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.Sticker != nil }, fn)
}

func (m *MessageHandlers) Document(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.Document != nil }, fn)
}

func (m *MessageHandlers) Reply(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.ReplyToMessage != nil }, fn)
}

func (m *MessageHandlers) Animation(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.Animation != nil }, fn)
}

func (m *MessageHandlers) VideoNote(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.VideoNote != nil }, fn)
}

func (m *MessageHandlers) Dice(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.Dice != nil }, fn)
}

func (m *MessageHandlers) Game(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.Game != nil }, fn)
}

func (m *MessageHandlers) Venue(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.Venue != nil }, fn)
}

func (m *MessageHandlers) NewChatMembers(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.NewChatMembers != nil }, fn)
}

func (m *MessageHandlers) LeftChatMember(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.LeftChatMember != nil }, fn)
}

func (m *MessageHandlers) PinnedMessage(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.PinnedMessage != nil }, fn)
}

func (m *MessageHandlers) ViaBot(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.ViaBot != nil }, fn)
}

func (m *MessageHandlers) Entities(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return len(msg.Entities) > 0 }, fn)
}

func (m *MessageHandlers) Caption(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.Caption != "" }, fn)
}

func (m *MessageHandlers) CaptionEntities(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return len(msg.CaptionEntities) > 0 }, fn)
}

func (m *MessageHandlers) Migrate(fn Handler) *MessageHandler {
	return m.bot.handleMessage(
		func(msg *gotgbot.Message) bool { return msg.MigrateFromChatId != 0 || msg.MigrateToChatId != 0 }, fn)
}

func (m *MessageHandlers) MediaGroup(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.MediaGroupId != "" }, fn)
}

func (m *MessageHandlers) IsAutomaticForward(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.IsAutomaticForward }, fn)
}

func (m *MessageHandlers) UsersShared(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.UsersShared != nil }, fn)
}

func (m *MessageHandlers) ChatShared(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.ChatShared != nil }, fn)
}

func (m *MessageHandlers) Story(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.Story != nil }, fn)
}

func (m *MessageHandlers) TopicCreated(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.ForumTopicCreated != nil }, fn)
}

func (m *MessageHandlers) TopicEdited(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.ForumTopicEdited != nil }, fn)
}

func (m *MessageHandlers) TopicClosed(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.ForumTopicClosed != nil }, fn)
}

func (m *MessageHandlers) TopicReopened(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.ForumTopicReopened != nil }, fn)
}

func (m *MessageHandlers) SuccessfulPayment(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.SuccessfulPayment != nil }, fn)
}

func (m *MessageHandlers) RefundedPayment(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.RefundedPayment != nil }, fn)
}

func (m *MessageHandlers) Checklist(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.Checklist != nil }, fn)
}

func (m *MessageHandlers) FromUser(id int64, fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.From != nil && msg.From.Id == id }, fn)
}

func (m *MessageHandlers) FromUsername(name String, fn Handler) *MessageHandler {
	return m.bot.handleMessage(
		func(msg *gotgbot.Message) bool { return msg.From != nil && msg.From.Username == name.Std() },
		fn,
	)
}

func (m *MessageHandlers) Forwarded(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.ForwardOrigin != nil }, fn)
}

func (m *MessageHandlers) ForwardFromUser(id int64, fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool {
		if msg.ForwardOrigin == nil {
			return false
		}

		u := msg.ForwardOrigin.MergeMessageOrigin().SenderUser

		return u != nil && u.Id == id
	}, fn)
}

func (m *MessageHandlers) ForwardFromChat(id int64, fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool {
		if msg.ForwardOrigin == nil {
			return false
		}

		c := msg.ForwardOrigin.MergeMessageOrigin().Chat

		return c != nil && c.Id == id
	}, fn)
}

func (m *MessageHandlers) Prefix(prefix String, fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return String(msg.GetText()).StartsWith(prefix) }, fn)
}

func (m *MessageHandlers) Suffix(suffix String, fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return String(msg.GetText()).EndsWith(suffix) }, fn)
}

func (m *MessageHandlers) Contains(substr String, fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return String(msg.GetText()).Contains(substr) }, fn)
}

func (m *MessageHandlers) Equal(str String, fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return String(msg.GetText()).Eq(str) }, fn)
}

func (m *MessageHandlers) MatchRegex(pattern *regexp.Regexp, fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return pattern.MatchString(msg.GetText()) }, fn)
}

func (m *MessageHandlers) ChatUsername(name String, fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.Chat.Username == name.Std() }, fn)
}

func (m *MessageHandlers) ChatID(id int64, fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.Chat.Id == id }, fn)
}

func (m *MessageHandlers) ChatType(chattype chat.ChatType, fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.Chat.Type == chattype.String() }, fn)
}

func (m *MessageHandlers) Private(fn Handler) *MessageHandler {
	return m.ChatType(chat.Private, fn)
}

func (m *MessageHandlers) Group(fn Handler) *MessageHandler {
	return m.ChatType(chat.Group, fn)
}

func (m *MessageHandlers) Supergroup(fn Handler) *MessageHandler {
	return m.ChatType(chat.Supergroup, fn)
}

func (m *MessageHandlers) Channel(fn Handler) *MessageHandler {
	return m.ChatType(chat.Channel, fn)
}

func (m *MessageHandlers) Business(fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.BusinessConnectionId != "" }, fn)
}

func (m *MessageHandlers) Entity(entType entity.EntityType, fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool {
		return Slice[gotgbot.MessageEntity](msg.Entities).Iter().
			Any(func(ent gotgbot.MessageEntity) bool { return ent.Type == entType.String() })
	}, fn)
}

func (m *MessageHandlers) CaptionEntity(entType entity.EntityType, fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool {
		return Slice[gotgbot.MessageEntity](msg.CaptionEntities).Iter().
			Any(func(ent gotgbot.MessageEntity) bool { return ent.Type == entType.String() })
	}, fn)
}

func (m *MessageHandlers) DiceValue(val int64, fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool { return msg.Dice != nil && msg.Dice.Value == val }, fn)
}

func (m *MessageHandlers) SuccessfulPaymentPrefix(prefix String, fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool {
		return msg.SuccessfulPayment != nil && String(msg.SuccessfulPayment.InvoicePayload).StartsWith(prefix)
	}, fn)
}

func (m *MessageHandlers) RefundedPaymentPrefix(prefix String, fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool {
		return msg.RefundedPayment != nil && String(msg.RefundedPayment.InvoicePayload).StartsWith(prefix)
	}, fn)
}

func (m *MessageHandlers) ChecklistTitleContains(substr String, fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool {
		return msg.Checklist != nil && String(msg.Checklist.Title).Contains(substr)
	}, fn)
}

func (m *MessageHandlers) ChecklistTitleEquals(title String, fn Handler) *MessageHandler {
	return m.bot.handleMessage(func(msg *gotgbot.Message) bool {
		return msg.Checklist != nil && msg.Checklist.Title == title.Std()
	}, fn)
}
