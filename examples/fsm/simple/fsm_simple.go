package main

import (
	"github.com/enetx/fsm"
	. "github.com/enetx/g"
	"github.com/enetx/tg"
)

// FSM state identifiers
const (
	StateGetEmail = "get_email" // Step 1: ask for user's email
	StateGetName  = "get_name"  // Step 2: ask for user's name
	StateSummary  = "summary"   // Final step: show summary
)

// Global FSM store: holds per-user FSM instances
var fsmStore = NewMapSafe[int64, *fsm.FSM]()

func main() {
	// Load the Telegram bot token from .env
	token := NewFile("../../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	bot := tg.NewBot(token).Build().Unwrap()

	// Define a template FSM — this will be cloned per user
	template := fsm.NewFSM(StateGetEmail).
		Transition(StateGetEmail, "next", StateGetName).
		Transition(StateGetName, "next", StateSummary)

	// Handler for stateGetEmail — triggered on /start and after reset
	template.OnEnter(StateGetEmail, func(ctx *fsm.Context) error {
		// Retrieve Telegram context passed in from /start or message
		tgctx := ctx.Values.Get("tgctx").Some().(*tg.Context)
		// Ask for email address
		return tgctx.Reply("Enter your email:").Send().Err()
	})

	// Handler for stateGetName — expects previous ctx.Input to contain email
	template.OnEnter(StateGetName, func(ctx *fsm.Context) error {
		// Extract email from Input (comes from previous message)
		email := ctx.Input.(string)
		ctx.Data.Set("email", email) // Save email to FSM context

		// Retrieve current Telegram context
		tgctx := ctx.Values.Get("tgctx").Some().(*tg.Context)
		// Ask for name
		return tgctx.Reply("Enter your name:").Send().Err()
	})

	// Handler for stateSummary — triggered after email + name collected
	template.OnEnter(StateSummary, func(ctx *fsm.Context) error {
		// Name comes from Input in this step
		name := ctx.Input.(string)
		// Email was stored earlier in FSM context
		email := ctx.Data.Get("email").UnwrapOr("<no email>")

		// Retrieve current Telegram context
		tgctx := ctx.Values.Get("tgctx").Some().(*tg.Context)

		// Clean up FSM for this user (conversation finished)
		defer fsmStore.Delete(tgctx.EffectiveUser.Id)

		// Send summary message
		return tgctx.Reply(Format("Got name: {} and email: {}", name, email)).Send().Err()
	})

	// /start command — initializes FSM and triggers first OnEnter
	bot.Command("start", func(ctx *tg.Context) error {
		// Get or create per-user FSM instance from template
		entry := fsmStore.Entry(ctx.EffectiveUser.Id)
		entry.OrSetBy(template.Clone)
		fsm := entry.Get().Some()

		// Reset FSM to initial state and bind current tg.Context
		fsm.Context().State = StateGetEmail
		fsm.Context().Values.Set("tgctx", ctx)

		// Trigger the OnEnter callback for StateGetEmail
		return fsm.CallEnter(StateGetEmail)
	})

	// Message handler — handles input for all steps
	bot.On.Message.Text(func(ctx *tg.Context) error {
		// Look up FSM for this user
		opt := fsmStore.Get(ctx.EffectiveUser.Id)
		if opt.IsNone() {
			// If no session, prompt to start again
			return ctx.Reply("Please type /start to begin.").Send().Err()
		}

		fsm := opt.Some()

		// Save user message into FSM context as Input (used by transitions)
		fsm.Context().Input = ctx.EffectiveMessage.Text
		// Save current tg.Context to be accessible during OnEnter
		fsm.Context().Values.Set("tgctx", ctx)

		// Trigger transition with fixed event name "next"
		return fsm.Trigger("next")
	})

	// Start receiving updates (long polling)
	bot.Polling().Start()
}
