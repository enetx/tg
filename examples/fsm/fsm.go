package main

import (
	"github.com/enetx/fsm"
	. "github.com/enetx/g"
	"github.com/enetx/tg"
	"github.com/enetx/tg/keyboard"
)

// FSM state identifiers define the steps of the user survey.
const (
	StateName     = "state_name"     // Step 1: ask for the user's name
	StateLike     = "state_like"     // Step 2: ask if the user likes writing bots (branching point)
	StateLanguage = "state_language" // Step 3: ask which language was used (only if they like bots)
	StateSummary  = "state_summary"  // Final step: show a summary of the answers
)

// fsmStore holds the active FSM instance for each user, keyed by their Telegram user ID.
// This allows each user to have their own independent state in the conversation.
var fsmStore = NewMapSafe[int64, *fsm.FSM]()

func main() {
	// Load the bot token from a local .env file.
	token := NewFile("../../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	// Initialize the Telegram bot and its helper components.
	bot := tg.NewBot(token).Build().Unwrap()

	// Define a master FSM template. Each new user will receive a clone of this template.
	// This ensures a consistent workflow while maintaining separate states and data.
	fsmachine := fsm.NewFSM(StateName).
		// Defines the first transition from asking the name to asking if they like bots.
		Transition(StateName, "next", StateLike).
		// Defines a conditional transition. This path is only taken if the user's input is "Yes".
		TransitionWhen(StateLike, "next", StateLanguage, func(ctx *fsm.Context) bool {
			return ctx.Input.(string) == "Yes"
		}).
		// Defines another conditional transition. This path is taken if the input is "No",
		// skipping the 'StateLanguage' step entirely.
		TransitionWhen(StateLike, "next", StateSummary, func(ctx *fsm.Context) bool {
			return ctx.Input.(string) == "No"
		}).
		// Defines the final transition from asking the language to showing the summary.
		Transition(StateLanguage, "next", StateSummary)

	// Callback executed upon entering StateName. It prompts the user for their name.
	fsmachine.OnEnter(StateName, func(ctx *fsm.Context) error {
		// Retrieve the tg.Context that was stored in the FSM's Values by the /start handler.
		tgctx := ctx.Values.Get("tgctx").Some().(*tg.Context)

		// Send a message asking for the user's name and force a reply.
		return tgctx.Reply("Hi there! What's your name?").ForceReply().Send().Err()
	})

	// Callback executed upon entering StateLike. It asks the user if they enjoy writing bots.
	fsmachine.OnEnter(StateLike, func(ctx *fsm.Context) error {
		// The user's name is retrieved from the input of the previous step.
		name := ctx.Input

		// The name is stored in the FSM's persistent Data map for later use in the summary.
		ctx.Data.Set("name", name)

		// Retrieve the latest tg.Context.
		tgctx := ctx.Values.Get("tgctx").Some().(*tg.Context)

		// Ask the user a yes/no question using a custom reply keyboard.
		return tgctx.Reply(Format("Did you <b>{}</b> like writing bots?", name)).
			HTML().
			Markup(keyboard.Reply().Row().Text("Yes").Text("No")).
			Send().Err()
	})

	// Callback executed upon entering StateLanguage. This state is only reached if the user answered "Yes".
	fsmachine.OnEnter(StateLanguage, func(ctx *fsm.Context) error {
		// Store the "Yes" answer in the FSM's Data map.
		ctx.Data.Set("like", ctx.Input)

		// Retrieve the latest tg.Context.
		tgctx := ctx.Values.Get("tgctx").Some().(*tg.Context)

		// Ask the user for their programming language and force a reply.
		return tgctx.Reply("Cool! I'm too!\nWhat programming language did you use for it?").ForceReply().Send().Err()
	})

	// Callback executed upon entering StateSummary. This is the final step for all branches.
	fsmachine.OnEnter(StateSummary, func(ctx *fsm.Context) error {
		// Retrieve the latest tg.Context.
		tgctx := ctx.Values.Get("tgctx").Some().(*tg.Context)

		// Crucial: Use defer to ensure the FSM instance is removed from the store after this
		// function completes, freeing memory and allowing the user to /start again.
		defer fsmStore.Delete(tgctx.EffectiveUser.Id)

		// Retrieve all collected data from the FSM's persistent storage.
		name := ctx.Data.Get("name").UnwrapOr("<no name>")
		like := ctx.Data.Get("like")

		// If the user answered "No", the 'like' field was never set.
		// We check for its absence to provide a different summary.
		if like.IsNone() {
			return tgctx.Reply(Format("Not bad, not terrible.\nSee you soon.\n\nSummary:\n- Name: {}\n- Like bots? No", name)).
				RemoveKeyboard().
				Send().
				Err()
		}

		// If we are here, the user answered "Yes" and provided a programming language.
		lang := ctx.Input.(string)

		// Add a playful, conditional greeting based on the language.
		var greeting String
		if lang == "go" {
			greeting = "Go? Nice choice – that really makes my circuits light up! 😉\n"
		}

		// Compose and send the final, detailed summary.
		return tgctx.Reply(Format("{}<b>Summary:</b>\n- Name: {}\n- Like bots? {}\n- Language: {}", greeting, name, like.Some(), lang)).
			HTML().
			RemoveKeyboard().
			Send().
			Err()
	})

	// Command handler for /start, which initializes or resets a user's workflow.
	bot.Command("start", func(ctx *tg.Context) error {
		// Get or create an FSM instance for the user.
		entry := fsmStore.Entry(ctx.EffectiveUser.Id)
		// If the user is new, clone the master template for them.
		entry.OrSetBy(fsmachine.Clone)
		fsm := entry.Get().Some()

		// Manually set the FSM's state to the beginning. This ensures that
		// a user can restart the flow cleanly even if they were halfway through.
		fsm.SetState(StateName)

		// Store the current Telegram context in the FSM's temporary Values.
		// This makes it accessible within the first OnEnter callback.
		fsm.Context().Values.Set("tgctx", ctx)

		// Manually trigger the entry callback for the initial state to begin the flow.
		return fsm.CallEnter(StateName)
	})

	// Command handler for /cancel to prematurely end the workflow and clean up.
	bot.Command("cancel", func(ctx *tg.Context) error {
		fsmStore.Delete(ctx.EffectiveUser.Id)
		return ctx.Reply("Cancelled.").RemoveKeyboard().Send().Err()
	})

	// This is the main handler for all subsequent text messages from the user.
	bot.On.Message.Text(func(ctx *tg.Context) error {
		// Attempt to retrieve the user's FSM instance.
		fsmUser := fsmStore.Get(ctx.EffectiveUser.Id)
		if fsmUser.IsNone() {
			// If no FSM exists, the user has not started the workflow.
			return ctx.Reply("Please type /start to begin.").Send().Err()
		}

		fsm := fsmUser.Some()

		// The user's text message is set as the 'Input' for the current FSM context.
		// This makes it available to GuardFuncs and OnEnter callbacks.
		fsm.Context().Input = ctx.EffectiveMessage.Text
		// The latest tg.Context is also updated in the FSM's Values.
		fsm.Context().Values.Set("tgctx", ctx)

		// Attempt to trigger a transition. The FSM will use its internal rules
		// (including GuardFuncs) to determine the next state.
		err := fsm.Trigger("next")
		if err != nil {
			// If a transition error occurs, we check if it's because the user
			// provided invalid input (e.g., not "Yes" or "No").
			if fsm.Current() == StateLike {
				return ctx.Reply("I don't understand you :(").Send().Err()
			}
		}

		// If the error is not a simple invalid input, it might be a bug. Propagate it.
		return err
	})

	// Start the bot's polling loop to listen for updates from Telegram.
	bot.Polling().Start()
}
