package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	. "github.com/enetx/g"
	"github.com/enetx/tg/core"
)

// PollAnswerHandlers provides methods to handle poll answer events.
type PollAnswerHandlers struct{ Bot core.BotAPI }

// handlePollAnswer registers a poll answer handler with the dispatcher.
func (h *PollAnswerHandlers) handlePollAnswer(f filters.PollAnswer, fn Handler) {
	h.Bot.Dispatcher().AddHandler(handlers.NewPollAnswer(f, wrap(h.Bot, middlewares(h.Bot), fn)))
}

// Any handles all poll answers.
func (h *PollAnswerHandlers) Any(fn Handler) *PollAnswerHandlers {
	h.handlePollAnswer(nil, fn)
	return h
}

// ID handles poll answers for a specific poll ID.
func (h *PollAnswerHandlers) ID(id String, fn Handler) *PollAnswerHandlers {
	h.handlePollAnswer(func(p *gotgbot.PollAnswer) bool {
		return p != nil && p.PollId == id.Std()
	}, fn)
	return h
}

// FromUserID handles poll answers from a specific user.
func (h *PollAnswerHandlers) FromUserID(id int64, fn Handler) *PollAnswerHandlers {
	h.handlePollAnswer(func(p *gotgbot.PollAnswer) bool {
		return p != nil && p.User.Id == id
	}, fn)
	return h
}
