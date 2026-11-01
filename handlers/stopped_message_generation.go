package handlers

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/enetx/tg/core"
	"github.com/enetx/tg/handlers/filters"
)

type StoppedMessageGenerationHandler struct {
	Filter   filters.StoppedMessageGeneration
	Response handlers.Response
}

// CheckUpdate checks if the update contains a stopped message generation that matches the filter.
func (smgh StoppedMessageGenerationHandler) CheckUpdate(_ *gotgbot.Bot, ctx *ext.Context) bool {
	return ctx.Update.StoppedMessageGeneration != nil &&
		(smgh.Filter == nil || smgh.Filter(ctx.Update.StoppedMessageGeneration))
}

// HandleUpdate processes the stopped message generation by calling the handler response.
func (smgh StoppedMessageGenerationHandler) HandleUpdate(b *gotgbot.Bot, ctx *ext.Context) error {
	return smgh.Response(b, ctx)
}

// Name returns a unique identifier for this handler instance.
func (smgh StoppedMessageGenerationHandler) Name() string {
	return fmt.Sprintf("stopped_message_generation_%p", smgh.Response)
}

// newStoppedMessageGeneration creates a new handler with the given filter and response.
func newStoppedMessageGeneration(f filters.StoppedMessageGeneration, r handlers.Response) StoppedMessageGenerationHandler {
	return StoppedMessageGenerationHandler{
		Filter:   f,
		Response: r,
	}
}

// StoppedMessageGeneration provides methods to handle users stopping message generation.
type StoppedMessageGeneration struct{ Bot core.BotAPI }

// handleStoppedMessageGeneration registers a stopped message generation handler with the dispatcher.
func (h *StoppedMessageGeneration) handleStoppedMessageGeneration(f filters.StoppedMessageGeneration, fn Handler) {
	h.Bot.Dispatcher().AddHandler(newStoppedMessageGeneration(f, wrap(h.Bot, middlewares(h.Bot), fn)))
}

// Any handles all stopped message generations.
func (h *StoppedMessageGeneration) Any(fn Handler) *StoppedMessageGeneration {
	h.handleStoppedMessageGeneration(nil, fn)
	return h
}

// ChatID handles stopped message generations in the specified chat.
func (h *StoppedMessageGeneration) ChatID(chatID int64, fn Handler) *StoppedMessageGeneration {
	h.handleStoppedMessageGeneration(func(s *gotgbot.MessageGenerationStopped) bool { return s.Chat.Id == chatID }, fn)
	return h
}

// Thread handles stopped message generations in the specified message thread.
func (h *StoppedMessageGeneration) Thread(threadID int64, fn Handler) *StoppedMessageGeneration {
	h.handleStoppedMessageGeneration(func(s *gotgbot.MessageGenerationStopped) bool { return s.MessageThreadId == threadID }, fn)
	return h
}

// DraftID handles stopped message generations for the specified draft.
func (h *StoppedMessageGeneration) DraftID(draftID int64, fn Handler) *StoppedMessageGeneration {
	h.handleStoppedMessageGeneration(func(s *gotgbot.MessageGenerationStopped) bool { return s.DraftId == draftID }, fn)
	return h
}
