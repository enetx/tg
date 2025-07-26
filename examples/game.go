package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func main() {
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// Set game score
	b.Command("setscore", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 2 {
			return ctx.Reply("Usage: /setscore <user_id> <score>").Send().Err()
		}

		userID := args[0].ToInt().Unwrap().Int64()
		score := args[1].ToInt().Unwrap().Int64()

		result := ctx.SetGameScore(userID, score).Send()
		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Score " + args[1] + " set for user " + args[0] + "!").Send().Err()
	})

	// Set game score with force
	b.Command("forcescore", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 2 {
			return ctx.Reply("Usage: /forcescore <user_id> <score>").Send().Err()
		}

		userID := args[0].ToInt().Unwrap().Int64()
		score := args[1].ToInt().Unwrap().Int64()

		ctx.SetGameScore(userID, score).Force().Send()

		return ctx.Reply("Score " + args[1] + " force-set for user " + args[0] + "!").Send().Err()
	})

	// Set game score for specific message
	b.Command("setscoreformsg", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 3 {
			return ctx.Reply("Usage: /setscoreformsg <message_id> <user_id> <score>").Send().Err()
		}

		messageID := args[0].ToInt().Unwrap().Int64()
		userID := args[1].ToInt().Unwrap().Int64()
		score := args[2].ToInt().Unwrap().Int64()

		ctx.SetGameScore(userID, score).MessageID(messageID).Send()

		return ctx.Reply("Score updated for message " + args[0] + "!").Send().Err()
	})

	// Get game high scores
	b.Command("gethighscores", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /gethighscores <user_id>").Send().Err()
		}

		userID := args[0].ToInt().Unwrap().Int64()

		result := ctx.GetGameHighScores(userID).Send()
		if result.IsErr() {
			return result.Err()
		}

		scores := result.Ok()
		if scores.Empty() {
			return ctx.Reply("No high scores found for user " + args[0]).Send().Err()
		}

		response := "High Scores for user " + args[0] + ":\n"
		for i, score := range scores {
			response += Int(i+1).String() + ". " + Int(score.Score).String() + " points\n"
		}

		return ctx.Reply(response).Send().Err()
	})

	// Get high scores for specific message
	b.Command("getscoresformsg", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 2 {
			return ctx.Reply("Usage: /getscoresformsg <message_id> <user_id>").Send().Err()
		}

		messageID := args[0].ToInt().Unwrap().Int64()
		userID := args[1].ToInt().Unwrap().Int64()

		result := ctx.GetGameHighScores(userID).
			MessageID(messageID).
			Send()

		if result.IsErr() {
			return result.Err()
		}

		scores := result.Ok()
		response := "Scores for message " + args[0] + ":\n"
		for i, score := range scores {
			response += Int(i+1).String() + ". " + Int(score.Score).String() + " points\n"
		}

		return ctx.Reply(String(response)).Send().Err()
	})

	// Set inline game score
	b.Command("setinlinescore", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 3 {
			return ctx.Reply("Usage: /setinlinescore <inline_message_id> <user_id> <score>").Send().Err()
		}

		inlineMessageID := args[0]
		userID := args[1].ToInt().Unwrap().Int64()
		score := args[2].ToInt().Unwrap().Int64()

		result := ctx.SetGameScore(userID, score).
			InlineMessageID(inlineMessageID).
			DisableEditMessage().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Inline game score updated!").Send().Err()
	})

	// Create achievement checklist with verification
	b.Command("achievements", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 2 {
			return ctx.Reply("Usage: /achievements <business_connection_id> <user_id>").Send().Err()
		}

		businessConnectionID := args[0]
		userID := args[1].ToInt().Unwrap().Int64()

		// First verify the user
		verifyResult := ctx.VerifyUser(userID).Send()
		if verifyResult.IsErr() {
			return verifyResult.Err()
		}

		// Then create their achievement checklist
		checklistResult := ctx.SendChecklist(businessConnectionID, args[1]+" Achievements").
			Task("Complete first game").Add().
			Task("Reach 1000 points").Add().
			Task("Win 5 games in a row").Add().
			Task("Get verified status").Add().
			Task("Unlock special features").Add().
			OthersCanMarkTasksAsDone().
			Send()

		if checklistResult.IsErr() {
			return checklistResult.Err()
		}

		return ctx.Reply("Achievement checklist created for verified user " + args[1] + "!").Send().Err()
	})

	// Game leaderboard with verification
	b.Command("leaderboard", func(ctx *ctx.Context) error {
		// This is a demo - in real implementation you'd query multiple users
		return ctx.Reply("VERIFIED PLAYERS LEADERBOARD\n\n" +
			"1. @alice - 2500 points *\n" +
			"2. @bob - 2200 points\n" +
			"3. @charlie - 1800 points *\n" +
			"4. @diana - 1500 points *\n" +
			"5. @eve - 1200 points\n\n" +
			"* = Verified Player").
			Send().Err()
	})

	// Admin panel with all features
	b.Command("adminpanel", func(ctx *ctx.Context) error {
		return ctx.Reply("ADMIN PANEL 🔧\n\n" +
			"Games:\n" +
			"• /setscore <user> <score> - Set game score\n" +
			"• /gethighscores <user> - Get high scores\n\n" +
			"Combined:\n" +
			"• /achievements <business_id> <user> - Create achievement list\n" +
			"• /leaderboard - Show verified leaderboard").
			Send().Err()
	})

	b.Polling().AllowedUpdates().Start()
}
