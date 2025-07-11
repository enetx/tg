package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg"
)

func main() {
	// Read the bot token from the .env file
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	bot := tg.NewBot(token).Build().Unwrap()

	// /start command - responds to !start as well, works on edited messages
	bot.Command("start", func(ctx *tg.Context) error { return ctx.Message("Start command triggered!").Send().Err() }).
		Triggers('!').  // reacts to commands starting with '!', e.g. !start
		AllowEdited().  // allows handling of edited messages
		AllowChannel(). // allows handling of messages from channels
		Register()

	// /announce command - works in channels
	bot.Command("announce", func(ctx *tg.Context) error { return ctx.Message("Received command from channel!").Send().Err() }).
		AllowChannel().
		Register()

	// Text message handler - responds to any text, including in channels
	bot.On.Message.Text(func(ctx *tg.Context) error {
		return ctx.Message("Received a text message!").Send().Err()
	}).
		AllowChannel(). // enable in channels
		Register()

	// Start polling with allowed updates
	bot.Polling().AllowedUpdates().Start()
}
