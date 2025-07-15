package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
)

func main() {
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// Middleware: restrict callbacks with "admin:" prefix to admins only
	b.Use(func(ctx *ctx.Context) error {
		if ctx.Callback == nil || !String(ctx.Callback.Data).StartsWith("admin:") {
			return nil
		}

		admin := ctx.IsAdmin()
		if admin.IsErr() {
			return admin.Err()
		}

		if !admin.Ok() {
			return ctx.Answer("Access restricted to admins only!").Alert().Send().Err()
		}

		return nil
	})

	// Middleware: log each incoming update type
	b.Use(func(ctx *ctx.Context) error {
		Println("[MW] Update: {}", ctx.Update.GetType())
		return nil
	})

	b.Command("start", func(ctx *ctx.Context) error {
		return ctx.Reply("Buttons:").Markup(keyboard.Inline().
			Text("Admin only", "admin:secure").
			Row().
			Text("Available to everyone", "public")).
			Send().Err()
	})

	// Callback handler for admin-only button
	b.On.Callback.Equal("admin:secure", func(ctx *ctx.Context) error {
		return ctx.Answer("Welcome, admin!").Send().Err()
	})

	// Callback handler for public button
	b.On.Callback.Equal("public", func(ctx *ctx.Context) error {
		return ctx.Answer("This is accessible to everyone!").Send().Err()
	})

	// Simple message handler: replies "Hello" to any text message
	b.On.Message.Text(func(ctx *ctx.Context) error { return ctx.Message("Hello").Send().Err() })

	b.Polling().DropPendingUpdates().Start()
}
