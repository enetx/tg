package tg

import (
	"github.com/PaulSonOfLars/gotgbot/v2"

	. "github.com/enetx/g"
)

type ReactionHandlers struct{ b *Bot }

func (h *ReactionHandlers) Any(fn Handler) *ReactionHandlers {
	h.b.handleReaction(nil, fn)
	return h
}

func (h *ReactionHandlers) FromPeer(id int64, fn Handler) *ReactionHandlers {
	h.b.handleReaction(func(mru *gotgbot.MessageReactionUpdated) bool {
		if mru == nil {
			return false
		}

		if mru.User != nil {
			return mru.User.Id == id
		}

		if mru.ActorChat != nil {
			return mru.ActorChat.Id == id
		}

		return false
	}, fn)

	return h
}

func (h *ReactionHandlers) ChatID(id int64, fn Handler) *ReactionHandlers {
	h.b.handleReaction(func(mru *gotgbot.MessageReactionUpdated) bool {
		return mru != nil && mru.Chat.Id == id
	}, fn)
	return h
}

func (h *ReactionHandlers) NewReactionEmoji(emoji String, fn Handler) *ReactionHandlers {
	h.b.handleReaction(func(mru *gotgbot.MessageReactionUpdated) bool {
		if mru == nil {
			return false
		}

		for _, r := range mru.NewReaction {
			if r.MergeReactionType().Emoji == emoji.Std() {
				return true
			}
		}

		return false
	}, fn)

	return h
}

func (h *ReactionHandlers) OldReactionEmoji(emoji String, fn Handler) *ReactionHandlers {
	h.b.handleReaction(func(mru *gotgbot.MessageReactionUpdated) bool {
		if mru == nil {
			return false
		}

		for _, r := range mru.OldReaction {
			if r.MergeReactionType().Emoji == emoji.Std() {
				return true
			}
		}

		return false
	}, fn)

	return h
}
