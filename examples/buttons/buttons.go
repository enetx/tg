package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg"
	"github.com/enetx/tg/keyboard"
)

func main() {
	token := NewFile("../../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	bot := tg.NewBot(token).Build().Unwrap()

	bot.Command("start", func(ctx *tg.Context) error {
		markup := keyboard.Inline().
			Row().
			Text("Callback1", "cb_1").
			Text("Callback2", "cb_2").
			Row().
			Text("Edit Callback1", "edit_cb1").
			Text("Remove Callback1", "delete_cb1").
			Row().
			Text("Remove Buttons", "clear")

		return ctx.Reply("Choose a button:").Markup(markup).Send().Err()
	})

	bot.On.Callback.Equal("cb_1", func(ctx *tg.Context) error {
		return ctx.Answer("clicked the callback1 button").Send().Err()
	})

	bot.On.Callback.Equal("cb_2", func(ctx *tg.Context) error {
		return ctx.Answer("clicked the callback2 button").Alert().Send().Err()
	})

	bot.On.Callback.Equal("edit_cb1", func(ctx *tg.Context) error {
		markup := keyboard.Inline(ctx.EffectiveMessage.ReplyMarkup).
			Edit(func(btn *keyboard.Button) {
				switch btn.Get.Callback() {
				case "cb_1":
					btn.Text("Edited 1")
				case "edit_cb1":
					btn.Delete()
				}
			})

		return ctx.EditMarkup(markup).Send().Err()
	})

	bot.On.Callback.Equal("delete_cb1", func(ctx *tg.Context) error {
		markup := keyboard.Inline(ctx.EffectiveMessage.ReplyMarkup).
			Edit(func(btn *keyboard.Button) {
				cb := btn.Get.Callback()

				if cb == "cb_1" || cb == "delete_cb1" {
					btn.Delete()
				}

				if btn.Get.Text() == "Edited 1" {
					btn.Delete()
				}
			})

		return ctx.EditMarkup(markup).Send().Err()
	})

	bot.On.Callback.Equal("clear", func(ctx *tg.Context) error {
		return ctx.EditMarkup(nil).Send().Err()
	})

	bot.Polling().AllowedUpdates().Start()
}
