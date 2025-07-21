package main

import (
	"github.com/enetx/fsm"
	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
)

// FSM state identifiers define the steps of the user survey.
const (
	StateName     = "state_name"     // Step 1: ask for the user's name
	StateLike     = "state_like"     // Step 2: ask if the user likes writing bots (branching point)
	StateLanguage = "state_language" // Step 3: ask for programming language (conditional)
	StateSummary  = "state_summary"  // Final step: show a summary of the answers
)

// fsmStore holds the active FSM instance for each user, keyed by their Telegram user ID.
var fsmStore = NewMapSafe[int64, *fsm.SyncFSM]()

func main() {
	// Load the bot token from a local .env file.
	token := NewFile("../../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	// Initialize the Telegram bot and its helper components.
	b := bot.New(token).Build().Unwrap()

	// Define a master FSM template. Each new user will receive a clone of this template.
	fsmachine := fsm.New(StateName).
		// Defines the first transition from asking for a name to the next step.
		Transition(StateName, "next", StateLike).
		// Defines a conditional transition. This path is taken only if the GuardFunc returns true (input is "Yes").
		TransitionWhen(StateLike, "next", StateLanguage, func(ctx *fsm.Context) bool {
			return ctx.Input.(string) == "Yes"
		}).
		// Defines another conditional transition. This path is taken if input is "No", skipping StateLanguage.
		TransitionWhen(StateLike, "next", StateSummary, func(ctx *fsm.Context) bool {
			return ctx.Input.(string) == "No"
		}).
		// Defines the final transition from the language step to the summary.
		Transition(StateLanguage, "next", StateSummary)

	// OnEnter StateName: Prompts the user for their name.
	fsmachine.OnEnter(StateName, func(fctx *fsm.Context) error {
		// Retrieve the tgctx stored in Meta to interact with the Telegram API.
		tgctx := fctx.Meta.Get("tgctx").Some().(*ctx.Context)

		// Ask for the user's name and force a reply.
		return tgctx.Reply("Hi there! What's your name?").ForceReply().Send().Err()
	})

	// OnEnter StateLike: Retrieves the user's name from Input, stores it, and asks the next question.
	fsmachine.OnEnter(StateLike, func(fctx *fsm.Context) error {
		// The user's name is retrieved from the input of the previous trigger.
		name := fctx.Input

		// Store the name in the FSM's persistent Data map for use in the summary.
		fctx.Data.Set("name", name)

		// Retrieve the latest Telegram context to send a reply.
		tgctx := fctx.Meta.Get("tgctx").Some().(*ctx.Context)

		// Ask the user a yes/no question using a reply keyboard.
		return tgctx.Reply(Format("Did you <b>{}</b> like writing bots?", name)).
			HTML().
			Markup(keyboard.Reply().Row().Text("Yes").Text("No")).
			Send().Err()
	})

	// OnEnter StateLanguage: This state is only reached if the user answered "Yes".
	fsmachine.OnEnter(StateLanguage, func(fctx *fsm.Context) error {
		// Store the "Yes" answer, which came from the previous trigger's input.
		fctx.Data.Set("like", fctx.Input)

		tgctx := fctx.Meta.Get("tgctx").Some().(*ctx.Context)

		// Ask for the user's programming language.
		return tgctx.Reply("Cool! I'm too!\nWhat programming language did you use for it?").ForceReply().Send().Err()
	})

	// OnEnter StateSummary: This is the final step, consolidating all collected data.
	fsmachine.OnEnter(StateSummary, func(fctx *fsm.Context) error {
		tgctx := fctx.Meta.Get("tgctx").Some().(*ctx.Context)
		// Use defer to ensure the FSM instance is removed after the interaction is complete.
		defer fsmStore.Delete(tgctx.EffectiveUser.Id)

		// Retrieve all collected data from persistent storage.
		name := fctx.Data.Get("name").UnwrapOr("<no name>")
		like := fctx.Data.Get("like")

		// If the user answered "No", the `like` data was never set.
		// Check for its absence to provide a different summary for the "No" branch.
		if like.IsNone() {
			return tgctx.Reply(Format("Not bad, not terrible.\nSee you soon.\n\nSummary:\n- Name: {}\n- Like bots? No", name)).
				RemoveKeyboard().
				Send().
				Err()
		}

		// If we are here, the user answered "Yes", and the final input is their programming language.
		lang := fctx.Input.(string)

		// Add a playful, conditional greeting.
		var greeting String
		if lang == "go" {
			greeting = "Go? Nice choice – that really makes my circuits light up! 😉\n"
		}

		// Compose and send the final, detailed summary for the "Yes" branch.
		return tgctx.Reply(Format("{}<b>Summary:</b>\n- Name: {}\n- Like bots? {}\n- Language: {}", greeting, name, like.Some(), lang)).
			HTML().
			RemoveKeyboard().
			Send().
			Err()
	})

	// Command handler for /start, which initializes or resets a user's workflow.
	b.Command("start", func(ctx *ctx.Context) error {
		// Get or create an FSM instance for the user.
		entry := fsmStore.Entry(ctx.EffectiveUser.Id)
		entry.OrSetBy(func() *fsm.SyncFSM { return fsmachine.Clone().Sync() })
		fsm := entry.Get().Some()

		// Manually reset the FSM to the initial state. This allows a user
		// to restart the flow cleanly even if they were halfway through.
		fsm.SetState(StateName)

		// Store the current Telegram context in the FSM's Meta store,
		// making it accessible within the first OnEnter callback.
		fsm.Context().Meta.Set("tgctx", ctx)

		// Manually trigger the entry callback for the initial state to begin the flow.
		return fsm.CallEnter(StateName)
	})

	// Command handler for /cancel to prematurely end the workflow and clean up.
	b.Command("cancel", func(ctx *ctx.Context) error {
		fsmStore.Delete(ctx.EffectiveUser.Id)
		return ctx.Reply("Cancelled.").RemoveKeyboard().Send().Err()
	})

	// Main handler for all subsequent text messages from the user.
	b.On.Message.Text(func(ctx *ctx.Context) error {
		// Attempt to retrieve the user's FSM instance.
		fsmUser := fsmStore.Get(ctx.EffectiveUser.Id)
		if fsmUser.IsNone() {
			return ctx.Reply("Please type /start to begin.").Send().Err()
		}

		fsm := fsmUser.Some()
		input := ctx.EffectiveMessage.Text

		// Update the Telegram context to ensure the next callback can reply.
		fsm.Context().Meta.Set("tgctx", ctx)

		// Trigger a transition with the user's text as input. The FSM will use its
		// GuardFuncs to determine the correct next state.
		err := fsm.Trigger("next", input)
		if err != nil {
			// If a transition error occurs (e.g., input is not "Yes" or "No"),
			// check if we are in the branching state and provide a helpful message.
			if fsm.Current() == StateLike {
				return ctx.Reply("I don't understand you :(").Send().Err()
			}
		}

		// Propagate any other, unexpected errors.
		return err
	})

	// Start the bot's polling loop.
	b.Polling().Start()
}
