package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg"
	"github.com/enetx/tg/keyboard"
)

const (
	stepColor   = "color"
	stepAnimal  = "animal"
	stepSummary = "summary"
)

func main() {
	token := NewFile("../../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	bot := tg.NewBot(token)

	tg.NewFlow(bot).
		Entry("start", stepColor).
		Step(stepColor, handleColor).
		Step(stepAnimal, handleAnimal).
		Step(stepSummary, handleFinish).
		Register("Please type /start to begin.")

	bot.Command("cancel", func(ctx *tg.Context) error {
		ctx.State.Clear()
		return ctx.Reply("Cancelled.").RemoveKeyboard().Send().Err()
	})

	bot.Polling().DropPendingUpdates().Start()
}

func handleColor(ctx *tg.Context) error {
	ctx.State.Set(stepAnimal)

	return ctx.Reply("🎨 Choose your favorite color:").
		Markup(keyboard.Inline().
			Row().Text("❤️ Red", "color:red").
			Row().Text("💙 Blue", "color:blue").
			Row().Text("💚 Green", "color:green")).
		Send().Err()
}

func handleAnimal(ctx *tg.Context) error {
	if ctx.Callback == nil {
		ctx.Delete()
		return nil
	}

	if String(ctx.Callback.Data).StartsWith("color:") {
		ctx.State.Set(stepSummary)

		color := String(ctx.Callback.Data).StripPrefix("color:")
		ctx.State.Data().Set("color", color)
		ctx.Reply("✅ Color selected: " + color).Send()
	}

	return ctx.Message("🐾 Pick your favorite animal:").
		Markup(keyboard.Inline().
			Row().Text("🐶 Dog", "animal:dog").
			Row().Text("🐱 Cat", "animal:cat").
			Row().Text("🦊 Fox", "animal:fox")).
		Send().Err()
}

func handleFinish(ctx *tg.Context) error {
	if ctx.Callback == nil {
		ctx.Delete()
		return nil
	}

	if String(ctx.Callback.Data).StartsWith("animal:") {
		animal := String(ctx.Callback.Data).StripPrefix("animal:")
		ctx.State.Data().Set("animal", animal)
		ctx.Reply("✅ Animal selected: " + animal).Send()
	}

	data := ctx.State.Data()
	color := data.Get("color").UnwrapOr("unknown")
	animal := data.Get("animal").UnwrapOr("unknown")

	ctx.State.Clear()

	return ctx.Message(Format("🧾 Your preferences:\n- Color: {}\n- Animal: {}", color, animal)).Send().Err()
}
