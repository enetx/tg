package main

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/fsm"
	. "github.com/enetx/g"
	"github.com/enetx/tg"
	"github.com/enetx/tg/keyboard"
)

// FSM state identifiers
const (
	StateGender   = "gender"   // Step 1: choose gender
	StatePhoto    = "photo"    // Step 2: send or skip photo
	StateLocation = "location" // Step 3: send or skip location
	StateBio      = "bio"      // Step 4: provide bio
	StateSummary  = "summary"  // Final step: show summary
	StateDone     = "done"     // Completion
)

// fsmStore holds the FSM instance per user
var fsmStore = NewMapSafe[int64, *fsm.FSM]()

func main() {
	// Load bot token from .env file
	token := NewFile("../../../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	bot := tg.NewBot(token).Build().Unwrap()

	// Create FSM template with transition definitions
	template := fsm.NewFSM(StateGender).
		Transition(StateGender, "next", StatePhoto).
		Transition(StatePhoto, "next", StateLocation).
		Transition(StateLocation, "next", StateBio).
		Transition(StateBio, "next", StateSummary).
		Transition(StateSummary, "next", StateDone)

	// Step: Ask for gender input
	template.OnEnter(StateGender, func(ctx *fsm.Context) error {
		// Retrieve tg.Context passed through FSM Values
		tgctx := ctx.Values.Get("tgctx").Some().(*tg.Context)

		// Ask gender using reply keyboard
		return tgctx.Reply("Are you a boy or a girl?").
			Markup(keyboard.Reply().Row().Text("Boy").Text("Girl").Text("Other")).
			Send().Err()
	})

	// Step: Ask for photo after saving gender input
	template.OnEnter(StatePhoto, func(ctx *fsm.Context) error {
		tgctx := ctx.Values.Get("tgctx").Some().(*tg.Context)
		ctx.Data.Set("gender", tgctx.EffectiveMessage.Text)

		// Prompt user to send photo or skip
		return tgctx.Reply("Send me your photo or type /skip").RemoveKeyboard().Send().Err()
	})

	// Step: Ask for location after processing photo
	template.OnEnter(StateLocation, func(ctx *fsm.Context) error {
		tgctx := ctx.Values.Get("tgctx").Some().(*tg.Context)

		// Save photo if sent
		if tgctx.EffectiveMessage.Photo != nil {
			ctx.Data.Set("photo", tgctx.EffectiveMessage.Photo)
			tgctx.Reply("✅ Photo received").Send()
		} else if ctx.Data.Get("photo").UnwrapOrDefault() == "skipped" {
			// Notify if skipped
			tgctx.Reply("⏭ Photo skipped").Send()
		}

		// Ask user to share location
		return tgctx.Reply("Now, share your location or type /skip").
			Markup(keyboard.Reply().Location("Location")).Send().Err()
	})

	// Step: Ask for bio and auto-trigger next
	template.OnEnter(StateBio, func(ctx *fsm.Context) error {
		// Retrieve tg.Context for response access
		tgctx := ctx.Values.Get("tgctx").Some().(*tg.Context)

		// Save location if provided
		if tgctx.EffectiveMessage.Location != nil {
			ctx.Data.Set("location", tgctx.EffectiveMessage.Location)
			tgctx.Reply("✅ Location received").Send()
		} else if ctx.Data.Get("location").UnwrapOrDefault() == "skipped" {
			tgctx.Reply("⏭ Location skipped").Send()
		}

		// Prompt final summary message
		tgctx.Message("Thanks! Let me summarize what you've told me...").RemoveKeyboard().Send()

		// Retrieve FSM instance to trigger next state
		fsm := ctx.Values.Get("fsm").Some().(*fsm.FSM)
		return fsm.Trigger("next")
	})

	// Step: Display collected summary
	template.OnEnter(StateSummary, func(ctx *fsm.Context) error {
		tgctx := ctx.Values.Get("tgctx").Some().(*tg.Context)
		data := ctx.Data

		// Retrieve stored user data
		gender := data.Get("gender").UnwrapOr("unknown")
		photo := data.Get("photo")
		location := data.Get("location")

		// Re-send user photo (if any)
		if photo.IsSome() {
			if sizes, ok := photo.Some().([]gotgbot.PhotoSize); ok && len(sizes) > 0 {
				fileID := sizes[len(sizes)-1].FileId
				tgctx.Photo(String(fileID).Prepend(tg.FileIDPrefix)).Caption("Your photo").Send()
			}
		}

		// Re-send location (if any)
		if location.IsSome() {
			if loc, ok := location.Some().(*gotgbot.Location); ok {
				tgctx.Bot.Raw.SendLocation(tgctx.EffectiveChat.Id, loc.Latitude, loc.Longitude, nil)
			}
		}

		// Compose text summary
		summary := "🧾 Summary:\n"
		summary += "👤 Gender: " + gender.(string) + "\n"

		tgctx.Message(String(summary)).Send()

		// Proceed to done
		fsm := ctx.Values.Get("fsm").Some().(*fsm.FSM)
		return fsm.Trigger("next")
	})

	// Step: Completion and cleanup
	template.OnEnter(StateDone, func(ctx *fsm.Context) error {
		tgctx := ctx.Values.Get("tgctx").Some().(*tg.Context)
		fsmStore.Delete(tgctx.EffectiveUser.Id)

		return tgctx.Message("Thank you! I hope we can talk again some day.").Send().Err()
	})

	// /start initializes FSM and starts flow
	bot.Command("start", func(ctx *tg.Context) error {
		entry := fsmStore.Entry(ctx.EffectiveUser.Id)
		entry.OrSetBy(template.Clone)
		fsm := entry.Get().Some()

		// Store Telegram context and FSM instance in values
		fsm.Context().Values.Set("tgctx", ctx)
		fsm.Context().Values.Set("fsm", fsm)

		return fsm.CallEnter(StateGender)
	})

	// /skip allows skipping optional steps like photo or location
	bot.Command("skip", func(ctx *tg.Context) error {
		fsmOpt := fsmStore.Get(ctx.EffectiveUser.Id)
		if fsmOpt.IsNone() {
			return ctx.Reply("Nothing to skip. Please type /start.").Send().Err()
		}

		fsm := fsmOpt.Some()
		fsm.Context().Values.Set("tgctx", ctx)
		fsm.Context().Values.Set("fsm", fsm)

		switch fsm.Current() {
		case StatePhoto:
			fsm.Context().Data.Set("photo", "skipped")
			return fsm.Trigger("next")
		case StateLocation:
			fsm.Context().Data.Set("location", "skipped")
			return fsm.Trigger("next")
		default:
			return ctx.Reply("Nothing to skip.").Send().Err()
		}
	})

	// /cancel clears current flow
	bot.Command("cancel", func(ctx *tg.Context) error {
		fsmStore.Delete(ctx.EffectiveUser.Id)
		return ctx.Reply("Bye! I hope we can talk again some day.").Send().Err()
	})

	// Handle text input and trigger FSM
	bot.On.Message.Text(func(ctx *tg.Context) error {
		fsmOpt := fsmStore.Get(ctx.EffectiveUser.Id)
		if fsmOpt.IsNone() {
			return ctx.Reply("Please type /start to begin.").Send().Err()
		}

		fsm := fsmOpt.Some()
		fsm.Context().Input = ctx.EffectiveMessage.Text
		fsm.Context().Values.Set("tgctx", ctx)
		fsm.Context().Values.Set("fsm", fsm)

		return fsm.Trigger("next")
	})

	// Handle photo input and trigger FSM
	bot.On.Message.Photo(func(ctx *tg.Context) error {
		fsmOpt := fsmStore.Get(ctx.EffectiveUser.Id)
		if fsmOpt.IsNone() {
			return ctx.Reply("Please type /start to begin.").Send().Err()
		}

		fsm := fsmOpt.Some()
		fsm.Context().Values.Set("tgctx", ctx)
		fsm.Context().Values.Set("fsm", fsm)

		return fsm.Trigger("next")
	})

	// Handle location input and trigger FSM
	bot.On.Message.Location(func(ctx *tg.Context) error {
		fsmOpt := fsmStore.Get(ctx.EffectiveUser.Id)
		if fsmOpt.IsNone() {
			return ctx.Reply("Please type /start to begin.").Send().Err()
		}

		fsm := fsmOpt.Some()
		fsm.Context().Values.Set("tgctx", ctx)
		fsm.Context().Values.Set("fsm", fsm)

		return fsm.Trigger("next")
	})

	// Start polling updates
	bot.Polling().Start()
}
