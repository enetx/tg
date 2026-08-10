package main

import (
	"github.com/enetx/g"
	"github.com/enetx/g/fs"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
)

// Ephemeral messages (Bot API 10.2+): messages sent to a group that only one member sees.
// Set the receiver with ReceiverUser, and pass the triggering callback query id when the
// message answers a button press. Editing and deleting them needs the same receiver id.
func main() {
	token := g.Slice[g.String](fs.NewFile("../.env").Read().Ok().Trim().Split("=")).Last().Some()
	b := bot.New(token).Build().Unwrap()

	// Advertise a command whose answer is only visible to its sender.
	b.SetMyCommands().
		AddEphemeralCommand("secret", "Answer visible only to you").
		Send()

	// /secret — reply privately inside the group.
	b.Command("secret", func(ctx *ctx.Context) error {
		kb := keyboard.Inline().Text("Update", "ephemeral_update").Row().Text("Dismiss", "ephemeral_delete")

		return ctx.SendMessage(g.String("Only you can see this.")).
			ReceiverUser(ctx.EffectiveUser.Id).
			Markup(kb).
			Send().
			Err()
	})

	// Editing an ephemeral message: the receiver and the ephemeral message id default to
	// the effective message, so a callback from the message itself needs no extra ids.
	b.On.Callback.Equal("ephemeral_update", func(ctx *ctx.Context) error {
		if err := ctx.AnswerCallbackQuery(g.String("Updated")).Send().Err(); err != nil {
			return err
		}

		return ctx.EditEphemeralMessageText(g.String("Still only you.")).Send().Err()
	})

	b.On.Callback.Equal("ephemeral_delete", func(ctx *ctx.Context) error {
		return ctx.DeleteEphemeralMessage().Send().Err()
	})

	b.Polling().Start()
}
