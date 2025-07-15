package main

import (
	"github.com/enetx/fsm"
	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

// FSM state identifiers define the steps of the user data collection workflow.
const (
	StateGetEmail = "get_email" // Step 1: ask for the user's email address
	StateGetName  = "get_name"  // Step 2: ask for the user's name
	StateSummary  = "summary"   // Final step: show a summary of the collected data
)

// fsmStore holds the active FSM instance for each user, keyed by their Telegram user ID.
// This allows each user to have their own independent state in the conversation.
var fsmStore = NewMapSafe[int64, *fsm.FSM]()

func main() {
	// Load the Telegram bot token from a local .env file.
	token := NewFile("../../../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	// Initialize the Telegram b and its helper components.
	b := bot.New(token).Build().Unwrap()

	// Define a master FSM template. Each new user will receive a clone of this template.
	// This ensures a consistent workflow while maintaining separate states and data for each user.
	template := fsm.NewFSM(StateGetEmail).
		// Defines the linear flow of the conversation from email to name.
		Transition(StateGetEmail, "next", StateGetName).
		// Defines the final transition from name to the summary.
		Transition(StateGetName, "next", StateSummary)

	// Callback executed upon entering StateGetEmail. This is the first step of the workflow.
	template.OnEnter(StateGetEmail, func(fctx *fsm.Context) error {
		// Retrieve the ctx.Context that was stored in the FSM's Values.
		// This happens inside the /start handler or the main message handler.
		tgctx := fctx.Values.Get("tgctx").Some().(*ctx.Context)

		// Send a message to the user asking for their email.
		return tgctx.Reply("Enter your email:").Send().Err()
	})

	// Callback executed upon entering StateGetName, after the user has provided their email.
	template.OnEnter(StateGetName, func(fctx *fsm.Context) error {
		// The user's email is retrieved from the input of the previous step.
		email := fctx.Input.(string)
		// The email is stored in the FSM's persistent Data map for later use in the summary.
		fctx.Data.Set("email", email)

		// Retrieve the latest ctx.Context.
		tgctx := fctx.Values.Get("tgctx").Some().(*ctx.Context)

		// Send a message to the user asking for their name.
		return tgctx.Reply("Enter your name:").Send().Err()
	})

	// Callback executed upon entering StateSummary. This is the final step.
	template.OnEnter(StateSummary, func(fctx *fsm.Context) error {
		// The user's name is retrieved from the input of the current step.
		name := fctx.Input.(string)
		// The email was stored in a previous step and is retrieved from the FSM's Data map.
		email := fctx.Data.Get("email").UnwrapOr("<no email>")

		// Retrieve the latest ctx.Context.
		tgctx := fctx.Values.Get("tgctx").Some().(*ctx.Context)

		// Crucial: Use defer to ensure the FSM instance is removed from the store after this
		// function completes, freeing memory and allowing the user to /start again.
		defer fsmStore.Delete(tgctx.EffectiveUser.Id)

		// Compose and send the final summary message to the user.
		return tgctx.Reply(Format("Got name: {} and email: {}", name, email)).Send().Err()
	})

	// Command handler for /start, which initializes or resets a user's workflow.
	b.Command("start", func(ctx *ctx.Context) error {
		// Get or create an FSM instance for the user.
		entry := fsmStore.Entry(ctx.EffectiveUser.Id)
		// If the user is new, clone the master template for them.
		entry.OrSetBy(template.Clone)
		fsm := entry.Get().Some()

		// Manually set the FSM's state to the beginning. This ensures that
		// a user can restart the flow cleanly even if they were halfway through.
		fsm.SetState(StateGetEmail)

		// Store the current Telegram context in the FSM's temporary Values.
		// This makes it accessible within the first OnEnter callback.
		fsm.Context().Values.Set("tgctx", ctx)

		// Manually trigger the entry callback for the initial state to begin the flow.
		return fsm.CallEnter(StateGetEmail)
	})

	// This is the main handler for all subsequent text messages from the user.
	b.On.Message.Text(func(ctx *ctx.Context) error {
		// Attempt to retrieve the user's FSM instance.
		opt := fsmStore.Get(ctx.EffectiveUser.Id)
		if opt.IsNone() {
			// If no FSM exists, the user has not started the workflow.
			return ctx.Reply("Please type /start to begin.").Send().Err()
		}

		fsm := opt.Some()

		// The user's text message is set as the 'Input' for the current FSM context.
		// This makes it available to the next OnEnter callback.
		fsm.Context().Input = ctx.EffectiveMessage.Text
		// The latest ctx.Context is also updated in the FSM's Values.
		fsm.Context().Values.Set("tgctx", ctx)

		// Trigger a transition using the "next" event, which will move the FSM
		// to the subsequent state as defined in the template.
		return fsm.Trigger("next")
	})

	// Start the bot's polling loop to listen for updates from Telegram.
	b.Polling().Start()
}
