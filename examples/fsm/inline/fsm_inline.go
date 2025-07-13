package main

import (
	"github.com/enetx/fsm"
	. "github.com/enetx/g"
	"github.com/enetx/tg"
	"github.com/enetx/tg/keyboard"
)

// FSM state identifiers
const (
	StateColor   = "color"   // Step 1: choose favorite color
	StateAnimal  = "animal"  // Step 2: choose favorite animal
	StateSummary = "summary" // Final step: show summary
)

// Global FSM store: holds per-user FSM instances
var fsmStore = NewMapSafe[int64, *fsm.FSM]()

func main() {
	// Load the Telegram bot token from .env
	token := NewFile("../../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	bot := tg.NewBot(token).Build().Unwrap()

	// Define FSM template (used per user)
	template := fsm.NewFSM(StateColor).
		Transition(StateColor, "color_selected", StateAnimal).
		Transition(StateAnimal, "animal_selected", StateSummary)

	// Step 1: ask user to select a color
	template.OnEnter(StateColor, func(ctx *fsm.Context) error {
		// Get Telegram context stored previously in /start
		tgctx := ctx.Values.Get("tgctx").Some().(*tg.Context)

		return tgctx.Reply("🎨 Choose your favorite color:").
			Markup(keyboard.Inline().
				Row().Text("❤️ Red", "color:red").
				Row().Text("💙 Blue", "color:blue").
				Row().Text("💚 Green", "color:green")).
			Send().Err()
	})

	// Step 2: ask user to select an animal
	template.OnEnter(StateAnimal, func(ctx *fsm.Context) error {
		// Get Telegram context again
		tgctx := ctx.Values.Get("tgctx").Some().(*tg.Context)
		// Color was passed as Input during transition
		color := ctx.Input.(String)
		ctx.Data.Set("color", color)

		tgctx.Reply("✅ Color selected: " + color).Send()

		return tgctx.Message("🐾 Pick your favorite animal:").
			Markup(keyboard.Inline().
				Row().Text("🐶 Dog", "animal:dog").
				Row().Text("🐱 Cat", "animal:cat").
				Row().Text("🦊 Fox", "animal:fox")).
			Send().Err()
	})

	// Step 3: show final summary
	template.OnEnter(StateSummary, func(ctx *fsm.Context) error {
		// Retrieve latest context and stored color
		tgctx := ctx.Values.Get("tgctx").Some().(*tg.Context)
		defer fsmStore.Delete(tgctx.EffectiveUser.Id) // Clear FSM after summary

		color := ctx.Data.Get("color").UnwrapOr("unknown")
		animal := ctx.Input.(String)

		tgctx.Reply("✅ Animal selected: " + animal).Send()

		return tgctx.Message(Format("🧾 Your preferences:\n- Color: {}\n- Animal: {}", color, animal)).Send().Err()
	})

	// /start command — initialize FSM and trigger first question
	bot.Command("start", func(ctx *tg.Context) error {
		entry := fsmStore.Entry(ctx.EffectiveUser.Id)
		entry.OrSetBy(template.Clone)
		fsm := entry.Get().Some()

		fsm.Context().Values.Set("tgctx", ctx) // Save Telegram context

		return fsm.CallEnter(StateColor) // Trigger first question
	})

	// /cancel command — exit flow and clear keyboard
	bot.Command("cancel", func(ctx *tg.Context) error {
		fsmStore.Delete(ctx.EffectiveUser.Id)
		return ctx.Reply("Cancelled.").RemoveKeyboard().Send().Err()
	})

	// Handle color selection via callback
	bot.On.Callback.Prefix("color:", func(ctx *tg.Context) error {
		fsmOpt := fsmStore.Get(ctx.EffectiveUser.Id)
		if fsmOpt.IsNone() {
			return ctx.Reply("Please type /start to begin.").Send().Err()
		}

		fsm := fsmOpt.Some()
		fsm.Context().Values.Set("tgctx", ctx)
		fsm.Context().Input = String(ctx.Callback.Data).StripPrefix("color:") // Pass selected color

		return fsm.Trigger("color_selected")
	})

	// Handle animal selection via callback
	bot.On.Callback.Prefix("animal:", func(ctx *tg.Context) error {
		fsmOpt := fsmStore.Get(ctx.EffectiveUser.Id)
		if fsmOpt.IsNone() {
			return ctx.Reply("Please type /start to begin.").Send().Err()
		}

		fsm := fsmOpt.Some()
		fsm.Context().Values.Set("tgctx", ctx)
		fsm.Context().Input = String(ctx.Callback.Data).StripPrefix("animal:") // Pass selected animal

		return fsm.Trigger("animal_selected")
	})

	// Cleanup handler — remove unsupported interactions
	bot.On.Any(func(ctx *tg.Context) error {
		if ctx.Callback == nil {
			return ctx.Delete().Send().Err()
		}
		return nil
	})

	// Optional: remove pending updates before starting polling
	bot.Polling().DropPendingUpdates().Start()
}
