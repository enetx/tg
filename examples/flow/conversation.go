package main

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg"
	"github.com/enetx/tg/keyboard"
)

const (
	gender   = "gender"
	photo    = "photo"
	location = "location"
	bio      = "bio"
	summary  = "summary"
)

func main() {
	token := NewFile("../../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	bot := tg.NewBot(token)

	flow := tg.NewFlow(bot)

	flow.Entry("start", gender).
		Step(gender, handleGender).
		Step(photo, handlePhoto).
		Step(location, handleLocation).
		Step(bio, handleBio).
		Step(summary, handleSummary).
		Step("done", done).
		Register("Please type /start to begin.")

	bot.Command("skip", handleSkip)
	bot.Command("cancel", handleCancel)

	bot.Polling().Start()
}

// Handlers

func handleGender(ctx *tg.Context) error {
	ctx.State.Set(photo)

	return ctx.Reply("Are you a boy or a girl?").
		Markup(keyboard.Reply().
			Row().
			Text("Boy").
			Text("Girl").
			Text("Other")).
		Send().Err()
}

func handlePhoto(ctx *tg.Context) error {
	ctx.State.Data().Set("gender", ctx.EffectiveMessage.Text)
	ctx.State.Set(location)

	return ctx.Reply("Send me your photo or type /skip").RemoveKeyboard().Send().Err()
}

func handleLocation(ctx *tg.Context) error {
	if ctx.EffectiveMessage.Photo != nil {
		ctx.State.Data().Set("photo", ctx.EffectiveMessage.Photo)
	}

	ctx.State.Set(bio)

	return ctx.Reply("Now, share your location or type /skip").
		Markup(keyboard.Reply().Location("Location")).Send().Err()
}

func handleBio(ctx *tg.Context) error {
	if ctx.EffectiveMessage.Location != nil {
		ctx.State.Data().Set("location", ctx.EffectiveMessage.Location)
	}

	defer ctx.State.Jump(summary)

	return ctx.Reply("Thanks! Let me summarize what you've told me...").RemoveKeyboard().Send().Err()
}

func handleSummary(ctx *tg.Context) error {
	data := ctx.State.Data()

	gender := data.Get("gender").UnwrapOr("unknown")
	photo := data.Get("photo")
	location := data.Get("location")

	if photo.IsSome() {
		if sizes, ok := photo.Some().([]gotgbot.PhotoSize); ok && len(sizes) > 0 {
			fileID := sizes[len(sizes)-1].FileId
			ctx.Photo(String(fileID).Prepend(tg.FileIDPrefix)).Caption("Your photo")
		}
	}

	if location.IsSome() {
		if loc, ok := location.Some().(*gotgbot.Location); ok {
			ctx.Bot.Std().SendLocation(ctx.EffectiveChat.Id, loc.Latitude, loc.Longitude, nil)
		}
	}

	summary := "🧾 Summary:\n"
	summary += "👤 Gender: " + gender.(string) + "\n"

	defer ctx.State.Jump("done")

	return ctx.Message(String(summary)).Send().Err()
}

func done(ctx *tg.Context) error {
	ctx.State.Clear()
	return ctx.Message("Thank you! I hope we can talk again some day.").Send().Err()
}

// Global command handlers
func handleSkip(ctx *tg.Context) error {
	state := ctx.State.Get().UnwrapOr("")
	switch state {
	case photo:
		ctx.State.Data().Set("photo", "skipped")
		ctx.State.Set(location)
		return ctx.Reply("Skipping photo. Now, share your location.").Send().Err()
	case location:
		ctx.State.Data().Set("location", "skipped")
		ctx.State.Set(bio)
		return ctx.Reply("Skipping location. Now, tell me something about yourself.").Send().Err()
	default:
		return ctx.Reply("Nothing to skip.").Send().Err()
	}
}

func handleCancel(ctx *tg.Context) error {
	ctx.State.Clear()
	return ctx.Reply("Bye! I hope we can talk again some day.").Send().Err()
}
