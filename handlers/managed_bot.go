package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	"github.com/enetx/tg/core"
)

// ManagedBot provides methods to handle managed-bot creation, token update, or owner update events.
type ManagedBot struct{ Bot core.BotAPI }

// newManagedBot creates a new managed-bot handler with the given filter and response.
func newManagedBot(f filters.ManagedBot, r handlers.Response) handlers.ManagedBot {
	return handlers.ManagedBot{
		Filter:   f,
		Response: r,
	}
}

// handleManagedBot registers a managed-bot handler with the dispatcher.
func (h *ManagedBot) handleManagedBot(f filters.ManagedBot, fn Handler) {
	h.Bot.Dispatcher().AddHandler(newManagedBot(f, wrap(h.Bot, middlewares(h.Bot), fn)))
}

// Any handles all managed-bot updates.
func (h *ManagedBot) Any(fn Handler) *ManagedBot {
	h.handleManagedBot(nil, fn)
	return h
}

// OwnedByUserID handles updates only for managed bots created by the given owner user ID.
func (h *ManagedBot) OwnedByUserID(userID int64, fn Handler) *ManagedBot {
	h.handleManagedBot(func(mbu *gotgbot.ManagedBotUpdated) bool {
		return mbu.User.Id == userID
	}, fn)
	return h
}

// AboutBotID handles updates only for the managed bot identified by the given bot ID.
func (h *ManagedBot) AboutBotID(botID int64, fn Handler) *ManagedBot {
	h.handleManagedBot(func(mbu *gotgbot.ManagedBotUpdated) bool {
		return mbu.Bot.Id == botID
	}, fn)
	return h
}
