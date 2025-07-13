package main

import (
	"github.com/enetx/fsm"
	. "github.com/enetx/g"
	"github.com/enetx/tg"
	"github.com/enetx/tg/keyboard"
)

// FSM state identifiers
const (
	StateName     = "state_name"     // Step: ask for user's name
	StateLike     = "state_like"     // Step: ask if user likes writing bots
	StateLanguage = "state_language" // Step: ask which language was used
	StateSummary  = "state_summary"  // Final summary screen
)

var fsmStore = NewMapSafe[int64, *fsm.FSM]()

func main() {
	// Load token from .env file
	token := NewFile("../../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	// Initialize Telegram bot
	bot := tg.NewBot(token).Build().Unwrap()

	// Define the FSM template (used per user)
	fsmachine := fsm.NewFSM(StateName).
		Transition(StateName, "next", StateLike).
		TransitionWhen(StateLike, "next", StateLanguage, func(ctx *fsm.Context) bool {
			// Only continue to StateLanguage if the user answered "Yes"
			return ctx.Input.(string) == "Yes"
		}).
		TransitionWhen(StateLike, "next", StateSummary, func(ctx *fsm.Context) bool {
			// If the user answered "No", go directly to summary
			return ctx.Input.(string) == "No"
		}).
		Transition(StateLanguage, "next", StateSummary)

	fsmachine.OnEnter(StateName, func(ctx *fsm.Context) error {
		// Retrieve the Telegram context previously stored in the FSM context.
		// This was done inside the /start command handler:
		//   fsm.Context().Values.Set("tgctx", ctx)
		tgctx := ctx.Values.Get("tgctx").Some().(*tg.Context)

		return tgctx.Reply("Hi there! What's your name?").ForceReply().Send().Err()
	})

	fsmachine.OnEnter(StateLike, func(ctx *fsm.Context) error {
		// Extract the user's name from ctx.Input.
		// This was set earlier inside bot.On.Message.Text, when the user responded to the previous question:
		//   fsm.Context().Input = ctx.EffectiveMessage.Text
		name := ctx.Input

		// Store the name in FSM context data for later use (e.g. in summary)
		ctx.Data.Set("name", name)

		// Get the Telegram context from ctx.Values.
		// This was previously stored in the bot.On.Message.Text handler:
		//   fsm.Context().Values.Set("tgctx", ctx)
		tgctx := ctx.Values.Get("tgctx").Some().(*tg.Context)

		return tgctx.Reply(Format("Did you <b>{}</b> like writing bots?", name)).
			HTML().
			Markup(keyboard.Reply().Row().Text("Yes").Text("No")).
			Send().Err()
	})

	fsmachine.OnEnter(StateLanguage, func(ctx *fsm.Context) error {
		// Store the like in FSM context data for later use (e.g. in summary)
		ctx.Data.Set("like", ctx.Input)

		// Get the Telegram context from ctx.Values.
		// This was previously stored in the bot.On.Message.Text handler:
		//   fsm.Context().Values.Set("tgctx", ctx)
		tgctx := ctx.Values.Get("tgctx").Some().(*tg.Context)

		return tgctx.Reply("Cool! I'm too!\nWhat programming language did you use for it?").ForceReply().Send().Err()
	})

	fsmachine.OnEnter(StateSummary, func(ctx *fsm.Context) error {
		// Get the Telegram context from ctx.Values.
		// This was previously stored in the bot.On.Message.Text handler:
		//   fsm.Context().Values.Set("tgctx", ctx)
		tgctx := ctx.Values.Get("tgctx").Some().(*tg.Context)

		// Clean up the FSM instance after summary is shown
		// This prevents stale data and allows /start to begin a fresh flow next time
		defer fsmStore.Delete(tgctx.EffectiveUser.Id)

		// Retrieve previously stored answers
		name := ctx.Data.Get("name").UnwrapOr("<no name>")
		like := ctx.Data.Get("like")

		// If user answered "No" earlier, we never asked for the language,
		// so we skip that part and show a simpler summary
		if like.IsNone() {
			return tgctx.Reply(Format("Not bad, not terrible.\nSee you soon.\n\nSummary:\n- Name: {}\n- Like bots? No", name)).
				RemoveKeyboard().
				Send().
				Err()
		}

		// User answered "Yes" — they were asked which language they used
		lang := ctx.Input.(string)

		// Add a playful greeting if the language is Go
		var greeting String
		if lang == "go" {
			greeting = "Go? Nice choice – that really makes my circuits light up! 😉\n"
		}

		return tgctx.Reply(Format("{}<b>Summary:</b>\n- Name: {}\n- Like bots? {}\n- Language: {}", greeting, name, like.Some(), lang)).
			HTML().
			RemoveKeyboard().
			Send().
			Err()
	})

	bot.Command("start", func(ctx *tg.Context) error {
		// Get or create an FSM instance for the current user (based on user ID)
		// If no FSM exists yet, clone the predefined template FSM
		entry := fsmStore.Entry(ctx.EffectiveUser.Id)
		entry.OrSetBy(fsmachine.Clone) // Clone template FSM if not exists
		fsm := entry.Get().Some()      // Retrieve the FSM (either existing or newly created)

		// Reset the FSM to the initial state for this flow
		// This is useful for re-entry or restarting after cancellation
		fsm.Context().State = StateName

		// Save the current Telegram context (tg.Context) into FSM context.Values,
		// so it can be accessed inside OnEnter callbacks
		// This ensures we have the most recent message/chat/user data
		fsm.Context().Values.Set("tgctx", ctx)

		// Manually invoke OnEnter callbacks for the initial state (StateName)
		// This triggers the first message in the conversation: "Hi there! What's your name?"
		return fsm.CallEnter(StateName)
	})

	bot.Command("cancel", func(ctx *tg.Context) error {
		fsmStore.Delete(ctx.EffectiveUser.Id)
		return ctx.Reply("Cancelled.").RemoveKeyboard().Send().Err()
	})

	bot.On.Message.Text(func(ctx *tg.Context) error {
		// Retrieve the FSM associated with this user (by user ID)
		fsmUser := fsmStore.Get(ctx.EffectiveUser.Id)
		if fsmUser.IsNone() {
			// If no FSM is running for this user — prompt them to start
			return ctx.Reply("Please type /start to begin.").Send().Err()
		}

		fsm := fsmUser.Some()

		// Save the user's message into FSM context (to be accessed in transitions or OnEnter)
		fsm.Context().Input = ctx.EffectiveMessage.Text

		// Update the FSM with the latest Telegram context (it may change between messages)
		fsm.Context().Values.Set("tgctx", ctx)

		// Attempt to perform a transition using the "next" event
		err := fsm.Trigger("next")
		if err != nil {
			// If we're on StateLike (expecting "Yes"/"No") and the input is invalid — show error
			if fsm.Current() == StateLike {
				return ctx.Reply("I don't understand you :(").Send().Err()
			}
		}

		// Otherwise, return the actual error (could be a logic bug)
		return err
	})

	// Start long-polling updates
	bot.Polling().Start()
}
