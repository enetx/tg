package handlers

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/enetx/g"
	"github.com/enetx/tg/core"
	"github.com/enetx/tg/handlers/filters"
)

type SubscriptionHandler struct {
	Filter   filters.Subscription
	Response handlers.Response
}

// CheckUpdate checks if the update contains a subscription change that matches the filter.
func (sh SubscriptionHandler) CheckUpdate(_ *gotgbot.Bot, ctx *ext.Context) bool {
	return ctx.Update.Subscription != nil &&
		(sh.Filter == nil || sh.Filter(ctx.Update.Subscription))
}

// HandleUpdate processes the subscription update by calling the handler response.
func (sh SubscriptionHandler) HandleUpdate(b *gotgbot.Bot, ctx *ext.Context) error {
	return sh.Response(b, ctx)
}

// Name returns a unique identifier for this handler instance.
func (sh SubscriptionHandler) Name() string {
	return fmt.Sprintf("subscription_%p", sh.Response)
}

// newSubscription creates a new subscription handler with the given filter and response.
func newSubscription(f filters.Subscription, r handlers.Response) SubscriptionHandler {
	return SubscriptionHandler{
		Filter:   f,
		Response: r,
	}
}

// Subscription provides methods to handle user payment subscription changes.
type Subscription struct{ Bot core.BotAPI }

// handleSubscription registers a subscription handler with the dispatcher.
func (h *Subscription) handleSubscription(f filters.Subscription, fn Handler) {
	h.Bot.Dispatcher().AddHandler(newSubscription(f, wrap(h.Bot, middlewares(h.Bot), fn)))
}

// Any handles all subscription changes.
func (h *Subscription) Any(fn Handler) *Subscription {
	h.handleSubscription(nil, fn)
	return h
}

// Active handles subscriptions re-enabled by the user.
func (h *Subscription) Active(fn Handler) *Subscription {
	h.handleSubscription(func(s *gotgbot.BotSubscriptionUpdated) bool { return s.State == "active" }, fn)
	return h
}

// Canceled handles subscriptions canceled by the user.
func (h *Subscription) Canceled(fn Handler) *Subscription {
	h.handleSubscription(func(s *gotgbot.BotSubscriptionUpdated) bool { return s.State == "canceled" }, fn)
	return h
}

// Failed handles subscriptions whose payment failed.
func (h *Subscription) Failed(fn Handler) *Subscription {
	h.handleSubscription(func(s *gotgbot.BotSubscriptionUpdated) bool { return s.State == "failed" }, fn)
	return h
}

// UserID handles subscription changes of a specific user.
func (h *Subscription) UserID(userID int64, fn Handler) *Subscription {
	h.handleSubscription(func(s *gotgbot.BotSubscriptionUpdated) bool { return s.User.Id == userID }, fn)
	return h
}

// Payload handles subscription changes with the specified invoice payload.
func (h *Subscription) Payload(payload g.String, fn Handler) *Subscription {
	h.handleSubscription(
		func(s *gotgbot.BotSubscriptionUpdated) bool { return s.InvoicePayload == payload.Std() }, fn)

	return h
}
