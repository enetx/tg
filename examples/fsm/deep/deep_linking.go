// # More info on what deep linking actually is (read this first if it's unclear to you):
// # https://core.telegram.org/bots/features#deep-linking

package main

import (
	"github.com/enetx/fsm"
	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/preview"
)

const (
	CheckThisOut   = "check-this-out"
	SoCool         = "so-cool"
	UsingEntities  = "using-entities-here"
	UsingKeyboard  = "using-keyboard-here"
	CallbackButton = "keyboard-callback-data"
)

var (
	// fsmStore keeps a per-user FSM instance
	fsmStore = g.NewMapSafe[int64, *fsm.SyncFSM]()

	// furl is the template for Telegram deep-link URLs
	furl = "https://t.me/{.Bot.Raw.Username}?start={}"
)

func main() {
	// Read the bot token from a .env file
	token := g.NewFile("../../../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()

	// Create the bot instance
	b := bot.New(token).Build().Unwrap()

	// Define the FSM template (will be cloned for each user)
	fsmachine := fsm.New("start").
		Transition("start", CheckThisOut, "deep1").
		Transition("deep1", SoCool, "deep2").
		Transition("deep2", UsingEntities, "deep3").
		Transition("deep3", UsingKeyboard, "deep4").
		OnEnter("deep1", handleDeep1).
		OnEnter("deep2", handleDeep2).
		OnEnter("deep3", handleDeep3).
		OnEnter("deep4", handleDeep4)

	// Команда /start
	b.Command("start", func(ctx *ctx.Context) error {
		// Retrieve or initialize FSM instance for the current user
		entry := fsmStore.Entry(ctx.EffectiveUser.Id)
		entry.OrSetBy(func() *fsm.SyncFSM { return fsmachine.Clone().Sync() })
		state := entry.Get().Some()

		// Extract payload from /start {payload}
		payload := ctx.Args().Last().UnwrapOrDefault().Trim()

		// Store the current Telegram context in FSM.Meta to access inside handlers
		state.Context().Meta.Set("tgctx", ctx)

		// Try to trigger a transition with the payload as the event
		if err := state.Trigger(fsm.Event(payload), payload); err != nil {
			return handleDefault(ctx)
		}

		return nil
	})

	// Handle callback button press
	b.On.Callback.Equal(CallbackButton, func(tgctx *ctx.Context) error {
		url := g.Format(furl, tgctx.Bot, UsingKeyboard)
		return tgctx.AnswerCallbackQuery("").URL(url).Send().Err()
	})

	// Start polling for updates
	b.Polling().DropPendingUpdates().Start()
}

// Fallback handler for unknown payloads or transitions
func handleDefault(c *ctx.Context) error {
	url := g.Format(furl, c, CheckThisOut)
	return c.SendMessage(g.String("Feel free to tell your friends about it:\n\n" + url)).Send().Err()
}

// FSM: Entered state "deep1"
func handleDeep1(fctx *fsm.Context) error {
	tgctx := fctx.Meta.Get("tgctx").Some().(*ctx.Context)
	url := g.Format(furl, tgctx, SoCool)

	return tgctx.SendMessage("Awesome, you just accessed hidden functionality! Now let's get back to the private chat.").
		Markup(keyboard.Inline().URL("Continue here!", url)).
		Send().
		Err()
}

// FSM: Entered state "deep2"
func handleDeep2(fctx *fsm.Context) error {
	tgctx := fctx.Meta.Get("tgctx").Some().(*ctx.Context)
	url := g.Format(furl, tgctx, UsingEntities)
	msg := g.Format(`You can also mask the deep-linked URLs as links: <a href="{}">▶️ CLICK HERE</a>`, url)

	return tgctx.SendMessage(msg).HTML().Preview(preview.New().Disable()).Send().Err()
}

// FSM: Entered state "deep3"
func handleDeep3(fctx *fsm.Context) error {
	tgctx := fctx.Meta.Get("tgctx").Some().(*ctx.Context)
	return tgctx.SendMessage("It is also possible to make deep-linking using InlineKeyboardButtons.").
		Markup(keyboard.Inline().Text("Like this!", CallbackButton)).Send().Err()
}

// FSM: Entered state "deep4" (final state)
func handleDeep4(fctx *fsm.Context) error {
	tgctx := fctx.Meta.Get("tgctx").Some().(*ctx.Context)

	// Clear FSM for the user (optional, since this is a final state)
	defer fsmStore.Delete(tgctx.EffectiveUser.Id)

	return tgctx.SendMessage(g.Format("Congratulations! This is as deep as it gets \n\nThe payload was: {}", fctx.Input)).
		Send().
		Err()
}
