package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	. "github.com/enetx/g"
	"github.com/enetx/tg/core"
	"github.com/enetx/tg/types/poll"
)

// PollHandlers provides methods to handle poll events.
type PollHandlers struct{ Bot core.BotAPI }

func (h *PollHandlers) handlePoll(f filters.Poll, fn Handler) {
	h.Bot.Dispatcher().AddHandler(handlers.NewPoll(f, wrap(h.Bot, middlewares(h.Bot), fn)))
}

// Any handles all polls.
func (h *PollHandlers) Any(fn Handler) *PollHandlers {
	h.handlePoll(nil, fn)
	return h
}

// ID handles polls with a specific ID.
func (h *PollHandlers) ID(id String, fn Handler) *PollHandlers {
	h.handlePoll(func(p *gotgbot.Poll) bool { return p != nil && p.Id == id.Std() }, fn)
	return h
}

// Type handles polls of the specified type.
func (h *PollHandlers) Type(t poll.PollType, fn Handler) *PollHandlers {
	h.handlePoll(func(p *gotgbot.Poll) bool { return p != nil && p.Type == t.String() }, fn)
	return h
}

// Regular handles regular polls.
func (h *PollHandlers) Regular(fn Handler) *PollHandlers { return h.Type(poll.Regular, fn) }

// Quiz handles quiz polls.
func (h *PollHandlers) Quiz(fn Handler) *PollHandlers { return h.Type(poll.Quiz, fn) }
