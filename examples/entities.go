package main

import (
	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/entities"
)

func main() {
	token := g.NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()

	// Create bot instance
	b := bot.New(token).Build().Unwrap()

	// Demo command showing various entity types
	b.Command("entities", func(ctx *ctx.Context) error {
		text := g.String("Hello bold italic code")

		e := entities.New(text).
			Bold("bold").     // "bold"
			Italic("italic"). // "italic"
			Code("code")      // "code"

		return ctx.Reply(text).
			Entities(e).
			Send().Err()
	})

	// Demo with URL and spoiler
	b.Command("url", func(ctx *ctx.Context) error {
		text := g.String("Click here to visit Google")

		e := entities.New(text).
			URL("here", "https://google.com"). // "here" as a hyperlink
			Spoiler("Google")                  // "Google" as a spoiler

		return ctx.Reply(text).
			Entities(e).
			Send().Err()
	})

	// Demo with preformatted code block
	b.Command("code", func(ctx *ctx.Context) error {
		code := g.String(`func main() {
	    fmt.Println("Hello")
	}`)
		codeText := g.Format("Check this Go code:{}", code)

		e := entities.New(codeText).
			Pre(code, "go") // Go code with syntax highlighting

		return ctx.Reply(g.String(codeText)).
			Entities(e).
			Send().Err()
	})

	// Demo combining multiple entity types
	b.Command("mixed", func(ctx *ctx.Context) error {
		text := g.String("Bold italic underline strikethrough spoiler")

		e := entities.New(text).
			Bold("Bold").                   // "Bold"
			Italic("italic").               // "italic"
			Underline("underline").         // "underline"
			Strikethrough("strikethrough"). // "strikethrough"
			Spoiler("spoiler")              // "spoiler"

		return ctx.Reply(text).
			Entities(e).
			Send().Err()
	})

	// Demo with blockquotes
	b.Command("quote", func(ctx *ctx.Context) error {
		quoteText := `Regular text
This is a blockquote
This is expandable quote`

		e := entities.New(quoteText).
			Blockquote("This is a blockquote").              // Regular blockquote
			ExpandableBlockquote("This is expandable quote") // Expandable blockquote

		return ctx.Reply(g.String(quoteText)).
			Entities(e).
			Send().Err()
	})

	// Demo for sending a gift with formatted text
	b.Command("gift", func(ctx *ctx.Context) error {
		text := g.String("Happy Birthday!")

		e := entities.New(text).
			Bold("Happy").     // "Happy"
			Italic("Birthday") // "Birthday"

		return ctx.SendGift("gift_premium_1").
			Text(text).
			TextEntities(e).
			Send().Err()
	})

	// Start polling for updates
	b.Polling().Start()
}
