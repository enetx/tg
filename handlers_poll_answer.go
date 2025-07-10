package tg

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

type PollAnswerHandlers struct{ b *Bot }

func (h *PollAnswerHandlers) Any(fn Handler) *PollAnswerHandlers {
	h.b.handlePollAnswer(nil, fn)
	return h
}

func (h *PollAnswerHandlers) ID(id String, fn Handler) *PollAnswerHandlers {
	h.b.handlePollAnswer(func(p *gotgbot.PollAnswer) bool {
		return p != nil && p.PollId == id.Std()
	}, fn)
	return h
}

func (h *PollAnswerHandlers) FromUserID(id int64, fn Handler) *PollAnswerHandlers {
	h.b.handlePollAnswer(func(p *gotgbot.PollAnswer) bool {
		return p != nil && p.User.Id == id
	}, fn)
	return h
}
