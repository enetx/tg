package handlers

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/enetx/g"
	"github.com/enetx/tg/core"
	"github.com/enetx/tg/handlers/filters"
	"github.com/enetx/tg/types/chat"
)

type BusinessMessagesDeletedHandler struct {
	Filter   filters.BusinessMessagesDeleted
	Response handlers.Response
}

// CheckUpdate checks if the update contains deleted business messages that match the filter.
func (bmdh BusinessMessagesDeletedHandler) CheckUpdate(_ *gotgbot.Bot, ctx *ext.Context) bool {
	return ctx.Update.DeletedBusinessMessages != nil &&
		(bmdh.Filter == nil || bmdh.Filter(ctx.Update.DeletedBusinessMessages))
}

// HandleUpdate processes the deleted business messages update by calling the handler response.
func (bmdh BusinessMessagesDeletedHandler) HandleUpdate(b *gotgbot.Bot, ctx *ext.Context) error {
	return bmdh.Response(b, ctx)
}

// Name returns a unique identifier for this handler instance.
func (bmdh BusinessMessagesDeletedHandler) Name() string {
	return fmt.Sprintf("businessmessagesdeleted_%p", bmdh.Response)
}

// newBusinessMessagesDeleted creates a new business messages deleted handler with the given filter and response.
func newBusinessMessagesDeleted(f filters.BusinessMessagesDeleted, r handlers.Response) BusinessMessagesDeletedHandler {
	return BusinessMessagesDeletedHandler{
		Filter:   f,
		Response: r,
	}
}

// BusinessMessagesDeleted provides methods to handle deleted business messages.
type BusinessMessagesDeleted struct{ Bot core.BotAPI }

// handleBusinessMessagesDelete registers a deleted business messages handler with the dispatcher.
func (h *BusinessMessagesDeleted) handleBusinessMessagesDelete(f filters.BusinessMessagesDeleted, fn Handler) {
	h.Bot.Dispatcher().AddHandler(newBusinessMessagesDeleted(f, wrap(h.Bot, middlewares(h.Bot), fn)))
}

// Any handles all deleted business messages.
func (h *BusinessMessagesDeleted) Any(fn Handler) *BusinessMessagesDeleted {
	h.handleBusinessMessagesDelete(nil, fn)
	return h
}

// ConnectionID handles deleted messages from a specific connection.
func (h *BusinessMessagesDeleted) ConnectionID(connectionID g.String, fn Handler) *BusinessMessagesDeleted {
	h.handleBusinessMessagesDelete(func(d *gotgbot.BusinessMessagesDeleted) bool {
		return d.BusinessConnectionId == connectionID.Std()
	}, fn)
	return h
}

// ChatID handles deleted messages from a specific chat.
func (h *BusinessMessagesDeleted) ChatID(chatID int64, fn Handler) *BusinessMessagesDeleted {
	h.handleBusinessMessagesDelete(func(d *gotgbot.BusinessMessagesDeleted) bool {
		return d.Chat.Id == chatID
	}, fn)
	return h
}

// Private handles deleted messages from private chats only.
func (h *BusinessMessagesDeleted) Private(fn Handler) *BusinessMessagesDeleted {
	h.handleBusinessMessagesDelete(func(d *gotgbot.BusinessMessagesDeleted) bool {
		return d.Chat.Type == chat.Private.String()
	}, fn)
	return h
}

// Group handles deleted messages from groups/supergroups only.
func (h *BusinessMessagesDeleted) Group(fn Handler) *BusinessMessagesDeleted {
	h.handleBusinessMessagesDelete(func(d *gotgbot.BusinessMessagesDeleted) bool {
		return d.Chat.Type == chat.Group.String() || d.Chat.Type == chat.Supergroup.String()
	}, fn)
	return h
}
