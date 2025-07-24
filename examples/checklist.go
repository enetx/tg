package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func main() {
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
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
			AddTask("Complete task 1").
			AddTask("Complete task 2").
			AddTask("Complete task 3").
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
			AddTask("Create repository").
			AddTask("Setup CI/CD").
			AddTask("Write documentation").
			AddTask("Add tests").
			AddTask("Deploy to production").
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
			AddTask("Milk").
			AddTask("Bread").
			AddTask("Eggs").
			AddTask("Apples").
			AddTask("Cheese").
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
			AddTask("New task 1").
			AddTask("New task 2").
			AddTask("New task 3").
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
			AddTask("@alice: Review code").
			AddTask("@bob: Update database").
			AddTask("@charlie: Test features").
			AddTask("@diana: Write documentation").
			OthersCanMarkTasksAsDone().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Team checklist created!").Send().Err()
	})
}
