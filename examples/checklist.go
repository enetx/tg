package main

import (
	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/entities"
)

func main() {
	token := g.NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// Send basic checklist
	b.Command("checklist", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 2 {
			return ctx.Reply("Usage: /checklist <business_connection_id> <title>").Send().Err()
		}

		businessConnectionID := args[0]
		title := args[1]

		result := ctx.SendChecklist(businessConnectionID, title).
			// Task with HTML formatting
			Task("<b>Bold task</b>").HTML().Add().
			// Task with Markdown formatting
			Task("*Italic task*").Markdown().Add().
			// Task with manually constructed entities
			Task("Underline and bold").
			Entities(
				entities.New("Underline and bold").
					Underline("Underline").
					Bold("bold")).
			Add().
			// Plain text tasks
			Task("Simple task 1").Add().
			Task("Simple task 2").Add().
			OthersCanMarkTasksAsDone().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Checklist created successfully!").Send().Err()
	})

	// Send checklist with some completed tasks
	b.Command("checklistpartial", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /checklistpartial <business_connection_id>").Send().Err()
		}

		businessConnectionID := args[0]

		result := ctx.SendChecklist(businessConnectionID, "Project Setup").
			Task("Create repository").Add().
			Task("Setup CI/CD").Add().
			Task("Write documentation").Add().
			Task("Add tests").Add().
			Task("Deploy to production").Add().
			OthersCanMarkTasksAsDone().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Project checklist created!").Send().Err()
	})

	// Send shopping list checklist
	b.Command("shopping", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /shopping <business_connection_id>").Send().Err()
		}

		businessConnectionID := args[0]

		result := ctx.SendChecklist(businessConnectionID, "Shopping List").
			Task("Milk").Add().
			Task("Bread").Add().
			Task("Eggs").Add().
			Task("Apples").Add().
			Task("Cheese").Add().
			OthersCanMarkTasksAsDone().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Shopping list created!").Send().Err()
	})

	// Edit checklist message
	b.Command("editchecklist", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 2 {
			return ctx.Reply("Usage: /editchecklist <business_connection_id> <message_id>").Send().Err()
		}

		businessConnectionID := args[0]
		messageID := args[1].ToInt().Unwrap().Int64()

		result := ctx.EditMessageChecklist(businessConnectionID).
			MessageID(messageID).
			Title("Updated Checklist").
			Task("New task 1").Add().
			Task("New task 2").Add().
			Task("New task 3").Add().
			OthersCanMarkTasksAsDone().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Checklist updated successfully!").Send().Err()
	})

	// Send team checklist
	b.Command("teamchecklist", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /teamchecklist <business_connection_id>").Send().Err()
		}

		businessConnectionID := args[0]

		result := ctx.SendChecklist(businessConnectionID, "Team Tasks").
			Task("@alice: Review code").Add().
			Task("@bob: Update database").Add().
			Task("@charlie: Test features").Add().
			Task("@diana: Write documentation").Add().
			OthersCanMarkTasksAsDone().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Team checklist created!").Send().Err()
	})
}
