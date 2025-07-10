package tg

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/types/poll"
)

type PollHandlers struct{ bot *Bot }

func (h *PollHandlers) Any(fn Handler) *PollHandlers {
	h.bot.handlePoll(nil, fn)
	return h
}

func (h *PollHandlers) ID(id String, fn Handler) *PollHandlers {
	h.bot.handlePoll(func(p *gotgbot.Poll) bool { return p != nil && p.Id == id.Std() }, fn)
	return h
}

func (h *PollHandlers) Type(t poll.PollType, fn Handler) *PollHandlers {
	h.bot.handlePoll(func(p *gotgbot.Poll) bool { return p != nil && p.Type == t.String() }, fn)
	return h
}

func (h *PollHandlers) Regular(fn Handler) *PollHandlers { return h.Type(poll.Regular, fn) }

func (h *PollHandlers) Quiz(fn Handler) *PollHandlers { return h.Type(poll.Quiz, fn) }
