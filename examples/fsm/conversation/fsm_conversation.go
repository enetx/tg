package main

import (
	"github.com/PaulSonOfLars/gotgbot/v2"

	"github.com/enetx/fsm"
	. "github.com/enetx/g"

	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/constants"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
)

// FSM state identifiers define the steps of the user registration workflow.
const (
	StateGender   = "gender"   // Step 1: choose gender
	StatePhoto    = "photo"    // Step 2: send or skip photo
	StateLocation = "location" // Step 3: send or skip location
	StateDone     = "done"     // Final state for completion and cleanup
)

// fsmStore holds the active FSM instance for each user, keyed by their Telegram user ID.
// This allows each user to have their own independent state in the conversation.
var fsmStore = NewMapSafe[int64, *fsm.FSM]()

func main() {
	// Load the bot token from a local .env file.
	token := NewFile("../../../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()

	// Initialize the Telegram bot and its helper components.
	b := bot.New(token).Build().Unwrap()

	// Create a master FSM fsmachine. Each new user will receive a clone of this fsmachine.
	// This ensures a consistent workflow while maintaining separate states and data for each user.
	fsmachine := fsm.NewFSM(StateGender).
		// Define the linear flow of the conversation. Each "next" event moves to the subsequent state.
		Transition(StateGender, "next", StatePhoto).
		Transition(StatePhoto, "next", StateLocation).
		Transition(StateLocation, "next", StateDone)

	// Callback for entering StateGender. This is the entry point for a new workflow.
	fsmachine.OnEnter(StateGender, func(fctx *fsm.Context) error {
		// Retrieve the Telegram context (`tgctx`) stored in the FSM's Meta store.
		// This context is needed to send replies to the user.
		tgctx := fctx.Meta.Get("tgctx").Some().(*ctx.Context)

		// Ask for the user's gender using a custom reply keyboard for easy input.
		return tgctx.Reply("Are you a boy or a girl?").
			Markup(keyboard.Reply().Row().Text("Boy").Text("Girl").Text("Other")).
			Send().Err()
	})

	// Callback for entering StatePhoto, executed after the user has provided their gender.
	fsmachine.OnEnter(StatePhoto, func(fctx *fsm.Context) error {
		// Retrieve the user's gender from `fctx.Input`. The text handler passed this value
		// when it called `fsm.Trigger("next", userText)`.
		gender := fctx.Input.(string)
		// Store the gender in the FSM's persistent Data store for later use in the summary.
		fctx.Data.Set("gender", gender)

		// Retrieve the tgctx to send the next prompt.
		tgctx := fctx.Meta.Get("tgctx").Some().(*ctx.Context)

		// Prompt the user to either send a photo or skip this optional step.
		return tgctx.Reply("Send me your photo or type /skip").RemoveKeyboard().Send().Err()
	})

	// Callback for entering StateLocation, executed after the user has sent a photo or skipped.
	fsmachine.OnEnter(StateLocation, func(fctx *fsm.Context) error {
		// Process the input from the 'photo' state. This input can be actual photo data
		// (passed by the photo handler) or a "skipped" string (passed by the /skip handler).
		photoInput := fctx.Input
		tgctx := fctx.Meta.Get("tgctx").Some().(*ctx.Context)

		// Check the type of input to determine if a photo was sent.
		if photo, ok := photoInput.([]gotgbot.PhotoSize); ok {
			fctx.Data.Set("photo", photo)
			tgctx.Reply("✅ Photo received").Send()
		} else if skipped, ok := photoInput.(string); ok && skipped == "skipped" {
			// Acknowledge that the photo step was skipped.
			tgctx.Reply("⏭ Photo skipped").Send()
		}

		// Ask the user to share their location, providing a keyboard button for convenience.
		return tgctx.Reply("Now, share your location or type /skip").
			Markup(keyboard.Reply().Location("Location")).Send().Err()
	})

	// This OnExit hook processes the user's final input (their location or a skip command).
	// It runs *after* the user has responded in the 'location' state but *before* the 'done' state's OnEnter callback.
	fsmachine.OnExit(StateLocation, func(fctx *fsm.Context) error {
		tgctx := fctx.Meta.Get("tgctx").Some().(*ctx.Context)

		// Retrieve the final input (location data or "skipped" string) from the context.
		locationInput := fctx.Input
		if loc, ok := locationInput.(*gotgbot.Location); ok {
			fctx.Data.Set("location", loc)
			tgctx.Reply("✅ Location received").Send()
		} else if skipped, ok := locationInput.(string); ok && skipped == "skipped" {
			fctx.Data.Set("location", "skipped")
			tgctx.Reply("⏭ Location skipped").Send()
		}

		// Inform the user that data collection is complete before showing the summary.
		return tgctx.Message("Thanks! Let me summarize what you've told me...").RemoveKeyboard().Send().Err()
	})

	// Callback for entering StateDone. This is the final step where the summary is displayed.
	fsmachine.OnEnter(StateDone, func(fctx *fsm.Context) error {
		tgctx := fctx.Meta.Get("tgctx").Some().(*ctx.Context)

		// Use `defer` to ensure the FSM instance is removed from the store after this function completes.
		// This frees memory and allows the user to run /start again.
		defer fsmStore.Delete(tgctx.EffectiveUser.Id)

		// Retrieve all collected data from the FSM's persistent storage.
		data := fctx.Data
		gender := data.Get("gender").UnwrapOr("unknown")
		photo := data.Get("photo")
		location := data.Get("location")

		// If a photo was provided, re-send the highest resolution version.
		if photo.IsSome() {
			if sizes, ok := photo.Some().([]gotgbot.PhotoSize); ok && len(sizes) > 0 {
				fileID := sizes[len(sizes)-1].FileId
				tgctx.Photo(String(fileID).Prepend(constants.FileIDPrefix)).Caption("Your photo").Send()
			}
		}

		// If a location was provided, re-send it as a map pin.
		if location.IsSome() {
			if loc, ok := location.Some().(*gotgbot.Location); ok {
				tgctx.Bot.Raw().SendLocation(tgctx.EffectiveChat.Id, loc.Latitude, loc.Longitude, nil)
			}
		}

		// Compose and send the final text summary.
		summary := "🧾 Summary:\n"
		summary += "👤 Gender: " + gender.(string) + "\n"
		tgctx.Message(String(summary)).Send()

		// Send the final goodbye message.
		return tgctx.Message("Thank you! I hope we can talk again some day.").Send().Err()
	})

	// Command handler for /start, which initializes or resets a user's workflow.
	b.Command("start", func(ctx *ctx.Context) error {
		// Get or create an FSM instance for the user.
		entry := fsmStore.Entry(ctx.EffectiveUser.Id)
		// If the user is new, clone the master template for them.
		entry.OrSetBy(fsmachine.Clone)
		fsm := entry.Get().Some()

		// Store the current Telegram context in the FSM's temporary Meta store.
		// This makes it accessible within all callbacks for sending API replies.
		fsm.Context().Meta.Set("tgctx", ctx)

		// Manually trigger the entry callback for the first state to begin the flow.
		return fsm.CallEnter(StateGender)
	})

	// Command handler for /skip, allowing users to bypass optional steps.
	b.Command("skip", func(ctx *ctx.Context) error {
		fsmOpt := fsmStore.Get(ctx.EffectiveUser.Id)
		if fsmOpt.IsNone() {
			return ctx.Reply("Nothing to skip. Please type /start.").Send().Err()
		}

		fsm := fsmOpt.Some()
		fsm.Context().Meta.Set("tgctx", ctx)

		// Trigger the FSM's "next" event, passing the string "skipped" as the input.
		// The appropriate OnEnter/OnExit callback will handle this special value.
		return fsm.Trigger("next", "skipped")
	})

	// Command handler for /cancel to prematurely end the workflow and clean up.
	b.Command("cancel", func(ctx *ctx.Context) error {
		fsmStore.Delete(ctx.EffectiveUser.Id)
		return ctx.Reply("Bye! I hope we can talk again some day.").Send().Err()
	})

	// Handler for incoming text messages.
	b.On.Message.Text(func(ctx *ctx.Context) error {
		fsmOpt := fsmStore.Get(ctx.EffectiveUser.Id)
		if fsmOpt.IsNone() {
			return ctx.Reply("Please type /start to begin.").Send().Err()
		}

		fsm := fsmOpt.Some()
		fsm.Context().Meta.Set("tgctx", ctx)

		// Trigger the "next" event, passing the message text directly as the FSM input.
		return fsm.Trigger("next", ctx.EffectiveMessage.Text)
	})

	// Handler for incoming photos.
	b.On.Message.Photo(func(ctx *ctx.Context) error {
		fsmOpt := fsmStore.Get(ctx.EffectiveUser.Id)
		if fsmOpt.IsNone() {
			return nil // Ignore photos if not in a workflow.
		}

		fsm := fsmOpt.Some()
		fsm.Context().Meta.Set("tgctx", ctx)

		// Trigger the "next" event, passing the photo data slice as the FSM input.
		return fsm.Trigger("next", ctx.EffectiveMessage.Photo)
	})

	// Handler for incoming locations.
	b.On.Message.Location(func(ctx *ctx.Context) error {
		fsmOpt := fsmStore.Get(ctx.EffectiveUser.Id)
		if fsmOpt.IsNone() {
			return nil // Ignore locations if not in a workflow.
		}

		fsm := fsmOpt.Some()
		fsm.Context().Meta.Set("tgctx", ctx)

		// Trigger the "next" event, passing the location object as the FSM input.
		return fsm.Trigger("next", ctx.EffectiveMessage.Location)
	})

	// Start the bot's polling loop to listen for updates from Telegram.
	b.Polling().Start()
}
