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
		ctx.State.Set("name")
		ctx.State.Data().Clear()
		return ctx.Reply("Hi there! What's your name?").ForceReply().Send().Err()
	})

	bot.Command("cancel", func(ctx *tg.Context) error {
		ctx.State.Clear()
		return ctx.Reply("Cancelled.").RemoveKeyboard().Send().Err()
	})

	bot.On.Message.Text(func(ctx *tg.Context) error {
		text := String(ctx.EffectiveMessage.Text)
		state := ctx.State.Get().Some()

		switch state {
		case "name":
			ctx.State.Data().Set("name", text)
			ctx.State.Set("like")

			return ctx.Reply(Format("Did you <b>{}</b> like writing bots?", text)).
				HTML().
				Markup(keyboard.Reply().
					Row().
					Text("Yes").
					Text("No")).Send().Err()
		case "like":
			ctx.State.Data().Set("like", text)

			switch text {
			case "Yes":
				ctx.State.Set("language")
				return ctx.Reply("Cool! I'm too!\nWhat programming language did you use for it?").
					ForceReply().Send().Err()
			case "No":
				data := ctx.State.Data()
				name := data.Get("name").UnwrapOr("<no name>")
				like := data.Get("like").UnwrapOr("<no choice>")
				ctx.State.Clear()

				return ctx.Reply(Format("Not bad, not terrible.\nSee you soon.\n\nSummary:\n- Name: {}\n- Like bots? {}", name, like)).
					RemoveKeyboard().
					Send().
					Err()
			default:
				return ctx.Reply("I don't understand you :(").Send().Err()
			}
		case "language":
			data := ctx.State.Data()
			defer ctx.State.Clear()

			name := data.Get("name").UnwrapOr("<no name>")
			like := data.Get("like").UnwrapOr("<no choice>")
			lang := text

			var greeting String
			if lang.EqFold("go") {
				greeting = "Go? Nice choice – that really makes my circuits light up! 😉\n"
			}

			summary := Format("{}Summary:\n- Name: {}\n- Like bots? {}\n- Language: {}", greeting, name, like, lang)
			return ctx.Reply(summary).RemoveKeyboard().Send().Err()
		default:
			ctx.State.Clear()
			return ctx.Reply("Please type /start to begin.").RemoveKeyboard().Send().Err()
		}
	})

	bot.Polling().Start()
}
