package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg"
	"github.com/enetx/tg/keyboard"
)

func main() {
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	bot := tg.NewBot(token)

	// Middleware: restrict callbacks with "admin:" prefix to admins only
	bot.Use(func(ctx *tg.Context) error {
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
	bot.Use(func(ctx *tg.Context) error {
		Println("[MW] Update: {}", ctx.Update.GetType())
		return nil
	})

	bot.Command("start", func(ctx *tg.Context) error {
		return ctx.Reply("Buttons:").Markup(keyboard.Inline().
			Text("Admin only", "admin:secure").
			Row().
			Text("Available to everyone", "public")).
			Send().Err()
	})

	// Callback handler for admin-only button
	bot.On.Callback.Equal("admin:secure", func(ctx *tg.Context) error {
		return ctx.Answer("Welcome, admin!").Send().Err()
	})

	// Callback handler for public button
	bot.On.Callback.Equal("public", func(ctx *tg.Context) error {
		return ctx.Answer("This is accessible to everyone!").Send().Err()
	})

	// Simple message handler: replies "Hello" to any text message
	bot.On.Message.Text(func(ctx *tg.Context) error { return ctx.Message("Hello").Send().Err() })

	bot.Polling().DropPendingUpdates().Start()
}
