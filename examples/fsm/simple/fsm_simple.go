package main

import (
	"github.com/enetx/fsm"
	"github.com/enetx/g"
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
var fsmStore = g.NewMapSafe[int64, *fsm.SyncFSM]()

func main() {
	// Load the Telegram bot token from a local .env file.
	token := g.NewFile("../../../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	// Initialize the Telegram bot and its helper components.
	b := bot.New(token).Build().Unwrap()

	// Define a master FSM template. Each new user will receive a clone of this template,
	// ensuring a consistent workflow while maintaining separate states and data for each user.
	template := fsm.New(StateGetEmail).
		// Defines the linear flow of the conversation from email to name.
		Transition(StateGetEmail, "next", StateGetName).
		// Defines the final transition from name to the summary.
		Transition(StateGetName, "next", StateSummary)

	// OnEnter StateGetEmail: This is the first step of the workflow.
	template.OnEnter(StateGetEmail, func(fctx *fsm.Context) error {
		// Retrieve the tgctx stored in Meta. This is needed to send a reply.
		tgctx := fctx.Meta.Get("tgctx").Some().(*ctx.Context)

		// Send a message to the user asking for their email.
		return tgctx.Reply("Enter your email:").Send().Err()
	})

	// OnEnter StateGetName: Executed after the user has provided their email.
	template.OnEnter(StateGetName, func(fctx *fsm.Context) error {
		// The user's email is retrieved from the input of the previous `Trigger` call.
		email := fctx.Input.(string)
		// The email is stored in the FSM's persistent Data map for use in the final summary.
		fctx.Data.Set("email", email)

		// Retrieve the latest tgctx to send the next prompt.
		tgctx := fctx.Meta.Get("tgctx").Some().(*ctx.Context)

		// Send a message to the user asking for their name.
		return tgctx.Reply("Enter your name:").Send().Err()
	})

	// OnEnter StateSummary: This is the final step, displaying the collected data.
	template.OnEnter(StateSummary, func(fctx *fsm.Context) error {
		// The user's name is retrieved from the input of the most recent `Trigger` call.
		name := fctx.Input.(string)
		// The email was stored in a previous step and is retrieved from the FSM's Data map.
		email := fctx.Data.Get("email").UnwrapOr("<no email>")

		// Retrieve the latest Telegram context.
		tgctx := fctx.Meta.Get("tgctx").Some().(*ctx.Context)

		// Use defer to ensure the FSM instance is removed from the store after this function completes,
		// freeing memory and allowing the user to /start again.
		defer fsmStore.Delete(tgctx.EffectiveUser.Id)

		// Compose and send the final summary message to the user.
		return tgctx.Reply(g.Format("Got name: {} and email: {}", name, email)).Send().Err()
	})

	// Command handler for /start, which initializes or resets a user's workflow.
	b.Command("start", func(ctx *ctx.Context) error {
		// Get or create an FSM instance for the user.
		entry := fsmStore.Entry(ctx.EffectiveUser.Id)
		entry.OrSetBy(func() *fsm.SyncFSM { return template.Clone().Sync() })
		fsm := entry.Get().Some()

		// Manually reset the FSM to the initial state. This allows a user
		// to restart the flow cleanly even if they were halfway through.
		fsm.SetState(StateGetEmail)

		// Store the current Telegram context in the FSM's Meta store,
		// making it accessible within the first OnEnter callback.
		fsm.Context().Meta.Set("tgctx", ctx)

		// Manually trigger the entry callback for the initial state to begin the flow.
		return fsm.CallEnter(StateGetEmail)
	})

	// Main handler for all subsequent text messages from the user.
	b.On.Message.Text(func(ctx *ctx.Context) error {
		// Attempt to retrieve the user's FSM instance.
		opt := fsmStore.Get(ctx.EffectiveUser.Id)
		if opt.IsNone() {
			// If no FSM exists, the user has not started the workflow.
			return ctx.Reply("Please type /start to begin.").Send().Err()
		}

		fsm := opt.Some()

		// Update the Telegram context to ensure the next callback can reply.
		fsm.Context().Meta.Set("tgctx", ctx)

		// Trigger the "next" event, passing the user's message text as input.
		// The FSM will move to the subsequent state as defined in the template.
		return fsm.Trigger("next", ctx.EffectiveMessage.Text)
	})

	// Start the bot's polling loop to listen for updates from Telegram.
	b.Polling().Start()
}
