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

	// Initialize the Telegram b and its helper components.
	b := bot.New(token).Build().Unwrap()

	// Create a master FSM template. Each new user will receive a clone of this template.
	// This ensures a consistent workflow while maintaining separate states and data for each user.
	template := fsm.NewFSM(StateGender).
		// Define the linear flow of the conversation.
		Transition(StateGender, "next", StatePhoto).
		Transition(StatePhoto, "next", StateLocation).
		Transition(StateLocation, "next", StateDone)

	// Step 1: Ask for gender. This is the entry point for a new workflow.
	template.OnEnter(StateGender, func(fctx *fsm.Context) error {
		// Retrieve the ctx.Context passed through FSM Values to interact with the Telegram API.
		tgctx := fctx.Values.Get("tgctx").Some().(*ctx.Context)

		// Ask for the user's gender using a custom reply keyboard for easy input.
		return tgctx.Reply("Are you a boy or a girl?").
			Markup(keyboard.Reply().Row().Text("Boy").Text("Girl").Text("Other")).
			Send().Err()
	})

	// Step 2: Ask for a photo after the user has provided their gender.
	template.OnEnter(StatePhoto, func(fctx *fsm.Context) error {
		tgctx := fctx.Values.Get("tgctx").Some().(*ctx.Context)

		// Save the gender from the previous step's input.
		fctx.Data.Set("gender", tgctx.EffectiveMessage.Text)

		// Prompt the user to either send a photo or skip this optional step.
		return tgctx.Reply("Send me your photo or type /skip").RemoveKeyboard().Send().Err()
	})

	// Step 3: Ask for location after the user has sent a photo or skipped.
	template.OnEnter(StateLocation, func(fctx *fsm.Context) error {
		tgctx := fctx.Values.Get("tgctx").Some().(*ctx.Context)

		// Check if a photo was sent in the previous step and save it.
		if tgctx.EffectiveMessage.Photo != nil {
			fctx.Data.Set("photo", tgctx.EffectiveMessage.Photo)
			tgctx.Reply("✅ Photo received").Send()
		} else if fctx.Data.Get("photo").UnwrapOrDefault() == "skipped" {
			// Acknowledge that the photo step was skipped.
			tgctx.Reply("⏭ Photo skipped").Send()
		}

		// Ask the user to share their location, providing a keyboard button for convenience.
		return tgctx.Reply("Now, share your location or type /skip").
			Markup(keyboard.Reply().Location("Location")).Send().Err()
	})

	// This hook processes the final user input (location) before transitioning to the 'done' state.
	// It acts as the final data processing step before the summary is shown.
	template.OnExit(StateLocation, func(fctx *fsm.Context) error {
		tgctx := fctx.Values.Get("tgctx").Some().(*ctx.Context)

		// Check if a location was provided and save it.
		if tgctx.EffectiveMessage.Location != nil {
			fctx.Data.Set("location", tgctx.EffectiveMessage.Location)
			tgctx.Reply("✅ Location received").Send()
		} else if fctx.Data.Get("location").UnwrapOrDefault() == "skipped" {
			// Acknowledge that the location step was skipped.
			tgctx.Reply("⏭ Location skipped").Send()
		}

		// Inform the user that the data collection is complete and a summary will be shown.
		return tgctx.Message("Thanks! Let me summarize what you've told me...").RemoveKeyboard().Send().Err()
	})

	// Final Step: Display a summary of all collected data and clean up the user's FSM instance.
	template.OnEnter(StateDone, func(fctx *fsm.Context) error {
		tgctx := fctx.Values.Get("tgctx").Some().(*ctx.Context)
		data := fctx.Data

		// Retrieve all collected data from the FSM's persistent storage.
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

		// Crucial: Remove the user's FSM instance from the store to free up memory
		// and allow them to start a new session later.
		fsmStore.Delete(tgctx.EffectiveUser.Id)

		// Send the final goodbye message.
		return tgctx.Message("Thank you! I hope we can talk again some day.").Send().Err()
	})

	// Command handler for /start, which initializes or resets a user's workflow.
	b.Command("start", func(ctx *ctx.Context) error {
		// Get or create an FSM instance for the user.
		entry := fsmStore.Entry(ctx.EffectiveUser.Id)
		// If the user is new, clone the master template for them.
		entry.OrSetBy(template.Clone)
		fsm := entry.Get().Some()

		// Store the current Telegram context in the FSM's temporary Values.
		// This makes it accessible within all callback functions for this interaction.
		fsm.Context().Values.Set("tgctx", ctx)

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
		fsm.Context().Values.Set("tgctx", ctx)

		// Check the user's current state to determine if skipping is allowed.
		switch fsm.Current() {
		case StatePhoto:
			fsm.Context().Data.Set("photo", "skipped")
			return fsm.Trigger("next")
		case StateLocation:
			fsm.Context().Data.Set("location", "skipped")
			return fsm.Trigger("next")
		default:
			// If the user is not in a skippable state, inform them.
			return ctx.Reply("Nothing to skip.").Send().Err()
		}
	})

	// Command handler for /cancel to prematurely end the workflow and clean up.
	b.Command("cancel", func(ctx *ctx.Context) error {
		fsmStore.Delete(ctx.EffectiveUser.Id)
		return ctx.Reply("Bye! I hope we can talk again some day.").Send().Err()
	})

	// Generic message handler for text, photo, and location inputs.
	// This function handles the common logic for advancing the FSM.
	handleMessage := func(ctx *ctx.Context) error {
		// Attempt to retrieve the user's FSM instance.
		fsmOpt := fsmStore.Get(ctx.EffectiveUser.Id)
		if fsmOpt.IsNone() {
			// If no FSM exists, the user has not started the workflow.
			return ctx.Reply("Please type /start to begin.").Send().Err()
		}

		fsm := fsmOpt.Some()

		// Store the latest Telegram context in the FSM's Values.
		fsm.Context().Values.Set("tgctx", ctx)

		// Trigger the "next" event to advance the state machine.
		// The FSM's OnEnter/OnExit callbacks will handle the specific logic.
		return fsm.Trigger("next")
	}

	// Register the generic handler for different message types.
	b.On.Message.Text(handleMessage)
	b.On.Message.Photo(handleMessage)
	b.On.Message.Location(handleMessage)

	// Start the bot's polling loop to listen for updates from Telegram.
	b.Polling().Start()
}
