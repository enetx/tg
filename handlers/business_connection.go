package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
	. "github.com/enetx/g"
	"github.com/enetx/tg/core"
)

// BusinessConnection provides methods to handle business connection events.
type BusinessConnection struct{ Bot core.BotAPI }

// newBusinessConnection creates a new business connection handler with the given filter and response.
func newBusinessConnection(f filters.BusinessConnection, r handlers.Response) handlers.BusinessConnection {
	return handlers.BusinessConnection{
		Filter:   f,
		Response: r,
	}
}

// handleBusinessConnection registers a business connection handler with the dispatcher.
func (h *BusinessConnection) handleBusinessConnection(f filters.BusinessConnection, fn Handler) {
	h.Bot.Dispatcher().AddHandler(newBusinessConnection(f, wrap(h.Bot, middlewares(h.Bot), fn)))
}

// Any handles all business connection events.
func (h *BusinessConnection) Any(fn Handler) *BusinessConnection {
	h.handleBusinessConnection(nil, fn)
	return h
}

// Enabled handles business connection enabled events.
func (h *BusinessConnection) Enabled(fn Handler) *BusinessConnection {
	h.handleBusinessConnection(func(bc *gotgbot.BusinessConnection) bool {
		return bc.IsEnabled
	}, fn)
	return h
}

// Disabled handles business connection disabled events.
func (h *BusinessConnection) Disabled(fn Handler) *BusinessConnection {
	h.handleBusinessConnection(func(bc *gotgbot.BusinessConnection) bool {
		return !bc.IsEnabled
	}, fn)
	return h
}

// FromUser handles business connections from a specific user.
func (h *BusinessConnection) FromUser(userID int64, fn Handler) *BusinessConnection {
	h.handleBusinessConnection(func(bc *gotgbot.BusinessConnection) bool {
		return bc.User.Id == userID
	}, fn)
	return h
}

// FromUsername handles business connections from a specific username.
func (h *BusinessConnection) FromUsername(username String, fn Handler) *BusinessConnection {
	h.handleBusinessConnection(func(bc *gotgbot.BusinessConnection) bool {
		return bc.User.Username == username.Std()
	}, fn)
	return h
}

// ConnectionID handles business connections with a specific connection ID.
func (h *BusinessConnection) ConnectionID(connectionID String, fn Handler) *BusinessConnection {
	h.handleBusinessConnection(func(bc *gotgbot.BusinessConnection) bool {
		return bc.Id == connectionID.Std()
	}, fn)
	return h
}

// CanReply handles business connections where the bot can reply.
func (h *BusinessConnection) CanReply(fn Handler) *BusinessConnection {
	h.handleBusinessConnection(func(bc *gotgbot.BusinessConnection) bool {
		return bc.Rights != nil && bc.Rights.CanReply
	}, fn)
	return h
}
