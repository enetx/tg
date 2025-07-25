package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func main() {
	// Read the bot token from the .env file
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// /start command - responds to !start as well, works on edited messages
	b.Command("start", func(ctx *ctx.Context) error { return ctx.SendMessage("Start command triggered!").Send().Err() }).

		// reacts to commands starting with '!', e.g. !start
		Triggers('!').

		// allows handling of edited messages
		AllowEdited().

		// allows handling of messages from channels
		AllowChannel().
		Register()

	// /announce command - works in channels
	b.Command("announce", func(ctx *ctx.Context) error { return ctx.SendMessage("Received command from channel!").Send().Err() }).
		AllowChannel().
		Register()

	// Text message handler - responds to any text, including in channels
	b.On.Message.Text(func(ctx *ctx.Context) error {
		return ctx.SendMessage("Received a text message!").Send().Err()
	}).
		AllowChannel(). // enable in channels
		Register()

	// Start polling with allowed updates
	b.Polling().AllowedUpdates().Start()
}
