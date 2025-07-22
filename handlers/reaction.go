package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	. "github.com/enetx/g"
	"github.com/enetx/tg/core"
)

// ReactionHandlers provides methods to handle message reaction events.
type ReactionHandlers struct{ Bot core.BotAPI }

// handleReaction registers a message reaction handler with the dispatcher.
func (h *ReactionHandlers) handleReaction(f filters.Reaction, fn Handler) {
	h.Bot.Dispatcher().AddHandler(handlers.NewReaction(f, wrap(h.Bot, middlewares(h.Bot), fn)))
}

// Any handles all message reaction updates.
func (h *ReactionHandlers) Any(fn Handler) *ReactionHandlers {
	h.handleReaction(nil, fn)
	return h
}

// FromPeer handles reactions from a specific user or chat.
func (h *ReactionHandlers) FromPeer(id int64, fn Handler) *ReactionHandlers {
	h.handleReaction(func(mru *gotgbot.MessageReactionUpdated) bool {
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

// ChatID handles reactions in a specific chat.
func (h *ReactionHandlers) ChatID(id int64, fn Handler) *ReactionHandlers {
	h.handleReaction(func(mru *gotgbot.MessageReactionUpdated) bool {
		return mru != nil && mru.Chat.Id == id
	}, fn)
	return h
}

// NewReactionEmoji handles reactions where the specified emoji was added.
func (h *ReactionHandlers) NewReactionEmoji(emoji String, fn Handler) *ReactionHandlers {
	h.handleReaction(func(mru *gotgbot.MessageReactionUpdated) bool {
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

// OldReactionEmoji handles reactions where the specified emoji was removed.
func (h *ReactionHandlers) OldReactionEmoji(emoji String, fn Handler) *ReactionHandlers {
	h.handleReaction(func(mru *gotgbot.MessageReactionUpdated) bool {
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
