package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg"
	"github.com/enetx/tg/keyboard"
)

const (
	StateName     = "state_name"
	StateLike     = "state_like"
	StateLanguage = "state_language"
	StateSummary  = "state_summary"
)

func main() {
	token := NewFile("../../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	bot := tg.NewBot(token)

	tg.NewFlow(bot).
		Entry("start", StateName).
		Step(StateName, askName).
		Step(StateLike, askLike).
		Step(StateLanguage, askLanguage).
		Step(StateSummary, showSummary).
		Register("Please type /start to begin.")

	bot.Command("cancel", func(ctx *tg.Context) error {
		ctx.State.Clear()
		return ctx.Reply("Cancelled.").RemoveKeyboard().Send().Err()
	})

	bot.Polling().Start()
}

func askName(ctx *tg.Context) error {
	ctx.State.Set(StateLike)
	return ctx.Reply("Hi there! What's your name?").ForceReply().Send().Err()
}

func askLike(ctx *tg.Context) error {
	text := ctx.EffectiveMessage.Text
	ctx.State.Data().Set("name", text)
	ctx.State.Set(StateLanguage)

	return ctx.Reply(Format("Did you <b>{}</b> like writing bots?", text)).
		HTML().
		Markup(keyboard.Reply().Row().Text("Yes").Text("No")).
		Send().Err()
}

func askLanguage(ctx *tg.Context) error {
	text := ctx.EffectiveMessage.Text
	data := ctx.State.Data()
	data.Set("like", text)

	switch text {
	case "Yes":
		ctx.State.Set(StateSummary)
		return ctx.Reply("Cool! I'm too!\nWhat programming language did you use for it?").ForceReply().Send().Err()
	case "No":
		ctx.State.Clear()
		name := data.Get("name").UnwrapOr("<no name>")
		like := data.Get("like").UnwrapOr("<no choice>")

		return ctx.Reply(Format("Not bad, not terrible.\nSee you soon.\n\nSummary:\n- Name: {}\n- Like bots? {}", name, like)).
			RemoveKeyboard().
			Send().
			Err()
	default:
		ctx.State.Set(StateLanguage)
		return ctx.Reply("I don't understand you :(").Send().Err()
	}
}

func showSummary(ctx *tg.Context) error {
	data := ctx.State.Data()
	name := data.Get("name").UnwrapOr("<no name>")
	like := data.Get("like").UnwrapOr("<no choice>")
	lang := ctx.EffectiveMessage.Text

	var greeting String
	if lang == "go" {
		greeting = "Go? Nice choice – that really makes my circuits light up! 😉\n"
	}

	ctx.State.Clear()

	return ctx.Reply(Format("{}<b>Summary</b>:\n- Name: {}\n- Like bots? {}\n- Language: {}", greeting, name, like, lang)).
		HTML().
		RemoveKeyboard().
		Send().
		Err()
}
