package main

import (
	"github.com/enetx/fsm"
	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
)

// FSM state identifiers define the steps of the user preference survey.
const (
	StateColor   = "color"   // Step 1: user chooses their favorite color
	StateAnimal  = "animal"  // Step 2: user chooses their favorite animal
	StateSummary = "summary" // Final step: show a summary of their choices
)

// fsmStore holds the active FSM instance for each user, keyed by their Telegram user ID.
// This allows each user to have their own independent state in the conversation.
var fsmStore = NewMapSafe[int64, *fsm.FSM]()

func main() {
	// Load the Telegram bot token from a local .env file.
	token := NewFile("../../../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()

	// Initialize the Telegram bot and its helper components.
	b := bot.New(token).Build().Unwrap()

	// Define a master FSM template. Each new user will receive a clone of this template,
	// ensuring a consistent workflow while maintaining separate states and data for each user.
	template := fsm.NewFSM(StateColor).
		// Defines the transition from choosing a color to choosing an animal.
		Transition(StateColor, "color_selected", StateAnimal).
		// Defines the final transition from choosing an animal to showing the summary.
		Transition(StateAnimal, "animal_selected", StateSummary)

	// Step 1: Callback executed upon entering StateColor. Asks the user to select a color.
	template.OnEnter(StateColor, func(fctx *fsm.Context) error {
		// Retrieve the tgctx stored in the FSM's Meta to interact with the Telegram API.
		tgctx := fctx.Meta.Get("tgctx").Some().(*ctx.Context)

		// Send a message with an inline keyboard. Each button's callback_data is prefixed
		// with "color:" to be handled by the corresponding callback handler.
		return tgctx.Reply("🎨 Choose your favorite color:").
			Markup(keyboard.Inline().
				Row().Text("❤️ Red", "color:red").
				Row().Text("💙 Blue", "color:blue").
				Row().Text("💚 Green", "color:green")).
			Send().Err()
	})

	// Step 2: Callback executed upon entering StateAnimal, after a color has been selected.
	template.OnEnter(StateAnimal, func(fctx *fsm.Context) error {
		// Retrieve the latest Telegram context.
		tgctx := fctx.Meta.Get("tgctx").Some().(*ctx.Context)

		// The selected color was passed as 'Input' from the color callback handler.
		color := fctx.Input.(String)
		// Store the color in the FSM's persistent Data map for later use.
		fctx.Data.Set("color", color)

		// Acknowledge the user's selection by sending a new message.
		tgctx.Reply("✅ Color selected: " + color).Send()

		// Send a new message asking for the user's favorite animal.
		return tgctx.Message("🐾 Pick your favorite animal:").
			Markup(keyboard.Inline().
				Row().Text("🐶 Dog", "animal:dog").
				Row().Text("🐱 Cat", "animal:cat").
				Row().Text("🦊 Fox", "animal:fox")).
			Send().Err()
	})

	// Step 3: Callback executed upon entering StateSummary, the final step.
	template.OnEnter(StateSummary, func(fctx *fsm.Context) error {
		// Retrieve the latest Telegram context.
		tgctx := fctx.Meta.Get("tgctx").Some().(*ctx.Context)
		// Use defer to ensure the FSM instance is removed from the store after this function completes,
		// freeing memory and allowing the user to /start again.
		defer fsmStore.Delete(tgctx.EffectiveUser.Id)

		// Retrieve all collected data: color from the Data map, and animal from the current Input.
		color := fctx.Data.Get("color").UnwrapOr("unknown")
		animal := fctx.Input.(String)

		// Acknowledge the final selection.
		tgctx.Reply("✅ Animal selected: " + animal).Send()

		// Compose and send the final summary message to the user.
		return tgctx.Message(Format("🧾 Your preferences:\n- Color: {}\n- Animal: {}", color, animal)).Send().Err()
	})

	// Command handler for /start, which initializes or resets a user's workflow.
	b.Command("start", func(ctx *ctx.Context) error {
		// Get or create an FSM instance for the user.
		entry := fsmStore.Entry(ctx.EffectiveUser.Id)
		entry.OrSetBy(template.Clone)
		fsm := entry.Get().Some()

		// Store the current Telegram context in the FSM's temporary Meta store.
		fsm.Context().Meta.Set("tgctx", ctx)

		// Manually trigger the entry callback for the initial state to begin the flow.
		return fsm.CallEnter(StateColor)
	})

	// Command handler for /cancel to prematurely end the workflow.
	b.Command("cancel", func(ctx *ctx.Context) error {
		fsmStore.Delete(ctx.EffectiveUser.Id)
		return ctx.Reply("Cancelled.").RemoveKeyboard().Send().Err()
	})

	// Callback handler for buttons with data prefixed by "color:".
	b.On.Callback.Prefix("color:", func(ctx *ctx.Context) error {
		// Retrieve the user's FSM instance.
		fsmOpt := fsmStore.Get(ctx.EffectiveUser.Id)
		if fsmOpt.IsNone() {
			return ctx.Reply("Please type /start to begin.").Send().Err()
		}

		fsm := fsmOpt.Some()
		fsm.Context().Meta.Set("tgctx", ctx)

		// Extract the color value from the callback data (e.g., "color:red" -> "red")
		// and use it as the input for the FSM transition.
		input := String(ctx.Callback.Data).StripPrefix("color:")

		// Trigger the transition associated with selecting a color.
		return fsm.Trigger("color_selected", input)
	})

	// Callback handler for buttons with data prefixed by "animal:".
	b.On.Callback.Prefix("animal:", func(ctx *ctx.Context) error {
		// Retrieve the user's FSM instance.
		fsmOpt := fsmStore.Get(ctx.EffectiveUser.Id)
		if fsmOpt.IsNone() {
			return ctx.Reply("Please type /start to begin.").Send().Err()
		}

		fsm := fsmOpt.Some()
		fsm.Context().Meta.Set("tgctx", ctx)

		// Extract the animal value from the callback data and use it as the FSM's input.
		input := String(ctx.Callback.Data).StripPrefix("animal:")

		// Trigger the transition associated with selecting an animal.
		return fsm.Trigger("animal_selected", input)
	})

	// A catch-all handler to clean up any unsupported user interactions.
	// This prevents the bot from appearing unresponsive if the user types text
	// when a button press is expected.
	b.On.Any(func(ctx *ctx.Context) error {
		// If the update is a text message (not a callback), delete it.
		if ctx.Callback == nil {
			return ctx.Delete().Send().Err()
		}
		// Ignore any other types of updates.
		return nil
	})

	// Start the bot's polling loop. DropPendingUpdates clears any old updates
	// that might have accumulated while the bot was offline.
	b.Polling().DropPendingUpdates().Start()
}
