// # More info on what deep linking actually is (read this first if it's unclear to you):
// # https://core.telegram.org/bots/features#deep-linking

package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/preview"
)

const (
	CheckThisOut   = "check-this-out"
	SoCool         = "so-cool"
	UsingEntities  = "using-entities-here"
	UsingKeyboard  = "using-keyboard-here"
	CallbackButton = "keyboard-callback-data"

	furl = "https://t.me/{.Bot.Raw.Username}?start={}"
)

func main() {
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	bot := tg.NewBot(token).Build().Unwrap()

	bot.Command("start", func(ctx *tg.Context) error {
		payload := ctx.Args().Last().UnwrapOrDefault()
		switch payload {
		case CheckThisOut:
			return deep1(ctx)
		case SoCool:
			return deep2(ctx)
		case UsingEntities:
			return deep3(ctx)
		case UsingKeyboard:
			return deep4(ctx)
		default:
			return start(ctx)
		}
	})

	bot.On.Callback.Equal(CallbackButton, func(ctx *tg.Context) error {
		url := Format(furl, ctx.Bot, UsingKeyboard)
		return ctx.Answer("").URL(url).Send().Err()
	})

	bot.Polling().DropPendingUpdates().Start()
}

func start(ctx *tg.Context) error {
	url := Format(furl, ctx, CheckThisOut)
	return ctx.Message(String("Feel free to tell your friends about it:\n\n" + url)).Send().Err()
}

func deep1(ctx *tg.Context) error {
	url := Format(furl, ctx, SoCool)
	return ctx.Message("Awesome, you just accessed hidden functionality! Now let's get back to the private chat.").
		Markup(keyboard.Inline().URL("Continue here!", url)).Send().Err()
}

func deep2(ctx *tg.Context) error {
	url := Format(furl, ctx, UsingEntities)
	msg := Format(`You can also mask the deep-linked URLs as links: <a href="{}">▶️ CLICK HERE</a>`, url)

	return ctx.Message(String(msg)).HTML().Preview(preview.New().Disable()).Send().Err()
}

func deep3(ctx *tg.Context) error {
	return ctx.Message("It is also possible to make deep-linking using InlineKeyboardButtons.").
		Markup(keyboard.Inline().Text("Like this!", CallbackButton)).Send().Err()
}

func deep4(ctx *tg.Context) error {
	msg := Format("Congratulations! This is as deep as it gets 👏🏻\n\nThe payload was: {}", ctx.Args())
	return ctx.Message(msg).Send().Err()
}
