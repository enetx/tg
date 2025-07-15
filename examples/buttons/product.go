package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
)

type Product struct {
	ID    Int
	Title String
	Price Float
}

func main() {
	token := NewFile("../../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	products := Slice[Product]{
		{ID: 1, Title: "iPhone", Price: 999.99},
		{ID: 2, Title: "MacBook", Price: 1999.00},
		{ID: 3, Title: "AirPods", Price: 199.00},
	}

	b.Command("start", func(ctx *ctx.Context) error {
		kb := keyboard.Inline()

		for _, p := range products {
			kb.Row().Text(p.Title, Format("buy:{}", p.ID))
		}

		return ctx.Reply("Select a product:").Markup(kb).Send().Err()
	})

	// buy:<id>
	b.On.Callback.Prefix("buy:", func(ctx *ctx.Context) error {
		data := String(ctx.Callback.Data)

		id := data.StripPrefix("buy:").ToInt().UnwrapOrDefault()

		message := String("Product not found")

		product := products.Iter().Find(func(p Product) bool { return p.ID == id })
		if product.IsSome() {
			message = Format("You selected: {1.Title} — ${1.Price}", product.Some())
		}

		return ctx.Answer(message).Send().Err()
	})

	b.Polling().Start()
}
