package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg"
)

func main() {
	token := NewFile("../../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	bot := tg.NewBot(token).Build().Unwrap()

	// Handler for /start command — begins the conversation
	bot.Command("start", func(ctx *tg.Context) error {
		// Set current state to expect an email
		ctx.State.Set("get_email")
		// Ask the user to enter their email
		return ctx.Reply("Enter your email:").Send().Err()
	})

	// Handler for incoming text messages
	bot.On.Message.Text(func(ctx *tg.Context) error {
		// Get the current user state
		switch ctx.State.Get().Some() {
		// If we're expecting an email
		case "get_email":
			// Save the entered email in state data
			ctx.State.Data().Set("email", ctx.EffectiveMessage.Text)
			// Update the state to expect the user's name
			ctx.State.Set("get_name")
			// Ask for the user's name
			return ctx.Reply("Enter your name:").Send().Err()

		// If we're expecting a name
		case "get_name":
			// Retrieve the previously entered email (default to "no email" if missing)
			email := ctx.State.Data().Get("email").UnwrapOr("no email")
			// Clear the state to end the conversation
			ctx.State.Clear()
			// Reply with the collected information
			return ctx.Reply(Format("Got name: {} and email: {}", ctx.EffectiveMessage.Text, email)).Send().Err()
		}

		// Default fallback response for unexpected state
		return ctx.Reply("Please type /start to begin.").Send().Err()
	})

	// Start the bot
	bot.Polling().Start()
}
